---
title: parsePath()
source: Remotion Documentation
last_updated: 2024-12-22
---

# parsePath()

- [Home page](/)
- [@remotion/paths](/docs/paths/)
- parsePath()

On this page

# parsePath()

_Part of the [`@remotion/paths`](/docs/paths) package. Available from v3.3.40_

Parses an SVG string path into an array of `Instruction`'s.

```

reset-path.ts
tsx

import { parsePath } from "@remotion/paths";

const newPath = parsePath("M 10 10 L 20 20");

/*
  [
    { type: "M", x: 10, y: 10 },
    { type: "L", x: 20, y: 20 },
  ]
*/
```

```

reset-path.ts
tsx

import { parsePath } from "@remotion/paths";

const newPath = parsePath("M 10 10 L 20 20");

/*
  [
    { type: "M", x: 10, y: 10 },
    { type: "L", x: 20, y: 20 },
  ]
*/
```

This function will throw if the SVG path is invalid.

## Return type type [​](\#return-type-type "Direct link to Return type type")

An array of `Instruction`'s. The `Instruction` type can also be imported from `@remotion/paths`:

```

ts

import type { Instruction } from "@remotion/paths";
```

```

ts

import type { Instruction } from "@remotion/paths";
```

The type has the following shape:

```

ts

export type Instruction =
  | {
      type: "M";
      x: number;
      y: number;
    }
  | {
      type: "L";
      x: number;
      y: number;
    }
  | {
      type: "H";
      x: number;
    }
  | {
      type: "V";
      y: number;
    }
  | {
      type: "C";
      cp1x: number;
      cp1y: number;
      cp2x: number;
      cp2y: number;
      x: number;
      y: number;
    }
  | {
      type: "S";
      cpx: number;
      cpy: number;
      x: number;
      y: number;
    }
  | {
      type: "Q";
      cpx: number;
      cpy: number;
      x: number;
      y: number;
    }
  | {
      type: "T";
      x: number;
      y: number;
    }
  | {
      type: "A";
      rx: number;
      ry: number;
      xAxisRotation: number;
      largeArcFlag: boolean;
      sweepFlag: boolean;
      x: number;
      y: number;
    }
  | {
      type: "m";
      dx: number;
      dy: number;
    }
  | {
      type: "l";
      dx: number;
      dy: number;
    }
  | {
      type: "h";
      dx: number;
    }
  | {
      type: "v";
      dy: number;
    }
  | {
      type: "c";
      cp1dx: number;
      cp1dy: number;
      cp2dx: number;
      cp2dy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "s";
      cpdx: number;
      cpdy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "q";
      cpdx: number;
      cpdy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "t";
      dx: number;
      dy: number;
    }
  | {
      type: "a";
      rx: number;
      ry: number;
      xAxisRotation: number;
      largeArcFlag: boolean;
      sweepFlag: boolean;
      dx: number;
      dy: number;
    }
  | {
      type: "Z";
    }
  | {
      type: "z";
    };
```

```

ts

export type Instruction =
  | {
      type: "M";
      x: number;
      y: number;
    }
  | {
      type: "L";
      x: number;
      y: number;
    }
  | {
      type: "H";
      x: number;
    }
  | {
      type: "V";
      y: number;
    }
  | {
      type: "C";
      cp1x: number;
      cp1y: number;
      cp2x: number;
      cp2y: number;
      x: number;
      y: number;
    }
  | {
      type: "S";
      cpx: number;
      cpy: number;
      x: number;
      y: number;
    }
  | {
      type: "Q";
      cpx: number;
      cpy: number;
      x: number;
      y: number;
    }
  | {
      type: "T";
      x: number;
      y: number;
    }
  | {
      type: "A";
      rx: number;
      ry: number;
      xAxisRotation: number;
      largeArcFlag: boolean;
      sweepFlag: boolean;
      x: number;
      y: number;
    }
  | {
      type: "m";
      dx: number;
      dy: number;
    }
  | {
      type: "l";
      dx: number;
      dy: number;
    }
  | {
      type: "h";
      dx: number;
    }
  | {
      type: "v";
      dy: number;
    }
  | {
      type: "c";
      cp1dx: number;
      cp1dy: number;
      cp2dx: number;
      cp2dy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "s";
      cpdx: number;
      cpdy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "q";
      cpdx: number;
      cpdy: number;
      dx: number;
      dy: number;
    }
  | {
      type: "t";
      dx: number;
      dy: number;
    }
  | {
      type: "a";
      rx: number;
      ry: number;
      xAxisRotation: number;
      largeArcFlag: boolean;
      sweepFlag: boolean;
      dx: number;
      dy: number;
    }
  | {
      type: "Z";
    }
  | {
      type: "z";
    };
```

## See also [​](\#see-also "Direct link to See also")

- [`serializeInstructions()`](/docs/paths/serialize-instructions)
- [`@remotion/paths`](/docs/paths)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/paths/src/parse-path.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/paths/parse-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
extendViewBox()](/docs/paths/extend-viewbox) [Next\
\
serializeInstructions()](/docs/paths/serialize-instructions)

- [Return type type](#return-type-type)
- [See also](#see-also)
