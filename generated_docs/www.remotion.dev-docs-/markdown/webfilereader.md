---
title: webFileReader
source: Remotion Documentation
last_updated: 2024-12-22
---

# webFileReader

- [Home page](/)
- [@remotion/media-parser](/docs/media-parser/)
- webFileReader

On this page

# webFileReader

warning

**Unstable API**: This package is experimental. We might change the API at any time, until we remove this notice.

A reader for `@remotion/media-parser` that reads from a [`File`](https://developer.mozilla.org/en-US/docs/Web/API/File) object.

It only works in the browser.

Useful if you want to parse a video from a `<input type="file">` element.

## Example [​](\#example "Direct link to Example")

```

Reading a local file
tsx

import React, {useCallback} from 'react';
import {parseMedia} from '@remotion/media-parser';
import {webFileReader} from '@remotion/media-parser/web-file';

const MyComp: React.FC = () => {
  const onFile = useCallback(async (file: File) => {
    const result = await parseMedia({
      src: file,
      fields: {
        durationInSeconds: true,
        dimensions: true,
      },
      reader: webFileReader,
    });
  }, []);

  return (
    <input type="file" onChange={(e) => onFile(e.target.files?.[0]!)} />
  );
}
```

```

Reading a local file
tsx

import React, {useCallback} from 'react';
import {parseMedia} from '@remotion/media-parser';
import {webFileReader} from '@remotion/media-parser/web-file';

const MyComp: React.FC = () => {
  const onFile = useCallback(async (file: File) => {
    const result = await parseMedia({
      src: file,
      fields: {
        durationInSeconds: true,
        dimensions: true,
      },
      reader: webFileReader,
    });
  }, []);

  return (
    <input type="file" onChange={(e) => onFile(e.target.files?.[0]!)} />
  );
}
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/web-file-reader.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
fetchReader](/docs/media-parser/fetch-reader) [Next\
\
Overview](/docs/webcodecs/)

- [Example](#example)
