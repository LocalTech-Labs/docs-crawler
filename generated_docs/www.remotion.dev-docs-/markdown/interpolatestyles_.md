---
title: interpolateStyles()
source: Remotion Documentation
last_updated: 2024-12-22
---

# interpolateStyles()

- [Home page](/)
- [@remotion/animation-utils](/docs/animation-utils/)
- interpolateStyles()

On this page

# interpolateStyles()

_Part of the [`@remotion/animation-utils`](/docs/animation-utils) package._

This function provides a convenient way to interpolate styles based on a specified range of values, allowing for smooth animations between different styles.

## Example [​](\#example "Direct link to Example")

```

tsx

import {
  interpolateStyles,
  makeTransform,
  translateY,
} from "@remotion/animation-utils";

const MyComponent: React.FC = () => {
  const animatedStyles = interpolateStyles(
    15,
    [0, 30, 60],
    [
      { opacity: 0, transform: makeTransform([translateY(-50)]) },
      { opacity: 1, transform: makeTransform([translateY(0)]) },
      { opacity: 0, transform: makeTransform([translateY(50)]) },
    ],
  );

  return <div style={animatedStyles} />;
};
```

```

tsx

import {
  interpolateStyles,
  makeTransform,
  translateY,
} from "@remotion/animation-utils";

const MyComponent: React.FC = () => {
  const animatedStyles = interpolateStyles(
    15,
    [0, 30, 60],
    [
      { opacity: 0, transform: makeTransform([translateY(-50)]) },
      { opacity: 1, transform: makeTransform([translateY(0)]) },
      { opacity: 0, transform: makeTransform([translateY(50)]) },
    ],
  );

  return <div style={animatedStyles} />;
};
```

## API [​](\#api "Direct link to API")

A function that takes 3-4 arguments:

1. The input value.
2. The range of values that you expect the input to assume.
3. The range of output styles that you want the input to map to.
4. Options object, same as the options of [`interpolate()`](/docs/interpolate#options) ( _optional_)

## Return value [​](\#return-value "Direct link to Return value")

- A style object representing the interpolated styles based on the current frame.

## Usage Notes [​](\#usage-notes "Direct link to Usage Notes")

- Ensure that the `inputRange` and `outputStylesRange` arrays contain at least two values to facilitate interpolation between styles.

- The `outputStylesRange` array must have the same number of elements as `inputRange`. Each style in `outputStylesRange` corresponds to a specific value in the input range.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this hook](https://github.com/remotion-dev/remotion/blob/main/packages/animation-utils/src/transformation-helpers/interpolate-styles/index.tsx)
- [`@remotion/animation-utils`](/docs/animation-utils)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/animation-utils/interpolate-styles.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
makeTransform()](/docs/animation-utils/make-transform) [Next\
\
Overview](/docs/animated-emoji/)

- [Example](#example)
- [API](#api)
- [Return value](#return-value)
- [Usage Notes](#usage-notes)
- [See also](#see-also)
