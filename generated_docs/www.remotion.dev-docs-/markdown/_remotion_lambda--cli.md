---
title: @remotion/lambda - CLI
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/lambda - CLI

- [Home page](/)
- [Command line](/docs/cli/)
- lambda

On this page

# @remotion/lambda - CLI

To use the Remotion Lambda CLI, you first need to [install it](/docs/lambda/setup).

## Commands [​](\#commands "Direct link to Commands")

- [sites](/docs/lambda/cli/sites)
- [functions](/docs/lambda/cli/functions)
- [render](/docs/lambda/cli/render)
- [still](/docs/lambda/cli/still)
- [policies](/docs/lambda/cli/policies)
- [compositions](/docs/lambda/cli/compositions)
- [regions](/docs/lambda/cli/regions)
- [quotas](/docs/lambda/cli/quotas)

## Global options [​](\#global-options "Direct link to Global options")

### `--region` [​](\#--region "Direct link to --region")

Selects an AWS region: For example:

```

--region=eu-central-1
```

```

--region=eu-central-1
```

The default region is `us-east-1`. You may also set a `REMOTION_AWS_REGION` environment variable directly or via `.env` file.

See: [Region selection](/docs/lambda/region-selection)

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Skips confirmation when doing a destructive action.

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Prints the minimal amount of logs.

### `--env-file` [​](\#--env-file "Direct link to --env-file")

Specify a location for a dotenv file - Default `.env`. [Read about how environment variables work in Remotion.](/docs/env-variables)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cli.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
bundle](/docs/cli/bundle) [Next\
\
sites](/docs/lambda/cli/sites)

- [Commands](#commands)
- [Global options](#global-options)
  - [`--region`](#--region)
  - [`--yes`, `-y`](#--yes--y)
  - [`--quiet`, `-q`](#--quiet--q)
  - [`--env-file`](#--env-file)
