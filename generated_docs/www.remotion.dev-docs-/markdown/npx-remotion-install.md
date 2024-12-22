---
title: npx remotion install
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion install

On this page

# npx remotion install

_removed in v4.0.0, available from v3.3_

_See ["No more FFmpeg installation"](/docs/4-0-migration#no-more-ffmpeg-install-ffmpegexecutable-option-removed)._

This page is for archival purpose.

Ensures that `ffmpeg` or `ffprobe` are installed by downloading them from the internet if they cannot be found.

```

bash

npx remotion install ffmpeg
```

```

bash

npx remotion install ffmpeg
```

```

bash

npx remotion install ffprobe
```

```

bash

npx remotion install ffprobe
```

You might not need to call this function. Remotion will automatically download `ffmpeg` and `ffprobe` if a render is attempted, and no binary was found.

These commands are useful if you need `ffmpeg` and `ffprobe` to be ready before the first render is started.

## See also [​](\#see-also "Direct link to See also")

- [Node.JS equivalent: `ensureFfmpeg()`](/docs/renderer/ensure-ffmpeg)
- [Node.JS equivalent: `ensureFfprobe()`](/docs/renderer/ensure-ffprobe)
- [Installing FFmpeg](/docs/ffmpeg)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/install.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

- [See also](#see-also)
