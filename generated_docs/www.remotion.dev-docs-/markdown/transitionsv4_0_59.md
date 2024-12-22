---
title: Transitionsv4.0.59
source: Remotion Documentation
last_updated: 2024-12-22
---

# Transitionsv4.0.59

- [Home page](/)
- Designing visuals
- Transitions

On this page

# Transitions [v4.0.59](https://github.com/remotion-dev/remotion/releases/v4.0.59)

To transition between two types of absolutely positioned content, you can use the [`<TransitionSeries>`](/docs/transitions/transitionseries) component.

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

A

B

direction

from-leftfrom-bottomfrom-rightfrom-top

## Enter and exit animations [​](\#enter-and-exit-animations "Direct link to Enter and exit animations")

You don't necessarily have to use `<TransitionSeries>` for transitions between scenes. You can also use it to animate the entrace or exit of a scene by putting a transition first or last in the `<TransitionSeries>`.

[See example](/docs/transitions/transitionseries#enter-and-exit-animations)

## Duration [​](\#duration "Direct link to Duration")

In the above example, `A` is visible for 40 frames, `B` for 60 frames and the duration of the transition is 30 frames.

During this time, both slides are rendered. This means the total duration of the animation is `60 + 40 - 30 = 70`.

## Presentations [​](\#presentations "Direct link to Presentations")

A presentation determines the visual of animation.

[**Overview** \
\
List of available presentations](/docs/transitions/presentations) [**Custom presentations** \
\
Implement your own effect](/docs/transitions/presentations/custom) [A\
\
B\
\
**`fade()`** \
\
Animate the opacity of the scenes](/docs/transitions/presentations/fade) [A\
\
B\
\
**`slide()`** \
\
Slide in and push out the previous scene](/docs/transitions/presentations/slide) [A\
\
B\
\
**`wipe()`** \
\
Slide over the previous scene](/docs/transitions/presentations/wipe) [A\
\
B\
\
**`flip()`** \
\
Rotate the previous scene](/docs/transitions/presentations/flip) [A\
\
B\
\
**`clockWipe()`** \
\
Reveal the new scene in a circular movement](/docs/transitions/presentations/clock-wipe) [A\
\
B\
\
**`cube()`** \
\
Paid\
\
Rotate both scenes with 3D perspective](/docs/transitions/presentations/cube) [A\
\
B\
\
**`none()`** \
\
Have no visual effect.](/docs/transitions/presentations/none) [**Audio transitions** \
\
Add a sound effect to a transition](/docs/transitions/audio-transitions)

## Timings [​](\#timings "Direct link to Timings")

A timing determines how long the animation takes and the animation curve.

[**`springTiming()`** \
\
Transition with a `spring()`](/docs/transitions/timings/springtiming) [**`linearTiming()`** \
\
Transition linearly with optional Easing](/docs/transitions/timings/lineartiming) [**Custom timings** \
\
Implement your own timing](/docs/transitions/timings/custom)

## Audio transitions [​](\#audio-transitions "Direct link to Audio transitions")

See here how to include [audio effects](/docs/transitions/audio-transitions) in your transitions.

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

- [`@remotion/transitions`](/docs/transitions)
- [`<TransitionSeries>`](/docs/transitions/transitionseries)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Layers](/docs/layers) [Next\
\
Audio](/docs/using-audio)

- [Enter and exit animations](#enter-and-exit-animations)
- [Duration](#duration)
- [Presentations](#presentations)
- [Timings](#timings)
- [Audio transitions](#audio-transitions)
- [Getting the duration of a transition](#getting-the-duration-of-a-transition)
- [Rules](#rules)
- [See also](#see-also)
