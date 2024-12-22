---
title: Fast Refresh not working
source: Remotion Documentation
last_updated: 2024-12-22
---

# Fast Refresh not working

- [Home page](/)
- Troubleshooting
- Fast Refresh not working

# Fast Refresh not working

If the Remotion Studio does not update after you have updated your file, it is usually due to one of two reasons:

[1](#1)

Studio server disconnected: This happens when you quit the process that was started when running `npx remotion studio`, usually using `Ctrl+C` or by quitting the terminal.

[2](#2)

Mismatched capitalization in the filenames: Ensure that the capitalization of filenames is correct in imports. For example, if your filename is `MyComp.tsx`, but you import it using `import {MyComp} from "./myComp.tsx"`, Fast Refresh may break.

note

This is a bug in Webpack for which [we've filed an issue](https://github.com/webpack/watchpack/issues/228).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/broken-fast-refresh.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Source could not be parsed](/docs/troubleshooting/could-not-be-parsed-as-a-value-list) [Next\
\
Timed out evaluating page function](/docs/troubleshooting/timed-out-page-function)
