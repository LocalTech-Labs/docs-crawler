---
title: Freeze portions of a sequence
source: Remotion Documentation
last_updated: 2024-12-22
---

# Freeze portions of a sequence

- [Home page](/)
- Embedding videos
- Freeze a portion

# Freeze portions of a sequence

With the following snippet, two portions of a sequence are frozen and resumed afterwards from the point they were paused.

```

tsx

import React, {useMemo} from 'react';
import {Freeze, Sequence, useCurrentFrame} from 'remotion';

export const Counter: React.FC = () => {
  return (
    <div className="flex h-full justify-center items-center text-6xl">
      {useCurrentFrame()}
    </div>
  );
};

const FREEZES = [
  {
    frame: 0,
    durationInFrames: 25,
  },
  {
    frame: 30,
    durationInFrames: 50,
  },
];

const getFreezes = () => {
  let summedUpFreezes = 0;
  const freezeFrames = [];
  for (const freeze of FREEZES) {
    freezeFrames.push({
      start: summedUpFreezes + freeze.frame,
      durationInFrames: freeze.durationInFrames,
      from: summedUpFreezes,
      frame: freeze.frame,
    });
    summedUpFreezes += freeze.durationInFrames;
  }
  return freezeFrames;
};

export const FreezePortion: React.FC = () => {
  const freezes = useMemo(() => {
    return getFreezes();
  }, []);
  const frame = useCurrentFrame();

  const nextFreeze = freezes.find(
    (freeze) => frame < freeze.start + freeze.durationInFrames,
  );
  const activeFreeze = freezes.find(
    (freeze) =>
      frame >= freeze.start && frame < freeze.start + freeze.durationInFrames,
  );

  const from = useMemo(() => {
    if (activeFreeze) {
      return 0;
    }

    if (nextFreeze) {
      return nextFreeze.from;
    }

    return (
      freezes[freezes.length - 1].from +
      freezes[freezes.length - 1].durationInFrames
    );
  }, [activeFreeze, freezes, nextFreeze]);

  return (
    <Freeze
      frame={activeFreeze ? activeFreeze.frame : 0}
      active={Boolean(activeFreeze)}
    >
      <Sequence layout="none" from={from}>
        <Counter />
      </Sequence>
    </Freeze>
  );
};
```

```

tsx

import React, {useMemo} from 'react';
import {Freeze, Sequence, useCurrentFrame} from 'remotion';

export const Counter: React.FC = () => {
  return (
    <div className="flex h-full justify-center items-center text-6xl">
      {useCurrentFrame()}
    </div>
  );
};

const FREEZES = [
  {
    frame: 0,
    durationInFrames: 25,
  },
  {
    frame: 30,
    durationInFrames: 50,
  },
];

const getFreezes = () => {
  let summedUpFreezes = 0;
  const freezeFrames = [];
  for (const freeze of FREEZES) {
    freezeFrames.push({
      start: summedUpFreezes + freeze.frame,
      durationInFrames: freeze.durationInFrames,
      from: summedUpFreezes,
      frame: freeze.frame,
    });
    summedUpFreezes += freeze.durationInFrames;
  }
  return freezeFrames;
};

export const FreezePortion: React.FC = () => {
  const freezes = useMemo(() => {
    return getFreezes();
  }, []);
  const frame = useCurrentFrame();

  const nextFreeze = freezes.find(
    (freeze) => frame < freeze.start + freeze.durationInFrames,
  );
  const activeFreeze = freezes.find(
    (freeze) =>
      frame >= freeze.start && frame < freeze.start + freeze.durationInFrames,
  );

  const from = useMemo(() => {
    if (activeFreeze) {
      return 0;
    }

    if (nextFreeze) {
      return nextFreeze.from;
    }

    return (
      freezes[freezes.length - 1].from +
      freezes[freezes.length - 1].durationInFrames
    );
  }, [activeFreeze, freezes, nextFreeze]);

  return (
    <Freeze
      frame={activeFreeze ? activeFreeze.frame : 0}
      active={Boolean(activeFreeze)}
    >
      <Sequence layout="none" from={from}>
        <Counter />
      </Sequence>
    </Freeze>
  );
};
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/freeze-portions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Jump Cuts](/docs/miscellaneous/snippets/jumpcuts) [Next\
\
HTTP Live Streaming](/docs/miscellaneous/snippets/hls)
