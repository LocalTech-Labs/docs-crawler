---
title: measureText()
source: Remotion Documentation
last_updated: 2024-12-22
---

# measureText()

- [Home page](/)
- [@remotion/layout-utils](/docs/layout-utils/)
- measureText()

On this page

# measureText()

_Part of the [`@remotion/layout-utils`](/docs/layout-utils) package._

Calculates the width and height of specified text to be used for layout calculations. Only works in the browser, not in Node.js or Bun.

## Example [​](\#example "Direct link to Example")

```

tsx

import {measureText} from '@remotion/layout-utils';

const text = 'remotion';
const fontFamily = 'Arial';
const fontWeight = '500';
const fontSize = 12;
const letterSpacing = '1px';

measureText({
  text,
  fontFamily,
  fontWeight,
  fontSize,
  letterSpacing,
}); // { height: 14, width: 20 }
```

```

tsx

import {measureText} from '@remotion/layout-utils';

const text = 'remotion';
const fontFamily = 'Arial';
const fontWeight = '500';
const fontSize = 12;
const letterSpacing = '1px';

measureText({
  text,
  fontFamily,
  fontWeight,
  fontSize,
  letterSpacing,
}); // { height: 14, width: 20 }
```

## API [​](\#api "Direct link to API")

This function has a cache. If there are two duplicate arguments inputs, the second one will return the first result without repeating the calculation.

The function takes the following options:

### `text` [​](\#text "Direct link to text")

_Required_ _string_

Any string.

### `fontFamily` [​](\#fontfamily "Direct link to fontfamily")

_Required_ _string_

Same as CSS style `font-family`

### `fontSize` [​](\#fontsize "Direct link to fontsize")

_Required_ _number_ / _string_

Same as CSS style `font-size`. Since v4.0.125, strings are allowed too, before only numbers.

### `fontWeight` [​](\#fontweight "Direct link to fontweight")

_string_

Same as CSS style `font-weight`

### `letterSpacing` [​](\#letterspacing "Direct link to letterspacing")

_string_

Same as CSS style `letter-spacing`.

### `fontVariantNumeric` [v4.0.57](https://github.com/remotion-dev/remotion/releases/v4.0.57) [​](\#fontvariantnumeric "Direct link to fontvariantnumeric")

_string_

Same as CSS style `font-variant-numeric`.

### `textTransform` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#texttransform "Direct link to texttransform")

_string_

Same as CSS style `text-transform`.

### `validateFontIsLoaded?` [v4.0.136](https://github.com/remotion-dev/remotion/releases/v4.0.136) [​](\#validatefontisloaded "Direct link to validatefontisloaded")

_boolean_

If set to `true`, will take a second measurement with the fallback font and if it produces the same measurements, it assumes the fallback font was used and will throw an error.

### `additionalStyles` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#additionalstyles "Direct link to additionalstyles")

_object, optional_

Additional CSS properties that affect the layout of the text.

## Return value [​](\#return-value "Direct link to Return value")

An object with `height` and `width`

## Important considerations [​](\#important-considerations "Direct link to Important considerations")

See [Best practices](/docs/layout-utils/best-practices) to ensure you get correct measurements.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/layout-utils/src/layouts/measure-text.ts)
- [`@remotion/layout-utils`](/docs/layout-utils)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/layout-utils/measure-text.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Best practices](/docs/layout-utils/best-practices) [Next\
\
fillTextBox()](/docs/layout-utils/fill-text-box)

- [Example](#example)
- [API](#api)
  - [`text`](#text)
  - [`fontFamily`](#fontfamily)
  - [`fontSize`](#fontsize)
  - [`fontWeight`](#fontweight)
  - [`letterSpacing`](#letterspacing)
  - [`fontVariantNumeric`](#fontvariantnumeric)
  - [`textTransform`](#texttransform)
  - [`validateFontIsLoaded?`](#validatefontisloaded)
  - [`additionalStyles`](#additionalstyles)
- [Return value](#return-value)
- [Important considerations](#important-considerations)
- [See also](#see-also)
