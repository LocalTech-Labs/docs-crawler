---
title: extendViewBox()
source: Remotion Documentation
last_updated: 2024-12-22
---

# extendViewBox()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- extendViewBox()

On this page

# extendViewBox()

_Part of the [`@remotion/paths`](/docs/paths) package. Available since v3.2.25_

Widens an SVG `viewBox` in all directions by a certain scale factor.

note

This function may be unnecessary: If you want the parts that go outside of the viewbox to be visible, you can also set `style={{overflow: 'visible'}}` on the SVG container.

```

tsx

import { extendViewBox } from "@remotion/paths";

const extended = extendViewBox("0 0 1000 1000", 2);
console.log(extended); // "-500 -500 2000 2000"
```

```

tsx

import { extendViewBox } from "@remotion/paths";

const extended = extendViewBox("0 0 1000 1000", 2);
console.log(extended); // "-500 -500 2000 2000"
```

The function will throw if the viewBox is invalid.

## Example: Displaying an SVG path that goes out of bounds [​](\#example-displaying-an-svg-path-that-goes-out-of-bounds "Direct link to Example: Displaying an SVG path that goes out of bounds")

Consider the following SVG:

The path will go from `0` to `1500` on the horizontal axis, but it will be cut off because it goes beyond the viewport area.

```

tsx

const viewBox = "0 0 1000 1000";

export const ViewBoxExample: React.FC = () => {
  return (
    <svg viewBox={viewBox}>
      <path d={"0 500 1500 500"} stroke="black" strokeWidth={4} />
    </svg>
  );
};
```

```

tsx

const viewBox = "0 0 1000 1000";

export const ViewBoxExample: React.FC = () => {
  return (
    <svg viewBox={viewBox}>
      <path d={"0 500 1500 500"} stroke="black" strokeWidth={4} />
    </svg>
  );
};
```

We can fix the cutoff by doing two things:

- Scaling the viewBox by a factor of 2
- Applying a 2x scale transform to the SVG.

```

tsx

import { extendViewBox } from "@remotion/paths";

const viewBox = "0 0 1000 1000";

export const ViewBoxExample: React.FC = () => {
  return (
    <svg style={{ scale: "2" }} viewBox={extendViewBox(viewBox, 2)}>
      <path d={"0 500 1500 500"} stroke="black" strokeWidth={4} />
    </svg>
  );
};
```

```

tsx

import { extendViewBox } from "@remotion/paths";

const viewBox = "0 0 1000 1000";

export const ViewBoxExample: React.FC = () => {
  return (
    <svg style={{ scale: "2" }} viewBox={extendViewBox(viewBox, 2)}>
      <path d={"0 500 1500 500"} stroke="black" strokeWidth={4} />
    </svg>
  );
};
```

By doing that, the each dimensions of the viewBox will be doubled, which will result in the picture being scaled down. By applying a scale transform, this can be corrected.

In this example, a factor of `2` was chosen because it is enough to fix the cutoff problem. The more the SVG path goes outside the container, the higher the factor needs to be to compensate.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/extend-viewbox.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/extend-viewbox.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getBoundingBox()](/docs/paths/get-bounding-box) [Next\
\
parsePath()](/docs/paths/parse-path)

- [Example: Displaying an SVG path that goes out of bounds](#example-displaying-an-svg-path-that-goes-out-of-bounds)
- [See also](#see-also)
