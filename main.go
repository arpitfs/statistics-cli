package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
)

func getContainerMemoryStats(cli *client.Client, containerID string) (string, error) {
	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return "", err
	}
	defer stats.Body.Close()

	var memUsage string
	for {
		var stat types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&stat); err != nil {
			return "", err
		}
		memUsage = fmt.Sprintf("%.2f MB", float64(stat.MemoryStats.Usage)/1024/1024)
		break
	}

	return memUsage, nil
}

func getContainerCPUStats(cli *client.Client, containerID string) (string, error) {
	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return "", err
	}
	defer stats.Body.Close()

	var cpuUsage string
	for {
		var stat types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&stat); err != nil {
			return "", err
		}

		cpuUsage = fmt.Sprintf("%.2f %%", float64(stat.CPUStats.CPUUsage.TotalUsage)/float64(stat.CPUStats.SystemCPUUsage)*100)
		break
	}

	return cpuUsage, nil
}

func listRunningContainers(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All:     false,
		Filters: filters.NewArgs(filters.Arg("status", "running")),
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func displayMemoryStats(containers []types.Container, cli *client.Client) {
	header := fmt.Sprintf("%-20s %-20s %-15s", "CONTAINER", "ID", "MEM USAGE")
	color.Set(color.FgYellow, color.Bold)
	fmt.Println(header)
	color.Unset()

	for _, container := range containers {
		memUsage, err := getContainerMemoryStats(cli, container.ID)
		if err != nil {
			log.Printf("Error fetching stats for container %s: %v", container.ID, err)
			continue
		}

		memStats := fmt.Sprintf("%-20s %-20s %-15s", container.Names[0], container.ID[:12], memUsage)
		fmt.Println(memStats)
	}
}

func displayCPUStats(containers []types.Container, cli *client.Client) {
	header := fmt.Sprintf("%-20s %-20s %-15s", "CONTAINER", "ID", "CPU USAGE")
	color.Set(color.FgYellow, color.Bold)
	fmt.Println(header)
	color.Unset()

	for _, container := range containers {
		cpuUsage, err := getContainerCPUStats(cli, container.ID)
		if err != nil {
			log.Printf("Error fetching stats for container %s: %v", container.ID, err)
			continue
		}

		cpuStats := fmt.Sprintf("%-20s %-20s %-15s", container.Names[0], container.ID[:12], cpuUsage)
		fmt.Println(cpuStats)
	}
}

func displayStatsOptions() {
	fmt.Println("Usage: stats <memory|cpu>")
	fmt.Println("Options:")
	fmt.Println("  memory  - Show memory usage stats")
	fmt.Println("  cpu     - Show CPU usage stats")
	fmt.Println("  help    - Show this help message")
}

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		displayStatsOptions()
		return
	}

	command := os.Args[1]

	containers, err := listRunningContainers(cli)
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
		os.Exit(1)
	}

	switch strings.ToLower(command) {
	case "stats":
		if len(os.Args) == 2 {
			displayStatsOptions()
		} else {
			option := os.Args[2]
			switch option {
			case "memory":
				displayMemoryStats(containers, cli)
			case "cpu":
				displayCPUStats(containers, cli)
			default:
				fmt.Printf("Unknown option: %s\n", option)
				displayStatsOptions()
			}
		}

	case "stats memory":
		displayMemoryStats(containers, cli)

	case "stats cpu":
		displayCPUStats(containers, cli)

	default:
		fmt.Printf("Invalid command: %s\n", command)
		displayStatsOptions()
	}
}
