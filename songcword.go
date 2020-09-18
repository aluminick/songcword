// TODO: testing
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var uniqueWords = make(map[string]int) // uniqueWords
var origWordCount = 0

// remove special characters on both ends
func removeSpecialCh(word string) string {
	return strings.ToLower(strings.Trim(word, "()\"', \n"))
}

// checks if word is already saved
func exists(word string) bool {
	_, ok := uniqueWords[word]
	if ok {
		return true
	}
	return false
}

// save word to uniqueWords
func insertWord(word string) {
	word = removeSpecialCh(word) // remove special characters
	if len(word) > 0 {           // if not empty
		origWordCount = origWordCount + 1 // increment for original word count
		if !exists(word) {                // not exists then insert to uniqueWords
			uniqueWords[word] = 1
		} else {
			uniqueWords[word] = uniqueWords[word] + 1
		}
	}
}

func main() {

	log.SetFlags(0)
	// Require file input
	inputFile := flag.String("input-file", "", "path to file input")
	format := flag.String("format", "txt", "format format: txt, json")

	flag.Parse()
	if *inputFile == "" {
		log.Println("Missing input file. \nExample: songcword -input-file=/path/to/lyrics.txt")
		return
	}

	dat, err := ioutil.ReadFile(*inputFile) //read file
	if err != nil {
		log.Println(err)
		return
	}

	strData := string(dat) //convert to string
	word := ""

	for _, ch := range strData {
		strCh := string(ch)                //convert to string
		word = word + strCh                // append to word
		if strCh == "\n" || strCh == " " { //per new word
			insertWord(word) // insert
			word = ""        //reset word
		}
	}

	uniqueWordsCount := len(uniqueWords)
	if *format == "json" {
		jsonUW, err := json.Marshal(uniqueWords)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(jsonUW))

	} else {
		for k, c := range uniqueWords {
			fmt.Println(k, "-", c)
		}
	}

	log.Println("\nOriginal Word Count: ", origWordCount)
	log.Printf("Unique Word Count: %v\n", len(uniqueWords))
	log.Printf("less %v%%\n", (100 - (float32(uniqueWordsCount)/float32(origWordCount))*100))
}
