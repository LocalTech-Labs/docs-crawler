---
title: getAvailableEmoji()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getAvailableEmoji()

- [Home page](/)
- [@remotion/animated-emoji](/docs/animated-emoji/)
- getAvailableEmoji()

On this page

# getAvailableEmoji()

_available from v4.0.187_

Gets a list of available emoji that you can use with [`<AnimatedEmoji>`](/docs/animated-emoji/animated-emoji).

```

get-emoji.ts
tsx

import {getAvailableEmojis} from "@remotion/animated-emoji";

const emojiList = getAvailableEmojis();

console.log(emojiList);
```

```

get-emoji.ts
tsx

import {getAvailableEmojis} from "@remotion/animated-emoji";

const emojiList = getAvailableEmojis();

console.log(emojiList);
```

## Return value [​](\#return-value "Direct link to Return value")

An array of objects with the following properties:

### `name` [​](\#name "Direct link to name")

The name of the emoji. You can pass the name to the [`emoji`](/docs/animated-emoji/animated-emoji#emoji) prop.

### `categories` [​](\#categories "Direct link to categories")

An array of categories that the emoji belongs to.

### `tags` [​](\#tags "Direct link to tags")

An array of tags that the emoji has.

### `durationInSeconds` [​](\#durationinseconds "Direct link to durationinseconds")

The duration of the emoji in seconds.

### `codepoint` [​](\#codepoint "Direct link to codepoint")

The Unicode codepoint of the emoji.

## See also [​](\#see-also "Direct link to See also")

- [`<AnimatedEmoji>`](/docs/animated-emoji/animated-emoji)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/animated-emoji/get-available-emoji.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<AnimatedEmoji>](/docs/animated-emoji/animated-emoji) [Next\
\
Overview](/docs/media-parser/)

- [Return value](#return-value)
  - [`name`](#name)
  - [`categories`](#categories)
  - [`tags`](#tags)
  - [`durationInSeconds`](#durationinseconds)
  - [`codepoint`](#codepoint)
- [See also](#see-also)
