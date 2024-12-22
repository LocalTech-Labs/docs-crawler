---
title: getAvailableVideoCodecs()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAvailableVideoCodecs()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- getAvailableVideoCodecs()

On this page

# getAvailableVideoCodecs()

Given a container, get a list of video codecs that the container can hold.

This does not mean that a any video stream of this codec can be put into the container.

Use [`canReencodeVideoTrack()`](/docs/webcodecs/can-reencode-video-track) and [`canCopyVideoTrack()`](/docs/webcodecs/can-copy-video-track) to determine this.

```

Get available video codecs for a container
tsx

import {getAvailableVideoCodecs} from '@remotion/webcodecs';

getAvailableVideoCodecs({container: 'webm'}); // ['vp8', 'vp9']
```

```

Get available video codecs for a container
tsx

import {getAvailableVideoCodecs} from '@remotion/webcodecs';

getAvailableVideoCodecs({container: 'webm'}); // ['vp8', 'vp9']
```

## See also [​](\#see-also "Direct link to See also")

- [Track Transformation](/docs/webcodecs/track-transformation)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/get-available-video-codecs.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/get-available-video-codecs.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAvailableAudioCodecs()](/docs/webcodecs/get-available-audio-codecs) [Next\
\
@remotion/captions](/docs/captions/)

- [See also](#see-also)
