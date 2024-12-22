---
title: serializeSrt()v4.0.216
source: Remotion Documentation
last_updated: 2024-12-22
---

# serializeSrt()v4.0.216

- [Home page](/)
- [@remotion/captions](/docs/captions/)
- serializeSrt()

On this page

# serializeSrt() [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216)

Converts a two-dimensional array of [`Caption`](/docs/captions/caption) items into a string in the SubRip format ( `.srt`).

```

Example usage
tsx

import {serializeSrt, Caption} from '@remotion/captions';

const captions: Caption[] = [
  {
    text: 'Welcome to the Example Subtitle File!',
    startMs: 0,
    endMs: 2500,
    timestampMs: 1250,
    confidence: 1,
  },
  {
    text: 'This is a demonstration of SRT subtitles.',
    startMs: 3000,
    endMs: 6000,
    timestampMs: 4500,
    confidence: 1,
  },
  {
    text: 'You can use SRT files to add subtitles to your videos.',
    startMs: 7000,
    endMs: 10500,
    timestampMs: 8750,
    confidence: 1,
  },
];

const lines = captions.map((caption) => [caption]);

const serialized = serializeSrt({lines});

/* serialized = `1
00:00:00,000 --> 00:00:02,500
Welcome to the Example Subtitle File!

2
00:00:03,000 --> 00:00:06,000
This is a demonstration of SRT subtitles.

3
00:00:07,000 --> 00:00:10,500
You can use SRT files to add subtitles to your videos.
`
*/
```

```

Example usage
tsx

import {serializeSrt, Caption} from '@remotion/captions';

const captions: Caption[] = [
  {
    text: 'Welcome to the Example Subtitle File!',
    startMs: 0,
    endMs: 2500,
    timestampMs: 1250,
    confidence: 1,
  },
  {
    text: 'This is a demonstration of SRT subtitles.',
    startMs: 3000,
    endMs: 6000,
    timestampMs: 4500,
    confidence: 1,
  },
  {
    text: 'You can use SRT files to add subtitles to your videos.',
    startMs: 7000,
    endMs: 10500,
    timestampMs: 8750,
    confidence: 1,
  },
];

const lines = captions.map((caption) => [caption]);

const serialized = serializeSrt({lines});

/* serialized = `1
00:00:00,000 --> 00:00:02,500
Welcome to the Example Subtitle File!

2
00:00:03,000 --> 00:00:06,000
This is a demonstration of SRT subtitles.

3
00:00:07,000 --> 00:00:10,500
You can use SRT files to add subtitles to your videos.
`
*/
```

## API [​](\#api "Direct link to API")

### `lines` [​](\#lines "Direct link to lines")

An two-dimensional array of [`Caption`](/docs/captions/caption) items.

Each top-level item represents a line in the SubRip file.

The second-level items represent the words in that line.

Words get concatenated together during serialization. No spaces are added between the words.

The start timestamp is determined from the `startMs` value of the first word in the line.

The end timestamp is determined from the `endMs` value of the last word in the line.

Arrays with no items will be ignored.

## Return value [​](\#return-value "Direct link to Return value")

A string in the SubRip format ( `.srt`).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/captions/src/serialize-srt.ts)
- [`@remotion/captions`](/docs/captions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/captions/serialize-srt.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
parseSrt()](/docs/captions/parse-srt) [Next\
\
createTikTokStyleCaptions()](/docs/captions/create-tiktok-style-captions)

- [API](#api)
  - [`lines`](#lines)
- [Return value](#return-value)
- [See also](#see-also)
