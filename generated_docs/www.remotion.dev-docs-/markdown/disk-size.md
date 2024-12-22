---
title: Disk size
source: Remotion Documentation
last_updated: 2024-12-22
---

# Disk size

- [Home page](/)
- [Lambda](/docs/lambda)
- Disk size

On this page

# Disk size

By default, each Lambda function comes with an ephemereal disk size of:

Remotion VersionDefault<5.0.02048MB>=5.0.010240MB

Increasing disk space will allow for longer videos, and speeds up renders because Chrome has access to more disk cache.

Disk.

Adding more disk space will almost not affect the cost of your renders at all (less than 1% more cost).

Therefore we recommend setting the disk size to the maximum possible value: 10240MB.

## Approximate maximum video length [​](\#approximate-maximum-video-length "Direct link to Approximate maximum video length")

Note that we recommend setting the disk size to the maximum possible value: 10240MB.

Disk sizeApproximate Maximum video length512 MB8 min - 1080p1024 MB16 min - 1080p2048 MB32 min - 1080p4096 MB1h 4min - 1080p8192 MB2h 8min - 1080p10240 MB2h 40min - 1080p

> These are approximate values and will not exactly match your scenario. Video output size is dependant on the video content and audio. Measure and find the values that work best for you.

## Setting the disk size [​](\#setting-the-disk-size "Direct link to Setting the disk size")

- Use the [`diskSizeInMb` option of `deployFunction()`](/docs/lambda/deployfunction#disksizeinmb) to set the disk size when you deploy.
- Use the [`--disk`](/docs/lambda/cli/functions) flag if you use the `remotion lambda functions deploy` command.

## Pricing [​](\#pricing "Direct link to Pricing")

Using more disk space costs marginally more.

Setting the disk size to the maximum possible value: 10240MB will cost less than 1% more.

See the [Lambda pricing page](https://aws.amazon.com/lambda/pricing/) "Lambda Ephemereal Storage Pricing" section for pricing.

The [`estimatePrice()`](/docs/lambda/estimateprice) API does also factor disk size into account.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/disk-size.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Runtime](/docs/lambda/runtime) [Next\
\
FAQ](/docs/lambda/faq)

- [Approximate maximum video length](#approximate-maximum-video-length)
- [Setting the disk size](#setting-the-disk-size)
- [Pricing](#pricing)
