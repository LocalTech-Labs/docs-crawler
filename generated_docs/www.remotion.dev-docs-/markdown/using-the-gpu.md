---
title: Using the GPU
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using the GPU

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- Using the GPU

On this page

# Using the GPU

Some types of content in Remotion can benefit from a GPU being available on the machine that is used for rendering.

By default in many cases, the GPU is disabled in headless mode, which can lead to a significant slowdown in rendering time.

## Content accelerated by the GPU [​](\#content-accelerated-by-the-gpu "Direct link to Content accelerated by the GPU")

- WebGL content (Three.JS, Skia, P5.js, Mapbox etc.)
- `box-shadow`
- `text-shadow`
- `background-image: linear-gradient()`
- `background-image: radial-gradient()`
- `filter: blur()`
- `filter: drop-shadow()`
- `transform`
- Many 2D Canvas operations

If a GPU is available, it should be enabled by default while in the Remotion Studio or Remotion Player.

However, in headless mode, Chromium disables the GPU, leading to a significant
slowdown in rendering time.

## Content not accelerated by the GPU [​](\#content-not-accelerated-by-the-gpu "Direct link to Content not accelerated by the GPU")

Contrary to popular belief, the following content is not accelerated by the GPU:

- `<Video>`
- `<OffthreadVideo>`
- [Canvas pixel manipulation](/docs/video-manipulation)

Furthermore, the encoding of the video is not accelerated by the GPU at this point.

## Use the `--gl` flag to enable the GPU during rendering [​](\#use-the---gl-flag-to-enable-the-gpu-during-rendering "Direct link to use-the---gl-flag-to-enable-the-gpu-during-rendering")

See [here](/docs/gl-options) for recommendations which OpenGL backend you should use during rendering.

## GPU for server-side rendering [​](\#gpu-for-server-side-rendering "Direct link to GPU for server-side rendering")

[See here](/docs/miscellaneous/cloud-gpu) for an example on how to use the GPU during server-side rendering.

## Using the GPU on Lambda [​](\#using-the-gpu-on-lambda "Direct link to Using the GPU on Lambda")

AWS Lambda instances have no GPU, so it is not possible to use it.

## What are your experiences? [​](\#what-are-your-experiences "Direct link to What are your experiences?")

We'd love to learn and document more findings about the GPU. Let us know and we will amend this document!

## See also [​](\#see-also "Direct link to See also")

- [OpenGL renderer backends](/docs/gl-options)
- [Hardware accelerated rendering](/docs/hardware-acceleration)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/gpu.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Multiple cores on Linux](/docs/miscellaneous/linux-single-process) [Next\
\
GPU in the cloud (bare)](/docs/miscellaneous/cloud-gpu)

- [Content accelerated by the GPU](#content-accelerated-by-the-gpu)
- [Content not accelerated by the GPU](#content-not-accelerated-by-the-gpu)
- [Use the `--gl` flag to enable the GPU during rendering](#use-the---gl-flag-to-enable-the-gpu-during-rendering)
- [GPU for server-side rendering](#gpu-for-server-side-rendering)
- [Using the GPU on Lambda](#using-the-gpu-on-lambda)
- [What are your experiences?](#what-are-your-experiences)
- [See also](#see-also)
