package enquiry

import (
	"kc/internal/fileutils"
	"log"
)

func GetEnquiryResults(filename string) (map[string]string, error) {
	var enquiries []string

	if err := fileutils.LoadJSON(filename, &enquiries); err != nil {
		log.Fatal(err)
	}

	// Create a slice to hold the answers
	answers := make(map[string]string, len(enquiries))

	// Iterate over the questions and prompt the user
	for _, question := range enquiries {
		ans, err := fileutils.Input(question, true)
		if err != nil {
			log.Fatalf("Error While Getting Input %v", err)
		}
		answers[question] = ans
	}

	return answers, nil
}
