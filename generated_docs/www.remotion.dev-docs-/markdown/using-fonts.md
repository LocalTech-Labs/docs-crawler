---
title: Using fonts
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using fonts

- [Home page](/)
- Designing visuals
- Fonts

On this page

# Using fonts

Here are some ways how you can use custom fonts in Remotion.

## Google Fonts using `@remotion/google-fonts` [​](\#google-fonts-using-remotiongoogle-fonts "Direct link to google-fonts-using-remotiongoogle-fonts")

_available from v3.2.40_

[`@remotion/google-fonts`](/docs/google-fonts) is a type-safe way to load Google fonts without having to create CSS files.

```

MyComp.tsx
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";

const { fontFamily } = loadFont();

const GoogleFontsComp: React.FC = () => {
  return <div style={{ fontFamily }}>Hello, Google Fonts</div>;
};
```

```

MyComp.tsx
tsx

import { loadFont } from "@remotion/google-fonts/TitanOne";

const { fontFamily } = loadFont();

const GoogleFontsComp: React.FC = () => {
  return <div style={{ fontFamily }}>Hello, Google Fonts</div>;
};
```

## Google Fonts using CSS [​](\#google-fonts-using-css "Direct link to Google Fonts using CSS")

Import the CSS that Google Fonts gives you.

note

From version 2.2 on, Remotion will automatically wait until the fonts are loaded.

```

font.css
css

@import url("https://fonts.googleapis.com/css2?family=Bangers");
```

```

font.css
css

@import url("https://fonts.googleapis.com/css2?family=Bangers");
```

```

MyComp.tsx
tsx

import "./font.css";

const MyComp: React.FC = () => {
  return <div style={{ fontFamily: "Bangers" }}>Hello</div>;
};
```

```

MyComp.tsx
tsx

import "./font.css";

const MyComp: React.FC = () => {
  return <div style={{ fontFamily: "Bangers" }}>Hello</div>;
};
```

## Local fonts using `@remotion/fonts` [​](\#local-fonts-using-remotionfonts "Direct link to local-fonts-using-remotionfonts")

_available from v4.0.164_

Put the font into your `public/` folder.

Put the [`loadFont()`](/docs/fonts-api/load-font) call somewhere in your app where it gets executed:

```

load-fonts.ts
tsx

import { loadFont } from "@remotion/fonts";
import { staticFile } from "remotion";

const fontFamily = "Inter";

loadFont({
  family: fontFamily,
  url: staticFile("Inter-Regular.woff2"),
  weight: "500",
}).then(() => {
  console.log("Font loaded!");
});
```

```

load-fonts.ts
tsx

import { loadFont } from "@remotion/fonts";
import { staticFile } from "remotion";

const fontFamily = "Inter";

loadFont({
  family: fontFamily,
  url: staticFile("Inter-Regular.woff2"),
  weight: "500",
}).then(() => {
  console.log("Font loaded!");
});
```

The font is now available for use:

```

MyComp.tsx
tsx

<div style={{ fontFamily: fontFamily }}>Some text</div>;
```

```

MyComp.tsx
tsx

<div style={{ fontFamily: fontFamily }}>Some text</div>;
```

## Local fonts (manually) [​](\#local-fonts-manually "Direct link to Local fonts (manually)")

You may load fonts by using the web-native [`FontFace`](https://developer.mozilla.org/en-US/docs/Web/API/FontFace) API.

```

load-fonts.ts
tsx

import { continueRender, delayRender, staticFile } from "remotion";

const waitForFont = delayRender();
const font = new FontFace(
  `Bangers`,
  `url('${staticFile("bangers.woff2")}') format('woff2')`,
);

font
  .load()
  .then(() => {
    document.fonts.add(font);
    continueRender(waitForFont);
  })
  .catch((err) => console.log("Error loading font", err));
```

```

load-fonts.ts
tsx

import { continueRender, delayRender, staticFile } from "remotion";

const waitForFont = delayRender();
const font = new FontFace(
  `Bangers`,
  `url('${staticFile("bangers.woff2")}') format('woff2')`,
);

font
  .load()
  .then(() => {
    document.fonts.add(font);
    continueRender(waitForFont);
  })
  .catch((err) => console.log("Error loading font", err));
```

note

If your Typescript types give errors, install the newest version of the `@types/web` package.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/fonts.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Audio](/docs/using-audio) [Next\
\
Measuring items](/docs/measuring)

- [Google Fonts using `@remotion/google-fonts`](#google-fonts-using-remotiongoogle-fonts)
- [Google Fonts using CSS](#google-fonts-using-css)
- [Local fonts using `@remotion/fonts`](#local-fonts-using-remotionfonts)
- [Local fonts (manually)](#local-fonts-manually)
