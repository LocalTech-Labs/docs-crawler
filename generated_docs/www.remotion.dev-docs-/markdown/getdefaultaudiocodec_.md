---
title: getDefaultAudioCodec()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getDefaultAudioCodec()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- getDefaultAudioCodec()

On this page

# getDefaultAudioCodec()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Gets the default audio codec for a container that `@remotion/webcodecs` uses if no other audio codec was specified.

```

Get the default audio codec for a container
tsx

import {getDefaultAudioCodec} from '@remotion/webcodecs';

getDefaultAudioCodec({container: 'webm'}); // 'opus'
```

```

Get the default audio codec for a container
tsx

import {getDefaultAudioCodec} from '@remotion/webcodecs';

getDefaultAudioCodec({container: 'webm'}); // 'opus'
```

Currently, the only supported container is `webm`, for which the default audio codec is `opus`.

## Default audio codecs [​](\#default-audio-codecs "Direct link to Default audio codecs")

ContainerDefault audio codecwebm`"opus"`mp4`"aac"`wav`"wav"`

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/get-default-audio-codec.ts)
- [`getDefaultVideoCodec()`](/docs/webcodecs/get-default-video-codec)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/get-default-audio-codec.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
defaultOnVideoTrackHandler()](/docs/webcodecs/default-on-video-track-handler) [Next\
\
getDefaultVideoCodec()](/docs/webcodecs/get-default-video-codec)

- [Default audio codecs](#default-audio-codecs)
- [See also](#see-also)
