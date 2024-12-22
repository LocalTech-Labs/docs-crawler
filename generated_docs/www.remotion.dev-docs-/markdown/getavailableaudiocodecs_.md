---
title: getAvailableAudioCodecs()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAvailableAudioCodecs()

- [Home page](/)
- [@remotion/webcodecs](/docs/webcodecs/)
- getAvailableAudioCodecs()

On this page

# getAvailableAudioCodecs()

Given a container, get a list of audio codecs that the container can hold.

This does not mean that a any audio stream of this codec can be put into the container.

Use [`canReencodeAudioTrack()`](/docs/webcodecs/can-reencode-audio-track) and [`canCopyAudioTrack()`](/docs/webcodecs/can-copy-audio-track) to determine this.

```

Get available audio codecs for a container
tsx

import {getAvailableAudioCodecs} from '@remotion/webcodecs';

getAvailableAudioCodecs({container: 'webm'}); // ['opus']
```

```

Get available audio codecs for a container
tsx

import {getAvailableAudioCodecs} from '@remotion/webcodecs';

getAvailableAudioCodecs({container: 'webm'}); // ['opus']
```

## See also [​](\#see-also "Direct link to See also")

- [Track Transformation](/docs/webcodecs/track-transformation)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/webcodecs/src/get-available-video-codecs.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/get-available-audio-codecs.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getDefaultVideoCodec()](/docs/webcodecs/get-default-video-codec) [Next\
\
getAvailableVideoCodecs()](/docs/webcodecs/get-available-video-codecs)

- [See also](#see-also)
