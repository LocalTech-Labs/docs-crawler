---
title: npx remotion upgrade
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion upgrade

- [Home page](/)
- [Command line](/docs/cli/)
- upgrade

On this page

# npx remotion upgrade

Upgrades all Remotion-related packages.

```

npx remotion upgrade
```

```

npx remotion upgrade
```

## Flags [​](\#flags "Direct link to Flags")

### `--package-manager` [v3.2.33](https://github.com/remotion-dev/remotion/releases/v3.2.33) [​](\#--package-manager "Direct link to --package-manager")

_optional_

Forces a specific package manager to be used. This is useful if you are using Remotion in a monorepo and you want to upgrade all packages at once. By default, Remotion will auto-detect the package manager.

Acceptable values are `npm`, `yarn` and `pnpm`

### `--version` [v4.0.15](https://github.com/remotion-dev/remotion/releases/v4.0.15) [​](\#--version "Direct link to --version")

Install a specific version. Also enables downgrading to an older version.

## Package manager support [​](\#package-manager-support "Direct link to Package manager support")

`npm`, `yarn` and `pnpm` are all supported.

## Difference to `npm update`, `yarn upgrade`, `pnpm up` [​](\#difference-to-npm-update-yarn-upgrade-pnpm-up "Direct link to difference-to-npm-update-yarn-upgrade-pnpm-up")

These commands, when executed without arguments will upgrade all dependencies in your project. We recommend against it because you may unintentionally break other parts of your project when you only wanted to upgrade Remotion.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/upgrade.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
versions](/docs/cli/versions) [Next\
\
ffmpeg](/docs/cli/ffmpeg)

- [Flags](#flags)
  - [`--package-manager`](#--package-manager)
  - [`--version`](#--version)
- [Package manager support](#package-manager-support)
- [Difference to `npm update`, `yarn upgrade`, `pnpm up`](#difference-to-npm-update-yarn-upgrade-pnpm-up)
