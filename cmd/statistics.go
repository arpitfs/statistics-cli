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

func statistics(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		isType, _ := cmd.Flags().GetString("type")

		if isType == CPU {
			cpuInfo, _ := cpu.Info()
			fmt.Println("CPU Information")
			for _, val := range cpuInfo {
				fmt.Println(val.ModelName)
			}

		}
	} else {
		cmd.Help()
	}

}
