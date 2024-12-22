---
title: Optimizing for speed
source: Remotion Documentation
last_updated: 2024-12-22
---

# Optimizing for speed

- [Home page](/)
- [Lambda](/docs/lambda)
- Optimizing for speed

On this page

# Optimizing for speed

On this page, a few strategies are presented for making a render on Lambda as fast as possible.

## Higher concurrency [​](\#higher-concurrency "Direct link to Higher concurrency")

Generally, a lower value for [`framesPerLambda`](/docs/lambda/rendermediaonlambda#framesperlambda) will result in higher concurrency and therefore an opportunity to finish the render in less absolute time. At the same time, more overhead will be produced, making the render more expensive. You will also experience diminishing returns, and adding too much concurrency can make the speed slower because the overhead of orchestrating many Lambda functions outweigh the gains.
See the [Lambda Concurrency](/docs/lambda/concurrency) page for more information.

## More memory [​](\#more-memory "Direct link to More memory")

Adding more memory on Lambda will also scale up the CPU power on Lambda proportionally, therefore making the render faster. At the same time, the cost also linearly increases with the memory you add.

## `concurrencyPerLambda` property [​](\#concurrencyperlambda-property "Direct link to concurrencyperlambda-property")

The [`concurrencyPerLambda`](/docs/lambda/rendermediaonlambda#concurrencyperlambda) property in [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) allows you to open multiple browser tabs in a single Lambda function, therefore opening an opportunity to do more work at once. If the Lambda function is too busy, increasing the concurrency might also be counterproductive.

## Use `speculateFunctionName()` [​](\#use-speculatefunctionname "Direct link to use-speculatefunctionname")

Instead of calling [`getFunctions()`](/docs/lambda/getfunctions), you can call [`speculateFunctionName()`](/docs/lambda/speculatefunctionname) to calculate the name of the function you are about to call to save an API call and save up to 1 second.

## Bucket naming [​](\#bucket-naming "Direct link to Bucket naming")

If you have a Remotion version before December 2022, then your bucket name might not include the region name in its name. This will result in Remotion having to list all bucket names and query their region before kicking off the render.

[See this article for more information](/docs/lambda/bucket-naming#aws-region-in-the-name). Consider renaming your bucket or re-setting up Remotion Lambda to gain speed. This especially applies if you are having many Remotion buckets across regions.

## Use MP3 as an audio codec [v4.0.16](https://github.com/remotion-dev/remotion/releases/v4.0.16) [​](\#use-mp3-as-an-audio-codec "Direct link to use-mp3-as-an-audio-codec")

By default, a video renders with the `h264` codec and the `aac` audio codec. Setting the [`audioCodec`](/docs/encoding/#audio-codec) to `mp3` will make "Combining videos" stage a lot faster, as the MP3 codec is much faster to encode than AAC. However, **the audio will not play in QuickTime Player and the file size is minimally higher**.

## Optimizing render performance [​](\#optimizing-render-performance "Direct link to Optimizing render performance")

See the [general performance tips](/docs/performance) which also apply to Lambda.

## See also [​](\#see-also "Direct link to See also")

- [Optimizing for cost](/docs/lambda/optimizing-cost)
- [Performance](/docs/performance)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/speed.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Optimizing for cost](/docs/lambda/optimizing-cost) [Next\
\
Limits](/docs/lambda/limits)

- [Higher concurrency](#higher-concurrency)
- [More memory](#more-memory)
- [`concurrencyPerLambda` property](#concurrencyperlambda-property)
- [Use `speculateFunctionName()`](#use-speculatefunctionname)
- [Bucket naming](#bucket-naming)
- [Use MP3 as an audio codec](#use-mp3-as-an-audio-codec)
- [Optimizing render performance](#optimizing-render-performance)
- [See also](#see-also)
