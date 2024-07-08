package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.reddit.com/r/dailyprogrammer/comments/98ufvz/20180820_challenge_366_easy_word_funnel_1/

func removeCharFromStr(str string, i int) string {
	return str[:i] + str[i+1:]
}

func funnel(str1 string, str2 string) bool {
	for i := range str1 {
		if removeCharFromStr(str1, i) == str2 {
			return true
		}
	}
	return false
}

var allWords = map[string]bool{}

func populateAllWords() {
	// read file with words => create set of all words
	var file, _ = os.Open("enable1.txt")
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		allWords[scanner.Text()] = true
	}
	file.Close()
}

func bonus(str string) []string {
	// remove single char from each position in str, add any that make valid words to set
	wordSet := map[string]bool{}
	for i := range str {
		strWithCharRemoved := removeCharFromStr(str, i)
		if _, ok := allWords[strWithCharRemoved]; ok {
			wordSet[strWithCharRemoved] = true
		}
	}

	// convert set of words to list
	wordList := []string{}
	for word := range wordSet {
		wordList = append(wordList, word)
	}
	return wordList
}

func bonus2() []string {
	wordList := []string{}
	for word := range allWords {
		if len(bonus(word)) == 5 {
			wordList = append(wordList, word)
		}
	}
	return wordList
}

func main() {
	// main challenge
	fmt.Println(funnel("leave", "eave"))
	fmt.Println(funnel("reset", "rest"))
	fmt.Println(funnel("dragoon", "dragon"))
	fmt.Println(funnel("eave", "leave"))
	fmt.Println(funnel("sleet", "lets"))
	fmt.Println(funnel("skiff", "ski"))

	// populate the set with all words
	populateAllWords()

	// bonus 1
	fmt.Println(bonus("dragoon"))
	fmt.Println(bonus("boats"))
	fmt.Println(bonus("affidavit"))

	// bonus 2
	fmt.Println(len(bonus2()))
}
