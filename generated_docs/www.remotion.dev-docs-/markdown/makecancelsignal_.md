---
title: makeCancelSignal()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeCancelSignal()

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- makeCancelSignal()

On this page

# makeCancelSignal()

_Available from v3.0.15_

This function returns a signal and a cancel function that allows to you cancel a render triggered using [`renderMedia()`](/docs/renderer/render-media), [`renderStill()`](/docs/renderer/render-still), [`renderFrames()`](/docs/renderer/render-frames) or [`stitchFramesToVideo()`](/docs/renderer/stitch-frames-to-video)
.

## Example [​](\#example "Direct link to Example")

```

tsx

import { makeCancelSignal, renderMedia } from "@remotion/renderer";

const { cancelSignal, cancel } = makeCancelSignal();

// Note that no `await` is used yet
const render = renderMedia({
  composition,
  codec: "h264",
  serveUrl: "https://silly-crostata-c4c336.netlify.app/",
  outputLocation: "out/render.mp4",
  cancelSignal,
});

// Cancel render after 10 seconds
setTimeout(() => {
  cancel();
}, 10000);

// If the render completed within 10 seconds, renderMedia() will resolve
await render;

// If the render did not complete, renderMedia() will reject
// ==> "[Error: renderMedia() got cancelled]"
```

```

tsx

import { makeCancelSignal, renderMedia } from "@remotion/renderer";

const { cancelSignal, cancel } = makeCancelSignal();

// Note that no `await` is used yet
const render = renderMedia({
  composition,
  codec: "h264",
  serveUrl: "https://silly-crostata-c4c336.netlify.app/",
  outputLocation: "out/render.mp4",
  cancelSignal,
});

// Cancel render after 10 seconds
setTimeout(() => {
  cancel();
}, 10000);

// If the render completed within 10 seconds, renderMedia() will resolve
await render;

// If the render did not complete, renderMedia() will reject
// ==> "[Error: renderMedia() got cancelled]"
```

## API [​](\#api "Direct link to API")

Calling `makeCancelSignal` returns an object with two properties:

- `cancelSignal`: An object to be passed to one of the above mentioned render functions
- `cancel`: A function you should call when you want to cancel the render.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/renderer/src/make-cancel-signal.ts)
- [`renderMedia()`](/docs/renderer/render-media)
- [`renderStill()`](/docs/renderer/render-still)
- [`renderFrames()`](/docs/renderer/render-frames)
- [`stitchFramesToVideo()`](/docs/renderer/stitch-frames-to-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/make-cancel-signal.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
ensureBrowser()](/docs/renderer/ensure-browser) [Next\
\
ensureFfmpeg()](/docs/renderer/ensure-ffmpeg)

- [Example](#example)
- [API](#api)
- [See also](#see-also)
