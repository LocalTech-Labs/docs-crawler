---
title: How Remotion Lambda works
source: Remotion Documentation
last_updated: 2024-12-22
---

# How Remotion Lambda works

- [Home page](/)
- [Lambda](/docs/lambda)
- How Remotion Lambda works

On this page

# How Remotion Lambda works

This document describes the procedure that gets executed when a Remotion Lambda video render is triggered.

note

This document explains the Lambda architecture from version 4.0.165 on.

Previously, Lambda functions did not use Response Streaming, and instead saved chunks to S3.

[1](#1)

A single Lambda function is invoked using [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) \- either directly or via the CLI which also calls this API. This invocation
is called the **main function**.

[2](#2) The main function visits the [Serve URL](/docs/terminology/serve-url) that it is being passed in a headless browser.

[3](#3) The main function finds the composition based on the composition ID
that was passed and runs the [props resolution](/docs/props-resolution) algorithm to determine the props that should be passed to the video as well
as the metadata (such as the duration in frames).

[4](#4) Based on the determined duration of the video and the [concurrency](/docs/lambda/concurrency), several **renderer functions** are spawned, which are tasked to
render a portion of the video.

[5](#5) The renderer functions use [AWS Lambda Response Streaming](https://aws.amazon.com/blogs/compute/introducing-aws-lambda-response-streaming/) to report progress as well as the binary video chunks.

[6](#6) The main function concatenates the progress reports into a
concise `progress.json` file and periodically uploads it to S3.

[7](#7) The [`getRenderProgress()`](/docs/lambda/getrenderprogress) API queries the S3 bucket for the `progress.json` file and returns
the progress of the render.

[8](#8) As soon as all chunks have arrived in the main function, they
get [seamlessly concatenated](/blog/faster-lambda). The concatenation
algorithm is not a public API at the moment.

[9](#9) The main function uploads the final video to S3 and shuts down.

## FAQ [​](\#faq "Direct link to FAQ")

### Can I roll my own distributed renderer? [​](\#can-i-roll-my-own-distributed-renderer "Direct link to Can I roll my own distributed renderer?")

The seamless concatenation of chunks is not a public API at the moment.

You may render chunks using [`frameRange`](/docs/renderer/render-media#framerange) and [`audioCodec: "pcm-16"`](/docs/renderer/render-media#audiocodec) which you can concatenate using FFmpeg.

Building a distributed renderer is hard, and not recommended for most.

### Will each chunk download all assets? [​](\#will-each-chunk-download-all-assets "Direct link to Will each chunk download all assets?")

Each chunk will download all assets that are referenced in this chunk.

This can lead to assets being downloaded many times at the same time, which may overwhelm a server or trigger rate limits.

In addition, you pay for the bandwidth, even if the assets are on S31).

Keep this in mind when designing your solution and consider using a CDN to serve assets.

* * *

1) An API for avoiding the S3 bandwidth charge is [planned](https://github.com/remotion-dev/remotion/issues/3817).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/how-lambda-works.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Multiple buckets](/docs/lambda/multiple-buckets) [Next\
\
Bucket naming](/docs/lambda/bucket-naming)

- [FAQ](#faq)
  - [Can I roll my own distributed renderer?](#can-i-roll-my-own-distributed-renderer)
  - [Will each chunk download all assets?](#will-each-chunk-download-all-assets)
