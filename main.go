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
package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

type Config struct {
	CustomJokesPath string `yaml:"custom-jokes-dir"`
}

func main() {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configPath := path.Join(usr.HomeDir, "/.cobra-jokes.yaml")
	defaultConfigData := []byte("hello\ngo\n")

	f, err := os.Open(configPath)
	if err != nil { // Check if config file exists
		switch err.(type) {
		case *fs.PathError: // Create the config file
			err := ioutil.WriteFile(configPath, defaultConfigData, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created config file @ ~/.cobra-jokes.yaml")
			f, err = os.Open(configPath) // Open the new config file
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatal(err)
		}
	}

	// Read the bytes of the config file
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("3")
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	// cmd.Execute()
}
