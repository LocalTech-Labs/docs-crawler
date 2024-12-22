---
title: defaultOnAudioTrackHandler()
source: Remotion Documentation
last_updated: 2024-12-22
---

# defaultOnAudioTrackHandler()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- defaultOnAudioTrackHandler()

On this page

# defaultOnAudioTrackHandler()

This is the default function if no [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) handler is provided to [`convertMedia()`](/docs/webcodecs/convert-media).

You may use this function if you want to customize part of the track transformation logic, but fall back to the default behavior for the rest.

```

Falling back to the default behavior
tsx

import {convertMedia, defaultOnVideoTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onVideoTrack: (params) => {
    // Custom logic for handling video tracks
    // ...

    // Fall back to the default behavior
    return defaultOnVideoTrackHandler(params);
  },
});
```

```

Falling back to the default behavior
tsx

import {convertMedia, defaultOnVideoTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onVideoTrack: (params) => {
    // Custom logic for handling video tracks
    // ...

    // Fall back to the default behavior
    return defaultOnVideoTrackHandler(params);
  },
});
```

## Algorithm [​](\#algorithm "Direct link to Algorithm")

The default behavior is as follows:

- Check if the track can be copied without re-encoding, if true, then do that.
- Determine the audio codec to be used - either [`audioCodec`](/docs/webcodecs/convert-media#audiocodec) which was passed to [`convertMedia()`](/docs/webcodecs/convert-media) or the default codec for the container.
- Check if the track can be re-encoded with the chosen audio codec and bitrate, if true, then do that.
- If the track can be neither copied nor re-encoded, then fail the render.

You may alternatively return `{type: 'drop'}` to remove the audio track, but still succeed the other tracks.

This is the source code for the default function. You may use this as a reference to create your own custom handler.

```

Source code for defaultOnAudioTrackHandler
tsx

import {
  canReencodeAudioTrack,
  AudioOperation,
  ConvertMediaOnAudioTrackHandler,
  getDefaultAudioCodec,
} from '@remotion/webcodecs';

const DEFAULT_BITRATE = 128_000;

export const defaultOnAudioTrackHandler: ConvertMediaOnAudioTrackHandler =
  async ({
    track,
    defaultAudioCodec,
    logLevel,
    inputContainer,
    outputContainer,
    canCopyTrack,
  }): Promise<AudioOperation> => {
    const bitrate = DEFAULT_BITRATE;

    if (canCopyTrack) {
      return Promise.resolve({type: 'copy'});
    }

    // In the future, we might support containers that don't support audio
    // (like GIF, animated WebP, etc.) - in that case, we should drop the audio
    if (defaultAudioCodec === null) {
      return Promise.resolve({type: 'drop'});
    }

    const canReencode = await canReencodeAudioTrack({
      audioCodec: defaultAudioCodec,
      track,
      bitrate,
    });

    if (canReencode) {
      return Promise.resolve({
        type: 'reencode',
        bitrate,
        audioCodec: defaultAudioCodec,
      });
    }

    return Promise.resolve({type: 'fail'});
  };
```

```

Source code for defaultOnAudioTrackHandler
tsx

import {
  canReencodeAudioTrack,
  AudioOperation,
  ConvertMediaOnAudioTrackHandler,
  getDefaultAudioCodec,
} from '@remotion/webcodecs';

const DEFAULT_BITRATE = 128_000;

export const defaultOnAudioTrackHandler: ConvertMediaOnAudioTrackHandler =
  async ({
    track,
    defaultAudioCodec,
    logLevel,
    inputContainer,
    outputContainer,
    canCopyTrack,
  }): Promise<AudioOperation> => {
    const bitrate = DEFAULT_BITRATE;

    if (canCopyTrack) {
      return Promise.resolve({type: 'copy'});
    }

    // In the future, we might support containers that don't support audio
    // (like GIF, animated WebP, etc.) - in that case, we should drop the audio
    if (defaultAudioCodec === null) {
      return Promise.resolve({type: 'drop'});
    }

    const canReencode = await canReencodeAudioTrack({
      audioCodec: defaultAudioCodec,
      track,
      bitrate,
    });

    if (canReencode) {
      return Promise.resolve({
        type: 'reencode',
        bitrate,
        audioCodec: defaultAudioCodec,
      });
    }

    return Promise.resolve({type: 'fail'});
  };
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/default-on-audio-track-handler.ts)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/default-on-audio-track-handler.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
canCopyVideoTrack()](/docs/webcodecs/can-copy-video-track) [Next\
\
defaultOnVideoTrackHandler()](/docs/webcodecs/default-on-video-track-handler)

- [Algorithm](#algorithm)
- [See also](#see-also)