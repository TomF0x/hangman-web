package hangman_web

import (
	"bufio"
	"os"
)

func Loadressource(filename string, size int) []string {
	file, _ := os.OpenFile(filename, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(file)
	cont := []string{""}
	count := 0
	index := 0
	for scanner.Scan() {
		count++
		if size == 7 || size == 9 {
			cont[index] += scanner.Text() + "\n"
		} else {
			cont[index] += scanner.Text()
		}
		if count == size {
			count = 0
			index++
			cont = append(cont, "")
		}
	}
	file.Close()
	return cont[:len(cont)-1]
}
