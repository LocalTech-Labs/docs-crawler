---
title: makeRect()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeRect()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makeRect()

On this page

# makeRect()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates an SVG rectangle.

## Example [​](\#example "Direct link to Example")

```

rect.ts
tsx

import { makeRect } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeRect({
  width: 100,
  height: 100,
});

console.log(path); // M 0 0 l 100 0 l 0 100 l -100 0 Z
console.log(width); // 100
console.log(height); // 100
console.log(transformOrigin); // '50 50'
```

```

rect.ts
tsx

import { makeRect } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeRect({
  width: 100,
  height: 100,
});

console.log(path); // M 0 0 l 100 0 l 0 100 l -100 0 Z
console.log(width); // 100
console.log(height); // 100
console.log(transformOrigin); // '50 50'
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `width`

_number_

The width of the rectangle.

### `height`

_number_

The height of the rectangle.

### `edgeRoundness`

_null \| number_

Allows to modify the shape by rounding the edges using bezier curves. Default `null`.

`0` will lead to a rotated rectangle being drawn inside the natural dimensions of the rectangle.`(4 * (Math.sqrt(2) - 1)) / 3` willdraw a circle.`1` will draw a squircle.Values below `0` and above `1` are possible and may result in interesting shapes. Pictured: `2`

Cannot be used together with `cornerRadius`.

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the rect. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the rect. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Rect />`](/docs/shapes/rect).

## See also [​](\#see-also "Direct link to See also")

- [<Rect />](/docs/shapes/rect)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-rect.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-rect.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Polygon />](/docs/shapes/polygon) [Next\
\
makeTriangle()](/docs/shapes/make-triangle)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [See also](#see-also)
