---
title: parseSrt()v4.0.216
source: Remotion Documentation
last_updated: 2024-12-22
---

# parseSrt()v4.0.216

- [Home page](/)
- [@remotion/captions](/docs/captions/)
- parseSrt()

On this page

# parseSrt() [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216)

Parses the contents of a SubRip file ( `.srt`) and returns an array of [`Caption`](/docs/captions) items.

```

Example usage
tsx

import {parseSrt} from '@remotion/captions';

const input = `
1
00:00:00,000 --> 00:00:02,500
Welcome to the Example Subtitle File!

2
00:00:03,000 --> 00:00:06,000
This is a demonstration of SRT subtitles.

3
00:00:07,000 --> 00:00:10,500
You can use SRT files to add subtitles to your videos.
`.trim();

const {captions} = parseSrt({input});

/* captions = [
  {
    confidence: 1,
    endMs: 2500,
    startMs: 0,
    text: 'Welcome to the Example Subtitle File!',
    timestampMs: 1250,
  },
  {
    confidence: 1,
    endMs: 6000,
    startMs: 3000,
    text: 'This is a demonstration of SRT subtitles.',
    timestampMs: 4500,
  },
  {
    confidence: 1,
    endMs: 10500,
    startMs: 7000,
    text: 'You can use SRT files to add subtitles to your videos.',
    timestampMs: 8750,
  },
]
*/
```

```

Example usage
tsx

import {parseSrt} from '@remotion/captions';

const input = `
1
00:00:00,000 --> 00:00:02,500
Welcome to the Example Subtitle File!

2
00:00:03,000 --> 00:00:06,000
This is a demonstration of SRT subtitles.

3
00:00:07,000 --> 00:00:10,500
You can use SRT files to add subtitles to your videos.
`.trim();

const {captions} = parseSrt({input});

/* captions = [
  {
    confidence: 1,
    endMs: 2500,
    startMs: 0,
    text: 'Welcome to the Example Subtitle File!',
    timestampMs: 1250,
  },
  {
    confidence: 1,
    endMs: 6000,
    startMs: 3000,
    text: 'This is a demonstration of SRT subtitles.',
    timestampMs: 4500,
  },
  {
    confidence: 1,
    endMs: 10500,
    startMs: 7000,
    text: 'You can use SRT files to add subtitles to your videos.',
    timestampMs: 8750,
  },
]
*/
```

## API [​](\#api "Direct link to API")

### `input` [​](\#input "Direct link to input")

The contents of a `.srt` file as a `string`.

## Return value [​](\#return-value "Direct link to Return value")

An object with the following properties:

### `captions` [​](\#captions "Direct link to captions")

An array of [`Caption`](/docs/captions/caption) items.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/captions/src/parse-srt.ts)
- [`@remotion/captions`](/docs/captions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/captions/parse-srt.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Caption](/docs/captions/caption) [Next\
\
serializeSrt()](/docs/captions/serialize-srt)

- [API](#api)
  - [`input`](#input)
- [Return value](#return-value)
  - [`captions`](#captions)
- [See also](#see-also)
