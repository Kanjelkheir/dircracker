package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/akamensky/argparse"
	"github.com/kanjelkheir/dircracker/internal/utils"
)

func main() {
	// read the command line arguments
	parser := argparse.NewParser("DirCracker", "triverses through all the directories provided in the wordlist and return all the valid ones found")
	wordlist := parser.String("w", "wordlist", &argparse.Options{Required: true, Help: "The wordlist that would be checked against"})
	target := parser.String("t", "target", &argparse.Options{Required: true, Help: "The target that you want to check the wordlist against"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	// read the directories
	content, err := utils.ReadWordList(wordlist)
	if err != nil {
		log.Fatal(err)
	}

	// get the directories from the content as []string
	directories := utils.ReturnDirectories(content)

	// Triverse through the directories and make a request for every single directory
	var wg sync.WaitGroup
	for _, directory := range directories {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := *target + directory
			success, err := utils.CheckDirectory(url)
			if success {
				fmt.Printf("%s found!\n", url)
			} else {
				if err != nil {
					if errors.Is(err, utils.ErrNotFound) {
						fmt.Printf("%s not found!\n", url)
					} else {
						fmt.Printf("%s not found! - %v\n", url, err)
					}
				} else {
					// This case should ideally not be reached if CheckDirectory always returns an error when success is false.
					// However, if it does happen, it implies the directory was not found.
					fmt.Printf("%s not found!\n", url)
				}
			}
		}()
	}

	wg.Wait()

}
