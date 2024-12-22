---
title: @remotion/google-fonts
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/google-fonts

- [Home page](/)
- @remotion/google-fonts

On this page

# @remotion/google-fonts

The `@remotion/google-fonts` package provides APIs for easily integrating [Google Fonts](https://fonts.google.com/) into Remotion.

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm

```

bash

npm i @remotion/google-fonts
```

```

bash

npm i @remotion/google-fonts
```

```

bash

yarn add @remotion/google-fonts
```

```

bash

yarn add @remotion/google-fonts
```

```

bash

pnpm i @remotion/google-fonts
```

```

bash

pnpm i @remotion/google-fonts
```

## Usage [​](\#usage "Direct link to Usage")

To load a font, import the package `@remotion/google-fonts/<FontName>` and call [`loadFont()`](/docs/google-fonts/load-font).

```

Load all font styles
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";
const { fontFamily } = loadFont(); // "Titan One"
```

```

Load all font styles
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";
const { fontFamily } = loadFont(); // "Titan One"
```

If you want to import multiple fonts and want to avoid a variable name collision, you can import the fonts using an `import * as` statement.

```

Scope loadFont() variable
tsx

import * as Montserrat from "@remotion/google-fonts/Montserrat";
Montserrat.loadFont();
```

```

Scope loadFont() variable
tsx

import * as Montserrat from "@remotion/google-fonts/Montserrat";
Montserrat.loadFont();
```

Call [`loadFont()`](/docs/google-fonts/load-font) to start the loading process. By default, every style, weight and subset is loaded.

You can pass a style (such as `normal`, `italic`) to only load that specific style. If you want multiple styles, call `loadFont()` multiple times.

```

Load just one style
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";

loadFont("normal");
```

```

Load just one style
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";

loadFont("normal");
```

Use the TypeScript autocomplete to see the available styles. To further narrow down what's being loaded, you can specify `weights` and `subsets`. These options are also typesafe.

```

Load a specific style with limit weights and subsets
tsx

import * as Montserrat from "@remotion/google-fonts/Montserrat";

Montserrat.loadFont("normal", {
  weights: ["400", "600", "800"],
  subsets: ["latin", "latin-ext"],
});
```

```

Load a specific style with limit weights and subsets
tsx

import * as Montserrat from "@remotion/google-fonts/Montserrat";

Montserrat.loadFont("normal", {
  weights: ["400", "600", "800"],
  subsets: ["latin", "latin-ext"],
});
```

`loadFonts()` returns an object with a `fontFamily` property. You can use the `style` attribute to render text in the font you loaded.

```

Use the fontFamily property
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";
import { AbsoluteFill } from "remotion";

const { fontFamily } = loadFont();

export const GoogleFontsDemoComposition = () => {
  return (
    <AbsoluteFill
      style={{
        fontFamily,
      }}
    >
      <div>Hallo Google Fonts</div>
    </AbsoluteFill>
  );
};
```

```

Use the fontFamily property
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";
import { AbsoluteFill } from "remotion";

const { fontFamily } = loadFont();

export const GoogleFontsDemoComposition = () => {
  return (
    <AbsoluteFill
      style={{
        fontFamily,
      }}
    >
      <div>Hallo Google Fonts</div>
    </AbsoluteFill>
  );
};
```

To get information about a font, you can import the `getInfo()` function:

```

Get info about the font
tsx

import { getInfo } from "@remotion/google-fonts/Montserrat";
console.log(getInfo());
```

```

Get info about the font
tsx

import { getInfo } from "@remotion/google-fonts/Montserrat";
console.log(getInfo());
```

```

Example value of info object
json

{
  "fontFamily": "Titan One",
  "importName": "TitanOne",
  "version": "v13",
  "url": "https://fonts.googleapis.com/css2?family=Titan+One:ital,wght@0,400",
  "unicodeRanges": {
    "latin-ext": "U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF",
    "latin": "U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD"
  },
  "fonts": {
    "normal": {
      "400": {
        "latin-ext": "https://fonts.gstatic.com/s/titanone/v13/mFTzWbsGxbbS_J5cQcjCmjgm6Es.woff2",
        "latin": "https://fonts.gstatic.com/s/titanone/v13/mFTzWbsGxbbS_J5cQcjClDgm.woff2"
      }
    }
  }
}
```

```

Example value of info object
json

{
  "fontFamily": "Titan One",
  "importName": "TitanOne",
  "version": "v13",
  "url": "https://fonts.googleapis.com/css2?family=Titan+One:ital,wght@0,400",
  "unicodeRanges": {
    "latin-ext": "U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF",
    "latin": "U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD"
  },
  "fonts": {
    "normal": {
      "400": {
        "latin-ext": "https://fonts.gstatic.com/s/titanone/v13/mFTzWbsGxbbS_J5cQcjCmjgm6Es.woff2",
        "latin": "https://fonts.gstatic.com/s/titanone/v13/mFTzWbsGxbbS_J5cQcjClDgm.woff2"
      }
    }
  }
}
```

To get a list of all available fonts, you can call [`getAvailableFonts()`](/docs/google-fonts/get-available-fonts) imported from `@remotion/google-fonts`:

```

ts

import { getAvailableFonts } from "@remotion/google-fonts";

console.log(getAvailableFonts());
```

```

ts

import { getAvailableFonts } from "@remotion/google-fonts";

console.log(getAvailableFonts());
```

## APIs [​](\#apis "Direct link to APIs")

[**loadFont()** \
\
Load a Google Font](/docs/google-fonts/load-font) [**getAvailableFonts()** \
\
Static list of available fonts](/docs/google-fonts/get-available-fonts) [**getInfo()** \
\
Metadata about a specific font](/docs/google-fonts/get-info)

## Credits [​](\#credits "Direct link to Credits")

Implemented by [Hidayatullah](https://github.com/ayatkyo).

## License [​](\#license "Direct link to License")

MIT

## See also [​](\#see-also "Direct link to See also")

- [Fonts](/docs/fonts)
- [`loadFont()`](/docs/google-fonts/load-font)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/google-fonts/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
noise4D()](/docs/noise/noise-4d) [Next\
\
loadFont()](/docs/google-fonts/load-font)

- [Installation](#installation)
- [Usage](#usage)
- [APIs](#apis)
- [Credits](#credits)
- [License](#license)
- [See also](#see-also)
