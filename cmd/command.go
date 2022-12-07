package cmd

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/nomnemonic/nomnemonic"

	"github.com/spf13/cobra"
)

//go:embed data/english.txt
var _words []byte

var (
	mnemonicer nomnemonic.Mnemonicer

	identifier string
	password   string
	passcode   string
	size       int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nomnemonic-cli",
	Short: "Deterministic mnemonic generator",
	Long: "Deterministic mnemonic generator client that uses 3 inputs and " +
		"cryptographic hash function to generate the words.",
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Library and algorithm versions",
	Long:  "Displays the library and algorithm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Library v" + nomnemonic.Version)
		fmt.Println("Algorithm v" + nomnemonic.VersionAlgorithm)
	},
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate deterministic mnemonic words",
	Long: "Generate deterministic mnemonic words based on 3 inputs an " +
		"'identifier', a 'password' and a 'passcode'.\n\n" +
		"WARNING: You have to use the same 3 inputs to regenerate the same " +
		"mnemonic words. The cli application auto trims the spaces from left " +
		"and right. Other than that even a character change results with a " +
		"different set of words. As a best practice, run the app twice to " +
		"validate your inputs and output mnemonic words.",
	Run: func(cmd *cobra.Command, args []string) {
		generateWords()
	},
}

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Generate deterministic mnemonic words in interactive mode",
	Long:  generateCmd.Long,
	Run: func(cmd *cobra.Command, args []string) {
		// Sentence size

		sizes := []int{24, 21, 18, 15, 12}
		promptSize := promptui.Select{
			Label: "Select sentence size",
			Items: sizes,
		}

		cursor, _, err := promptSize.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		size = sizes[cursor]

		// Identifier

		validateLen := func(length int) func(input string) error {
			return func(input string) error {
				input = strings.Trim(input, " ")
				if len(input) < length {
					return fmt.Errorf("Must be at least %d chars", length)
				}
				return nil
			}
		}

		promptIdentifier := promptui.Prompt{
			Label:    "Identifier (like username, email, phone, etc...)",
			Validate: validateLen(2),
		}

		identifier, err = promptIdentifier.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Password

		promptPassword := promptui.Prompt{
			Label:       "Password",
			Validate:    validateLen(12),
			HideEntered: true,
			Mask:        '*',
		}

		password, err = promptPassword.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		validateRepeat := func(entry string) func(input string) error {
			return func(input string) error {
				input = strings.Trim(input, " ")
				if input != entry {
					return errors.New("Does not match the previous entry")
				}
				return nil
			}
		}

		promptPasswordRepeat := promptui.Prompt{
			Label:       "Password (repeat)",
			Validate:    validateRepeat(password),
			HideEntered: true,
			Mask:        '*',
		}

		_, err = promptPasswordRepeat.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Passcode

		validateExactLen := func(length int) func(input string) error {
			return func(input string) error {
				input = strings.Trim(input, " ")
				if len(input) != length {
					return fmt.Errorf("Must be exactly %d digits", length)
				}
				return nil
			}
		}

		promptPasscode := promptui.Prompt{
			Label:       "Passcode",
			Validate:    validateExactLen(6),
			HideEntered: true,
			Mask:        '#',
		}

		passcode, err = promptPasscode.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		promptPasscodeRepeat := promptui.Prompt{
			Label:       "Passcode (repeat)",
			Validate:    validateRepeat(passcode),
			HideEntered: true,
			Mask:        '#',
		}

		_, err = promptPasscodeRepeat.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		generateWords()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initCommands()
	initWords()
	initFlags()
}

func initCommands() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(versionCmd)
}

func initWords() {
	words := strings.Split(string(_words), "\n")
	var err error
	mnemonicer, err = nomnemonic.New(words)
	if err != nil {
		panic(err)
	}
}

func initFlags() {
	generateCmd.Flags().
		IntVarP(&size, "size", "s", 24, "The number of words needs to be generated, valid values are 12, 15, 18, 21, 24")
	generateCmd.Flags().
		StringVarP(&identifier, "identifier", "i", "", "Any identifier like username, email, phone number, etc.. (min 2 chars)")
	generateCmd.Flags().
		StringVarP(&password, "password", "p", "", "A password (min 8 chars)")
	generateCmd.Flags().
		StringVarP(&passcode, "passcode", "c", "", "6 digit numeric code")

	generateCmd.MarkFlagRequired("identifier")
	generateCmd.MarkFlagRequired("password")
	generateCmd.MarkFlagRequired("passcode")
}

func generateWords() {
	sentence, err := mnemonicer.Generate(
		strings.Trim(identifier, " "),
		strings.Trim(password, " "),
		strings.Trim(passcode, " "),
		size,
	)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return
	}

	fmt.Printf("Generated %d words deterministic mnemonic sentence\n\n", len(sentence))
	fmt.Println("Sentence")
	fmt.Println("---------------------------------------------------------")
	fmt.Println(strings.Join(sentence, " "))
	fmt.Println("")
	for i, w := range sentence {
		fmt.Printf("%2d. %s\n", i+1, w)
	}
}
