---
title: Chrome Headless Shellv4.0.208
source: Remotion Documentation
last_updated: 2024-12-22
---

# Chrome Headless Shellv4.0.208

- [Home page](/)
- Miscellaneous
- Chrome Headless Shell

On this page

# Chrome Headless Shell [v4.0.208](https://github.com/remotion-dev/remotion/releases/v4.0.208)

From v4.0.208, Remotion will download [Chrome Headless Shell](https://developer.chrome.com/blog/chrome-headless-shell) before the first render.

This will future-proof your Remotion project, because headless mode will be removed in a future version of Chrome.

## Supported platforms [​](\#supported-platforms "Direct link to Supported platforms")

The following platforms are supported:

- macOS (x64 and arm64)
- Windows (x64)
- Linux (x64 and arm64) - [install the Linux dependencies](/docs/miscellaneous/linux-dependencies)

## Ensure Headless Shell is available [​](\#ensure-headless-shell-is-available "Direct link to Ensure Headless Shell is available")

There are two ways to ensure Chrome Headless Shell is available:

- [`npx remotion browser ensure`](/docs/cli/browser/ensure) on the command line
- [`ensureBrowser()`](/docs/renderer/ensure-browser) as a Node.js / Bun API

It is recommended use call these functions if you do server-side rendering.

That way, if a request comes in wanting to render a video, the browser is already downloaded and ready to go.

## Bring your own binary [​](\#bring-your-own-binary "Direct link to Bring your own binary")

If you don't want Chrome Headless Shell to get installed or your platform is not supported, you need to specify your own Chromium-based browser:

- Using the [`setBrowserExecutable()`](/docs/config#setbrowserexecutable) option in the config file (for the CLI)
- Using the [`browserExecutable`](/docs/renderer/render-media) option in [`renderMedia()`](/docs/renderer/render-media) and other SSR APIs

In [Lambda](/docs/lambda) and [Cloud Run](/docs/cloudrun), a version of Chrome is already installed, so you don't need to do anything.

note

In a future version of Chrome, headless mode in the desktop browser will stop being supported and you will need to use the Chrome Headless Shell.

## Download location [​](\#download-location "Direct link to Download location")

Chrome Headless Shell will download to

```

node_modules/.remotion/chrome-headless-shell/[platform]/chrome-headless-shell-[platform]
```

```

node_modules/.remotion/chrome-headless-shell/[platform]/chrome-headless-shell-[platform]
```

`platform` can be one of `mac-arm64`, `mac-x64`, `linux64`, `linux-arm64` or `win64`.

At this path, a folder with the necessary files will be created.

An executable `./chrome-headless-shell` ( `.\chrome-headless-shell.exe` on Windows) will be created.

## Why this change? [​](\#why-this-change "Direct link to Why this change?")

Remotion previously used the desktop version of Chrome that many users had already installed.

This workflow is bound to break at some point, which is why Remotion is not picking up the desktop version of Chrome anymore.

## Best practices [​](\#best-practices "Direct link to Best practices")

To ensure your project does not get disrupted by an upcoming Chrome change, you should use the Remotion mechanisms which uses and pins the version of Chrome Headless Shell.

- Use Remotion v4.0.208 or later to not pick up an externally installed browser.
- Use [`npx remotion browser ensure`](/docs/cli/browser/ensure) to ensure Chrome Headless Shell is available.
- Do not download Chrome in your Dockerfile, but do install [Linux dependencies](/docs/miscellaneous/linux-dependencies) if you use Linux.
- Do not use `--browser-executable`, `browserExecutable` or `setBrowserExecutable()` options to override the Headless Shell with an incompatible Chrome version.

warning

Note: Most Linux distros do not allow you to pin a Chrome package.

If you use a Remotion version below v4.0.208, you are at risk of Chrome automatically being upgraded to a version that does not ship with a headless mode.

## What is Chrome Headless Shell? [​](\#what-is-chrome-headless-shell "Direct link to What is Chrome Headless Shell?")

Chrome used to ship with a `--headless` flag, which Remotion would use.

As of Chrome 123, the headless mode is split up into:

- `--headless=old`, which is ideal for screenshotting (and therefore Remotion)
- `--headless=new`, which is ideal for browser testing

`--headless=old` will stop working in a future version of Chrome.

The old headless mode is being extracted into "Chrome Headless Shell".

Hence we encourage you to use Chrome Headless Shell to future-proof your Remotion application.

## Version [​](\#version "Direct link to Version")

Remotion will download a well-tested Chrome version from the `123.0.6312` release line.

Upgrades may happen in a patch release and will be listed here.

## On Lambda and Cloud Run [​](\#on-lambda-and-cloud-run "Direct link to On Lambda and Cloud Run")

If you are using [Remotion Lambda](/docs/lambda) or [Cloud Run](/docs/cloudrun), you don't need to worry about installing a browser - it is included in the runtime already.

## Previous changes [​](\#previous-changes "Direct link to Previous changes")

### Thorium (v4.0.18 - v4.0.135) [​](\#thorium-v4018---v40135 "Direct link to Thorium (v4.0.18 - v4.0.135)")

In these versions, if no local browser can be found, an instance of [Thorium](https://thorium.rocks/) is downloaded.

Thorium is a free and open-source browser forked off Chromium, which includes the codecs needed to render videos.

### Chromium (before v4.0.18) [​](\#chromium-before-v4018 "Direct link to Chromium (before v4.0.18)")

In previous versions, Remotion would download the free version of Chromium, which would not include codecs for the proprietary H.264 and H.265 codecs.
This would often lead to problems when using the [`<Video>`](/docs/video) tag.

## See also [​](\#see-also "Direct link to See also")

- [`ensureBrowser()`](/docs/renderer/ensure-browser)
- [Media playback error](/docs/media-playback-error)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/chrome-headless-shell.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Temporary directory](/docs/miscellaneous/changing-temp-dir) [Next\
\
Linux Dependencies](/docs/miscellaneous/linux-dependencies)

- [Supported platforms](#supported-platforms)
- [Ensure Headless Shell is available](#ensure-headless-shell-is-available)
- [Bring your own binary](#bring-your-own-binary)
- [Download location](#download-location)
- [Why this change?](#why-this-change)
- [Best practices](#best-practices)
- [What is Chrome Headless Shell?](#what-is-chrome-headless-shell)
- [Version](#version)
- [On Lambda and Cloud Run](#on-lambda-and-cloud-run)
- [Previous changes](#previous-changes)
  - [Thorium (v4.0.18 - v4.0.135)](#thorium-v4018---v40135)
  - [Chromium (before v4.0.18)](#chromium-before-v4018)
- [See also](#see-also)
