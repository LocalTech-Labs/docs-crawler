---
title: Dockerizing a Remotion app
source: Remotion Documentation
last_updated: 2024-12-22
---

# Dockerizing a Remotion app

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- Docker

On this page

# Dockerizing a Remotion app

We recommend the following structure for your Dockerfile. Read below about the individual steps and whether you need to adjust them.

```

Dockerfile
docker

FROM node:22-bookworm-slim

# Install Chrome dependencies
RUN apt-get update
RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libatk-bridge2.0-0 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2

# Copy everything from your project to the Docker image. Adjust if needed.
COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src

# If you have a public folder:
COPY public ./public

# Install the right package manager and dependencies - see below for Yarn/PNPM
RUN npm i

# Install Chrome
RUN npx remotion browser ensure

# Run your application
COPY render.mjs render.mjs
CMD ["node", "render.mjs"]
```

```

Dockerfile
docker

FROM node:22-bookworm-slim

# Install Chrome dependencies
RUN apt-get update
RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libatk-bridge2.0-0 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2

# Copy everything from your project to the Docker image. Adjust if needed.
COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src

# If you have a public folder:
COPY public ./public

# Install the right package manager and dependencies - see below for Yarn/PNPM
RUN npm i

# Install Chrome
RUN npx remotion browser ensure

# Run your application
COPY render.mjs render.mjs
CMD ["node", "render.mjs"]
```

note

[Click here](#example-render-script) to see an example for a `render.mjs` script you can use.

## Line-by-line [​](\#line-by-line "Direct link to Line-by-line")

[1](#1)

Specify the base image for the Dockerfile. In this case, we use
Debian.

```

docker

FROM node:22-bookworm-slim
```

```

docker

FROM node:22-bookworm-slim
```

[2](#2)

Update the package lists on the Debian system and install the
needed [shared libraries](/docs/miscellaneous/linux-dependencies) for Chrome
Headless Shell.

```

docker

RUN apt-get update
RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libatk-bridge2.0-0 \
  libcups2
```

```

docker

RUN apt-get update
RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libatk-bridge2.0-0 \
  libcups2
```

[3](#3)

Copy the files from your project. If you have additional source
files, add them here. If some files do not exist, remove them. The `COPY` syntax allows multiple files, but at least one file must
exist. It is assumed `package.json`, `src` and `public` exist in your project, but you can adjust this to your
needs.

```

docker

COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src
COPY public ./public
```

```

docker

COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src
COPY public ./public
```

[4](#4)

Install the right package manager and dependencies.

- If you use NPM, put the following in your Dockerfile:



```

docker

RUN npm i
```


```

docker

RUN npm i
```

- If you use Yarn or PNPM, add the `packageManager` field to your `package.json` (example: `"packageManager": "pnpm@7.7.1"`) and remove the `npm` line from step 3. Then put following in your Dockerfile:



```

If you use PNPM
docker

RUN corepack enable
RUN pnpm i
```


```

If you use PNPM
docker

RUN corepack enable
RUN pnpm i
```






```

If you use Yarn
docker

RUN corepack enable
RUN yarn
```


```

If you use Yarn
docker

RUN corepack enable
RUN yarn
```

[5](#5)

Install [Chrome Headless\
Shell](/docs/miscellaneous/chrome-headless-shell).

```

docker

RUN npx remotion browser ensure
```

```

docker

RUN npx remotion browser ensure
```

[6](#6)

Run your code. It can be a CLI command or a Node.JS app.

```

docker

COPY render.mjs render.mjs
CMD ["node", "render.mjs"]
```

```

docker

COPY render.mjs render.mjs
CMD ["node", "render.mjs"]
```

## Example render script [​](\#example-render-script "Direct link to Example render script")

Assuming you want to render the composition `MyComp`:

```

render.mjs
tsx

import {bundle} from '@remotion/bundler';
import {renderMedia, selectComposition} from '@remotion/renderer';
import {createRequire} from 'node:module';

const require = createRequire(import.meta.url);

const bundled = await bundle({
  entryPoint: require.resolve('./src/index.ts'),
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: (config) => config,
});

const inputProps = {};

const composition = await selectComposition({
  serveUrl: bundled,
  id: 'MyComp',
  inputProps,
});

console.log('Starting to render composition');

await renderMedia({
  codec: 'h264',
  composition,
  serveUrl: bundled,
  outputLocation: `out/${composition.id}.mp4`,
  chromiumOptions: {
    enableMultiProcessOnLinux: true,
  },
  inputProps,
});

console.log(`Rendered composition ${composition.id}.`);
```

```

render.mjs
tsx

import {bundle} from '@remotion/bundler';
import {renderMedia, selectComposition} from '@remotion/renderer';
import {createRequire} from 'node:module';

const require = createRequire(import.meta.url);

const bundled = await bundle({
  entryPoint: require.resolve('./src/index.ts'),
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: (config) => config,
});

const inputProps = {};

const composition = await selectComposition({
  serveUrl: bundled,
  id: 'MyComp',
  inputProps,
});

console.log('Starting to render composition');

await renderMedia({
  codec: 'h264',
  composition,
  serveUrl: bundled,
  outputLocation: `out/${composition.id}.mp4`,
  chromiumOptions: {
    enableMultiProcessOnLinux: true,
  },
  inputProps,
});

console.log(`Rendered composition ${composition.id}.`);
```

note

We recommend setting the `enableMultiProcessOnLinux` option for this Docker image, available from v4.0.42. [Read more](/docs/miscellaneous/linux-single-process)

## Building the Docker image [​](\#building-the-docker-image "Direct link to Building the Docker image")

Run

```

sh

docker build -t remotion-app .
```

```

sh

docker build -t remotion-app .
```

to build a Docker image called `remotion-app`.

Use the following command to run the image:

```

sh

docker run remotion-app
```

```

sh

docker run remotion-app
```

## Giving access to the CPUs [​](\#giving-access-to-the-cpus "Direct link to Giving access to the CPUs")

By default, Docker containers are not allowed to use all memory CPUs . Consider:

- Changing the resource settings in Docker Desktop settings
- Using the `--cpus` and `--cpuset-cpus` flags with the `docker run` command. Example: `--cpus=16 --cpuset-cpus=0-15`

## Emojis [​](\#emojis "Direct link to Emojis")

No emojis are installed by default. If you want to use emojis, install an emoji font:

```

docker

RUN apt-get install fonts-noto-color-emoji
```

```

docker

RUN apt-get install fonts-noto-color-emoji
```

## Japanese, Chinese, Korean, etc. [​](\#japanese-chinese-korean-etc "Direct link to Japanese, Chinese, Korean, etc.")

Those fonts may have limited Character support enabled by default. If you need full support, install the following fonts:

```

docker

RUN apt-get install fonts-noto-cjk
```

```

docker

RUN apt-get install fonts-noto-cjk
```

## Why are the packages not pinned? [​](\#why-are-the-packages-not-pinned "Direct link to Why are the packages not pinned?")

In Debian (and also Alpine), old packages are removed from the repositories once new versions are released. This means that pinning the versions will actually cause the Dockerfiles to break in the future. We choose Debian as the distribution because the packages get well tested before they get released into the repository.

## Notes for older versions [​](\#notes-for-older-versions "Direct link to Notes for older versions")

- If you are on a lower version than `v4.0.0`, add `ffmpeg` to the list of packages to install:



```

docker

RUN apt-get install -y nodejs ffmpeg npm chromium
```


```

docker

RUN apt-get install -y nodejs ffmpeg npm chromium
```

- If you are on Remotion `v3.3.80` or lower, tell Remotion where Chrome is installed:



```

docker

ENV PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium
```


```

docker

ENV PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium
```

## Recommendation: Don't use Alpine Linux [​](\#recommendation-dont-use-alpine-linux "Direct link to Recommendation: Don't use Alpine Linux")

Alpine Linux is a lightweight distribution often used in Docker. There are two known issues with it when used in conjunction with Remotion:

- The launch of the Rust parts of Remotion may be very slow (>10sec slowdown per render)
- If a new version of Chrome gets released in the registry, you might be unable to downgrade because old versions are not kept and breaking changes can not be ruled out.

## Changelog [​](\#changelog "Direct link to Changelog")

**November 6th, 2024**: Use node:22-bookworm-slim over node:20-bookworm to update to LTS and get a much smaller image.
**October 11th, 2023**: Used the node:20-bookworm, which is faster to deploy and also Debian.

**September 25th, 2023**: Recommend setting `enableMultiProcessOnLinux`.

**May 30th, 2023**: Update document for Remotion 4.0.

**April 15th, 2023**: Unpinning the versions in Debian since it would cause breakage.

**April 3rd, 2023**: Changed the Alpine Docker image to a Debian one, since the versions of Alpine packages cannot be pinned. This makes the Debian one less likely to break.

## See also [​](\#see-also "Direct link to See also")

- [Deploying the Remotion Studio](/docs/studio/deploy-server)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/docker.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Example](/docs/ssr-node) [Next\
\
Multiple cores on Linux](/docs/miscellaneous/linux-single-process)

- [Line-by-line](#line-by-line)
- [Example render script](#example-render-script)
- [Building the Docker image](#building-the-docker-image)
- [Giving access to the CPUs](#giving-access-to-the-cpus)
- [Emojis](#emojis)
- [Japanese, Chinese, Korean, etc.](#japanese-chinese-korean-etc)
- [Why are the packages not pinned?](#why-are-the-packages-not-pinned)
- [Notes for older versions](#notes-for-older-versions)
- [Recommendation: Don't use Alpine Linux](#recommendation-dont-use-alpine-linux)
- [Changelog](#changelog)
- [See also](#see-also)
