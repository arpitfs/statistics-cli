package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/spf13/cobra"
)

const (
	CPU    = "cpu"
	Memory = "memory"
	Disk   = "disk"
)

// var cpuCmd = &cobra.Command{
// 	Use:   "cpu",
// 	Short: "A CLI to generate machine statisitics",
// 	Long: ` A CLI to get the cpu, memory and disk statisitcs"
// 	For example:
// 	statistics cpu percent
// 	statistics memory `,

// 	Run: statistics,
// }

// func init() {
// 	rootCmd.AddCommand(cpuCmd)
// 	rootCmd.Flags().StringP("type", "t", "", "Provide the type of statisitcs")
// }

func statistics(cmd *cobra.Command, args []string) {
	isType, _ := cmd.Flags().GetString("type")
	if isType == "" {
		cmd.Help()
		return
	}
	switch isType {
	case CPU:
		cpuInfo, _ := cpu.Info()
		fmt.Println("CPU Information:")
		for _, val := range cpuInfo {
			fmt.Println(val.ModelName)
		}
	default:
		fmt.Println("Unknown type:", isType)
		cmd.Help()
	}

}
