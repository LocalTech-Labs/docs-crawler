---
title: npx remotion browser ensurev4.0.137
source: Remotion Documentation
last_updated: 2024-12-22
---

# npx remotion browser ensurev4.0.137

- [Home page](/)
- [Command line](/docs/cli/)
- [browser](/docs/cli/browser/)
- ensure

On this page

# npx remotion browser ensure [v4.0.137](https://github.com/remotion-dev/remotion/releases/v4.0.137)

Ensures that Remotion has a browser it can use for rendering.

```

npx remotion browser ensure
```

```

npx remotion browser ensure
```

## Arguments [​](\#arguments "Direct link to Arguments")

### `--browser-executable` [​](\#--browser-executable "Direct link to --browser-executable")

[Path to a custom Chrome executable](/docs/config#setbrowserexecutable). If not specified and Remotion cannot find one, one will be downloaded by this command.

### `--log` [​](\#--log "Direct link to --log")

One of `verbose`, `info`, `warn`, `error`.

Determines how much is being logged to the console.

`verbose` will also log `console.log`'s from the browser.

Default `info`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this command](https://github.com/remotion-dev/remotion/blob/main/packages/cli/src/browser/ensure.ts)
- [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell)
- [`ensureBrowser()`](/docs/renderer/ensure-browser)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cli/browser/ensure.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
browser](/docs/cli/browser/) [Next\
\
versions](/docs/cli/versions)

- [Arguments](#arguments)
  - [`--browser-executable`](#--browser-executable)
  - [`--log`](#--log)
- [See also](#see-also)
