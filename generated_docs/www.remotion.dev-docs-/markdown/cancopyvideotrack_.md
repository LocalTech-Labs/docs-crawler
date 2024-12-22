---
title: canCopyVideoTrack()
source: Remotion Documentation
last_updated: 2024-12-22
---

# canCopyVideoTrack()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- canCopyVideoTrack()

On this page

# canCopyVideoTrack()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

🚧 Unstable API

This package is experimental.

We might change the API at any time, until we remove this notice.

Given a `VideoTrack`, determine if it can be copied to the output without re-encoding.

You can obtain a `VideoTrack` using [`parseMedia()`](/docs/media-parser/parse-media) or during the conversion process using the [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) callback of [`convertMedia()`](/docs/webcodecs/convert-media).

## Examples [​](\#examples "Direct link to Examples")

```

Check if a video tracks can be copied
tsx

import {parseMedia} from '@remotion/media-parser';
import {canCopyVideoTrack} from '@remotion/webcodecs';

const {tracks, container} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.webm',
  fields: {
    tracks: true,
    container: true,
  },
});

for (const track of tracks.videoTracks) {
  canCopyVideoTrack({
    outputContainer: 'webm',
    inputTrack: track,
    inputContainer: container,
    rotationToApply: 0,
    resizeOperation: null,
  }); // boolean
}
```

```

Check if a video tracks can be copied
tsx

import {parseMedia} from '@remotion/media-parser';
import {canCopyVideoTrack} from '@remotion/webcodecs';

const {tracks, container} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.webm',
  fields: {
    tracks: true,
    container: true,
  },
});

for (const track of tracks.videoTracks) {
  canCopyVideoTrack({
    outputContainer: 'webm',
    inputTrack: track,
    inputContainer: container,
    rotationToApply: 0,
    resizeOperation: null,
  }); // boolean
}
```

```

Copy a video track to VP8, otherwise drop it
tsx

import {convertMedia, canCopyVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.webm',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onVideoTrack: async ({track, inputContainer, outputContainer}) => {
    const canCopy = canCopyVideoTrack({
      outputContainer,
      inputTrack: track,
      inputContainer,
      rotationToApply: 0,
      resizeOperation: null,
    });

    if (canCopy) {
      return {type: 'copy'};
    }

    // In reality, you would re-encode the track here
    return {type: 'drop'};
  },
});
```

```

Copy a video track to VP8, otherwise drop it
tsx

import {convertMedia, canCopyVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.webm',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onVideoTrack: async ({track, inputContainer, outputContainer}) => {
    const canCopy = canCopyVideoTrack({
      outputContainer,
      inputTrack: track,
      inputContainer,
      rotationToApply: 0,
      resizeOperation: null,
    });

    if (canCopy) {
      return {type: 'copy'};
    }

    // In reality, you would re-encode the track here
    return {type: 'drop'};
  },
});
```

## API [​](\#api "Direct link to API")

### `inputTrack` [​](\#inputtrack "Direct link to inputtrack")

_string_ `VideoTrack`

The input video track.

### `rotationToApply` [​](\#rotationtoapply "Direct link to rotationtoapply")

_number_

The number of degrees to rotate the video track.

### `inputContainer` [​](\#inputcontainer "Direct link to inputcontainer")

_string_ `ParseMediaContainer`

The container format of the input media.

### `outputContainer` [​](\#outputcontainer "Direct link to outputcontainer")

_string_ `ConvertMediaContainer`

The container format of the output media.

### `resizeOperation` [​](\#resizeoperation "Direct link to resizeoperation")

_string_ `ResizeOperation`

The [resize operation](/docs/webcodecs/resize-a-video) to apply to the video track.

## Rotation behavior [​](\#rotation-behavior "Direct link to Rotation behavior")

Any `rotationToApply` is in addition to an auto-rotation that is applied by default to fix the orientation of the video track.

If `rotationToApply` is not the same amount of rotation as `inputRotation`, this function will always return `false`, because rotation cannot be performed without re-encoding.

See: [Rotating a video](/docs/webcodecs/rotate-a-video)

## Return value [​](\#return-value "Direct link to Return value")

Returns a `boolean`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function on GitHub](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/can-copy-video-track.ts)
- [`canReencodeVideoTrack()`](/docs/webcodecs/can-reencode-video-track)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/can-copy-video-track.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
canCopyAudioTrack()](/docs/webcodecs/can-copy-audio-track) [Next\
\
defaultOnAudioTrackHandler()](/docs/webcodecs/default-on-audio-track-handler)

- [Examples](#examples)
- [API](#api)
  - [`inputTrack`](#inputtrack)
  - [`rotationToApply`](#rotationtoapply)
  - [`inputContainer`](#inputcontainer)
  - [`outputContainer`](#outputcontainer)
  - [`resizeOperation`](#resizeoperation)
- [Rotation behavior](#rotation-behavior)
- [Return value](#return-value)
- [See also](#see-also)
