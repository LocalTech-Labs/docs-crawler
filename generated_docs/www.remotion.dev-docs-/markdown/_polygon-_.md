---
title: <Polygon />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Polygon />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Polygon />

On this page

# <Polygon />

_Part of the [` @remotion/shapes`](/docs/shapes) package._

Renders an SVG element containing a polygon.

## Explorer [​](\#explorer "Direct link to Explorer")

points

`3`

radius

`100`

cornerRadius

`0`

edgeRoundness

``

## Example [​](\#example "Direct link to Example")

```

src/Polygon.tsx
tsx

import { Polygon } from "@remotion/shapes";
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
      <Polygon points={5} radius={80} />
    </AbsoluteFill>
  );
};
```

```

src/Polygon.tsx
tsx

import { Polygon } from "@remotion/shapes";
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
      <Polygon points={5} radius={80} />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `points`

_number_

The number of points in the polygon.

### `radius`

_number_

The radius of the polygon.

### `edgeRoundness`

_number \| null_

Allows to modify the shape by rounding the edges using bezier curves. Default null.

### `cornerRadius`

_number_

Rounds the corner using an arc. Similar to CSS's border-radius. Cannot be used together with edgeRoundness.

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

### Other props

All other props that can be passed to a `<path>` are accepted and will be forwarded.

## See also [​](\#see-also "Direct link to See also")

- [makePolygon()](/docs/shapes/polygon)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/polygon.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/polygon.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Pie />](/docs/shapes/pie) [Next\
\
makeRect()](/docs/shapes/make-rect)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
