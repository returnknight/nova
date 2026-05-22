package config

import (
	"os"
	"path/filepath"
)

// Config 保存 Nova 的全局配置
type Config struct {
	OpenAIAPIKey        string `toml:"openai_api_key"`
	OpenAIBaseURL       string `toml:"openai_base_url"`
	OpenAIModel         string `toml:"openai_model"`
	SkillsDir           string `toml:"skills_dir"`
	NovaDir             string `toml:"nova_dir"`
	Workspace           string `toml:"workspace"`
	ResumeLastWorkspace bool   `toml:"-"`
}

// LoadWithWorkspace 在已知 workspace 时读取分层配置（默认 < 用户级 < 工作区级 < 环境变量）。
func LoadWithWorkspace(workspace string) (*Config, LayeredSettings, error) {
	novaDir := os.Getenv("NOVA_DIR")
	if novaDir == "" {
		novaDir = defaultNovaDir()
	}
	novaDir = normalizePath(novaDir)

	layered, err := LoadLayered(novaDir, workspace)
	if err != nil {
		return nil, LayeredSettings{}, err
	}

	s := layered.Effective
	cfg := &Config{
		OpenAIAPIKey:        s.OpenAIAPIKey,
		OpenAIBaseURL:       s.OpenAIBaseURL,
		OpenAIModel:         s.OpenAIModel,
		SkillsDir:           s.SkillsDir,
		NovaDir:             novaDir,
		Workspace:           workspace,
		ResumeLastWorkspace: true,
	}

	// 环境变量始终最高优先级
	overrideFromEnv(cfg)

	if cfg.Workspace != "" {
		if abs, err := filepath.Abs(cfg.Workspace); err == nil {
			cfg.Workspace = abs
		}
	}
	if cfg.SkillsDir != "" {
		cfg.SkillsDir = normalizePath(cfg.SkillsDir)
	}
	if cfg.NovaDir == "" {
		cfg.NovaDir = normalizePath(defaultNovaDir())
	} else {
		cfg.NovaDir = normalizePath(cfg.NovaDir)
	}
	return cfg, layered, nil
}

// Load 兼容旧入口：以当前目录作为 workspace 加载分层配置。
func Load() *Config {
	cwd, _ := os.Getwd()
	if cwd == "" {
		cwd = "."
	}
	cfg, _, err := LoadWithWorkspace(cwd)
	if err != nil || cfg == nil {
		// fallback：返回纯默认值 + env，保持启动不挂
		d := DefaultSettings()
		cfg = &Config{
			OpenAIBaseURL:       d.OpenAIBaseURL,
			OpenAIModel:         d.OpenAIModel,
			SkillsDir:           d.SkillsDir,
			NovaDir:             normalizePath(d.NovaDir),
			Workspace:           cwd,
			ResumeLastWorkspace: true,
		}
		overrideFromEnv(cfg)
		if abs, err := filepath.Abs(cfg.Workspace); err == nil {
			cfg.Workspace = abs
		}
		if cfg.SkillsDir != "" {
			cfg.SkillsDir = normalizePath(cfg.SkillsDir)
		}
	}
	return cfg
}

// overrideFromEnv 用环境变量覆盖配置
func overrideFromEnv(cfg *Config) {
	if v := os.Getenv("OPENAI_API_KEY"); v != "" {
		cfg.OpenAIAPIKey = v
	}
	if v := os.Getenv("OPENAI_BASE_URL"); v != "" {
		cfg.OpenAIBaseURL = v
	}
	if v := os.Getenv("OPENAI_MODEL"); v != "" {
		cfg.OpenAIModel = v
	}
	if v := os.Getenv("NOVA_SKILLS_DIR"); v != "" {
		cfg.SkillsDir = v
	}
	if v := os.Getenv("NOVA_DIR"); v != "" {
		cfg.NovaDir = v
	}
	if v := os.Getenv("NOVA_WORKSPACE"); v != "" {
		cfg.Workspace = v
	}
}

func defaultNovaDir() string {
	return "~/.nova"
}

func normalizePath(path string) string {
	path = expandHome(path)
	if abs, err := filepath.Abs(path); err == nil {
		return abs
	}
	return path
}

func expandHome(path string) string {
	if path == "~" {
		if home, err := os.UserHomeDir(); err == nil && home != "" {
			return home
		}
		return path
	}
	if len(path) > 2 && path[:2] == "~/" {
		if home, err := os.UserHomeDir(); err == nil && home != "" {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}
