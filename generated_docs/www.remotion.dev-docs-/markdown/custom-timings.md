---
title: Custom timings
source: Remotion Documentation
last_updated: 2024-12-22
---

# Custom timings

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Timings](/docs/transitions/timings/)
- Custom timings

On this page

# Custom timings

This page describes how to create your own custom timings for [`<TransitionSeries>`](/docs/transitions/transitionseries).

## Concept [​](\#concept "Direct link to Concept")

A timing is a configuration which includes:

[1](#1)

The duration of the transition

[2](#2) A progress function.

## Implementation [​](\#implementation "Direct link to Implementation")

To implement a custom timing function, create a function which returns an object `TransitionTiming`.

```

custom-timing.ts
tsx

import type { TransitionTiming } from "@remotion/transitions";

const customTiming = (): TransitionTiming => {
  return {
    getDurationInFrames: ({ fps }) => fps, // 1 second
    getProgress: ({ frame, fps }) => Math.min(1, frame / fps),
  };
};
```

```

custom-timing.ts
tsx

import type { TransitionTiming } from "@remotion/transitions";

const customTiming = (): TransitionTiming => {
  return {
    getDurationInFrames: ({ fps }) => fps, // 1 second
    getProgress: ({ frame, fps }) => Math.min(1, frame / fps),
  };
};
```

In this example, the timing function will last for 1 second and will be linear.

## Advanced example [​](\#advanced-example "Direct link to Advanced example")

A

In the following example:

[1](#1)

a spring animation will push the progress to 50%

[2](#2) a pause with customizable duration follows

[3](#3) a second spring animation will push the progress to 100%

```

spring-timing-with-pause.ts
tsx

import type { TransitionTiming } from "@remotion/transitions";
import { measureSpring, spring, SpringConfig } from "remotion";

const springTimingWithPause = ({
  pauseDuration,
}: {
  pauseDuration: number;
}): TransitionTiming => {
  const firstHalf: Partial<SpringConfig> = {};
  const secondPush: Partial<SpringConfig> = {
    damping: 200,
  };

  return {
    getDurationInFrames: ({ fps }) => {
      return (
        measureSpring({ fps, config: firstHalf }) +
        measureSpring({ fps, config: secondPush }) +
        pauseDuration
      );
    },
    getProgress({ fps, frame }) {
      const first = spring({ fps, frame, config: firstHalf });
      const second = spring({
        fps,
        frame,
        config: secondPush,
        delay: pauseDuration + measureSpring({ fps, config: firstHalf }),
      });

      return first / 2 + second / 2;
    },
  };
};
```

```

spring-timing-with-pause.ts
tsx

import type { TransitionTiming } from "@remotion/transitions";
import { measureSpring, spring, SpringConfig } from "remotion";

const springTimingWithPause = ({
  pauseDuration,
}: {
  pauseDuration: number;
}): TransitionTiming => {
  const firstHalf: Partial<SpringConfig> = {};
  const secondPush: Partial<SpringConfig> = {
    damping: 200,
  };

  return {
    getDurationInFrames: ({ fps }) => {
      return (
        measureSpring({ fps, config: firstHalf }) +
        measureSpring({ fps, config: secondPush }) +
        pauseDuration
      );
    },
    getProgress({ fps, frame }) {
      const first = spring({ fps, frame, config: firstHalf });
      const second = spring({
        fps,
        frame,
        config: secondPush,
        delay: pauseDuration + measureSpring({ fps, config: firstHalf }),
      });

      return first / 2 + second / 2;
    },
  };
};
```

The duration needs to be calculated deterministically by adding the duration of the two springs and the pause duration, so that the previous and next scenes are timed correctly.

A `spring()` animation by default goes from 0 to 1, which is why they [can be added together](/docs/animation-math).

## Using a custom timing [​](\#using-a-custom-timing "Direct link to Using a custom timing")

You may use a custom timing like any other timing by calling it and passing it to the `timing` prop of `<TransitionSeries.Transition>`.

```

using-custom-timing.tsx
tsx

import { TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";
import { useVideoConfig } from "remotion";

export const CustomTransition: React.FC = () => {
  const { width, height } = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={80}>
        <Letter color="orange">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide({ direction: "from-left" })}
        timing={customTiming({ pauseDuration: 10 })}
      />
      <TransitionSeries.Sequence durationInFrames={200}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

```

using-custom-timing.tsx
tsx

import { TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";
import { useVideoConfig } from "remotion";

export const CustomTransition: React.FC = () => {
  const { width, height } = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={80}>
        <Letter color="orange">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide({ direction: "from-left" })}
        timing={customTiming({ pauseDuration: 10 })}
      />
      <TransitionSeries.Sequence durationInFrames={200}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

## Getting the duration of a timing [​](\#getting-the-duration-of-a-timing "Direct link to Getting the duration of a timing")

Call `.getDurationInFrames({fps})` on any timing function and pass [`fps`](/docs/use-video-config) to get the duration in frames.

## References [​](\#references "Direct link to References")

See the source code for already implemented timings [here](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/timings).

## See also [​](\#see-also "Direct link to See also")

- [Custom presentations](/docs/transitions/presentations/custom)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/timings/custom.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
linearTiming()](/docs/transitions/timings/lineartiming) [Next\
\
Presentations](/docs/transitions/presentations/)

- [Concept](#concept)
- [Implementation](#implementation)
- [Advanced example](#advanced-example)
- [Using a custom timing](#using-a-custom-timing)
- [Getting the duration of a timing](#getting-the-duration-of-a-timing)
- [References](#references)
- [See also](#see-also)
