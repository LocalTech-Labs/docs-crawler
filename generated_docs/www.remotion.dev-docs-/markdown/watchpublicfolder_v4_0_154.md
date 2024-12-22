---
title: watchPublicFolder()v4.0.154
source: Remotion Documentation
last_updated: 2024-12-22
---

# watchPublicFolder()v4.0.154

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- watchPublicFolder()

On this page

# watchPublicFolder() [v4.0.154](https://github.com/remotion-dev/remotion/releases/v4.0.154)

Watches for changes in the [public directory](/docs/terminology/public-dir) and calls a callback function when a file is added, removed or modified.

note

This feature is only available within the Remotion Studio environment. In the Player, events will never fire.

## Example [​](\#example "Direct link to Example")

```

example.tsx
tsx

import { StaticFile, watchPublicFolder } from "@remotion/studio";

// Watch for changes in a specific static file
const { cancel } = watchPublicFolder((newFiles: StaticFile[]) => {
  console.log("The public folder now contains:", newFiles);
});

// To stop watching for changes, call the cancel function
cancel();
```

```

example.tsx
tsx

import { StaticFile, watchPublicFolder } from "@remotion/studio";

// Watch for changes in a specific static file
const { cancel } = watchPublicFolder((newFiles: StaticFile[]) => {
  console.log("The public folder now contains:", newFiles);
});

// To stop watching for changes, call the cancel function
cancel();
```

## Arguments [​](\#arguments "Direct link to Arguments")

Takes one argument and returns a function that can be used to `cancel` the event listener.

### `callback` [​](\#callback "Direct link to callback")

A callback function that will be called when the directory is modified. As an argument, an array of [`StaticFile`](/docs/getstaticfiles#api)'s is passed.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/watch-public-folder.ts)
- [`watchStaticFile()`](/docs/studio/watch-static-file)
- [`staticFile()`](/docs/staticfile)
- [`getStaticFiles()`](/docs/studio/get-static-files)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/watch-public-folder.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getStaticFiles()](/docs/studio/get-static-files) [Next\
\
watchStaticFile()](/docs/studio/watch-static-file)

- [Example](#example)
- [Arguments](#arguments)
  - [`callback`](#callback)
- [See also](#see-also)
