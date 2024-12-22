---
title: useTransitionProgress()v4.0.177
source: Remotion Documentation
last_updated: 2024-12-22
---

# useTransitionProgress()v4.0.177

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- useTransitionProgress()

On this page

# useTransitionProgress() [v4.0.177](https://github.com/remotion-dev/remotion/releases/v4.0.177)

A hook that can be used inside a child of a [`<TransitionSeries.Sequence>`](/docs/transitions/transitionseries) to get the progress of the transition to directly manipulate the objects inside the scene.

It is meant to be used together with the [`none()`](/docs/transitions/presentations/none) presentation, but can be used with any other presentation.

## Example [​](\#example "Direct link to Example")

```

useTransitionProgress()
tsx

import { useTransitionProgress } from "@remotion/transitions";
import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { none } from "@remotion/transitions/none";

const A: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Always `1`
  console.log(progress.exiting); // Going from 0 to 1
  console.log(progress.isInTransitionSeries); //  `true`

  return <div>A</div>;
};

const B: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Going from 0 to 1
  console.log(progress.exiting); // Always `0`
  console.log(progress.isInTransitionSeries); //  `true`

  return <div>B</div>;
};

const C: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Always `1`
  console.log(progress.exiting); // Always `0`
  console.log(progress.isInTransitionSeries); //  `false`

  return <div>C</div>;
};

const Transition: React.FC = () => {
  return (
    <>
      <TransitionSeries>
        <TransitionSeries.Sequence durationInFrames={40}>
          <A />
        </TransitionSeries.Sequence>
        <TransitionSeries.Transition
          presentation={none()}
          timing={linearTiming({ durationInFrames: 30 })}
        />
        <TransitionSeries.Sequence durationInFrames={60}>
          <B />
        </TransitionSeries.Sequence>
      </TransitionSeries>
      <C />
    </>
  );
};
```

```

useTransitionProgress()
tsx

import { useTransitionProgress } from "@remotion/transitions";
import { linearTiming, TransitionSeries } from "@remotion/transitions";
import { none } from "@remotion/transitions/none";

const A: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Always `1`
  console.log(progress.exiting); // Going from 0 to 1
  console.log(progress.isInTransitionSeries); //  `true`

  return <div>A</div>;
};

const B: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Going from 0 to 1
  console.log(progress.exiting); // Always `0`
  console.log(progress.isInTransitionSeries); //  `true`

  return <div>B</div>;
};

const C: React.FC = () => {
  const progress = useTransitionProgress();
  console.log(progress.entering); // Always `1`
  console.log(progress.exiting); // Always `0`
  console.log(progress.isInTransitionSeries); //  `false`

  return <div>C</div>;
};

const Transition: React.FC = () => {
  return (
    <>
      <TransitionSeries>
        <TransitionSeries.Sequence durationInFrames={40}>
          <A />
        </TransitionSeries.Sequence>
        <TransitionSeries.Transition
          presentation={none()}
          timing={linearTiming({ durationInFrames: 30 })}
        />
        <TransitionSeries.Sequence durationInFrames={60}>
          <B />
        </TransitionSeries.Sequence>
      </TransitionSeries>
      <C />
    </>
  );
};
```

## API [​](\#api "Direct link to API")

A React hook that returns an object with the following properties:

### `entering` [​](\#entering "Direct link to entering")

The progress of the entering scene. If not inside a [`<TransitionSeries.Sequence>`](/docs/transitions/transitionseries), it will always be `1`.

### `exiting` [​](\#exiting "Direct link to exiting")

The progress of the exiting scene. If not inside a [`<TransitionSeries.Sequence>`](/docs/transitions/transitionseries), it will always be `0`.

### `isInTransitionSeries` [​](\#isintransitionseries "Direct link to isintransitionseries")

Whether the current scene is inside a [`<TransitionSeries.Sequence>`](/docs/transitions/transitionseries).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this hook](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/use-transition-progress.ts)
- [`<TransitionSeries>`](/docs/transitions/transitionseries)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/use-transition-progress.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<TransitionSeries>](/docs/transitions/transitionseries) [Next\
\
Timings](/docs/transitions/timings/)

- [Example](#example)
- [API](#api)
  - [`entering`](#entering)
  - [`exiting`](#exiting)
  - [`isInTransitionSeries`](#isintransitionseries)
- [See also](#see-also)
