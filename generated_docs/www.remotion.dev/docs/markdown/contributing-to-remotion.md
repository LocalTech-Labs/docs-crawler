---
title: Contributing to Remotion
source: Remotion Documentation
last_updated: 2024-12-21
---

# Contributing to Remotion

- [Home page](/)
- Contributing
- General information

On this page

# Contributing to Remotion

Issues and pull requests of all sorts are welcome!

For bigger projects, please coordinate with Jonny Burger ( [jonny@remotion.dev](mailto:jonny@remotion.dev), [Discord](https://remotion.dev/discord): `@jonnyburger`) to make sure your changes get merged.

Please note that since we charge for Remotion when companies are using it, this is a **commercial project**.

By sending pull requests, you agree that we can use your code changes in a commercial context.

Furthermore, also note that you **cannot redistribute** this project. Please see [LICENSE.md](https://remotion.dev/license) for what's allowed and what's not.

This project is released with a [Contributor Code of Conduct](https://remotion.dev/coc). By participating in this project you agree to abide by its terms.

## Setup [​](\#setup "Direct link to Setup")

[1](#1)

Remotion uses [`pnpm` v8.10.2](https://pnpm.io/) as the package manager for development in this repository. We recommend
using Corepack so you get the same version of pnpm as we have.

```

sh

corepack enable
```

```

sh

corepack enable
```

If you don't have `corepack`, install `pnpm@8.10.2` manually:

```

sh

npm i -g pnpm@8.10.2
```

```

sh

npm i -g pnpm@8.10.2
```

Prefix with `sudo` if necessary.

The version must be `8.10.2`.

We use [Bun](https://bun.sh) to speed up some parts of the build.

Install at least Bun 1.1.7 on your system - see [https://bun.sh/](https://bun.sh/) for instructions.

[2](#2)

Clone the Remotion repository:

```

sh

git clone --depth=1 https://github.com/remotion-dev/remotion.git && cd remotion
```

```

sh

git clone --depth=1 https://github.com/remotion-dev/remotion.git && cd remotion
```

note

The full Git history of Remotion is large. To save time and disk space, we recommend adding `--depth=1` to only clone the most recent `main` branch.

[3](#3)

Install all dependencies:

```

sh

pnpm i
```

```

sh

pnpm i
```

[4](#4)

Build the project initially:

```

sh

pnpm build
```

```

sh

pnpm build
```

[5](#5)

Rebuild whenever a file changes:

```

sh

pnpm watch
```

```

sh

pnpm watch
```

[6](#6)

You can start making changes!

## Testing your changes [​](\#testing-your-changes "Direct link to Testing your changes")

You can start the Testbed using

```

sh

cd packages/example
pnpm run dev
```

```

sh

cd packages/example
pnpm run dev
```

You can render a test video using

```

sh

cd packages/example
pnpm exec remotion render
```

```

sh

cd packages/example
pnpm exec remotion render
```

You can run tests using

```

sh

pnpm test
```

```

sh

pnpm test
```

in either a subpackage to run tests for that package or in the root to run all tests.

## Running the `@remotion/player` testbed [​](\#running-the-remotionplayer-testbed "Direct link to running-the-remotionplayer-testbed")

You can test changes to [@remotion/player](https://remotion.dev/docs/player) by starting the Player testbed:

```

sh

cd packages/player-example
pnpm run dev
```

```

sh

cd packages/player-example
pnpm run dev
```

For information about testing, please consult [TESTING.md](https://github.com/remotion-dev/remotion/blob/main/TESTING.md). Issues and pull requests of all sorts are welcome!

## Running documentation [​](\#running-documentation "Direct link to Running documentation")

You can run the Docusaurus server that powers our docs using:

```

sh

cd packages/docs
pnpm start
```

```

sh

cd packages/docs
pnpm start
```

## Running the CLI [​](\#running-the-cli "Direct link to Running the CLI")

You can test changes to the CLI by moving to `packages/examples` directory and using `pnpm exec` to execute the CLI:

```

sh

cd packages/examples
# Example - Get available compositions
pnpm exec remotion compositions
# Example - Render command
pnpm exec remotion render ten-frame-tester --output ../../out/video.mp4
```

```

sh

cd packages/examples
# Example - Get available compositions
pnpm exec remotion compositions
# Example - Render command
pnpm exec remotion render ten-frame-tester --output ../../out/video.mp4
```

## Testing Remotion Lambda [​](\#testing-remotion-lambda "Direct link to Testing Remotion Lambda")

In `packages/example`, there is a `runlambda.sh` script that will rebuild the code for the Lambda function, remove any deployed Lambda functions, deploy a new one and render a video.

You need to put you [AWS credentials](/docs/lambda/authentication) in a `.env` file of the `packages/example` directory.

```

sh

cd packages/example
sh ./runlambda.sh
```

```

sh

cd packages/example
sh ./runlambda.sh
```

note

This will delete any Lambda functions in your account!

## Testing Remotion Cloud Run [​](\#testing-remotion-cloud-run "Direct link to Testing Remotion Cloud Run")

In `packages/example`, there is a `runcloudrun.sh` script that will rebuild the code for the Cloud Run function, remove any deployed Cloud Run services, deploy a new one and render a video.

You need to put you [GCP credentials](/docs/cloudrun/generate-env) in a `.env` file of the `packages/example` directory.

```

sh

cd packages/example
sh ./runcloudrun.sh
```

```

sh

cd packages/example
sh ./runcloudrun.sh
```

note

This will delete any Cloud Run services in your account!

## Troubleshooting [​](\#troubleshooting "Direct link to Troubleshooting")

If your `pnpm build` throws errors, oftentimes it is because of caching issues. You can resolve many of these errors by running

```

ts

pnpm run clean
```

```

ts

pnpm run clean
```

in the root directory. Make sure to beforehand kill any `pnpm watch` commands, as those might regenerate files as you clean them!

## Developing Rust parts [​](\#developing-rust-parts "Direct link to Developing Rust parts")

To develop the Rust parts of Remotion, see [here](/docs/contributing/rust).

## See also [​](\#see-also "Direct link to See also")

- [Implementing a new feature](/docs/contributing/feature)
- [Implementing a new option](/docs/contributing/option)
- [Writing documentation](/docs/contributing/docs)
- [How to take a bounty issue](/docs/contributing/bounty)
- [Formatting guidelines](/docs/contributing/formatting)
- [Authoring Remotion libraries](/docs/authoring-packages)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Difference to Motion Canvas](/docs/compare/motion-canvas) [Next\
\
Implementing a new feature](/docs/contributing/feature)

- [Setup](#setup)
- [Testing your changes](#testing-your-changes)
- [Running the `@remotion/player` testbed](#running-the-remotionplayer-testbed)
- [Running documentation](#running-documentation)
- [Running the CLI](#running-the-cli)
- [Testing Remotion Lambda](#testing-remotion-lambda)
- [Testing Remotion Cloud Run](#testing-remotion-cloud-run)
- [Troubleshooting](#troubleshooting)
- [Developing Rust parts](#developing-rust-parts)
- [See also](#see-also)
