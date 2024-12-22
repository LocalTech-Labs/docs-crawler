---
title: Installing FFmpeg
source: Remotion Documentation
last_updated: 2024-12-22
---

# Installing FFmpeg

On this page

# Installing FFmpeg

info

Since Remotion v4.0, Remotion comes bundled with a lightweight version of FFmpeg. An installation of FFmpeg is no longer needed.

## FFmpeg in V3 of Remotion [​](\#ffmpeg-in-v3-of-remotion "Direct link to FFmpeg in V3 of Remotion")

**The following documentation is an archival for how FFmpeg worked in v3.0.**

Remotion requires FFmpeg to encode videos. Since v3.3, you do not need to install FFmpeg manually. This page documents the behavior of Remotion for developers needing advanced control.

### `ffmpeg` and `ffprobe` [​](\#ffmpeg-and-ffprobe "Direct link to ffmpeg-and-ffprobe")

Two binaries are required for Remotion: `ffmpeg` and `ffprobe`. When talking about FFmpeg in the documentation, it may also refer to FFprobe.

### Auto-install [​](\#auto-install "Direct link to Auto-install")

When rendering a video and binaries are not found, Remotion will download them from the internet and put it inside your `node_modules` folder. The binary will not get added to your `PATH`, so if you type in `ffmpeg` into your Terminal, it may not be found. However, Remotion will be able to use it

#### Supported architectures [​](\#supported-architectures "Direct link to Supported architectures")

Auto-install is supported on the following platforms:

- Linux, x86\_64,
- macOS, Intel
- macOS, Apple Silicon
- Windows, x86\_64

For other platforms, you need to supply your own binaries.

#### Triggering auto-install [​](\#triggering-auto-install "Direct link to Triggering auto-install")

By rendering a video, the download of FFmpeg will be triggered automatically.

On servers, it might be of use to install the binaries before the first render, so no time is wasted once the first render begins.

- Using the CLI, you can run [`npx remotion install ffmpeg`](/docs/cli/install) and `npx remotion install ffprobe` to trigger auto-install of binaries. If the binaries exist, the command will do nothing. This requires `@remotion/cli` to be installed.
- The [`@remotion/renderer`](/docs/renderer) package exposes [`ensureFfmpeg()`](/docs/renderer/ensure-ffmpeg) and [`ensureFfprobe()`](/docs/renderer/ensure-ffprobe) functions

### Order of priority [​](\#order-of-priority "Direct link to Order of priority")

In case of multiple binaries being supplied, they priority order is the following:

- If a binary was supplied using the `ffmpegExecutable` or `ffprobeExecutable` option, it will be used.
- If `ffmpeg` or `ffprobe` is in the `PATH`, it will be used.
- If a binary was previously installed by Remotion into `node_modules`, it will be used.
- If a binary can be downloaded from the internet, Remotion will do so and use it.
- Failure if no binary was found using the logic above.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/ffmpeg.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

- [FFmpeg in V3 of Remotion](#ffmpeg-in-v3-of-remotion)
  - [`ffmpeg` and `ffprobe`](#ffmpeg-and-ffprobe)
  - [Auto-install](#auto-install)
  - [Order of priority](#order-of-priority)