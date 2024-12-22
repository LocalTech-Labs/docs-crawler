---
title: useVideoConfig()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useVideoConfig()

- [Home page](/)
- [remotion](/docs/remotion)
- useVideoConfig()

On this page

# useVideoConfig()

With this hook, you can retrieve some info about the composition you are in.

## Example [​](\#example "Direct link to Example")

```

tsx

import React from "react";
import { useVideoConfig } from "remotion";

export const MyComp: React.FC = () => {
  const { width, height, fps, durationInFrames } = useVideoConfig();
  console.log(width); // 1920
  console.log(height); // 1080
  console.log(fps); // 30;
  console.log(durationInFrames); // 300

  return <div>Hello World!</div>;
};
```

```

tsx

import React from "react";
import { useVideoConfig } from "remotion";

export const MyComp: React.FC = () => {
  const { width, height, fps, durationInFrames } = useVideoConfig();
  console.log(width); // 1920
  console.log(height); // 1080
  console.log(fps); // 30;
  console.log(durationInFrames); // 300

  return <div>Hello World!</div>;
};
```

## Return value [​](\#return-value "Direct link to Return value")

A object with the following properties:

### `width` [​](\#width "Direct link to width")

The width of the composition in pixels, or the `width` of a [`<Sequence>`](/docs/sequence) if the component that calls `useVideoConfig()` is a child of a [`<Sequence>`](/docs/sequence) that defines a width.

### `height` [​](\#height "Direct link to height")

The height of the composition in pixels, or the `height` of a [`<Sequence>`](/docs/sequence) if the component that calls `useVideoConfig()` is a child of a [`<Sequence>`](/docs/sequence) that defines a height.

### `fps` [​](\#fps "Direct link to fps")

The frame rate of the composition, in frames per seconds.

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

The duration of the composition in frames or the `durationInFrames` of a [`<Sequence>`](/docs/sequence) if the component that calls `useVideoConfig()` is a child of a [`<Sequence>`](/docs/sequence) that defines a `durationInFrames`.

### `id` [​](\#id "Direct link to id")

The ID of the composition. This is the same as the `id` prop of the [`<Composition>`](/docs/composition) component.

### `defaultProps` [​](\#defaultprops "Direct link to defaultprops")

The object that you have defined as the `defaultProps` in your composition.

### `props` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#props "Direct link to props")

The props that were passed to the composition, after [all transformations](/docs/props-resolution).

### `defaultCodec` [v4.0.54](https://github.com/remotion-dev/remotion/releases/v4.0.54) [​](\#defaultcodec "Direct link to defaultcodec")

The default codec that is used for rendering this composition. Use [`calculateMetadata()`](/docs/calculate-metadata) to modify it.

## See also [​](\#see-also "Direct link to See also")

These properties are controlled by passing them as props to [`<Composition>`](/docs/composition). Read the page about [the fundamentals](/docs/the-fundamentals) to read how to setup a Remotion project.

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/use-video-config.ts)
- [useCurrentFrame()](/docs/use-current-frame)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-video-config.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useCurrentScale()](/docs/use-current-scale) [Next\
\
VERSION](/docs/version)

- [Example](#example)
- [Return value](#return-value)
  - [`width`](#width)
  - [`height`](#height)
  - [`fps`](#fps)
  - [`durationInFrames`](#durationinframes)
  - [`id`](#id)
  - [`defaultProps`](#defaultprops)
  - [`props`](#props)
  - [`defaultCodec`](#defaultcodec)
- [See also](#see-also)
