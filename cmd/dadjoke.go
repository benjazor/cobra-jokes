/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// dadCmd represents the dad command
var dadJokeCmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "Tells you a dad joke",
	Long:  `Tells you a dad joke from the API icanhazdadjoke.com  `,
	Run: func(cmd *cobra.Command, args []string) {
		getDadJoke()
	},
}

type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getDadJoke() {
	url := "https://icanhazdadjoke.com"
	responseBytes := getDadJokeData(url)
	dadJoke := DadJoke{}

	if err := json.Unmarshal(responseBytes, &dadJoke); err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dadJoke.Joke))
}

func getDadJokeData(apiURL string) []byte {
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
