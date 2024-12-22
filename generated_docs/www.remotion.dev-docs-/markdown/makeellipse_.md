---
title: makeEllipse()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeEllipse()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makeEllipse()

On this page

# makeEllipse()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates an ellipse SVG path.

## Example [​](\#example "Direct link to Example")

```

ellipse.ts
tsx

import { makeEllipse } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeEllipse({
  rx: 100,
  ry: 50,
});

console.log(path); // M 100 0 a 100 100 0 1 0 1 0
console.log(width); // 200
console.log(height); // 100
console.log(transformOrigin); // '100 50'
```

```

ellipse.ts
tsx

import { makeEllipse } from "@remotion/shapes";

const { path, width, height, transformOrigin } = makeEllipse({
  rx: 100,
  ry: 50,
});

console.log(path); // M 100 0 a 100 100 0 1 0 1 0
console.log(width); // 200
console.log(height); // 100
console.log(transformOrigin); // '100 50'
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `rx`

_number_

The radius of the ellipse on the X axis.

### `ry`

_number_

The radius of the ellipse on the Y axis.

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the ellipse. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the ellipse. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Ellipse />`](/docs/shapes/ellipse).

## See also [​](\#see-also "Direct link to See also")

- [<Ellipse />](/docs/shapes/ellipse)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-ellipse.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-ellipse.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeCircle()](/docs/shapes/make-circle) [Next\
\
makeStar()](/docs/shapes/make-star)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [See also](#see-also)
