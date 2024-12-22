---
title: @remotion/captions
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/captions

- [Home page](/)
- @remotion/captions

On this page

# @remotion/captions

_Available from v4.0.216_

The `@remotion/captions` package provides utilities for dealing with subtitles.

At the centre of this caption is the [`Caption`](/docs/captions/caption) type, which defines a standard shape for captions from different sources.

Captions generated from [`@remotion/install-whisper-cpp`](/docs/install-whisper-cpp) [can be converted](/docs/install-whisper-cpp/to-captions) into the `Caption` type.

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/captions@4.0.242Copy
```

```

npm i --save-exact @remotion/captions@4.0.242Copy
```

```

pnpm i @remotion/captions@4.0.242Copy
```

```

pnpm i @remotion/captions@4.0.242Copy
```

```

bun i @remotion/captions@4.0.242Copy
```

```

bun i @remotion/captions@4.0.242Copy
```

```

yarn --exact add @remotion/captions@4.0.242Copy
```

```

yarn --exact add @remotion/captions@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## APIs [​](\#apis "Direct link to APIs")

[**Caption** \
\
An object shape for captions](/docs/captions/caption) [**parseSrt()** \
\
Parse a .srt file into a `Caption` array](/docs/captions/parse-srt) [**serializeSrt()** \
\
Serialize a .srt file into a `Caption` array](/docs/captions/serialize-srt) [**createTikTokStyleCaptions()** \
\
Structure the captions for TikTok-style display](/docs/captions/create-tiktok-style-captions)

## License [​](\#license "Direct link to License")

MIT

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/captions/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getAvailableVideoCodecs()](/docs/webcodecs/get-available-video-codecs) [Next\
\
Caption](/docs/captions/caption)

- [APIs](#apis)
- [License](#license)
