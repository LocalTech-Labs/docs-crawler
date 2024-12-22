---
title: getDefaultVideoCodec()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getDefaultVideoCodec()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- getDefaultVideoCodec()

On this page

# getDefaultVideoCodec()

_Part of the [`@remotion/webcodecs`](/docs/webcodecs) package._

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Gets the default video codec for a container that `@remotion/webcodecs` uses if no other audio codec was specified.

```

Get the default video codec for a container
tsx

import {getDefaultVideoCodec} from '@remotion/webcodecs';

getDefaultVideoCodec({container: 'webm'}); // 'vp8'
```

```

Get the default video codec for a container
tsx

import {getDefaultVideoCodec} from '@remotion/webcodecs';

getDefaultVideoCodec({container: 'webm'}); // 'vp8'
```

## Default video codecs [​](\#default-video-codecs "Direct link to Default video codecs")

ContainerDefault video codecwebm`"vp8"`mp4`"h264"`wav`null`

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/get-default-video-codec.ts)
- [`getDefaultAudioCodec()`](/docs/webcodecs/get-default-audio-codec)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/get-default-video-codec.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getDefaultAudioCodec()](/docs/webcodecs/get-default-audio-codec) [Next\
\
getAvailableAudioCodecs()](/docs/webcodecs/get-available-audio-codecs)

- [Default video codecs](#default-video-codecs)
- [See also](#see-also)
