---
title: watchStaticFile()v4.0.144
source: Remotion Documentation
last_updated: 2024-12-22
---

# watchStaticFile()v4.0.144

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- watchStaticFile()

On this page

# watchStaticFile() [v4.0.144](https://github.com/remotion-dev/remotion/releases/v4.0.144)

Watches for changes in a specific [static file](/docs/staticfile) and invokes a callback function when the file changes, enabling dynamic updates in your Remotion projects.

note

This API is being moved from the `remotion` package.

Prefer this API over the old one.

note

This feature is only available within the Remotion Studio environment. In the Player, events will never fire.

## Example [​](\#example "Direct link to Example")

```

example.tsx
tsx

import { StaticFile, watchStaticFile } from "@remotion/studio";

// Watch for changes in a specific static file
const { cancel } = watchStaticFile(
  "your-static-file.jpg",
  (newData: StaticFile | null) => {
    if (newData) {
      console.log(`File ${newData.name} has been added or modified.`);
    } else {
      console.log("File has been deleted.");
    }
  },
);

// To stop watching for changes, call the cancel function
cancel();
```

```

example.tsx
tsx

import { StaticFile, watchStaticFile } from "@remotion/studio";

// Watch for changes in a specific static file
const { cancel } = watchStaticFile(
  "your-static-file.jpg",
  (newData: StaticFile | null) => {
    if (newData) {
      console.log(`File ${newData.name} has been added or modified.`);
    } else {
      console.log("File has been deleted.");
    }
  },
);

// To stop watching for changes, call the cancel function
cancel();
```

## Arguments [​](\#arguments "Direct link to Arguments")

Takes two arguments and returns a function that can be used to `cancel` the event listener.

### `filename` [​](\#filename "Direct link to filename")

A name of the file in `/public` folder to watch for changes.

### `callback` [​](\#callback "Direct link to callback")

A callback function that will be called when the file is modified. As an argument, a [`StaticFile`](/docs/getstaticfiles#api) or `null` is passed.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/watch-static-file.ts)
- [`watchPublicFolder()`](/docs/studio/watch-public-folder)
- [`staticFile()`](/docs/staticfile)
- [`getStaticFiles()`](/docs/studio/get-static-files)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/watch-static-file.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
watchPublicFolder()](/docs/studio/watch-public-folder) [Next\
\
writeStaticFile()](/docs/studio/write-static-file)

- [Example](#example)
- [Arguments](#arguments)
  - [`filename`](#filename)
  - [`callback`](#callback)
- [See also](#see-also)
