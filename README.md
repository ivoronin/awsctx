# awsctx

Switch AWS SDK configuration profiles by copying profile settings to default

[![CI](https://github.com/ivoronin/awsctx/actions/workflows/test.yml/badge.svg)](https://github.com/ivoronin/awsctx/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/ivoronin/awsctx)](https://github.com/ivoronin/awsctx/releases)

[Overview](#overview) · [Features](#features) · [Installation](#installation) · [Usage](#usage) · [Configuration](#configuration) · [Requirements](#requirements) · [License](#license)

```
awsctx
prod
dev
stage
drp

awsctx dev
Switched to profile "dev"

awsctx -c
dev
```

## Overview

awsctx manages AWS SDK configuration profiles by modifying `~/.aws/config`. When switching profiles, it copies all settings from the selected profile section to the `[default]` section. This makes the selected profile active for all AWS SDKs and tools that use the default profile. The current profile is determined by comparing the default section contents against all named profiles.

## Features

- Lists all named profiles from `~/.aws/config`
- Highlights the current profile (matching default section)
- Switches profiles by copying settings to default section
- Shows current profile with `-c` flag
- Inspired by kubectx workflow

## Installation

### GitHub Releases

Download from [Releases](https://github.com/ivoronin/awsctx/releases).

### Homebrew

```bash
brew install ivoronin/tap/awsctx
```

### Build from source

```bash
go install github.com/ivoronin/awsctx/cmd/awsctx@latest
```

## Usage

### List profiles

```bash
awsctx
```

Lists all named profiles from the config file. The current profile (matching the default section) is highlighted in green.

### Switch profile

```bash
awsctx prod
```

Copies all settings from `[profile prod]` to `[default]` in `~/.aws/config`.

### Show current profile

```bash
awsctx -c
awsctx --current
```

Prints the name of the profile whose settings match the default section. Returns "unknown" if no match is found.

### Version

```bash
awsctx --version
```

## Configuration

awsctx reads and modifies `~/.aws/config`. Profiles must be defined using the standard AWS SDK format:

```ini
[default]
region = us-east-1
output = json

[profile prod]
region = us-east-1
output = json
sso_session = my-sso
sso_account_id = 123456789012
sso_role_name = AdminRole

[profile dev]
region = us-west-2
output = json
```

See [AWS SDK Shared config and credentials files](https://docs.aws.amazon.com/sdkref/latest/guide/file-format.html) for full format documentation.

## Requirements

- `~/.aws/config` file with named profiles
- Write access to `~/.aws/config` for profile switching

## License

[GPL-3.0](LICENSE)
