---
title: Track Transformation
source: Remotion Documentation
last_updated: 2024-12-22
---

# Track Transformation

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- Track Transformation

On this page

# Track Transformation

When transforming media, there are multiple thing that can be done for each audio or video track:

- Copying the track without re-encoding
- Re-encoding the track into a different codec
- Removing the track

[`@remotion/webcodecs`](/docs/webcodecs) allows you to decide for each track what to do with it.

## Using the defaults [​](\#using-the-defaults "Direct link to Using the defaults")

The minimum amount of configuration is to only specify a [`src`](/docs/webcodecs/convert-media#src) and an output [`container`](/docs/webcodecs/convert-media#container).

```

Using the default codecs
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

```

Using the default codecs
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

The default codecs are defined by [`getDefaultAudioCodec()`](/docs/webcodecs/get-default-audio-codec) and [`getDefaultVideoCodec()`](/docs/webcodecs/get-default-video-codec).

## Choosing codecs [​](\#choosing-codecs "Direct link to Choosing codecs")

You can use the [`videoCodec`](/docs/webcodecs/convert-media#videocodec) and [`audioCodec`](/docs/webcodecs/convert-media#audiocodec) options to transform all tracks to the codecs you specify.

```

Choosing video and audio codecs
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
});
```

```

Choosing video and audio codecs
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
});
```

## Handle each track individually [​](\#handle-each-track-individually "Direct link to Handle each track individually")

With the [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) and [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) callbacks, you can decide for each track what to do with it.

```

Using the onVideoTrack() API
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: ({track}) => {
    if (track.codecWithoutConfig === 'vp8') {
      return {type: 'copy'};
    }

    return {type: 'reencode', videoCodec: 'vp8'};
  },
});
```

```

Using the onVideoTrack() API
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: ({track}) => {
    if (track.codecWithoutConfig === 'vp8') {
      return {type: 'copy'};
    }

    return {type: 'reencode', videoCodec: 'vp8'};
  },
});
```

[`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) and [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) have a higher priority than [`videoCodec`](/docs/webcodecs/convert-media#videocodec) and [`audioCodec`](/docs/webcodecs/convert-media#audiocodec).

The options for video codecs are:

- `{"type": "copy"}` \- Copy the track without re-encoding
- `{"type": "reencode", "videoCodec": ConvertMediaVideoCodec}` \- Re-encode the track into the specified codec
- `{"type": "drop"}` \- Remove the track from the output
- `{"type": "fail"}` \- Fail and stop the conversion process

The options for audio codecs are:

- `{"type": "copy"}` \- Copy the track without re-encoding
- `{"type": "reencode", "audioCodec": ConvertMediaAudioCodec; bitrate: number}` \- Re-encode the track into the specified codec. The suggested bitrate to use is `128000`.
- `{"type": "drop"}` \- Remove the track from the output
- `{"type": "fail"}` \- Fail and stop the conversion process

The enums `ConvertMediaVideoCodec` and `ConvertMediaAudioCodec` can be imported from `@remotion/webcodecs`.

## Checking if a track can be copied [​](\#checking-if-a-track-can-be-copied "Direct link to Checking if a track can be copied")

To check if it is possible to return `{"type": "copy"}`, you can use `canCopyTrack` property you get from `onVideoTrack`.

```

Using the canCopyVideoTrack() API
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: ({track, inputContainer, outputContainer, canCopyTrack}) => {
    if (canCopyTrack) {
      return {type: 'copy'};
    }

    return {type: 'reencode', videoCodec: 'vp8'};
  },
});
```

```

Using the canCopyVideoTrack() API
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: ({track, inputContainer, outputContainer, canCopyTrack}) => {
    if (canCopyTrack) {
      return {type: 'copy'};
    }

    return {type: 'reencode', videoCodec: 'vp8'};
  },
});
```

To check outside of a `onVideoTrack` handler, you can also use the the [`canCopyVideoTrack()`](/docs/webcodecs/can-copy-video-track) and [`canCopyAudioTrack()`](/docs/webcodecs/can-copy-audio-track) APIs

## Checking if a track can be re-encoded [​](\#checking-if-a-track-can-be-re-encoded "Direct link to Checking if a track can be re-encoded")

To check if it is possible to return `{"type": "reencode"}`, you can use the [`canReencodeVideoTrack()`](/docs/webcodecs/can-reencode-video-track) and [`canReencodeAudioTrack()`](/docs/webcodecs/can-reencode-audio-track) APIs.

```

Using the canReencodeVideoTrack() API
tsx

import {convertMedia, canReencodeVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: async ({track}) => {
    const canReencode = await canReencodeVideoTrack({
      videoCodec: 'vp8',
      track,
    });

    if (canReencode) {
      return {type: 'reencode', videoCodec: 'vp8'};
    }

    return {type: 'drop'};
  },
});
```

```

Using the canReencodeVideoTrack() API
tsx

import {convertMedia, canReencodeVideoTrack} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  audioCodec: 'opus',
  onVideoTrack: async ({track}) => {
    const canReencode = await canReencodeVideoTrack({
      videoCodec: 'vp8',
      track,
    });

    if (canReencode) {
      return {type: 'reencode', videoCodec: 'vp8'};
    }

    return {type: 'drop'};
  },
});
```

## Asynchronously determining action [​](\#asynchronously-determining-action "Direct link to Asynchronously determining action")

The [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) and [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) callbacks can be asynchronous.

While the operations are unresolved, reading of the input fill is paused.

## Decide behavior upfront [​](\#decide-behavior-upfront "Direct link to Decide behavior upfront")

If you want to display a UI letting the user choose codec settings before the conversion starts, you can do so.

Use [`parseMedia()`](/docs/media-parser/parse-media) to get video and audio tracks respectively:

```

Using parseMedia() to get tracks upfront.
tsx

import {parseMedia} from '@remotion/media-parser';

const {tracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});
```

```

Using parseMedia() to get tracks upfront.
tsx

import {parseMedia} from '@remotion/media-parser';

const {tracks} = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  fields: {
    tracks: true,
  },
});
```

You now have an object of two arrays of `VideoTrack` and `AudioTrack` objects.

You can now use [`canReencodeAudioTrack()`](/docs/webcodecs/can-reencode-audio-track), [`canReencodeVideoTrack()`](/docs/webcodecs/can-reencode-video-track), [`canCopyAudioTrack()`](/docs/webcodecs/can-copy-audio-track), and [`canCopyVideoTrack()`](/docs/webcodecs/can-copy-video-track) to determine which options to show.

Use the [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) and [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) callbacks to return the user selection.

You can use the `trackId` field as the unique key for each track.

## Falling back to default [​](\#falling-back-to-default "Direct link to Falling back to default")

The default values for [`onVideoTrack`](/docs/webcodecs/convert-media#onvideotrack) and [`onAudioTrack`](/docs/webcodecs/convert-media#onaudiotrack) are the functions [`defaultOnVideoTrackHandler`](/docs/webcodecs/default-on-video-track-handler) and [`defaultOnAudioTrackHandler`](/docs/webcodecs/default-on-audio-track-handler) respectively.

If you only want to override part of the logic, you can return the default resolver functions at the end of your logic.

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

## Debugging [​](\#debugging "Direct link to Debugging")

Pass [`logLevel: "verbose"`](/docs/webcodecs/convert-media#loglevel) to [`convertMedia()`](/docs/webcodecs/convert-media) to see debug information in the console, including how the defaults have decided which operations to take.

## Reference implementation [​](\#reference-implementation "Direct link to Reference implementation")

Visit the [source code](https://github.com/remotion-dev/remotion/tree/main/packages/convert) for [convert.remotion.dev](https://convert.remotion.dev) to see a reference implementation for an online video converter that displays a user interface for all possible options.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/webcodecs`](/docs/webcodecs)
- [`convertMedia()`](/docs/webcodecs/convert-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/track-transformation.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/webcodecs/) [Next\
\
convertMedia()](/docs/webcodecs/convert-media)

- [Using the defaults](#using-the-defaults)
- [Choosing codecs](#choosing-codecs)
- [Handle each track individually](#handle-each-track-individually)
- [Checking if a track can be copied](#checking-if-a-track-can-be-copied)
- [Checking if a track can be re-encoded](#checking-if-a-track-can-be-re-encoded)
- [Asynchronously determining action](#asynchronously-determining-action)
- [Decide behavior upfront](#decide-behavior-upfront)
- [Falling back to default](#falling-back-to-default)
- [Debugging](#debugging)
- [Reference implementation](#reference-implementation)
- [See also](#see-also)
