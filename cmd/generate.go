package cmd

import (
	"errors"
	"fmt"
	"kc/internal/fileutils"
	"kc/pkg/enquiry"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Define flags for the command
var (
	lengthFlag     int
	useSpecialFlag bool
	useNumFlag     bool
	interactive    bool
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate Random password",
	Long:    "Generate Random Password",
	Example: "kc generate -l 16 -s  -n \nkc generate -i",
	Run:     generate,
}

func init() {
	// Add flags for length, special characters, numbers, and interactive mode
	generateCmd.Flags().IntVarP(&lengthFlag, "length", "l", 0, "Length of the password")
	generateCmd.Flags().BoolVarP(&useSpecialFlag, "special", "s", false, "Include special characters in the password")
	generateCmd.Flags().BoolVarP(&useNumFlag, "numbers", "n", false, "Include numbers in the password")
	generateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Run in interactive mode")

	rootCmd.AddCommand(generateCmd)
}

func processInput(answers map[string]string) (int, bool, bool, error) {
	var length int
	var spl bool
	var num bool
	for key, value := range answers {
		if strings.Contains(key, "length") {
			var err error
			length, err = strconv.Atoi(value)
			if err != nil {
				return 0, false, false, errors.New("error while converting length: " + err.Error())
			}
		}
		if strings.Contains(key, "special") {
			var err error
			spl, err = strconv.ParseBool(value)
			if err != nil {
				return 0, false, false, errors.New("error while converting special: " + err.Error())
			}
		}
		if strings.Contains(key, "numbers") {
			var err error
			num, err = strconv.ParseBool(value)
			if err != nil {
				return 0, false, false, errors.New("error while converting numbers: " + err.Error())
			}
		}
	}
	return length, spl, num, nil
}

func getParamsfromUser() (int, bool, bool, error) {
	questionsFilePath := "config/enquiry/enquiry.json"
	answer, _ := enquiry.GetEnquiryResults(questionsFilePath)
	return processInput(answer)
}

func generate(cmd *cobra.Command, args []string) {
	var length int
	var useSpecial bool
	var useNum bool
	var err error

	if len(args) == 0 && !cmd.Flags().Changed("length") && !cmd.Flags().Changed("special") && !cmd.Flags().Changed("numbers") && !cmd.Flags().Changed("interactive") {
		cmd.Help()
		return
	}

	if interactive {
		// If interactive mode is enabled
		length, useSpecial, useNum, err = getParamsfromUser()
	} else {
		// If flags are provided
		length = lengthFlag
		useSpecial = useSpecialFlag
		useNum = useNumFlag
	}
	password, err := fileutils.GeneratePassword(length, true, useSpecial, useNum)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated password:", password)
}
