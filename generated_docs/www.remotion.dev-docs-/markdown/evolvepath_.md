---
title: evolvePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# evolvePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- evolvePath()

On this page

# evolvePath()

_Part of the [`@remotion/paths`](/docs/paths) package._

Animates an SVG path from being invisible to it's full length. The function takes two arguments:

- `progress` is the progress towards full evolution, where `0` means the path being invisible, and `1` meaning the path being fully drawn out.


note



Passing in a value above 1 will result in the start of the path getting devolved. Passing in a value below 0 will result in the path getting evolved from the end.

- `path` must be a valid SVG path.

The return value will be an object containing [`strokeDasharray`](https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dasharray) and [`strokeDashoffset`](https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/stroke-dashoffset) values, which should be passed to the `<path>` element.

```

tsx

import { evolvePath } from "@remotion/paths";

const path = "M 0 0 L 100 0";
const evolution = evolvePath(0.5, path);
console.log(evolution); // { strokeDasharray: '100 100',  strokeDashoffset: 50 }

const element = (
  <path
    d={path}
    strokeDasharray={evolution.strokeDasharray}
    strokeDashoffset={evolution.strokeDashoffset}
  />
);
```

```

tsx

import { evolvePath } from "@remotion/paths";

const path = "M 0 0 L 100 0";
const evolution = evolvePath(0.5, path);
console.log(evolution); // { strokeDasharray: '100 100',  strokeDashoffset: 50 }

const element = (
  <path
    d={path}
    strokeDasharray={evolution.strokeDasharray}
    strokeDashoffset={evolution.strokeDashoffset}
  />
);
```

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [reversePath()](/docs/paths/reverse-path)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/evolve-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/evolve-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
interpolatePath()](/docs/paths/interpolate-path) [Next\
\
resetPath()](/docs/paths/reset-path)

- [See also](#see-also)
