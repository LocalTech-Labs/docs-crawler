---
title: canReencodeVideoTrack()
source: Remotion Documentation
last_updated: 2024-12-22
---

# canReencodeVideoTrack()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- canReencodeVideoTrack()

On this page

# canReencodeVideoTrack()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Given a `VideoTrack`, determine if it can be re-encoded to another track.

You can obtain a `VideoTrack` using [`parseMedia()`](/docs/media-parser/parse-media) or during the conversion process using the [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) callback of [`convertMedia()`](/docs/webcodecs/convert-media).

## Examples [​](\#examples "Direct link to Examples")

```

Check if video tracks can be re-encoded to VP8
tsx

import {parseMedia} from '@remotion/media-parser';
import {canReencodeVideoTrack} from '@remotion/webcodecs';

const {tracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});

for (const track of tracks.videoTracks) {
  await canReencodeVideoTrack({
    track,
    videoCodec: 'vp8',
  });
}
```

```

Check if video tracks can be re-encoded to VP8
tsx

import {parseMedia} from '@remotion/media-parser';
import {canReencodeVideoTrack} from '@remotion/webcodecs';

const {tracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});

for (const track of tracks.videoTracks) {
  await canReencodeVideoTrack({
    track,
    videoCodec: 'vp8',
  });
}
```

```

Convert a video track to VP8, otherwise drop it
tsx

import {convertMedia, canReencodeVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onVideoTrack: async ({track}) => {
    const canReencode = await canReencodeVideoTrack({
      track,
      videoCodec: 'vp8',
    });

    if (canReencode) {
      return {type: 'reencode', videoCodec: 'vp8'};
    }

    return {type: 'drop'};
  },
});
```

```

Convert a video track to VP8, otherwise drop it
tsx

import {convertMedia, canReencodeVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onVideoTrack: async ({track}) => {
    const canReencode = await canReencodeVideoTrack({
      track,
      videoCodec: 'vp8',
    });

    if (canReencode) {
      return {type: 'reencode', videoCodec: 'vp8'};
    }

    return {type: 'drop'};
  },
});
```

## API [​](\#api "Direct link to API")

### `track` [​](\#track "Direct link to track")

A `VideoTrack` object.

### `videoCodec` [​](\#videocodec "Direct link to videocodec")

_string_ `ConvertMediaVideoCodec`

One of the supported video codecs: `"vp8"`, `"vp9"`.

## Return value [​](\#return-value "Direct link to Return value")

Returns a `Promise<boolean>`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function on GitHub](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/can-reencode-video-track.ts)
- [`convertMedia()`](/docs/webcodecs/convert-media)
- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/can-reencode-video-track.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
canReencodeAudioTrack()](/docs/webcodecs/can-reencode-audio-track) [Next\
\
canCopyAudioTrack()](/docs/webcodecs/can-copy-audio-track)

- [Examples](#examples)
- [API](#api)
  - [`track`](#track)
  - [`videoCodec`](#videocodec)
- [Return value](#return-value)
- [See also](#see-also)
