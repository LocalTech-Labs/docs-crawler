---
title: convertMedia()v4.0.229
source: Remotion Documentation
last_updated: 2024-12-22
---

# convertMedia()v4.0.229

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- convertMedia()

On this page

# convertMedia() [v4.0.229](https://github.com/remotion-dev/remotion/releases/v4.0.229)

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

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

Re-encodes a video using [WebCodecs](https://developer.mozilla.org/en-US/docs/Web/API/WebCodecs_API) and [`@remotion/media-parser`](/docs/media-parser).

```

Converting a video
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

```

Converting a video
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

## API [​](\#api "Direct link to API")

### `src` [​](\#src "Direct link to src")

A `string` or `File` or `Blob`.

If it is a `string`, it must be a URL.

If it is a `File`, the `reader` field must be set to [`webFileReader`](/docs/media-parser/web-file-reader).

### `reader` [​](\#reader "Direct link to reader")

A reader interface.

Default value: [`fetchReader`](/docs/media-parser/fetch-reader), which uses `fetch()` to read the video.

If you pass [`webFileReader`](/docs/media-parser/web-file-reader), you must also pass a `File` as the `src` argument.

### `container` [​](\#container "Direct link to container")

_string_ `ConvertMediaContainer`

The container format to convert to. Currently, `"mp4"`, `"webm"` and `"wav"` is supported.

### `videoCodec?` [​](\#videocodec "Direct link to videocodec")

_string_ `ConvertMediaVideoCodec`

The video codec to convert to. Currently, `"h264"`, `"vp8"` and `"vp9"` are supported.

The default is defined by [`getDefaultVideoCodec()`](/docs/webcodecs/get-default-video-codec).

If a [`onVideoTrack`](#onvideotrack) handler is provided, it will override this setting.

### `audioCodec?` [​](\#audiocodec "Direct link to audiocodec")

_string_ `ConvertMediaAudioCodec`

The audio codec to convert to. Currently, only `"opus"` is supported.

The default is defined by [`getDefaultAudioCodec()`](/docs/webcodecs/get-default-audio-codec).

If an [`onAudioTrack`](#onaudiotrack) handler is provided, it will override this setting.

### `rotate?` [​](\#rotate "Direct link to rotate")

_number_

The number of degrees to rotate the video. See [Rotate a video](/docs/webcodecs/rotate-a-video) for more information.

### `resize?` [​](\#resize "Direct link to resize")

_object_ `ResizeOperation`

Resize the video. See [Resize a video](/docs/webcodecs/resize-a-video) for more information.

### `apiKey?` [​](\#apikey "Direct link to apikey")

If you are a customer of a [Remotion Company License](https://remotion.pro/license), you can provide an API key from your dashboard to track your conversions.

note

[Telemetry is enabled](/docs/webcodecs/telemetry) even if you don't provide an API key.

### `logLevel?` [​](\#loglevel "Direct link to loglevel")

_string_ `LogLevel`

One of `"error"`, `"warn"`, `"info"`, `"debug"`, `"trace"`.

Default value: `"info"`, which logs only important information.

### `onProgress?` [​](\#onprogress "Direct link to onprogress")

_Function_ `ConvertMediaOnProgress`

Allows receiving progress updates. The following fields are available:

```

tsx

import type {
  ConvertMediaOnProgress,
  ConvertMediaProgress,
} from '@remotion/webcodecs';

export const onProgress: ConvertMediaOnProgress = ({
  decodedVideoFrames,
  decodedAudioFrames,
  encodedVideoFrames,
  encodedAudioFrames,
  bytesWritten,
  millisecondsWritten,
  expectedOutputDurationInMs,
  overallProgress,
}: ConvertMediaProgress) => {
  console.log(decodedVideoFrames);

(parameter) decodedVideoFrames: number
  console.log(decodedAudioFrames);

(parameter) decodedAudioFrames: number
  console.log(encodedVideoFrames);

(parameter) encodedVideoFrames: number
  console.log(encodedAudioFrames);

(parameter) encodedAudioFrames: number
  console.log(bytesWritten);

(parameter) bytesWritten: number
  console.log(millisecondsWritten);

(parameter) millisecondsWritten: number
  console.log(expectedOutputDurationInMs);

(parameter) expectedOutputDurationInMs: number | null
  console.log(overallProgress);

(parameter) overallProgress: number | null
};
```

```

tsx

import type {
  ConvertMediaOnProgress,
  ConvertMediaProgress,
} from '@remotion/webcodecs';

export const onProgress: ConvertMediaOnProgress = ({
  decodedVideoFrames,
  decodedAudioFrames,
  encodedVideoFrames,
  encodedAudioFrames,
  bytesWritten,
  millisecondsWritten,
  expectedOutputDurationInMs,
  overallProgress,
}: ConvertMediaProgress) => {
  console.log(decodedVideoFrames);

(parameter) decodedVideoFrames: number
  console.log(decodedAudioFrames);

(parameter) decodedAudioFrames: number
  console.log(encodedVideoFrames);

(parameter) encodedVideoFrames: number
  console.log(encodedAudioFrames);

(parameter) encodedAudioFrames: number
  console.log(bytesWritten);

(parameter) bytesWritten: number
  console.log(millisecondsWritten);

(parameter) millisecondsWritten: number
  console.log(expectedOutputDurationInMs);

(parameter) expectedOutputDurationInMs: number | null
  console.log(overallProgress);

(parameter) overallProgress: number | null
};
```

### `onVideoFrame?` [​](\#onvideoframe "Direct link to onvideoframe")

_Function_ `ConvertMediaOnVideoFrame`

Allows you to hook into the video frames. The frames are [`VideoFrame`](https://developer.mozilla.org/en-US/docs/Web/API/VideoFrame) objects.

```

tsx

import type {ConvertMediaOnVideoFrame} from '@remotion/webcodecs';

export const onVideoFrame: ConvertMediaOnVideoFrame = ({frame}) => {
  console.log(frame);

(parameter) frame: VideoFrame

  // Do something with the frame, for example:
  // - Draw to a canvas
  // - Modify the frame

  // Then return the frame to be used for encoding.
  return frame;
};
```

```

tsx

import type {ConvertMediaOnVideoFrame} from '@remotion/webcodecs';

export const onVideoFrame: ConvertMediaOnVideoFrame = ({frame}) => {
  console.log(frame);

(parameter) frame: VideoFrame

  // Do something with the frame, for example:
  // - Draw to a canvas
  // - Modify the frame

  // Then return the frame to be used for encoding.
  return frame;
};
```

The callback function may be async.

When the function returns, the returned frame is used for video encoding.

You may return the same frame or replace it with a new `VideoFrame` object.

After the function returns, `convertMedia()` will call [`.close()`](https://developer.mozilla.org/en-US/docs/Web/API/VideoFrame/close) on the input and output frames.

This will destroy the frame and free up memory.

If you need a reference to the frame that lasts longer than the lifetime of this function, you must call [`.clone()`](https://developer.mozilla.org/en-US/docs/Web/API/VideoFrame/clone) on it.

If you return a different frame than the one you received, it must have the same values for `codedWidth`, `codedHeight`, `displayWidth` and `displayHeight`, `timestamp` and `duration` as the input frame.

### `signal?` [​](\#signal "Direct link to signal")

An [`AbortSignal`](https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal) to cancel the conversion process whenever the signal is aborted.

### `writer?` [​](\#writer "Direct link to writer")

_object_ `WriterInterface`

A writer interface. The following interfaces are available:

```

Buffer writer
tsx

import {bufferWriter} from '@remotion/media-parser/buffer';

(alias) const bufferWriter: WriterInterface
import bufferWriter
```

```

Buffer writer
tsx

import {bufferWriter} from '@remotion/media-parser/buffer';

(alias) const bufferWriter: WriterInterface
import bufferWriter
```

Write to a resizable Array Buffer.

```

Web File System writer
tsx

import {canUseWebFsWriter, webFsWriter} from '@remotion/media-parser/web-fs';

(alias) const webFsWriter: WriterInterface
import webFsWriter

await canUseWebFsWriter(); // boolean
```

```

Web File System writer
tsx

import {canUseWebFsWriter, webFsWriter} from '@remotion/media-parser/web-fs';

(alias) const webFsWriter: WriterInterface
import webFsWriter

await canUseWebFsWriter(); // boolean
```

Use the Web File System API to write to a file.

By default the writer is `webFsWriter`.

### `onAudioTrack?` [​](\#onaudiotrack "Direct link to onaudiotrack")

_Function_ `ConvertMediaOnAudioTrackHandler`

Take control of what to do for each audio track in the file: Re-encoding, copying, or dropping.

See [Track Transformation](/docs/webcodecs/track-transformation) for a guide on how to use this callback.

### `onVideoTrack?` [​](\#onvideotrack "Direct link to onvideotrack")

_Function_ `ConvertMediaOnVideoTrackHandler`

Take control of what to do for each video track in the file: Re-encoding, copying, or dropping.

See [Track Transformation](/docs/webcodecs/track-transformation) for a guide on how to use this callback.

### `progressIntervalInMs?` [​](\#progressintervalinms "Direct link to progressintervalinms")

_number_

The interval in milliseconds at which the `onProgress` callback is called.

Default `100`. Set to `0` for unthrottled updates.

Note that updates are fired very often and updating the DOM often may slow down the conversion process.

### `fields?` and Callbacks [​](\#fields-and-callbacks "Direct link to fields-and-callbacks")

You can obtain information about the video file while you are converting it.

This feature is inherited from [`parseMedia()`](/docs/media-parser/parse-media), but only the callback-style API is available.

```

Converting a video
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  fields: {
    size: true,
  },
  onSize: (size) => {
    console.log(size);

(parameter) size: number | null
  },
});
```

```

Converting a video
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp8',
  audioCodec: 'opus',
  fields: {
    size: true,
  },
  onSize: (size) => {
    console.log(size);

(parameter) size: number | null
  },
});
```

## License [​](\#license "Direct link to License")

[See notes about `@remotion/webcodecs`](/docs/webcodecs#-license-disclaimer).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/convert-media.ts)
- [`@remotion/webcodecs`](/docs/webcodecs)
- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/convert-media.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Track Transformation](/docs/webcodecs/track-transformation) [Next\
\
getAvailableContainers()](/docs/webcodecs/get-available-containers)

- [API](#api)
  - [`src`](#src)
  - [`reader`](#reader)
  - [`container`](#container)
  - [`videoCodec?`](#videocodec)
  - [`audioCodec?`](#audiocodec)
  - [`rotate?`](#rotate)
  - [`resize?`](#resize)
  - [`apiKey?`](#apikey)
  - [`logLevel?`](#loglevel)
  - [`onProgress?`](#onprogress)
  - [`onVideoFrame?`](#onvideoframe)
  - [`signal?`](#signal)
  - [`writer?`](#writer)
  - [`onAudioTrack?`](#onaudiotrack)
  - [`onVideoTrack?`](#onvideotrack)
  - [`progressIntervalInMs?`](#progressintervalinms)
  - [`fields?` and Callbacks](#fields-and-callbacks)
- [License](#license)
- [See also](#see-also)
