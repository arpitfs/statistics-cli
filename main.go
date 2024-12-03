package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(0, true)
	for _, cpui := range cpuInfo {
		fmt.Println(cpui.ModelName, cpui.Cores)
	}

	fmt.Println(cpuPercent)

}
