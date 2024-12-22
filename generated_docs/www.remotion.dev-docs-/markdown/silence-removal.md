---
title: Silence removal
source: Remotion Documentation
last_updated: 2024-12-22
---

# Silence removal

- [Home page](/)
- Recorder
- Editing
- Silence removal

On this page

# Silence removal

When captions are generated, the Whisper.cpp timestamps are taken to determine the start and end of the transcription.

Therefore, you should make sure captions are generated before you manually trim the video.

## Edit the transcript to fix mistakes [​](\#edit-the-transcript-to-fix-mistakes "Direct link to Edit the transcript to fix mistakes")

The easiest way to fix mistakes is to [click on a caption](/docs/recorder/editing/captions) and edit the transcript.

The first non-empty token determines the start trim.

The last non-empty token determines the end trim.

## Apply an offset [​](\#apply-an-offset "Direct link to Apply an offset")

In the right hand side default props editor, apply a `startOffset` to the scene.

The offset is in frames.

A negative offset will start the video earlier, a positive offset will start the video later.

To change the end trim, apply a `endOffset` to the scene. It follows the same logic as the start offset.

## Customizing the logic [​](\#customizing-the-logic "Direct link to Customizing the logic")

The logic for determining the start and end timestamp is in [`remotion/calculate-metadata/add-metadata-to-scene.ts`](https://github.com/remotion-dev/recorder/blob/main/remotion/calculate-metadata/add-metadata-to-scene.ts#L56).

The default start padding is a 0.25 seconds and the default end padding is a 0.5 seconds.

These values are defined in [`remotion/calculate-metadata/get-start-end-frame.ts`](https://github.com/remotion-dev/recorder/blob/main/remotion/calculate-metadata/get-start-end-frame.ts#L8)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/silence-removal.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Start editing](/docs/recorder/editing/) [Next\
\
Captions](/docs/recorder/editing/captions)

- [Edit the transcript to fix mistakes](#edit-the-transcript-to-fix-mistakes)
- [Apply an offset](#apply-an-offset)
- [Customizing the logic](#customizing-the-logic)
