---
title: nodeReader
source: Remotion Documentation
last_updated: 2024-12-22
---

# nodeReader

- [Home page](/)
- [@remotion/media-parser](/docs/media-parser/)
- nodeReader

On this page

# nodeReader

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

A reader for `@remotion/media-parser` that reads from the local file system using Node.js `fs` module.

It also works with Bun.

## Example [​](\#example "Direct link to Example")

```

Reading a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});
```

```

Reading a local file
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/Users/jonnyburger/Downloads/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/node-reader.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
parseMedia()](/docs/media-parser/parse-media) [Next\
\
fetchReader](/docs/media-parser/fetch-reader)

- [Example](#example)
