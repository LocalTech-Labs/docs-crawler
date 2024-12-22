---
title: makeTriangle()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeTriangle()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makeTriangle()

On this page

# makeTriangle()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates an SVG path for a triangle with the same length on all sides.

## Example [​](\#example "Direct link to Example")

```

triangle.ts
tsx

import { makeTriangle } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeTriangle({
  length: 100,
  direction: "right",
});

console.log(path); // M 0 0 L 0 100 L 86.60254037844386 50 z
console.log(width); // 86.60254037844386
console.log(height); // 100
console.log(transformOrigin); // '28.867513459481287 50'
```

```

triangle.ts
tsx

import { makeTriangle } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeTriangle({
  length: 100,
  direction: "right",
});

console.log(path); // M 0 0 L 0 100 L 86.60254037844386 50 z
console.log(width); // 86.60254037844386
console.log(height); // 100
console.log(transformOrigin); // '28.867513459481287 50'
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `length`

_number_

The length of one triangle side.

### `direction`

_"left" \| "right" \| "up" \| "down"_

The direction of the triangle.

### `edgeRoundness`

_null \| number_

Allows to modify the shape by rounding the edges using bezier curves. Default `null`.

`0` will lead to a rotated triangle being drawn inside the natural dimensions of the triangle.`Math.sqrt(2) - 1` willdraw a circle.`1` will draw a shape similar to a "squircle" but as a triangle.Values below `0` and above `1` may result in other interesting shapes. Pictured: `2`.

Cannot be used together with `cornerRadius`.

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the triangle. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the triangle. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Triangle />`](/docs/shapes/triangle).

## Credits [​](\#credits "Direct link to Credits")

Source code partially taken from [this StackBlitz](https://stackblitz.com/edit/react-triangle-svg?file=index.js).

## See also [​](\#see-also "Direct link to See also")

- [<Triangle />](/docs/shapes/triangle)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-triangle.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-triangle.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeRect()](/docs/shapes/make-rect) [Next\
\
makeCircle()](/docs/shapes/make-circle)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [Credits](#credits)
- [See also](#see-also)
