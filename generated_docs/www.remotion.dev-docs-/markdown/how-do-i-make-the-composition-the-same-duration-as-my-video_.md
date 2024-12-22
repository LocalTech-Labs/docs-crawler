---
title: How do I make the composition the same duration as my video?
source: Remotion Documentation
last_updated: 2024-12-22
---

# How do I make the composition the same duration as my video?

- [Home page](/)
- Embedding videos
- Make timeline duration the same

On this page

# How do I make the composition the same duration as my video?

If you have a component rendering a video:

```

MyComp.tsx
tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} />;
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile('video.mp4')} />;
};
```

and you want to make the composition the same duration as the video, first make the video source a React prop:

```

MyComp.tsx
tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

type MyCompProps = {
  src: string;
};

export const MyComp: React.FC<MyCompProps> = ({src}) => {
  return <OffthreadVideo src={src} />;
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {OffthreadVideo, staticFile} from 'remotion';

type MyCompProps = {
  src: string;
};

export const MyComp: React.FC<MyCompProps> = ({src}) => {
  return <OffthreadVideo src={src} />;
};
```

Then, define a [`calculateMetadata()`](/docs/calculate-metadata) function that calculates the duration of the composition based on the video.

[Install `@remotion/media-parser`](/docs/media-parser) if necessary.

```

MyComp.tsx
tsx

import {CalculateMetadataFunction} from 'remotion';
import {parseMedia} from '@remotion/media-parser';

export const calculateMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({props}) => {
  const {durationInSeconds, dimensions} = await parseMedia({
    src: props.src,
    fields: {
      durationInSeconds: true,
      dimensions: true,
    },
  });

  if (durationInSeconds === null) {
    throw new Error('Could not get video duration');
  }

  const fps = 30;

  return {
    durationInFrames: Math.floor(durationInSeconds * fps),
    fps,
    width: dimensions.width,
    height: dimensions.height,
  };
};
```

```

MyComp.tsx
tsx

import {CalculateMetadataFunction} from 'remotion';
import {parseMedia} from '@remotion/media-parser';

export const calculateMetadata: CalculateMetadataFunction<
  MyCompProps
> = async ({props}) => {
  const {durationInSeconds, dimensions} = await parseMedia({
    src: props.src,
    fields: {
      durationInSeconds: true,
      dimensions: true,
    },
  });

  if (durationInSeconds === null) {
    throw new Error('Could not get video duration');
  }

  const fps = 30;

  return {
    durationInFrames: Math.floor(durationInSeconds * fps),
    fps,
    width: dimensions.width,
    height: dimensions.height,
  };
};
```

note

If your asset is not CORS-enabled, you can use the [`getVideoMetadata`](/docs/get-video-metadata) function from [`@remotion/media-utils`](/docs/media-utils) instead of `parseMedia()`.

Finally, pass the `calculateMetadata` function to the `Composition` component and define the previously hardcoded `src` as a default prop:

```

Root.tsx
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComp, calculateMetadata} from './MyComp';

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      }}
      calculateMetadata={calculateMetadata}
    />
  );
};
```

```

Root.tsx
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComp, calculateMetadata} from './MyComp';

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      }}
      calculateMetadata={calculateMetadata}
    />
  );
};
```

## How do I make the composition the same duration as my audio? [​](\#how-do-i-make-the-composition-the-same-duration-as-my-audio "Direct link to How do I make the composition the same duration as my audio?")

Follow the same steps, but instead of [`parseMedia()`](/docs/media-parser), use [`getAudioDurationInSeconds()`](/docs/get-audio-duration-in-seconds) from `@remotion/media-utils` to calculate the duration of the composition based on the audio file.

## See Also [​](\#see-also "Direct link to See Also")

- [Variable duration](/docs/dynamic-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/align-duration.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Adding a video](/docs/videos/) [Next\
\
Playing videos in sequence](/docs/videos/sequence)

- [How do I make the composition the same duration as my audio?](#how-do-i-make-the-composition-the-same-duration-as-my-audio)
- [See Also](#see-also)
