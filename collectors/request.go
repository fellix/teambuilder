package collectors

import (
	"io/ioutil"
	"net/http"
)

func ApiRequest(url string) []byte {
	resp, err := http.Get(url)
	check(err)

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body
}
