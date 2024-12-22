---
title: serializeInstructions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# serializeInstructions()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- serializeInstructions()

On this page

# serializeInstructions()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.40_

Takes an array of [`Instruction`](/docs/paths/parse-path)'s and serializes it into an SVG path string.

```

serialize-instructions.ts
tsx

import { serializeInstructions } from "@remotion/paths";

const newPath = serializeInstructions([
  {
    type: "M",
    x: 10,
    y: 10,
  },
  {
    type: "L",
    x: 20,
    y: 20,
  },
]); // M 10 10 L 20 20
```

```

serialize-instructions.ts
tsx

import { serializeInstructions } from "@remotion/paths";

const newPath = serializeInstructions([
  {
    type: "M",
    x: 10,
    y: 10,
  },
  {
    type: "L",
    x: 20,
    y: 20,
  },
]); // M 10 10 L 20 20
```

This function may throw if the instructions don't match the [`Instruction`](/docs/paths/parse-path) type, but it does not explicitly check for invalid input.

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [`parsePath()`](/docs/paths/parse-path)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/serialize-instructions.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/serialize-instructions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
parsePath()](/docs/paths/parse-path) [Next\
\
reduceInstructions()](/docs/paths/reduce-instructions)

- [See also](#see-also)
