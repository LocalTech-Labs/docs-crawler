---
title: getCanExtractFramesFast()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getCanExtractFramesFast()

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- getCanExtractFramesFast()

On this page

# getCanExtractFramesFast()

warning

From v4.0 on, frames can always be extracted fast and therefore the function has been removed. The information in this document only applies to older versions of Remotion and is preserved for people who are still using them.

_Available since v3.3.2, removed in v4.0 - Part of the `@remotion/renderer` package._

Probes whether frames of a video can be efficiently extracted when using [`<OffthreadVideo>`](/docs/offthreadvideo).

```

ts

// @module: ESNext
// @target: ESNext
import { getCanExtractFramesFast } from "@remotion/renderer";

const result = await getCanExtractFramesFast({
  src: "/var/path/to/video.mp4",
});

console.log(result.canExtractFramesFast); // false
console.log(result.shouldReencode); // true
```

```

ts

// @module: ESNext
// @target: ESNext
import { getCanExtractFramesFast } from "@remotion/renderer";

const result = await getCanExtractFramesFast({
  src: "/var/path/to/video.mp4",
});

console.log(result.canExtractFramesFast); // false
console.log(result.shouldReencode); // true
```

info

Pass an absolute path to `getCanExtractFramesFast()`. URLs are not supported.

## When to use this API [​](\#when-to-use-this-api "Direct link to When to use this API")

If you are using [`<OffthreadVideo>`](/docs/offthreadvideo), you might get a warning ["Using a slow method to extract the frame"](/docs/slow-method-to-extract-frame) if a video is included which does not include enough metadata to efficiently extract a certain frame of a video. This might result in the render becoming slow.

Using this API, you can probe whether this issue affects your video file. It will try to extract the last frame of a video and if it succeeds, your video is not affected. Otherwise, `canExtractFramesFast` will be `false`.

## How to act on the results [​](\#how-to-act-on-the-results "Direct link to How to act on the results")

When `canExtractFramesFast` is `false`, you should check the `shouldReencode` flag. If it is true, you can re-encode the video to make the render faster. Note that it is not always faster to re-encode the video than it is to deal with a slow render.

Videos with a VP8 codec don't support fast frame extraction at all, and therefore `shouldReencode` can be false even if `canExtractFramesFast` is false.

## Reencoding a video [​](\#reencoding-a-video "Direct link to Reencoding a video")

You can re-encode a video using FFmpeg:

```

sh

ffmpeg -i inputvideo.mp4 outputvideo.mp4
```

```

sh

ffmpeg -i inputvideo.mp4 outputvideo.mp4
```

## Arguments [​](\#arguments "Direct link to Arguments")

An object containing one or more of the following options:

### `src` [​](\#src "Direct link to src")

Pointing to a video file. Must be an absolute file path.

### `ffmpegExecutable?` [​](\#ffmpegexecutable "Direct link to ffmpegexecutable")

_removed in v4.0_

_string - optional_

An absolute path overriding the `ffmpeg` executable to use.

### `ffprobeExecutable?` [​](\#ffprobeexecutable "Direct link to ffprobeexecutable")

_removed in v4.0, \_string, optional_

An absolute path overriding the `ffprobe` executable to use.

## Return value [​](\#return-value "Direct link to Return value")

Returns a promise which resolves to an object with the following parameters:

- `canExtractFramesFast`: _boolean_ Whether it will be fast to extract a frame from a video.
- `shouldReencode`: _boolean_ Whether the video can be re-encoded to make the render faster.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/get-can-extract-frames-fast.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
ensureFfprobe()](/docs/renderer/ensure-ffprobe) [Next\
\
getVideoMetadata()](/docs/renderer/get-video-metadata)

- [When to use this API](#when-to-use-this-api)
- [How to act on the results](#how-to-act-on-the-results)
- [Reencoding a video](#reencoding-a-video)
- [Arguments](#arguments)
  - [`src`](#src)
  - [`ffmpegExecutable?`](#ffmpegexecutable)
  - [`ffprobeExecutable?`](#ffprobeexecutable)
- [Return value](#return-value)
