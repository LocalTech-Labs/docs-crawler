---
title: getStaticFiles()
source: Remotion Documentation
last_updated: 2024-12-22
---

# getStaticFiles()

- [Home page](/)
- [remotion](/docs/remotion)
- getStaticFiles()

On this page

# getStaticFiles()

_Available from v3.3.26._

note

This API is being moved to the `@remotion/studio` package. Prefer importing the API from [`@remotion/studio`](/docs/studio/get-static-files) instead of `remotion`.

Gets an array containing all files in the `public/` folder.

You can reference them by using [`staticFile()`](/docs/staticfile).

warning

This feature _only_ works in Remotion Studio and during rendering, otherwise it returns an empty array.

note

On Linux, watching for changes in subdirectories is only supported from Node.js v19.1.0. If you use a version earlier than that, you need to refresh the Remotion Studio browser tab manually.

```

example.ts
tsx

import { getStaticFiles, StaticFile, Video } from "remotion";

const files = getStaticFiles();
/*
[
  {
    "name": "video.mp4",
    "src": "/static-7n5spa/video.mp4",
    "sizeInBytes": 432944,
    "lastModified": 1670170466865
  },
  {
    "name": "assets/data.json",
    "src": "/static-7n5spa/assets/data.json",
    "sizeInBytes": 1311,
    "lastModified": 1670170486089
  },
]
*/

// ❗ Don't pass the `name` directly to the `src` of a media element
const videoName = files[0].name;

// ✅ Wrap it in staticFile() instead or use `src`
const videoSrc = files[0].src;

// Find a file by it's name and import it
const data = files.find((f) => {
  return f.name === "video.mp4";
}) as StaticFile; // Use `as StaticFile` to assert the file exists

// Use the `src` property to get a src to pass to a media element
<Video src={data.src} />;
```

```

example.ts
tsx

import { getStaticFiles, StaticFile, Video } from "remotion";

const files = getStaticFiles();
/*
[
  {
    "name": "video.mp4",
    "src": "/static-7n5spa/video.mp4",
    "sizeInBytes": 432944,
    "lastModified": 1670170466865
  },
  {
    "name": "assets/data.json",
    "src": "/static-7n5spa/assets/data.json",
    "sizeInBytes": 1311,
    "lastModified": 1670170486089
  },
]
*/

// ❗ Don't pass the `name` directly to the `src` of a media element
const videoName = files[0].name;

// ✅ Wrap it in staticFile() instead or use `src`
const videoSrc = files[0].src;

// Find a file by it's name and import it
const data = files.find((f) => {
  return f.name === "video.mp4";
}) as StaticFile; // Use `as StaticFile` to assert the file exists

// Use the `src` property to get a src to pass to a media element
<Video src={data.src} />;
```

## API [​](\#api "Direct link to API")

Takes no arguments and returns an array of object, each of which have four entries:

- `name`: The path relative to the public folder.


note



Contains forward slashes `/` even on Windows.



note

- `src`: The path with a prefix. The prefix changes whenever the Studio server restarts.

- `sizeInBytes`: The file size. If it is a symbolic link, the file size of the original is reported.

- `lastModified`: Last modified date in Unix timestamp in milliseconds.

## Maximum files [​](\#maximum-files "Direct link to Maximum files")

For performance, only the first 10000 items are fetched and displayed.

**Changelog**

Before v4.0.64, only the first 1000 items were displayed.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/get-static-files.ts)
- [`staticFile()`](/docs/staticfile)
- [`watchStaticFile()`](/docs/watchstaticfile)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-static-files.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getRemotionEnvironment()](/docs/get-remotion-environment) [Next\
\
<IFrame>](/docs/iframe)

- [API](#api)
- [Maximum files](#maximum-files)
- [See also](#see-also)
