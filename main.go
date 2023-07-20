package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	generateCmd.PersistentFlags().IntVar(&seed, "seed", seed, "Set seed value")
	generateCmd.MarkFlagRequired("seed")
	rootCmd.AddCommand(generateCmd)

	generateCmd.AddCommand(oneCmd)

	generateCmd.AddCommand(twoCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var seed int

var rootCmd = &cobra.Command{
	Use:   "cobra-1986",
	Run:   func(cmd *cobra.Command, args []string) {
		fmt.Printf("root called\n")
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("generate: seed %d\n", seed)
	},
}

var oneCmd = &cobra.Command{
	Use:   "one",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("generate one: seed %d\n", seed)
	},
}

var twoCmd = &cobra.Command{
	Use:   "two",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("generate two: seed %d\n", seed)
	},
}
