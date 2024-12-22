---
title: Best practices for @remotion/layout-utils
source: Remotion Documentation
last_updated: 2024-12-22
---

# Best practices for @remotion/layout-utils

- [Home page](/)
- [@remotion/layout-utils](/docs/layout-utils/)
- Best practices

On this page

# Best practices for @remotion/layout-utils

Take note of the following points to ensure correct measurements when using the [`@remotion/layout-utils`](/docs/layout-utils) package.

These tips apply to all of [`measureText()`](/docs/layout-utils/measure-text), [`fitText()`](/docs/layout-utils/fit-text), and [`fillTextBox()`](/docs/layout-utils/fill-text-box).

## Wait until the font is loaded [​](\#wait-until-the-font-is-loaded "Direct link to Wait until the font is loaded")

Only call [`measureText()`](/docs/layout-utils/measure-text) after the font is loaded. This applies to Google Fonts (example below) or any other font loading mechanism.

### Example with `useEffect` [​](\#example-with-useeffect "Direct link to example-with-useeffect")

```

MyComp.tsx
tsx

import {useState, useEffect} from 'react';
import {Dimensions, measureText} from '@remotion/layout-utils';
import {loadFont, fontFamily} from '@remotion/google-fonts/Inter';

const {waitUntilDone} = loadFont('normal');

const MyComp: React.FC = () => {
  const [dimensions, setDimensions] = useState<Dimensions | null>(null);

  useEffect(() => {
    // Wait until the font is loaded before measuring text
    waitUntilDone().then(() => {
      const measurement = measureText({
        fontFamily: fontFamily,
        fontSize: 14,
        fontWeight: '400',
        text: 'Hello world',
      });

      // We don't need to use delayRender() here, because
      // font loading from @remotion/google-fonts is already wrapped in it
      setDimensions(measurement);
    });
  }, []);

  return null;
};
```

```

MyComp.tsx
tsx

import {useState, useEffect} from 'react';
import {Dimensions, measureText} from '@remotion/layout-utils';
import {loadFont, fontFamily} from '@remotion/google-fonts/Inter';

const {waitUntilDone} = loadFont('normal');

const MyComp: React.FC = () => {
  const [dimensions, setDimensions] = useState<Dimensions | null>(null);

  useEffect(() => {
    // Wait until the font is loaded before measuring text
    waitUntilDone().then(() => {
      const measurement = measureText({
        fontFamily: fontFamily,
        fontSize: 14,
        fontWeight: '400',
        text: 'Hello world',
      });

      // We don't need to use delayRender() here, because
      // font loading from @remotion/google-fonts is already wrapped in it
      setDimensions(measurement);
    });
  }, []);

  return null;
};
```

### Example with high-order component [​](\#example-with-high-order-component "Direct link to Example with high-order component")

This following logic is borrowed from the [Remotion Recorder](/docs/recorder).

It keeps the code clean by exposing a high-order component which only mounts its children when the font is loaded.

A file is defined which loads some fonts:

```

fonts.ts
tsx

import {
  fontFamily as regularFont,
  loadFont as loadRegular,
} from '@remotion/google-fonts/Inter';

import {
  fontFamily as monospaceFont,
  loadFont as loadMonospace,
} from '@remotion/google-fonts/RobotoMono';

import {cancelRender, continueRender, delayRender} from 'remotion';

const regular = loadRegular();
const monospace = loadMonospace();

export const waitForFonts = async () => {
  await regular.waitUntilDone();
  await monospace.waitUntilDone();
};

const delay = delayRender('Loading fonts');

waitForFonts()
  .then(() => continueRender(delay))
  .catch((err) => cancelRender(err));
```

```

fonts.ts
tsx

import {
  fontFamily as regularFont,
  loadFont as loadRegular,
} from '@remotion/google-fonts/Inter';

import {
  fontFamily as monospaceFont,
  loadFont as loadMonospace,
} from '@remotion/google-fonts/RobotoMono';

import {cancelRender, continueRender, delayRender} from 'remotion';

const regular = loadRegular();
const monospace = loadMonospace();

export const waitForFonts = async () => {
  await regular.waitUntilDone();
  await monospace.waitUntilDone();
};

const delay = delayRender('Loading fonts');

waitForFonts()
  .then(() => continueRender(delay))
  .catch((err) => cancelRender(err));
```

Then a higher order component is defined which only renders it's children when the font is loaded:

```

WaitForFonts.tsx
tsx

import React, {useEffect, useState} from 'react';
import {cancelRender, continueRender, delayRender} from 'remotion';
import {waitForFonts} from './fonts';

export const WaitForFonts: React.FC<{
  children: React.ReactNode;
}> = ({children}) => {
  const [fontsLoaded, setFontsLoaded] = useState(false);
  const [handle] = useState(() => delayRender());

  useEffect(() => {
    const delay = delayRender('Waiting for fonts to be loaded');

    waitForFonts()
      .then(() => {
        continueRender(handle);
        continueRender(delay);
        setFontsLoaded(true);
      })
      .catch((err) => {
        cancelRender(err);
      });
  }, [fontsLoaded, handle]);

  if (!fontsLoaded) {
    return null;
  }

  return <>{children}</>;
};
```

```

WaitForFonts.tsx
tsx

import React, {useEffect, useState} from 'react';
import {cancelRender, continueRender, delayRender} from 'remotion';
import {waitForFonts} from './fonts';

export const WaitForFonts: React.FC<{
  children: React.ReactNode;
}> = ({children}) => {
  const [fontsLoaded, setFontsLoaded] = useState(false);
  const [handle] = useState(() => delayRender());

  useEffect(() => {
    const delay = delayRender('Waiting for fonts to be loaded');

    waitForFonts()
      .then(() => {
        continueRender(handle);
        continueRender(delay);
        setFontsLoaded(true);
      })
      .catch((err) => {
        cancelRender(err);
      });
  }, [fontsLoaded, handle]);

  if (!fontsLoaded) {
    return null;
  }

  return <>{children}</>;
};
```

Then the component can be wrap any other component that calls text measurement APIs:

```

MyComp.tsx
tsx

import React from 'react';
import {regular} from './fonts';
import {WaitForFonts} from './WaitForFonts';
import {measureText} from '@remotion/layout-utils';

const MyCompInner: React.FC = () => {
  // Safe to call measureText() here
  const measurement = measureText({
    fontFamily: regular,
    fontSize: 14,
    fontWeight: '400',
    text: 'Hello world',
  });

  return null;
};

export const MyComp: React.FC = () => {
  return (
    <WaitForFonts>
      <MyCompInner />
    </WaitForFonts>
  );
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {regular} from './fonts';
import {WaitForFonts} from './WaitForFonts';
import {measureText} from '@remotion/layout-utils';

const MyCompInner: React.FC = () => {
  // Safe to call measureText() here
  const measurement = measureText({
    fontFamily: regular,
    fontSize: 14,
    fontWeight: '400',
    text: 'Hello world',
  });

  return null;
};

export const MyComp: React.FC = () => {
  return (
    <WaitForFonts>
      <MyCompInner />
    </WaitForFonts>
  );
};
```

### Use the `validateFontIsLoaded` option [v4.0.136](https://github.com/remotion-dev/remotion/releases/v4.0.136) [​](\#use-the-validatefontisloaded-option "Direct link to use-the-validatefontisloaded-option")

Pass `validateFontIsLoaded: true` to any of [`measureText()`](/docs/layout-utils/measure-text), [`fitText()`](/docs/layout-utils/fit-text), and [`fillTextBox()`](/docs/layout-utils/fill-text-box) to validate that the font family you passed is actually loaded.

This will take a second measurement with the fallback font and if it produces the same measurements, it assumes the fallback font was used and will throw an error.

note

In Remotion v5, this will become the default.

## Match all font properties [​](\#match-all-font-properties "Direct link to Match all font properties")

When measuring text, ensure that all font properties match the ones you are going to use in your video.

This includes `fontFamily`, `fontSize`, and `fontWeight`, `letterSpacing` and `fontVariantNumeric`.

You could make reusable variables that you reference in both the measuring function and the actual component.

```

Using variables for font properties
tsx

import {measureText} from '@remotion/layout-utils';

const text = 'Hello world';
const fontFamily = 'Inter';
const fontWeight = 'bold';
const fontSize = 16;

// Use the variable in the measurement function:
measureText({
  text,
  fontFamily,
  fontWeight,
  fontSize,
});

// As well as in markup
<div style={{fontFamily, fontWeight, fontSize}}>{text}</div>;
```

```

Using variables for font properties
tsx

import {measureText} from '@remotion/layout-utils';

const text = 'Hello world';
const fontFamily = 'Inter';
const fontWeight = 'bold';
const fontSize = 16;

// Use the variable in the measurement function:
measureText({
  text,
  fontFamily,
  fontWeight,
  fontSize,
});

// As well as in markup
<div style={{fontFamily, fontWeight, fontSize}}>{text}</div>;
```

## Be aware of borders and padding [​](\#be-aware-of-borders-and-padding "Direct link to Be aware of borders and padding")

Adding a `padding` or a `border` to a word will skew the measurements.

Avoid using `padding` altogether and use the natural spacing between words.

Instead of `border`, use an `outline` to add a line outside the text without affecting its layout.

## Whitespace [​](\#whitespace "Direct link to Whitespace")

When measuring, the Layout utils will wrap the text in a `<span>` with `display: inline-block` and `white-space: pre` applied.

This will also measure the width of the whitespace characters.

Add those two CSS properties to your markup as well to match it with the measurements.

## Server-side rendering [​](\#server-side-rendering "Direct link to Server-side rendering")

The layout utilities need to be run in a browser.

If you are using them in a component that will be server-side rendered, then we recommend you call the functions in a `useEffect`, like on the first example on this page.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/layout-utils/best-practices.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/layout-utils](/docs/layout-utils/) [Next\
\
measureText()](/docs/layout-utils/measure-text)

- [Wait until the font is loaded](#wait-until-the-font-is-loaded)
  - [Example with `useEffect`](#example-with-useeffect)
  - [Example with high-order component](#example-with-high-order-component)
  - [Use the `validateFontIsLoaded` option](#use-the-validatefontisloaded-option)
- [Match all font properties](#match-all-font-properties)
- [Be aware of borders and padding](#be-aware-of-borders-and-padding)
- [Whitespace](#whitespace)
- [Server-side rendering](#server-side-rendering)
