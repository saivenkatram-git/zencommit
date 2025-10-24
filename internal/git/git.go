package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run() == nil
}

func HasStagedChanges() bool {
	cmd := exec.Command("git", "diff", "--cached", "--quiet")
	return cmd.Run() != nil
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func GetLastCommitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func BuildCommitMessage(commitType, scope, desc, body string, breaking bool, breakingDesc string, ticket string) string {
	var msg strings.Builder

	// First line: type(scope): description
	msg.WriteString(commitType)
	if scope != "" {
		msg.WriteString("(" + scope + ")")
	}
	if breaking {
		msg.WriteString("!")
	}
	msg.WriteString(": " + desc)

	// Body (optional)
	if body != "" {
		msg.WriteString("\n\n" + body)
	}

	if ticket != "" {
		msg.WriteString("\n\n" + ticket)
	}

	// Breaking change footer
	if breaking && breakingDesc != "" {
		msg.WriteString("\n\nBREAKING CHANGE: " + breakingDesc)
	}

	return msg.String()
}

func ExecuteHook(command string) error {
	if command == "" {
		return nil
	}

	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()

	if len(output) > 0 {
		fmt.Println(string(output))
	}

	return err
}
