---
title: Implementing a new option
source: Remotion Documentation
last_updated: 2024-12-22
---

# Implementing a new option

- [Home page](/)
- Contributing
- Implementing a new option

On this page

# Implementing a new option

Adding a new option to a feature is a great way to contribute to Remotion and enable more use cases.

## Ways of specifying an option [​](\#ways-of-specifying-an-option "Direct link to Ways of specifying an option")

There are multiple ways to specify an option:

- APIs that can be called via Node.JS can directly accept a parameter.
- Options influencing a render can be added to the render dialog in the [Remotion Studio](/docs/terminology/studio).
- If the action is available as a [CLI](/docs/cli) command, the option should also be added as a CLI flag.
- The [config](/docs/config) file can be also be used to specify multiple options.



note



The config file should not be read when executing a CLI command, otherwise it's values should be ignored.

## Option Resolution [​](\#option-resolution "Direct link to Option Resolution")

The option should be resolved in the following order:

[1](#1)

Directly passed to the Node.JS API

[2](#2) Specified via the render UI in the [Remotion Studio](/docs/terminology/studio)

[3](#3) Specified via [CLI flag](/docs/cli)

[4](#4) Specified via [config file](/docs/config)

[5](#5) Fallback to a default value

## Naming [​](\#naming "Direct link to Naming")

The option should be named in `camelCase` for options in Node.JS and in `hyphen-case` for options accepted by the CLI.

Options accepting numerical values should include the unit in the name. For example `durationInFrames` instead of `duration` or `timeoutInMilliseconds` instead of `timeout`.

## Documentation [​](\#documentation "Direct link to Documentation")

The option should be documented in the [API reference](/docs/api) and the [CLI reference](/docs/cli).

Note from which version the option is available.

## Testing changes to the CLI [​](\#testing-changes-to-the-cli "Direct link to Testing changes to the CLI")

Check out the [Running the CLI](/docs/contributing) section in the contributing docs.

## Add CLI autocompletion [​](\#add-cli-autocompletion "Direct link to Add CLI autocompletion")

In the `cli-autocomplete` package under `src/source.ts`, you can add new options. People who use [Fig](https://fig.io) will then be able to get accurate autocompletion for your feature.

## See also [​](\#see-also "Direct link to See also")

- [Implementing a new feature](/docs/contributing/feature)
- [Writing documentation](/docs/contributing/docs)
- [How to take a bounty issue](/docs/contributing/bounty)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/option.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Implementing a new feature](/docs/contributing/feature) [Next\
\
Writing docs](/docs/contributing/docs)

- [Ways of specifying an option](#ways-of-specifying-an-option)
- [Option Resolution](#option-resolution)
- [Naming](#naming)
- [Documentation](#documentation)
- [Testing changes to the CLI](#testing-changes-to-the-cli)
- [Add CLI autocompletion](#add-cli-autocompletion)
- [See also](#see-also)
