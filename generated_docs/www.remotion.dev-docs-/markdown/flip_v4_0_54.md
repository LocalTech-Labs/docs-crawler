---
title: flip()v4.0.54
source: Remotion Documentation
last_updated: 2024-12-22
---

# flip()v4.0.54

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- flip()

On this page

# flip() [v4.0.54](https://github.com/remotion-dev/remotion/releases/v4.0.54)

A presentation where the exiting slide flips by 180 degrees, revealing the next slide on the back side.

A

B

direction

from-leftfrom-bottomfrom-rightfrom-top

## Example [​](\#example "Direct link to Example")

```

SlideTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { flip } from "@remotion/transitions/flip";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={flip()}
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
import { flip } from "@remotion/transitions/flip";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={flip()}
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

import { FlipDirection } from "@remotion/transitions/flip";

const flipDirection: FlipDirection = "from-left";
```

```

TypeScript type
tsx

import { FlipDirection } from "@remotion/transitions/flip";

const flipDirection: FlipDirection = "from-left";
```

### `perspective?` [​](\#perspective "Direct link to perspective")

The CSS `perspective` of the flip animation. Defaults to `1000`.

### `outerEnterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#outerenterstyle "Direct link to outerenterstyle")

The style of the outer element when the scene is is entering.

### `outerExitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#outerexitstyle "Direct link to outerexitstyle")

The style of the outer element when the scene is exiting.

### `innerEnterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#innerenterstyle "Direct link to innerenterstyle")

The style of the inner element when the scene is entering.

### `innerExitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#innerexitstyle "Direct link to innerexitstyle")

The style of the inner element when the scene is exiting.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this presentation](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/presentations/flip.tsx)
- [Presentations](/docs/transitions/presentations)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/flip.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
wipe()](/docs/transitions/presentations/wipe) [Next\
\
clockWipe()](/docs/transitions/presentations/clock-wipe)

- [Example](#example)
- [API](#api)
  - [`direction`](#direction)
  - [`perspective?`](#perspective)
  - [`outerEnterStyle?`](#outerenterstyle)
  - [`outerExitStyle?`](#outerexitstyle)
  - [`innerEnterStyle?`](#innerenterstyle)
  - [`innerExitStyle?`](#innerexitstyle)
- [See also](#see-also)
