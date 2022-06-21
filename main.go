package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/debdut/textnet/pkg/fetch"
)

func main() {
	fmt.Println(os.Args[1])
	links := fetch.GetSearchLinks(os.Args[1], 1)

	json, err := json.Marshal(links)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
	// fmt.Println()
}
