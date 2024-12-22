---
title: getInfo()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getInfo()

- [Home page](/)
- [@remotion/google-fonts](/docs/google-fonts/)
- getInfo()

On this page

# getInfo()

_Part of the [`@remotion/google-fonts`](/docs/google-fonts) package_

Each font exports an `getInfo()` function that can be used to get various information about the font at runtime, such as which weights, styles and subsets are available.

## Usage [​](\#usage "Direct link to Usage")

To get information about a font, you can import the `info` property:

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

## See also [​](\#see-also "Direct link to See also")

- [Fonts](/docs/fonts)
- [`@remotion/google-fonts`](/docs/google-fonts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/google-fonts/get-info.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAvailableFonts()](/docs/google-fonts/get-available-fonts) [Next\
\
@remotion/shapes](/docs/shapes/)

- [Usage](#usage)
- [See also](#see-also)
