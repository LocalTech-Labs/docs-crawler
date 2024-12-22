---
title: @remotion/webcodecs
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/webcodecs

- [Home page](/)
- WebCodecs

On this page

# @remotion/webcodecs

_available from v4.0.229_

This package provides APIs for converting videos in the browser.

It leverages [`@remotion/media-parser`](/docs/media-parser) to parse the video and audio data, and then uses the [WebCodecs API](https://developer.mozilla.org/en-US/docs/Web/API/WebCodecs_API) to encode the video.

## What can you to with this package? [​](\#what-can-you-to-with-this-package "Direct link to What can you to with this package?")

In browsers that implement WebCodecs, you can use this package to:

- [Convert videos from one format to another](/docs/webcodecs/convert-a-video) (.mp4 and .webm bidirectional, .avi, .mkv, .mov as import)
- [Rotate videos](/docs/webcodecs/rotate-a-video)
- Extract audio from a video
- Manipulate the pixels of a video
- [Fix videos that were recorded with `MediaRecorder`](/docs/webcodecs/fix-mediarecorder-video)
- Soon: Compress, trim, crop videos

## 💼 License Disclaimer [​](\#-license-disclaimer "Direct link to 💼 License Disclaimer")

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

## 🚧 Unstable API Warning [​](\#-unstable-api-warning "Direct link to 🚧 Unstable API Warning")

This package is experimental.

We might change the API at any time, until we remove this notice.

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/webcodecs@4.0.242Copy
```

```

npm i --save-exact @remotion/webcodecs@4.0.242Copy
```

```

pnpm i @remotion/webcodecs@4.0.242Copy
```

```

pnpm i @remotion/webcodecs@4.0.242Copy
```

```

bun i @remotion/webcodecs@4.0.242Copy
```

```

bun i @remotion/webcodecs@4.0.242Copy
```

```

yarn --exact add @remotion/webcodecs@4.0.242Copy
```

```

yarn --exact add @remotion/webcodecs@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## Guide [​](\#guide "Direct link to Guide")

[**Convert a video** \
\
from one format to another](/docs/webcodecs/convert-a-video) [**Rotate a video** \
\
Fix bad orientation](/docs/webcodecs/rotate-a-video) [**Fix a MediaRecorder video** \
\
Fix missing video duration and poor seeking performance](/docs/webcodecs/fix-mediarecorder-video)

## APIs [​](\#apis "Direct link to APIs")

The following APIs are available:

[**convertMedia()** \
\
Converts a video using WebCodecs and Media Parser](/docs/webcodecs/convert-media) [**getAvailableContainers()** \
\
Get a list of containers `@remotion/webcodecs` supports.](/docs/webcodecs/get-available-containers) [**canReencodeVideoTrack()** \
\
Determine if a video track can be re-encoded](/docs/webcodecs/can-reencode-video-track) [**canReencodeAudioTrack()** \
\
Determine if a audio track can be re-encoded](/docs/webcodecs/can-reencode-audio-track) [**canCopyVideoTrack()** \
\
Determine if a video track can be copied without re-encoding](/docs/webcodecs/can-copy-video-track) [**canCopyAudioTrack()** \
\
Determine if a audio track can be copied without re-encoding](/docs/webcodecs/can-copy-audio-track) [**getDefaultAudioCodec()** \
\
Gets the default audio codec for a container if no other audio codec is specified.](/docs/webcodecs/get-default-audio-codec) [**getDefaultVideoCodec()** \
\
Gets the default video codec for a container if no other audio codec is specified.](/docs/webcodecs/get-default-video-codec) [**defaultOnAudioTrackHandler()** \
\
The default track transformation function for audio tracks.](/docs/webcodecs/default-on-audio-track-handler) [**defaultOnVideoTrackHandler()** \
\
The default track transformation function for video tracks.](/docs/webcodecs/default-on-video-track-handler) [**getAvailableAudioCodecs()** \
\
Get the audio codecs that can fit in a container.](/docs/webcodecs/get-available-audio-codecs) [**getAvailableVideoCodecs()** \
\
Get the video codecs that can fit in a container.](/docs/webcodecs/get-available-video-codecs)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Foreign file types](/docs/media-parser/foreign-file-types) [Next\
\
Convert a video](/docs/webcodecs/convert-a-video)

- [What can you to with this package?](#what-can-you-to-with-this-package)
- [💼 License Disclaimer](#-license-disclaimer)
- [🚧 Unstable API Warning](#-unstable-api-warning)
- [Installation](#installation)
- [Guide](#guide)
- [APIs](#apis)
