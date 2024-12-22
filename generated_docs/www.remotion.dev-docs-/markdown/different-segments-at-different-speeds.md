---
title: Different segments at different speeds
source: Remotion Documentation
last_updated: 2024-12-22
---

# Different segments at different speeds

- [Home page](/)
- Snippets
- Different segments at different speeds

On this page

# Different segments at different speeds

If you have a video and want to show different sections of the video at different speeds, use the following snippet.

```

SegmentSpeeds.tsx
tsx

import {OffthreadVideo, Sequence, staticFile, useCurrentFrame} from 'remotion';

const segments = [
  {
    duration: 100,
    speed: 0.5,
  },
  {
    duration: 100,
    speed: 1,
  },
  {
    duration: 200,
    speed: 2,
  },
  {
    duration: 400,
    speed: 4,
  },
];

type AccumulatedSegment = {
  start: number;
  passedVideoTime: number;
  end: number;
  speed: number;
};

export const accumulateSegments = () => {
  const accumulatedSegments: AccumulatedSegment[] = [];
  let accumulatedDuration = 0;
  let accumulatedPassedVideoTime = 0;

  for (const segment of segments) {
    const duration = segment.duration / segment.speed;
    accumulatedSegments.push({
      end: accumulatedDuration + duration,
      speed: segment.speed,
      start: accumulatedDuration,
      passedVideoTime: accumulatedPassedVideoTime,
    });

    accumulatedPassedVideoTime += segment.duration;
    accumulatedDuration += duration;
  }

  return accumulatedSegments;
};

export const SpeedSegments = () => {
  const frame = useCurrentFrame();
  const accumulated = accumulateSegments();

  const currentSegment = accumulated.find(
    (segment) => frame > segment.start && frame <= segment.end,
  );

  if (!currentSegment) {
    return;
  }

  return (
    <Sequence from={currentSegment.start}>
      <OffthreadVideo
        pauseWhenBuffering
        startFrom={currentSegment.passedVideoTime}
        // Remotion will automatically add a time fragment to the end of the video URL
        // based on `startFrom`. Opt out of this by adding one yourself.
        // https://www.remotion.dev/docs/media-fragments
        src={`${staticFile('bigbuckbunny.mp4')}#t=0,`}
        playbackRate={currentSegment.speed}
      />
    </Sequence>
  );
};
```

```

SegmentSpeeds.tsx
tsx

import {OffthreadVideo, Sequence, staticFile, useCurrentFrame} from 'remotion';

const segments = [
  {
    duration: 100,
    speed: 0.5,
  },
  {
    duration: 100,
    speed: 1,
  },
  {
    duration: 200,
    speed: 2,
  },
  {
    duration: 400,
    speed: 4,
  },
];

type AccumulatedSegment = {
  start: number;
  passedVideoTime: number;
  end: number;
  speed: number;
};

export const accumulateSegments = () => {
  const accumulatedSegments: AccumulatedSegment[] = [];
  let accumulatedDuration = 0;
  let accumulatedPassedVideoTime = 0;

  for (const segment of segments) {
    const duration = segment.duration / segment.speed;
    accumulatedSegments.push({
      end: accumulatedDuration + duration,
      speed: segment.speed,
      start: accumulatedDuration,
      passedVideoTime: accumulatedPassedVideoTime,
    });

    accumulatedPassedVideoTime += segment.duration;
    accumulatedDuration += duration;
  }

  return accumulatedSegments;
};

export const SpeedSegments = () => {
  const frame = useCurrentFrame();
  const accumulated = accumulateSegments();

  const currentSegment = accumulated.find(
    (segment) => frame > segment.start && frame <= segment.end,
  );

  if (!currentSegment) {
    return;
  }

  return (
    <Sequence from={currentSegment.start}>
      <OffthreadVideo
        pauseWhenBuffering
        startFrom={currentSegment.passedVideoTime}
        // Remotion will automatically add a time fragment to the end of the video URL
        // based on `startFrom`. Opt out of this by adding one yourself.
        // https://www.remotion.dev/docs/media-fragments
        src={`${staticFile('bigbuckbunny.mp4')}#t=0,`}
        playbackRate={currentSegment.speed}
      />
    </Sequence>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Change the speed of a video over time](/docs/miscellaneous/snippets/accelerated-video)
- [Jump Cutting](/docs/miscellaneous/snippets/jumpcuts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/different-segments-at-different-speeds.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Stuck renders](/docs/troubleshooting/stuck-render) [Next\
\
Embedding a <Player> into an <iframe>](/docs/miscellaneous/snippets/player-in-iframe)

- [See also](#see-also)