---
title: Multiple cores on Linux
source: Remotion Documentation
last_updated: 2024-12-22
---

# Multiple cores on Linux

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- Multiple cores on Linux

On this page

# Multiple cores on Linux

info

This document is outdated. As of 4.0.137, the default behavior is to use multi-process mode.

The setting still exists, but is `true` by default.

By default, Remotion starts the Chromium browser on Linux with the `--single-process` flag. This is because:

- The sandboxing feature of Chromium is not supported on all Linux distributions
- Older versions of Chromium used to crash when starting headlessly with multiple processes
- If some desktop libraries are missing, Chromium could crash on startup on some Distros.

This can lead to degraded rendering performance, especially when trying to leverage high-core CPUs and high Remotion concurrency.

## Enable multi-process rendering on Linux [​](\#enable-multi-process-rendering-on-linux "Direct link to Enable multi-process rendering on Linux")

_available from v4.0.42_

For now, this option is opt-in. To enable multi-process mode for Chromium during rendering on Linux:

- **renderMedia()** / **openBrowser()** / **renderFrames()**: Use the [`chromiumOptions.enableMultiProcessOnLinux`](/docs/renderer/render-media#enablemultiprocessonlinux) option.
- **CLI**: Use the [`--enable-multi-process-on-linux`](/docs/cli/render#--enable-multiprocess-on-linux) flag.
- **Config file**: Use the [`Config.setChromiumMultiProcessInLinux()`](/docs/config#setchromiummultiprocessonlinux) option.
- **Remotion Studio**: Check the toggle in the advanced settings.
- **Cloud Run**: From v4.0.42, this option is automatically enabled.
- **Lambda**: This option is disabled with no option to enable it because on Amazon Linux, Chrome will crash if the `--single-process` flag is not used. We recommend using more Lambdas instead of more concurrency per Lambda.

## Should I enable multi-process mode? [​](\#should-i-enable-multi-process-mode "Direct link to Should I enable multi-process mode?")

The answer depends on the Linux environment. If you are running Remotion on a server, you should try to enable multi-process mode and see if it works.

Our recommended [Docker](/docs/docker) image is confirmed working with multi-process mode enabled.

## See also [​](\#see-also "Direct link to See also")

- [Chromium flags](/docs/chromium-flags)
- [Dockerizing a Remotion app](/docs/docker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/linux-single-process.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Docker](/docs/docker) [Next\
\
Using the GPU](/docs/gpu)

- [Enable multi-process rendering on Linux](#enable-multi-process-rendering-on-linux)
- [Should I enable multi-process mode?](#should-i-enable-multi-process-mode)
- [See also](#see-also)
