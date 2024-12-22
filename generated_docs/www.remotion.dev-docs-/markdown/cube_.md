---
title: cube()
source: Remotion Documentation
last_updated: 2024-12-22
---

# cube()

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- cube()

On this page

# cube()

note

This is a paid item which you can [buy here](https://remotion.pro/cube-transition).

A presentation where both the entering and exiting scene rotate with 3D perspective.

A

B

direction

from-leftfrom-topfrom-rightfrom-bottom

## Example [​](\#example "Direct link to Example")

```

CubeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { cube } from "@remotion-dev/cube-presentation";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={cube({ direction: "from-left" })}
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

CubeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { cube } from "@remotion-dev/cube-presentation";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={cube({ direction: "from-left" })}
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

The `cube()` function does take an object with the following properties:

### `direction` [​](\#direction "Direct link to direction")

One of `from-left`, `from-right`, `from-top`, `from-bottom`.

```

tsx

import { CubeDirection } from "@remotion-dev/cube-presentation";

const flipDirection: CubeDirection = "from-left";
```

```

tsx

import { CubeDirection } from "@remotion-dev/cube-presentation";

const flipDirection: CubeDirection = "from-left";
```

### `perspective?` [​](\#perspective "Direct link to perspective")

The CSS `perspective` of the flip animation. Defaults to `1000`.

## See also [​](\#see-also "Direct link to See also")

- [Presentations](/docs/transitions/presentations)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/cube.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
none()](/docs/transitions/presentations/none) [Next\
\
Custom presentations](/docs/transitions/presentations/custom)

- [Example](#example)
- [API](#api)
  - [`direction`](#direction)
  - [`perspective?`](#perspective)
- [See also](#see-also)
