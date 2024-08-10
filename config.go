package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	ConfigPath         = filepath.Join(os.Getenv("HOME"), ".aws", "config")
	ErrProfileNotFound = fmt.Errorf("profile not found")
)

const (
	DefaultSectionName   = "default"
	ProfileSectionPrefix = "profile "
)

type Config struct {
	ini *ini.File
}

type Profile struct {
	Name            string
	IsEqualsDefault bool
}

func LoadConfig() (*Config, error) {
	cfg, err := ini.Load(ConfigPath)
	if err != nil {
		return nil, err
	}

	return &Config{cfg}, nil
}

func (cfg *Config) GetProfiles() []Profile {
	defaultSection := cfg.ini.Section(DefaultSectionName)
	profiles := []Profile{}

	for _, profileSection := range cfg.ini.Sections() {
		if strings.HasPrefix(profileSection.Name(), ProfileSectionPrefix) {
			name := profileSection.Name()[len(ProfileSectionPrefix):]
			isEqualsDefault := maps.Equal(profileSection.KeysHash(), defaultSection.KeysHash())
			profiles = append(profiles, Profile{name, isEqualsDefault})
		}
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Name < profiles[j].Name
	})

	return profiles
}

func (cfg *Config) GuessCurrentProfile() string {
	for _, profileSection := range cfg.GetProfiles() {
		if profileSection.IsEqualsDefault {
			return profileSection.Name
		}
	}

	return "unknown"
}

func (cfg *Config) CopyProfileToDefault(profile string) error {
	sectionName := ProfileSectionPrefix + profile
	if !cfg.ini.HasSection(sectionName) {
		return fmt.Errorf("%w: %s", ErrProfileNotFound, profile)
	}

	cfg.ini.DeleteSection(DefaultSectionName)

	defaultSection, err := cfg.ini.NewSection(DefaultSectionName)
	if err != nil {
		return err
	}

	profileSection := cfg.ini.Section(sectionName)
	for _, key := range profileSection.KeyStrings() {
		defaultSection.Key(key).SetValue(profileSection.Key(key).String())
	}

	err = cfg.ini.SaveTo(ConfigPath)
	if err != nil {
		return err
	}

	return nil
}
