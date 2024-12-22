---
title: Which video formats does Remotion support?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Which video formats does Remotion support?

- [Home page](/)
- [Rendering](/docs/render)
- Video formats

# Which video formats does Remotion support?

Remotion supports a variety of video codecs.

**Output formats**: Videos can be rendered to `H.264` (MP4), `H.265`, (HEVC), `VP8` and `VP9` (WebM), `ProRes` as well as GIFs. See the [Encoding guide](/docs/encoding) for more detail.

**Codecs supported for [`<Video>`](/docs/video):** Remotion uses the default `<video>` tag for playback and therefore inherits codec support from the browser.

During rendering this means that the Chrome codecs are supported. In the [Remotion Studio](/docs/terminology/studio) and in the [Player](/docs/terminology/player), the codec support from the browser in which the webpage is hosted in applies.

Codec support varies between browsers and changes from time to time. Refer to this [MDN article](https://developer.mozilla.org/en-US/docs/Web/Media/Formats/Video_codecs) for browser support.

note

Prior to v4.0.18, if you did not have a local copy of Chrome, Remotion would download a copy of Chromium which [would not support the proprietary H.264 and H.265 codecs](/docs/miscellaneous/chrome-headless-shell).

**Codecs supported for [`<OffthreadVideo>`](/docs/offthreadvideo)**: The same support that applies to `<Video>` applies to `<OffthreadVideo>` as well, however, more codecs can be supported during rendering since FFmpeg is used under the hood to read the video file. We however don't maintain an exact list of supported video codecs.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/video-formats.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Render all compositions](/docs/render-all) [Next\
\
Emitting Artifacts](/docs/artifacts)
