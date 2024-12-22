---
title: @remotion/three
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/three

- [Home page](/)
- @remotion/three

On this page

# @remotion/three

is a package for integrating [React Three Fiber](https://github.com/pmndrs/react-three-fiber) with Remotion.

- [`<ThreeCanvas />`](/docs/three-canvas) will allow you to use `useCurrentFrame()` and other Remotion hooks within a R3F Canvas. Animations are now not inside a `useFrame()` hook, but directly rendered into the markup.

- [`useVideoTexture()`](/docs/use-video-texture) allows you to use a Remotion [`<Video />`](/docs/video) as a texture map.

- [`useOffthreadVideoTexture()`](/docs/use-offthread-video-texture) is an equivalent to [`useVideoTexture()`](/docs/use-video-texture), but displays the exact frame as an `Three.ImageTexture` during rendering, simillarly to [`<OffthreadVideo />`](/docs/offthreadvideo).

These are the only two APIs provided - for everything else you can use the standard [React Three Fiber](https://github.com/pmndrs/react-three-fiber) APIs.

## Starter template [​](\#starter-template "Direct link to Starter template")

Check out [remotion-template-three](https://github.com/remotion-dev/template-three), a minimal boilerplate for Remotion and React Three Fiber. It is a template repository, you can click "Use this template" on the GitHub repo to get started.

0:00 / 0:10

The template features a 3D phone with a video inside which you can effortlessly swap out. Just as easily, you can change properties like the color, size, thickness, corner radius of the phone.

The template serves as a soft introduction on how to use `<ThreeCanvas />` and `useVideoTexture()`. You can easily delete everything inside the canvas to start working on a different idea.

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm

```

bash

npm i three @react-three/fiber @remotion/three @types/three
```

```

bash

npm i three @react-three/fiber @remotion/three @types/three
```

```

bash

yarn add three @react-three/fiber @remotion/three @types/three
```

```

bash

yarn add three @react-three/fiber @remotion/three @types/three
```

```

bash

pnpm i three @react-three/fiber @remotion/three @types/three
```

```

bash

pnpm i three @react-three/fiber @remotion/three @types/three
```

You are now set up and can render a [`<ThreeCanvas />`](/docs/three-canvas) in your project.

## Note on `<Sequence>` [​](\#note-on-sequence "Direct link to note-on-sequence")

A [`<Sequence>`](/docs/sequence) by default will return a `<div>` component, which is not allowed inside a `<ThreeCanvas>`. To avoid an error, pass `layout="none"` to `<Sequence>`.

## Note on server-side rendering [​](\#note-on-server-side-rendering "Direct link to Note on server-side rendering")

Three.JS does not render with the default OpenGL renderer - we recommend to set it to `angle`. The config file of new projects includes by default:

```

ts

import {Config} from '@remotion/cli/config';

Config.setChromiumOpenGlRenderer('angle');
```

```

ts

import {Config} from '@remotion/cli/config';

Config.setChromiumOpenGlRenderer('angle');
```

Since the config file does not apply to server-side rendering, you need to explicitly add

```

json

"chromiumOptions": {
  "gl": "angle"
}
```

```

json

"chromiumOptions": {
  "gl": "angle"
}
```

to server-side rendering APIs like [`renderMedia()`](/docs/renderer/render-media), [`renderFrames()`](/docs/renderer/render-frames), [`getCompositions()`](/docs/renderer/get-compositions) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda).

## Thanks [​](\#thanks "Direct link to Thanks")

A big thanks to [Björn Zeutzheim](https://github.com/olee) for researching and discovering the techniques needed for React Three Fiber integration and for doing the initial implementation of the @remotion/three APIs.

## APIs [​](\#apis "Direct link to APIs")

[**<ThreeCanvas>** \
\
A wrapper for React Three Fiber' Canvas](/docs/three-canvas) [**useVideoTexture(** \
\
Use a video in React Three Fiber](/docs/use-video-texture) [**useOffthreadVideoTexture()** \
\
Use an <OffthreadVideo> in React Three Fiber](/docs/use-offthread-video-texture)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/three.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
enableScss()](/docs/enable-scss/enable-scss) [Next\
\
<ThreeCanvas>](/docs/three-canvas)

- [Starter template](#starter-template)
- [Installation](#installation)
- [Note on `<Sequence>`](#note-on-sequence)
- [Note on server-side rendering](#note-on-server-side-rendering)
- [Thanks](#thanks)
- [APIs](#apis)
