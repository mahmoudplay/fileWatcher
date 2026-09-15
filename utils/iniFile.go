package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
)

func CreateConfig(interval int) error {
	if _, err := os.Stat("config.ini"); err == nil {
		return nil
	}

	if interval <= 0 {
		interval = 10
	}

	cfg := ini.Empty()

	cfg.Section("backup").
		Key("interval").
		SetValue(strconv.Itoa(interval))

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
	number, err := timeTranslater(value)
	if err != nil {
		return err
	}
	cfg, err := ini.Load("config.ini")
	if err != nil {
		if err := CreateConfig(number); err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}
		cfg, err = ini.Load("config.ini")
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
	}

	cfg.Section("backup").
		Key("interval").
		SetValue(strconv.Itoa(number))

	err = cfg.SaveTo("config.ini")
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Interval changed successfully to %s", value)

	return nil
}

func timeTranslater(value string) (int, error) {
	if len(value) < 2 {
		return 0, fmt.Errorf("invalid input: %q must be a number followed by s, m, or h", value)
	}

	unit := value[len(value)-1]
	num := value[:len(value)-1]

	multiplier := 1
	switch unit {
	case 's':
		multiplier = 1
	case 'm':
		multiplier = 60
	case 'h':
		multiplier = 60 * 60
	default:
		return 0, fmt.Errorf("invalid unit %q: use s, m, or h", string(unit))
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %q is not a number", num)
	}

	return n * multiplier, nil
}
