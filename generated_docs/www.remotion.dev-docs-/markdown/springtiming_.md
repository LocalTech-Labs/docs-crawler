---
title: springTiming()
source: Remotion Documentation
last_updated: 2024-12-22
---

# springTiming()

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Timings](/docs/transitions/timings/)
- springTiming()

On this page

# springTiming()

A timing function for [`<TransitionSeries>`](/docs/transitions/transitionseries) based on [`spring()`](/docs/spring).

```

SlideTransition.tsx
tsx

import { springTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
        timing={springTiming({
          config: {
            damping: 200,
          },
          durationInFrames: 30,
          durationRestThreshold: 0.001,
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

import { springTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
        timing={springTiming({
          config: {
            damping: 200,
          },
          durationInFrames: 30,
          durationRestThreshold: 0.001,
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

### `config` [​](\#config "Direct link to config")

_optional_

A spring `config`, see [`spring()`](/docs/spring).

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

_optional_

Stretch the timing curve of the animation to a fixed duration.

### `durationRestThreshold?` [​](\#durationrestthreshold "Direct link to durationrestthreshold")

_optional_

At which point the animation is considered to be finished. Defaults: `0.005`

### `reverse?` [v4.0.144](https://github.com/remotion-dev/remotion/releases/v4.0.144) [​](\#reverse "Direct link to reverse")

_optional_

Reverse the timing curve.

## Recommendation: Set a low duration rest threshold [​](\#recommendation-set-a-low-duration-rest-threshold "Direct link to Recommendation: Set a low duration rest threshold")

The default `durationRestThreshold` is `0.005` (same as `spring()`). This means that if the animations has progressed 99.5%, it is considered to be finished.

For a transition, this can lead to a slightly noticeable cutoff of the animation:

A

B

To avoid this, set a low `durationRestThreshold`, we recommend `0.001`.

Consider that this will also increase the duration of the animation.
If you set a fixed `durationInFrames`, the animation will feel faster because it for longer it is not considered to be finished.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/timings/springtiming.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Timings](/docs/transitions/timings/) [Next\
\
linearTiming()](/docs/transitions/timings/lineartiming)

- [API](#api)
  - [`config`](#config)
  - [`durationInFrames`](#durationinframes)
  - [`durationRestThreshold?`](#durationrestthreshold)
  - [`reverse?`](#reverse)
- [Recommendation: Set a low duration rest threshold](#recommendation-set-a-low-duration-rest-threshold)
