---
title: Setup
source: Remotion Documentation
last_updated: 2024-12-22
---

# Setup

- [Home page](/)
- Recorder
- Setup

On this page

# Setup

First, [purchase](/docs/recorder/buy) the Recorder to [get access to the repo](https://github.com/remotion-dev/recorder).

## 1\. Fork and clone [​](\#1-fork-and-clone "Direct link to 1. Fork and clone")

[Create a fork](https://github.com/remotion-dev/recorder/fork) of the `remotion-dev/recorder` repository.

Clone the repository using `git clone` and `cd` into the directory.

note

You need to keep your fork **private** to not leak the source code of the Remotion Recorder.

Then, clone your fork to get the project locally.

## 2\. Install Bun [​](\#2-install-bun "Direct link to 2. Install Bun")

The Remotion Recorder requires [Bun](https://bun.sh/), with the minimum version being 1.1.11. To install it, run

```

bash

curl -fsSL https://bun.sh/install | bash
```

```

bash

curl -fsSL https://bun.sh/install | bash
```

## 3\. Install dependencies [​](\#3-install-dependencies "Direct link to 3. Install dependencies")

Once Bun is installed, run

```

bash

bun i
```

```

bash

bun i
```

to install all the neccessary dependencies.

## 4\. Install Whisper.cpp [​](\#4-install-whispercpp "Direct link to 4. Install Whisper.cpp")

_optional_

This is a good time to initialize Whisper.cpp for captioning.

This will install Whisper.cpp and download the model (the default being 1.5GB).

```

bun sub.ts
```

```

bun sub.ts
```

## 5\. Run the Remotion Recorder [​](\#5-run-the-remotion-recorder "Direct link to 5. Run the Remotion Recorder")

Once done, you are ready to start the Remotion Recorder:

```

bash

bun start
```

```

bash

bun start
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/setup.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Demo](/docs/recorder/demo) [Next\
\
Create a new video](/docs/recorder/create)

- [1\. Fork and clone](#1-fork-and-clone)
- [2\. Install Bun](#2-install-bun)
- [3\. Install dependencies](#3-install-dependencies)
- [4\. Install Whisper.cpp](#4-install-whispercpp)
- [5\. Run the Remotion Recorder](#5-run-the-remotion-recorder)
