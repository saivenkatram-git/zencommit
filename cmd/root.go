package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zencommit",
	Short: "A easy to use CLI tool for conventional commits",
	Long: `
✨ Zen Commit - Conventional Commits Made Easy ✨

A fast, interactive CLI tool for creating perfect Git commits.

🚀 Quick Start:
  1. zencommit init        Initialize your project
  2. git add .             Stage your changes
  3. zencommit commit      Create a beautiful commit

📖 Learn more: https://github.com/saivenkatram-git/zencommit
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(versionCmd)
}
