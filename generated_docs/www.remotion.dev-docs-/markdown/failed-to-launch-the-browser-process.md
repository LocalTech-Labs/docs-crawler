---
title: Failed to launch the browser process
source: Remotion Documentation
last_updated: 2024-12-22
---

# Failed to launch the browser process

- [Home page](/)
- Troubleshooting
- Failed to launch the browser process

On this page

# Failed to launch the browser process

If you get an error message:

```

Failed to launch the browser process!

TROUBLESHOOTING: https://github.com/puppeteer/puppeteer/blob/main/docs/troubleshooting.md // or a URL pointing to this page
```

```

Failed to launch the browser process!

TROUBLESHOOTING: https://github.com/puppeteer/puppeteer/blob/main/docs/troubleshooting.md // or a URL pointing to this page
```

it means the browser was not able to start. The most common reasons are

## Missing shared libraries [​](\#missing-shared-libraries "Direct link to Missing shared libraries")

Install the [missing shared libraries for your OS](/docs/miscellaneous/linux-dependencies).

## Wrong OS or architecture [​](\#wrong-os-or-architecture "Direct link to Wrong OS or architecture")

If a x64 browser binary is run on arm64 or vice versa, you will get this error.

If the browser was compiled for a different OS or Linux distribution, it might also be incompatible.
Make sure you are using the correct binary for your OS and architecture.

## Debug using `--log=verbose` [​](\#debug-using---logverbose "Direct link to debug-using---logverbose")

Add the `--log=verbose` flag to the `npx remotion render` command.

This will log all browser process output, helping you to debug the issue.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/browser-launch.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Calling bundle() in bundled code](/docs/troubleshooting/bundling-bundle) [Next\
\
SIGKILL](/docs/troubleshooting/sigkill)

- [Missing shared libraries](#missing-shared-libraries)
- [Wrong OS or architecture](#wrong-os-or-architecture)
- [Debug using `--log=verbose`](#debug-using---logverbose)
