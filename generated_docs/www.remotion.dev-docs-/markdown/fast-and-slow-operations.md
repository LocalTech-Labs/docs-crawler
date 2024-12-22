---
title: Fast and slow operations
source: Remotion Documentation
last_updated: 2024-12-22
---

# Fast and slow operations

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- Fast and slow operations

On this page

# Fast and slow operations

[`@remotion/media-parser`](/docs/media-parser) allows you to specify the information that you want to obtain.

It then reads as little data as possible to achieve the goal.

There are three types of fields:

- Header-only: Only requires the first few bits of the file to be read.
- Metadata-only: Only requires the metadata section to be parsed
- Full parse required: The entire file is read and processed.

Obviously, processing less of the file is faster and you should aim to read only the information you require.

## Full parsing operations [​](\#full-parsing-operations "Direct link to Full parsing operations")

The following [`fields`](/docs/media-parser#fields) require the full file to be read:

- [`structure`](/docs/media-parser/parse-media#structure)
- [`slowKeyframes`](/docs/media-parser/parse-media#slowkeyframes)
- [`slowFps`](/docs/media-parser/parse-media#slowfps)
- [`slowDurationInSeconds`](/docs/media-parser/parse-media#slowdurationinseconds)
- [`slowNumberOfFrames`](/docs/media-parser/parse-media#slownumberofframes)

Also, if an [`onVideoTrack`](/docs/media-parser/parse-media#onvideotrack) or [`onAudioTrack`](/docs/media-parser/parse-media#onvideotrack) handler is passed, and the handler function returns an callback function for each sample, full parsing is required.

Also, if [`convertMedia()`](/docs/webcodecs/convert-media) is used, full parsing is always required.

## Metadata-only operations [​](\#metadata-only-operations "Direct link to Metadata-only operations")

The following [`fields`](/docs/media-parser#fields) require only the metadata section of the video to be parsed:

- [`dimensions`](/docs/media-parser/parse-media#dimensions)
- [`durationInSeconds`](/docs/media-parser/parse-media#durationinseconds)
- [`fps`](/docs/media-parser/parse-media#fps)
- [`videoCodec`](/docs/media-parser/parse-media#videocodec)
- [`audioCodec`](/docs/media-parser/parse-media#audiocodec)
- [`tracks`](/docs/media-parser/parse-media#tracks)
- [`unrotatedDimensions`](/docs/media-parser/parse-media#unrotateddimensions)
- [`isHdr`](/docs/media-parser/parse-media#ishdr)
- [`rotation`](/docs/media-parser/parse-media#rotation)
- [`metadata`](/docs/media-parser/parse-media#metadata)
- [`location`](/docs/media-parser/parse-media#location)
- [`keyframes`](/docs/media-parser/parse-media#keyframes)

Also, if an [`onVideoTrack`](/docs/media-parser/parse-media#onvideotrack) or [`onAudioTrack`](/docs/media-parser/parse-media#onvideotrack) handler is passed, only the parsing of the metadata section is required if `null` is returned from the handler function.

## Header-only operations [​](\#header-only-operations "Direct link to Header-only operations")

The following [`fields`](/docs/media-parser/parse-media#fields) require only the first few bytes of the media to be parsed:

- [`name`](/docs/media-parser/parse-media#name)
- [`size`](/docs/media-parser/parse-media#size)
- [`container`](/docs/media-parser/parse-media#container)
- [`mimeType`](/docs/media-parser/parse-media#mimetype)

## Seeking required [​](\#seeking-required "Direct link to Seeking required")

If you load videos from a URL, make sure that they support the `Range` header.

Otherwise, `@remotion/media-parser` has no choice but to read the full file if the metadata is at the end of it.

## Example [​](\#example "Direct link to Example")

```

Reading header only
tsx

// Some fields only require the first few bytes of the file to be read:
const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    size: true,
    container: true,
    internalStats: true,
  },
});

console.log(result.internalStats.finalCursorOffset); // 12

// Reading the metadata of the video will only require the metadata section to be parsed.
// You can also use onVideoTrack() and return null to retrieve track information but to not get the samples.
const result2 = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
    internalStats: true,
  },
  onVideoTrack: ({track}) => {
    console.log(track);
    return null;
  },
});

console.log(result2.internalStats.finalCursorOffset); // 4000
console.log(result2.dimensions);

// Asking for all video samples requires parsing the whole file
const result3 = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    internalStats: true,
  },
  onVideoTrack: () => {
    return (videoSample) => console.log(videoSample);
  },
});

console.log(result3.internalStats.finalCursorOffset); // 1870234802
```

```

Reading header only
tsx

// Some fields only require the first few bytes of the file to be read:
const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    size: true,
    container: true,
    internalStats: true,
  },
});

console.log(result.internalStats.finalCursorOffset); // 12

// Reading the metadata of the video will only require the metadata section to be parsed.
// You can also use onVideoTrack() and return null to retrieve track information but to not get the samples.
const result2 = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
    internalStats: true,
  },
  onVideoTrack: ({track}) => {
    console.log(track);
    return null;
  },
});

console.log(result2.internalStats.finalCursorOffset); // 4000
console.log(result2.dimensions);

// Asking for all video samples requires parsing the whole file
const result3 = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    internalStats: true,
  },
  onVideoTrack: () => {
    return (videoSample) => console.log(videoSample);
  },
});

console.log(result3.internalStats.finalCursorOffset); // 1870234802
```

## See also [​](\#see-also "Direct link to See also")

- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/fast-and-slow.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Getting video metadata](/docs/media-parser/metadata) [Next\
\
Extract ID3 tags](/docs/media-parser/tags)

- [Full parsing operations](#full-parsing-operations)
- [Metadata-only operations](#metadata-only-operations)
- [Header-only operations](#header-only-operations)
- [Seeking required](#seeking-required)
- [Example](#example)
- [See also](#see-also)