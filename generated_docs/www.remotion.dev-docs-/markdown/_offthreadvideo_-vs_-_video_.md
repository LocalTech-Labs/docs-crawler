---
title: <OffthreadVideo> vs. <Video>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <OffthreadVideo> vs. <Video>

- [Home page](/)
- Embedding videos
- <OffthreadVideo> vs. <Video>

On this page

# <OffthreadVideo> vs. <Video>

We offer two components for including other videos in your video: [`<Video />`](/docs/video) and [`<OffthreadVideo />`](/docs/offthreadvideo).

This page offers a comparison to help you decide which tag to use.

## [`<OffthreadVideo />`](/docs/offthreadvideo) [​](\#offthreadvideo- "Direct link to offthreadvideo-")

info

Prefer [`<OffthreadVideo />`](/docs/offthreadvideo) whenever possible.

A more sophisticated way of embedding a video, which:

- **During preview:** renders a HTML5 `<video>` tag.
- **While rendering:** Extracts the exact frame outside the browser and displays it in an [`<Img>`](/docs/img) tag.

**Pros**

✅   More videos can be displayed simulatenously as Chrome will not apply throttling.

✅   No flickers or duplicate frames in the output video can occur.

✅   Supports more codecs such as ProRes (only during render-time).

✅   Is usually faster to render.

**Cons**

⛔   The video needs to be downloaded fully before a frame can be rendered.

⛔   No ref can be attached to this element.

⛔   More work is required to loop a video. Check out: [Looping an `<Offthread>` video](/docs/offthreadvideo#looping-a-video).

⛔   Supports transparent videos only if the [`transparent`](/docs/offthreadvideo#transparent) is set which is a bit slower.

## [`<Video />`](/docs/video) [​](\#video- "Direct link to video-")

Is based on the native HTML5 `<video>` element and therefore behaves similar to it.

**Pros**

✅   Can render a video without having it to be downloaded fully (if you don't pass the [`muted`](/docs/video/#muted) prop, the video will still be downloaded fully to extract its audio).

✅   You can attach a ref to the `<video>` element.

**Cons**

⛔   Fewer codecs are supported.

⛔   Chrome may throttle video tags if the page is heavy.

⛔   If too many video tags get rendered simultaneously, a timeout may occur.

⛔   If the input video framerate does not match with the output framerate, some duplicate frames may occur in the output.

⛔   A Google Chrome build with proprietary codecs is required.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/video-vs-offthreadvideo.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
HTTP Live Streaming](/docs/miscellaneous/snippets/hls) [Next\
\
Passing props](/docs/parameterized-rendering)

- [`<OffthreadVideo />`](#offthreadvideo-)
- [`<Video />`](#video-)
