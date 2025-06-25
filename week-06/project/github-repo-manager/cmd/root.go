package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gh-manager",
	Short: "GitHub Repository Manager",
	Long: `A CLI tool to manage GitHub repositories.
This application allows you to list, create, inspect, and clone
GitHub repositories from the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, print help
		cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define persistent flags and configuration settings
	// These flags are available to all subcommands
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	
	// Add subcommands
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(infoCmd)
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(cloneCmd)
	RootCmd.AddCommand(authCmd)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your GitHub repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing GitHub repositories...")
		// TODO: Implement repository listing
	},
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [username/repo]",
	Short: "Show information about a GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]
		fmt.Printf("Showing info for repository: %s\n", repo)
		// TODO: Implement repository info
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [repo-name]",
	Short: "Create a new GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		desc, _ := cmd.Flags().GetString("description")
		private, _ := cmd.Flags().GetBool("private")
		
		fmt.Printf("Creating repository: %s (private: %v)\n", repoName, private)
		fmt.Printf("Description: %s\n", desc)
		// TODO: Implement repository creation
	},
}

// init function to set up create command flags
func init() {
	createCmd.Flags().StringP("description", "d", "", "Repository description")
	createCmd.Flags().BoolP("private", "p", false, "Make repository private")
}

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone [username/repo]",
	Short: "Clone a GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]
		dir, _ := cmd.Flags().GetString("directory")
		
		if dir == "" {
			// Default to repo name as the directory
			dir = repo[len(repo)-len(repo):]
		}
		
		fmt.Printf("Cloning repository %s into %s\n", repo, dir)
		// TODO: Implement repository cloning
	},
}

// init function to set up clone command flags
func init() {
	cloneCmd.Flags().StringP("directory", "d", "", "Directory to clone into")
}

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth --token [token]",
	Short: "Set your GitHub API token",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		if token == "" {
			fmt.Println("Error: GitHub token is required")
			cmd.Help()
			return
		}
		
		fmt.Println("Setting GitHub API token...")
		// TODO: Implement token storage
	},
}

// init function to set up auth command flags
func init() {
	authCmd.Flags().StringP("token", "t", "", "GitHub API token")
	authCmd.MarkFlagRequired("token")
}
