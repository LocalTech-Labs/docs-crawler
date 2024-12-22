---
title: Playing videos in sequence
source: Remotion Documentation
last_updated: 2024-12-22
---

# Playing videos in sequence

- [Home page](/)
- Embedding videos
- Playing videos in sequence

On this page

# Playing videos in sequence

If you would like to play multiple videos in sequence, you can:

[1](#1)

Define a component that renders a [`<Series>`](/docs/series) of [`<OffthreadVideo>`](/docs/offthreadvideo) components.

[2](#2) Create a [`calculateMetadata()`](/docs/calculate-metadata) function that fetches the duration of each video.

[3](#3) Register a [`<Composition>`](/docs/composition) that specifies a list of videos.

## Basic example [​](\#basic-example "Direct link to Basic example")

Start off by creating a component that renders a list of videos using the [`<Series>`](/docs/series) and [`<OffthreadVideo>`](/docs/offthreadvideo) component:

```

VideosInSequence.tsx
tsx

import React from 'react';
import {OffthreadVideo, Series} from 'remotion';

type VideoToEmbed = {
  src: string;
  durationInFrames: number | null;
};

type Props = {
  videos: VideoToEmbed[];
};

export const VideosInSequence: React.FC<Props> = ({videos}) => {
  return (
    <Series>
      {videos.map((vid) => {
        if (vid.durationInFrames === null) {
          throw new Error('Could not get video duration');
        }

        return (
          <Series.Sequence
            key={vid.src}
            durationInFrames={vid.durationInFrames}
          >
            <OffthreadVideo src={vid.src} />
          </Series.Sequence>
        );
      })}
    </Series>
  );
};
```

```

VideosInSequence.tsx
tsx

import React from 'react';
import {OffthreadVideo, Series} from 'remotion';

type VideoToEmbed = {
  src: string;
  durationInFrames: number | null;
};

type Props = {
  videos: VideoToEmbed[];
};

export const VideosInSequence: React.FC<Props> = ({videos}) => {
  return (
    <Series>
      {videos.map((vid) => {
        if (vid.durationInFrames === null) {
          throw new Error('Could not get video duration');
        }

        return (
          <Series.Sequence
            key={vid.src}
            durationInFrames={vid.durationInFrames}
          >
            <OffthreadVideo src={vid.src} />
          </Series.Sequence>
        );
      })}
    </Series>
  );
};
```

In the same file, create a function that calculates the metadata for the composition:

[1](#1)

Calls [`parseMedia()`](/docs/media-parser) to get the duration of
each video.

[2](#2)

Create a [`calculateMetadata()`](/docs/calculate-metadata)
function that fetches the duration of each video.

[3](#3)

Sums up all durations to get the total duration of the
composition.

```

VideosInSequence.tsx
tsx

export const calculateMetadata: CalculateMetadataFunction<Props> = async ({
  props,
}) => {
  const fps = 30;
  const videos = await Promise.all([
    ...props.videos.map(async (video): Promise<VideoToEmbed> => {
      const {durationInSeconds} = await parseMedia({
        src: video.src,
        fields: {
          durationInSeconds: true,
        },
      });

      if (durationInSeconds === null) {
        throw new Error(`Could not get video duration of ${durationInSeconds}`);
      }

      return {
        durationInFrames: Math.floor(durationInSeconds * fps),
        src: video.src,
      };
    }),
  ]);

  const totalDurationInFrames = videos.reduce(
    (acc, video) => acc + (video.durationInFrames ?? 0),
    0,
  );

  return {
    props: {
      ...props,
      videos,
    },
    fps,
    durationInFrames: totalDurationInFrames,
  };
};
```

```

VideosInSequence.tsx
tsx

export const calculateMetadata: CalculateMetadataFunction<Props> = async ({
  props,
}) => {
  const fps = 30;
  const videos = await Promise.all([
    ...props.videos.map(async (video): Promise<VideoToEmbed> => {
      const {durationInSeconds} = await parseMedia({
        src: video.src,
        fields: {
          durationInSeconds: true,
        },
      });

      if (durationInSeconds === null) {
        throw new Error(`Could not get video duration of ${durationInSeconds}`);
      }

      return {
        durationInFrames: Math.floor(durationInSeconds * fps),
        src: video.src,
      };
    }),
  ]);

  const totalDurationInFrames = videos.reduce(
    (acc, video) => acc + (video.durationInFrames ?? 0),
    0,
  );

  return {
    props: {
      ...props,
      videos,
    },
    fps,
    durationInFrames: totalDurationInFrames,
  };
};
```

In your [root file](/docs/terminology/root-file), create a [`<Composition>`](/docs/composition) that uses the `VideosInSequence` component and the exported `calculateMetadata` function:

```

Root.tsx
tsx

import React from 'react';
import {Composition, staticFile} from 'remotion';
import {VideosInSequence, calculateMetadata} from './VideosInSequence';

export const Root: React.FC = () => {
  return (
    <Composition
      id="VideosInSequence"
      component={VideosInSequence}
      width={1920}
      height={1080}
      defaultProps={{
        videos: [
          {
            durationInFrames: null,
            src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
          },
          {
            durationInFrames: null,
            src: staticFile('localvideo.mp4'),
          },
        ],
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
import {Composition, staticFile} from 'remotion';
import {VideosInSequence, calculateMetadata} from './VideosInSequence';

export const Root: React.FC = () => {
  return (
    <Composition
      id="VideosInSequence"
      component={VideosInSequence}
      width={1920}
      height={1080}
      defaultProps={{
        videos: [
          {
            durationInFrames: null,
            src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
          },
          {
            durationInFrames: null,
            src: staticFile('localvideo.mp4'),
          },
        ],
      }}
      calculateMetadata={calculateMetadata}
    />
  );
};
```

## Adding premounting [​](\#adding-premounting "Direct link to Adding premounting")

If you only care about the video looking smooth when rendered, you may skip this step.

If you also want smooth preview playback in the Player, consider this:

A video will only load when it is about to be played.

To create a smoother preview playback, we should do two things to all videos:

[1](#1)

Add a [`premountFor`](/docs/series#premountfor) prop to [`<Series.Sequence>`](/docs/series). This will
invisibly mount the video tag before it is played, giving it some time to load.

[2](#2)

Add the [`pauseWhenBuffering`](/docs/series#premount) prop. This will transition the Player into a [buffering state](/docs/player/buffer-state), should the video still need to load.

```

VideosInSequence.tsx
tsx

export const VideosInSequence: React.FC<Props> = ({videos}) => {
  const {fps} = useVideoConfig();

  return (
    <Series>
      {videos.map((vid) => {
        if (vid.durationInFrames === null) {
          throw new Error('Could not get video duration');
        }

        return (
          <Series.Sequence
            key={vid.src}
            premountFor={4 * fps}
            durationInFrames={vid.durationInFrames}
          >
            <OffthreadVideo pauseWhenBuffering src={vid.src} />
          </Series.Sequence>
        );
      })}
    </Series>
  );
};
```

```

VideosInSequence.tsx
tsx

export const VideosInSequence: React.FC<Props> = ({videos}) => {
  const {fps} = useVideoConfig();

  return (
    <Series>
      {videos.map((vid) => {
        if (vid.durationInFrames === null) {
          throw new Error('Could not get video duration');
        }

        return (
          <Series.Sequence
            key={vid.src}
            premountFor={4 * fps}
            durationInFrames={vid.durationInFrames}
          >
            <OffthreadVideo pauseWhenBuffering src={vid.src} />
          </Series.Sequence>
        );
      })}
    </Series>
  );
};
```

## Browser autoplay policies [​](\#browser-autoplay-policies "Direct link to Browser autoplay policies")

Mobile browsers are more aggressive in blocking autoplaying videos that enter after the start of the composition.

If you want to ensure a smooth playback experience for all videos, also [read the notes about browser autoplay behavior](/docs/player/autoplay#media-that-enters-the-video-after-the-start) and customize the behavior if needed.

## See also [​](\#see-also "Direct link to See also")

- [`<OffthreadVideo>`](/docs/offthreadvideo)
- [`<Series>`](/docs/series)
- [`calculateMetadata()`](/docs/calculate-metadata)
- [`parseMedia()`](/docs/media-parser)
- [`<Composition>`](/docs/composition)
- [Root file](/docs/terminology/root-file)
- [Player buffering state](/docs/player/buffer-state)
- [Avoiding flickering in the Player](/docs/troubleshooting/player-flicker)
- [Combatting autoplay issues](/docs/player/autoplay)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/videos/sequence.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Make timeline duration the same](/docs/miscellaneous/snippets/align-duration) [Next\
\
Adding a transparent video](/docs/videos/transparency)

- [Basic example](#basic-example)
- [Adding premounting](#adding-premounting)
- [Browser autoplay policies](#browser-autoplay-policies)
- [See also](#see-also)
