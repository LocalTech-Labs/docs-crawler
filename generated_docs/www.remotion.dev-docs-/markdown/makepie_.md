---
title: makePie()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makePie()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makePie()

On this page

# makePie()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates a piece of pie SVG path.

## Example [​](\#example "Direct link to Example")

```

pie.ts
tsx

import { makePie } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makePie({
  radius: 100,
  progress: 0.5,
});

console.log(path); // M 100 0 A 100 100 0 0 1 100 200 L 100 100 z
console.log(width); // 200
console.log(height); // 200
console.log(transformOrigin); // '100 100'
```

```

pie.ts
tsx

import { makePie } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makePie({
  radius: 100,
  progress: 0.5,
});

console.log(path); // M 100 0 A 100 100 0 0 1 100 200 L 100 100 z
console.log(width); // 200
console.log(height); // 200
console.log(transformOrigin); // '100 100'
```

## Arguments [​](\#arguments "Direct link to Arguments")

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

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the pie. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the pie. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Pie />`](/docs/shapes/pie).

## See also [​](\#see-also "Direct link to See also")

- [<Pie />](/docs/shapes/pie)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-pie.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-pie.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeStar()](/docs/shapes/make-star) [Next\
\
makePolygon()](/docs/shapes/make-polygon)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [See also](#see-also)
