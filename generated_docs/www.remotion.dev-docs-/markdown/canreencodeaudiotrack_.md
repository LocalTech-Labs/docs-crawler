---
title: canReencodeAudioTrack()
source: Remotion Documentation
last_updated: 2024-12-22
---

# canReencodeAudioTrack()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- canReencodeAudioTrack()

On this page

# canReencodeAudioTrack()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Given an `AudioTrack`, determine if it can be re-encoded to another track.

You can obtain an `AudioTrack` using [`parseMedia()`](/docs/media-parser/parse-media) or during the conversion process using the [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) callback of [`convertMedia()`](/docs/webcodecs/convert-media).

## Examples [​](\#examples "Direct link to Examples")

```

Check if audio tracks can be re-encoded to Opus
tsx

import {parseMedia} from '@remotion/media-parser';
import {canReencodeAudioTrack} from '@remotion/webcodecs';

const {audioTracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});

for (const track of audioTracks) {
  await canReencodeAudioTrack({
    track,
    audioCodec: 'opus',
    bitrate: 128000,
  });
}
```

```

Check if audio tracks can be re-encoded to Opus
tsx

import {parseMedia} from '@remotion/media-parser';
import {canReencodeAudioTrack} from '@remotion/webcodecs';

const {audioTracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});

for (const track of audioTracks) {
  await canReencodeAudioTrack({
    track,
    audioCodec: 'opus',
    bitrate: 128000,
  });
}
```

```

Convert an audio track to Opus, otherwise drop it
tsx

import {convertMedia, canReencodeAudioTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onAudioTrack: async ({track}) => {
    const canReencode = await canReencodeAudioTrack({
      track,
      audioCodec: 'opus',
      bitrate: 128000,
    });

    if (canReencode) {
      return {type: 'reencode', audioCodec: 'opus', bitrate: 128000};
    }

    return {type: 'drop'};
  },
});
```

```

Convert an audio track to Opus, otherwise drop it
tsx

import {convertMedia, canReencodeAudioTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onAudioTrack: async ({track}) => {
    const canReencode = await canReencodeAudioTrack({
      track,
      audioCodec: 'opus',
      bitrate: 128000,
    });

    if (canReencode) {
      return {type: 'reencode', audioCodec: 'opus', bitrate: 128000};
    }

    return {type: 'drop'};
  },
});
```

## API [​](\#api "Direct link to API")

### `track` [​](\#track "Direct link to track")

A `AudioTrack` object.

### `audioCodec` [​](\#audiocodec "Direct link to audiocodec")

_string_ `ConvertMediaAudioCodec`

### `bitrate` [​](\#bitrate "Direct link to bitrate")

_number_

The bitrate with which you'd like to re-encode the audio track.

## Return value [​](\#return-value "Direct link to Return value")

Returns a `Promise<boolean>`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function on GitHub](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/can-reencode-audio-track.ts)
- [`canReencodeVideoTrack()`](/docs/webcodecs/can-reencode-video-track)
- [`convertMedia()`](/docs/webcodecs/convert-media)
- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/can-reencode-audio-track.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAvailableContainers()](/docs/webcodecs/get-available-containers) [Next\
\
canReencodeVideoTrack()](/docs/webcodecs/can-reencode-video-track)

- [Examples](#examples)
- [API](#api)
  - [`track`](#track)
  - [`audioCodec`](#audiocodec)
  - [`bitrate`](#bitrate)
- [Return value](#return-value)
- [See also](#see-also)
