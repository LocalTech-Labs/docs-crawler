---
title: noise4D()
source: Remotion Documentation
last_updated: 2024-12-22
---

# noise4D()

- [Home page](/)
- [@remotion/noise](/docs/noise/)
- noise4D()

On this page

# noise4D()

_Part of the [`@remotion/noise`](/docs/noise) package._

Creates 4D noise.

## API [​](\#api "Direct link to API")

The function takes five arguments:

### `seed` [​](\#seed "Direct link to seed")

Pass any _string_ or _number_. If the seed is the same, you will get the same result for same `x`, `y`, `z` and `w` values. Change the seed to get different results for your `x`, `y`, `z`, `w` values.

### `x` [​](\#x "Direct link to x")

_number_

The first dimensional value.

### `y` [​](\#y "Direct link to y")

_number_

The second dimensional value.

### `z` [​](\#z "Direct link to z")

_number_

The third dimensional value.

### `w` [​](\#w "Direct link to w")

_number_

The fourth dimensional value.

## Return value [​](\#return-value "Direct link to Return value")

A value between `-1` and `1`, swinging as your `x`, `y`, `z` and `w` values change.

## Example [​](\#example "Direct link to Example")

```

tsx

import { noise4D } from "@remotion/noise";

const x = 32;
const y = 40;
const z = 50;
const w = 64;
console.log(noise4D("my-seed", x, y, z, w));
```

```

tsx

import { noise4D } from "@remotion/noise";

const x = 32;
const y = 40;
const z = 50;
const w = 64;
console.log(noise4D("my-seed", x, y, z, w));
```

## Credits [​](\#credits "Direct link to Credits")

Dependency: [simplex-noise](https://www.npmjs.com/package/simplex-noise)

## See also [​](\#see-also "Direct link to See also")

- [Example: Noise visualization](/docs/noise-visualization)
- [noise2D()](/docs/noise/noise-2d)
- [noise3D()](/docs/noise/noise-3d)
- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/noise/src/index.ts)
- [`@remotion/noise`](/docs/noise)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/noise/noise-4d.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
noise3D()](/docs/noise/noise-3d) [Next\
\
Overview](/docs/google-fonts/)

- [API](#api)
  - [`seed`](#seed)
  - [`x`](#x)
  - [`y`](#y)
  - [`z`](#z)
  - [`w`](#w)
- [Return value](#return-value)
- [Example](#example)
- [Credits](#credits)
- [See also](#see-also)
