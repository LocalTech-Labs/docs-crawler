---
title: Debugging font measurement bugs
source: Remotion Documentation
last_updated: 2024-12-22
---

# Debugging font measurement bugs

- [Home page](/)
- [@remotion/layout-utils](/docs/layout-utils/)
- Debugging

On this page

# Debugging font measurement bugs

When you find that the measurements that you are getting from [`measureText()`](/docs/layout-utils/measure-text), [`fillTextBox()`](/docs/layout-utils/fill-text-box) or [`fitText()`](/docs/layout-utils/fit-text) are off and are causing layout issues, follow the debugging instructions on this page.

[1](#1)

Open your project in the Remotion Studio and select the
composition in which you observe the issue.

[2](#2) Set the zoom to 100% to rule out any issues that arise from the
zoom level.

[3](#3) Render at the bottom of your component a `<div>` that contains the string and all the properties that you passed to [`measureText()`](/docs/layout-utils/measure-text).

```

Sample markup
tsx

import {AbsoluteFill} from 'remotion';

const MyComp: React.FC = () => {
  const fontSize = 100;
  const fontWeight = 'bold';
  const fontFamily = 'Roboto';
  const text = 'Hello World ';
  const letterSpacing = undefined;
  const fontVariantNumeric = undefined;
  const textTransform = undefined;

  return (
    <AbsoluteFill>
      <div
        id="remotion-measurer"
        style={{
          display: 'inline-block',
          whiteSpace: 'pre',
          fontSize,
          fontWeight,
          fontFamily,
          letterSpacing,
          fontVariantNumeric,
          textTransform,
        }}
      >
        {text}
      </div>
    </AbsoluteFill>
  );
};
```

```

Sample markup
tsx

import {AbsoluteFill} from 'remotion';

const MyComp: React.FC = () => {
  const fontSize = 100;
  const fontWeight = 'bold';
  const fontFamily = 'Roboto';
  const text = 'Hello World ';
  const letterSpacing = undefined;
  const fontVariantNumeric = undefined;
  const textTransform = undefined;

  return (
    <AbsoluteFill>
      <div
        id="remotion-measurer"
        style={{
          display: 'inline-block',
          whiteSpace: 'pre',
          fontSize,
          fontWeight,
          fontFamily,
          letterSpacing,
          fontVariantNumeric,
          textTransform,
        }}
      >
        {text}
      </div>
    </AbsoluteFill>
  );
};
```

[4](#4)

Open the Chrome DevTools and select the `<div id="remotion-measurer">` node in the "Elements" tab.

[5](#5) The node lights up, and you can see its measurements. See if there are any unexpected deviations.

[6](#6) Use the "Computed" tabs and go through all properties that may affect the layout, and compare them with the node in your composition that is causing issue. If the measurements are different, there will be at least 1 computed proeprty that has a different value.

[7](#7) Align your markup so that it has the same computed properties as the node in your composition.

## Remember [​](\#remember "Direct link to Remember")

- The text gets measured with `whiteSpace: "pre"`, which includes whitespace.
  - If you don't want to measure whitespace. call `.trim()` on your string.
  - Make sure `whiteSpace: "pre"` is applied to your whole container, not just the individual words.
  - Make sure you don't remove any whitespace while splitting your text that you leave in while measuring.
- Make sure the font is loaded when you are calling `measureText()`.
  - Use the `validateFontIsLoaded` option to ensure the font is loaded.
- External styles may apply. See if you have CSS Resets, TailwindCSS stylesheets or other external resources messing up your layout.
  - If this is the case, you can see it in the Computed Properties panel in the Chrome Dev Tools.
- Don't use `padding` or `border` on your text. Use `outline` instead of `border`.
- Don't multiply the font size with [`useCurrentScale()`](/docs/use-current-scale). It will lead to a broken layout.

## See also [​](\#see-also "Direct link to See also")

You can also check out the source code of `measureText()`, paste and customize it in your project to aid debugging.

- [Source code for `measureText()`](https://github.com/remotion-dev/remotion/blob/main/packages/layout-utils/src/layouts/measure-text.ts)
- [Best practices](/docs/layout-utils/best-practices)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/layout-utils/debug.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
fitText()](/docs/layout-utils/fit-text) [Next\
\
@remotion/animation-utils](/docs/animation-utils/)

- [Remember](#remember)
- [See also](#see-also)
