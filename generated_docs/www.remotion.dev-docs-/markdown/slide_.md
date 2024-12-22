---
title: slide()
source: Remotion Documentation
last_updated: 2024-12-22
---

# slide()

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- slide()

On this page

# slide()

A presentation where the entering slide pushes out the exiting slide.

A

B

direction

from-leftfrom-bottomfrom-rightfrom-top

## Example [​](\#example "Direct link to Example")

```

SlideTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
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

SlideTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { slide } from "@remotion/transitions/slide";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={slide()}
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

Takes an object with the following properties:

### `direction` [​](\#direction "Direct link to direction")

One of `from-left`, `from-right`, `from-top`, `from-bottom`.

```

TypeScript type
tsx

import { SlideDirection } from "@remotion/transitions/slide";

const slideDirection: SlideDirection = "from-left";
```

```

TypeScript type
tsx

import { SlideDirection } from "@remotion/transitions/slide";

const slideDirection: SlideDirection = "from-left";
```

### `enterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#enterstyle "Direct link to enterstyle")

The style of the container when the scene is is entering.

### `exitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#exitstyle "Direct link to exitstyle")

The style of the container when the scene is exiting.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this presentation](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/presentations/slide.tsx)
- [Presentations](/docs/transitions/presentations)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/slide.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
fade()](/docs/transitions/presentations/fade) [Next\
\
wipe()](/docs/transitions/presentations/wipe)

- [Example](#example)
- [API](#api)
  - [`direction`](#direction)
  - [`enterStyle?`](#enterstyle)
  - [`exitStyle?`](#exitstyle)
- [See also](#see-also)
