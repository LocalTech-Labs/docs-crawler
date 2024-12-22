---
title: makeStar()
source: Remotion Documentation
last_updated: 2024-12-22
---

# makeStar()

- [Home page](/)
- [@remotion/shapes](/docs/shapes/)
- makeStar()

On this page

# makeStar()

_Part of the [`@remotion/shapes`](/docs/shapes) package._

Generates an star SVG path.

## Example [​](\#example "Direct link to Example")

```

star.ts
tsx

import { makeStar } from "@remotion/shapes";

const { path, width, height, transformOrigin, instructions } = makeStar({
  innerRadius: 200,
  outerRadius: 150,
  points: 5,
});

console.log(path); // M 200 0 L 288.167787843871 78.64745084375788 L 390.21130325903073 138.19660112501052 L 342.658477444273 246.3525491562421 L 317.55705045849464 361.8033988749895 L 200 350 L 82.4429495415054 361.8033988749895 L 57.34152255572698 246.35254915624213 L 9.788696740969272 138.19660112501055 L 111.83221215612902 78.6474508437579 L 200 0
console.log(width); // 400
console.log(height); // 400
console.log(transformOrigin); // '200 200'
console.log(instructions); // '[{type: "M"}, ...]'
```

```

star.ts
tsx

import { makeStar } from "@remotion/shapes";

const { path, width, height, transformOrigin, instructions } = makeStar({
  innerRadius: 200,
  outerRadius: 150,
  points: 5,
});

console.log(path); // M 200 0 L 288.167787843871 78.64745084375788 L 390.21130325903073 138.19660112501052 L 342.658477444273 246.3525491562421 L 317.55705045849464 361.8033988749895 L 200 350 L 82.4429495415054 361.8033988749895 L 57.34152255572698 246.35254915624213 L 9.788696740969272 138.19660112501055 L 111.83221215612902 78.6474508437579 L 200 0
console.log(width); // 400
console.log(height); // 400
console.log(transformOrigin); // '200 200'
console.log(instructions); // '[{type: "M"}, ...]'
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `points`

_number_

The amount of points of the star.

### `innerRadius`

_number_

The inner radius of the star.

### `outerRadius`

_number_

The outer radius of the star.

## Return type [​](\#return-type "Direct link to Return type")

### `path`

A string that is suitable as an argument for `d` in a `<path>` element.

### `width`

The width of the star. Suitable for defining the `viewBox` of an `<svg>` tag.

### `height`

The height of the star. Suitable for defining the `viewBox` of an `<svg>` tag.

### `instructions`

An array with SVG instructions. The type for a instruction can be seen by importing `Instruction` from `@remotion/shapes`.

### `transformOrigin`

A string representing the point of origin if a shape should be rotated around itself.

If you want to rotate the shape around its center, use the `transform-origin` CSS property and pass this value, and also add `transform-box: fill-box`. This is the default for [`<Star />`](/docs/shapes/star).

## See also [​](\#see-also "Direct link to See also")

- [<Star />](/docs/shapes/star)
- [`@remotion/shapes`](/docs/shapes)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/shapes/src/utils/make-star.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/shapes/make-star.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeEllipse()](/docs/shapes/make-ellipse) [Next\
\
makePie()](/docs/shapes/make-pie)

- [Example](#example)
- [Arguments](#arguments)
- [Return type](#return-type)
- [See also](#see-also)
