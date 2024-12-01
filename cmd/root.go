package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:  "solve",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dayNum, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		partNum, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		dirName := fmt.Sprintf("day%d", dayNum)
		inputPath := filepath.Join("days", dirName, "input.txt")
		res, err := Solve(dayNum, partNum, inputPath)
		if err != nil {
			return err
		}
		fmt.Printf("Day [%s], Part [%s] solution: [%s]", args[0], args[1], res)

		return nil
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("Fatal error")
	}
}
