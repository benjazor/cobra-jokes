/*
Copyright © 2021 Benjamin Ludwig benjazor@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// dadCmd represents the dad command
var ojapiCmd = &cobra.Command{
	Use:   "ojapi",
	Short: "Tells you a joke from the official joke api",
	Long: `Tells you a random joke from the official joke api.
	You can find more informations about this api here:
	https://github.com/15Dkatz/official_joke_api`,
	Run: func(cmd *cobra.Command, args []string) {
		if Count < 1 {
			log.Fatal("Count can't be lower than 1")
		}
		for i := 0; i < Count; i++ {
			getJoke()
		}
	},
}

var Type string
var Count int

func init() {
	ojapiCmd.Flags().StringVarP(&Type, "type", "t", "random", "Sets the joke type between: random, general, programming")
	ojapiCmd.Flags().IntVarP(&Count, "count", "c", 1, "How many jokes are displayed, value as to be >= 1")
}

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func getJoke() {
	url := "https://official-joke-api.appspot.com/jokes/"
	switch strings.ToLower(Type) {
	case "random":
	case "general":
		url += "general/"
	case "programming":
		url += "programming/"
	default:
		log.Fatal(Type + " is not a valid joke type")
	}
	url += "random"

	responseBytes := getJokeData(url)
	if strings.ToLower(Type) == "random" {
		joke := Joke{}
		if err := json.Unmarshal(responseBytes, &joke); err != nil {
			log.Fatal(err)
		}
		fmt.Println(joke.Setup + " " + joke.Punchline)
	} else {
		joke := []Joke{}
		if err := json.Unmarshal(responseBytes, &joke); err != nil {
			log.Fatal(err)
		}
		fmt.Println(joke[0].Setup + " " + joke[0].Punchline)
	}
}

func getJokeData(apiURL string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		apiURL,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "cobra-jokes CLI (github.com/benjazor/cobra-jokes)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBytes
}
