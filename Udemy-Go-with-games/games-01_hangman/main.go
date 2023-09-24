package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Zombie",
	"Gopher",
	"Brasil",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	//targetWord = "United States of America"
	fmt.Println(targetWord)
	guessedLetters := bagOfLetters(targetWord)
	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue // keep the game as it is
		}

		letter := rune(input[0]) // converting to rune
		if validateGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	fmt.Print("Game over...")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win! You found out the word!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("invalid state. Game is over and there is no winner")
	}
}

// Derive a word we have to guess
func getRandomWord() string {

	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

// Read Input
func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

// Validating guessed letter
func validateGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)

}

func bagOfLetters(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true                 // convert targetWord to a rune
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true // get the last letter
	//guessedLetters[rune('s')] = true

	return guessedLetters
}

// Check if word is guessed
func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

// Check if Hangman is complete
func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9

}

// Game Over
func isGameOver(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) bool {
	return isWordGuessed(targetWord, guessedLetters) ||
		isHangmanComplete(hangmanState)
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println(getHangmanDrawing(hangmanState))

}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	result := ""
	for _, ch := range targetWord { // get the index (i) and the character (ch) inside the string. Bcoz wont need the index (i), replaces to (_)
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch) // Printf reads formatation e.g. %c $.3f
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := os.ReadFile(fmt.Sprintf("states/hm%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return (string(data))
}
