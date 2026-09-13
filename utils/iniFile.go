package utils

import (
	"fmt"
	"os"
	"strconv"
	"gopkg.in/ini.v1"
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

func timeTranslater(time string) (int, error){
	num := time[:len(time)-1]

	switch(time[len(time) - 1]){
		case 's': {
			value, err := strconv.Atoi(num)

			if err != nil {
				return 0, fmt.Errorf("invalid input: %q is not a number", num)
			}

			return value, nil
		}

		case 'm': {
			value, err := strconv.Atoi(num)

			if err != nil {
				return 0, fmt.Errorf("invalid input: %q is not a number", num)
			}

			return value * 60, nil
		}

		case 'h': {
			value, err := strconv.Atoi(num)

			if err != nil {
				return 0, fmt.Errorf("invalid input: %q is not a number", num)
			}

			return value  * 60 * 60, nil
		}

		default: return 0, fmt.Errorf("Missing extension")
	}
}