---
title: <Series>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Series>

- [Home page](/)
- [remotion](/docs/remotion)
- <Series>

On this page

# <Series>

_Available from v.2.3.1_

Using this component, you can easily stitch together scenes that should play sequentially after another.

## Example [​](\#example "Direct link to Example")

### Code [​](\#code "Direct link to Code")

```

src/Example.tsx
tsx

import { Series } from "remotion";

export const Example: React.FC = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={40}>
        <Square color={"#3498db"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={20}>
        <Square color={"#5ff332"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={70}>
        <Square color={"#fdc321"} />
      </Series.Sequence>
    </Series>
  );
};
```

```

src/Example.tsx
tsx

import { Series } from "remotion";

export const Example: React.FC = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={40}>
        <Square color={"#3498db"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={20}>
        <Square color={"#5ff332"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={70}>
        <Square color={"#fdc321"} />
      </Series.Sequence>
    </Series>
  );
};
```

### Result [​](\#result "Direct link to Result")

0

0:00 / 0:05

## API [​](\#api "Direct link to API")

The `<Series />` component takes no props may only contain a list of `<Series.Sequence />` instances. A `<Series.Sequence />` component takes the following props:

This component is a high order component, and accepts, besides it's children, the following props:

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

For how many frames the sequence should be displayed. Children are unmounted if they are not within the time range of display.

Only the last `<Series.Sequence />` instance is allowed to have `Infinity` as a duration, all previous one must have a positive integer.

### `offset` [​](\#offset "Direct link to offset")

_optional_

Pass a positive number to delay the beginning of the sequence. Pass a negative number to start the sequence earlier, and to overlay the sequence with the one that comes before.

The offset does not apply to sequences that come before, but the sequences that come after it will also be shifted.

**Example 1**: Pass `10` to delay the sequence by 10 frames and create a blank space of 10 frames before it.

**Example 2**: Pass `-10` to start the sequence earlier and overlay the sequence on top of the previous one for 10 frames.

### `layout` [​](\#layout "Direct link to layout")

_optional_

Either `"absolute-fill"` _(default)_ or `"none"` By default, your sequences will be absolutely positioned, so they will overlay each other. If you would like to opt out of it and handle layouting yourself, pass `layout="none"`.

### `style` [v3.3.4](https://github.com/remotion-dev/remotion/releases/v3.3.4) [​](\#style "Direct link to style")

_optional_

CSS styles to be applied to the container. If `layout` is set to `none`, there is no container and setting this style is not allowed.

### `className` [v3.3.45](https://github.com/remotion-dev/remotion/releases/v3.3.45) [​](\#classname "Direct link to classname")

_optional_

A class name to be applied to the container. If `layout` is set to `none`, there is no container and setting this style is not allowed.

### `premountFor` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#premountfor "Direct link to premountfor")

_optional_

[Premount](/docs/player/premounting) the sequence for a set number of frames.

### `ref` [v3.3.4](https://github.com/remotion-dev/remotion/releases/v3.3.4) [​](\#ref "Direct link to ref")

_optional_

You can add a [React ref](https://react.dev/learn/manipulating-the-dom-with-refs) to a `<Series.Sequence>`. If you use TypeScript, you need to type it with `HTMLDivElement`:

```

src/Example.tsx
tsx

import React, { useRef } from "react";
import { Series } from "remotion";

export const Example: React.FC = () => {
  const first = useRef<HTMLDivElement>(null);
  const second = useRef<HTMLDivElement>(null);

  return (
    <Series>
      <Series.Sequence durationInFrames={40} ref={first}>
        <Square color={"#3498db"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={20} ref={second}>
        <Square color={"#5ff332"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={70}>
        <Square color={"#fdc321"} />
      </Series.Sequence>
    </Series>
  );
};
```

```

src/Example.tsx
tsx

import React, { useRef } from "react";
import { Series } from "remotion";

export const Example: React.FC = () => {
  const first = useRef<HTMLDivElement>(null);
  const second = useRef<HTMLDivElement>(null);

  return (
    <Series>
      <Series.Sequence durationInFrames={40} ref={first}>
        <Square color={"#3498db"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={20} ref={second}>
        <Square color={"#5ff332"} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={70}>
        <Square color={"#fdc321"} />
      </Series.Sequence>
    </Series>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/series/index.tsx)
- [`<Sequence />`](/docs/sequence)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/series.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Sequence>](/docs/sequence) [Next\
\
spring()](/docs/spring)

- [Example](#example)
  - [Code](#code)
  - [Result](#result)
- [API](#api)
  - [`durationInFrames`](#durationinframes)
  - [`offset`](#offset)
  - [`layout`](#layout)
  - [`style`](#style)
  - [`className`](#classname)
  - [`premountFor`](#premountfor)
  - [`ref`](#ref)
- [See also](#see-also)
