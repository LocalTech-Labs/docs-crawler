---
title: Layers
source: Remotion Documentation
last_updated: 2024-12-22
---

# Layers

- [Home page](/)
- Designing visuals
- Layers

On this page

# Layers

Unlike normal websites, a video has fixed dimensions. That means, it is okay to use `position: "absolute"`!

In Remotion, you often want to "layer" elements on top of each other.

This is a common pattern in video editors, and in Photoshop.

An easy way to do it is using the [`<AbsoluteFill>`](/docs/absolute-fill) component:

```

MyComp.tsx
tsx

import React from 'react';
import {AbsoluteFill, Img, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <Img src={staticFile('bg.png')} />
      </AbsoluteFill>
      <AbsoluteFill>
        <h1>This text appears on top of the video!</h1>
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {AbsoluteFill, Img, staticFile} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <Img src={staticFile('bg.png')} />
      </AbsoluteFill>
      <AbsoluteFill>
        <h1>This text appears on top of the video!</h1>
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

This will render the text on top of the image.

If you want to only show an element for a certain duration, you can use a [`<Sequence>`](/docs/sequence) component, which by default is also absolutely positioned.

```

MyComp.tsx
tsx

import React from 'react';
import {AbsoluteFill, Img, staticFile, Sequence} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Sequence>
        <Img src={staticFile('bg.png')} />
      </Sequence>
      <Sequence from={60} durationInFrames={40}>
        <h1>This text appears after 60 frames!</h1>
      </Sequence>
    </AbsoluteFill>
  );
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {AbsoluteFill, Img, staticFile, Sequence} from 'remotion';

export const MyComp: React.FC = () => {
  return (
    <AbsoluteFill>
      <Sequence>
        <Img src={staticFile('bg.png')} />
      </Sequence>
      <Sequence from={60} durationInFrames={40}>
        <h1>This text appears after 60 frames!</h1>
      </Sequence>
    </AbsoluteFill>
  );
};
```

The lower the absolutely positioned element is in the tree, the higher it will be in the layer stack.

If you are aware of this behavior, you can use it to your advantage and avoid using `z-index` in most cases.

## See also [​](\#see-also "Direct link to See also")

- [`<AbsoluteFill>`](/docs/absolute-fill)
- [`<Sequence>`](/docs/sequence)
- [Build a timeline-based video editor](/docs/building-a-timeline)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/layers.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Assets](/docs/assets) [Next\
\
Transitions](/docs/transitioning)

- [See also](#see-also)
