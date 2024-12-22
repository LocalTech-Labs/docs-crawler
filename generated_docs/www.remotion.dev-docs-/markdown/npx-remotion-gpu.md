---
title: npx remotion gpu
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion gpu

- [Home page](/)
- [Command line](/docs/cli/)
- gpu

On this page

# npx remotion gpu

_Available from Remotion v4.0.52_

Prints out how the Chrome browser uses the GPUs.

```

bash

npx remotion gpu --gl=angle
```

```

bash

npx remotion gpu --gl=angle
```

The command takes the same arguments for `--gl` as `npx remotion render` and also picks up the [`Config.setChromiumOpenGlRenderer()`](/docs/config#setchromiumopenglrenderer) option.

Try out different values to find which one is the best for your system.

```

Example output
bash

Canvas: Hardware accelerated
Canvas out-of-process rasterization: Enabled
Direct Rendering Display Compositor: Disabled
Compositing: Hardware accelerated
Multiple Raster Threads: Enabled
OpenGL: Enabled
Rasterization: Hardware accelerated
Raw Draw: Disabled
Skia Graphite: Disabled
Video Decode: Hardware accelerated
Video Encode: Hardware accelerated
WebGL: Hardware accelerated
WebGL2: Hardware accelerated
WebGPU: Hardware accelerated
```

```

Example output
bash

Canvas: Hardware accelerated
Canvas out-of-process rasterization: Enabled
Direct Rendering Display Compositor: Disabled
Compositing: Hardware accelerated
Multiple Raster Threads: Enabled
OpenGL: Enabled
Rasterization: Hardware accelerated
Raw Draw: Disabled
Skia Graphite: Disabled
Video Decode: Hardware accelerated
Video Encode: Hardware accelerated
WebGL: Hardware accelerated
WebGL2: Hardware accelerated
WebGPU: Hardware accelerated
```

The output should not be used for automated parsing, as it may change inbetween any Remotion and Chrome versions.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/cli/src/gpu.ts)
- [Using the GPU](/docs/gpu)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/gpu.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
ffprobe](/docs/cli/ffprobe) [Next\
\
help](/docs/cli/help)

- [See also](#see-also)
