---
title: interpolateColors()
source: Remotion Documentation
last_updated: 2024-12-22
---

# interpolateColors()

- [Home page](/)
- [remotion](/docs/remotion)
- interpolateColors()

On this page

# interpolateColors()

_Available from v2.0.3_

Allows you to map a range of values to colors using a concise syntax.

## Reference [​](\#reference "Direct link to Reference")

### Params [​](\#params "Direct link to Params")

1. The input value.
2. The range of values that you expect the input to assume.
3. The range of output colors that you want the input to map to.

### Returns [​](\#returns "Direct link to Returns")

`interpolateColors()` returns a `rgba` color string. eg: `rgba(255, 100, 12, 1)`

## Example: interpolate colors [​](\#example-interpolate-colors "Direct link to Example: interpolate colors")

In this example, we are interpolating colors from red to yellow. At frame 0 (the start of the video), we want the color to be `red`. At frame 20, we want the color to be `yellow`.

Using the following snippet, we can calculate the current color for any frame:

```

tsx

import { interpolateColors, useCurrentFrame } from "remotion";

const frame = useCurrentFrame() / 10;

const color = interpolateColors(frame, [0, 20], ["red", "yellow"]); // rgba(255, 128, 0, 1)

const color2 = interpolateColors(frame, [0, 20], ["#ff0000", "#ffff00"]); // rgba(255, 128, 0, 1)
```

```

tsx

import { interpolateColors, useCurrentFrame } from "remotion";

const frame = useCurrentFrame() / 10;

const color = interpolateColors(frame, [0, 20], ["red", "yellow"]); // rgba(255, 128, 0, 1)

const color2 = interpolateColors(frame, [0, 20], ["#ff0000", "#ffff00"]); // rgba(255, 128, 0, 1)
```

## Example: interpolate `rgb` or `rgba` colors [​](\#example-interpolate-rgb-or-rgba-colors "Direct link to example-interpolate-rgb-or-rgba-colors")

In this example, we are interpolating colors from red to yellow. At frame 0 (the start of the video), we want the color to be `red` ( `rgb(255, 0, 0)`). At frame 20, we want the color to be `yellow` ( `rgba(255, 255, 0)`).

Using the following snippet, we can calculate the current color for any frame:

```

tsx

import { interpolateColors, useCurrentFrame } from "remotion";

const frame = useCurrentFrame(); // 10

// RGB colors
const color = interpolateColors(
  frame,
  [0, 20],
  ["rgb(255, 0, 0)", "rgb(255, 255, 0)"]
); // rgba(255, 128, 0, 1)

// RGBA colors
const color2 = interpolateColors(
  frame,
  [0, 20],
  ["rgba(255, 0, 0, 1)", "rgba(255, 255, 0, 0)"]
); // rgba(255, 128, 0, 0.5)
```

```

tsx

import { interpolateColors, useCurrentFrame } from "remotion";

const frame = useCurrentFrame(); // 10

// RGB colors
const color = interpolateColors(
  frame,
  [0, 20],
  ["rgb(255, 0, 0)", "rgb(255, 255, 0)"]
); // rgba(255, 128, 0, 1)

// RGBA colors
const color2 = interpolateColors(
  frame,
  [0, 20],
  ["rgba(255, 0, 0, 1)", "rgba(255, 255, 0, 0)"]
); // rgba(255, 128, 0, 0.5)
```

## Example: interpolate `hsl` or `hsla` colors [​](\#example-interpolate-hsl-or-hsla-colors "Direct link to example-interpolate-hsl-or-hsla-colors")

In this example, we are interpolating colors from red to yellow. At frame 0 (the start of the video), we want the color to be `red` ( `hsl(0, 100%, 50%)`). At frame 20, we want the color to be `yellow` ( `hsl(60, 100%, 50%)`).

Using the following snippet, we can calculate the current color for any frame:

```

ts

import { useCurrentFrame, interpolateColors } from "remotion";

const frame = useCurrentFrame(); // 10
//hsl example
const color = interpolateColors(
  frame,
  [0, 20],
  ["hsl(0, 100%, 50%)", "hsl(60, 100%, 50%)"]
); // rgba(255, 128, 0, 1)

//hsla example
const color2 = interpolateColors(
  frame,
  [0, 20],
  ["hsla(0, 100%, 50%, 1)", "hsla(60, 100%, 50%, 1)"]
); // rgba(255, 128, 0, 1)
```

```

ts

import { useCurrentFrame, interpolateColors } from "remotion";

const frame = useCurrentFrame(); // 10
//hsl example
const color = interpolateColors(
  frame,
  [0, 20],
  ["hsl(0, 100%, 50%)", "hsl(60, 100%, 50%)"]
); // rgba(255, 128, 0, 1)

//hsla example
const color2 = interpolateColors(
  frame,
  [0, 20],
  ["hsla(0, 100%, 50%, 1)", "hsla(60, 100%, 50%, 1)"]
); // rgba(255, 128, 0, 1)
```

## Example: interpolate color names [​](\#example-interpolate-color-names "Direct link to Example: interpolate color names")

Interpolating CSS color names is also supported.

```

ts

import { useCurrentFrame, interpolateColors } from "remotion";

const frame = useCurrentFrame(); // 10

const color = interpolateColors(frame, [0, 20], ["red", "yellow"]); // rgba(255, 128, 0, 1)
```

```

ts

import { useCurrentFrame, interpolateColors } from "remotion";

const frame = useCurrentFrame(); // 10

const color = interpolateColors(frame, [0, 20], ["red", "yellow"]); // rgba(255, 128, 0, 1)
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/interpolate-colors.ts)
- [interpolate()](/docs/interpolate)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/interpolate-colors.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Img>](/docs/img) [Next\
\
interpolate()](/docs/interpolate)

- [Reference](#reference)
  - [Params](#params)
  - [Returns](#returns)
- [Example: interpolate colors](#example-interpolate-colors)
- [Example: interpolate `rgb` or `rgba` colors](#example-interpolate-rgb-or-rgba-colors)
- [Example: interpolate `hsl` or `hsla` colors](#example-interpolate-hsl-or-hsla-colors)
- [Example: interpolate color names](#example-interpolate-color-names)
- [See also](#see-also)
