---
title: Captionv4.0.216
source: Remotion Documentation
last_updated: 2024-12-22
---

# Captionv4.0.216

- [Home page](/)
- [@remotion/captions](/docs/captions/)
- Caption

On this page

# Caption [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216)

This is a simple data structure for a caption.

```

tsx

import type {Caption} from '@remotion/captions';

(alias) type Caption = {
    text: string;
    startMs: number;
    endMs: number;
    timestampMs: number | null;
    confidence: number | null;
}
import Caption
```

```

tsx

import type {Caption} from '@remotion/captions';

(alias) type Caption = {
    text: string;
    startMs: number;
    endMs: number;
    timestampMs: number | null;
    confidence: number | null;
}
import Caption
```

By establishing a standard data structure, we allow many operations that involve captions to be interoperable:

- **Transcribing**: Using the [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp) or [`@remotion/openai-whisper`](/docs/openai-whisper) packages
- **Formatting**: For example, creating pages using [`createTikTokStyleCaptions()`](/docs/captions/create-tiktok-style-captions)
- **Parsing**: Using the [`parseSrt()`](/docs/captions/parse-srt) function
- **Serializing**: For example to a `.srt` file using [`serializeSrt()`](/docs/captions/serialize-srt)

## Fields [​](\#fields "Direct link to Fields")

### `text` [​](\#text "Direct link to text")

The text of the caption.

### `startMs` [​](\#startms "Direct link to startms")

The start time of the caption in milliseconds.

### `endMs` [​](\#endms "Direct link to endms")

The end time of the caption in milliseconds.

### `timestampMs` [​](\#timestampms "Direct link to timestampms")

The timestamp of the caption as a singular timestamp in milliseconds.

When using [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp), this the `t_dtw` value.

Otherwise, it is not defined, but may be the average of the start and end timestamps.

### `confidence` [​](\#confidence "Direct link to confidence")

A number between 0 and 1 that indicates how confident the transcription is.

## Whitespace sensitivity [​](\#whitespace-sensitivity "Direct link to Whitespace sensitivity")

The `text` field is white-space sensitive. You should include spaces in it, as none are added automatically.

While rendering, apply the `white-space: pre` CSS property to the container of the caption to ensure that the spaces are preserved.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this type](https://github.com/remotion-dev/remotion/blob/main/packages/captions/src/caption.ts)
- [`@remotion/captions`](/docs/captions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/captions/caption.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/captions](/docs/captions/) [Next\
\
parseSrt()](/docs/captions/parse-srt)

- [Fields](#fields)
  - [`text`](#text)
  - [`startMs`](#startms)
  - [`endMs`](#endms)
  - [`timestampMs`](#timestampms)
  - [`confidence`](#confidence)
- [Whitespace sensitivity](#whitespace-sensitivity)
- [See also](#see-also)
