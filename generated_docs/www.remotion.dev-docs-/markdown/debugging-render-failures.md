---
title: Debugging render failures
source: Remotion Documentation
last_updated: 2024-12-22
---

# Debugging render failures

- [Home page](/)
- Troubleshooting
- Debugging render failures

On this page

# Debugging render failures

Since JavaScript code is executing, it may happen that a render may due to an exception. Here are general tips to troubleshoot the issue.

[1](#1) **Enable verbose logging**

By enabling more detailed logging, all `console.log` statements from your code will be made visible alongside other debugging information.

From the CLI: Add the [`--log=verbose`](/docs/cli/render#--log) flag to your render command.

From Node.JS: Add the [`verbose: true`](/docs/renderer/render-media#verbose) options to `renderMedia()`.

note

If you see a log multiple times, it is because the render is split up to multiple threads. Set `--concurrency=1` temporarily to only see each log once.

[2](#2) **Adding logs to your project**

Use `console.log` in your code to understand the order things are executing in and verify your assumptions about how your code should behave.

[3](#3) **Remove components one by one**

Remove components until the video is empty. At which point does the error disappear? This can help you identify the component responsible for the render failure.

[4](#4) **Search issues and documentation**

It's also helpful to check for [issues on GitHub](https://github.com/remotion-dev/remotion/issues) to see if other people have encountered similar problems. If you find an issue that matches your problem, you can add a comment to the issue to help the community troubleshoot the problem.

Also [search the documentation](/search) which has over 300 pages and contains troubleshooting instructions for many common problems.

[5](#5) **Ask for help**

You can ask for help on GitHub and Discord. [Read on to see how to get help](/docs/get-help)!

## See also [​](\#see-also "Direct link to See also")

- [Debug failed Lambda renders](/docs/lambda/troubleshooting/debug)
- [Ask for help](/docs/get-help)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/debug-failed-render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Import from After Effects](/docs/after-effects) [Next\
\
Timeouts](/docs/timeout)

- [See also](#see-also)
