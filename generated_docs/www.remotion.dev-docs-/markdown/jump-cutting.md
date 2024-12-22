---
title: Jump Cutting
source: Remotion Documentation
last_updated: 2024-12-22
---

# Jump Cutting

- [Home page](/)
- Embedding videos
- Jump Cuts

On this page

# Jump Cutting

You might wanna use a "jump cut" to skip parts of a video.

Use the following snippet to skip certain sections of a video, without re-mounting it.

```

tsx

import React, {useMemo} from 'react';
import {
  CalculateMetadataFunction,
  OffthreadVideo,
  staticFile,
  useCurrentFrame,
} from 'remotion';

const fps = 30;

type Section = {
  startFrom: number;
  endAt: number;
};

export const SAMPLE_SECTIONS: Section[] = [
  {startFrom: 0, endAt: 5 * fps},
  {
    startFrom: 7 * fps,
    endAt: 10 * fps,
  },
  {
    startFrom: 13 * fps,
    endAt: 18 * fps,
  },
];

type Props = {
  sections: Section[];
};

export const calculateMetadata: CalculateMetadataFunction<Props> = ({
  props,
}) => {
  const durationInFrames = props.sections.reduce((acc, section) => {
    return acc + section.endAt - section.startFrom;
  }, 0);

  return {
    fps,
    durationInFrames,
  };
};

export const JumpCuts: React.FC<Props> = ({sections}) => {
  const frame = useCurrentFrame();

  const startFrom = useMemo(() => {
    let summedUpDurations = 0;
    for (const section of sections) {
      summedUpDurations += section.endAt - section.startFrom;
      if (summedUpDurations > frame) {
        return section.endAt - summedUpDurations;
      }
    }

    return null;
  }, [frame, sections]);

  if (startFrom === null) {
    return null;
  }

  return (
    <OffthreadVideo
      pauseWhenBuffering
      startFrom={startFrom}
      // Remotion will automatically add a time fragment to the end of the video URL
      // based on `startFrom` and `endAt`. Opt out of this by adding one yourself.
      // https://www.remotion.dev/docs/media-fragments
      src={`${staticFile('time.mp4')}#t=0,`}
    />
  );
};
```

```

tsx

import React, {useMemo} from 'react';
import {
  CalculateMetadataFunction,
  OffthreadVideo,
  staticFile,
  useCurrentFrame,
} from 'remotion';

const fps = 30;

type Section = {
  startFrom: number;
  endAt: number;
};

export const SAMPLE_SECTIONS: Section[] = [
  {startFrom: 0, endAt: 5 * fps},
  {
    startFrom: 7 * fps,
    endAt: 10 * fps,
  },
  {
    startFrom: 13 * fps,
    endAt: 18 * fps,
  },
];

type Props = {
  sections: Section[];
};

export const calculateMetadata: CalculateMetadataFunction<Props> = ({
  props,
}) => {
  const durationInFrames = props.sections.reduce((acc, section) => {
    return acc + section.endAt - section.startFrom;
  }, 0);

  return {
    fps,
    durationInFrames,
  };
};

export const JumpCuts: React.FC<Props> = ({sections}) => {
  const frame = useCurrentFrame();

  const startFrom = useMemo(() => {
    let summedUpDurations = 0;
    for (const section of sections) {
      summedUpDurations += section.endAt - section.startFrom;
      if (summedUpDurations > frame) {
        return section.endAt - summedUpDurations;
      }
    }

    return null;
  }, [frame, sections]);

  if (startFrom === null) {
    return null;
  }

  return (
    <OffthreadVideo
      pauseWhenBuffering
      startFrom={startFrom}
      // Remotion will automatically add a time fragment to the end of the video URL
      // based on `startFrom` and `endAt`. Opt out of this by adding one yourself.
      // https://www.remotion.dev/docs/media-fragments
      src={`${staticFile('time.mp4')}#t=0,`}
    />
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Different segments at different speeds](/docs/miscellaneous/snippets/different-segments-at-different-speeds)
- [Change the speed of a video over time](/docs/miscellaneous/snippets/accelerated-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/jumpcuts.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Changing speed over time](/docs/miscellaneous/snippets/accelerated-video) [Next\
\
Freeze a portion](/docs/miscellaneous/snippets/freeze-portions)

- [See also](#see-also)
