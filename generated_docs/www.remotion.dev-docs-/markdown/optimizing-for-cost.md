---
title: Optimizing for cost
source: Remotion Documentation
last_updated: 2024-12-22
---

# Optimizing for cost

- [Home page](/)
- [Lambda](/docs/lambda)
- Optimizing for cost

On this page

# Optimizing for cost

On this page, a few strategies for optimizing the cost of using Remotion Lambda are presented.

## Lower concurrency [​](\#lower-concurrency "Direct link to Lower concurrency")

A lot of compute time is spent on warming up the Lambda function, opening browsers, uploading and downloading assets. By using less Lambda functions, less overhead is being produced which will ultimately result in lower cost, however also in slower render speeds. See the [Lambda Concurrency](/docs/lambda/concurrency) page for more information.

## Using cheaper regions [​](\#using-cheaper-regions "Direct link to Using cheaper regions")

Not all AWS regions have the same cost. See the [AWS Lambda](https://aws.amazon.com/lambda/pricing/) pricing chart to see if you are using a region that is more expensive.

## Pre-compute data [​](\#pre-compute-data "Direct link to Pre-compute data")

Consider whether your computation will run on every invoked Lambda function. If possible, pre-compute data once and pass it as [input props](/docs/passing-props#passing-input-props-in-the-cli) to the render.

## Less memory [​](\#less-memory "Direct link to Less memory")

Reducing the memory will linearly decrease the cost. Try out different values for the memory when deploying a Lambda function and see how low you can set it without experiencing a crash.

## Make the render more performant [​](\#make-the-render-more-performant "Direct link to Make the render more performant")

Making your render more efficient will also reduce the cost of using Lambda. See the [general performance tips](/docs/performance).

## Real life optimization example [​](\#real-life-optimization-example "Direct link to Real life optimization example")

See [this video](https://www.youtube.com/watch?v=GUsjj1jsLhw) of how we optimized the cost of a Lambda function by finding the bottlenecks and eliminating them.

## See also [​](\#see-also "Direct link to See also")

- [How much does Remotion Lambda cost?](/docs/lambda/cost-example)
- [Optimizing for speed](/docs/lambda/optimizing-speed)
- [Performance](/docs/performance)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/cost.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Bucket naming](/docs/lambda/bucket-naming) [Next\
\
Optimizing for speed](/docs/lambda/optimizing-speed)

- [Lower concurrency](#lower-concurrency)
- [Using cheaper regions](#using-cheaper-regions)
- [Pre-compute data](#pre-compute-data)
- [Less memory](#less-memory)
- [Make the render more performant](#make-the-render-more-performant)
- [Real life optimization example](#real-life-optimization-example)
- [See also](#see-also)
