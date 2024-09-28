package main

import "fmt"

/*
Pointers are a way to share memory with other parts of our program, which is useful for two major reasons:
1. When we have large amounts of data, making copies to pass between functions is very inefficient.
By passing the memory location of where the data is stored instead,
we can dramatically reduce the resource-footprint of our programs.
2. By passing pointers between functions, we can access and modify the single copy of the data directly,
meaning that any changes made by one function are immediately visible to other parts of the program
when the function ends.
*/

func NewVoteCounter(initialVotes int) *int {
	/*
	   Create a function NewVoteCounter that accepts the number of initial votes for a candidate and
	   returns a pointer referring to an int,initialized with the given number of initial votes.
	*/
	var counter *int        // the pointer
	counter = &initialVotes // Initialized : the pointer gets the value from 'initialVotes'
	return counter
}

func VoteCount(counter *int) int {
	// VoteCount extracts the number of votes from a counter.
	// return the number of votes in the counter.
	//If the counter is nil you should assume the counter has no votes:
	if counter == nil {
		return 0
	}
	var readCounter int
	readCounter = *counter //accessing pointer -- Deferencing
	return readCounter
}

func IncrementVoteCount(counter *int, increment int) {
	//Create a function IncrementVoteCount that will take a counter (*int) as an argument and a number of votes,
	//and will increment the counter by that number of votes.
	//Assume the pointer passed will never be nil.
	*counter = *counter + increment
}

type ElectionResult struct {
	Name  string
	Votes int
}

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	// Create a function NewElectionResult that receives the name of a candidate and their number of votes and
	//returns a new election result.
	return &ElectionResult{Name: candidateName, Votes: VoteCount(&votes)}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	//fmt.Printf("%v (%v)", result.Name, result.Votes)
	return fmt.Sprintf("%v (%v)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	println("Who is the Decrementing candidate: ", candidate, ", current votes: ", results[candidate])
	if validCandidate, exists := results[candidate]; exists {
		results[candidate] = validCandidate - 1
	}
}

func main() {
	//initialVotes := 2
	//NewVoteCounter(initialVotes)
	var votes int
	//votes = 3
	var voteCounter *int
	voteCounter = &votes
	VoteCount(voteCounter)
	//var nilVoteCounter *int // this will panic
	//VoteCount(nilVoteCounter) // this will panic
	votes = 3
	voteCounter = &votes
	IncrementVoteCount(voteCounter, 2)
	NewElectionResult("Peter", 3)
	var result *ElectionResult
	result = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}
	DisplayResult(result)
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}
	DecrementVotesOfCandidate(finalResults, "Mary")
	//finalResults["Mary"]
}
