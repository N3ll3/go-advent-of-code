package cmd

import (
	"errors"
	"go-advent-of-code/inputs"
	"go-advent-of-code/puzzles"

	"github.com/spf13/cobra"
)

var puzzleCmd = &cobra.Command{
	Use: "puzzle <date>",
	Short: "Lance le puzzle d'un <jour>",
	Long: ` "Lance le puzzle d'un <jour> qui represente le path dans le projet Exemple :01`,
	Run: func(cmd *cobra.Command, args []string) {
		var test string
		var part string
		var day string
		var path string
		var inputPath string
		var input []string

		day = args[0]
		test, err := cmd.Flags().GetString("test")
		
		if err != nil {
			error := errors.New("Lancer le test ?")
			panic(error)
		}

		part, err = cmd.Flags().GetString("part")
		if err != nil {
			error := errors.New("Quelle partie du puzzle souhaitez-vous lancer ?")
			panic(error)
		}

		if test == "true" {
			inputPath = "input-test.txt"
		} else {
			inputPath = "input.txt"
		}
		
		path = day + "/"+ part +"/"+inputPath

		input = inputs.Get(path)

		puzzle := puzzles.Puzzle{
				Name: day,
				Part: part,
				Input: input,
		}

		puzzles.resolve()
		

		
	},
}

func init() {
	rootCmd.AddCommand(puzzleCmd)
}