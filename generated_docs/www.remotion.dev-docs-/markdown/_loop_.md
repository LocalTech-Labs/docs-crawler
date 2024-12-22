---
title: <Loop>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Loop>

- [Home page](/)
- [remotion](/docs/remotion)
- <Loop>

On this page

# <Loop>

_available from v2.5.0_

The `<Loop />` component allows you to quickly lay out an animation so it repeats itself.

```

MyComp.tsx
tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50} times={2}>
      <BlueSquare />
    </Loop>
  );
};
```

```

MyComp.tsx
tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50} times={2}>
      <BlueSquare />
    </Loop>
  );
};
```

note

Good to know: You can nest loops within each other and they will cascade.

## API [​](\#api "Direct link to API")

The Loop component is a high order component and accepts, besides it's children, the following props:

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

How many frames one iteration of the loop should be long.

### `times` [​](\#times "Direct link to times")

_optional_

How many times to loop the content. By default it will be `Infinity`.

### `layout` [​](\#layout "Direct link to layout")

_optional_

Either `"absolute-fill"` _(default)_ or `"none"`.

By default, your content will be absolutely positioned.

If you would like to disable layout side effects, pass `layout="none"`.

### `style` [v3.3.4](https://github.com/remotion-dev/remotion/releases/v3.3.4) [​](\#style "Direct link to style")

_optional_

CSS styles to be applied to the container. If `layout` is set to `none`, there is no container and setting this style is not allowed.

## Examples [​](\#examples "Direct link to Examples")

All the examples below are based on the following animation of a blue square:

0

0:00 / 0:05

```

tsx

const MyComp = () => {
  return <BlueSquare />;
};
```

```

tsx

const MyComp = () => {
  return <BlueSquare />;
};
```

### Continuous loop [​](\#continuous-loop "Direct link to Continuous loop")

0

0:00 / 0:05

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50}>
      <BlueSquare />
    </Loop>
  );
};
```

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50}>
      <BlueSquare />
    </Loop>
  );
};
```

### Fixed count loop [​](\#fixed-count-loop "Direct link to Fixed count loop")

0

0:00 / 0:05

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50} times={2}>
      <BlueSquare />
    </Loop>
  );
};
```

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={50} times={2}>
      <BlueSquare />
    </Loop>
  );
};
```

### Nested loop [​](\#nested-loop "Direct link to Nested loop")

0

0:00 / 0:05

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={75}>
      <Loop durationInFrames={30}>
        <BlueSquare />
      </Loop>
    </Loop>
  );
};
```

```

tsx

const MyComp = () => {
  return (
    <Loop durationInFrames={75}>
      <Loop durationInFrames={30}>
        <BlueSquare />
      </Loop>
    </Loop>
  );
};
```

## `useLoop()` [v4.0.142](https://github.com/remotion-dev/remotion/releases/v4.0.142) [​](\#useloop "Direct link to useloop")

A child component can use the `Loop.useLoop()` hook to get information about the current loop.

You should check for `null`, which is the case if the component is not inside a loop.

If inside a loop, an object with two properties is returned:

- `durationInFrames`: The duration of the loop in frames as passed to the `<Loop />` component.
- `iteration`: The current iteration of the loop, starting at 0.

```

tsx

import React from "react";
import { Loop, useCurrentFrame } from "remotion";

const LoopedVideo: React.FC = () => {
  return (
    <Loop durationInFrames={50} times={3} name="MyLoop">
      <Child />
    </Loop>
  );
};

const Child = () => {
  const frame = useCurrentFrame(); // 75
  const loop = Loop.useLoop();

  if (loop === null) {
    // Not inside a loop
  } else {
    console.log(loop.durationInFrames); // 50
    console.log(loop.iteration); // 1
  }

  return <div>{frame}</div>;
};
```

```

tsx

import React from "react";
import { Loop, useCurrentFrame } from "remotion";

const LoopedVideo: React.FC = () => {
  return (
    <Loop durationInFrames={50} times={3} name="MyLoop">
      <Child />
    </Loop>
  );
};

const Child = () => {
  const frame = useCurrentFrame(); // 75
  const loop = Loop.useLoop();

  if (loop === null) {
    // Not inside a loop
  } else {
    console.log(loop.durationInFrames); // 50
    console.log(loop.iteration); // 1
  }

  return <div>{frame}</div>;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/loop/index.tsx)
- [`<Sequence>`](/docs/sequence)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/loop.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
interpolate()](/docs/interpolate) [Next\
\
measureSpring()](/docs/measure-spring)

- [API](#api)
  - [`durationInFrames`](#durationinframes)
  - [`times`](#times)
  - [`layout`](#layout)
  - [`style`](#style)
- [Examples](#examples)
  - [Continuous loop](#continuous-loop)
  - [Fixed count loop](#fixed-count-loop)
  - [Nested loop](#nested-loop)
- [`useLoop()`](#useloop)
- [See also](#see-also)
