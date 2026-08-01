package utils

import (
	"fmt"
	"os"
	"gopkg.in/ini.v1"
)

func CreateConfig() error {
	if _, err := os.Stat("config.ini"); err == nil {
		return nil
	}

	cfg := ini.Empty()

	cfg.Section("backup").
		Key("interval").
		SetValue("10")

	return cfg.SaveTo("config.ini")
}

func GetConfigTime() (int, error) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return 0, fmt.Errorf("failed to load config: %w", err)
	}

	number, err := cfg.Section("backup").Key("interval").Int()
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}

	return number, nil
}

func SetConfigTime(value string) error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	cfg.Section("backup").
		Key("interval").
		SetValue(value)

	err = cfg.SaveTo("config.ini")
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Interval changed successfully to %s", value)

	return nil
}