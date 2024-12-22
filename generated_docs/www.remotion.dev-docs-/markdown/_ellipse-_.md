---
title: <Ellipse />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Ellipse />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Ellipse />

On this page

# <Ellipse />

_Part of the [` @remotion/shapes`](/docs/shapes) package._

Renders an SVG element drawing an ellipse.

## Explorer [​](\#explorer "Direct link to Explorer")

rx

`150`

ry

`200`

## Example [​](\#example "Direct link to Example")

```

src/Ellipse.tsx
tsx

import { Ellipse } from "@remotion/shapes";
import { AbsoluteFill } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill
      style={{
        backgroundColor: "white",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Ellipse rx={100} ry={50} fill="green" stroke="red" strokeWidth={1} />
    </AbsoluteFill>
  );
};
```

```

src/Ellipse.tsx
tsx

import { Ellipse } from "@remotion/shapes";
import { AbsoluteFill } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill
      style={{
        backgroundColor: "white",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Ellipse rx={100} ry={50} fill="green" stroke="red" strokeWidth={1} />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `rx`

_number_

The radius of the ellipse on the X axis.

### `ry`

_number_

The radius of the ellipse on the Y axis.

### `fill`

_string_

The color of the shape.

### `stroke`

_string_

The color of the stroke. Should be used together with `strokeWidth`.

### `strokeWidth`

_string_

The width of the stroke. Should be used together with `stroke`.

### `style`

_string_

CSS properties that will be applied to the `<svg>` tag. Default style: `overflow: 'visible'`

### `pathStyle`

_string_

CSS properties that will be applied to the `<path>` tag. Default style: `transform-box: 'fill-box'` and a dynamically calculated `transform-origin` which is the center of the shape, so that the shape rotates around its center by default.

### `strokeDasharray`

_string_

Allows to animate a path. See [evolvePath()](/docs/paths/evolve-path) for an example.

### `strokeDashoffset`

_string_

Allows to animate a path. See [evolvePath()](/docs/paths/evolve-path) for an example.

### Other props

All other props that can be passed to a `<path>` are accepted and will be forwarded.

## See also [​](\#see-also "Direct link to See also")

- [makeEllipse()](/docs/shapes/ellipse)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/ellipse.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/ellipse.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Circle />](/docs/shapes/circle) [Next\
\
<Star />](/docs/shapes/star)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
