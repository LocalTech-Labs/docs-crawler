---
title: @remotion/cloudrun - CLI
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/cloudrun - CLI

- [Home page](/)
- [Command line](/docs/cli/)
- cloudrun

On this page

# @remotion/cloudrun - CLI

EXPERIMENTAL

Cloud Run is in [Alpha](/docs/cloudrun-alpha), which means APIs may change in any version and documentation is not yet finished. See the [changelog to stay up to date with breaking changes](https://remotion.dev/changelog).

To use the Remotion Cloud Run CLI, you first need to [install it](/docs/cloudrun/setup).

## Commands [​](\#commands "Direct link to Commands")

- [sites](/docs/cloudrun/cli/sites)
- [services](/docs/cloudrun/cli/services)
- [render](/docs/cloudrun/cli/render)
- [still](/docs/cloudrun/cli/still)
- [permissions](/docs/cloudrun/cli/permissions)
- [regions](/docs/cloudrun/cli/regions)

## Global options [​](\#global-options "Direct link to Global options")

### `--region` [​](\#--region "Direct link to --region")

Selects a GCP region: For example:

```

--region=us-central1
```

```

--region=us-central1
```

The default region is `us-east1`. You may also set a `REMOTION_GCP_REGION` environment variable directly or via `.env` file.

See: [Region selection](/docs/cloudrun/region-selection)

### `--yes`, `-y` [​](\#--yes--y "Direct link to --yes--y")

Skips confirmation when doing a destructive action.

### `--quiet`, `-q` [​](\#--quiet--q "Direct link to --quiet--q")

Prints the minimal amount of logs.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cloudrun/cli.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
quotas](/docs/lambda/cli/quotas) [Next\
\
sites](/docs/cloudrun/cli/sites)

- [Commands](#commands)
- [Global options](#global-options)
  - [`--region`](#--region)
  - [`--yes`, `-y`](#--yes--y)
  - [`--quiet`, `-q`](#--quiet--q)
