---
title: ensureFfprobe()
source: Remotion Documentation
last_updated: 2024-12-22
---

# ensureFfprobe()

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- ensureFfprobe()

On this page

# ensureFfprobe()

_Available from v3.3, removed from v4.0_

warning

This API has been removed in v4.0 and is not necessary to call anymore. This page remains for archival purposes.

Checks if the `ffprobe` binary is installed and if it is not, [downloads it and puts it into your `node_modules` folder](/docs/ffmpeg).

```

ensure.mjs
ts

import { ensureFfprobe } from "@remotion/renderer";

await ensureFfprobe();
```

```

ensure.mjs
ts

import { ensureFfprobe } from "@remotion/renderer";

await ensureFfprobe();
```

You might not need to call this function. Remotion will automatically download `ffprobe` if a render is attempted, and no binary was found.

This function is useful if you need `ffprobe` to be ready before the first render is started.

Also call [`ensureFfmpeg()`](/docs/renderer/ensure-ffmpeg) to get both binaries that Remotion requires.

## Options [​](\#options "Direct link to Options")

Optionally, you can pass an object and pass the following options:

### `remotionRoot` [​](\#remotionroot "Direct link to remotionroot")

_string_

The directory in which your `node_modules` is located.

## Return value [​](\#return-value "Direct link to Return value")

A promise which resolves an object with the following properties:

- `wasAlreadyInstalled`: Boolean whether the binary was downloaded because of this function call.
- `result`: A string, either `found-in-path`, `found-in-node-modules` or `installed`.

## Exceptions [​](\#exceptions "Direct link to Exceptions")

This function throws if no binary was found, the download fails or no binaries are available for your platform.

## See also [​](\#see-also "Direct link to See also")

- CLI equivalent: [`npx remotion install ffprobe`](/docs/cli/install)
- [`ensureFfmpeg()`](/docs/renderer/ensure-ffmpeg)
- [Installing FFmpeg](/docs/ffmpeg)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/ensure-ffprobe.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
ensureFfmpeg()](/docs/renderer/ensure-ffmpeg) [Next\
\
getCanExtractFramesFast()](/docs/renderer/get-can-extract-frames-fast)

- [Options](#options)
  - [`remotionRoot`](#remotionroot)
- [Return value](#return-value)
- [Exceptions](#exceptions)
- [See also](#see-also)
