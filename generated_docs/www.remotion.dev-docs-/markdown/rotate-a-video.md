---
title: Rotate a video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Rotate a video

- [Home page](/)
- [WebCodecs](/docs/webcodecs/)
- Rotate a video

On this page

# Rotate a video

You can rotate a video in the browser using the [`@remotion/webcodecs`](/docs/webcodecs) package to fix a video that has a bad orientation.

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

Rotate 90 degrees clockwise
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  rotate: 90,
});
```

```

Rotate 90 degrees clockwise
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  rotate: 90,
});
```

All rotation values that are multiples of 90 are supported. The video will be rotated clockwise.

## Videos with rotation [​](\#videos-with-rotation "Direct link to Videos with rotation")

Videos that have rotation metadata embedded will by default be rotated when re-encoded in order to produce a video with the correct orientation.

This means you should not try to detect rotation metadata and apply the correction yourself, because it will be done automatically.

If you supply a rotation, it will be in addition to the automatic rotation correction that [`convertMedia()`](/docs/webcodecs/convert-media) applies by default.

## Disabling automatic rotation [​](\#disabling-automatic-rotation "Direct link to Disabling automatic rotation")

If you want the video to not automatically be rotated, you need to re-apply the rotation that is the same as the rotation metadata in the video.

The rotations will negate each other, and [`convertMedia()`](/docs/webcodecs/convert-media) will not execute any code to apply rotation.

```

Prevent automatic video rotation
tsx

import {convertMedia, defaultOnVideoTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onVideoTrack: async (params) => {
    const action = await defaultOnVideoTrackHandler(params);

    if (action.type !== 'reencode') {
      return action;
    }

    return {
      ...action,
      rotate: params.track.rotation,
    };
  },
});
```

```

Prevent automatic video rotation
tsx

import {convertMedia, defaultOnVideoTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onVideoTrack: async (params) => {
    const action = await defaultOnVideoTrackHandler(params);

    if (action.type !== 'reencode') {
      return action;
    }

    return {
      ...action,
      rotate: params.track.rotation,
    };
  },
});
```

See [Track Transformation](/docs/webcodecs/track-transformation) for more information on how to customize the video track transformation.

## Order of operations [​](\#order-of-operations "Direct link to Order of operations")

If you apply both a `rotate` and a [`resize`](/docs/webcodecs/resize-a-video) operation, the `rotate` operation will be applied first.

## Reference implementation [​](\#reference-implementation "Direct link to Reference implementation")

Use [remotion.dev/rotate](https://remotion.dev/rotate) to rotate a video online using WebCodecs.

See the [source code](https://github.com/remotion-dev/remotion/tree/main/packages/convert) for a reference on how to implement it.

## See also [​](\#see-also "Direct link to See also")

- [`convertMedia()`](/docs/webcodecs/convert-media)
- [Resize a video](/docs/webcodecs/resize-a-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/rotation.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Convert a video](/docs/webcodecs/convert-a-video) [Next\
\
Resize a video](/docs/webcodecs/resize-a-video)

- [Simple Example](#simple-example)
- [Videos with rotation](#videos-with-rotation)
- [Disabling automatic rotation](#disabling-automatic-rotation)
- [Order of operations](#order-of-operations)
- [Reference implementation](#reference-implementation)
- [See also](#see-also)
