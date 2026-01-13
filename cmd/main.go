package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/SamJohn04/verse-of-the-day/internal"
)

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

	result := internal.Result{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Printf("error while decoding the result: %v", err)
		os.Exit(1)
	}

	fmt.Printf(
		"%v\n%v\n\n%v\n",
		internal.Bold(result.Verse.Details.Reference),
		internal.CommonBibleVersionNames(result.Verse.Details.Version),
		result.Verse.Details.Text,
	)
}
