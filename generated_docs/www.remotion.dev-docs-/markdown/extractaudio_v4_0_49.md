---
title: extractAudio()v4.0.49
source: Remotion Documentation
last_updated: 2024-12-22
---

# extractAudio()v4.0.49

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- extractAudio()

On this page

# extractAudio() [v4.0.49](https://github.com/remotion-dev/remotion/releases/v4.0.49)

note

This function is meant to be used **in Node.js applications**. It cannot run in the browser.

Extracts the audio from a video source and saves it to the specified output path. It does not convert the audio to a different format.

## Example [​](\#example "Direct link to Example")

```

ts

import { resolve } from "node:path";
import { extractAudio, getVideoMetadata } from "@remotion/renderer";

const videoSource = resolve(process.cwd(), "/Users/john/path-to-video.mp4");

const videoMetadata = await getVideoMetadata(videoSource);
const audioOutput = resolve(
  process.cwd(),
  `./output-audio-path.${videoMetadata.audioFileExtension}`,
);

await extractAudio({
  videoSource,
  audioOutput,
});
```

```

ts

import { resolve } from "node:path";
import { extractAudio, getVideoMetadata } from "@remotion/renderer";

const videoSource = resolve(process.cwd(), "/Users/john/path-to-video.mp4");

const videoMetadata = await getVideoMetadata(videoSource);
const audioOutput = resolve(
  process.cwd(),
  `./output-audio-path.${videoMetadata.audioFileExtension}`,
);

await extractAudio({
  videoSource,
  audioOutput,
});
```

info

Pass an absolute path to `extractAudio()`. URLs are not supported.

## Arguments [​](\#arguments "Direct link to Arguments")

An object containing the following properties:

### `videoSource` [​](\#videosource "Direct link to videosource")

_string_

The path to the video source from which the audio will be extracted.

### `outputPath` [​](\#outputpath "Direct link to outputpath")

_string_

The path where the extracted audio will be saved. The file extension must match the audio codec. To find the appropriate file extension, use [`getVideoMetadata()`](/docs/renderer/get-video-metadata) to read the field `audioFileExtension`.

### `logLevel?` [​](\#loglevel "Direct link to loglevel")

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

### `binariesDirectory?` [v4.0.120](https://github.com/remotion-dev/remotion/releases/v4.0.120) [​](\#binariesdirectory "Direct link to binariesdirectory")

The directory where the platform-specific binaries and libraries that Remotion needs are located. Those include an `ffmpeg` and `ffprobe` binary, a Rust binary for various tasks, and various shared libraries. If the value is set to `null`, which is the default, then the path of a platform-specific package located at `node_modules/@remotion/compositor-*` is selected.

This option is useful in environments where Remotion is not officially supported to run like bundled serverless functions or Electron.

## Return Value [​](\#return-value "Direct link to Return Value")

The function returns a `Promise<void>`, which resolves once the audio extraction is complete.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/renderer/src/extract-audio.ts)
- [getVideoMetadata()](/docs/renderer/get-video-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/extract-audio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getSilentParts()](/docs/renderer/get-silent-parts) [Next\
\
Installation](/docs/player/installation)

- [Example](#example)
- [Arguments](#arguments)
  - [`videoSource`](#videosource)
  - [`outputPath`](#outputpath)
  - [`logLevel?`](#loglevel)
  - [`binariesDirectory?`](#binariesdirectory)
- [Return Value](#return-value)
- [See also](#see-also)
