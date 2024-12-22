---
title: fade()
source: Remotion Documentation
last_updated: 2024-12-22
---

# fade()

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- fade()

On this page

# fade()

A simple fade animation. The incoming slide fades in over the outgoing slide, while the outgoing slide does not change. Works only if the incoming slide is fully opaque.

A

B

## Example [​](\#example "Direct link to Example")

```

FadeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { fade } from "@remotion/transitions/fade";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={fade()}
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

FadeTransition.tsx
tsx

import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { fade } from "@remotion/transitions/fade";

const BasicTransition = () => {
  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={40}>
        <Letter color="#0b84f3">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={fade()}
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

An object which takes:

### `enterStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#enterstyle "Direct link to enterstyle")

The style of the container element when the scene is is entering.

### `exitStyle?` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#exitstyle "Direct link to exitstyle")

The style of the container element when the scene is exiting.

### `shouldFadeOutExitingScene?` [v4.0.166](https://github.com/remotion-dev/remotion/releases/v4.0.166) [​](\#shouldfadeoutexitingscene "Direct link to shouldfadeoutexitingscene")

Whether the exiting scene should fade out or not. Default `false`.

note

The default is `false` because if both the entering and existing scene are semi-opaque, the whole scene will be semi-opaque, which will make the content underneath shine though.

We recommend for transitioning between fully opaque scenes setting this to `false`.

If the scenes are not fully covered (like fading between overlays), we recommend setting this to `false`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this presentation](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/presentations/fade.tsx)
- [Presentations](/docs/transitions/presentations)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/fade.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Presentations](/docs/transitions/presentations/) [Next\
\
slide()](/docs/transitions/presentations/slide)

- [Example](#example)
- [API](#api)
  - [`enterStyle?`](#enterstyle)
  - [`exitStyle?`](#exitstyle)
  - [`shouldFadeOutExitingScene?`](#shouldfadeoutexitingscene)
- [See also](#see-also)
