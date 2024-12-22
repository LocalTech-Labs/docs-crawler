---
title: Runtime
source: Remotion Documentation
last_updated: 2024-12-22
---

# Runtime

- [Home page](/)
- [Lambda](/docs/lambda)
- Runtime

On this page

# Runtime

This page describes the environment that the Lambda function is running in.

## Node.JS Version [​](\#nodejs-version "Direct link to Node.JS Version")

The Lambda function uses a NodeJS version from the `18.x` release line.

The Lambda runtime will get locked to

```

arn:aws:lambda:${region}::runtime:b97ad873eb5228db2e7d5727cd116734cc24c92ff1381739c4400c095404a2d3
```

```

arn:aws:lambda:${region}::runtime:b97ad873eb5228db2e7d5727cd116734cc24c92ff1381739c4400c095404a2d3
```

if your user policy includes `lambda:PutRuntimeManagementConfig`, which is recommended.

Otherwise, future updates to the runtime by AWS have the potential to break the function. If you don't have this permission in your policy, a warning will be printed.

Details

Changelog
Before Remotion v4.0.0, the Node.JS version was `14.x`.

## Memory size [​](\#memory-size "Direct link to Memory size")

The default is 2048 MB. You can configure it by passing an argument to [`deployFunction()`](/docs/lambda/deployfunction) or by passing a `--memory` flag to the CLI when deploying a function.

## Timeout [​](\#timeout "Direct link to Timeout")

The default is 120 seconds. You can configure it when calling [`deployFunction()`](/docs/lambda/deployfunction) or by passing a `--timeout` flag to the CLI when deploying a function.

Note that you probably don't need to increase it - Since the video is rendered by splitting it into many parts and those parts are rendered in parallel, there are rare cases where you need more than 120 seconds.

## Storage space [​](\#storage-space "Direct link to Storage space")

The function has between [512MB and 10GB of storage space](/docs/lambda/disk-size) in total available for video rendering depending on your configuration. Keep in mind that the concatenations of various chunks into one video takes place within a Lambda function, so the space must suffice for both the chunks and the output video.

## Core count / vCPUs [​](\#core-count--vcpus "Direct link to Core count / vCPUs")

The amount of cores inside a Lambda is dependent on the amount of memory you give it. According to [this research](https://web.archive.org/web/20230331040434/https://www.sentiatechblog.com/aws-re-invent-2020-day-3-optimizing-lambda-cost-with-multi-threading), these are the tiers:

MemoryvCPUs128 - 3008 MB23009 - 5307 MB35308 - 7076 MB47077 - 8845 MB58846+ MB6

You can render multiple frames at once inside a Lambda function by using the [`concurrencyPerLambda`](/docs/lambda/rendermediaonlambda#concurrencyperlambda) option.

## Chrome [​](\#chrome "Direct link to Chrome")

The function already includes a running version of Chrome.
The browser was compiled including the proprietary codecs, so you can include MP4 videos into your project.

Remotion versionChrome versionFrom 5.0.0 (upcoming)123.0.6312.86From 4.0.0114.0.5731.1From 3.2.0104.0.5112.64From 3.0.8101.0.4951.68From 3.0.098.0.4758.139

## FFMPEG [​](\#ffmpeg "Direct link to FFMPEG")

Since Remotion 4.0, the function already includes an FFMPEG 6.0 binary that was compiled to support all [codecs](/docs/encoding) that Remotion supports.

Before v4.0.0

Version built from source: `N-104627-g40cf317d09` (corresponds to v4.4)

Configuration:

```
--prefix=/home/ec2-user/ffmpeg_build --pkg-config-flags=--static
--extra-cflags=-I/home/ec2-user/ffmpeg_build/include
--extra-ldflags=-L/home/ec2-user/ffmpeg_build/lib --extra-libs=-lpthread
--extra-libs=-lm --bindir=/home/ec2-user/bin --enable-gpl --enable-libfdk-aac
--enable-libfreetype --enable-libmp3lame --enable-libopus --enable-libvpx
--enable-libx264 --enable-libx265 --enable-nonfree
```

## Fonts [​](\#fonts "Direct link to Fonts")

The function includes the following fonts:

- Noto Color Emoji
- Noto Sans Black
- Noto Sans Bold
- Noto Sans Regular
- Noto Sans SemiBold
- Noto Sans Thin
- Noto Sans Arabic Regular
- Noto Sans Devanagari Regular
- Noto Sans Hebrew Regular
- Noto Sans Tamil Regular
- Noto Sans Thai Regular

Since December 2021 the following fonts are also available **only on the `arm64` version of Remotion Lambda:**

- Noto Sans Simplified Chinese Regular
- Noto Sans Simplified Chinese Bold
- Noto Sans Traditional Chinese Regular
- Noto Sans Traditional Chinese Bold
- Noto Sans Korean Regular
- Noto Sans Korean Bold
- Noto Sans Japanese Regular
- Noto Sans Japanese Bold

If you'd like to use different fonts, we recommend using Webfonts.

While the set of default fonts that we can include must be kept small in order to save space, we are happy to hear feedback if you encounter a scenario where characters cannot be rendered.

## Customize layers [​](\#customize-layers "Direct link to Customize layers")

See: [Customize Lambda layers](/docs/lambda/custom-layers) to learn about how you can customize this stack.

## Runtime changes in Remotion 5.0 [​](\#runtime-changes-in-remotion-50 "Direct link to Runtime changes in Remotion 5.0")

Remotion 5.0 will use the following runtime:

- Node.js 20.x instead of 18.x
- Chrome 123.0.6312.86 instead of 114.0.5731.1.
- AWS Lambda runtime locked to `arn:aws:lambda:${region}::runtime:b97ad873eb5228db2e7d5727cd116734cc24c92ff1381739c4400c095404a2d3`

Available from v4.0.148: You may opt to use the new runtime by setting `enableV5Runtime: true` in the [`deployFunction()`](/docs/lambda/deployfunction) function.

If you use the CLI, pass `--enable-v5-runtime` to the [`npx remotion lambda functions deploy`](/docs/lambda/cli/functions#deploy) command.

## See also [​](\#see-also "Direct link to See also")

- [Customize Lambda layers](/docs/lambda/custom-layers)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/runtime.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Concurrency](/docs/lambda/concurrency) [Next\
\
Disk size](/docs/lambda/disk-size)

- [Node.JS Version](#nodejs-version)
- [Memory size](#memory-size)
- [Timeout](#timeout)
- [Storage space](#storage-space)
- [Core count / vCPUs](#core-count--vcpus)
- [Chrome](#chrome)
- [FFMPEG](#ffmpeg)
- [Fonts](#fonts)
- [Customize layers](#customize-layers)
- [Runtime changes in Remotion 5.0](#runtime-changes-in-remotion-50)
- [See also](#see-also)
