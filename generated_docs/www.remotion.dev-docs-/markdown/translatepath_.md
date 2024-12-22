---
title: translatePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# translatePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- translatePath()

On this page

# translatePath()

_Part of the [`@remotion/paths`](/docs/paths) package._

Translates the path by the given `x` and `y` coordinates.

## Arguments [​](\#arguments "Direct link to Arguments")

The function takes three arguments:

- `path`, the original SVG path.
- `x`, the amount of horizontal translation.
- `y` the amount of vertical translation.

## Return value [​](\#return-value "Direct link to Return value")

Returns a new `string` containing a path if it is valid:

```

translate-x.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath("M 50 50 L 150 50", 10, 0);
console.log(translatedPath); // "M 50 50 L 150 50"
```

```

translate-x.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath("M 50 50 L 150 50", 10, 0);
console.log(translatedPath); // "M 50 50 L 150 50"
```

```

translate-y.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath("M10 10 L15 15", 10, 10);
console.log(translatedPath); // "M 20 20 L 25 25"
```

```

translate-y.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath("M10 10 L15 15", 10, 10);
console.log(translatedPath); // "M 20 20 L 25 25"
```

```

translate-x-and-y.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath(
  "M 35,50 a 25,25,0,1,1,50,0 a 25,25,0,1,1,-50,0",
  10,
  20
);
console.log(translatedPath); // "M 45 70 a 25 25 0 1 1 50 0 a 25, 5 0 1 1 -50 0"
```

```

translate-x-and-y.ts
tsx

import { translatePath } from "@remotion/paths";

const translatedPath = translatePath(
  "M 35,50 a 25,25,0,1,1,50,0 a 25,25,0,1,1,-50,0",
  10,
  20
);
console.log(translatedPath); // "M 45 70 a 25 25 0 1 1 50 0 a 25, 5 0 1 1 -50 0"
```

The function will throw if the path is invalid:

```

tsx

translatePath("remotion", 10, 0); // Malformed path data: "m" ...
```

```

tsx

translatePath("remotion", 10, 0); // Malformed path data: "m" ...
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [translate-svg-path](https://github.com/michaelrhodes/translate-svg-path) and [serialize-svg-path](https://github.com/jkroso/serialize-svg-path).

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/translate-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/translate-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getSubpaths()](/docs/paths/get-subpaths) [Next\
\
warpPath()](/docs/paths/warp-path)

- [Arguments](#arguments)
- [Return value](#return-value)
- [Credits](#credits)
- [See also](#see-also)
