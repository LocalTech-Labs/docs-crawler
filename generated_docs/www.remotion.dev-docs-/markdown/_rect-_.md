---
title: <Rect />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Rect />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Rect />

On this page

# <Rect />

_Part of the [` @remotion/shapes`](/docs/shapes) package._

Renders an SVG element containing a rectangle.

## Explorer [​](\#explorer "Direct link to Explorer")

width

`200`

height

`200`

cornerRadius

`0`

edgeRoundness

``

debug

## Example [​](\#example "Direct link to Example")

```

src/Rect.tsx
tsx

import { Rect } from "@remotion/shapes";
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
      <Rect width={200} height={200} fill="red" />
    </AbsoluteFill>
  );
};
```

```

src/Rect.tsx
tsx

import { Rect } from "@remotion/shapes";
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
      <Rect width={200} height={200} fill="red" />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `width`

_number_

The width of the rectangle.

### `height`

_number_

The height of the rectangle.

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

`0` will lead to a rotated rectangle being drawn inside the natural dimensions of the rectangle.`(4 * (Math.sqrt(2) - 1)) / 3` willdraw a circle.`1` will draw a squircle.Values below `0` and above `1` are possible and may result in interesting shapes. Pictured: `2`

Cannot be used together with `cornerRadius`.

### `debug`

_boolean_

If enabled, draws the lines for Bézier curves. This is meant for debugging, note that the visuals may change in any version.

### Other props

All other props that can be passed to a `<path>` are accepted and will be forwarded.

## See also [​](\#see-also "Direct link to See also")

- [makeRect()](/docs/shapes/rect)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/rect.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/rect.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/shapes](/docs/shapes/) [Next\
\
<Triangle />](/docs/shapes/triangle)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
