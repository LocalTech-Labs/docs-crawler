---
title: fetchReader
source: Remotion Documentation
last_updated: 2024-12-22
---

# fetchReader

- [Home page](/)
- [@remotion/media-parser](/docs/media-parser/)
- fetchReader

On this page

# fetchReader

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

A reader for `@remotion/media-parser` that reads from a URL using `fetch()`.

It works in the browser, Node.js and Bun.

This is the default reader for `@remotion/media-parser`, so you don't need to explicitly specify it.

## Example [​](\#example "Direct link to Example")

```

Reading a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {fetchReader} from '@remotion/media-parser/fetch';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: fetchReader,
});
```

```

Reading a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {fetchReader} from '@remotion/media-parser/fetch';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: fetchReader,
});
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/fetch-reader.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
nodeReader](/docs/media-parser/node-reader) [Next\
\
webFileReader](/docs/media-parser/web-file-reader)

- [Example](#example)
