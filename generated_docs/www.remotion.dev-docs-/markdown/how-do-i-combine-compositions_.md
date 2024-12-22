---
title: How do I combine compositions?
source: Remotion Documentation
last_updated: 2024-12-22
---

# How do I combine compositions?

- [Home page](/)
- Snippets
- Combining compositions

On this page

# How do I combine compositions?

Say you have two compositions, `One` and `Two`:

```

Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { One } from "./One";
import { Two } from "./Two";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="One"
        component={One}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Two"
        component={Two}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
    </>
  );
};
```

```

Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { One } from "./One";
import { Two } from "./Two";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="One"
        component={One}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Two"
        component={Two}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
    </>
  );
};
```

To combine them, create another component, let's call it `Main`:

```

Main.tsx
tsx

import React from "react";
import { Series } from "remotion";
import { One } from "./One";
import { Two } from "./Two";

export const Main: React.FC = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={120}>
        <One />
      </Series.Sequence>
      <Series.Sequence durationInFrames={120}>
        <Two />
      </Series.Sequence>
    </Series>
  );
};
```

```

Main.tsx
tsx

import React from "react";
import { Series } from "remotion";
import { One } from "./One";
import { Two } from "./Two";

export const Main: React.FC = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={120}>
        <One />
      </Series.Sequence>
      <Series.Sequence durationInFrames={120}>
        <Two />
      </Series.Sequence>
    </Series>
  );
};
```

Then, in your `Root` component, add a `Main` composition:

```

Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { One } from "./One";
import { Two } from "./Two";
import { Main } from "./Main";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="One"
        component={One}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Two"
        component={Two}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Main"
        component={Main}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={240}
      />
    </>
  );
};
```

```

Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { One } from "./One";
import { Two } from "./Two";
import { Main } from "./Main";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="One"
        component={One}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Two"
        component={Two}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={120}
      />
      <Composition
        id="Main"
        component={Main}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={240}
      />
    </>
  );
};
```

## How do I avoid hardcoding the duration? [​](\#how-do-i-avoid-hardcoding-the-duration "Direct link to How do I avoid hardcoding the duration?")

You can define variables alongside your components:

```

tsx

export const ONE_DURATION = 120;
export const TWO_DURATION = 120;
export const MAIN_DURATION = ONE_DURATION + TWO_DURATION;
```

```

tsx

export const ONE_DURATION = 120;
export const TWO_DURATION = 120;
export const MAIN_DURATION = ONE_DURATION + TWO_DURATION;
```

And then pass the variables them to [`<Composition>`](/docs/composition) and [`<Series.Sequence>`](/docs/series).

## How do I transition between compositions? [​](\#how-do-i-transition-between-compositions "Direct link to How do I transition between compositions?")

You can use [`<TransitionSeries>`](/docs/transitions/transitionseries). If you use it, your main composition will get shorter because for a period of time, both compositions are mounted.

Subtract the [duration of the transition](/docs/transitions/timings/custom#getting-the-duration-of-a-timing) to get a correct timing.

## How does this scale? [​](\#how-does-this-scale "Direct link to How does this scale?")

As more scenes are added, consider:

- using [`calculateMetadata()`](/docs/calculate-metadata) to calculate the duration of a composition based on its props.
- creating an array of scenes and using JavaScripts [`.map()`](https://www.freecodecamp.org/news/how-to-render-lists-in-react/) function to render them.

## See also [​](\#see-also "Direct link to See also")

- [`<Series>` component](/docs/series)
- [`<TransitionSeries>` component](/docs/transitions/transitionseries)
- [Making components reusable](/docs/reusability)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/combine-compositions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<OffthreadVideo /> while rendering](/docs/miscellaneous/snippets/offthread-video-while-rendering) [Next\
\
Upgrading Remotion](/docs/upgrading)

- [How do I avoid hardcoding the duration?](#how-do-i-avoid-hardcoding-the-duration)
- [How do I transition between compositions?](#how-do-i-transition-between-compositions)
- [How does this scale?](#how-does-this-scale)
- [See also](#see-also)
