---
title: resetPath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# resetPath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- resetPath()

On this page

# resetPath()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.40_

Translates an SVG path so that the top-left corner of the [bounding box](/docs/paths/get-bounding-box) is at `0, 0`. Useful for simplifying the math when transforming the coordinates of an SVG path.

```

reset-path.ts
tsx

import { resetPath } from "@remotion/paths";

const newPath = resetPath("M 10 10 L 20 20"); // M 0 0 L 10 10
```

```

reset-path.ts
tsx

import { resetPath } from "@remotion/paths";

const newPath = resetPath("M 10 10 L 20 20"); // M 0 0 L 10 10
```

This function will throw if the SVG path is invalid.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/reset-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/reset-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
evolvePath()](/docs/paths/evolve-path) [Next\
\
getSubpaths()](/docs/paths/get-subpaths)

- [See also](#see-also)
