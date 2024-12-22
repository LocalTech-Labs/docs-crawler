---
title: Example
source: Remotion Documentation
last_updated: 2024-12-22
---

# Example

- [Home page](/)
- [Command line](/docs/cli/)
- ffprobe

On this page

_available since v4.0_

In order to use `ffprobe` without having to directly install it, Remotion provides it via `npx remotion ffprobe`.

Note that in order to keep the binary size small, the FFprobe binary only understand the codecs that Remotion itself supports: H.264, H.265, VP8, VP9 and ProRes. A binary from the 7.1 release line of FFprobe is used.

# Example

```

bash

npx remotion ffprobe your_video.mp4
```

```

bash

npx remotion ffprobe your_video.mp4
```

To find out more about FFprobe, visit their [docs](https://ffmpeg.org/ffprobe.html).

## See also [​](\#see-also "Direct link to See also")

- [`npx remotion ffmpeg`](/docs/cli/ffmpeg)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/ffprobe.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
ffmpeg](/docs/cli/ffmpeg) [Next\
\
gpu](/docs/cli/gpu)

- [See also](#see-also)
