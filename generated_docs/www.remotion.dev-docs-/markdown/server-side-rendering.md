---
title: Server-Side Rendering
source: Remotion Documentation
last_updated: 2024-12-22
---

# Server-Side Rendering

- [Home page](/)
- Server-side rendering

On this page

# Server-Side Rendering

Remotion's rendering engine is built with Node.JS, which makes it easy to render a video in the cloud.

## Render a video on AWS Lambda [​](\#render-a-video-on-aws-lambda "Direct link to Render a video on AWS Lambda")

The easiest and fastest way to render videos in the cloud is to use [`@remotion/lambda`](/docs/lambda).

## Render a video using Node.js APIs [​](\#render-a-video-using-nodejs-apis "Direct link to Render a video using Node.js APIs")

We provide a set of APIs to render videos using Node.js and Bun.

See an [example](/docs/ssr-node) or the [API reference](/docs/renderer) for more information.

## Render using GitHub Actions [​](\#render-using-github-actions "Direct link to Render using GitHub Actions")

The [Hello World starter template](https://github.com/remotion-dev/template-helloworld) includes a GitHub Actions workflow file [`.github/workflows/render-video.yml`](https://github.com/remotion-dev/template-helloworld/blob/main/.github/workflows/render-video.yml).

[1](#1)

Commit the template to a GitHub repository.

[2](#2) On GitHub, click the `Actions` tab.

[3](#3) Select the `Render video` workflow on the left.

[4](#4) A `Run workflow` button should appear. Click it.

[5](#5) Fill in the props of the root component and click `Run workflow`.

[6](#6) After the rendering is finished, you can download the video under `Artifacts`.

Note that running the workflow may incur costs. However, the workflow will only run if you actively trigger it.

See also: [Passing input props in GitHub Actions](/docs/passing-props#passing-input-props-in-github-actions)

## Render a video using Docker [​](\#render-a-video-using-docker "Direct link to Render a video using Docker")

See: [Dockerizing a Remotion project](/docs/docker)

## Render a video using GCP Cloud Run (Alpha) [​](\#render-a-video-using-gcp-cloud-run-alpha "Direct link to Render a video using GCP Cloud Run (Alpha)")

Check out the experimental [Cloud Run](/docs/cloudrun-alpha) package.

Our plan is to port the Lambda runtime to Cloud Run instead of maintaining a separate implementation.

## API reference [​](\#api-reference "Direct link to API reference")

[**getCompositions()** \
\
List available compositions](/docs/renderer/get-compositions) [**selectComposition()** \
\
Get a composition](/docs/renderer/select-composition) [**renderMedia()** \
\
Render a video or audio](/docs/renderer/render-media) [**renderFrames()** \
\
Render a series of images](/docs/renderer/render-frames) [**renderStill()** \
\
Render a single image](/docs/renderer/render-still) [**stitchFramesToVideo()** \
\
Turn images into a video](/docs/renderer/stitch-frames-to-video) [**openBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/open-browser) [**ensureBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/ensure-browser) [**makeCancelSignal()** \
\
Create token to later cancel a render](/docs/renderer/make-cancel-signal) [**getVideoMetadata()** \
\
Get metadata from a video file in Node.js](/docs/renderer/get-video-metadata) [**getSilentParts()** \
\
Obtain silent portions of a video or audio](/docs/renderer/get-silent-parts) [**ensureFfmpeg()** \
\
Check for ffmpeg binary and install if not existing](/docs/renderer/ensure-ffmpeg) [**ensureFfprobe()** \
\
Check for ffprobe binary and install if not existing](/docs/renderer/ensure-ffprobe) [**getCanExtractFramesFast()** \
\
Probes for fast extraction for <OffthreadVideo>](/docs/renderer/get-can-extract-frames-fast)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/ssr.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Deploy to a VPS](/docs/studio/deploy-server) [Next\
\
Overview](/docs/ssr)

- [Render a video on AWS Lambda](#render-a-video-on-aws-lambda)
- [Render a video using Node.js APIs](#render-a-video-using-nodejs-apis)
- [Render using GitHub Actions](#render-using-github-actions)
- [Render a video using Docker](#render-a-video-using-docker)
- [Render a video using GCP Cloud Run (Alpha)](#render-a-video-using-gcp-cloud-run-alpha)
- [API reference](#api-reference)
