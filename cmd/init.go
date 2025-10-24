package cmd

import (
	"fmt"
	"github.com/saivenkatram-git/zc/internal/git"
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize zen commit configuration",
	Long:  "Creates a zc.config.toml configuration file in your project root",
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	if !git.IsGitRepo() {
		fmt.Println("Error: Not a git repository")
		fmt.Println(" Run 'git init' first")
		os.Exit(1)
	}

	configPath := "zc.config.toml"

	// CHECK FOR EXISTING CONFIG..
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println("zc.config.toml already exists")
		fmt.Print("   Overwrite? (y/N): ")

		var response string
		fmt.Scanln(&response)

		if response != "y" && response != "Y" {
			fmt.Println("Initialization cancelled")
			os.Exit(0)
		}
	}

	// TOML - For Config
	configContent := `# Zen Commit Configuration
# https://github.com/yourusername/zc

[project]
name = "my-project"
description = "Project description"

[scopes]
# Define scopes specific to your project
# These will be suggested when creating commits
enabled = true
list = [
  "api",
  "ui",
  "auth",
  "database",
  "docs",
  "config",
  "tests"
]

[commit]
# Maximum length for commit subject line
max_length = 72

# Require scope for all commits
require_scope = false

# Add ticket/issue number to commits
# Example: feat(api): add endpoint [TICKET-123]
ticket_prefix = ""  # e.g., "JIRA-", "GH-", etc.

[breaking]
# Require description for breaking changes
require_description = true

[emoji]
# Enable emoji in commit messages
enabled = false

# Emoji mapping for commit types
[emoji.types]
feat = "✨"
fix = "🐛"
docs = "📝"
style = "💄"
refactor = "♻️"
perf = "⚡"
test = "✅"
build = "🏗️"
ci = "👷"
chore = "🔧"
revert = "⏪"

[hooks]
# Run commands before/after commit
# pre_commit = "npm test"
# post_commit = "echo 'Committed successfully!'"
`

	// CONFIG FILE CREATION
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		fmt.Printf("❌ Failed to create config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Created zc.config.toml")
	fmt.Println()
	fmt.Println("Configuration created with:")
	fmt.Println("   • Default scopes (api, ui, auth, etc.)")
	fmt.Println("   • Commit message length limit (72 chars)")
	fmt.Println("   • Optional emoji support")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("   1. Edit zc.config.toml to customize your settings")
	fmt.Println("   2. Run: git add zc.config.toml")
	fmt.Println("   3. Run: zc commit")
	fmt.Println()
	fmt.Println("Quick start:")
	fmt.Println("   zc commit    # Interactive commit")
	fmt.Println("   zc c         # Shorthand")
}
