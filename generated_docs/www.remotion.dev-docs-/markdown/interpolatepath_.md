---
title: interpolatePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# interpolatePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- interpolatePath()

On this page

# interpolatePath()

_Part of the [`@remotion/paths`](/docs/paths) package._

Interpolates between two SVG paths. The function takes three arguments:

- `value`, which is a number.
  - If it is `0`, the first path is returned.
  - If it is `1`, the second path is returned.
  - If it is inbetween or outside the range, the path is interpolated.
- `firstPath`, which must be a valid SVG path.
- `secondPath`, which must be a valid SVG path.

```

tsx

import { interpolatePath } from "@remotion/paths";

const interpolated = interpolatePath(0.5, "M 0 0 L 100 0", "M 100 0 L 0 0");
console.log(interpolated); // "M 50 0 L 50 0"
```

```

tsx

import { interpolatePath } from "@remotion/paths";

const interpolated = interpolatePath(0.5, "M 0 0 L 100 0", "M 100 0 L 0 0");
console.log(interpolated); // "M 50 0 L 50 0"
```

## Credits [​](\#credits "Direct link to Credits")

Source code stems mostly from [d3-interpolate-path](https://github.com/pbeshai/d3-interpolate-path).

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/interpolate-path/interpolate-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/interpolate-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
normalizePath()](/docs/paths/normalize-path) [Next\
\
evolvePath()](/docs/paths/evolve-path)

- [Credits](#credits)
- [See also](#see-also)
