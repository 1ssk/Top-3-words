package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Print("Введите текст: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	wordCounts := make(map[string]int)
	for _, word := range strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", r)
	}) {
		wordCounts[word]++
	}

	type wordCount struct {
		word  string
		count int
	}

	wordCountsSlice := make([]wordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		wordCountsSlice = append(wordCountsSlice, wordCount{word, count})
	}

	sort.Slice(wordCountsSlice, func(i, j int) bool {
		return wordCountsSlice[i].count > wordCountsSlice[j].count
	})

	fmt.Println("Топ-3 самых часто встречающихся слов:")
	for i := 0; i < 3 && i < len(wordCountsSlice); i++ {
		fmt.Printf("%d. %s\n", i+1, wordCountsSlice[i].word)
	}
}
