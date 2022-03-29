package hangman_web

import (
	"math/rand"
	"strconv"
)

func Finder(letter, randomword, wordtofind string, usedletter *[]string) (string, string) {
	for _, let := range *usedletter {
		if let == Nomalize(letter) {
			return wordtofind, "usedletter"
		}
	}
	if len([]rune(letter)) > 1 {
		if len([]rune(randomword)) == len([]rune(letter)) {
			if Nomalize(randomword) == Nomalize(letter) {
				return wordtofind, "wordgood"
			} else {
				return wordtofind, "wordwrong"
			}
		} else {
			return wordtofind, "wordinvalid"
		}
	}
	*usedletter = append(*usedletter, Nomalize(letter))
	if Nomalize(letter) >= "a" && Nomalize(letter) <= "z" {
		word := ""
		for k := 0; k < len([]rune(randomword)); k++ {
			if Nomalize(letter) == Nomalize(string([]rune(randomword)[k])) {
				word += string([]rune(randomword)[k])
			} else if []rune(wordtofind)[k] != '_' {
				word += string([]rune(wordtofind)[k])
			} else {
				word += "_"
			}
		}
		if word == wordtofind {
			return word, "fail"
		}
		return word, "good"
	} else {
		return wordtofind, "error"
	}
}

func GenerateWord(word string) string {
	var wordtofind []string
	for k := 0; k < len([]rune(word)); k++ {
		if word[k] == '-' {
			wordtofind = append(wordtofind, "-")
		} else {
			wordtofind = append(wordtofind, "_")
		}
	}
	for i := 0; i < (len([]rune(word))/2 - 1); i++ {
		tempr := rand.Intn(len([]rune(word)))
		if wordtofind[tempr] == "_" {
			wordtofind[tempr] = string([]rune(word)[tempr])
		} else {
			i--
		}
	}
	myrep := ""
	for _, letter := range wordtofind {
		myrep += letter
	}
	return myrep
}

func Message(state, randomword, choose string, try *int) string {
	if state == "fail" {
		*try--
		return "The letter " + Nomalize(choose) + " is not included in the word, you only have " + strconv.Itoa(*try) + " tries left"
	} else if state == "usedletter" {
		return "Letter already used"
	} else if state == "good" {
		return "The letter " + Nomalize(choose) + " is included in the word"
	} else if state == "wordinvalid" {
		return "The format is not valid, please enter a letter or a word with a good size"
	} else if state == "wordgood" {
		return "You found the word, you had " + strconv.Itoa(*try) + " tries left, the word was: " + randomword
	} else if state == "error" {
		return "The letter is invalid, please try again"
	} else if state == "wordwrong" {
		*try -= 2
		return "The word is not the right one, you have " + strconv.Itoa(*try) + " tries left"
	}
	return ""
}
