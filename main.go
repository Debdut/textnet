package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/debdut/textnet/pkg/fetch"
)

func main() {
	site, err := fetch.GetSite(os.Args[1])
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(site)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
