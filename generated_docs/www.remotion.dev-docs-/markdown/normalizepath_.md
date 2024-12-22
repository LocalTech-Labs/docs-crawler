---
title: normalizePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# normalizePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- normalizePath()

On this page

# normalizePath()

_Part of the [`@remotion/paths`](/docs/paths) package._

Removes all relative coordinates from a path and converts them into absolute coordinates.

Returns a string if the path is valid:

```

tsx

import { normalizePath } from "@remotion/paths";

const normalizedPath = normalizePath("M 50 50 l 100 0");
console.log(normalizedPath); // "M 50 50 L 150 50"
```

```

tsx

import { normalizePath } from "@remotion/paths";

const normalizedPath = normalizePath("M 50 50 l 100 0");
console.log(normalizedPath); // "M 50 50 L 150 50"
```

The function will throw if the path is invalid:

```

tsx

normalizePath("remotion"); // Error: Malformed path data: ...
```

```

tsx

normalizePath("remotion"); // Error: Malformed path data: ...
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [svg-path-reverse](https://www.npmjs.com/package/svg-path-reverse).

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [`reduceInstructions()`](/docs/paths/reduce-instructions)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/normalize-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/normalize-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
reversePath()](/docs/paths/reverse-path) [Next\
\
interpolatePath()](/docs/paths/interpolate-path)

- [Credits](#credits)
- [See also](#see-also)
