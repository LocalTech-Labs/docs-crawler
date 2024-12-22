---
title: openBrowser()
source: Remotion Documentation
last_updated: 2024-12-22
---

# openBrowser()

- [Home page](/)
- [@remotion/renderer](/docs/renderer)
- openBrowser()

On this page

# openBrowser()

_Available since v3.0 - Part of the `@remotion/renderer` package._

Opens a Chrome or Chromium browser instance. By reusing an instance across [`renderFrames()`](/docs/renderer/render-frames), [`renderStill()`](/docs/renderer/render-still), [`renderMedia()`](/docs/renderer/render-media) and [`getCompositions()`](/docs/renderer/get-compositions) calls, you can save time by not opening and closing browsers for each call.

```

ts

const openBrowser: (
  browser: Browser,
  options: {
    shouldDumpIo?: boolean;
    browserExecutable?: string | null;
    chromiumOptions?: ChromiumOptions;
  },
) => Promise<puppeteer.Browser>;
```

```

ts

const openBrowser: (
  browser: Browser,
  options: {
    shouldDumpIo?: boolean;
    browserExecutable?: string | null;
    chromiumOptions?: ChromiumOptions;
  },
) => Promise<puppeteer.Browser>;
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `browser` [​](\#browser "Direct link to browser")

Currently the only valid option is `"chrome"`. This field is reserved for future compatibility with other browsers.

### `options?` [​](\#options "Direct link to options")

_optional_

An object containing one or more of the following options:

#### `shouldDumpIo?` [​](\#shoulddumpio "Direct link to shoulddumpio")

_optional, deprecated since v4.0.189, scheduled for removal in v5.0_

If set to `true`, logs and other browser diagnostics are being printed to standard output. This setting is useful for debugging.

**Will be removed in 5.0:** Use `logLevel` instead.

#### `logLevel?` [v4.0.189](https://github.com/remotion-dev/remotion/releases/v4.0.189) [​](\#loglevel "Direct link to loglevel")

_optional_

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

#### `browserExecutable?` [​](\#browserexecutable "Direct link to browserexecutable")

_optional_

A string defining the absolute path on disk of the browser executable that should be used. By default Remotion will try to detect it automatically and download one if none is available. If `puppeteerInstance` is defined, it will take precedence over `browserExecutable`.

#### `chromiumOptions?` [​](\#chromiumoptions "Direct link to chromiumoptions")

_optional_

Allows you to set certain Chromium / Google Chrome flags. See: [Chromium flags](/docs/chromium-flags).

note

Chromium flags need to be set at browser launch. If you pass an instance to SSR APIs like [`renderMedia()`](/docs/renderer/render-media), the `chromiumOptions` option of that API will not apply, but rather the flags that have been passed to `openBrowser()`.

#### `forceDeviceScaleFactor?` [​](\#forcedevicescalefactor "Direct link to forcedevicescalefactor")

Set a [scale](/docs/scaling). If you plan to use scaling, you already need to set it when opening the browser.

#### `onBrowserDownload?` [v4.0.137](https://github.com/remotion-dev/remotion/releases/v4.0.137) [​](\#onbrowserdownload "Direct link to onbrowserdownload")

Gets called when no compatible local browser is detected on the system and this API needs to download a browser. Return a callback to observe progress. [See here for how to use this option.](/docs/renderer/ensure-browser#onbrowserdownload)

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/renderer/src/open-browser.ts)
- [Server-Side rendering](/docs/ssr)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/renderer/open-browser.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
stitchFramesToVideo()](/docs/renderer/stitch-frames-to-video) [Next\
\
ensureBrowser()](/docs/renderer/ensure-browser)

- [Arguments](#arguments)
  - [`browser`](#browser)
  - [`options?`](#options)
- [See also](#see-also)
