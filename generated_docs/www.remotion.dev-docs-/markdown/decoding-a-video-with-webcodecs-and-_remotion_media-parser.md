---
title: Decoding a video with WebCodecs and @remotion/media-parser
source: Remotion Documentation
last_updated: 2024-12-22
---

# Decoding a video with WebCodecs and @remotion/media-parser

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- WebCodecs

On this page

# Decoding a video with WebCodecs and @remotion/media-parser

[`parseMedia()`](/docs/media-parser/parse-media) is able to extract tracks and samples from audio and video in a format that is suitable for usage with WebCodecs APIs.

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

## With `@remotion/webcodecs` [​](\#with-remotionwebcodecs "Direct link to with-remotionwebcodecs")

[`@remotion/webcodecs`](/docs/webcodecs) is a package that uses `@remotion/media-parser` to provide video processing in the browser.

It handles various browser quirks and hard implementation details for you.

We recommend it if you want to build a solution based on WebCodecs.

It is a higher-level API that is easier to use than `@remotion/media-parser`.

## Using `@remotion/media-parser` [​](\#using-remotionmedia-parser "Direct link to using-remotionmedia-parser")

`@remotion/media-parser` is a lower-level API that allows you to interface with WebCodecs directly.

```

Reading video frames
tsx

import {parseMedia, OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  videoDecoder.configure(track);

  return (sample) => {
    videoDecoder.decode(new EncodedVideoChunk(sample));
  };
};

const result = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  onVideoTrack,
});
```

```

Reading video frames
tsx

import {parseMedia, OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  videoDecoder.configure(track);

  return (sample) => {
    videoDecoder.decode(new EncodedVideoChunk(sample));
  };
};

const result = await parseMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  onVideoTrack,
});
```

## Why is this useful? [​](\#why-is-this-useful "Direct link to Why is this useful?")

WebCodecs is the fastest way to decode videos in the browser.

WebAssembly solutions need to strip the CPU-specific optimizations and cannot benefit from hardware acceleration.

To decode a video using WebCodecs, the binary format of a video needs to be understood and parsed.

This is a meticulous task that requires a lot of domain knowledge.

Video files usually come in one of two container formats: ISO BMFF (.mp4, .mov) or Matroska (.webm, .mkv).

Libraries like mp4box.js do a good job of parsing these containers, but are scoped to the specific container format, meaning you need to mix multiple libraries.

[`parseMedia()`](/docs/media-parser/parse-media) allows to to read an arbitrary video file (in the future: an arbitrary media file) and interface with it regardless of container, video codec and audio codec.

It uses modern Web APIs like `fetch()`, `ReadableStream` and resizeable ArrayBuffers, and returns data structures that are designed to be used with WebCodecs APIs.

## Will Remotion switch to WebCodecs? [​](\#will-remotion-switch-to-webcodecs "Direct link to Will Remotion switch to WebCodecs?")

Not in the foreseeable future - Remotion currently renders videos with FFmpeg and a headless browser.

FFmpeg is just as fast as WebCodecs (they share the same code) - therefore it is not necessary to switch to WebCodecs.

Remotion cannot export videos in the browser, because browsers don't have APIs for capturing the viewport.

An exception is the `canvas` element, however Remotion supports all ways of drawing to the viewport: HTML, CSS, SVG, and Canvas.

We are interested in WebCodecs because it still has the potential to solve a lot of problems for developers and `@remotion/media-parser` as a whole can solve a lot of problems for users.

See also: [Can I render videos in the browser?](/docs/miscellaneous/render-in-browser).

## Practical considerations [​](\#practical-considerations "Direct link to Practical considerations")

If you use `parseMedia()` with codecs, make the following considerations for your implementation.

### Check browser support for `@remotion/media-parser` [​](\#check-browser-support-for-remotionmedia-parser "Direct link to check-browser-support-for-remotionmedia-parser")

Remotion requires the `fetch()` and Resizeable ArrayBuffer APIs to be present.

Check if your runtime supports these APIs before you use `parseMedia()`.

```

tsx

const canUseMediaParser =
  typeof fetch === 'function' && typeof new ArrayBuffer().resize === 'function';
```

```

tsx

const canUseMediaParser =
  typeof fetch === 'function' && typeof new ArrayBuffer().resize === 'function';
```

### Check if browser has `VideoDecoder` and `AudioDecoder` [​](\#check-if-browser-has-videodecoder-and-audiodecoder "Direct link to check-if-browser-has-videodecoder-and-audiodecoder")

Chrome has both `VideoDecoder` and `AudioDecoder`.

Firefox has support for `VideoDecoder` and `AudioDecoder` only if the `dom.media.webcodecs.enabled` flag is enabled.

Safari has support for `VideoDecoder`, but not `AudioDecoder`. You can decode the video track but not the audio track.

note

Please help [improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/webcodecs.mdx) if this information is outdated.

You can choose to not receive samples if the corresponding decoder is not supported in the browser.

```

Rejecting samples
tsx

import type {OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  if (typeof VideoDecoder === 'undefined') {
    return null;
  }

  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  // ...
};

const onAudioTrack: OnAudioTrack = ({track}) => {
  if (typeof AudioDecoder === 'undefined') {
    return null;
  }

  const audioDecoder = new AudioDecoder({
    output: console.log,
    error: console.error,
  });
  // ...
};
```

```

Rejecting samples
tsx

import type {OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  if (typeof VideoDecoder === 'undefined') {
    return null;
  }

  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  // ...
};

const onAudioTrack: OnAudioTrack = ({track}) => {
  if (typeof AudioDecoder === 'undefined') {
    return null;
  }

  const audioDecoder = new AudioDecoder({
    output: console.log,
    error: console.error,
  });
  // ...
};
```

### Check if the browser supports the codec [​](\#check-if-the-browser-supports-the-codec "Direct link to Check if the browser supports the codec")

Not all browsers support all codecs that `parseMedia()` emits.

The best way is to use `AudioDecoder.isConfigSupported()` and `VideoDecoder.isConfigSupported()` to check if the browser supports the codec.

These are async APIs, fortunately `onAudioTrack` and `onVideoTrack` allow async code as well.

```

Checking if the browser supports the codec
tsx

import type {OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  const {supported} = await VideoDecoder.isConfigSupported(track);
  if (!supported) {
    return null;
  }

  // ...
};

const onAudioTrack: OnAudioTrack = async ({track}) => {
  const audioDecoder = new AudioDecoder({
    output: console.log,
    error: console.error,
  });

  const {supported} = await AudioDecoder.isConfigSupported(track);
  if (!supported) {
    return null;
  }

  // ...
};
```

```

Checking if the browser supports the codec
tsx

import type {OnAudioTrack, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  const {supported} = await VideoDecoder.isConfigSupported(track);
  if (!supported) {
    return null;
  }

  // ...
};

const onAudioTrack: OnAudioTrack = async ({track}) => {
  const audioDecoder = new AudioDecoder({
    output: console.log,
    error: console.error,
  });

  const {supported} = await AudioDecoder.isConfigSupported(track);
  if (!supported) {
    return null;
  }

  // ...
};
```

note

Perform these checks in addition to the previously mentioned ones.

### Error handling [​](\#error-handling "Direct link to Error handling")

If an error occurs, you get the error in the `error` callback that you passed to the `VideoDecoder` or `AudioDecoder` constructor.

The decoder `state` will switch to `"closed"`, however, you will still receive samples.

If the decoder is in `"closed"` state, you should stop passing them to VideoDecoder.

```

Error handling
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  return async (sample) => {
    if (videoDecoder.state === 'closed') {
      return;
    }
  };
};
```

```

Error handling
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  return async (sample) => {
    if (videoDecoder.state === 'closed') {
      return;
    }
  };
};
```

note

- The same logic goes for `AudioDecoder`.
- You should still perform the checks previously mentioned, but they are omitted from this example.

### Queuing samples [​](\#queuing-samples "Direct link to Queuing samples")

Extracting samples is the fast part, decoding them is the slow part.

If too many samples are in the queue, it will negatively impact the performance of the page.

Fortunately, the parsing process can be temporarily paused while the decoder is busy.

For this, make the sample processing function async. Remotion will await it before processing the file further.

This will make it so that samples that are not yet needed are not kept in memory, keeping the decoding process efficient.

```

Only keeping 10 samples in the queue at a time
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  return async (sample) => {
    if (videoDecoder.decodeQueueSize > 10) {
      let resolve = () => {};

      const cb = () => {
        resolve();
      };

      await new Promise<void>((r) => {
        resolve = r;
        videoDecoder.addEventListener('dequeue', cb);
      });
      videoDecoder.removeEventListener('dequeue', cb);
    }

    videoDecoder.decode(new EncodedVideoChunk(sample));
  };
};
```

```

Only keeping 10 samples in the queue at a time
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });

  return async (sample) => {
    if (videoDecoder.decodeQueueSize > 10) {
      let resolve = () => {};

      const cb = () => {
        resolve();
      };

      await new Promise<void>((r) => {
        resolve = r;
        videoDecoder.addEventListener('dequeue', cb);
      });
      videoDecoder.removeEventListener('dequeue', cb);
    }

    videoDecoder.decode(new EncodedVideoChunk(sample));
  };
};
```

note

- The same logic goes for `AudioDecoder`.
- You should still perform the checks previously mentioned, but they are omitted from this example.

### Handling stretched videos [​](\#handling-stretched-videos "Direct link to Handling stretched videos")

Some videos don't have the same dimensions internally as they are presented.

For example, [this sample video](https://github.com/remotion-dev/remotion/blob/main/packages/example/public/stretched-vp8.webm) has a coded width of 1440, but a presentation width of 1920.

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  videoDecoder.configure(track);

  return async (sample) => {
    console.log(sample);
    // {
    //   codedWidth: 1440,
    //   codedHeight: 1080,
    //   displayAspectWidth: 1920,
    //   displayAspectHeight: 1080,
    //   ...
    // }
  };
};
```

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  const videoDecoder = new VideoDecoder({
    output: console.log,
    error: console.error,
  });
  videoDecoder.configure(track);

  return async (sample) => {
    console.log(sample);
    // {
    //   codedWidth: 1440,
    //   codedHeight: 1080,
    //   displayAspectWidth: 1920,
    //   displayAspectHeight: 1080,
    //   ...
    // }
  };
};
```

This means the frame is internally encoded in a 4:3 aspect ratio, but the frame should as a 16:9.

By passing all of `codedWidth`, `codedHeight`, `displayAspectWidth` and `displayAspectHeight` to `new EncodedVideoChunk()`, the decoder should handle the stretching correctly.

### Handling rotation [​](\#handling-rotation "Direct link to Handling rotation")

WebCodecs do not seem to consider rotation.

For example, this [video recorded with an iPhone](https://github.com/remotion-dev/remotion/blob/main/packages/example/public/iphone-hevc.mov) has metadata that it should be displayed at 90 degrees rotation.

`VideoDecoder` is not able to rotate the video for you, so you might need to do it yourself, for example by drawing it to a canvas.

Fortunately `parseMedia()` returns the rotation of the track:

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  console.log(track.rotation); // -90
  return null;
};
```

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  console.log(track.rotation); // -90
  return null;
};
```

See here for an example of how a [video frame is turned into a bitmap](https://github.com/remotion-dev/remotion/blob/1a99dadcd8e700bf94abdfa3f5506329d6c9c182/packages/example/src/Encoder/SrcEncoder.tsx#L98-L102) and rotated.

### Understanding the different dimensions of a video [​](\#understanding-the-different-dimensions-of-a-video "Direct link to Understanding the different dimensions of a video")

As just mentioned, some videos might be stretched or rotated.

In an extreme case, it is possible that you stumble opon a video that has three different dimensions.

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  console.log(track);
  // {
  //   codedWidth: 1440,
  //   codedHeight: 1080,
  //   displayAspectWidth: 1920,
  //   displayAspectHeight: 1080,
  //   width: 1080,
  //   height: 1900,
  //   ...
  // }

  return null;
};
```

```

Handling stretched videos
tsx

import type {OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = async ({track}) => {
  console.log(track);
  // {
  //   codedWidth: 1440,
  //   codedHeight: 1080,
  //   displayAspectWidth: 1920,
  //   displayAspectHeight: 1080,
  //   width: 1080,
  //   height: 1900,
  //   ...
  // }

  return null;
};
```

The meaning of it is:

- `codedWidth` and `codedHeight` are the dimensions of the video in the codec's internal format.
- `displayAspectWidth` and `displayAspectHeight` are scaled dimensions of how the video should be displayed, but with rotation not yet applied.

note



These are not necessarily the actual dimensions of how a video is presented to the user, because rotation is not yet applied.

The fields are named like this because they correspond with what should be passed to `new EncodedVideoChunk()`.

- `width` and `height` are the dimensions of the video how it would be displayed by a Player.

### Google Chrome quirks [​](\#google-chrome-quirks "Direct link to Google Chrome quirks")

We find that as of now, `AudioDecoder.isConfigSupported()` are not 100% reliable. For example, Chrome marks this config as supported, but then throws an error nonetheless.

```

tsx

const config = {codec: 'opus', numberOfChannels: 6, sampleRate: 44100};
console.log(await AudioDecoder.isConfigSupported(config)); // {supported: true}

const decoder = new AudioDecoder({error: console.error, output: console.log});

decoder.configure(config); // Unsupported configuration. Check isConfigSupported() prior to calling configure().
```

```

tsx

const config = {codec: 'opus', numberOfChannels: 6, sampleRate: 44100};
console.log(await AudioDecoder.isConfigSupported(config)); // {supported: true}

const decoder = new AudioDecoder({error: console.error, output: console.log});

decoder.configure(config); // Unsupported configuration. Check isConfigSupported() prior to calling configure().
```

Consider this in your implementation.

### Safari performance [​](\#safari-performance "Direct link to Safari performance")

We find that with our reference implementation, Safari chokes on decoding the full Big Buck Bunny video. Tips are welcome, and otherwise we encourage to consider if and which parts of WebCodecs APIs you want to support.

## Reference implementation [​](\#reference-implementation "Direct link to Reference implementation")

A testbed with many different codecs and edge cases is available [here](https://github.com/remotion-dev/remotion/blob/main/packages/example/src/DecoderDemo.tsx).

Follow [these instructions](https://www.remotion.dev/docs/contributing#testing-your-changes) to run the testbed locally.

## License reminder [​](\#license-reminder "Direct link to License reminder")

Like Remotion itself, this package is licensed under the [Remotion License](https://remotion.dev/license).

TL;DR: Individuals and small teams can use this package, but teams of 4+ people [need a company license](https://remotion.pro/license).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/webcodecs.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Runtime support](/docs/media-parser/support) [Next\
\
Foreign file types](/docs/media-parser/foreign-file-types)

- [With `@remotion/webcodecs`](#with-remotionwebcodecs)
- [Using `@remotion/media-parser`](#using-remotionmedia-parser)
- [Why is this useful?](#why-is-this-useful)
- [Will Remotion switch to WebCodecs?](#will-remotion-switch-to-webcodecs)
- [Practical considerations](#practical-considerations)
  - [Check browser support for `@remotion/media-parser`](#check-browser-support-for-remotionmedia-parser)
  - [Check if browser has `VideoDecoder` and `AudioDecoder`](#check-if-browser-has-videodecoder-and-audiodecoder)
  - [Check if the browser supports the codec](#check-if-the-browser-supports-the-codec)
  - [Error handling](#error-handling)
  - [Queuing samples](#queuing-samples)
  - [Handling stretched videos](#handling-stretched-videos)
  - [Handling rotation](#handling-rotation)
  - [Understanding the different dimensions of a video](#understanding-the-different-dimensions-of-a-video)
  - [Google Chrome quirks](#google-chrome-quirks)
  - [Safari performance](#safari-performance)
- [Reference implementation](#reference-implementation)
- [License reminder](#license-reminder)