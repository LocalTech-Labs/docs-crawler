---
title: Resize a video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Resize a video

- [Home page](/)
- [WebCodecs](/docs/webcodecs/)
- Resize a video

On this page

# Resize a video

You can resize a video in the browser using [`convertMedia()`](/docs/webcodecs/convert-media) function from the [`@remotion/webcodecs`](/docs/webcodecs) package.

💼 Important License Disclaimer

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

🚧 Unstable API

This package is experimental.

We might change the API at any time, until we remove this notice.

## Simple Example [​](\#simple-example "Direct link to Simple Example")

```

Scale a video down to 480p
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'mp4',
  resize: {
    mode: 'max-height',
    maxHeight: 480,
  },
});
```

```

Scale a video down to 480p
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'mp4',
  resize: {
    mode: 'max-height',
    maxHeight: 480,
  },
});
```

## Resize modes [​](\#resize-modes "Direct link to Resize modes")

For the [`resize`](/docs/webcodecs/convert-media#resize) value, you can specify the following modes:

### `max-height` [​](\#max-height "Direct link to max-height")

```

Scale a video down to 480p
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-height',
  maxHeight: 480,
};
```

```

Scale a video down to 480p
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-height',
  maxHeight: 480,
};
```

Scales a video down so it is at most `maxHeight` pixels tall.

### `max-width` [​](\#max-width "Direct link to max-width")

```

Scale a video down to 640 pixels wide
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-width',
  maxWidth: 640,
};
```

```

Scale a video down to 640 pixels wide
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-width',
  maxWidth: 640,
};
```

Scales a video down so it is at most `maxWidth` pixels wide.

### `width` [​](\#width "Direct link to width")

```

Scale a video to 640 pixels wide
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'width',
  width: 640,
};
```

```

Scale a video to 640 pixels wide
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'width',
  width: 640,
};
```

Scales a video to exactly `width` pixels wide.

### `height` [​](\#height "Direct link to height")

```

Scale a video to 480 pixels tall
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'height',
  height: 480,
};
```

```

Scale a video to 480 pixels tall
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'height',
  height: 480,
};
```

Scales a video to exactly `height` pixels tall.

### `max-height-width` [​](\#max-height-width "Direct link to max-height-width")

```

Scale a video down to 480px height or 640 pixels width
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-height-width',
  maxHeight: 480,
  maxWidth: 640,
};
```

```

Scale a video down to 480px height or 640 pixels width
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'max-height-width',
  maxHeight: 480,
  maxWidth: 640,
};
```

Scales a video down so it is at most `maxHeight` pixels tall or `maxWidth` pixels wide.

The video will be scaled down to the biggest size that fits both constraints.

### `scale` [​](\#scale "Direct link to scale")

```

Scale a video to 50%
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'scale',
  scale: 0.5,
};
```

```

Scale a video to 50%
tsx

import {ResizeOperation} from '@remotion/webcodecs';

const resize: ResizeOperation = {
  mode: 'scale',
  scale: 0.5,
};
```

Scales a video by a factor of `scale`. Factor must be greater than `0`.

## Order of operations [​](\#order-of-operations "Direct link to Order of operations")

If you apply both a [`rotate`](/docs/webcodecs/rotate-a-video) and a `resize` operation, the `rotate` operation will be applied first.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/webcodecs`](/docs/webcodecs)
- [Rotate a video](/docs/webcodecs/rotate-a-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/resizing.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Rotate a video](/docs/webcodecs/rotate-a-video) [Next\
\
Fixing a MediaRecorder video](/docs/webcodecs/fix-mediarecorder-video)

- [Simple Example](#simple-example)
- [Resize modes](#resize-modes)
  - [`max-height`](#max-height)
  - [`max-width`](#max-width)
  - [`width`](#width)
  - [`height`](#height)
  - [`max-height-width`](#max-height-width)
  - [`scale`](#scale)
- [Order of operations](#order-of-operations)
- [See also](#see-also)
