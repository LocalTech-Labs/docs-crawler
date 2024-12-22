---
title: Debugging stuck renders
source: Remotion Documentation
last_updated: 2024-12-22
---

# Debugging stuck renders

- [Home page](/)
- Troubleshooting
- Stuck renders

On this page

# Debugging stuck renders

If your render is stuck, proceed as follows:

## Check for long `delayRender()` calls [​](\#check-for-long-delayrender-calls "Direct link to check-for-long-delayrender-calls")

If the `delayRender()` timeout is too big and does not resolve, a render may get stuck indefinitely without an error message.

This issue may be more confusing on Remotion Lambda, whose runtime also has a timeout.

If the `delayRender()` timeout is bigger than the Lambda function timeout, the Lambda function will first time out without the `delayRender()` call failing first, leading to no error message.

To debug the issue:

- Set the `delayRender()` timeout value to a smaller value, so that if it does not resolve, an error message is shown.
- Always sse [`cancelRender()`](/docs/cancel-render) to handle any error that prevents you from calling `continueRender()`.

## Enable verbose logging [​](\#enable-verbose-logging "Direct link to Enable verbose logging")

[Enable verbose logging](/docs/troubleshooting/debug-failed-render) to see if there are any failures from the browser or other child processes.

## Ensure Chrome Headless Shell is used [​](\#ensure-chrome-headless-shell-is-used "Direct link to Ensure Chrome Headless Shell is used")

Newer Chrome versions may not work with Remotion anymore and cause a render to get stuck, because they have removed the Headless mode.

Instead, you need to migrate to the [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell).

note

This issue does not occur on Remotion Lambda.

Follow these best practices:

- Do not download Chrome from a Linux package manager.
- Do not set the `--chrome-executable` option and let Remotion download a compatible version of Chrome Headless Shell for you.
- Use [`npx remotion browser ensure`](/docs/cli/browser/ensure) to make sure a compatible versions of Chrome Headless Shell is ready before the first render.

## Lambda: Are you reading the `errors` field? [​](\#lambda-are-you-reading-the-errors-field "Direct link to lambda-are-you-reading-the-errors-field")

If the [`overallProgress`](/docs/lambda/getrenderprogress#overallprogress) field is stuck on Lambda, ensure that you are:

- Reading the [`errors`](/docs/lambda/getrenderprogress#errors) field to see if there are any errors.
- not further polling [`getRenderProgress()`](/docs/lambda/getrenderprogress) if the [`fatalErrorEncountered`](/docs/lambda/getrenderprogress#fatalerrorencountered) field is `true`.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/render-stuck.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Could not determine executable to run](/docs/troubleshooting/could-not-find-executable-to-run) [Next\
\
Different segments at different speeds](/docs/miscellaneous/snippets/different-segments-at-different-speeds)

- [Check for long `delayRender()` calls](#check-for-long-delayrender-calls)
- [Enable verbose logging](#enable-verbose-logging)
- [Ensure Chrome Headless Shell is used](#ensure-chrome-headless-shell-is-used)
- [Lambda: Are you reading the `errors` field?](#lambda-are-you-reading-the-errors-field)
