---
title: reduceInstructions()
source: Remotion Documentation
last_updated: 2024-12-22
---

# reduceInstructions()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- reduceInstructions()

On this page

# reduceInstructions()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.40_

Takes an array of [`Instruction`](/docs/paths/parse-path)'s and reduces the amount of instruction types them so the path only consists of `M`, `L`, `C` and `Z` instructions.

note

In versions before v4.0.168, `Q` instructions were reduced away.

In reverse logic, this function will eliminate all `H`, `V`, `S`, `T`, `A`, `Q` `m`, `l`, `h`, `v`, `c`, `s`, `q`, `t`, `a`, `q` and `z` instructions.

This is useful if you want to manually edit a path and want to make sure it's as simple as possible.

Note that this may result in a longer path.

```

ts

import { reduceInstructions, ReducedInstruction } from "@remotion/paths";

const simplified: ReducedInstruction[] = reduceInstructions([
  { type: "m", dx: 10, dy: 10 },
  { type: "h", dx: 100 },
]);

/*
  [
    {type: 'M', x: 10, y: 10},
    {type: 'L', x: 110, y: 10},
  ]
*/
```

```

ts

import { reduceInstructions, ReducedInstruction } from "@remotion/paths";

const simplified: ReducedInstruction[] = reduceInstructions([
  { type: "m", dx: 10, dy: 10 },
  { type: "h", dx: 100 },
]);

/*
  [
    {type: 'M', x: 10, y: 10},
    {type: 'L', x: 110, y: 10},
  ]
*/
```

## `ReducedInstruction` type [​](\#reducedinstruction-type "Direct link to reducedinstruction-type")

If you want a type which includes only reduced instructions, you can import the `ReducedInstruction` type.

```

ts

import { ReducedInstruction } from "@remotion/paths";
```

```

ts

import { ReducedInstruction } from "@remotion/paths";
```

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/paths`](/docs/paths)
- [`parsePath()`](/docs/paths/parse-path)
- [`normalizePath()`](/docs/paths/normalize-path)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/reduce-instructions.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/reduce-instructions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
serializeInstructions()](/docs/paths/serialize-instructions) [Next\
\
getParts()](/docs/paths/get-parts)

- [`ReducedInstruction` type](#reducedinstruction-type)
- [See also](#see-also)
