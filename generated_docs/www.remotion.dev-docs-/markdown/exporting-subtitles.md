---
title: Exporting subtitles
source: Remotion Documentation
last_updated: 2024-12-22
---

# Exporting subtitles

- [Home page](/)
- Recorder
- Exporting
- Subtitles

On this page

# Exporting subtitles

## Non-landscape layout [​](\#non-landscape-layout "Direct link to Non-landscape layout")

If you export a video in the `square` layout, then the subtitles will be burned into the video.

The idea behind this is that the square videos are posted on platforms like x.com and LinkedIn, where the video appears in a feed, muted by default.

It is important that the subtitles are always visible and have a large font size.

## Landscape layout [​](\#landscape-layout "Direct link to Landscape layout")

If you export a video in the `landscape` layout, then the subtitles will be by default exported as a `.srt` file, located in `out/[composition-id]/captions.srt`.

This is caused by the use of the [`<Artifact>`](/docs/artifact) API in [`remotion/captions/srt/EmitSrtFile.tsx`](https://github.com/remotion-dev/recorder/blob/main/remotion/captions/srt/EmitSrtFile.tsx).

When posting to YouTube, you should upload the SRT file alongside your video, so that viewers can use the native YouTube feature to toggle the subtitles and adjust the font size.

## Burn subtitles [​](\#burn-subtitles "Direct link to Burn subtitles")

If you want to burn the subtitles into the video in landscape layout, edit the [`remotion/captions/srt/SrtPreviewAndEditor/SrtPreviewAndEditor.tsx`](https://github.com/remotion-dev/recorder/blob/main/remotion/captions/srt/SrtPreviewAndEditor/SrtPreviewAndEditor.tsx) file and delete the following lines:

```

remotion/captions/srt/SrtPreviewAndEditor/SrtPreviewAndEditor.tsx
diff

-  // During rendering, you will get the actual .srt file instead of the preview.
-  if (getRemotionEnvironment().isRendering) {
-    return null;
-  }
```

```

remotion/captions/srt/SrtPreviewAndEditor/SrtPreviewAndEditor.tsx
diff

-  // During rendering, you will get the actual .srt file instead of the preview.
-  if (getRemotionEnvironment().isRendering) {
-    return null;
-  }
```

Now your video will include the subtitles as an overlay on top of the video.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/exporting-subtitles.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Render on Lambda](/docs/recorder/lambda-rendering) [Next\
\
Source Control](/docs/recorder/source-control)

- [Non-landscape layout](#non-landscape-layout)
- [Landscape layout](#landscape-layout)
- [Burn subtitles](#burn-subtitles)
