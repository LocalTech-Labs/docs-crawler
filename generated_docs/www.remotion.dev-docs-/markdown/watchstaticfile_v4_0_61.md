---
title: watchStaticFile()v4.0.61
source: Remotion Documentation
last_updated: 2024-12-22
---

# watchStaticFile()v4.0.61

- [Home page](/)
- [remotion](/docs/remotion)
- watchStaticFile()

On this page

# watchStaticFile() [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61)

note

This API is being moved to the `@remotion/studio` package. Prefer importing the API from [`@remotion/studio`](/docs/studio/watch-static-file) instead of `remotion`.

Watches for changes in a specific [static file](/docs/staticfile) and invokes a callback function when the file changes, enabling dynamic updates in your Remotion projects.

warning

This feature is only available within the Remotion Studio environment. In the Player, events will never fire.

## Example [​](\#example "Direct link to Example")

```

example.tsx
tsx

import { StaticFile, watchStaticFile } from "remotion";

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

import { StaticFile, watchStaticFile } from "remotion";

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

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/watch-static-file.ts)
- [`staticFile()`](/docs/staticfile)
- [`getStaticFiles()`](/docs/getstaticfiles)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/watch-static-file.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Video>](/docs/video) [Next\
\
@remotion/bundler](/docs/bundler)

- [Example](#example)
- [Arguments](#arguments)
  - [`filename`](#filename)
  - [`callback`](#callback)
- [See also](#see-also)
