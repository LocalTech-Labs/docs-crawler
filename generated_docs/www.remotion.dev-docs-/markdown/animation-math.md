---
title: Animation math
source: Remotion Documentation
last_updated: 2024-12-22
---

# Animation math

- [Home page](/)
- Designing visuals
- Animation math

# Animation math

You can add, subtract and multiply animation values to create more complex animations.

Consider the following example:

```

Enter and exit
tsx

import { spring, useCurrentFrame, useVideoConfig } from "remotion";

const frame = useCurrentFrame();
const { fps, durationInFrames } = useVideoConfig();

const enter = spring({
  fps,
  frame,
  config: {
    damping: 200,
  },
});

const exit = spring({
  fps,
  config: {
    damping: 200,
  },
  durationInFrames: 20,
  delay: durationInFrames - 20,
  frame,
});

const scale = enter - exit;
```

```

Enter and exit
tsx

import { spring, useCurrentFrame, useVideoConfig } from "remotion";

const frame = useCurrentFrame();
const { fps, durationInFrames } = useVideoConfig();

const enter = spring({
  fps,
  frame,
  config: {
    damping: 200,
  },
});

const exit = spring({
  fps,
  config: {
    damping: 200,
  },
  durationInFrames: 20,
  delay: durationInFrames - 20,
  frame,
});

const scale = enter - exit;
```

- At the beginning of the animation, the value of `enter` is `0`, it goes to `1` over the course of the animation.
- Before the sequence ends, we create an `exit` animation that goes from `0` to `1`.
- Subtracting the `exit` animation from the `enter` animation gives us the overall state of the animation which we use to animate `scale`.

0

```

Full snippet
tsx

import React from "react";
import {
  AbsoluteFill,
  spring,
  useCurrentFrame,
  useVideoConfig,
} from "remotion";

export const AnimationMath: React.FC = () => {
  const frame = useCurrentFrame();
  const { fps, durationInFrames } = useVideoConfig();

  const enter = spring({
    fps,
    frame,
    config: {
      damping: 200,
    },
  });

  const exit = spring({
    fps,
    config: {
      damping: 200,
    },
    durationInFrames: 20,
    delay: durationInFrames - 20,
    frame,
  });

  const scale = enter - exit;

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "white",
      }}
    >
      <div
        style={{
          height: 100,
          width: 100,
          backgroundColor: "#4290f5",
          borderRadius: 20,
          transform: `scale(${scale})`,
          justifyContent: "center",
          alignItems: "center",
          display: "flex",
          fontSize: 50,
          color: "white",
        }}
      >
        {frame}
      </div>
    </AbsoluteFill>
  );
};
```

```

Full snippet
tsx

import React from "react";
import {
  AbsoluteFill,
  spring,
  useCurrentFrame,
  useVideoConfig,
} from "remotion";

export const AnimationMath: React.FC = () => {
  const frame = useCurrentFrame();
  const { fps, durationInFrames } = useVideoConfig();

  const enter = spring({
    fps,
    frame,
    config: {
      damping: 200,
    },
  });

  const exit = spring({
    fps,
    config: {
      damping: 200,
    },
    durationInFrames: 20,
    delay: durationInFrames - 20,
    frame,
  });

  const scale = enter - exit;

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "white",
      }}
    >
      <div
        style={{
          height: 100,
          width: 100,
          backgroundColor: "#4290f5",
          borderRadius: 20,
          transform: `scale(${scale})`,
          justifyContent: "center",
          alignItems: "center",
          display: "flex",
          fontSize: 50,
          color: "white",
        }}
      >
        {frame}
      </div>
    </AbsoluteFill>
  );
};
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/animation-math.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Noise visualization](/docs/noise-visualization) [Next\
\
Adding a video](/docs/videos/)
