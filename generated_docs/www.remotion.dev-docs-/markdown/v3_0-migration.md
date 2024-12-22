---
title: v3.0 Migration
source: Remotion Documentation
last_updated: 2024-12-22
---

# v3.0 Migration

- [Home page](/)
- Migration guides
- v3.0 Migration

On this page

# v3.0 Migration

When upgrading from Remotion 2 to Remotion 3, note the following changes and apply them to your project.

## How to upgrade [​](\#how-to-upgrade "Direct link to How to upgrade")

Upgrade `remotion` and all packages starting with `@remotion` to `^3.0.0`:

```

diff

- "remotion": "^2.6.15"
- "@remotion/bundler": "^2.6.15"
- "@remotion/eslint-config": "^2.6.15"
- "@remotion/eslint-plugin": "^2.6.15"
- "@remotion/cli": "^2.6.15"
- "@remotion/renderer": "^2.6.15"
+ "remotion": "^3.0.0"
+ "@remotion/bundler": "^3.0.0"
+ "@remotion/eslint-config": "^3.0.0"
+ "@remotion/eslint-plugin": "^3.0.0"
+ "@remotion/cli": "^3.0.0"
+ "@remotion/renderer": "^3.0.0"
```

```

diff

- "remotion": "^2.6.15"
- "@remotion/bundler": "^2.6.15"
- "@remotion/eslint-config": "^2.6.15"
- "@remotion/eslint-plugin": "^2.6.15"
- "@remotion/cli": "^2.6.15"
- "@remotion/renderer": "^2.6.15"
+ "remotion": "^3.0.0"
+ "@remotion/bundler": "^3.0.0"
+ "@remotion/eslint-config": "^3.0.0"
+ "@remotion/eslint-plugin": "^3.0.0"
+ "@remotion/cli": "^3.0.0"
+ "@remotion/renderer": "^3.0.0"
```

Run `npm i `, `yarn` or `pnpm i` respectively afterwards.

**Optional: Upgrade to React 18**

You need to upgrade both `react` and `react-dom`:

```

diff

- "react": "17.0.1"
- "react-dom": "17.0.1"
+ "react": "18.2.0"
+ "react-dom": "18.2.0"
```

```

diff

- "react": "17.0.1"
- "react-dom": "17.0.1"
+ "react": "18.2.0"
+ "react-dom": "18.2.0"
```

If you use TypeScript, update to the newest types as well:

```

diff

- "@types/react": "17.0.3"
+ "@types/react": "18.0.0"
```

```

diff

- "@types/react": "17.0.3"
+ "@types/react": "18.0.0"
```

**Optional: Upgrade to ESLint 8**

```

diff

- "eslint": "^7.15.0"
+ "eslint": "^8.13.0"
```

```

diff

- "eslint": "^7.15.0"
+ "eslint": "^8.13.0"
```

## Breaking changes [​](\#breaking-changes "Direct link to Breaking changes")

### Node 14 is required [​](\#node-14-is-required "Direct link to Node 14 is required")

Previously, we supported Node 12 and 13, but we no longer do.

**Upgrade path:** Upgrade to at least Node 14.

### Remotion is built with React 18 types [​](\#remotion-is-built-with-react-18-types "Direct link to Remotion is built with React 18 types")

The Remotion components adhere to [version 18 of `@types/react` and `@types/react-dom`](https://github.com/DefinitelyTyped/DefinitelyTyped/pull/56210). Minor type errors may occur that can be resolved by adjusting your components. We recommend that you upgrade the `@types/react` package yourself.

**Upgrade path:** [Update to React 18](/docs/react-18) \- optional but recommended.

### `renderFrames()`/ `renderStill()`: `compositionId` parameter removed [​](\#renderframesrenderstill-compositionid-parameter-removed "Direct link to renderframesrenderstill-compositionid-parameter-removed")

Instead, the composition ID is now embedded in the `composition` object (previously `config`).

**Upgrade path:** Remove the `compositionId` property from `renderFrames()`. Add the `composition.id` property to `renderFrames()` or pull the object from `getCompositions()`.

```

Previously
tsx

renderFrames({
  compositionId: "my-com",
  config: {
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Previously
tsx

renderFrames({
  compositionId: "my-com",
  config: {
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Now
tsx

renderFrames({
  composition: {
    id: "my-com",
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Now
tsx

renderFrames({
  composition: {
    id: "my-com",
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

### `renderFrames()`/ `renderStill()`: renamed `config` in to `composition` [​](\#renderframesrenderstill-renamed-config-in-to-composition "Direct link to renderframesrenderstill-renamed-config-in-to-composition")

The `config` parameter in `renderFrames()` and `renderStill()` was renamed to `composition` and now requires an additional `id` property.

```

Previously
tsx

renderFrames({
  compositionId: "my-com",
  config: {
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Previously
tsx

renderFrames({
  compositionId: "my-com",
  config: {
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Now
tsx

renderFrames({
  composition: {
    id: "my-com",
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

```

Now
tsx

renderFrames({
  composition: {
    id: "my-com",
    width: 1920,
    height: 1080,
    fps: 30,
    durationInFrames: 300,
  },
});
```

**Upgrade path:** Rename the `config` property of `renderFrames()` and `renderStill()`. Add the `composition.id` property to or pull the object from `getCompositions()`.

### `renderFrames()`/ `renderStill()`/ `getCompositions()`: Errors thrown in your app make the render fail [​](\#renderframesrenderstillgetcompositions-errors-thrown-in-your-app-make-the-render-fail "Direct link to renderframesrenderstillgetcompositions-errors-thrown-in-your-app-make-the-render-fail")

Previously, you could catch errors being thrown in your Remotion code using the `onError` property of `getCompositions()`, `renderFrames()` and `renderStill()`.

The new behavior of Remotion 3.0 is that if an error occurs, these functions reject instead.

**Upgrade path**: Remove the `onError` property from your `getCompositions()`, `renderFrames()` and `renderStill()` calls and catch errors in a try / catch instead. Eliminate any error being thrown in your application.

```

Previously
tsx

await renderStill({
  // ...
  output: "/tmp/still.png",
  onError: (err) => {
    console.log("Error occured in browser", err);
  },
});
```

```

Previously
tsx

await renderStill({
  // ...
  output: "/tmp/still.png",
  onError: (err) => {
    console.log("Error occured in browser", err);
  },
});
```

```

Now
tsx

try {
  await renderStill({
    // ...
    output: "/tmp/still.png",
  });
} catch (err) {
  console.log("Error occured in browser", err);
}
```

```

Now
tsx

try {
  await renderStill({
    // ...
    output: "/tmp/still.png",
  });
} catch (err) {
  console.log("Error occured in browser", err);
}
```

### `renderFrames()`/ `renderStill()`/ `getCompositions()`: Renamed `webpackBundle` to `serveUrl` [​](\#renderframesrenderstillgetcompositions-renamed-webpackbundle-to-serveurl "Direct link to renderframesrenderstillgetcompositions-renamed-webpackbundle-to-serveurl")

The `webpackBundle` property was renamed to `serveUrl` to reflect the fact that now a URL is also supported. Functionally, the same behavior is supported as before.

**Upgrade path**: Rename `webpackBundle` to `serveUrl`.

```

Previously
tsx

await renderStill({
  output: "/tmp/still.png",
  webpackBundle: "/tmp/react-motion-graphics8zfs9d/index.html",
  // ...
});
```

```

Previously
tsx

await renderStill({
  output: "/tmp/still.png",
  webpackBundle: "/tmp/react-motion-graphics8zfs9d/index.html",
  // ...
});
```

```

Now
tsx

await renderStill({
  output: "/tmp/still.png",
  serveUrl: "/tmp/react-motion-graphics8zfs9d/index.html",
  // ...
});
```

```

Now
tsx

await renderStill({
  output: "/tmp/still.png",
  serveUrl: "/tmp/react-motion-graphics8zfs9d/index.html",
  // ...
});
```

### `stitchFramesToVideo()`: removed `imageFormat`, `parallelism` and `webpackBundle` from API [​](\#stitchframestovideo-removed-imageformat-parallelism-and-webpackbundle-from-api "Direct link to stitchframestovideo-removed-imageformat-parallelism-and-webpackbundle-from-api")

- The parameters `imageFormat` and `webpackBundle` are not necessary anymore to pass into `stitchFramesToVideo()`, because the necessary information is now embedded in `assetsInfo`, retrieved by calling `renderFrames()`.
- The `parallelism` parameter did probably not to what you thought it would. To avoid any confusion, we removed it without any replacement.

**Upgrade path**: Remove the `imageFormat`, `parallelism` and `webpackBundle` parameters from `stitchFramesToVideo()`.

### `useVideoConfig()`: returns additional `id` and `defaultProps` properties [​](\#usevideoconfig-returns-additional-id-and-defaultprops-properties "Direct link to usevideoconfig-returns-additional-id-and-defaultprops-properties")

The [`useVideoConfig()`](/docs/use-video-config) hook now returns two additional properties:

```

diff

{
  "width": 1920,
  "height": 1080,
  "fps": 30,
  "durationInFrames": 30
+ "id": "my-comp",
+ "defaultProps": {}
}
```

```

diff

{
  "width": 1920,
  "height": 1080,
  "fps": 30,
  "durationInFrames": 30
+ "id": "my-comp",
+ "defaultProps": {}
}
```

**Upgrade path**: Ensure you don't rely on the new properties not being there.

### `getCompositions()`: renamed `browserInstance` to `puppeteerInstance` [​](\#getcompositions-renamed-browserinstance-to-puppeteerinstance "Direct link to getcompositions-renamed-browserinstance-to-puppeteerinstance")

To align the name with `renderStill()`, `renderFrames()` and `renderMedia()`, `browserInstance` was renamed to `puppeteerInstance`.

### `overrideWebpackConfig()`: Config file option is removed [​](\#overridewebpackconfig-config-file-option-is-removed "Direct link to overridewebpackconfig-config-file-option-is-removed")

The import `overrideWebpackConfig` was deprecated and now removed. Use `Config.overrideWebpackConfig()` instead.

**Upgrade path**: Change import in `remotion.config.ts` and use `overrideWebpackConfig()`.

```

Previously
tsx

```

```

Previously
tsx

```

```

Now
tsx

import { Config } from "remotion";

Config.overrideWebpackConfig();
```

```

Now
tsx

import { Config } from "remotion";

Config.overrideWebpackConfig();
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/3-0-migration.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
v4.0 Migration](/docs/4-0-migration) [Next\
\
v2.0 Migration](/docs/2-0-migration)

- [How to upgrade](#how-to-upgrade)
- [Breaking changes](#breaking-changes)
  - [Node 14 is required](#node-14-is-required)
  - [Remotion is built with React 18 types](#remotion-is-built-with-react-18-types)
  - [`renderFrames()`/ `renderStill()`: `compositionId` parameter removed](#renderframesrenderstill-compositionid-parameter-removed)
  - [`renderFrames()`/ `renderStill()`: renamed `config` in to `composition`](#renderframesrenderstill-renamed-config-in-to-composition)
  - [`renderFrames()`/ `renderStill()`/ `getCompositions()`: Errors thrown in your app make the render fail](#renderframesrenderstillgetcompositions-errors-thrown-in-your-app-make-the-render-fail)
  - [`renderFrames()`/ `renderStill()`/ `getCompositions()`: Renamed `webpackBundle` to `serveUrl`](#renderframesrenderstillgetcompositions-renamed-webpackbundle-to-serveurl)
  - [`stitchFramesToVideo()`: removed `imageFormat`, `parallelism` and `webpackBundle` from API](#stitchframestovideo-removed-imageformat-parallelism-and-webpackbundle-from-api)
  - [`useVideoConfig()`: returns additional `id` and `defaultProps` properties](#usevideoconfig-returns-additional-id-and-defaultprops-properties)
  - [`getCompositions()`: renamed `browserInstance` to `puppeteerInstance`](#getcompositions-renamed-browserinstance-to-puppeteerinstance)
  - [`overrideWebpackConfig()`: Config file option is removed](#overridewebpackconfig-config-file-option-is-removed)
