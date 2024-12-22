---
title: Change the speed of a video over time
source: Remotion Documentation
last_updated: 2024-12-22
---

# Change the speed of a video over time

- [Home page](/)
- Embedding videos
- Changing speed over time

On this page

# Change the speed of a video over time

To speed up a video over time - for example to start with regular speed and then increasingly make it faster - we need to write some code to ensure we stay within Remotion's rules.

It is not as easy as interpolating the [`playbackRate`](/docs/video#playbackrate):

```

❌ Does not work
tsx

<OffthreadVideo
  playbackRate={interpolate(frame, [0, 100], [1, 5])}
  src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4#disable"
/>;
```

```

❌ Does not work
tsx

<OffthreadVideo
  playbackRate={interpolate(frame, [0, 100], [1, 5])}
  src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4#disable"
/>;
```

This is because Remotion will evaluate each frame independently from the other frames. If `frame` is 100, the `playbackRate` evaluates as 5 and Remotion will render the 500th frame of the video, which is undesired because it does not take into account that the speed has been building up to 5 until now.

The correct way is to calculate the total time that has elapsed until the current frame, by accumulating the playback rates of the previous frames. Then we start the video from that frame and set the playback rate of that current frame to ensure the video plays smoothly.

```

✅ AcceleratedVideo.tsx
tsx

import React from 'react';
import {interpolate, Sequence, useCurrentFrame, OffthreadVideo} from 'remotion';

const remapSpeed = (frame: number, speed: (fr: number) => number) => {
  let framesPassed = 0;
  for (let i = 0; i <= frame; i++) {
    framesPassed += speed(i);
  }

  return framesPassed;
};

export const AcceleratedVideo: React.FC = () => {
  const frame = useCurrentFrame();

  const speedFunction = (f: number) => interpolate(f, [0, 500], [1, 5]);

  const remappedFrame = remapSpeed(frame, speedFunction);

  return (
    <Sequence from={frame}>
      <OffthreadVideo
        startFrom={Math.round(remappedFrame)}
        playbackRate={speedFunction(frame)}
        src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4#disable"
      />
    </Sequence>
  );
};
```

```

✅ AcceleratedVideo.tsx
tsx

import React from 'react';
import {interpolate, Sequence, useCurrentFrame, OffthreadVideo} from 'remotion';

const remapSpeed = (frame: number, speed: (fr: number) => number) => {
  let framesPassed = 0;
  for (let i = 0; i <= frame; i++) {
    framesPassed += speed(i);
  }

  return framesPassed;
};

export const AcceleratedVideo: React.FC = () => {
  const frame = useCurrentFrame();

  const speedFunction = (f: number) => interpolate(f, [0, 500], [1, 5]);

  const remappedFrame = remapSpeed(frame, speedFunction);

  return (
    <Sequence from={frame}>
      <OffthreadVideo
        startFrom={Math.round(remappedFrame)}
        playbackRate={speedFunction(frame)}
        src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4#disable"
      />
    </Sequence>
  );
};
```

note

We are [disabling the media fragment hints](/docs/media-fragments) being appended to the video URL by adding `#disable` to the end of the URL.

Note that the timeline in the Remotion Studio might move as you play the video because we reposition the sequence over time. This is okay as long as we calculate the frame in an idempotent way.

## Demo [​](\#demo "Direct link to Demo")

0:00 / 0:50

## See also [​](\#see-also "Direct link to See also")

- [`<OffthreadVideo>`](/docs/video)
- [Different segments at different speeds](/docs/miscellaneous/snippets/different-segments-at-different-speeds)
- [Jump Cutting](/docs/miscellaneous/snippets/jumpcuts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/accelerated-video.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Manipulating pixels](/docs/video-manipulation) [Next\
\
Jump Cuts](/docs/miscellaneous/snippets/jumpcuts)

- [Demo](#demo)
- [See also](#see-also)
