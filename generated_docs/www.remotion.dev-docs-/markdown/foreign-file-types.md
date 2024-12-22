---
title: Foreign file types
source: Remotion Documentation
last_updated: 2024-12-22
---

# Foreign file types

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- Foreign file types

On this page

# Foreign file types

If you pass a file that is not one of the supported formats, [`parseMedia()`](/docs/media-parser/parse-media) will throw an error.

Sometimes this error can contain useful information about the file type nonetheless.

This is useful if you allow users to upload an arbitrary file and want to get information about it with just one pass.

```

Get information about a video, or get the file type if it's not a video
tsx

import {
  parseMedia,
  IsAnImageError,
  IsAGifError,
  IsAPdfError,
  IsAnUnsupportedFileTypeError,
  IsAnUnsupportedAudioTypeError,
} from '@remotion/media-parser';

try {
  await parseMedia({
    src: 'https://example.com/my-video.png',
    fields: {},
  });
} catch (e) {
  if (e instanceof IsAnImageError) {
    console.log(
      'The file is an image of format:',
      e.imageType,
      'dimensions:',
      e.dimensions,
    );
  } else if (e instanceof IsAGifError) {
    console.log('The file is a GIF');
  } else if (e instanceof IsAPdfError) {
    console.log('The file is a PDF');
  } else if (e instanceof IsAnUnsupportedAudioTypeError) {
    console.log('The file is an audio file of format:', e.audioType);
  } else if (e instanceof IsAnUnsupportedFileTypeError) {
    console.log('The file is of an unsupported type');
  }
}
```

```

Get information about a video, or get the file type if it's not a video
tsx

import {
  parseMedia,
  IsAnImageError,
  IsAGifError,
  IsAPdfError,
  IsAnUnsupportedFileTypeError,
  IsAnUnsupportedAudioTypeError,
} from '@remotion/media-parser';

try {
  await parseMedia({
    src: 'https://example.com/my-video.png',
    fields: {},
  });
} catch (e) {
  if (e instanceof IsAnImageError) {
    console.log(
      'The file is an image of format:',
      e.imageType,
      'dimensions:',
      e.dimensions,
    );
  } else if (e instanceof IsAGifError) {
    console.log('The file is a GIF');
  } else if (e instanceof IsAPdfError) {
    console.log('The file is a PDF');
  } else if (e instanceof IsAnUnsupportedAudioTypeError) {
    console.log('The file is an audio file of format:', e.audioType);
  } else if (e instanceof IsAnUnsupportedFileTypeError) {
    console.log('The file is of an unsupported type');
  }
}
```

## Available errors [​](\#available-errors "Direct link to Available errors")

### `IsAnImageError` [​](\#isanimageerror "Direct link to isanimageerror")

An error if the image is a `png`, `jpeg`, `bmp` or `webp`.

- `fileName`: The name of the file - from file name or `Content-Disposition` header.
- `sizeInBytes`: The size of the file in bytes.
- `mimeType`: The MIME type of the file.
- `imageType`: The type of the image ( `png`, `jpeg`, `bmp` or `webp`).
- `dimensions`: The dimensions of the image ( `width` and `height`).

### `IsAnUnsupportedAudioTypeError` [​](\#isanunsupportedaudiotypeerror "Direct link to isanunsupportedaudiotypeerror")

An error if the file is an audio file that `@remotion/media-parser` does not yet support.

- `fileName`: The name of the file - from file name or `Content-Disposition` header.
- `sizeInBytes`: The size of the file in bytes.
- `mimeType`: The MIME type of the file.
- `audioType`: The type of the audio file - one of `mp3`, `wav` and `aac`.

note

`@remotion/media-parser` aims to support these audio formats in the future, in which case they will not throw an error but return audio tracks.

### `IsAGifError` [​](\#isagiferror "Direct link to isagiferror")

An error if the image is a GIF.

- `fileName`: The name of the file - from file name or `Content-Disposition` header.
- `sizeInBytes`: The size of the file in bytes.
- `mimeType`: The MIME type of the file.

### `IsAPdfError` [​](\#isapdferror "Direct link to isapdferror")

An error if the file is a PDF.

- `fileName`: The name of the file - from file name or `Content-Disposition` header.
- `sizeInBytes`: The size of the file in bytes.
- `mimeType`: The MIME type of the file.

### `IsAnUnsupportedFileTypeError` [​](\#isanunsupportedfiletypeerror "Direct link to isanunsupportedfiletypeerror")

An error if the file is of a type that `@remotion/media-parser` does not recognize.

- `fileName`: The name of the file - from file name or `Content-Disposition` header.
- `sizeInBytes`: The size of the file in bytes.
- `mimeType`: The MIME type of the file.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/foreign-file-types.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
WebCodecs](/docs/media-parser/webcodecs) [Next\
\
Overview](/docs/webcodecs/)

- [Available errors](#available-errors)
  - [`IsAnImageError`](#isanimageerror)
  - [`IsAnUnsupportedAudioTypeError`](#isanunsupportedaudiotypeerror)
  - [`IsAGifError`](#isagiferror)
  - [`IsAPdfError`](#isapdferror)
  - [`IsAnUnsupportedFileTypeError`](#isanunsupportedfiletypeerror)
