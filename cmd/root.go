/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "statistics",
	Short: "A CLI to generate machine statisitics",
	Long: ` A CLI to get the cpu, memory and disk statisitcs"
	For example:
	statistics -t cpu
	statistics -t memory `,

	Run: statistics,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("type", "t", "", "Provide the type of statistics (cpu, memory, disk)")
}
