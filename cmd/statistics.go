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

var getStatisitcs = &cobra.Command{
	Use:   "statistics",
	Short: "A CLI to generate machine statisitics",
	Long: ` A CLI to get the cpu, memory and disk statisitcs"
	For example:
	statistics cpu percent
	statistics memory `,
	Run: statistics,
}

func init() {
	rootCmd.AddCommand(getStatisitcs)

	getStatisitcs.Flags().StringP("type", "t", "cpu", "Provide the type of statisitcs")
}

func statistics(cmd *cobra.Command, args []string) {
	isType, _ := cmd.Flags().GetString("type")

	if isType == CPU {
		cpuInfo, _ := cpu.Info()
		fmt.Println("CPU Information")
		for _, val := range cpuInfo {
			fmt.Println(val.ModelName)
		}

	}
}
