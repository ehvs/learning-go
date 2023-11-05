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

// Declare a variable globaly outside main function

var inputReader = bufio.NewReader(os.Stdin)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	guessedLetters := initializedGuess(targetWord)
	hangmanState := 0

	fmt.Println(targetWord)

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Too long input]")
			continue // to keep the game state as it is. Do not reset anything.
		}

		letter := rune(input[0])
		//fmt.Printf("\n ...You inserted: %c \n", letter)
		for l := range guessedLetters {
			if l == letter {
				fmt.Printf("Letter %c, already used \n", l)
			}
		}
		if isCorrectGuess(targetWord, letter) {
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

// Function structure - nome (entrada) saida { codigo }
func getRandomWord() string {
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func initializedGuess(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

// Em um loop, para cada caractere da palavra em 'targetWord', se esse caractere nao existe no map 'guessedLetters'
func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, character := range targetWord {
		if !guessedLetters[unicode.ToLower(character)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

// RUNE eh uma letra/caractere
func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, character := range targetWord {
		if character == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(character)] == true {
			result += fmt.Sprintf("%c", character) //Printf to format  || Sprintf format anrd returns as string
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

func readInput() string {
	fmt.Printf("> ")
	// r := bufio.NewReader(os.Stdin) -- instead, declare globaly for recycling purposes
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	// checking if a letter is inside a word 'containsRune'
	return strings.ContainsRune(targetWord, letter)
}

/* Extra challenge
- if a letter was already guessed, print "letter already used" -- DONE
- implement 'hint', print a letter that was not used yet
  - limit the user to one hint only
*/
