---
title: Timings
source: Remotion Documentation
last_updated: 2024-12-22
---

# Timings

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- Timings

On this page

# Timings

A **timing** together with a presentation forms a transition.

Remotion provides certain timings out of the box, but you can also [create your own](/docs/transitions/timings/custom).

## Available timings [​](\#available-timings "Direct link to Available timings")

[**`springTiming()`** \
\
Transition with a `spring()`](/docs/transitions/timings/springtiming) [**`linearTiming()`** \
\
Transition linearly with optional Easing](/docs/transitions/timings/lineartiming) [**Custom timings** \
\
Implement your own timing](/docs/transitions/timings/custom)

## Getting the duration of a timing [​](\#getting-the-duration-of-a-timing "Direct link to Getting the duration of a timing")

You can get the duration of a transition by calling `getDurationInFrames()` on the timing:

```

Assuming a framerate of 30fps
tsx

import { springTiming } from "@remotion/transitions";

springTiming({ config: { damping: 200 } }).getDurationInFrames({ fps: 30 }); // 23
```

```

Assuming a framerate of 30fps
tsx

import { springTiming } from "@remotion/transitions";

springTiming({ config: { damping: 200 } }).getDurationInFrames({ fps: 30 }); // 23
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/timings/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useTransitionProgress()](/docs/transitions/use-transition-progress) [Next\
\
springTiming()](/docs/transitions/timings/springtiming)

- [Available timings](#available-timings)
- [Getting the duration of a timing](#getting-the-duration-of-a-timing)
