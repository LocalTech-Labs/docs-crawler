---
title: ENAMETOOLONG
source: Remotion Documentation
last_updated: 2024-12-22
---

# ENAMETOOLONG

- [Home page](/)
- Troubleshooting
- ENAMETOOLONG

On this page

# ENAMETOOLONG

```

Command failed with ENAMETOOLONG: ffmpeg ...
```

```

Command failed with ENAMETOOLONG: ffmpeg ...
```

This error occurs if Remotion is used on Windows and too many audio layers are in the video so that an FFmpeg command gets generated that is longer than the maximum allowed Windows command length (8192 characters).

Unfortunately, FFmpeg does not have any other alternative way to specify a massive amount of inputs, so Remotion can not fix this issue. Here are some recommendations instead:

## Mute videos that have no sound [​](\#mute-videos-that-have-no-sound "Direct link to Mute videos that have no sound")

If you have videos without sound, add the `muted` property, so they won't be added to the mix.

## Use a different operating system [​](\#use-a-different-operating-system "Direct link to Use a different operating system")

macOS and Linux have a much longer maximum command length. Render the same video on a different operating system or use Linux for Windows Subsystem.

## Render the video partially [​](\#render-the-video-partially "Direct link to Render the video partially")

Render only a portion of a video using the [`--frames`](/docs/cli/render#--frames) property, and add portions together using an [FFmpeg concat command](https://stackoverflow.com/a/11175851)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/enametoolong.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Version mismatch](/docs/version-mismatch) [Next\
\
Slow method to extract frame](/docs/slow-method-to-extract-frame)

- [Mute videos that have no sound](#mute-videos-that-have-no-sound)
- [Use a different operating system](#use-a-different-operating-system)
- [Render the video partially](#render-the-video-partially)
