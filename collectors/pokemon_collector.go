package collectors

import (
	".././entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func CollectPokemons() {
	dat, err := ioutil.ReadFile("./data/pokemon.json")
	check(err)

	var availablePokemon []int

	json.Unmarshal(dat, &availablePokemon)

	moveData, err := ioutil.ReadFile("./data/moves.json")
	check(err)

	moves := map[string]entities.Move{}
	json.Unmarshal(moveData, &moves)

	waitGroup := &sync.WaitGroup{}

	for index, number := range availablePokemon {
		waitGroup.Add(1)
		go collectPokemon(number, moves, waitGroup)

		if index%3 == 0 {
			time.Sleep(1000 * time.Millisecond)
		}
	}

	waitGroup.Wait()
}

func collectPokemon(index int, availableMoves map[string]entities.Move, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	body := ApiRequest("http://pokeapi.co/api/v1/pokemon/" + strconv.Itoa(index))

	var pokemon entities.Pokemon
	json.Unmarshal(body, &pokemon)

	storedMoves := pokemon.Moves
	pokemon.Moves = fixMoveset(availableMoves, storedMoves)

	fileName := strings.ToLower(pokemon.Name) + ".json"

	j, err := json.Marshal(pokemon)
	check(err)

	f, err := os.OpenFile("./data/pokemon/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	check(err)

	defer f.Close()

	n, err := f.WriteString(string(j))
	check(err)

	fmt.Printf("%v: wrote  %d bytes\n", pokemon.Name, n)

	f.Sync()
}

func fixMoveset(availableMoves map[string]entities.Move, storedMoveset []entities.Move) []entities.Move {
	var validMoves []entities.Move

	for _, move := range storedMoveset {
		name := strings.ToLower(move.Name)
		foundMove, keyFound := availableMoves[name]
		if !keyFound {
			name := strings.Replace(name, "-", " ", -1)
			foundMove, keyFound := availableMoves[name]

			if keyFound {
				validMoves = append(validMoves, foundMove)
			}
		}

		if keyFound {
			validMoves = append(validMoves, foundMove)
		}
	}

	return validMoves
}
