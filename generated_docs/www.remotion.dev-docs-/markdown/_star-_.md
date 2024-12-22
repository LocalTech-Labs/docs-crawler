---
title: <Star />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Star />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Star />

On this page

# <Star />

_Part of the [` @remotion/shapes`](/docs/shapes) package._

Renders an SVG element containing a star.

## Explorer [​](\#explorer "Direct link to Explorer")

innerRadius

`100`

outerRadius

`200`

edgeRoundness

``

points

`5`

cornerRadius

`0`

## Example [​](\#example "Direct link to Example")

```

src/Star.tsx
tsx

import { Star } from "@remotion/shapes";
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
      <Star points={5} innerRadius={150} outerRadius={200} fill="red" />
    </AbsoluteFill>
  );
};
```

```

src/Star.tsx
tsx

import { Star } from "@remotion/shapes";
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
      <Star points={5} innerRadius={150} outerRadius={200} fill="red" />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `points`

_number_

The amount of points of the star.

### `innerRadius`

_number_

The inner radius of the star.

### `outerRadius`

_number_

The outer radius of the star.

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

- [makeStar()](/docs/shapes/star)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/star.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/star.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Ellipse />](/docs/shapes/ellipse) [Next\
\
<Pie />](/docs/shapes/pie)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
