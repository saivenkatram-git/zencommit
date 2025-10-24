# Zen Commit

A beautiful, interactive tool for creating conventional commits.

## ✨ Features

- 🎨 Beautiful interactive prompts
- 📝 Conventional commits standard
- ⚙️ Configurable via TOML
- 🎯 Custom scopes per project
- 🔧 Pre/post-commit hooks
- ✅ Emoji support

## 🚀 Installation

### macOS/Linux (Homebrew)

\`\`\`bash
brew tap yourusername/tap
brew install zencommit
\`\`\`

### Manual Installation

Download from [releases](https://github.com/saivenkatram-git/zencommit/releases)

## 📖 Usage

\`\`\`bash

# Initialize in your project

zencommit init

# Create a commit

zencommit commit

# Or use the shorthand

zencommit c

# Check version

zencommit version
\`\`\`

## ⚙️ Configuration

Run `zencommit init` to create `zencommit.config.toml`:

\`\`\`toml
[scopes]
list = ["api", "ui", "docs"]

[commit]
max_length = 72

[emoji]
enabled = true
\`\`\`

## 📝 Examples

\`\`\`bash

# Interactive commit

zencommit commit

# Output:

# ╔════════════════════════════════════════════════╗

# ║ ZC - Zen Commit Creator ║

# ╚════════════════════════════════════════════════╝

#

# ? Select commit type: feat

# ? Select scope: api

# ? Short description: add user endpoint

# ...

\`\`\`

## 🤝 Contributing

Contributions welcome! Please open an issue or PR.

## 📄 License

MIT © Sai Venkatram
