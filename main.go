package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	resp, err := http.Get("")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body))
}