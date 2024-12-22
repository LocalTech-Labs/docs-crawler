---
title: getBoundingBox()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getBoundingBox()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- getBoundingBox()

On this page

# getBoundingBox()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.40_

Returns the bounding box of the given path, suitable for calculating the `viewBox` value that you need to pass to an SVG.

The bounding box is the smallest rectangle which can contain the shape in full.

```

get-bounding-box.ts
tsx

import { getBoundingBox } from "@remotion/paths";

const boundingBox = getBoundingBox(
  "M 35,50 a 25,25,0,1,1,50,0 a 25,25,0,1,1,-50,0"
);
// { x1: 35, x2: 85, y1: 24.999999999999993, y2: 75 };
```

```

get-bounding-box.ts
tsx

import { getBoundingBox } from "@remotion/paths";

const boundingBox = getBoundingBox(
  "M 35,50 a 25,25,0,1,1,50,0 a 25,25,0,1,1,-50,0"
);
// { x1: 35, x2: 85, y1: 24.999999999999993, y2: 75 };
```

This function will throw if the SVG path is invalid.

## Return type [​](\#return-type "Direct link to Return type")

Includes the following properties:

- `x1`: The leftmost x coordinate of the bounding box
- `x2`: The rightmost x coordinate of the bounding box
- `y1`: The topmost y coordinate of the bounding box
- `y2`: The bottommost y coordinate of the bounding box
- `width`: The width of the bounding box, _returned from v3.3.97_
- `height`: The height of the bounding box, _returned from v3.3.97_
- `viewBox`: The `viewBox` value that you can pass to an SVG, _returned from v3.3.97_

## `BoundingBox` type [​](\#boundingbox-type "Direct link to boundingbox-type")

In TypeScript, you can get the shape of the return value by importing the `BoundingBox` type:

```

ts

import type { BoundingBox } from "@remotion/paths";
```

```

ts

import type { BoundingBox } from "@remotion/paths";
```

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/get-bounding-box.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/get-bounding-box.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
scalePath()](/docs/paths/scale-path) [Next\
\
extendViewBox()](/docs/paths/extend-viewbox)

- [Return type](#return-type)
- [`BoundingBox` type](#boundingbox-type)
- [See also](#see-also)
