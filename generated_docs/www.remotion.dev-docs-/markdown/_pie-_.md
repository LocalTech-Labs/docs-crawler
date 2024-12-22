---
title: <Pie />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Pie />

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- <Pie />

On this page

# <Pie />

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Renders an SVG element drawing a pie piece.

## Explorer [​](\#explorer "Direct link to Explorer")

radius

`200`

progress

`0.5`

rotation

`0`

closePath

counterClockwise

showStrokeInsteadPreviewOnly

## Example [​](\#example "Direct link to Example")

```

src/Pie.tsx
tsx

import { Pie } from "@remotion/shapes";
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
      <Pie
        radius={100}
        progress={0.5}
        fill="green"
        stroke="red"
        strokeWidth={1}
      />
    </AbsoluteFill>
  );
};
```

```

src/Pie.tsx
tsx

import { Pie } from "@remotion/shapes";
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
      <Pie
        radius={100}
        progress={0.5}
        fill="green"
        stroke="red"
        strokeWidth={1}
      />
    </AbsoluteFill>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `radius`

_number_

The radius of the circle.

### `progress`

_number_

The percentage of the circle that is filled. `0` means fully empty, `1` means fully filled.

### `counterClockwise`

_boolean_

If set, the circle gets filled counterclockwise instead of clockwise. Default false.

### `closePath`

_boolean_

If set to `false`, no path to the middle of the circle will be drawn, leading to an open arc. Default `true`.

### `rotation`

_boolean_

Add rotation to the path. `0` means no rotation, `Math.PI * 2` means 1 full clockwise rotation

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

- [makePie()](/docs/shapes/pie)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/components/pie.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/pie.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Star />](/docs/shapes/star) [Next\
\
<Polygon />](/docs/shapes/polygon)

- [Explorer](#explorer)
- [Example](#example)
- [Props](#props)
- [See also](#see-also)
