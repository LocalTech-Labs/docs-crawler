---
title: Entry point
source: Remotion Documentation
last_updated: 2024-12-22
---

# Entry point

- [Home page](/)
- [Terminology](/docs/terminology)
- Entry point

On this page

# Entry point

The entry point is the file where the [Remotion CLI](/docs/cli) and [`@remotion/renderer`](/docs/renderer) APIs will look for a Remotion project.

- By default in most templates, it is `src/index.ts`.
- In older projects, it may have an `.tsx` extension instead of `.ts`.
- The entry point can be passed to the render command, for example: `npx remotion render src/index.ts`.
- The entry point should call [`registerRoot()`](/docs/register-root).
- If you render a video using [`npx remotion render`](/docs/cli/render), the entry point is printed in grey.
- You can customize the default entry point in the config file using [`Config.setEntryPoint()`](/docs/config#setentrypoint).

## Which entry point is being used? [​](\#which-entry-point-is-being-used "Direct link to Which entry point is being used?")

If you call [`npx remotion render`](/docs/cli/render), the entry point and the reason why it was chosen will be printed. The algorithm is as follows:

[1](#1)

If the entry point is directly passed, for example `npx remotion render src/index.ts`, use that.

[2](#2) Otherwise, if a path was specified in the [Configuration file](/docs/config) using [`Config.setEntryPoint()`](/docs/config#setentrypoint), use that.

[3](#3) Go through the following common paths and if one exists, use that:

- `src/index.ts`
- `src/index.tsx`
- `src/index.js`
- `src/index.mjs`
- `remotion/index.tsx`
- `remotion/index.ts`
- `remotion/index.js`
- `remotion/index.mjs`

## In the Player [​](\#in-the-player "Direct link to In the Player")

In the [Remotion Player](/docs/terminology/player), there is no concept of an entry point.

You directly pass a React component and metadata to the [`<Player>`](/docs/player/player).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/entry-point.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Service Name](/docs/terminology/service-name) [Next\
\
Root file](/docs/terminology/root-file)

- [Which entry point is being used?](#which-entry-point-is-being-used)
- [In the Player](#in-the-player)
