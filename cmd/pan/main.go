package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/shanebarnes/panapi/snb/paol/dossier"
	"github.com/shanebarnes/panapi/snb/paol/search"
)

var (
	cookie string
	street string
)

func init() {
	flag.StringVar(&cookie, "cookie", "", "")
	flag.StringVar(&street, "street", "", "street name search terms (ex: 1 main st)")
	flag.Parse()
}

func main() {
	client := &http.Client{}

	if searchRequest, err := search.NewRequest(street, cookie); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create search request: %v\n", err)
	} else if searchResults, err := search.DoRequest(client, searchRequest); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to do search request: %v\n", err)
	} else {
		fmt.Fprintf(os.Stdout, "Search results: %s\n", prettyJson(searchResults))
		for _, group := range searchResults.ResultsByGroup {
			for _, result := range group.Results {
				if dossierRequest, err := dossier.NewRequest(result.Pan, cookie); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to create dossier request: %v\n", err)
				} else if dossierResults, err := dossier.DoRequest(client, dossierRequest); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to do dossier request: %v\n", err)
				} else {
					fmt.Fprintf(os.Stdout, "Dossier results: %s\n", prettyJson(dossierResults))
				}
			}
		}
	}
}

func prettyJson(f interface{}) string {
	if buf, err := json.MarshalIndent(f, "", "  "); err == nil {
		return string(buf)
	}
	return ""
}
