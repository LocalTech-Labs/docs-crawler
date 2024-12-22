---
title: @remotion/media-parser
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/media-parser

- [Home page](/)
- Media Parser

On this page

# @remotion/media-parser

_available from v4.0.190_

This is an experimental package that parses video and other media files in order to extract relevant metadata.

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

Design goals:

- Support as many relevant containers as possible - currently: `.mp4`, `.mov`, `.webm`, `.mkv`, `.avi`, `.ts`
- [Work in the browser, and server runtimes (Node.js, Bun, Edge, etc.)](/docs/media-parser/support)
- [Enable powerful WebCodecs use cases](/docs/media-parser/webcodecs)
- [Fine-grained querying, only download as much data as necessary](/docs/media-parser/fast-and-slow)
- [Functional TypeScript API](/docs/media-parser/parse-media)
- [Be useful when passing unsupported media](/docs/media-parser/foreign-file-types)
- Solve problems that are relevant for Remotion users

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/media-parser@4.0.242Copy
```

```

npm i --save-exact @remotion/media-parser@4.0.242Copy
```

```

pnpm i @remotion/media-parser@4.0.242Copy
```

```

pnpm i @remotion/media-parser@4.0.242Copy
```

```

bun i @remotion/media-parser@4.0.242Copy
```

```

bun i @remotion/media-parser@4.0.242Copy
```

```

yarn --exact add @remotion/media-parser@4.0.242Copy
```

```

yarn --exact add @remotion/media-parser@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## Guide [​](\#guide "Direct link to Guide")

[**Getting video metadata** \
\
Simple examples of extracting video metadata](/docs/media-parser/metadata) [**Fast and slow operations** \
\
Efficently use parseMedia()](/docs/media-parser/fast-and-slow) [**Extract ID3 tags and EXIF data** \
\
Get embedded tags from video files](/docs/media-parser/tags) [**Runtime supports** \
\
Where you can run Media Parser](/docs/media-parser/support) [**WebCodecs** \
\
How Media Parser integrates with WebCodecs](/docs/media-parser/webcodecs)

## APIs [​](\#apis "Direct link to APIs")

The following APIs are available:

[**parseMedia()** \
\
Parse a media file.](/docs/media-parser/parse-media) [**nodeReader** \
\
Read a file from the local file system.](/docs/media-parser/node-reader) [**fetchReader** \
\
Read a file using `fetch()`.](/docs/media-parser/fetch-reader) [**webFileReader** \
\
Read a file from `<input type="file">`.](/docs/media-parser/web-file-reader)

## License [​](\#license "Direct link to License")

[Remotion License](https://remotion.dev/license)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Uninstall Cloud Run](/docs/cloudrun/uninstall) [Next\
\
Getting video metadata](/docs/media-parser/metadata)

- [Installation](#installation)
- [Guide](#guide)
- [APIs](#apis)
- [License](#license)
