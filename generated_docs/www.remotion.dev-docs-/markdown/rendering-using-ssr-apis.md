---
title: Rendering using SSR APIs
source: Remotion Documentation
last_updated: 2024-12-22
---

# Rendering using SSR APIs

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- Example

On this page

# Rendering using SSR APIs

The NPM package [`@remotion/renderer`](/docs/renderer) provides you with " **S** erver- **S** ide **R** endering" APIs for rendering media programmatically.

These functions can be used in Node.js and Bun.

Rendering a video takes three steps:

[1](#1)

Creating a [Remotion Bundle](/docs/terminology/bundle)

[2](#2)

Select the composition to render and calculate its metadata

[3](#3)

Render the video, audio, still or image sequence.

## Example script [​](\#example-script "Direct link to Example script")

Follow this commented example to see how to render a video:

```

render.mjs
tsx

import {bundle} from '@remotion/bundler';
import {renderMedia, selectComposition} from '@remotion/renderer';
import path from 'path';

// The composition you want to render
const compositionId = 'HelloWorld';

// You only have to create a bundle once, and you may reuse it
// for multiple renders that you can parametrize using input props.
const bundleLocation = await bundle({
  entryPoint: path.resolve('./src/index.ts'),
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: (config) => config,
});

// Parametrize the video by passing props to your component.
const inputProps = {
  foo: 'bar',
};

// Get the composition you want to render. Pass `inputProps` if you
// want to customize the duration or other metadata.
const composition = await selectComposition({
  serveUrl: bundleLocation,
  id: compositionId,
  inputProps,
});

// Render the video. Pass the same `inputProps` again
// if your video is parametrized with data.
await renderMedia({
  composition,
  serveUrl: bundleLocation,
  codec: 'h264',
  outputLocation: `out/${compositionId}.mp4`,
  inputProps,
});

console.log('Render done!');
```

```

render.mjs
tsx

import {bundle} from '@remotion/bundler';
import {renderMedia, selectComposition} from '@remotion/renderer';
import path from 'path';

// The composition you want to render
const compositionId = 'HelloWorld';

// You only have to create a bundle once, and you may reuse it
// for multiple renders that you can parametrize using input props.
const bundleLocation = await bundle({
  entryPoint: path.resolve('./src/index.ts'),
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: (config) => config,
});

// Parametrize the video by passing props to your component.
const inputProps = {
  foo: 'bar',
};

// Get the composition you want to render. Pass `inputProps` if you
// want to customize the duration or other metadata.
const composition = await selectComposition({
  serveUrl: bundleLocation,
  id: compositionId,
  inputProps,
});

// Render the video. Pass the same `inputProps` again
// if your video is parametrized with data.
await renderMedia({
  composition,
  serveUrl: bundleLocation,
  codec: 'h264',
  outputLocation: `out/${compositionId}.mp4`,
  inputProps,
});

console.log('Render done!');
```

This flow is customizable. Click on one of the SSR APIs to read about its options:

- [`getCompositions()`](/docs/renderer/get-compositions) \- Get a list of available compositions from a Remotion project.
- [`selectComposition()`](/docs/renderer/select-composition) \- Get a list of available compositions from a Remotion project.
- [`renderMedia()`](/docs/renderer/render-media) \- Render a video or audio.
- [`renderFrames()`](/docs/renderer/render-frames) \- Render an image sequence.
- [`renderStill()`](/docs/renderer/render-still) \- Render a still image.
- [`stitchFramesToVideo()`](/docs/renderer/stitch-frames-to-video) \- Encode a video based on an image sequence
- [`openBrowser()`](/docs/renderer/open-browser) \- Share a browser instance across function calls for improved performance.

## Linux Dependencies [​](\#linux-dependencies "Direct link to Linux Dependencies")

If you are on Linux, Chrome Headless Shell requires some shared libraries to be installed. See [Linux Dependencies](/docs/miscellaneous/linux-dependencies).

## SSR APIs in Next.js [​](\#ssr-apis-in-nextjs "Direct link to SSR APIs in Next.js")

If you are using Next.js, you will not be able to use `@remotion/bundler` because of the limitations explained in [Can I render videos in Next.js?](https://www.remotion.dev/docs/miscellaneous/vercel-functions#can-i-render-videos-in-nextjs) Refer to the page for possible alternatives.

We recommend [Lambda](/docs/lambda) for use in Next.js.

## See also [​](\#see-also "Direct link to See also")

- [Server-side rendering](/docs/ssr)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/ssr-node.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/ssr) [Next\
\
Docker](/docs/docker)

- [Example script](#example-script)
- [Linux Dependencies](#linux-dependencies)
- [SSR APIs in Next.js](#ssr-apis-in-nextjs)
- [See also](#see-also)