package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Zones []string `json:"zones"`
}

func configFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("could not get user config dir: " + err.Error())
	}
	return filepath.Join(configDir, "tz", "config.json")
}

func LoadZones() []string {
	var cfg Config
	path := configFilePath()

	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return []string{}
	}

	return cfg.Zones
}

func AddZone(zone string) error {
	cfg := Config{Zones: LoadZones()}

	for _, z := range cfg.Zones {
		if z == zone {
			return errors.New("timezone already exists")
		}
	}

	cfg.Zones = append(cfg.Zones, zone)

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	path := configFilePath()
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func RemoveZone(zone string) error {
	cfg := Config{Zones: LoadZones()}

	newZones := []string{}
	found := false
	for _, z := range cfg.Zones {
		if z == zone {
			found = true
			continue
		}
		newZones = append(newZones, z)
	}
	if !found {
		return errors.New("timezone not found in config")
	}
	cfg.Zones = newZones

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath(), data, 0644)
}

func ConfigFilePath() string {
	return configFilePath()
}
