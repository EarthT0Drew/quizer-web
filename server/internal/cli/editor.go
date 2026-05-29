package cli

import "fmt"

func EditorMode() {
	for {
		fmt.Println("Select an editor mode:")
		var choice = Menu("Create a new quiz", "Edit an existing quiz", "Back to main menu")

		switch choice {
		case 0:
			NewQuiz()
		case 1:
			// Handle edit an existing quiz mode
		case 2:
			return
		}
	}
}

func NewQuiz() {
	var quizID int = 0 // TODO: Get the actual quiz ID after creating the quiz in the database

	var quizName string
	fmt.Println("Enter the name of the new quiz: ")
	fmt.Scan(&quizName)

	var quizDescription string
	fmt.Println("Enter a description for the quiz (optional): ")
	fmt.Scan(&quizDescription)

	// TODO: Create a new quiz with the provided name and description

	for {
		fmt.Println("Select an option:")
		var choice = Menu("Add a question", "Edit a question", "Finish quiz")

		switch choice {
		case 0:
			NewQuestion(quizID)
		case 1:
			// Handle edit a question mode
		case 2:
			fmt.Println("Returning to editor menu...")
			return
		}
	}
}

func NewQuestion(quizID int) {
	var questionID int = 0 // TODO: Get the actual question ID after creating the question in the database

	var prompt string
	fmt.Println("Enter the question: ")
	fmt.Scan(&prompt)

	var typeChoice int
	fmt.Println("Select the type of question:")
	typeChoice = Menu("Multiple choice", "True/False", "Definition")

	var questionType string
	_ = questionType

	switch typeChoice {
	case 0:
		var finished bool = false
		for !finished {
			fmt.Println("Select an option:")
			var option int = Menu("Create new MCQ answer", "Finish question")
			switch option {
			case 0:
				NewMCQ(questionID, quizID)
			case 1:
				finished = true
			}
		}
		questionType = "mcq"

	case 1:
		var answerOption int = Menu("True", "False")
		switch answerOption {
		case 0:
			questionType = "true"
		case 1:
			questionType = "false"
		}
		questionType = "true_false"

	case 2:
		var answer string
		fmt.Println("Type the definition: ")
		fmt.Scan(&answer)

		questionType = "answer"
	}

	// TODO: Create a new question
}

func NewMCQ(questionID int, quizID int) {
	var answer string
	fmt.Println("Type the answer: ")
	fmt.Scan(&answer)

	fmt.Println("Is this answer correct?")
	var isCorrectOption int = Menu("Correct", "Incorrect")

	var isCorrect bool
	_ = isCorrect
	switch isCorrectOption {
	case 0:
		isCorrect = true
	case 1:
		isCorrect = false
	}

	// TODO: Create a new MCQ answer
}
