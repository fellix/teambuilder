package collectors

import (
	".././entities"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	slice := strings.Split(ctx.NormalizedURL().String(), "/")
	var fileNumber = slice[len(slice)-1]
	//var list = make([]interface{}, 0)
	moves := map[string]entities.Move{}

	doc.Find("table#moves tbody tr").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a.ent-name").Text()
		typing := s.Find("a.type-icon").Text()
		category := s.Find("i.icon-move-cat").Text()
		var powerStr, accuracyStr, ppStr string
		s.Find("td.num").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				powerStr = s.Text()
			}

			if i == 1 {
				accuracyStr = s.Text()
			}

			if i == 2 {
				ppStr = s.Text()
			}
		})
		description := s.Find("td.long-text").Text()

		power, err := strconv.Atoi(powerStr)
		if err != nil {
			power = -1
		}
		accuracy, err := strconv.Atoi(accuracyStr)
		if err != nil {
			accuracy = -1
		}
		pp, err := strconv.Atoi(ppStr)
		if err != nil {
			pp = -1
		}

		move := entities.Move{Name: name, Typing: typing, Category: category, Power: power, Accuracy: accuracy, Pp: pp, Description: description}

		moves[strings.ToLower(name)] = move
	})

	j, err := json.Marshal(moves)
	check(err)

	f, err := os.OpenFile("./data/moves"+fileNumber+".json", os.O_WRONLY|os.O_CREATE, 0777)
	check(err)

	defer f.Close()

	n, err := f.WriteString(string(j))
	check(err)

	fmt.Printf("wrote %d bytes\n", n)

	f.Sync()

	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CollectMoves() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = false
	opts.MaxVisits = 1

	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("http://pokemondb.net/move/generation/1")
	c.Run("http://pokemondb.net/move/generation/2")
	c.Run("http://pokemondb.net/move/generation/3")
}
