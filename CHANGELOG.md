# Changelog

All notable changes to Zen Commit will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned Features

- Commit message templates
- Interactive scope creation
- Commit history viewer
- Statistics and analytics dashboard
- Plugin system for extensibility

## [0.2.0] - 2025-10-24

### Added

- **Git Hooks Integration** - Run custom commands before and after commits
  - Pre-commit hooks: Execute commands before committing (e.g., `npm test`, `npm run lint`)
  - Post-commit hooks: Execute commands after successful commit (e.g., update changelog)
  - Automatic commit abort when pre-commit hook fails
  - Hook output displayed in terminal for full transparency
  - Support for complex shell commands with `&&`, `||`, and pipes
- Hook execution with proper error handling and exit codes
- Detailed error messages when hooks fail

### Changed

- Improved commit message preview formatting
- Enhanced terminal output with better visual hierarchy
- Clearer error messages for hook failures

### Documentation

- Added comprehensive hooks documentation to README
- Added hook configuration examples for different project types
- Added troubleshooting section for common hook issues
- Included real-world hook examples (testing, linting, building)

### Example Configuration

```toml
[hooks]
pre_commit = "npm run lint && npm test"
post_commit = "npm run update-changelog"
```

### Migration Notes

- No breaking changes - fully backwards compatible
- Hooks section is optional in config
- Existing configurations work without modification

## [0.1.1] - 2025-10-24

### Fixed

- **Homebrew Installation**: Resolved SHA256 checksum mismatch during installation
- Cache-related issues causing installation failures
- Brew upgrade command not detecting new versions

### Documentation

- Added cache clearing instructions for Homebrew
- Improved troubleshooting section in README
- Added manual installation workarounds

### Technical Details

- Issue occurred when tags were recreated with same version number
- Solution: Clear Homebrew cache or use `brew reinstall`

## [0.1.0] - 2025-10-24

### Added

- 🎉 **Initial Release of Zen Commit**
- **Interactive Commit Builder**
  - Beautiful terminal UI with color and formatting
  - Arrow key navigation for selections
  - Step-by-step guided prompts
  - Real-time input validation
- **Conventional Commits Support**
  - Full support for conventional commit types:
    - `feat` - New features
    - `fix` - Bug fixes
    - `docs` - Documentation changes
    - `style` - Code style changes
    - `refactor` - Code refactoring
    - `perf` - Performance improvements
    - `test` - Adding or updating tests
    - `build` - Build system changes
    - `ci` - CI configuration changes
    - `chore` - Other changes
    - `revert` - Revert commits
- **Scope Management**
  - Custom scopes per project
  - Interactive scope selection
  - Optional scope support
  - Configurable scope requirement
- **Commit Message Features**
  - Commit message length validation (default: 72 characters)
  - Multi-line body support
  - Breaking changes detection with `!` notation
  - Required descriptions for breaking changes
  - Commit message preview before confirming
  - Subject line statistics
- **Configuration System**
  - TOML-based configuration (`zencommit.config.toml`)
  - Project-level customization
  - Initialization command: `zencommit init`
  - Sensible defaults out of the box
- **Emoji Support** (Optional)
  - Customizable emoji for each commit type
  - Default emoji mappings included
  - Easy enable/disable toggle
  - Example: ✨ feat, 🐛 fix, 📝 docs
- **Ticket/Issue Integration**
  - Support for ticket number prefixes
  - Configurable ticket format (e.g., JIRA-, GH-, etc.)
  - Automatic ticket number prompts
  - Appends ticket references to commit messages
- **Validation**
  - Git repository detection
  - Staged changes verification
  - Commit message length validation
  - Required field validation
  - Breaking change description validation
- **Cross-Platform Support**
  - macOS (Intel x86_64)
  - macOS (Apple Silicon ARM64)
  - Linux (x86_64)
  - Linux (ARM64)
  - Windows (x86_64)
  - FreeBSD (x86_64)
- **Installation Methods**
  - Homebrew tap for macOS/Linux
  - Manual binary download from GitHub Releases
  - Go install for developers
  - Automated installation via GoReleaser
- **User Experience**
  - Beautiful ASCII header
  - Color-coded terminal output
  - Clear success/error messages
  - Helpful tips and next steps
  - Graceful error handling
  - Cancel-friendly (Ctrl+C support)

### Configuration Options

```toml
[project]
name = "my-project"
description = "Project description"

[scopes]
enabled = true
list = ["api", "ui", "auth", "database", "docs"]

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
```

### Commands

- `zencommit init` - Initialize configuration file
- `zencommit commit` (alias: `zencommit c`) - Create a commit
- `zencommit version` - Show version information
- `zencommit help` - Show help message

### Documentation

- Comprehensive README with installation guide
- Step-by-step getting started tutorial
- Complete configuration reference
- Real-world usage examples
- Commit type reference table
- Troubleshooting guide
- Contributing guidelines
- MIT License

### Technical Details

- Written in Go for speed and efficiency
- Built with Cobra CLI framework
- Interactive prompts using promptui
- TOML parsing with BurntSushi/toml
- Automated releases via GoReleaser
- CI/CD with GitHub Actions
- Cross-compilation for multiple platforms
- Single binary with no dependencies

## Release Links

[Unreleased]: https://github.com/saivenkatram-git/zencommit/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/saivenkatram-git/zencommit/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/saivenkatram-git/zencommit/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/saivenkatram-git/zencommit/releases/tag/v0.1.0
