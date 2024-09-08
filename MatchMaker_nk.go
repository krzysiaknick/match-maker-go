package main

import (
	"fmt"
	"strconv"
)

const (
	TRUELOVE = 20
	FRIENDS  = 10
)

var questions = []string{
	"Country music is great",
	"Cats are better than dogs",
	"Windows Is Better Than Mac OS",
	"Baseball Is the Best Sport",
	"Porsche Makes The Best Cars",
}

var correctAnswers = []string{
	"4", "1", "5", "1", "5",
}

var userAnswers = make([]int, len(questions))

func main() {
	// Welcome message and instructions
	fmt.Println("Welcome to Matchmaker!")
	fmt.Println("Answer the following questions by entering a number between 1 and 5:")
	fmt.Println("1 = Strongly Disagree, 5 = Strongly Agree")

	// Get user answers with validation
	for i := 0; i < len(questions); i++ {
		fmt.Printf("Question %d: %s\n", i+1, questions[i])
		userAnswers[i] = validate()
	}

	// Output the user's answers and the correct answers
	fmt.Println("\nThank you! Let's see how compatible you are based on your answers.")
	for i, answer := range userAnswers {
		correctAnswer, _ := strconv.Atoi(correctAnswers[i]) // Convert correctAnswers[i] to an integer
		fmt.Printf("Q%d: %s -> Your answer: %d, Desired answer: %d\n", i+1, questions[i], answer, correctAnswer)
	}

	calculate()
	closeProgram()
}

func calculate() {
	var score = make([]int, len(questions))

	for i := 0; i < len(questions); i++ {
		correctAnswer, _ := strconv.Atoi(correctAnswers[i]) // Convert correctAnswers[i] to an integer
		score[i] = 5 - abs(userAnswers[i]-correctAnswer)

	}
	sum := 0
	for _, num := range score {
		sum += num

	}
	fmt.Println("Your score: ", percent(sum), "%")

	classification := ""

	if sum >= TRUELOVE {
		classification = "Its True Love!!!"
	} else if sum >= FRIENDS || sum <= TRUELOVE {
		classification = "We should just be friends."
	} else {
		classification = "Run Away!"
	}

	fmt.Println(classification)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func percent(sum int) int {
	return (sum * 100) / 25
}

func validate() int {
	var answer int
	_, err := fmt.Scan(&answer)

	for err != nil || answer < 1 || answer > 5 {
		fmt.Println("Invalid input! Please enter a number between 1 and 5.")
		_, err = fmt.Scan(&answer)
	}

	return answer
}

func closeProgram() {
	fmt.Println("\nPress Enter to close the program.")
	fmt.Scanln()
	fmt.Scanln()
}
