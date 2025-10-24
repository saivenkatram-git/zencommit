package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/saivenkatram-git/zc/internal/config"
	"github.com/saivenkatram-git/zc/internal/git"
	"github.com/saivenkatram-git/zc/internal/types"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Create a zen commit",
	Long:    "Create a zen commit with interactive prompts following the Conventional Commits specification",
	Aliases: []string{"c"},
	Run:     runCommit,
}

func runCommit(cmd *cobra.Command, args []string) {
	if !git.IsGitRepo() {
		fmt.Println("Error: Not a git repository")
		fmt.Println("   Run 'git init' first")
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		fmt.Printf("   Run 'zc init' first")
		os.Exit(1)
	}

	if !git.HasStagedChanges() {
		fmt.Println("No staged changes found")
		fmt.Println("   Stage files first with: git add <files>")
		fmt.Println()
		fmt.Println("Tip: Use 'git status' to see unstaged changes")
		os.Exit(1)
	}

	// Show header
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║         ZC - Zen Commit Creator                ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println()

	// COMMIT TYPE
	typePrompt := promptui.Select{
		Label: "Select commit type",
		Items: types.CommitTypes,
		Size:  11,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "▸ {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	idx, _, err := typePrompt.Run()
	if err != nil {
		fmt.Printf("\n Selection cancelled\n")
		os.Exit(0)
	}

	commitType := types.GetCommitType(idx)
	fmt.Println()

	// SCOPES
	var scope string

	if cfg.Scopes.Enabled && len(cfg.GetScopes()) > 0 {
		// Show scopes from config
		scopeItems := append([]string{"(none) - No scope"}, cfg.GetScopes()...)

		scopePrompt := promptui.Select{
			Label: "Select scope",
			Items: scopeItems,
			Size:  10,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "▸ {{ . | cyan }}",
				Inactive: "  {{ . }}",
				Selected: "✓ {{ . | green }}",
			},
		}

		_, selectedScope, err := scopePrompt.Run()
		if err != nil {
			fmt.Printf("\nSelection cancelled\n")
			os.Exit(0)
		}

		if !strings.HasPrefix(selectedScope, "(none)") {
			scope = selectedScope
		}
	} else {
		// Manual scope entry
		scopePrompt := promptui.Prompt{
			Label:   "Scope (optional, press Enter to skip)",
			Default: "",
		}

		scope, err = scopePrompt.Run()
		if err != nil {
			fmt.Printf("\n❌ Input cancelled\n")
			os.Exit(0)
		}
	}

	// Check if scope is required
	if cfg.Commit.RequireScope && scope == "" {
		fmt.Println("❌ Error: Scope is required by project configuration")
		fmt.Println("   Edit zc.config.toml to change this setting")
		os.Exit(1)
	}

	fmt.Println()

	// SHORT DESCRIPTION
	maxLength := cfg.Commit.MaxLength
	descPrompt := promptui.Prompt{
		Label: fmt.Sprintf("Short description (max %d characters)", maxLength),
		Validate: func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("description is required")
			}
			if len(input) > maxLength {
				return fmt.Errorf("description too long (%d/%d characters)", len(input), maxLength)
			}
			return nil
		},
	}

	description, err := descPrompt.Run()
	if err != nil {
		fmt.Printf("\n❌ Input cancelled\n")
		os.Exit(0)
	}

	// Add emoji if enabled
	emoji := cfg.GetEmoji(commitType)
	if emoji != "" {
		description = emoji + description
	}

	fmt.Println()

	// LONGER BODY (OPTIONAL)
	bodyPrompt := promptui.Prompt{
		Label:   "Longer description (optional, press Enter to skip)",
		Default: "",
	}

	body, err := bodyPrompt.Run()
	if err != nil {
		fmt.Printf("\n❌ Input cancelled\n")
		os.Exit(0)
	}

	fmt.Println()

	// BREAKING CHANGES
	breakingPrompt := promptui.Prompt{
		Label:     "Are there any breaking changes?",
		IsConfirm: true,
	}

	breakingResult, err := breakingPrompt.Run()
	breaking := (err == nil && (breakingResult == "y" || breakingResult == "Y"))

	// IF BREAKING CHANGES IS TRUE, THEN DESCRIPTION
	breakingDesc := ""
	if breaking {
		if cfg.Breaking.RequireDescription {
			breakingDescPrompt := promptui.Prompt{
				Label: "Describe the breaking change",
				Validate: func(input string) error {
					if len(input) == 0 {
						return fmt.Errorf("breaking change description is required")
					}
					return nil
				},
			}

			breakingDesc, err = breakingDescPrompt.Run()
			if err != nil {
				fmt.Printf("\n❌ Input cancelled\n")
				os.Exit(0)
			}
		} else {
			breakingDescPrompt := promptui.Prompt{
				Label:   "Describe the breaking change (optional)",
				Default: "",
			}

			breakingDesc, _ = breakingDescPrompt.Run()
		}
		fmt.Println()
	}

	// TICKER DETAILS (JIRA, ETC) (OPTIONAL)
	ticket := ""
	if cfg.Commit.TicketPrefix != "" {
		ticketPrompt := promptui.Prompt{
			Label:   fmt.Sprintf("Ticket/Issue number (optional, prefix: %s)", cfg.Commit.TicketPrefix),
			Default: "",
		}

		ticketNumber, err := ticketPrompt.Run()
		if err == nil && ticketNumber != "" {
			ticket = cfg.Commit.TicketPrefix + ticketNumber
		}
		fmt.Println()
	}

	message := git.BuildCommitMessage(commitType, scope, description, body, breaking, breakingDesc, ticket)

	// COMMIT MESSAGE PREVIEW
	fmt.Println()
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("Commit Message Preview:")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	lines := strings.Split(message, "\n")
	for i, line := range lines {
		if i == 0 {
			fmt.Printf("  %s\n", line)
		} else if line == "" {
			fmt.Println()
		} else {
			fmt.Printf("  %s\n", line)
		}
	}

	fmt.Println()
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	// CONFIRMATION
	confirmPrompt := promptui.Prompt{
		Label:     "Proceed with this commit?",
		IsConfirm: true,
	}

	_, err = confirmPrompt.Run()
	if err != nil {
		fmt.Println()
		fmt.Println("❌ Commit cancelled")
		os.Exit(0)
	}

	fmt.Println()

	// 13. Execute the commit
	fmt.Println("Committing changes...")
	if err := git.Commit(message); err != nil {
		fmt.Printf("❌ Commit failed: %v\n", err)
		os.Exit(1)
	}

	// 15. Success message
	fmt.Println("✅ Commit successful!")
	fmt.Println()

	// Show the commit hash
	hash, err := git.GetLastCommitHash()
	if err == nil {
		fmt.Printf("📌 Commit: %s\n", hash)
	}

	fmt.Println()
	fmt.Println("💡 Next steps:")
	fmt.Println("   • Review: git log -1")
	fmt.Println("   • Push:   git push")
	fmt.Println()
}
