---
title: clockWipe()v4.0.74
source: Remotion Documentation
last_updated: 2024-12-22
---

# clockWipe()v4.0.74

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- clockWipe()

On this page

# clockWipe() [v4.0.74](https://github.com/remotion-dev/remotion/releases/v4.0.74)

A presentation where the exiting slide is wiped out in a circular movement, revealing the next slide underneath it.

A

B

## Example [​](\#example "Direct link to Example")

```

ClockWipeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { clockWipe } from "@remotion/transitions/clock-wipe";
import { useVideoConfig } from "remotion";

const BasicTransition = () => {
  const { width, height } = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={clockWipe({ width, height })}
        timing={linearTiming({ durationInFrames: 30 })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

```

ClockWipeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { clockWipe } from "@remotion/transitions/clock-wipe";
import { useVideoConfig } from "remotion";

const BasicTransition = () => {
  const { width, height } = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={clockWipe({ width, height })}
        timing={linearTiming({ durationInFrames: 30 })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

## API [​](\#api "Direct link to API")

Accepts an object with the following options:

### `width` [​](\#width "Direct link to width")

Should be set to the width of the video.

### `height` [​](\#height "Direct link to height")

Should be set to the height of the video.

### `outerEnterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#outerenterstyle "Direct link to outerenterstyle")

The style of the outer element when the scene is is entering.

### `outerExitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#outerexitstyle "Direct link to outerexitstyle")

The style of the outer element when the scene is exiting.

### `innerEnterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#innerenterstyle "Direct link to innerenterstyle")

The style of the inner element when the scene is entering.

### `innerExitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#innerexitstyle "Direct link to innerexitstyle")

The style of the inner element when the scene is exiting.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this presentation](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/presentations/clock-wipe.tsx)
- [Presentations](/docs/transitions/presentations)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/clock-wipe.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
flip()](/docs/transitions/presentations/flip) [Next\
\
none()](/docs/transitions/presentations/none)

- [Example](#example)
- [API](#api)
  - [`width`](#width)
  - [`height`](#height)
  - [`outerEnterStyle?`](#outerenterstyle)
  - [`outerExitStyle?`](#outerexitstyle)
  - [`innerEnterStyle?`](#innerenterstyle)
  - [`innerExitStyle?`](#innerexitstyle)
- [See also](#see-also)
