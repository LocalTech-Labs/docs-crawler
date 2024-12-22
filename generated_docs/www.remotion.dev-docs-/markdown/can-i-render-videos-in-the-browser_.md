---
title: Can I render videos in the browser?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Can I render videos in the browser?

- [Home page](/)
- FAQ
- Browser rendering

On this page

# Can I render videos in the browser?

**Rendering videos in the browser is not supported.** In order to render videos, you need to hook up [server-side rendering](/docs/ssr), [Remotion Lambda](/docs/lambda), or [render videos locally](/docs/render).

## Will it be supported in the future? [​](\#will-it-be-supported-in-the-future "Direct link to Will it be supported in the future?")

Currently there is no browser API that allows to capture the viewport. There [are](https://x.com/fserb/status/1794058245901824349) [some](https://github.com/WICG/proposals/issues/73) proposed APIs for it.

If such an API gets introduced, we can consider supporting browser rendering in the future.

## Why not... [​](\#why-not "Direct link to Why not...")

### a Chrome extension? [​](\#a-chrome-extension "Direct link to a Chrome extension?")

Chrome extensions do get the privilege of capturing the viewport. We may explore this in the future, but the combination of asking the user to install an extension and slow render times means we are not prioritizing this feature.

### [`canvas.getImageData()`](https://developer.mozilla.org/de/docs/Web/API/CanvasRenderingContext2D/getImageData)? [​](\#canvasgetimagedata "Direct link to canvasgetimagedata")

It allows to capture the pixels from a canvas, however Remotion videos can be written with any web technology including HTML and SVG. Canvas elements are just a subset of what's supported in Remotion.

### [`html2canvas`](https://html2canvas.hertzen.com/)? [​](\#html2canvas "Direct link to html2canvas")

It implements it's own rendering engine which only supports a subset of CSS properties.

You would only have access to a very limited feature set.

### SVG `<foreignElement>`? [​](\#svg-foreignelement "Direct link to svg-foreignelement")

You can render HTML markup inside an SVG `<foreignElement>` and then render that SVG to a canvas and then call [`getImageData()`](https://developer.mozilla.org/de/docs/Web/API/CanvasRenderingContext2D/getImageData) to turn that into an image.

This is the closest thing to browser rendering, however there are problems with `<img>` tags and potentially other technologies as well. It's not yet fully out of the questions, but seems hacky so far.

## How can I avoid server-rendering? [​](\#how-can-i-avoid-server-rendering "Direct link to How can I avoid server-rendering?")

We cannot fully avoid server-rendering, but we are taking initiatives to make it easier and efficient:

- [Remotion Lambda](/lambda) is a batteries-included renderer for Remotion that you only have to pay for when you use it.

## See also [​](\#see-also "Direct link to See also")

- [`<Player>`](/player): Displaying a Remotion video in the browser without encoding it

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/render-in-browser.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
AI System Prompt](/docs/system-prompt) [Next\
\
Automatic duration](/docs/miscellaneous/automatic-duration)

- [Will it be supported in the future?](#will-it-be-supported-in-the-future)
- [Why not...](#why-not)
  - [a Chrome extension?](#a-chrome-extension)
  - [`canvas.getImageData()`?](#canvasgetimagedata)
  - [`html2canvas`?](#html2canvas)
  - [SVG `<foreignElement>`?](#svg-foreignelement)
- [How can I avoid server-rendering?](#how-can-i-avoid-server-rendering)
- [See also](#see-also)
