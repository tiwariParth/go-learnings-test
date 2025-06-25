package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var (
	dockerClient *client.Client
	rootCmd      *cobra.Command
	
	// Flags
	ageFlag     string
	allFlag     bool
	formatFlag  string
	verboseFlag bool
)

func init() {
	var err error
	
	// Initialize Docker client
	dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Error creating Docker client: %v\n", err)
		os.Exit(1)
	}
	
	// Setup root command
	rootCmd = &cobra.Command{
		Use:   "container-cleaner",
		Short: "A tool to manage and clean Docker containers",
		Long: `Container Cleaner is a CLI tool that helps you manage your Docker containers.
It allows you to list, filter, and remove containers that are no longer needed.`,
	}
	
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "Enable verbose output")
	
	// List command
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List containers",
		Run:   listContainers,
	}
	
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Show all containers (default shows just running)")
	listCmd.Flags().StringVarP(&ageFlag, "age", "", "", "Filter by age (e.g., 24h)")
	listCmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "Output format (table, json)")
	
	// Remove command
	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove containers",
		Run:   removeContainers,
	}
	
	removeCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Remove all stopped containers")
	removeCmd.Flags().StringVarP(&ageFlag, "age", "", "", "Remove containers older than specified age (e.g., 24h)")
	
	// Stats command
	statsCmd := &cobra.Command{
		Use:   "stats",
		Short: "Show container disk usage statistics",
		Run:   showStats,
	}
	
	// Add commands to root
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(statsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listContainers(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	
	// Set up filters
	filterArgs := filters.NewArgs()
	
	if ageFlag != "" {
		// TODO: Implement age filtering
	}
	
	// List containers
	containers, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{
		All:     allFlag,
		Filters: filterArgs,
	})
	
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
		os.Exit(1)
	}
	
	if len(containers) == 0 {
		fmt.Println("No containers found")
		return
	}
	
	// Display containers
	if formatFlag == "json" {
		// TODO: Implement JSON output
	} else {
		// Table output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "CONTAINER ID\tNAME\tSTATUS\tAGE\tSIZE")
		
		for _, container := range containers {
			// TODO: Calculate container age and size
			names := strings.TrimPrefix(container.Names[0], "/")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", 
				container.ID[:12], names, container.Status, "N/A", "N/A")
		}
		
		w.Flush()
	}
}

func removeContainers(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	
	// Set up filters
	filterArgs := filters.NewArgs()
	filterArgs.Add("status", "exited")
	
	if ageFlag != "" {
		// TODO: Implement age filtering
	}
	
	// List containers to be removed
	containers, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterArgs,
	})
	
	if err != nil {
		fmt.Printf("Error listing containers: %v\n", err)
		os.Exit(1)
	}
	
	if len(containers) == 0 {
		fmt.Println("No containers to remove")
		return
	}
	
	fmt.Printf("The following %d containers will be removed:\n", len(containers))
	for _, container := range containers {
		fmt.Printf(" - %s (%s)\n", container.ID[:12], strings.TrimPrefix(container.Names[0], "/"))
	}
	
	// Confirm removal
	fmt.Print("Are you sure you want to remove these containers? [y/N] ")
	var confirm string
	fmt.Scanln(&confirm)
	
	if strings.ToLower(confirm) != "y" {
		fmt.Println("Operation cancelled")
		return
	}
	
	// Remove containers
	for _, container := range containers {
		err := dockerClient.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
			Force: false,
		})
		
		if err != nil {
			fmt.Printf("Error removing container %s: %v\n", container.ID[:12], err)
		} else {
			fmt.Printf("Removed container %s\n", container.ID[:12])
		}
	}
}

func showStats(cmd *cobra.Command, args []string) {
	// TODO: Implement container statistics display
	fmt.Println("Stats feature not implemented yet")
}
