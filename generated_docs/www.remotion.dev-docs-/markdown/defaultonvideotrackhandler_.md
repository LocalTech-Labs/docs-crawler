---
title: defaultOnVideoTrackHandler()
source: Remotion Documentation
last_updated: 2024-12-22
---

# defaultOnVideoTrackHandler()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- defaultOnVideoTrackHandler()

On this page

# defaultOnVideoTrackHandler()

This is the default function if no [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) handler is provided to [`convertMedia()`](/docs/webcodecs/convert-media).

You may use this function if you want to customize part of the track transformation logic, but fall back to the default behavior for the rest.

```

Falling back to the default behavior
tsx

import {convertMedia, defaultOnAudioTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onAudioTrack: (params) => {
    // Custom logic for handling video tracks
    // ...

    // Fall back to the default behavior
    return defaultOnAudioTrackHandler(params);
  },
});
```

```

Falling back to the default behavior
tsx

import {convertMedia, defaultOnAudioTrackHandler} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  onAudioTrack: (params) => {
    // Custom logic for handling video tracks
    // ...

    // Fall back to the default behavior
    return defaultOnAudioTrackHandler(params);
  },
});
```

## Algorithm [​](\#algorithm "Direct link to Algorithm")

The default behavior is as follows:

- Check if the track can be copied without re-encoding, if true, then do that.
- Determine the video codec to be used - either [`videoCodec`](/docs/webcodecs/convert-media#videocodec) which was passed to [`convertMedia()`](/docs/webcodecs/convert-media) or the default codec for the container.
- Check if the track can be re-encoded with the chosen video codec, if true, then do that.
- If the track can be neither copied nor re-encoded, then fail the render.

You may alternatively return `{type: 'drop'}` to remove the video track, but still succeed the other tracks.

```

Source code for defaultOnVideoTrackHandler
tsx

import {
  canReencodeVideoTrack,
  getDefaultVideoCodec,
  ConvertMediaOnVideoTrackHandler,
  VideoOperation,
} from '@remotion/webcodecs';

export const defaultOnVideoTrackHandler: ConvertMediaOnVideoTrackHandler =
  async ({
    track,
    defaultVideoCodec,
    logLevel,
    outputContainer,
    rotate,
    inputContainer,
    canCopyTrack,
  }): Promise<VideoOperation> => {
    if (canCopyTrack) {
      return Promise.resolve({type: 'copy'});
    }

    // If for example exporting to audio, the default video codec will be null
    if (defaultVideoCodec === null) {
      return Promise.resolve({type: 'drop'});
    }

    const canReencode = await canReencodeVideoTrack({
      videoCodec: defaultVideoCodec,
      track,
    });

    if (canReencode) {
      return Promise.resolve({
        type: 'reencode',
        videoCodec: defaultVideoCodec,
        rotation: rotate - track.rotation,
      });
    }

    return Promise.resolve({type: 'fail'});
  };
```

```

Source code for defaultOnVideoTrackHandler
tsx

import {
  canReencodeVideoTrack,
  getDefaultVideoCodec,
  ConvertMediaOnVideoTrackHandler,
  VideoOperation,
} from '@remotion/webcodecs';

export const defaultOnVideoTrackHandler: ConvertMediaOnVideoTrackHandler =
  async ({
    track,
    defaultVideoCodec,
    logLevel,
    outputContainer,
    rotate,
    inputContainer,
    canCopyTrack,
  }): Promise<VideoOperation> => {
    if (canCopyTrack) {
      return Promise.resolve({type: 'copy'});
    }

    // If for example exporting to audio, the default video codec will be null
    if (defaultVideoCodec === null) {
      return Promise.resolve({type: 'drop'});
    }

    const canReencode = await canReencodeVideoTrack({
      videoCodec: defaultVideoCodec,
      track,
    });

    if (canReencode) {
      return Promise.resolve({
        type: 'reencode',
        videoCodec: defaultVideoCodec,
        rotation: rotate - track.rotation,
      });
    }

    return Promise.resolve({type: 'fail'});
  };
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/default-on-video-track-handler.ts)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/default-on-video-track-handler.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
defaultOnAudioTrackHandler()](/docs/webcodecs/default-on-audio-track-handler) [Next\
\
getDefaultAudioCodec()](/docs/webcodecs/get-default-audio-codec)

- [Algorithm](#algorithm)
- [See also](#see-also)
