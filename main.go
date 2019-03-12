package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"word-frequecy-counter/wordcounter"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)

	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	//Clean up input text
	matches := wordcounter.CleanUpText(bs)

	//Get words with their frequencies
	var wordCounts = wordcounter.GetWordFrequencies(matches)

	//Sort the word in desc order of frequencies
	wordCounts = wordcounter.SortWordsCountDesc(wordCounts)

	fmt.Println("Word  Frequency")
	fmt.Println("====  =========")
	for rank := 1; rank <= 20; rank++ {
		word := wordCounts[rank-1].Word
		freq := wordCounts[rank-1].Count
		fmt.Printf("%-4s    %5d\n",  word, freq)
	}
}
