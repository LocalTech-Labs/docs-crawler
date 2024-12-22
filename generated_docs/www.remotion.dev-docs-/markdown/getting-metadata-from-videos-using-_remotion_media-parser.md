---
title: Getting metadata from videos using @remotion/media-parser
source: Remotion Documentation
last_updated: 2024-12-22
---

# Getting metadata from videos using @remotion/media-parser

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- Getting video metadata

On this page

# Getting metadata from videos using @remotion/media-parser

Using [`@remotion/media-parser`](/docs/media-parser), you can get metadata from a variety of video formats: MP4, MKV, AVI, WebM and MOV.

## Getting metadata from a URL [​](\#getting-metadata-from-a-url "Direct link to Getting metadata from a URL")

Specify the [`fields`](/docs/media-parser/parse-media#fields) you would like to read and the URl as [`src`](/docs/media-parser/parse-media#src).

This works in the browser as well as in Node.js and Bun.

```

Parsing a hosted video
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a hosted video
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

## Getting metadata from a local path [​](\#getting-metadata-from-a-local-path "Direct link to Getting metadata from a local path")

Use [`nodeReader`](/docs/media-parser/node-reader) to read a file from a filesystem.

This can be used in Node.js and Bun.

```

Parsing a video from a file path
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/tmp/video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a video from a file path
tsx

import {parseMedia} from '@remotion/media-parser';
import {nodeReader} from '@remotion/media-parser/node';

const result = await parseMedia({
  src: '/tmp/video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: nodeReader,
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

## Getting metadata from a `File` [​](\#getting-metadata-from-a-file "Direct link to getting-metadata-from-a-file")

If you take user uploads in the browser, they will be in the form of a [`File`](https://developer.mozilla.org/en-US/docs/Web/API/File).

Use [`webFileReader`](/docs/media-parser/web-file-reader) to read a video from a [`File`](https://developer.mozilla.org/en-US/docs/Web/API/File).

```

Parsing a video from a file path
tsx

import {parseMedia} from '@remotion/media-parser';
import {webFileReader} from '@remotion/media-parser/web-file';

// You would get this from a `<input type="file">`
const file = new File([], 'video.mp4');

const result = await parseMedia({
  src: file,
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: webFileReader,
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a video from a file path
tsx

import {parseMedia} from '@remotion/media-parser';
import {webFileReader} from '@remotion/media-parser/web-file';

// You would get this from a `<input type="file">`
const file = new File([], 'video.mp4');

const result = await parseMedia({
  src: file,
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  reader: webFileReader,
});

console.log(result.durationInSeconds); // 10
console.log(result.dimensions); // {width: 1920, height: 1080}
```

## Available fields [​](\#available-fields "Direct link to Available fields")

You can read the duration, the dimensions, the framerate, the container format, the codecs, the ID3 tags and more information from a video file.

See [Fields](/docs/media-parser/parse-media#fields) for a complete list.

## Getting metadata as soon as possible [​](\#getting-metadata-as-soon-as-possible "Direct link to Getting metadata as soon as possible")

```

Parsing a video from a File
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  onDurationInSeconds: (durationInSeconds) => {
    console.log(durationInSeconds); // 10
  },
});

console.log(result.dimensions); // {width: 1920, height: 1080}
```

```

Parsing a video from a File
tsx

import {parseMedia} from '@remotion/media-parser';

const result = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    durationInSeconds: true,
    dimensions: true,
  },
  onDurationInSeconds: (durationInSeconds) => {
    console.log(durationInSeconds); // 10
  },
});

console.log(result.dimensions); // {width: 1920, height: 1080}
```

## See also [​](\#see-also "Direct link to See also")

- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/media-parser/) [Next\
\
Fast and slow operations](/docs/media-parser/fast-and-slow)

- [Getting metadata from a URL](#getting-metadata-from-a-url)
- [Getting metadata from a local path](#getting-metadata-from-a-local-path)
- [Getting metadata from a `File`](#getting-metadata-from-a-file)
- [Available fields](#available-fields)
- [Getting metadata as soon as possible](#getting-metadata-as-soon-as-possible)
- [See also](#see-also)
