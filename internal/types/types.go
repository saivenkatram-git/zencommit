package types

var CommitTypes = []string{
	"feat     - A new feature",
	"fix      - A bug fix",
	"docs     - Documentation changes",
	"style    - Code style changes (formatting, etc)",
	"refactor - Code refactoring",
	"perf     - Performance improvements",
	"test     - Adding or updating tests",
	"build    - Build system changes",
	"ci       - CI configuration changes",
	"chore    - Other changes",
	"revert   - Revert a previous commit",
}

// GetCommitType extracts just the type keyword from the selected item
func GetCommitType(idx int) string {
	types := []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "build", "ci", "chore", "revert"}
	return types[idx]
}
