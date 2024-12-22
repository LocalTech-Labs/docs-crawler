---
title: @remotion/renderer
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/renderer

- [Home page](/)
- @remotion/renderer

On this page

# @remotion/renderer

The `@remotion/renderer` package provides APIs for rendering video server-side.
The package is also internally used by the Remotion CLI and [Remotion Lambda](/docs/lambda).

warning

The configuration file has no effect when using these APIs.

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/renderer@4.0.242Copy
```

```

npm i --save-exact @remotion/renderer@4.0.242Copy
```

```

pnpm i @remotion/renderer@4.0.242Copy
```

```

pnpm i @remotion/renderer@4.0.242Copy
```

```

bun i @remotion/renderer@4.0.242Copy
```

```

bun i @remotion/renderer@4.0.242Copy
```

```

yarn --exact add @remotion/renderer@4.0.242Copy
```

```

yarn --exact add @remotion/renderer@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## Server-side rendering examples [​](\#server-side-rendering-examples "Direct link to Server-side rendering examples")

See the [Server-side rendering](/docs/ssr) for some examples of how to use server-side rendering.

## Available functions [​](\#available-functions "Direct link to Available functions")

The following APIs are available in the `@remotion/renderer` package:

[**getCompositions()** \
\
List available compositions](/docs/renderer/get-compositions) [**selectComposition()** \
\
Get a composition](/docs/renderer/select-composition) [**renderMedia()** \
\
Render a video or audio](/docs/renderer/render-media) [**renderFrames()** \
\
Render a series of images](/docs/renderer/render-frames) [**renderStill()** \
\
Render a single image](/docs/renderer/render-still) [**stitchFramesToVideo()** \
\
Turn images into a video](/docs/renderer/stitch-frames-to-video) [**openBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/open-browser) [**ensureBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/ensure-browser) [**makeCancelSignal()** \
\
Create token to later cancel a render](/docs/renderer/make-cancel-signal) [**getVideoMetadata()** \
\
Get metadata from a video file in Node.js](/docs/renderer/get-video-metadata) [**getSilentParts()** \
\
Obtain silent portions of a video or audio](/docs/renderer/get-silent-parts) [**ensureFfmpeg()** \
\
Check for ffmpeg binary and install if not existing](/docs/renderer/ensure-ffmpeg) [**ensureFfprobe()** \
\
Check for ffprobe binary and install if not existing](/docs/renderer/ensure-ffprobe) [**getCanExtractFramesFast()** \
\
Probes for fast extraction for <OffthreadVideo>](/docs/renderer/get-can-extract-frames-fast)

## What's the difference between `renderMedia()` and `renderFrames()`? [​](\#whats-the-difference-between-rendermedia-and-renderframes "Direct link to whats-the-difference-between-rendermedia-and-renderframes")

In Remotion 3.0, we added the [`renderMedia()`](/docs/renderer/render-media) API which combines `renderFrames()` and `stitchFramesToVideo()` into one simplified step and performs the render faster. Prefer `renderMedia()` if you can.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
bundle()](/docs/bundle) [Next\
\
getCompositions()](/docs/renderer/get-compositions)

- [Installation](#installation)
- [Server-side rendering examples](#server-side-rendering-examples)
- [Available functions](#available-functions)
- [What's the difference between `renderMedia()` and `renderFrames()`?](#whats-the-difference-between-rendermedia-and-renderframes)
