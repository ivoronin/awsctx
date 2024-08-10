package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
)

var (
	ColorGreen     = color.New(color.FgGreen).SprintFunc()
	ColorRed       = color.New(color.FgRed).SprintFunc()
	CheckMarkGreen = ColorGreen("✔")

	version = ""
	commit  = ""
	date    = ""
)

type Options struct {
	Current bool   `arg:"-c,--current" help:"Show the current profile"`
	Profile string `arg:"positional" help:"Profile name"`
}

func (Options) Description() string {
	return `awsctx is a tool to manage AWS SDK configuration profiles`
}

func (Options) Version() string {
	return fmt.Sprintf("version %s (commit: %s, build date: %s)", version, commit, date)
}

func (Options) Epilogue() string {
	return "For more information see https://github.com/ivoronin/awsctx"
}

func listProfiles(cfg *Config) {
	for _, profileSection := range cfg.GetProfiles() {
		profile := profileSection.Name
		if profileSection.IsEqualsDefault {
			profile = ColorGreen(profile)
		}

		fmt.Println(profile)
	}
}

func showCurentProfile(cfg *Config) {
	fmt.Println(cfg.GuessCurrentProfile())
}

func setProfile(cfg *Config, profile string) {
	err := cfg.CopyProfileToDefault(profile)
	if err != nil {
		fatalError(err)
	}

	fmt.Printf("%s Switched to profile \"%s\"\n", CheckMarkGreen, profile)
}

func fatalError(err error) {
	fmt.Printf("%s %s\n", ColorRed("error:"), err)
	os.Exit(1)
}

func main() {
	var opts Options

	arg.MustParse(&opts)

	cfg, err := LoadConfig()
	if err != nil {
		fatalError(err)
	}

	switch {
	case opts.Current:
		showCurentProfile(cfg)
	case opts.Profile == "":
		listProfiles(cfg)
	default:
		setProfile(cfg, opts.Profile)
	}
}
