package http

import (
	"io/ioutil"
	"net/http"
)

var HTTPGetBody = httpGetBody

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
