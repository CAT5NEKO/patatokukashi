package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var wordA string
	var wordB string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("単語Aを入力してください：")
	wordA, _ = reader.ReadString('\n')
	wordA = strings.TrimSpace(wordA)

	fmt.Print("単語Bを入力してください：")
	wordB, _ = reader.ReadString('\n')
	wordB = strings.TrimSpace(wordB)

	var additionalWords []string
	for {
		fmt.Print("追加する単語を入力してください (終了する場合はnya!): ")
		additionalWord, _ := reader.ReadString('\n')
		additionalWord = strings.TrimSpace(additionalWord)

		if additionalWord == "nya!" {
			break
		}

		additionalWords = append(additionalWords, additionalWord)
	}

	result := mergeWords(wordA, wordB, additionalWords...)
	fmt.Println("合体した単語:", result)
}

func mergeWords(wordA string, wordB string, additionalWords ...string) string {
	var mergedWord strings.Builder

	runesA := []rune(wordA)
	runesB := []rune(wordB)

	maxLen := len(runesA)
	if len(runesB) > maxLen {
		maxLen = len(runesB)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(runesA) {
			mergedWord.WriteRune(runesA[i])
		}
		if i < len(runesB) {
			mergedWord.WriteRune(runesB[i])
		}
		for _, word := range additionalWords {
			runes := []rune(word)
			if i < len(runes) {
				mergedWord.WriteRune(runes[i])
			}
		}
	}

	return mergedWord.String()
}
