package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	TargetChatID int64  `json:"target_chat_id"`
	Token        string `json:"token"`
	Debug        bool   `json:"debug"`
}

func (cfg *Config) Validate() error {
	if cfg.Token == "" {
		return fmt.Errorf("config: token not present")
	}

	if cfg.TargetChatID == 0 {
		return fmt.Errorf("config: target_chat_id not present")
	}

	return nil
}

// ParseFromFile читает файл конфигурации с диска и загружает его содержимое
// в структуру конфига приложения. Поддерживаются файлы в формате YAML и JSON.
func (cfg *Config) ParseFromFile(name string) error {
	file, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	return json.Unmarshal(file, cfg)
}
