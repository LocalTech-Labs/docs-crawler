---
title: createTikTokStyleCaptions()v4.0.216
source: Remotion Documentation
last_updated: 2024-12-22
---

# createTikTokStyleCaptions()v4.0.216

- [Home page](/)
- [@remotion/captions](/docs/captions/)
- createTikTokStyleCaptions()

On this page

# createTikTokStyleCaptions() [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216)

Using this function, you can segment tokens to create "pages" of captions, as commonly seen on TikTok videos.

You may specify how often pages switch.
A high value for `combineTokensWithinMilliseconds` will fit many words on 1 page, while a low value will lead to word-by-word animation.

This function is safe to use in the browser, Node.js and Bun.

```

Create pages from captions
tsx

import {createTikTokStyleCaptions, Caption} from '@remotion/captions';

const captions: Caption[] = [
  {
    text: 'Using',
    startMs: 40,
    endMs: 300,
    timestampMs: 200,
    confidence: null,
  },
  {
    text: " Remotion's",
    startMs: 300,
    endMs: 900,
    timestampMs: 440,
    confidence: null,
  },
  {
    text: ' TikTok',
    startMs: 900,
    endMs: 1260,
    timestampMs: 1080,
    confidence: null,
  },
  {
    text: ' template,',
    startMs: 1260,
    endMs: 1950,
    timestampMs: 1600,
    confidence: null,
  },
];

const {pages} = createTikTokStyleCaptions({
  captions,
  combineTokensWithinMilliseconds: 1200,
});

/* pages: [
  {
    text: "Using Remotion's",
    startMs: 40,
    tokens: [
      {
        text: 'Using',
        fromMs: 40,
        toMs: 300,
      },
      {
        text: " Remotion's",
        fromMs: 300,
        toMs: 900,
      },
    ],
  },
  {
    text: 'TikTok template,',
    startMs: 900,
    tokens: [
      {
        text: 'TikTok',
        fromMs: 900,
        toMs: 1260,
      },
      {
        text: ' template,',
        fromMs: 1260,
        toMs: 1950,
      },
    ],
  }
] */
```

```

Create pages from captions
tsx

import {createTikTokStyleCaptions, Caption} from '@remotion/captions';

const captions: Caption[] = [
  {
    text: 'Using',
    startMs: 40,
    endMs: 300,
    timestampMs: 200,
    confidence: null,
  },
  {
    text: " Remotion's",
    startMs: 300,
    endMs: 900,
    timestampMs: 440,
    confidence: null,
  },
  {
    text: ' TikTok',
    startMs: 900,
    endMs: 1260,
    timestampMs: 1080,
    confidence: null,
  },
  {
    text: ' template,',
    startMs: 1260,
    endMs: 1950,
    timestampMs: 1600,
    confidence: null,
  },
];

const {pages} = createTikTokStyleCaptions({
  captions,
  combineTokensWithinMilliseconds: 1200,
});

/* pages: [
  {
    text: "Using Remotion's",
    startMs: 40,
    tokens: [
      {
        text: 'Using',
        fromMs: 40,
        toMs: 300,
      },
      {
        text: " Remotion's",
        fromMs: 300,
        toMs: 900,
      },
    ],
  },
  {
    text: 'TikTok template,',
    startMs: 900,
    tokens: [
      {
        text: 'TikTok',
        fromMs: 900,
        toMs: 1260,
      },
      {
        text: ' template,',
        fromMs: 1260,
        toMs: 1950,
      },
    ],
  }
] */
```

## API [​](\#api "Direct link to API")

### `captions` [​](\#captions "Direct link to captions")

An array of [`Caption`](/docs/captions/caption) objects.

### `combineTokensWithinMilliseconds` [​](\#combinetokenswithinmilliseconds "Direct link to combinetokenswithinmilliseconds")

Words that are closer than this value will be combined into a single page.

## Return value [​](\#return-value "Direct link to Return value")

An object with the following properties:

### `pages` [​](\#pages "Direct link to pages")

An array of `TikTokPage` objects.

A page consists of:

- `text`: The text of the page.
- `startMs`: The start time of the page in milliseconds.
- `tokens`: An array of objects, if you want to animate word-per-word:
  - `text`: The text of the token.
  - `fromMs`: The absolute start time of the token in milliseconds.
  - `toMs`: The absolute end time of the token in milliseconds.

## Whitespace sensitivity [​](\#whitespace-sensitivity "Direct link to Whitespace sensitivity")

The `text` field is white-space sensitive. You should include spaces in it, as none are added automatically.

While rendering, apply the `white-space: pre` CSS property to the container of the caption to ensure that the spaces are preserved.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/captions/src/create-tiktok-style-captions.ts)
- [`@remotion/captions`](/docs/captions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/captions/create-tiktok-style-captions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
serializeSrt()](/docs/captions/serialize-srt) [Next\
\
@remotion/install-whisper-cpp](/docs/install-whisper-cpp/)

- [API](#api)
  - [`captions`](#captions)
  - [`combineTokensWithinMilliseconds`](#combinetokenswithinmilliseconds)
- [Return value](#return-value)
  - [`pages`](#pages)
- [Whitespace sensitivity](#whitespace-sensitivity)
- [See also](#see-also)
