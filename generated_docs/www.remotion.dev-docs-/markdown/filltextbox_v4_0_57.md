---
title: fillTextBox()v4.0.57
source: Remotion Documentation
last_updated: 2024-12-22
---

# fillTextBox()v4.0.57

- [Home page](/)
- [@remotion/layout-utils](/docs/layout-utils/)
- fillTextBox()

On this page

# fillTextBox() [v4.0.57](https://github.com/remotion-dev/remotion/releases/v4.0.57)

_Part of the [`@remotion/layout-utils`](/docs/layout-utils) package._

Calculate whether the text exceeds the box and wraps within the container. Only works in the browser, not in Node.js or Bun.

## Example [​](\#example "Direct link to Example")

```

tsx

import { fillTextBox } from "@remotion/layout-utils";

const fontFamily = "Arial";
const fontSize = 12;

const box = fillTextBox({ maxLines: 4, maxBoxWidth: 100 });
box.add({ text: "Hello", fontFamily, fontSize }); // {exceedsBox: false, newLine: false}
box.add({ text: "World!", fontFamily, fontSize }); // {exceedsBox: false, newLine: false}
// Doesn't fit on the previous line anymore
box.add({ text: "How", fontFamily, fontSize }); // {exceedsBox: false, newLine: true}
// ...
// Doesn't fix in the box anymore
box.add({ text: "the end", fontFamily, fontSize }); // {exceedsBox: true, newLine: false}
```

```

tsx

import { fillTextBox } from "@remotion/layout-utils";

const fontFamily = "Arial";
const fontSize = 12;

const box = fillTextBox({ maxLines: 4, maxBoxWidth: 100 });
box.add({ text: "Hello", fontFamily, fontSize }); // {exceedsBox: false, newLine: false}
box.add({ text: "World!", fontFamily, fontSize }); // {exceedsBox: false, newLine: false}
// Doesn't fit on the previous line anymore
box.add({ text: "How", fontFamily, fontSize }); // {exceedsBox: false, newLine: true}
// ...
// Doesn't fix in the box anymore
box.add({ text: "the end", fontFamily, fontSize }); // {exceedsBox: true, newLine: false}
```

## API [​](\#api "Direct link to API")

The function takes the following options:

### `maxBoxWidth` [​](\#maxboxwidth "Direct link to maxboxwidth")

_Required_ _number_

The max box width, unit `px`.

### `maxLines` [​](\#maxlines "Direct link to maxlines")

_Required_ _number_

The max lines of the box.

## Return value [​](\#return-value "Direct link to Return value")

An object with an `add()` method, which can be used to add words to the text box.

### Arguments [​](\#arguments "Direct link to Arguments")

#### `text` [​](\#text "Direct link to text")

_Required_ _string_

Any string.

#### `fontFamily` [​](\#fontfamily "Direct link to fontfamily")

_Required_ _string_

Same as CSS style `font-family`.

#### `fontSize` [​](\#fontsize "Direct link to fontsize")

_Required_ _number_

Same as CSS style `font-size`. Only _numbers_ allowed, unit `px`.

#### `fontWeight` [​](\#fontweight "Direct link to fontweight")

_string_

Same as CSS style `font-weight`.

#### `fontVariantNumeric` [​](\#fontvariantnumeric "Direct link to fontvariantnumeric")

_string_

Same as CSS style `font-variant-numeric`.

#### `textTransform` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#texttransform "Direct link to texttransform")

_string_

Same as CSS style `text-transform`.

#### `validateFontIsLoaded?` [v4.0.136](https://github.com/remotion-dev/remotion/releases/v4.0.136) [​](\#validatefontisloaded "Direct link to validatefontisloaded")

_boolean_

If set to `true`, will take a second measurement with the fallback font and if it produces the same measurements, it assumes the fallback font was used and will throw an error.

#### `additionalStyles` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#additionalstyles "Direct link to additionalstyles")

_object, optional_

Additional CSS properties that affect the layout of the text.

### Return value [​](\#return-value-1 "Direct link to Return value")

The add method returns an object with two properties:

- `exceedsBox`:
_Boolean_, whether adding the word would cause the text to exceed the width of the box.
- `newLine`:
_Boolean_, whether adding the word would require starting a new line in the text box.

## Important considerations [​](\#important-considerations "Direct link to Important considerations")

See [Best practices](/docs/layout-utils/best-practices) to ensure you get correct measurements.

## See also [​](\#see-also "Direct link to See also")

- [measureText()](/docs/layout-utils/measure-text)
- [`@remotion/layout-utils`](/docs/layout-utils)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/layout-utils/fill-text-box.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
measureText()](/docs/layout-utils/measure-text) [Next\
\
fitText()](/docs/layout-utils/fit-text)

- [Example](#example)
- [API](#api)
  - [`maxBoxWidth`](#maxboxwidth)
  - [`maxLines`](#maxlines)
- [Return value](#return-value)
  - [Arguments](#arguments)
  - [Return value](#return-value-1)
- [Important considerations](#important-considerations)
- [See also](#see-also)
