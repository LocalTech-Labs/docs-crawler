---
title: @remotion/skia
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/skia

- [Home page](/)
- @remotion/skia

On this page

# @remotion/skia

This package provides utilities useful for integrating [React Native Skia](https://github.com/Shopify/react-native-skia) with Remotion.

## Installation [​](\#installation "Direct link to Installation")

Install, `@remotion/skia` as well as `@shopify/react-native-skia`.

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

npm i --save-exact @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

pnpm i @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

pnpm i @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

bun i @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

bun i @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

yarn --exact add @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

```

yarn --exact add @remotion/skia@4.0.242 @shopify/react-native-skiaCopy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

note

Since Remotion `v3.3.93` and React Native Skia `0.1.191`, `react-native-web` is not a dependency anymore. You may remove it from your project.

Also update **all the other Remotion packages** to have the same version: `remotion`, `@remotion/cli` and others.

note

Make sure no package version number has a `^` character in front of it as it can lead to a version conflict.

[Override the Webpack config](/docs/webpack) by using [`enableSkia()`](/docs/skia/enable-skia).

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableSkia} from '@remotion/skia/enable';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia(currentConfiguration);
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableSkia} from '@remotion/skia/enable';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia(currentConfiguration);
});
```

note

Prior to `v3.3.39`, the option was called `Config.Bundling.overrideWebpackConfig()`.

Next, you need to refactor the [entry point](/docs/terminology/entry-point) file to first load the Skia WebAssembly binary before calling registerRoot():

```

src/index.ts
ts

import { LoadSkia } from "@shopify/react-native-skia/src/web";
import { registerRoot } from "remotion";

(async () => {
  await LoadSkia();
  const { RemotionRoot } = await import("./Root");
  registerRoot(RemotionRoot);
})();
```

```

src/index.ts
ts

import { LoadSkia } from "@shopify/react-native-skia/src/web";
import { registerRoot } from "remotion";

(async () => {
  await LoadSkia();
  const { RemotionRoot } = await import("./Root");
  registerRoot(RemotionRoot);
})();
```

You can now start using the [`<SkiaCanvas>`](/docs/skia/skia-canvas) in your Remotion project.

## Templates [​](\#templates "Direct link to Templates")

You can find the [starter template](https://github.com/remotion-dev/template-skia) here or install it using:

- npm
- bun
- pnpm
- yarn

```

bash

npx create-video --skia
```

```

bash

npx create-video --skia
```

```

bash

bun create video --skia
```

```

bash

bun create video --skia
```

```

bash

yarn create video --skia
```

```

bash

yarn create video --skia
```

```

bash

pnpm create video --skia
```

```

bash

pnpm create video --skia
```

## Rendering [​](\#rendering "Direct link to Rendering")

By default Remotion rendering are done on the CPU. Some Skia effects rely on advanced GPU features, which may be slow to run on the CPU depending on the kind of effect you are using. If your Skia export is extremely slow, we found that enabling the GPU via the `--gl=angle` option improves things substantially. Please check out the documentation on [GPU rendering](/docs/gpu).

```

sh

remotion render Main out/video.mp4 --gl=angle
```

```

sh

remotion render Main out/video.mp4 --gl=angle
```

## Resources [​](\#resources "Direct link to Resources")

- [Example project by William Candillon](https://github.com/wcandillon/remotion-skia-tutorial)
- [Tutorial for the example project](https://www.youtube.com/watch?v=-7MOoWN2_nk)

## APIs [​](\#apis "Direct link to APIs")

[**enableSkia()** \
\
Webpack override for enabling Skia](/docs/skia/enable-skia) [**<SkiaCanvas>** \
\
React Native Skia <Canvas> wrapper](/docs/skia/skia-canvas)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/skia/skia.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useOffthreadVideoTexture()](/docs/use-offthread-video-texture) [Next\
\
enableSkia()](/docs/skia/enable-skia)

- [Installation](#installation)
- [Templates](#templates)
- [Rendering](#rendering)
- [Resources](#resources)
- [APIs](#apis)
