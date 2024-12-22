---
title: linearTiming()
source: Remotion Documentation
last_updated: 2024-12-22
---

# linearTiming()

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Timings](/docs/transitions/timings/)
- linearTiming()

On this page

# linearTiming()

A timing function for [`<TransitionSeries>`](/docs/transitions/transitionseries) based on [`interpolate()`](/docs/interpolate).

```

SlideTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";
import { Easing } from "remotion";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
        timing={linearTiming({
          durationInFrames: 30,
          easing: Easing.in(Easing.ease),
        })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

```

SlideTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";
import { Easing } from "remotion";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
        timing={linearTiming({
          durationInFrames: 30,
          easing: Easing.in(Easing.ease),
        })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

## API [​](\#api "Direct link to API")

An object with the following properties:

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

The duration of the transition in frames.

### `easing?` [​](\#easing "Direct link to easing")

_optional_

An easing function, see [`Easing`](/docs/easing).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this presentation](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/timings/linear-timing.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/timings/lineartiming.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
springTiming()](/docs/transitions/timings/springtiming) [Next\
\
Custom timings](/docs/transitions/timings/custom)

- [API](#api)
  - [`durationInFrames`](#durationinframes)
  - [`easing?`](#easing)
- [See also](#see-also)
