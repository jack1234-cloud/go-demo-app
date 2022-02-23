package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//Plain URL
	response, err := http.Get("http://localhost/")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
