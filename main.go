package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func containsWord(s []string, v string) bool {
	for _, w := range s {
		if w == v {
			return true
		}
	}
	return false
}

func containsRune(s []string, r rune) bool {
	for _, w := range s {
		if strings.ContainsRune(w, r) {
			return true
		}
	}
	return false
}

func validateWord(word string) (bool, error) {
	word = strings.TrimSpace(word)
	if word == "" {
		return false, nil
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return false, fmt.Errorf("validateWord: unable to resolve caller path")
	}
	words_file := string(filepath.Dir(currentFile)) + "/valid-words.txt"
	data, err := os.ReadFile(words_file)
	if err != nil {
		return false, fmt.Errorf("validateWord: read word list: %w", err)
	}

	words := strings.Split(strings.TrimSpace(strings.ToUpper(string(data))), "\n")

	if containsWord(words, word) {
		return true, nil
	} else {
		return false, nil
	}

}

func getRandomWord() string {

	_, currentFile, _, _ := runtime.Caller(0)
	words_file := string(filepath.Dir(currentFile)) + "/words.txt"
	data, _ := os.ReadFile(words_file)
	words := strings.Split(strings.TrimSpace(string(data)), "\n")

	return words[rand.IntN(len(words))]
}

func errorScreen(screen [6]string) {

	ClearScreen()
	printScreen(screen)

	fmt.Println("")
	fmt.Println("Error!")
	fmt.Println("")
	time.Sleep(500 * time.Millisecond)
}

func printScreen(screen [6]string) {
	//Print screen
	for _, line := range screen {
		fmt.Println(line)
	}
}

func main() {

	// solution := strings.ToUpper(getRandomWord())
	solution := "SLATE"

	screen := [6]string{
		"1. _ _ _ _ _",
		"2. _ _ _ _ _",
		"3. _ _ _ _ _",
		"4. _ _ _ _ _",
		"5. _ _ _ _ _",
		"6. _ _ _ _ _",
	}

	try := 0

	for true {

		ClearScreen()

		if try == 6 {
			fmt.Println("You loose!")
			fmt.Println("Solution was " + solution)
			break
		}

		printScreen(screen)

		var input string
		fmt.Println("")
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		validated, _ := validateWord(input)

		if validated == true {

			letters := strings.Split(input, "")
			for index, letter := range input {
				guessed_letter := string(letter)
				solution_letter := string(solution[index])

				if guessed_letter == solution_letter {
					letters[index] = "\033[1;32m" + guessed_letter + "\033[0m"
				} else if strings.ContainsRune(solution, letter) {
					letters[index] = "\033[1;33m" + guessed_letter + "\033[0m"
				}

			}

			screen[try] = strconv.Itoa(try+1) + ". " + strings.Join(letters, " ")

			if input == solution {
				ClearScreen()
				printScreen(screen)

				break
			}

			try = try + 1

		} else {

			errorScreen(screen)

		}
	}

}
