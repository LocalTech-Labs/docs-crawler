---
title: Embedding transparent videos
source: Remotion Documentation
last_updated: 2024-12-22
---

# Embedding transparent videos

- [Home page](/)
- Embedding videos
- Adding a transparent video

On this page

# Embedding transparent videos

You can embed transparent videos in Remotion.

## With an alpha channel [​](\#with-an-alpha-channel "Direct link to With an alpha channel")

To embed a video which has an alpha channel, use the [`<OffthreadVideo>`](/docs/offthreadvideo) component with the [`transparent`](/docs/offthreadvideo#transparent) prop.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('transparent.webm')} transparent />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('transparent.webm')} transparent />;
};
```

## Without an alpha channel [​](\#without-an-alpha-channel "Direct link to Without an alpha channel")

To embed a video which does not have an alpha channel but just a black background, add `mixBlendMode: "screen"` to the [`style`](/docs/offthreadvideo#style) prop of the [`<OffthreadVideo>`](/docs/offthreadvideo) component.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo
      src={staticFile('nottransparent.mp4')}
      style={{
        mixBlendMode: 'screen',
      }}
    />
  );
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo
      src={staticFile('nottransparent.mp4')}
      style={{
        mixBlendMode: 'screen',
      }}
    />
  );
};
```

## Greenscreen effect [​](\#greenscreen-effect "Direct link to Greenscreen effect")

To remove a background based on the color of the pixels, see the [Greenscreen example](/docs/video-manipulation#greenscreen-example).

## Rendering transparent videos [​](\#rendering-transparent-videos "Direct link to Rendering transparent videos")

To export a video with transparency, see [Transparent videos](/docs/transparent-videos/)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/videos/transparency.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Playing videos in sequence](/docs/videos/sequence) [Next\
\
Manipulating pixels](/docs/video-manipulation)

- [With an alpha channel](#with-an-alpha-channel)
- [Without an alpha channel](#without-an-alpha-channel)
- [Greenscreen effect](#greenscreen-effect)
- [Rendering transparent videos](#rendering-transparent-videos)
