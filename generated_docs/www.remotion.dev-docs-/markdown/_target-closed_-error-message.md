---
title: 'Target closed' error message
source: Remotion Documentation
last_updated: 2024-12-22
---

# 'Target closed' error message

- [Home page](/)
- Troubleshooting
- Target closed

On this page

# 'Target closed' error message

This error message **comes up when a Chromium Tab crashes**. Read on for possible reasons for this message.

## Missing dependencies (Linux users) [​](\#missing-dependencies-linux-users "Direct link to Missing dependencies (Linux users)")

Make sure that you [installed the necessary dependencies](/docs/miscellaneous/linux-dependencies) in Linux for running Chrome.

## Out of memory or CPU overload [​](\#out-of-memory-or-cpu-overload "Direct link to Out of memory or CPU overload")

A Chrome tab can crash if the process runs out of memory, or if the website stops responding because it uses the CPU too heavily.

- [Decrease the concurrency so each tab has more resources](/docs/cli/render#--concurrency)
- In Remotion Lambda, increase the memory size.

## Corrupt Google Chrome binary [​](\#corrupt-google-chrome-binary "Direct link to Corrupt Google Chrome binary")

Locate the binary that you are using to render the video. Run a test whether you can open Google Chrome / Chromium normally on this machine. Potentially reinstall the browser to fix the problem.

note

Add `--log=verbose` during a render to print the browser executable that is being used (from Remotion v2.6.7 on).

Did none of these tips help? File an issue in the Remotion repository.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/target-closed.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Timeouts](/docs/timeout) [Next\
\
Media playback error](/docs/media-playback-error)

- [Missing dependencies (Linux users)](#missing-dependencies-linux-users)
- [Out of memory or CPU overload](#out-of-memory-or-cpu-overload)
- [Corrupt Google Chrome binary](#corrupt-google-chrome-binary)
