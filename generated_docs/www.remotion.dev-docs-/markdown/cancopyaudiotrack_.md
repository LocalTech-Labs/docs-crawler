---
title: canCopyAudioTrack()
source: Remotion Documentation
last_updated: 2024-12-22
---

# canCopyAudioTrack()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- canCopyAudioTrack()

On this page

# canCopyAudioTrack()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Given an `AudioTrack`, determine if it can be copied to the output without re-encoding.

You can obtain an `AudioTrack` using [`parseMedia()`](/docs/media-parser/parse-media) or during the conversion process using the [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) callback of [`convertMedia()`](/docs/webcodecs/convert-media).

## Examples [​](\#examples "Direct link to Examples")

```

Check if an audio track can be copied
tsx

import {parseMedia} from '@remotion/media-parser';
import {canCopyAudioTrack} from '@remotion/webcodecs';

const {tracks, container} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
    container: true,
  },
});

for (const track of tracks.audioTracks) {
  canCopyAudioTrack({
    inputCodec: track.codecWithoutConfig,
    outputContainer: 'webm',
    inputContainer: container,
  }); // bool
}
```

```

Check if an audio track can be copied
tsx

import {parseMedia} from '@remotion/media-parser';
import {canCopyAudioTrack} from '@remotion/webcodecs';

const {tracks, container} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
    container: true,
  },
});

for (const track of tracks.audioTracks) {
  canCopyAudioTrack({
    inputCodec: track.codecWithoutConfig,
    outputContainer: 'webm',
    inputContainer: container,
  }); // bool
}
```

```

Copy an audio track to Opus, otherwise drop it
tsx

import {convertMedia, canCopyAudioTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onAudioTrack: async ({track, outputContainer, inputContainer}) => {
    const canCopy = canCopyAudioTrack({
      inputCodec: track.codecWithoutConfig,
      outputContainer,
      inputContainer,
    });

    if (canCopy) {
      return {type: 'copy'};
    }

    // Just to keep the example brief, in reality, you would re-encode the track here
    return {type: 'drop'};
  },
});
```

```

Copy an audio track to Opus, otherwise drop it
tsx

import {convertMedia, canCopyAudioTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  onAudioTrack: async ({track, outputContainer, inputContainer}) => {
    const canCopy = canCopyAudioTrack({
      inputCodec: track.codecWithoutConfig,
      outputContainer,
      inputContainer,
    });

    if (canCopy) {
      return {type: 'copy'};
    }

    // Just to keep the example brief, in reality, you would re-encode the track here
    return {type: 'drop'};
  },
});
```

## API [​](\#api "Direct link to API")

### `inputCodec` [​](\#inputcodec "Direct link to inputcodec")

_string_ `MediaParserAudioCodec`

The codec of the input audio track.

### `inputContainer` [​](\#inputcontainer "Direct link to inputcontainer")

_string_ `ParseMediaContainer`

The container format of the input media.

### `outputContainer` [​](\#outputcontainer "Direct link to outputcontainer")

_string_ `ConvertMediaContainer`

The container format of the output media.

## Return value [​](\#return-value "Direct link to Return value")

Returns a `boolean`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function on GitHub](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/can-copy-audio-track.ts)
- [`canCopyVideoTrack()`](/docs/webcodecs/can-copy-video-track)
- [`canReencodeAudioTrack()`](/docs/webcodecs/can-reencode-audio-track)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/can-copy-audio-track.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
canReencodeVideoTrack()](/docs/webcodecs/can-reencode-video-track) [Next\
\
canCopyVideoTrack()](/docs/webcodecs/can-copy-video-track)

- [Examples](#examples)
- [API](#api)
  - [`inputCodec`](#inputcodec)
  - [`inputContainer`](#inputcontainer)
  - [`outputContainer`](#outputcontainer)
- [Return value](#return-value)
- [See also](#see-also)
