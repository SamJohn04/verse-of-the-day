package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Details struct {
	Text      string `json:"text"`
	Reference string `json:"reference"`
	Version   string `json:"version"`
	VerseUrl  string `json:"verseurl"`
}

type Verse struct {
	Details Details `json:"details"`
	Notice  string  `json:"notice"`
}

type Result struct {
	Verse Verse `json:"verse"`
}

func main() {
	onlineVerseSource := "https://beta.ourmanna.com/api/v1/get?format=json&order=daily"
	response, err := http.Get(onlineVerseSource)
	if err != nil {
		fmt.Printf("error while getting verse: %v", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf(
			"error in staus code: expected %v but received %v",
			http.StatusOK,
			response.StatusCode,
		)
		os.Exit(1)
	}

	result := Result{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Printf("error while decoding the result: %v", err)
		os.Exit(1)
	}

	fmt.Printf(
		"%v, %v\n\n%v\n",
		result.Verse.Details.Reference,
		result.Verse.Details.Version,
		result.Verse.Details.Text,
	)
}
