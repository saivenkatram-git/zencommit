![image alt](https://github.com/saivenkatram-git/zencommit/blob/f28cad5f91077c964de5aeaddb72c184a65250da/assets/logo.png)

# Zen Commit

A beautiful, interactive CLI tool for creating [conventional commits](https://www.conventionalcommits.org/).

[![Release](https://img.shields.io/github/v/release/saivenkatram-git/zencommit)](https://github.com/saivenkatram-git/zencommit/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## ✨ Features

- **Beautiful interactive prompts** - Intuitive UI with arrow key navigation
- **Conventional commits** - Industry-standard commit message format
- **Configurable via TOML** - Customize per project
- **Custom scopes** - Define scopes specific to your project
- **Git hooks** - Optional pre/post-commit hooks
- **Emoji support** - Add visual flair to your commits
- **Fast** - Written in Go, starts instantly

---

## Prerequisites

Before using Zen Commit, ensure you have:

- **Git installed** - [Download Git](https://git-scm.com/downloads)
- **A Git repository** - Your project must be initialized with `git init`
- **Staged changes** - Files must be staged with `git add` before committing

---

## Installation

### macOS / Linux (Homebrew)

```bash
brew tap saivenkatram-git/tap
brew install zencommit
```

### Manual Installation

1. Download the latest release from [releases](https://github.com/saivenkatram-git/zencommit/releases)
2. Extract the archive
3. Move the binary to your PATH:
   ```bash
   sudo mv zencommit /usr/local/bin/
   ```
4. Verify installation:
   ```bash
   zencommit version
   ```

### Using Go

```bash
go install github.com/saivenkatram-git/zencommit@latest
```

---

## Getting Started

Follow these steps to start using Zen Commit:

### Step 1: Navigate to Your Git Repository

```bash
cd your-project
```

**Note:** Your project must already be a Git repository. If not, initialize it:

```bash
git init
```

### Step 2: Initialize Zen Commit

```bash
zencommit init
```

This creates a `zencommit.config.toml` file in your project root with default configuration.

**Output:**

```
✅ Created zencommit.config.toml

📝 Configuration created with:
   • Default scopes (api, ui, auth, etc.)
   • Commit message length limit (72 chars)
   • Optional emoji support

📖 Next steps:
   1. Edit zencommit.config.toml to customize your settings
   2. Run: git add zencommit.config.toml
   3. Run: zencommit commit
```

### Step 3: Make Changes to Your Code

```bash
# Edit files in your project
echo "new feature" > feature.txt
```

### Step 4: Stage Your Changes

```bash
# Stage specific files
git add feature.txt

# Or stage all changes
git add .
```

**Important:** You must stage files before running `zencommit commit`.

### Step 5: Create a Commit

```bash
zencommit commit
```

Or use the shorthand:

```bash
zencommit c
```

### Step 6: Follow the Interactive Prompts

```
▸ feat     - A new feature
  fix      - A bug fix
  docs     - Documentation changes
  style    - Code style changes
  refactor - Code refactoring
  perf     - Performance improvements
  test     - Adding or updating tests
  build    - Build system changes
  ci       - CI configuration changes
  chore    - Other changes

✓ Type: feat

▸ (none) - No scope
  api
  ui
  auth
  database
  docs

✓ Scope: api

Short description (max 72 characters): add user authentication endpoint

Longer description (optional):

Any breaking changes? (y/N): n

════════════════════════════════════════════════════════════
📝 Commit Message Preview:
════════════════════════════════════════════════════════════

  feat(api): add user authentication endpoint

════════════════════════════════════════════════════════════

📊 Stats: Subject length: 42/72 characters

Proceed with commit? (y/N): y

✅ Commit successful!
📌 Commit: a1b2c3d
```

---

## Configuration

The `zencommit.config.toml` file allows you to customize Zen Commit for your project.

### Example Configuration

```toml
[project]
name = "my-awesome-app"
description = "A web application"

[scopes]
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
max_length = 72
require_scope = false
ticket_prefix = ""

[breaking]
require_description = true

[emoji]
enabled = false

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
# pre_commit = "npm test"
# post_commit = "echo 'Committed successfully!'"
```

### Configuration Options

| Section      | Option                | Description                              | Default              |
| ------------ | --------------------- | ---------------------------------------- | -------------------- |
| `[scopes]`   | `enabled`             | Enable scope selection                   | `true`               |
| `[scopes]`   | `list`                | Available scopes for your project        | `["api", "ui", ...]` |
| `[commit]`   | `max_length`          | Maximum subject line length              | `72`                 |
| `[commit]`   | `require_scope`       | Make scope mandatory                     | `false`              |
| `[commit]`   | `ticket_prefix`       | Add ticket numbers (e.g., "JIRA-")       | `""`                 |
| `[breaking]` | `require_description` | Require description for breaking changes | `true`               |
| `[emoji]`    | `enabled`             | Add emoji to commit messages             | `false`              |
| `[hooks]`    | `pre_commit`          | Command to run before commit             | `""`                 |
| `[hooks]`    | `post_commit`         | Command to run after commit              | `""`                 |

### Customizing Scopes

Edit `zencommit.config.toml` to define scopes specific to your project:

**Frontend project:**

```toml
[scopes]
list = ["components", "pages", "styles", "hooks", "utils"]
```

**Backend project:**

```toml
[scopes]
list = ["routes", "controllers", "models", "middleware", "services"]
```

**Full-stack project:**

```toml
[scopes]
list = ["api", "ui", "auth", "database", "docs", "config"]
```

---

## Commit Types

Zen Commit follows the [Conventional Commits](https://www.conventionalcommits.org/) specification:

| Type       | Description              | Example                               |
| ---------- | ------------------------ | ------------------------------------- |
| `feat`     | A new feature            | `feat(auth): add login page`          |
| `fix`      | A bug fix                | `fix(ui): resolve navbar bug`         |
| `docs`     | Documentation only       | `docs: update installation guide`     |
| `style`    | Code style changes       | `style: format with prettier`         |
| `refactor` | Code refactoring         | `refactor(api): simplify user routes` |
| `perf`     | Performance improvements | `perf: optimize image loading`        |
| `test`     | Adding tests             | `test: add auth unit tests`           |
| `build`    | Build system changes     | `build: update webpack config`        |
| `ci`       | CI configuration         | `ci: add GitHub Actions workflow`     |
| `chore`    | Other changes            | `chore: update dependencies`          |
| `revert`   | Revert a commit          | `revert: undo feature X`              |

---

## 💡 Examples

### Example 1: Simple Feature

```bash
$ zencommit commit

? Select commit type: feat
? Select scope: (none)
? Short description: add dark mode toggle
✅ Commit successful!

# Result: feat: add dark mode toggle
```

### Example 2: Bug Fix with Scope

```bash
$ zencommit commit

? Select commit type: fix
? Select scope: api
? Short description: resolve CORS issue
✅ Commit successful!

# Result: fix(api): resolve CORS issue
```

### Example 3: Breaking Change

```bash
$ zencommit commit

? Select commit type: feat
? Select scope: auth
? Short description: redesign authentication flow
? Any breaking changes? Yes
? Describe the breaking change: Old session format no longer supported

# Result:
# feat(auth)!: redesign authentication flow
#
# BREAKING CHANGE: Old session format no longer supported
```

### Example 4: With Emoji (when enabled)

```toml
[emoji]
enabled = true
```

```bash
$ zencommit commit

? Select commit type: feat
? Short description: add user profile page
✅ Commit successful!

# Result: ✨ feat: add user profile page
```

### Example 5: With Ticket Number

```toml
[commit]
ticket_prefix = "JIRA-"
```

```bash
$ zencommit commit

? Select commit type: feat
? Select scope: api
? Short description: add user endpoint
? Ticket/Issue number: 1234

# Result: feat(api): add user endpoint [JIRA-1234]
```

---

## 🔧 Advanced Usage

### Git Hooks

Run commands automatically before or after commits:

```toml
[hooks]
pre_commit = "npm test"           # Run tests before committing
post_commit = "npm run changelog"  # Update changelog after commit
```

**Note:** If `pre_commit` fails, the commit will be aborted.

### Enforce Scope Requirement

Make scopes mandatory for all commits:

```toml
[commit]
require_scope = true
```

Now users must select a scope - the "(none)" option is removed.

### Custom Message Length

Adjust the maximum commit message length:

```toml
[commit]
max_length = 50  # Shorter messages
```

---

## 🚨 Troubleshooting

### Error: "Not a git repository"

**Problem:** Zen Commit requires a Git repository.

**Solution:**

```bash
git init
```

### Error: "No staged changes"

**Problem:** You need to stage files before committing.

**Solution:**

```bash
git add .
```

### Error: "config file not found"

**Problem:** Zen Commit requires initialization.

**Solution:**

```bash
zencommit init
```

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/saivenkatram-git/zencommit.git
cd zencommit

# Install dependencies
go mod download

# Build
go build -o zencommit

# Test
./zencommit version
```

---

## 📄 License

MIT © 2025 [Sai Venkatram](https://github.com/saivenkatram-git)

---

## 🙏 Acknowledgments

- Inspired by [Commitizen](https://github.com/commitizen/cz-cli)
- Built with [Cobra](https://github.com/spf13/cobra) and [promptui](https://github.com/manifoldco/promptui)
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification

---

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

---

**[Documentation](https://github.com/saivenkatram-git/zencommit)** • **[Report Bug](https://github.com/saivenkatram-git/zencommit/issues)** • **[Request Feature](https://github.com/saivenkatram-git/zencommit/issues)**
