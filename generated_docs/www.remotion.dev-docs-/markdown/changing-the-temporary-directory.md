---
title: Changing the temporary directory
source: Remotion Documentation
last_updated: 2024-12-22
---

# Changing the temporary directory

- [Home page](/)
- Miscellaneous
- Temporary directory

# Changing the temporary directory

During rendering, Remotion will write items into the temporary directory of the system. This includes rendered frames, uncompressed audio and other artifacts.

Changing the temporary directory may be useful if the default temporary directory is on a disk that has limited space available.

If you want to customize the directory that Remotion uses, you can set the `TMPDIR` (Linux, macOS) or `TEMP` (Windows) environment variable to specify a directory to your liking.

```

bash

TMPDIR=/var/tmp npx remotion render
```

```

bash

TMPDIR=/var/tmp npx remotion render
```

Remotion will make a new temporary directory in the path that you have specified.
This is because Remotion uses the the Node.JS [`os.tmpdir()`](https://github.com/nodejs/node/blob/58431c0e6bb1829b6ccafc5cf6340226c15da790/lib/os.js#L181) API, which checks for environment variables in the following order:

- `TMPDIR`, `TMP`, `TEMP` on non-Windows platforms
- `TEMP` and `TMP` on Windows platforms

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/changing-temp-dir.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Chromium flags](/docs/chromium-flags) [Next\
\
Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell)
