---
title: npx remotion bundle
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion bundle

- [Home page](/)
- [Command line](/docs/cli/)
- bundle

On this page

# npx remotion bundle

_available from v4.0.89_

Creates a [Remotion Bundle](/docs/terminology/bundle) on the command line.

Equivalent to the [`bundle()`](/docs/bundle) Node.JS API.

```

bash

npx remotion bundle <serve-url|entry-file>?
```

```

bash

npx remotion bundle <serve-url|entry-file>?
```

You may pass a [Serve URL](/docs/terminology/serve-url) or an [entry point](/docs/terminology/entry-point) as the first argument, otherwise the entry point will be [determined](/docs/terminology/entry-point#which-entry-point-is-being-used).

## Flags [​](\#flags "Direct link to Flags")

### `--config` [​](\#--config "Direct link to --config")

Specify a location for the Remotion config file.

### `--log` [​](\#--log "Direct link to --log")

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

### `--public-dir` [​](\#--public-dir "Direct link to --public-dir")

Define the location of the [`public/ directory`](/docs/terminology/public-dir). If not defined, Remotion will assume the location is the \`public\` folder in your Remotion root.

### `--out-dir` [​](\#--out-dir "Direct link to --out-dir")

Define the location of the resulting bundle. By default it is a folder called `./build`, adjacent to the [Remotion Root](/docs/terminology/remotion-root).

### `--public-path` [v4.0.127](https://github.com/remotion-dev/remotion/releases/v4.0.127) [​](\#--public-path "Direct link to --public-path")

The path of the URL where the bundle is going to be hosted. By default it is `/`, meaning that the bundle is going to be hosted at the root of the domain (e.g. `https://localhost:3000/`). If you are deploying to a subdirectory (e.g. `/sites/my-site/`), you should set this to the subdirectory.

### `--disable-git-source` [v4.0.182](https://github.com/remotion-dev/remotion/releases/v4.0.182) [​](\#--disable-git-source "Direct link to --disable-git-source")

Disables the Git Source being connected to the Remotion Studio. Clicking on stack traces and certain menu items will be disabled.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/bundle.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
compositions](/docs/cli/compositions) [Next\
\
Overview](/docs/lambda/cli)

- [Flags](#flags)
  - [`--config`](#--config)
  - [`--log`](#--log)
  - [`--public-dir`](#--public-dir)
  - [`--out-dir`](#--out-dir)
  - [`--public-path`](#--public-path)
  - [`--disable-git-source`](#--disable-git-source)
