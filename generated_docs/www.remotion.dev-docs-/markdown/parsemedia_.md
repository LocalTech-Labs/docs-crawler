---
title: parseMedia()
source: Remotion Documentation
last_updated: 2024-12-22
---

# parseMedia()

- [Home page](/)
- [@remotion/media-parser](/docs/media-parser/)
- parseMedia()

On this page

# parseMedia()

_Part of the [`@remotion/media-parser`](/docs/media-parser) package._ _available from v4.0.190_

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

## Examples [​](\#examples "Direct link to Examples")

```

Parsing a hosted video
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a hosted video
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});
```

```

Parsing a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});
```

## API [​](\#api "Direct link to API")

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

### `src` [​](\#src "Direct link to src")

Either a local file path, or a URL, or a `File` or `Blob` object.

If you pass a local file path, you must also pass [`nodeReader`](/docs/media-parser/node-reader) as the `reader` argument.

If you pass a `File` object, you must also pass [`webFileReader`](/docs/media-parser/web-file-reader) as the `reader` argument.

### `fields?` [​](\#fields "Direct link to fields")

An object specifying which fields you'd like to receive.

If you like to receive the field, pass `true` as the value.

Possible fields are:

#### `dimensions` [​](\#dimensions "Direct link to dimensions")

_`{width: number, height: number}`_

The dimensions of the video.

Any rotation is already applied - the dimensions are like a media player would show them.

Use `unrotatedDimensions` to get the dimensions before rotation.

#### `durationInSeconds` [​](\#durationinseconds "Direct link to durationinseconds")

_number \| null_

The duration of the video in seconds.

Only returns a non-null value if the duration is stored in the metadata.

#### `slowDurationInSeconds` [​](\#slowdurationinseconds "Direct link to slowdurationinseconds")

_number_

The duration of the video in seconds, but it is guaranteed to return a value.

If needed, the entire video file is read to determine the duration.

#### `name` [​](\#name "Direct link to name")

_string_

The name of the file.

#### `container` [​](\#container "Direct link to container")

_"mp4" \| "webm" \| "avi" \| null_

The container of the file.

#### `size` [​](\#size "Direct link to size")

_number \| null_

The size of the input in bytes.

#### `mimeType` [​](\#mimetype "Direct link to mimetype")

_string \| null_

The MIME type of the file that was returned when the file was fetched.

Only available if using the `fetchReader` or `webFileReader`.

#### `structure` [​](\#structure "Direct link to structure")

The internal structure of the video. Unstable, internal data structure, refer to the TypeScript types to see what's inside.

#### `fps` [​](\#fps "Direct link to fps")

_number \| null_

The frame rate of the video.

Only returns a non-null value if the frame rate is stored in the metadata.

#### `slowFps` [​](\#slowfps "Direct link to slowfps")

_number_

The frame rate of the video, but it is guaranteed to return a value.

If needed, the entire video file is read to determine the frame rate.

#### `videoCodec` [​](\#videocodec "Direct link to videocodec")

The video codec of the file.

If multiple video tracks are present, this will be the first video track.

One of `"h264"`, `"h265"`, `"vp8"`, `"vp9"`, `"av1"`, `"prores"` or `null` (in case of an unknown codec).

#### `audioCodec` [​](\#audiocodec "Direct link to audiocodec")

The audio codec of the file.

If multiple audio tracks are present, this will be the first audio track.

One of `'aac'`, `'mp3'`, `'aiff'`, `'opus'`, `'pcm'`, `'unknown'` (audio is there but not recognized) or `null` (in case of no audio detected).

#### `metadata` [​](\#metadata "Direct link to metadata")

Metadata fields such as ID3 tags or EXIF data.

See [metadata](/docs/media-parser/tags) for more information.

#### `location` [​](\#location "Direct link to location")

The location of the video was shot. Either `null` if not available or:

- `latitude`: The latitude of the location
- `longitude`: The longitude of the location
- `altitude`: The altitude of the location (can be `null`)
- `horizontalAccuracy`: The horizontal accuracy of the location (can be `null`)

#### `tracks` [​](\#tracks "Direct link to tracks")

Returns an object of two two arrays `videoTracks` and `audioTracks`.

The data structure of them is not yet stable.

#### `keyframes` [​](\#keyframes "Direct link to keyframes")

Return type: `MediaParserKeyframe[] | null`

An array of keyframes. Each keyframe has the following structure:

- `presentationTimeInSeconds`: The time in seconds when the keyframe should be presented
- `decodingTimeInSeconds`: The time in seconds when the keyframe should be decoded
- `positionInBytes`: The position in bytes where the keyframe is located in the file
- `sizeInBytes`: The size of the keyframe in bytes
- `trackId`: The ID of the track the frame belongs to

Only being returned if the keyframe information are stored in the metadata, otherwise `null`.

#### `slowKeyframes` [​](\#slowkeyframes "Direct link to slowkeyframes")

Return type: `MediaParserKeyframe[]`

An array of keyframes, same as [`keyframes`](#keyframes), but it is guaranteed to return a value.

Will read the entire video file to determine the keyframes.

#### `slowNumberOfFrames` [​](\#slownumberofframes "Direct link to slownumberofframes")

_number_

The number of video frames in the media.

Will read the entire video file to determine the number of frames.

#### `unrotatedDimensions` [​](\#unrotateddimensions "Direct link to unrotateddimensions")

_`{width: number, height: number}`_

The dimensions of the video before rotation.

#### `isHdr` [​](\#ishdr "Direct link to ishdr")

_`boolean`_

Whether the video is in HDR (High dynamic range).

#### `rotation` [​](\#rotation "Direct link to rotation")

_number_

The rotation of the video in degrees (e.g. `-90` for a 90° counter-clockwise rotation).

### `reader?` [​](\#reader "Direct link to reader")

A reader interface.

Default value: `fetchReader`, which uses `fetch()` to read the video.

If you pass [`nodeReader`](/docs/media-parser/node-reader), you must also pass a local file path as the `src` argument.
If you pass [`webFileReader`](/docs/media-parser/web-file-reader), you must also pass a `File` as the `src` argument.

### `onVideoTrack?` [​](\#onvideotrack "Direct link to onvideotrack")

A callback that is called when a video track is detected.

It receives an object with `track` and `container` (API not yet stable).

You must return either `null` or a callback that is called for each sample that corresponds to the video track.

The `sample` has the type `VideoSample` and while not all fields are stable, it has all the mandatory fields for the [`EncodedVideoChunk`](https://developer.mozilla.org/en-US/docs/Web/API/EncodedVideoChunk) constructor.

```

Reading video frames
tsx

import {parseMedia, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  console.log(track);

  return (sample) => {
    console.log(new EncodedVideoChunk(sample));
  };
};
```

```

Reading video frames
tsx

import {parseMedia, OnVideoTrack} from '@remotion/media-parser';

const onVideoTrack: OnVideoTrack = ({track}) => {
  console.log(track);

  return (sample) => {
    console.log(new EncodedVideoChunk(sample));
  };
};
```

### `onAudioTrack?` [​](\#onaudiotrack "Direct link to onaudiotrack")

A callback that is called when an audio track is detected.

It receives an object with `track` and `container` (API not yet stable).

You must return either `null` or a callback that is called for each sample that corresponds to the audio track.

The `sample` has the type `AudioSample` and while not all fields are stable, it has all the mandatory fields for the [`EncodedAudioChunk`](https://developer.mozilla.org/en-US/docs/Web/API/EncodedAudioChunk) constructor.

```

Reading audio frames
tsx

import {parseMedia, OnAudioTrack} from '@remotion/media-parser';

const onAudioTrack: OnAudioTrack = ({track}) => {
  console.log(track);

  return (sample) => {
    console.log(new EncodedAudioChunk(sample));
  };
};
```

```

Reading audio frames
tsx

import {parseMedia, OnAudioTrack} from '@remotion/media-parser';

const onAudioTrack: OnAudioTrack = ({track}) => {
  console.log(track);

  return (sample) => {
    console.log(new EncodedAudioChunk(sample));
  };
};
```

### `onParseProgress?` [​](\#onparseprogress "Direct link to onparseprogress")

A callback that is called from time to time when bytes have been read from the media.

It includes the following data:

```

tsx

import {ParseMediaProgress} from '@remotion/media-parser';

(alias) type ParseMediaProgress = {
    bytes: number;
    percentage: number | null;
    totalBytes: number | null;
}
import ParseMediaProgress
```

```

tsx

import {ParseMediaProgress} from '@remotion/media-parser';

(alias) type ParseMediaProgress = {
    bytes: number;
    percentage: number | null;
    totalBytes: number | null;
}
import ParseMediaProgress
```

note

You may make this an async function, and while it is not resolved, **the parsing process will be paused**.

This is useful if you want to add a small delay inbetween progress steps to keep the UI interactive.

### `logLevel?` [​](\#loglevel "Direct link to loglevel")

One of `"error"`, `"warn"`, `"info"`, `"debug"`, `"trace"`.

Default value: `"info"`, which logs only important information.

### `signal?` [​](\#signal "Direct link to signal")

An [`AbortSignal`](https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal) instance.

If the signal is aborted, the parser will stop reading the media and stop the decoding process and throw an error.

### Callbacks [​](\#callbacks "Direct link to Callbacks")

Each field also has a callback that allows you to retrieve the value as soon as it is obtained without waiting for the function to resolve.

You do not have to add the field to the [`fields`](#fields) object if you use the callback.

However, just like with [`fields`](#fields), adding a callback for a [slow field](/docs/media-parser/fast-and-slow) may require reading more of the file.

```

Using a callback
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  onDurationInSeconds: (durationInSeconds) => {
    console.log(durationInSeconds);
  },
  onDimensions: (dimensions) => {
    console.log(dimensions);
  },
});
```

```

Using a callback
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  onDurationInSeconds: (durationInSeconds) => {
    console.log(durationInSeconds);
  },
  onDimensions: (dimensions) => {
    console.log(dimensions);
  },
});
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/media-parser/src/parse-media.ts)
- [`@remotion/media-parser`](/docs/media-parser)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/parse-media.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Overview](/docs/media-parser/) [Next\
\
nodeReader](/docs/media-parser/node-reader)

- [Examples](#examples)
- [API](#api)
  - [`src`](#src)
  - [`fields?`](#fields)
  - [`reader?`](#reader)
  - [`onVideoTrack?`](#onvideotrack)
  - [`onAudioTrack?`](#onaudiotrack)
  - [`onParseProgress?`](#onparseprogress)
  - [`logLevel?`](#loglevel)
  - [`signal?`](#signal)
  - [Callbacks](#callbacks)
- [See also](#see-also)