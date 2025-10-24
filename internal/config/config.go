package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Project  ProjectConfig  `toml:"project"`
	Scopes   ScopesConfig   `toml:"scopes"`
	Commit   CommitConfig   `toml:"commit"`
	Breaking BreakingConfig `toml:"breaking"`
	Emoji    EmojiConfig    `toml:"emoji"`
	Hooks    HooksConfig    `toml:"hooks"`
}

type ProjectConfig struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
}

type ScopesConfig struct {
	Enabled bool     `toml:"enabled"`
	List    []string `toml:"list"`
}

type CommitConfig struct {
	MaxLength    int    `toml:"max_length"`
	RequireScope bool   `toml:"require_scope"`
	TicketPrefix string `toml:"ticket_prefix"`
}

type BreakingConfig struct {
	RequireDescription bool `toml:"require_description"`
}

type EmojiConfig struct {
	Enabled bool              `toml:"enabled"`
	Types   map[string]string `toml:"types"`
}

type HooksConfig struct {
	PreCommit  string `toml:"pre_commit"`
	PostCommit string `toml:"post_commit"`
}

func Load() (*Config, error) {
	configPath := "zc.config.toml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) GetScopes() []string {
	if !c.Scopes.Enabled {
		return []string{}
	}
	return c.Scopes.List
}

func (c *Config) GetEmoji(commitType string) string {
	if !c.Emoji.Enabled {
		return ""
	}
	if emoji, ok := c.Emoji.Types[commitType]; ok {
		return emoji + " "
	}
	return ""
}
