---
title: Convert a video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Convert a video

- [Home page](/)
- [WebCodecs](/docs/webcodecs/)
- Convert a video

On this page

# Convert a video

You can convert a video in the browser from one format to another using the [`convertMedia()`](/docs/webcodecs/convert-media) function from [`@remotion/webcodecs`](/docs/webcodecs).

💼 Important License Disclaimer

This package is licensed under the [Remotion License](/docs/license).

We consider a team of 4 or more people a "company".

**For "companies"**: A Remotion Company license needs to be obtained to use this package.

In a future version of`@remotion/webcodecs`, this package will also require the purchase of a newly created "WebCodecs Conversion Seat". [Get in touch](/contact) with us if you are planning to use this package.

**For individuals and teams up to 3:** You can use this package for free.

This is a short, non-binding explanation of our license. See the [License](/docs/license) itself for more details.

The following input formats are supported:

- ISO Base Media ( `.mp4`, `.mov`)
- Matroska ( `.mkv`, `.webm`)
- `.avi`
- MPEG Transport Stream ( `.ts`)

The following output formats are supported:

- MP4
- WebM
- WAV

The following output video codecs are supported:

- VP8 (WebM only)
- VP9 (WebM only)
- H.264 (MP4 only)

The following output audio codecs are supported:

- Opus (WebM only)
- AAC (MP4 only)
- PCM (WAV only)

## Installation [​](\#installation "Direct link to Installation")

Install the `@remotion/webcodecs` and `@remotion/media-parser` packages:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

npm i --save-exact @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

pnpm i @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

pnpm i @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

bun i @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

bun i @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

yarn --exact add @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

```

yarn --exact add @remotion/webcodecs@4.0.242 @remotion/media-parser@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

🚧 Unstable API

This package is experimental.

We might change the API at any time, until we remove this notice.

## Basic conversions [​](\#basic-conversions "Direct link to Basic conversions")

### Converting from an URL [​](\#converting-from-an-url "Direct link to Converting from an URL")

(needs to be CORS-enabled)

```

Converting an MP4 to a WebM
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

```

Converting an MP4 to a WebM
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
});
```

### Converting from a [`File`](https://developer.mozilla.org/en-US/docs/Web/API/File): [​](\#converting-from-a-file "Direct link to converting-from-a-file")

```

Converting a file
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// Get an actual file from an <input type="file"> element
const file = new File([], 'video.mp4');

await convertMedia({
  src: file,
  container: 'webm',
  reader: webFileReader,
});
```

```

Converting a file
tsx

import {convertMedia} from '@remotion/webcodecs';
import {webFileReader} from '@remotion/media-parser/web-file';

// Get an actual file from an <input type="file"> element
const file = new File([], 'video.mp4');

await convertMedia({
  src: file,
  container: 'webm',
  reader: webFileReader,
});
```

## Specifying the output codec [​](\#specifying-the-output-codec "Direct link to Specifying the output codec")

You can specify the output codec by passing the [`videoCodec`](/docs/webcodecs/convert-media#videocodec) and [`audioCodec`](/docs/webcodecs/convert-media#audiocodec) options:

```

Converting to VP9
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp9',
  audioCodec: 'opus',
});
```

```

Converting to VP9
tsx

import {convertMedia} from '@remotion/webcodecs';

await convertMedia({
  src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
  container: 'webm',
  videoCodec: 'vp9',
  audioCodec: 'opus',
});
```

## Advanced conversions [​](\#advanced-conversions "Direct link to Advanced conversions")

See [Track Transformation](/docs/webcodecs/track-transformation) for how you can get more control over the conversion.

## See also [​](\#see-also "Direct link to See also")

- [`convertMedia()`](/docs/webcodecs/convert-media)
- [Rotate a video](/docs/webcodecs/rotate-a-video)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/webcodecs/convert-a-video.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/webcodecs/) [Next\
\
Rotate a video](/docs/webcodecs/rotate-a-video)

- [Installation](#installation)
- [Basic conversions](#basic-conversions)
  - [Converting from an URL](#converting-from-an-url)
  - [Converting from a `File`:](#converting-from-a-file)
- [Specifying the output codec](#specifying-the-output-codec)
- [Advanced conversions](#advanced-conversions)
- [See also](#see-also)
