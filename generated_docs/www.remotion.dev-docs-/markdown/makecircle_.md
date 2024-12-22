---
title: makeCircle()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeCircle()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makeCircle()

On this page

# makeCircle()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates a circle SVG path.

## Example [​](\#example "Direct link to Example")

```

circle.ts
tsx

import { makeCircle } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeCircle({ radius: 50 });

console.log(path); // M 0 0 m -50, 0 a 50,50 0 1,0 100,0  50,50 0 1,0 -100,0
console.log(width); // 100
console.log(height); // 100
console.log(transformOrigin); // '50 50'
```

```

circle.ts
tsx

import { makeCircle } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeCircle({ radius: 50 });

console.log(path); // M 0 0 m -50, 0 a 50,50 0 1,0 100,0  50,50 0 1,0 -100,0
console.log(width); // 100
console.log(height); // 100
console.log(transformOrigin); // '50 50'
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `radius`

_number_

The radius of the circle.

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the circle. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the circle. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Circle />`](/docs/shapes/circle).

## See also [​](\#see-also "Direct link to See also")

- [<Circle />](/docs/shapes/circle)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-circle.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-circle.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeTriangle()](/docs/shapes/make-triangle) [Next\
\
makeEllipse()](/docs/shapes/make-ellipse)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [See also](#see-also)
