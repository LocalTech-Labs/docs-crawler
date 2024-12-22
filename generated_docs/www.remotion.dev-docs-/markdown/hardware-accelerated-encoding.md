---
title: Hardware accelerated encoding
source: Remotion Documentation
last_updated: 2024-12-22
---

# Hardware accelerated encoding

- [Home page](/)
- [Rendering](/docs/render)
- Hardware acceleration

On this page

# Hardware accelerated encoding

Encoding is the process of converting a sequence of images into a video file.

Besides rendering frames, encoding is one of the two steps required to create a video.

From Remotion v4.0.228, Remotion supports hardware-accelerated encoding in some cases.

Since encoding is platform- and codec-specific, only a few scenarios are supported at the moment.

- Currently, only macOS is supported (Acceleration using VideoToolbox)
- ProRes is supported from v4.0.228, H.264 and H.265 are supported from v4.0.236

## Enable hardware accelerated encoding [​](\#enable-hardware-accelerated-encoding "Direct link to Enable hardware accelerated encoding")

By default, hardware acceleration is `"disabled"`.

You can set the `hardwareAcceleration` option to `"if-possible"` to enable hardware acceleration if it is available.

If you want the render to fail if hardware acceleration is not possible, set the option to `"required"`.

### In SSR APIs [​](\#in-ssr-apis "Direct link to In SSR APIs")

Use the [`hardwareAcceleration`](/docs/renderer/render-media#hardwareacceleration) option in the [`renderMedia()`](/docs/renderer/render-media) function.

```

tsx

await renderMedia({
  composition,
  serveUrl,
  codec: 'prores',
  outputLocation,
  inputProps,
  hardwareAcceleration: 'if-possible',
});
```

```

tsx

await renderMedia({
  composition,
  serveUrl,
  codec: 'prores',
  outputLocation,
  inputProps,
  hardwareAcceleration: 'if-possible',
});
```

### In the CLI [​](\#in-the-cli "Direct link to In the CLI")

Use the [`--hardware-acceleration`](/docs/cli/render#--hardware-acceleration) option in the `npx remotion studio` command.

```

bash

npx remotion render MyComp --codec prores --hardware-acceleration if-possible
```

```

bash

npx remotion render MyComp --codec prores --hardware-acceleration if-possible
```

### In the Studio [​](\#in-the-studio "Direct link to In the Studio")

In the [Remotion Studio](/docs/studio), you can set the hardware acceleration option in the "Advanced" tab when rendering a video.

### With the config file [​](\#with-the-config-file "Direct link to With the config file")

You can set the [`setHardwareAcceleration()`](/docs/config#sethardwareacceleration) option in the [config file](/docs/config).

```

ts

import {Config} from '@remotion/cli/config';

Config.setHardwareAcceleration('if-possible');
```

```

ts

import {Config} from '@remotion/cli/config';

Config.setHardwareAcceleration('if-possible');
```

### In Remotion Lambda and Cloud Run [​](\#in-remotion-lambda-and-cloud-run "Direct link to In Remotion Lambda and Cloud Run")

These options are not supported in Remotion Lambda and Cloud Run, because those cloud services do not support hardware acceleration.

## File size [​](\#file-size "Direct link to File size")

Note that the file size is significantly larger by default when using hardware acceleration, likely because less compression is applied.

We recommend that you use the `--video-bitrate` flag to control the file size.

We find that `--video-bitrate=8M` achieves a similar file size than software encoding when exporting a H.264 Full HD video.

## Tell if hardware acceleration is being used [​](\#tell-if-hardware-acceleration-is-being-used "Direct link to Tell if hardware acceleration is being used")

[Run the render with verbose logging.](/docs/troubleshooting/debug-failed-render)
If the render is using hardware acceleration, you will see a log message like this:

```

Encoder: prores_videotoolbox, hardware accelerated: true
```

```

Encoder: prores_videotoolbox, hardware accelerated: true
```

Don't rely on the exact wording of the log message to determine if hardware acceleration is being used.

## See also [​](\#see-also "Direct link to See also")

- [Encoding guide](/docs/encoding)
- [Using the GPU](/docs/gpu)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/hardware-acceleration.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Setting video metadata](/docs/metadata) [Next\
\
Starting the Studio](/docs/studio/)

- [Enable hardware accelerated encoding](#enable-hardware-accelerated-encoding)
  - [In SSR APIs](#in-ssr-apis)
  - [In the CLI](#in-the-cli)
  - [In the Studio](#in-the-studio)
  - [With the config file](#with-the-config-file)
  - [In Remotion Lambda and Cloud Run](#in-remotion-lambda-and-cloud-run)
- [File size](#file-size)
- [Tell if hardware acceleration is being used](#tell-if-hardware-acceleration-is-being-used)
- [See also](#see-also)
