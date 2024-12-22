---
title: <TransitionSeries>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <TransitionSeries>

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- <TransitionSeries>

On this page

# <TransitionSeries>

_available from v4.0.59_

The `<TransitionSeries>` component behaves the same as the [`<Series>`](/docs/series) component, but allows for `<TransitionSeries.Transition>` components to be rendered between `<TransitionSeries.Sequence>` components.

Each transition consists of two parts:

[1](#1)

`presentation`: The effect that is being used, for
example [`fade()`](/docs/transitions/presentations/fade) or [`wipe()`](/docs/transitions/presentations/wipe).

[2](#2)

`timing`: The duration and the progress curve, for
example [`springTiming()`](/docs/transitions/timings/springtiming) or [`linearTiming()`](/docs/transitions/timings/lineartiming)

This package has some presentations and timings built-in, but custom ones can be created as well.

```

MyComp.tsx
tsx

import {
  linearTiming,
  springTiming,
  TransitionSeries,
} from "@remotion/transitions";

import { fade } from "@remotion/transitions/fade";
import { wipe } from "@remotion/transitions/wipe";

export const MyComp: React.FC = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="blue" />
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        timing={springTiming({ config: { damping: 200 } })}
        presentation={fade()}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="black" />
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        timing={linearTiming({ durationInFrames: 30 })}
        presentation={wipe()}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="white" />
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

```

MyComp.tsx
tsx

import {
  linearTiming,
  springTiming,
  TransitionSeries,
} from "@remotion/transitions";

import { fade } from "@remotion/transitions/fade";
import { wipe } from "@remotion/transitions/wipe";

export const MyComp: React.FC = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="blue" />
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        timing={springTiming({ config: { damping: 200 } })}
        presentation={fade()}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="black" />
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        timing={linearTiming({ durationInFrames: 30 })}
        presentation={wipe()}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Fill color="white" />
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

## API [​](\#api "Direct link to API")

### `<TransitionSeries>` [​](\#transitionseries "Direct link to transitionseries")

Inherits the [`from`](/docs/sequence#from), [`name`](/docs/sequence#name), [`className`](/docs/sequence#classname), and [`style`](/docs/sequence#style) props from [`<Sequence>`](/docs/sequence).

The `<TransitionSeries>` component can only contain `<TransitionSeries.Sequence>` and `<TransitionSeries.Transition>` components.

### `<TransitionSeries.Sequence>` [​](\#transitionseriessequence "Direct link to transitionseriessequence")

Inherits the [`durationInFrames`](/docs/sequence#durationinframes), [`name`](/docs/sequence#name), [`className`](/docs/sequence#classname), [`style`](/docs/sequence#style), [`premountFor`](/docs/sequence#premountfor) and [`layout`](/docs/sequence#layout) props from [`<Sequence>`](/docs/sequence).

As children, put the contents of your scene.

### `<TransitionSeries.Transition>` [​](\#transitionseriestransition "Direct link to transitionseriestransition")

Takes two props:

- `timing`: A timing of type `TransitionTiming`. See [Timings](/docs/transitions/timings) for more information.
- `presentation?`: A presentation of type `TransitionPresentation`. If not specified, the default value is `slide()`. See [Presentations](/docs/transitions/presentations) for more information.

Must be a direct child of `<TransitionSeries>`.

At least one `<TransitionSeries.Sequence>` component must come before or after the `<TransitionSeries.Transition>` component.

It is not allowed for two `<TransitionSeries.Transition>` components to be next to each other.

## Enter and exit animations [​](\#enter-and-exit-animations "Direct link to Enter and exit animations")

You don't necessarily have to use `<TransitionSeries>` for transitions between scenes. You can also use it to animate the entrace or exit of a scene by putting a transition first or last in the `<TransitionSeries>`.

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
    </TransitionSeries>
  );
};
```

## Duration of a `<TransitionSeries>` [​](\#duration-of-a-transitionseries "Direct link to duration-of-a-transitionseries")

Consider this example:

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

`A` is visible for 40 frames, `B` for 60 frames and the duration of the transition is 30 frames.

During this time, both slides are rendered. This means the total duration of the animation is `60 + 40 - 30 = 70`.

Example with 3 slides

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
      <TransitionSeries.Transition
        presentation={slide()}
        timing={linearTiming({ durationInFrames: 45 })}
      />
      <TransitionSeries.Sequence durationInFrames={90}>
        <Letter color="green">C</Letter>
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
      <TransitionSeries.Transition
        presentation={slide()}
        timing={linearTiming({ durationInFrames: 45 })}
      />
      <TransitionSeries.Sequence durationInFrames={90}>
        <Letter color="green">C</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

- The first slide is shown for 40 frames
- The second slide is shown for 60 frames
- The third slide is shown for 90 frames
- There are two transitions, one 30 frames long and one 45 frames long

1. Sum up the durations: `40 + 60 + 90 = 190`
2. Subtract the duration of the transitions: `190 - 30 - 45 = 115`

## Getting the duration of a transition [​](\#getting-the-duration-of-a-transition "Direct link to Getting the duration of a transition")

You can get the duration of a transition by calling `getDurationInFrames()` on the timing:

```

Assuming a framerate of 30fps
tsx

import { springTiming } from "@remotion/transitions";

springTiming({ config: { damping: 200 } }).getDurationInFrames({ fps: 30 }); // 23
```

```

Assuming a framerate of 30fps
tsx

import { springTiming } from "@remotion/transitions";

springTiming({ config: { damping: 200 } }).getDurationInFrames({ fps: 30 }); // 23
```

## Rules [​](\#rules "Direct link to Rules")

[1](#1)

It is forbidden to have a transition that is longer than the duration
of the previous or next sequence.

[2](#2) There can be no two transitions next to each other.

[3](#3) There must be at least one sequence before or after a transition.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/TransitionSeries.tsx)
- [Transitions](/docs/transitioning)
- [Timings](/docs/transitions/timings)
- [Presentations](/docs/transitions/presentations)
- [`<Series>`](/docs/series)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/transitionseries.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/transitions/) [Next\
\
useTransitionProgress()](/docs/transitions/use-transition-progress)

- [API](#api)
  - [`<TransitionSeries>`](#transitionseries)
  - [`<TransitionSeries.Sequence>`](#transitionseriessequence)
  - [`<TransitionSeries.Transition>`](#transitionseriestransition)
- [Enter and exit animations](#enter-and-exit-animations)
- [Duration of a `<TransitionSeries>`](#duration-of-a-transitionseries)
- [Getting the duration of a transition](#getting-the-duration-of-a-transition)
- [Rules](#rules)
- [See also](#see-also)
