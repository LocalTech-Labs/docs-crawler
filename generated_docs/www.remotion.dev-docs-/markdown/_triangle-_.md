---
title: <Triangle />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Triangle />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Triangle />

On this page

# <Triangle />

_Part of the [` @remotion/shapes`](/docs/shapes) package._

Renders an SVG element containing a triangle with same length on all sides.

## Explorer [​](\#explorer "Direct link to Explorer")

length

`200`

edgeRoundness

``

cornerRadius

`0`

direction

updownleftright

debug

## Example [​](\#example "Direct link to Example")

```

src/Triangle.tsx
tsx

import { Triangle } from "@remotion/shapes";
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
      <Triangle length={100} fill="red" direction="left" />
    </AbsoluteFill>
  );
};
```

```

src/Triangle.tsx
tsx

import { Triangle } from "@remotion/shapes";
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
      <Triangle length={100} fill="red" direction="left" />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `length`

_number_

The length of one triangle side.

### `direction`

_"left" \| "right" \| "up" \| "down"_

The direction of the triangle.

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

### `cornerRadius`

_number_

Rounds the corner using an arc. Similar to CSS's `border-radius`. Cannot be used together with `edgeRoundness`.

### `edgeRoundness`

_null \| number_

Allows to modify the shape by rounding the edges using bezier curves. Default `null`.

`0` will lead to a rotated triangle being drawn inside the natural dimensions of the triangle.`Math.sqrt(2) - 1` willdraw a circle.`1` will draw a shape similar to a "squircle" but as a triangle.Values below `0` and above `1` may result in other interesting shapes. Pictured: `2`.

Cannot be used together with `cornerRadius`.

### `debug`

_boolean_

If enabled, draws the lines for Bézier curves. This is meant for debugging, note that the visuals may change in any version.

### Other props

All other props that can be passed to a `<path>` are accepted and will be forwarded.

## See also [​](\#see-also "Direct link to See also")

- [makeTriangle()](/docs/shapes/triangle)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/triangle.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/triangle.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Rect />](/docs/shapes/rect) [Next\
\
<Circle />](/docs/shapes/circle)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
