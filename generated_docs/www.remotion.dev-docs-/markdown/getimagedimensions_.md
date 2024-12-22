---
title: getImageDimensions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getImageDimensions()

- [Home page](/)
- [@remotion/media-utils](/docs/media-utils/)
- getImageDimensions()

On this page

# getImageDimensions()

_Part of the `@remotion/media-utils` package of helper functions. Available from v4.0.143._

Takes an image `src`, retrieves the dimensions of an image.

## Arguments [​](\#arguments "Direct link to Arguments")

### `src` [​](\#src "Direct link to src")

A string that specifies the URL or path of the image.

## Return value [​](\#return-value "Direct link to Return value")

_`Promise<ImageDimensions>`_

An object with information about the image dimensions:

### `width` [​](\#width "Direct link to width")

_number_

The image width, in pixels (px).

### `height` [​](\#height "Direct link to height")

_number_

The image height, in pixels (px).

## Example [​](\#example "Direct link to Example")

```

ts

import { getImageDimensions } from "@remotion/media-utils";

const { width, height } = await getImageDimensions(
  "https://example.com/remote-image.png",
);
console.log(width, height);
```

```

ts

import { getImageDimensions } from "@remotion/media-utils";

const { width, height } = await getImageDimensions(
  "https://example.com/remote-image.png",
);
console.log(width, height);
```

## Caching behavior [​](\#caching-behavior "Direct link to Caching behavior")

This function is memoizing the results it returns.

If you pass in the same argument to `src` multiple times, it will return a cached version from the second time on, regardless of if the file has changed.

To clear the cache, you have to reload the page.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-utils/src/get-image-dimensions.ts)
- [Preload image](/docs/preload/preload-image)
- [`<Img />`](/docs/img)
- [`getAudioData()`](/docs/get-audio-data)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-image-dimensions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
visualizeAudio()](/docs/visualize-audio) [Next\
\
@remotion/motion-blur](/docs/motion-blur/)

- [Arguments](#arguments)
  - [`src`](#src)
- [Return value](#return-value)
  - [`width`](#width)
  - [`height`](#height)
- [Example](#example)
- [Caching behavior](#caching-behavior)
- [See also](#see-also)
