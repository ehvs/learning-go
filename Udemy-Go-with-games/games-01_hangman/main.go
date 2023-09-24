package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

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
	hangmanState := 1
	printGameState(targetWord, guessedLetters, hangmanState)

	//guessedLetters['o'] = true
	//printGameState(targetWord, guessedLetters)

}
func getRandomWord() string {
	// Derive a word we have to guess
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func bagOfLetters(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true                 // convert targetWord to a rune
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true // get the last letter
	//guessedLetters[rune('s')] = true

	return guessedLetters
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

// Print game state
//  * word is: H _ _ _ m _ n
//  * print hangman state
// read user input
//  * validate (e.g. only letters)
// Determine if letter is correct or not
//  * if correct, update the guessed letter
//  * if incorrect, update the hangman state
// If word is guessed -> game over, you win
// If hagman is complete -> game over, you lose
// prompt > k
// guess incorrect!
// H a _ _ m a n
//
// --------
// |      |
// |      o
// |
// |
// -...
