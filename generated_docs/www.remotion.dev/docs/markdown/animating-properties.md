---
title: Animating properties
source: Remotion Documentation
last_updated: 2024-12-21
---

# Animating properties

- [Home page](/)
- Getting started
- Animating properties

On this page

# Animating properties

Animation works by changing properties over time.

Let's create a simple fade in animation.

If we want to fade the text in over 60 frames, we need to gradually change the `opacity` over time so that it goes from 0 to 1.

```

FadeIn.tsx
tsx

export const FadeIn = () => {
  const frame = useCurrentFrame();

  const opacity = Math.min(1, frame / 60);

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "white",
        fontSize: 80,
      }}
    >
      <div style={{ opacity: opacity }}>Hello World!</div>
    </AbsoluteFill>
  );
};
```

```

FadeIn.tsx
tsx

export const FadeIn = () => {
  const frame = useCurrentFrame();

  const opacity = Math.min(1, frame / 60);

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "white",
        fontSize: 80,
      }}
    >
      <div style={{ opacity: opacity }}>Hello World!</div>
    </AbsoluteFill>
  );
};
```

## Using the interpolate helper function [​](\#using-the-interpolate-helper-function "Direct link to Using the interpolate helper function")

Using the [`interpolate()`](/docs/interpolate) function can make animations more readable. The above animation can also be written as:

```

tsx

import { interpolate } from "remotion";

const opacity = interpolate(frame, [0, 60], [0, 1], {
  /*                        ^^^^^   ^^^^^    ^^^^
  Variable to interpolate ----|       |       |
  Input range ------------------------|       |
  Output range -------------------------------|  */
  extrapolateRight: "clamp",
});
```

```

tsx

import { interpolate } from "remotion";

const opacity = interpolate(frame, [0, 60], [0, 1], {
  /*                        ^^^^^   ^^^^^    ^^^^
  Variable to interpolate ----|       |       |
  Input range ------------------------|       |
  Output range -------------------------------|  */
  extrapolateRight: "clamp",
});
```

In this example, we map the frames 0 to 60 to their opacity values `(0, 0.0166, 0.033, 0.05 ...`) and use the [`extrapolateRight`](/docs/interpolate#extrapolateright) setting to clamp the output so that it never becomes bigger than 1.

## Using spring animations [​](\#using-spring-animations "Direct link to Using spring animations")

Spring animations are a natural animation primitive. By default, they animate from 0 to 1 over time. This time, let's animate the scale of the text.

```

tsx

import { spring, useCurrentFrame, useVideoConfig } from "remotion";

export const MyVideo = () => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  const scale = spring({
    fps,
    frame,
  });

  return (
    <div
      style={{
        flex: 1,
        textAlign: "center",
        fontSize: "7em",
      }}
    >
      <div style={{ transform: `scale(${scale})` }}>Hello World!</div>
    </div>
  );
};
```

```

tsx

import { spring, useCurrentFrame, useVideoConfig } from "remotion";

export const MyVideo = () => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  const scale = spring({
    fps,
    frame,
  });

  return (
    <div
      style={{
        flex: 1,
        textAlign: "center",
        fontSize: "7em",
      }}
    >
      <div style={{ transform: `scale(${scale})` }}>Hello World!</div>
    </div>
  );
};
```

You should see the text jump in.

The default spring configuration leads to a little bit of overshoot, meaning the text will bounce a little bit. See the documentation page for [`spring()`](/docs/spring) to learn how to customize it.

## Always animate using `useCurrentFrame()` [​](\#always-animate-using-usecurrentframe "Direct link to always-animate-using-usecurrentframe")

Watch out for flickering issues during rendering that arise if you write animations that are not driven by [`useCurrentFrame()`](/docs/use-current-frame) – for example CSS transitions.

[Read more about how Remotion's rendering works](/docs/flickering) \- understanding it will help you avoid issues down the road.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/animating-properties.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
The fundamentals](/docs/the-fundamentals) [Next\
\
Reuse components](/docs/reusability)

- [Using the interpolate helper function](#using-the-interpolate-helper-function)
- [Using spring animations](#using-spring-animations)
- [Always animate using `useCurrentFrame()`](#always-animate-using-usecurrentframe)
