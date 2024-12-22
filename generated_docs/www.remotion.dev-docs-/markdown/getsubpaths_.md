---
title: getSubpaths()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getSubpaths()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- getSubpaths()

On this page

# getSubpaths()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.6_

Takes an SVG path and returns an array of subpaths.

Each `M` and `m` statement in a path creates a new subpath.

Example of a path that has two straight lines:

```

tsx

import { getSubpaths } from "@remotion/paths";

const parts = getSubpaths(`
  M 0 0 L 100 0
  M 0 100 L 200 100
`);
```

```

tsx

import { getSubpaths } from "@remotion/paths";

const parts = getSubpaths(`
  M 0 0 L 100 0
  M 0 100 L 200 100
`);
```

An array is returned containing two parts.

```

tsx

console.log(parts[0]); // "M 0 0 L 100 0"
console.log(parts[1]); // "M 0 100 L 200 100"
```

```

tsx

console.log(parts[0]); // "M 0 0 L 100 0"
console.log(parts[1]); // "M 0 100 L 200 100"
```

Paths containing relative `m` elements will be converted into `M` elements.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/get-subpaths.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/get-subpaths.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
resetPath()](/docs/paths/reset-path) [Next\
\
translatePath()](/docs/paths/translate-path)

- [See also](#see-also)
