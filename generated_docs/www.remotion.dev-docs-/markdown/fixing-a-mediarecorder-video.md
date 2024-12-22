---
title: Fixing a MediaRecorder video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Fixing a MediaRecorder video

- [Home page](/)
- [WebCodecs](/docs/webcodecs/)
- Fixing a MediaRecorder video

On this page

# Fixing a MediaRecorder video

When recording a video with the [`MediaRecorder`](https://developer.mozilla.org/en-US/docs/Web/API/MediaRecorder) and [`getUserMedia()`](https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia) API, a WebM file gets created that may have some playback and compatibility issues.

Namely, it:

- does not show the video duration in the browser
- can be slow to seek around
- does not play in all browsers (Safari)

To fix these issues, you should either re-encode or re-mux the video.

- **Re-encoding** will decode and re-encode the frames (if you want to a different codec).

It will be slower, but allows to use a more compatible format like MP4.
- **Re-muxing** will leave the frames, but disassemble and reassemble the file to allow for seeking points and duration metadata to be inserted.

It is very fast, but does not allow for codec changes.

## Why does this happen? [​](\#why-does-this-happen "Direct link to Why does this happen?")

The reason is that while recording, browsers open a file and append video chunks to the end.

However, the important metadata such as duration and seeking points should be at the beginning of the file for them to be useful.

By placing the metadata at the beginning, the video player has the information it needs to seek around the video and display the duration.

## Re-encoding on the server [​](\#re-encoding-on-the-server "Direct link to Re-encoding on the server")

The traditional way is to use a server and run [FFmpeg](https://ffmpeg.org) on it:

```

sh

ffmpeg -i input.webm -c:v libx264 -c:a aac output.mp4
```

```

sh

ffmpeg -i input.webm -c:v libx264 -c:a aac output.mp4
```

or for re-muxing:

```

sh

ffmpeg -i input.webm -c copy output.webm
```

```

sh

ffmpeg -i input.webm -c copy output.webm
```

## Re-encoding using `@remotion/webcodecs` [​](\#re-encoding-using-remotionwebcodecs "Direct link to re-encoding-using-remotionwebcodecs")

You can also re-encode the video in the browser using the new and experimental [`@remotion/webcodecs`](/docs/webcodecs) package.

```

Re-encoding a video
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// The video get from the MediaRecorder as a Blob
const blob = new Blob([], {type: 'video/webm'});

await convertMedia({
  src: blob,
  container: 'mp4',
  videoCodec: 'h264',
  audioCodec: 'aac',
  reader: webFileReader,
});
```

```

Re-encoding a video
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// The video get from the MediaRecorder as a Blob
const blob = new Blob([], {type: 'video/webm'});

await convertMedia({
  src: blob,
  container: 'mp4',
  videoCodec: 'h264',
  audioCodec: 'aac',
  reader: webFileReader,
});
```

💼 Important License Disclaimer about `@remotion/webcodecs`

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

🚧 Unstable API

This package is experimental.

We might change the API at any time, until we remove this notice.

## Re-muxing using `@remotion/webcodecs` [​](\#re-muxing-using-remotionwebcodecs "Direct link to re-muxing-using-remotionwebcodecs")

Instead of re-encoding to an MP4, you can also re-mux the video to a new WebM file:

```

Re-muxing a video
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// The video get from the MediaRecorder as a Blob
const blob = new Blob([], {type: 'video/webm'});

await convertMedia({
  src: blob,
  container: 'webm',
  reader: webFileReader,
});
```

```

Re-muxing a video
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// The video get from the MediaRecorder as a Blob
const blob = new Blob([], {type: 'video/webm'});

await convertMedia({
  src: blob,
  container: 'webm',
  reader: webFileReader,
});
```

[`convertMedia()`](/docs/webcodecs/convert-media) will move the metadata and seek information to the beginning of the video.

💼 Important License Disclaimer about `@remotion/webcodecs`

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

🚧 Unstable API

This package is experimental.

We might change the API at any time, until we remove this notice.

## Sample application [​](\#sample-application "Direct link to Sample application")

Use [remotion.dev/convert](https://remotion.dev/convert) to fix a MediaRecorder video online using WebCodecs.

See the [source code](https://github.com/remotion-dev/remotion/tree/main/packages/convert) for a reference on how to implement it.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/webcodecs`](/docs/webcodecs)
- [Convert a video](/docs/webcodecs/convert-a-video)
- [`convertMedia()`](/docs/webcodecs/convert-media)
- [`@remotion/media-parser`](/docs/media-parser)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/fix-a-mediarecorder-video.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Resize a video](/docs/webcodecs/resize-a-video) [Next\
\
Telemetry in @remotion/webcodecs](/docs/webcodecs/telemetry)

- [Why does this happen?](#why-does-this-happen)
- [Re-encoding on the server](#re-encoding-on-the-server)
- [Re-encoding using `@remotion/webcodecs`](#re-encoding-using-remotionwebcodecs)
- [Re-muxing using `@remotion/webcodecs`](#re-muxing-using-remotionwebcodecs)
- [Sample application](#sample-application)
- [See also](#see-also)
