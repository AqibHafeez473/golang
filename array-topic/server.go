package main

import (
	wordcount "array-topic/golang-wordcount-array"
)

func main() {
	words, count := wordcount.WordCount()
	wordcount.PrintWords(words, count)
}
