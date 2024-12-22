---
title: ensureBrowser()v4.0.137
source: Remotion Documentation
last_updated: 2024-12-22
---

# ensureBrowser()v4.0.137

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- ensureBrowser()

On this page

# ensureBrowser() [v4.0.137](https://github.com/remotion-dev/remotion/releases/v4.0.137)

Ensures a browser is locally installed so a Remotion render can be executed.

```

Simple usage
tsx

import { ensureBrowser } from "@remotion/renderer";

await ensureBrowser();
```

```

Simple usage
tsx

import { ensureBrowser } from "@remotion/renderer";

await ensureBrowser();
```

```

Setting a specific Chrome version and listening to progress
tsx

import { ensureBrowser } from "@remotion/renderer";

await ensureBrowser({
  onBrowserDownload: () => {
    console.log("Downloading browser");

    return {
      version: "123.0.6312.86",
      onProgress: ({ percent }) => {
        console.log(`${Math.round(percent * 100)}% downloaded`);
      },
    };
  },
});
```

```

Setting a specific Chrome version and listening to progress
tsx

import { ensureBrowser } from "@remotion/renderer";

await ensureBrowser({
  onBrowserDownload: () => {
    console.log("Downloading browser");

    return {
      version: "123.0.6312.86",
      onProgress: ({ percent }) => {
        console.log(`${Math.round(percent * 100)}% downloaded`);
      },
    };
  },
});
```

## API [​](\#api "Direct link to API")

An object with the following properties, all of which are optional:

### `browserExecutable` [​](\#browserexecutable "Direct link to browserexecutable")

Pass a path to a browser executable that you want to use instead of downloading.

If the path does not exist, this function will throw.

Pass the same path to any other API that supports the `browserExecutable` option.

### `logLevel` [​](\#loglevel "Direct link to loglevel")

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

### `onBrowserDownload` [​](\#onbrowserdownload "Direct link to onbrowserdownload")

Specify a specific version of Chrome that should be used and hook into the download progress.

See the example below for the function signature.

```

init.ts
tsx

import {
  ensureBrowser,
  OnBrowserDownload,
  DownloadBrowserProgressFn,
} from "@remotion/renderer";

const onProgress: DownloadBrowserProgressFn = ({
  percent,
  downloadedBytes,
  totalSizeInBytes,
}) => {
  console.log(`${Math.round(percent * 100)}% downloaded`);
};

const onBrowserDownload: OnBrowserDownload = () => {
  console.log("Downloading browser");

  return {
    // Pass `null` to use Remotion's recommendation.
    version: "123.0.6312.86",
    onProgress,
  };
};

await ensureBrowser({
  onBrowserDownload,
});
```

```

init.ts
tsx

import {
  ensureBrowser,
  OnBrowserDownload,
  DownloadBrowserProgressFn,
} from "@remotion/renderer";

const onProgress: DownloadBrowserProgressFn = ({
  percent,
  downloadedBytes,
  totalSizeInBytes,
}) => {
  console.log(`${Math.round(percent * 100)}% downloaded`);
};

const onBrowserDownload: OnBrowserDownload = () => {
  console.log("Downloading browser");

  return {
    // Pass `null` to use Remotion's recommendation.
    version: "123.0.6312.86",
    onProgress,
  };
};

await ensureBrowser({
  onBrowserDownload,
});
```

## Return value [​](\#return-value "Direct link to Return value")

A promise with no value.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/renderer/src/ensure-browser.ts)
- [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell)
- [Server-Side rendering](/docs/ssr)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/ensure-browser.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
openBrowser()](/docs/renderer/open-browser) [Next\
\
makeCancelSignal()](/docs/renderer/make-cancel-signal)

- [API](#api)
  - [`browserExecutable`](#browserexecutable)
  - [`logLevel`](#loglevel)
  - [`onBrowserDownload`](#onbrowserdownload)
- [Return value](#return-value)
- [See also](#see-also)
