---
title: scalePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# scalePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- scalePath()

On this page

# scalePath()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.43_

Allows you to grow or shrink the size of a path.

```

scale-path.ts
tsx

import { scalePath } from "@remotion/paths";

const newPath = scalePath("M 0 0 L 100 100", 1, 2); // "M 0 0 L 100 200";
```

```

scale-path.ts
tsx

import { scalePath } from "@remotion/paths";

const newPath = scalePath("M 0 0 L 100 100", 1, 2); // "M 0 0 L 100 200";
```

The origin of the transform is the top left corner of the path. To use a different origin, first use [`translatePath()`](/docs/paths/translate-path) to move the path to the desired origin, then scale it, and finally move it back to the original origin.

## Arguments [​](\#arguments "Direct link to Arguments")

### `path` [​](\#path "Direct link to path")

_string_

A valid SVG Path string.

### `xScale` [​](\#xscale "Direct link to xscale")

_number_

The factor of which to scale the path horizontally. `1` will leave the path unchanged.

### `yScale` [​](\#yscale "Direct link to yscale")

_number_

The factor of which to scale the path vertically. `1` will leave the path unchanged.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/scale-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/scale-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
warpPath()](/docs/paths/warp-path) [Next\
\
getBoundingBox()](/docs/paths/get-bounding-box)

- [Arguments](#arguments)
  - [`path`](#path)
  - [`xScale`](#xscale)
  - [`yScale`](#yscale)
- [See also](#see-also)
