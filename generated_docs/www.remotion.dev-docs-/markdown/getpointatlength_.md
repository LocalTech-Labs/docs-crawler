---
title: getPointAtLength()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getPointAtLength()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- getPointAtLength()

On this page

# getPointAtLength()

_Part of the [`@remotion/paths`](/docs/paths) package._

Gets the coordinates of a point which is on an SVG path.
The first argument is an SVG path, the second one is the length at which the point should be sampled. It must be between `0` and the return value of [`getLength()`](/docs/paths/get-length).

An object containing `x` and `y` is returned if the path is valid:

```

tsx

import { getPointAtLength } from "@remotion/paths";

const point = getPointAtLength("M 0 0 L 100 0", 50);
console.log(point); // { x: 50, y: 0 }
```

```

tsx

import { getPointAtLength } from "@remotion/paths";

const point = getPointAtLength("M 0 0 L 100 0", 50);
console.log(point); // { x: 50, y: 0 }
```

The function will throw if the path is invalid:

```

tsx

getPointAtLength("remotion", 50); // Error: Malformed path data: ...
```

```

tsx

getPointAtLength("remotion", 50); // Error: Malformed path data: ...
```

## Example: Getting the middle point of a path [​](\#example-getting-the-middle-point-of-a-path "Direct link to Example: Getting the middle point of a path")

Use [`getLength()`](/docs/paths/get-length) to get the total length of a path and then multiply it with a number between 0 and 1 to get any point on the path. For example, `length * 0.5` to get the coordinates in the middle of the path.

```

tsx

import { getLength, getPointAtLength } from "@remotion/paths";

const path = "M 0 0 L 100 0";
const length = getLength(path);
const point = getPointAtLength(path, length * 0.5);

console.log(point); // { x: 50, y: 0 }
```

```

tsx

import { getLength, getPointAtLength } from "@remotion/paths";

const path = "M 0 0 L 100 0";
const length = getLength(path);
const point = getPointAtLength(path, length * 0.5);

console.log(point); // { x: 50, y: 0 }
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [svg-path-properties](https://www.npmjs.com/package/svg-path-properties).

## See also [​](\#see-also "Direct link to See also")

- [getLength()](/docs/paths/get-length)
- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/get-point-at-length.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/get-point-at-length.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getLength()](/docs/paths/get-length) [Next\
\
getTangentAtLength()](/docs/paths/get-tangent-at-length)

- [Example: Getting the middle point of a path](#example-getting-the-middle-point-of-a-path)
- [Credits](#credits)
- [See also](#see-also)
