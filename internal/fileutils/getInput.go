package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ValidateYesOrNo prompts the user for a yes/no answer and validates the input.
func validateYesOrNo(prompt string) (bool, error) {
	validResponses := []string{"y", "n"}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		lowerInput := strings.ToLower(input)

		// Check if the input is a valid response
		for _, validResponse := range validResponses {
			if lowerInput == validResponse {
				return validResponse == "y", nil
			}
		}

		// Prompt again if the input is not valid
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

func validateInput(prompt string, scanner *bufio.Scanner) (string, error) {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input != "" {
			return input, nil
		}
		// Prompt again if the input is not valid
		fmt.Println("Required Input")
	}
}

func Input(question string, required bool) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if !required {
		return scanner.Text(), nil
	}

	if strings.HasSuffix(question, "(y/n):") {
		// For yes/no questions, use ValidateYesOrNo
		response, err := validateYesOrNo(question)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%t", response), nil
	} else {
		response, err := validateInput(question, scanner)
		if err != nil {
			return "", err
		}
		return response, nil
	}
}
