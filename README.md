# awsctx
![GitHub release (with filter)](https://img.shields.io/github/v/release/ivoronin/awsctx)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivoronin/awsctx)](https://goreportcard.com/report/github.com/ivoronin/awsctx)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/ivoronin/awsctx/main)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/ivoronin/awsctx/main.yml)
![GitHub top language](https://img.shields.io/github/languages/top/ivoronin/awsctx)

## Description

`awsctx` is a tool to manage AWS SDK configuration profiles inspired by kubectx.

[AWS SDKs and Tools: Shared config and credentials files](https://docs.aws.amazon.com/sdkref/latest/guide/file-format.html)

## Usage

<pre>
<b># awsctx</b>
prod
dev
<b>stage</b>
drp

<b># awsctx dev</b>
✔ Switched to profile "dev"

<b># awsctx -c</b>
dev
</pre>
