---
title: Rust development
source: Remotion Documentation
last_updated: 2024-12-22
---

# Rust development

- [Home page](/)
- Contributing
- Rust code

On this page

# Rust development

To participate in the development of the Rust parts of Remotion, you need to do additional steps.

These are entirely optional if you only touch the TypeScript parts.

## Setup [​](\#setup "Direct link to Setup")

First, install Cargo, if you don't have it, or upgrade to a version that supports `edition-2021`:

```

sh

curl https://sh.rustup.rs -sSf | sh
```

```

sh

curl https://sh.rustup.rs -sSf | sh
```

## Building [​](\#building "Direct link to Building")

To build the Rust parts for your operating system, run:

```

sh

cd packages/compositor
bun build.ts
```

```

sh

cd packages/compositor
bun build.ts
```

## Building for all platforms [​](\#building-for-all-platforms "Direct link to Building for all platforms")

These instructions currently are for macOS. Contributions for other platforms are appreciated.

To build the Rust binaries for all supported platforms, you need to install their toolchains:

```

sh

cd packages/compositor
node install-toolchains.mjs
```

```

sh

cd packages/compositor
node install-toolchains.mjs
```

You can then build all binaries with:

```

pnpm build-all
```

```

pnpm build-all
```

The resulting artifacts should be checked into Git.

## Architecture [​](\#architecture "Direct link to Architecture")

### Development workflow [​](\#development-workflow "Direct link to Development workflow")

- Remotion has platform specific packages like `compositor-darwin-arm64` and a user will only install packages that are specific to their operating system.
- The Rust code will be compiled into your native `compositor` package by default, unless you compile for all platforms.
- Resulting binaries are committed to Git so that people who don't setup Rust can develop too.

### FFmpeg bindings [​](\#ffmpeg-bindings "Direct link to FFmpeg bindings")

- Both the Rust binary as well as the FFmpeg binaries use the shared libraries located in the folder of each compositor package.
- Those shared libraries and FFmpeg binaries come from the [`rust-ffmpeg-splitter`](https://github.com/remotion-dev/rust-ffmpeg-splitter) repository. They are built in CircleCI as well as on a M1 Macbook and pasted into the [`rust-ffmpeg-sys`](https://github.com/remotion-dev/rust-ffmpeg-sys) project. The `rust-ffmpeg-sys` project is a dependency of [`rust-ffmpeg`](https://github.com/remotion-dev/rust-ffmpeg) which is a dependency of the main Remotion project. Refer to those repositories to learn how to build them.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/contributing/rust.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Bounty issues](/docs/contributing/bounty) [Next\
\
Adding new presentations](/docs/contributing/presentation)

- [Setup](#setup)
- [Building](#building)
- [Building for all platforms](#building-for-all-platforms)
- [Architecture](#architecture)
  - [Development workflow](#development-workflow)
  - [FFmpeg bindings](#ffmpeg-bindings)
