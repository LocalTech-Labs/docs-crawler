---
title: Embedding videos into Remotion
source: Remotion Documentation
last_updated: 2024-12-22
---

# Embedding videos into Remotion

- [Home page](/)
- Embedding videos
- Adding a video

On this page

# Embedding videos into Remotion

You can embed existing videos into Remotion by using the [`<OffthreadVideo>`](/docs/offthreadvideo) component.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

## Using a local file [​](\#using-a-local-file "Direct link to Using a local file")

Put a file into the [`public` folder](/docs/terminology/public-dir) and reference it using [`staticFile()`](/docs/staticfile).

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} />;
};
```

## Trimming [​](\#trimming "Direct link to Trimming")

By using the [`startFrom`](/docs/offthreadvideo#startfrom) prop, you can remove the first few seconds of the video.

In the example below, the first two seconds of the video are skipped (assuming a composition FPS of 30).

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} />;
};
```

Similarly, you can use [`endAt`](/docs/offthreadvideo#endat) to trim the end of the video.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} endAt={120} />
  );
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src={staticFile('video.mp4')} startFrom={60} endAt={120} />
  );
};
```

## Delaying [​](\#delaying "Direct link to Delaying")

Use the [`<Sequence>`](/docs/sequence) component to delay the appearance of a video.

In the example below, the video will start playing at frame 60.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile, Sequence} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <Sequence from={60}>
      <OffthreadVideo src={staticFile('video.mp4')} />
    </Sequence>
  );
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile, Sequence} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <Sequence from={60}>
      <OffthreadVideo src={staticFile('video.mp4')} />
    </Sequence>
  );
};
```

## Size and Position [​](\#size-and-position "Direct link to Size and Position")

You can size and position the video using CSS.

You'll find the properties `width`, `height`, `position`, `left`, `top`, `right`, `bottom`, `margin`, `aspectRatio` and `objectFit` useful.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo
      src={staticFile('video.mp4')}
      style={{
        width: 640,
        height: 360,
        position: 'absolute',
        top: 100,
        left: 100,
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
      src={staticFile('video.mp4')}
      style={{
        width: 640,
        height: 360,
        position: 'absolute',
        top: 100,
        left: 100,
      }}
    />
  );
};
```

## Volume [​](\#volume "Direct link to Volume")

You can set the volume of the video using the [`volume`](/docs/offthreadvideo#volume) prop.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} volume={0.5} />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} volume={0.5} />;
};
```

You can also mute a video using the [`muted`](/docs/offthreadvideo#muted) prop.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} muted />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} muted />;
};
```

See [Using Audio](/docs/using-audio#controlling-volume) for more ways you can control volume.

## Speed [​](\#speed "Direct link to Speed")

You can use the [`playbackRate`](/docs/offthreadvideo#playbackrate) prop to change the speed of the video.

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} playbackRate={2} />;
};
```

```

tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} playbackRate={2} />;
};
```

This only works if the speed is constant. See also: [Changing the speed of a video over time](/docs/miscellaneous/snippets/accelerated-video).

## Looping [​](\#looping "Direct link to Looping")

See: [Looping an `<OffthreadVideo>`](/docs/offthreadvideo#looping-a-video)

## GIFs [​](\#gifs "Direct link to GIFs")

See: [Using GIFs](/docs/gif)

## See also [​](\#see-also "Direct link to See also")

- [`<OffthreadVideo>`](/docs/offthreadvideo)
- [Using Audio](/docs/using-audio)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/videos/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Animation math](/docs/animation-math) [Next\
\
Make timeline duration the same](/docs/miscellaneous/snippets/align-duration)

- [Using a local file](#using-a-local-file)
- [Trimming](#trimming)
- [Delaying](#delaying)
- [Size and Position](#size-and-position)
- [Volume](#volume)
- [Speed](#speed)
- [Looping](#looping)
- [GIFs](#gifs)
- [See also](#see-also)
