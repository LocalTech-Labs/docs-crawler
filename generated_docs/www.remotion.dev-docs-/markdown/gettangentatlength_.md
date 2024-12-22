---
title: getTangentAtLength()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getTangentAtLength()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- getTangentAtLength()

On this page

# getTangentAtLength()

_Part of the [`@remotion/paths`](/docs/paths) package._

Gets tangent values `x` and `y` of a point which is on an SVG path. The first argument is an SVG path, the second one is the length at which the point should be sampled. It must be between 0 and the return value of [`getLength()`](/docs/paths/get-length).

Returns a point if the path is valid:

```

tsx

import { getTangentAtLength } from "@remotion/paths";

const tangent = getTangentAtLength("M 50 50 L 150 50", 50);
console.log(tangent); // { x: 1, y: 0}
```

```

tsx

import { getTangentAtLength } from "@remotion/paths";

const tangent = getTangentAtLength("M 50 50 L 150 50", 50);
console.log(tangent); // { x: 1, y: 0}
```

The function will throw if the path is invalid:

```

tsx

getTangentAtLength("remotion", 50); // Error: Malformed path data: ...
```

```

tsx

getTangentAtLength("remotion", 50); // Error: Malformed path data: ...
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [svg-path-properties](https://www.npmjs.com/package/svg-path-properties).

## See also [​](\#see-also "Direct link to See also")

- [getLength()](/docs/paths/get-length)
- [getPointAtLength()](/docs/paths/get-point-at-length)
- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/get-tangent-at-length.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/get-tangent-at-length.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getPointAtLength()](/docs/paths/get-point-at-length) [Next\
\
getInstructionIndexAtLength()](/docs/paths/get-instruction-index-at-length)

- [Credits](#credits)
- [See also](#see-also)
