---
title: Flickering
source: Remotion Documentation
last_updated: 2024-12-22
---

# Flickering

- [Home page](/)
- Troubleshooting
- Flickering

On this page

# Flickering

[![](https://i.ytimg.com/vi/M7BOPECeqV8/hq720.jpg)\
\
Also available as a 6min video\
\
**Avoid this common mistake in Remotion**](https://www.youtube.com/watch?v=M7BOPECeqV8)

If your video flickers or suffers from choppiness during rendering, it is an indication that:

[1](#1)

there is a **multi-threading issue** or

[2](#2) the renderer does not **wait for data or assets to be loaded**
.

## Multi-threading issue [​](\#multi-threading-issue "Direct link to Multi-threading issue")

The rendering process of Remotion works as follows:

![](/img/parallel-rendering.png)

We open multiple tabs to render the video to speed up the process dramatically.

Tabs don't share state and animations that run independent of [`useCurrentFrame()`](/docs/use-current-frame) will break.

### Solution [​](\#solution "Direct link to Solution")

Code your video in a way that animations run purely off the value of [`useCurrentFrame()`](/docs/use-current-frame).

Think of your component as a function that transforms a frame number into an image.

Does your component satisfy the following criteria?

[1](#1)

A component should always display the same visual when called multiple
times.

[2](#2) A component should not rely on frames being rendered in order.

[3](#3) A component should not animate when the video is paused.

[4](#4) A component should not rely on randomness - Exception: [`random()`](/docs/random)

### Bypass multithreading [​](\#bypass-multithreading "Direct link to Bypass multithreading")

If your animation will not break if the frames are rendered in order, users often use the [`--concurrency=1`](/docs/cli/render#--concurrency) flag. This will fix flickering / choppiness in many cases and is a viable path if the effort of refactoring is too big.

The drawback of this technique is that it is way slower and that the correct timing of the animations is still not guaranteed. It will also block rendering via [Remotion Lambda](/lambda).

## Asset loading issue [​](\#asset-loading-issue "Direct link to Asset loading issue")

Remotion needs to know that an asset is not loaded yet so it can block rendering.

Otherwise, it will take a screenshot of a loading state.

### Solution [​](\#solution-1 "Direct link to Solution")

[1](#1)

Use the `<Img>`
, `<Video>`
, `<OffthreadVideo>`
, `<Audio>`
, `<Iframe>` and `<Gif>` as they wait for the assets to be loaded.

[2](#2) When loading data, use the [`delayRender()`](/docs/delay-render) function.

[3](#3) Ensure you correctly wait [for fonts to load](/docs/fonts).

[4](#4) Only call [`fitText()`](/docs/layout-utils/fit-text), [`fillTextBox()`](/docs/layout-utils/fill-text-box) and [`measureText()`](/docs/layout-utils/measure-text) once fonts are loaded.

[5](#5) Avoid using the [`background-image` and `mask-image` CSS properties](/docs/troubleshooting/background-image).

## Flickering `<Video>` tag [​](\#flickering-video-tag "Direct link to flickering-video-tag")

Adding many [`<Video>`](/docs/video) tags can lead to stutters.

If you are experiencing the problem, consider using the [`<OffthreadVideo>`](/docs/offthreadvideo) component for frame-perfect rendering.

## Integrations [​](\#integrations "Direct link to Integrations")

See the list of [third-party integrations](/docs/third-party) to see if there is a solution for synchronizing your animation with [`useCurrentFrame()`](/docs/use-current-frame).

## Why Remotion works this way [​](\#why-remotion-works-this-way "Direct link to Why Remotion works this way")

- Rendering speed is important, especially with server-side rendering.

Rendering each frame sequentially would be detrimental for speed, a sacrifice that is not worth it when it's possible to write concurrency-safe videos.

- Setting `--concurrency=1` on a video that would be choppy otherwise does not fully fix the problem.

Often the result looks okay only because of coincidence, because the rendering speed is approximately the same as the animation speed. There is no real timing synchronization and results will differ across machines.

- Deterministic videos enable distributed video renders like [Remotion Lambda](/docs/lambda), which can render a video much faster than real-time.

## See Also [​](\#see-also "Direct link to See Also")

- [Avoiding flickering in `<Player>`](/docs/troubleshooting/player-flicker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/flickering.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Non-seekable media](/docs/non-seekable-media) [Next\
\
Version mismatch](/docs/version-mismatch)

- [Multi-threading issue](#multi-threading-issue)
  - [Solution](#solution)
  - [Bypass multithreading](#bypass-multithreading)
- [Asset loading issue](#asset-loading-issue)
  - [Solution](#solution-1)
- [Flickering `<Video>` tag](#flickering-video-tag)
- [Integrations](#integrations)
- [Why Remotion works this way](#why-remotion-works-this-way)
- [See Also](#see-also)
