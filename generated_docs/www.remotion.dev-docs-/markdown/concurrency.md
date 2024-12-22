---
title: Concurrency
source: Remotion Documentation
last_updated: 2024-12-22
---

# Concurrency

- [Home page](/)
- [Lambda](/docs/lambda)
- Concurrency

On this page

# Concurrency

Remotion Lambda is a highly concurrent distributed video rendering system. That means that the video rendering work is split up across many Lambda functions. How many Lambda functions exactly can be defined by you, or you let Remotions defaults decide.

## Setting the concurrency [​](\#setting-the-concurrency "Direct link to Setting the concurrency")

You can set the concurrency via the [`framesPerLambda`](/docs/lambda/rendermediaonlambda#framesperlambda) option (or [`--frames-per-lambda`](/docs/lambda/cli/render#--frames-per-lambda) via CLI).

The concurrency is defined as `frameCount / framesPerLambda`. That means that the higher you set `framesPerLambda`, the lower the concurrency gets.

note

Example: You render a video that has a `durationInFrames` of `300` with a `framePerLambda` setting of `15`. The concurrency is `300 / 15 = 20`.

## Default values [​](\#default-values "Direct link to Default values")

By default, Remotion chooses a value between 20 and ∞ for `framesPerLambda`. The longer the video, the higher the concurrency. As a baseline, no matter how short the video, always at least 20 frames are rendered per Lambda.

The following chart shows how the `framesPerLambda` and the implied concurrency is chosen based on the frame count:

![](/img/concurrency-chart.svg)

The code for determining the `framesPerLambda` parameter is:

```

tsx

import { interpolate } from "remotion";

const bestFramesPerLambdaParam = (frameCount: number) => {
  // Between 0 and 10 minutes (at 30fps), interpolate the concurrency from 75 to 150
  const concurrency = interpolate(frameCount, [0, 18000], [75, 150], {
    extrapolateRight: "clamp",
  });

  // At least have 20 as a `framesPerLambda` value
  const framesPerLambda = Math.max(frameCount / concurrency, 20);

  // Evenly distribute: For 21 frames over 2 lambda functions, distribute as 11 + 10 ==> framesPerLambda = 11
  const lambdasNeeded = Math.ceil(frameCount / framesPerLambda);

  return Math.ceil(frameCount / lambdasNeeded);
};
```

```

tsx

import { interpolate } from "remotion";

const bestFramesPerLambdaParam = (frameCount: number) => {
  // Between 0 and 10 minutes (at 30fps), interpolate the concurrency from 75 to 150
  const concurrency = interpolate(frameCount, [0, 18000], [75, 150], {
    extrapolateRight: "clamp",
  });

  // At least have 20 as a `framesPerLambda` value
  const framesPerLambda = Math.max(frameCount / concurrency, 20);

  // Evenly distribute: For 21 frames over 2 lambda functions, distribute as 11 + 10 ==> framesPerLambda = 11
  const lambdasNeeded = Math.ceil(frameCount / framesPerLambda);

  return Math.ceil(frameCount / lambdasNeeded);
};
```

Enter number of frames:

The default `framesPerLambda` is 20

## Concurrency limits [​](\#concurrency-limits "Direct link to Concurrency limits")

Ensure that you only set parameter within these limits to ensure the renders don't throw any errors:

- Minimum `framesPerLambda`: 4
- Maximum concurrency: 200

The Remotion Lambda defaults will never go outside these bounds.

## "Too many functions" [​](\#too-many-functions "Direct link to \"Too many functions\"")

If you get an error:

> Too many functions: This render would cause \[X\] functions to spawn. We limit this amount to 200 functions as more would result in diminishing returns.

You have set a value for `framesPerLambda` that is very low and would cause many functions to be spawned. In our experience, renders will not become faster if the concurrency is increased beyond this point.

- We recommend setting the `framesPerLambda` value to `null`. Remotion will choose a reasonable value that stays within the bounds.
- If you don't want to use the default, ensure that you don't set values that go outside of the bounds defined above.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/concurrency.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Region selection](/docs/lambda/region-selection) [Next\
\
Runtime](/docs/lambda/runtime)

- [Setting the concurrency](#setting-the-concurrency)
- [Default values](#default-values)
- [Concurrency limits](#concurrency-limits)
- ["Too many functions"](#too-many-functions)
