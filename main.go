package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data []struct {
	UserID string `json:"userId"`
	Id     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	var r Data

	err = json.Unmarshal(body, &r)

	fmt.Println(r[0].Body)
}
