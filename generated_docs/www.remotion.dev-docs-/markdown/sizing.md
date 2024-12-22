---
title: Sizing
source: Remotion Documentation
last_updated: 2024-12-22
---

# Sizing

- [Home page](/)
- [Player](/docs/player/)
- Sizing

On this page

# Sizing

The following algorithm is used for calculating the size of the Player:

[1](#1)

By default, the Player is as big as the composition height,
defined by the `compositionHeight` and `compositionWidth` props.

[2](#2) If `height` and `width` is defined using
the `style` property, the player will assume the dimensions you have
passed.

[3](#3) If a `height` is passed using the `style` property,
the player will assume that height, and calculate the width based on the aspect
ratio of the video.

[4](#4) If `width` is passed using the `style` property,
the player will assume that width and calculate the height based on the aspect
ration of the video.

note

Before v3.3.43, if case

3

or

4

happened, a layout shift would occur during mounting because the element was measured. Using a newer version of Remotion will fix this, because it uses the `aspect-ratio` CSS property.

## Full width [​](\#full-width "Direct link to Full width")

By setting the following style:

```

tsx

style={{ width: "100%" }}
```

```

tsx

style={{ width: "100%" }}
```

The video will scale to the full width of the parent container, while the height will be calculated based on the aspect ratio of the video.

## Fitting to a container [​](\#fitting-to-a-container "Direct link to Fitting to a container")

This is how you can make the Player fill out a container but keep the aspect ratio of the video:

```

Canvas.tsx
tsx

import {Player} from '@remotion/player';
import React from 'react';
import {AbsoluteFill} from 'remotion';

const MyComp: React.FC = () => {
  return <AbsoluteFill style={{backgroundColor: 'black'}} />;
};

export const FullscreenPlayer = () => {
  const compositionWidth = 300;
  const compositionHeight = 200;

  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        backgroundColor: 'gray',
        // Any container with "position: relative" will work
        position: 'relative',
      }}
    >
      <div
        style={{
          position: 'absolute',
          top: 0,
          left: 0,
          right: 0,
          bottom: 0,
          margin: 'auto',
          aspectRatio: `${compositionWidth} / ${compositionHeight}`,
          maxHeight: '100%',
          maxWidth: '100%',
        }}
      >
        <Player
          controls
          component={MyComp}
          compositionWidth={compositionWidth}
          compositionHeight={compositionHeight}
          durationInFrames={200}
          fps={30}
          style={{
            width: '100%',
          }}
        />
      </div>
    </div>
  );
};
```

```

Canvas.tsx
tsx

import {Player} from '@remotion/player';
import React from 'react';
import {AbsoluteFill} from 'remotion';

const MyComp: React.FC = () => {
  return <AbsoluteFill style={{backgroundColor: 'black'}} />;
};

export const FullscreenPlayer = () => {
  const compositionWidth = 300;
  const compositionHeight = 200;

  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        backgroundColor: 'gray',
        // Any container with "position: relative" will work
        position: 'relative',
      }}
    >
      <div
        style={{
          position: 'absolute',
          top: 0,
          left: 0,
          right: 0,
          bottom: 0,
          margin: 'auto',
          aspectRatio: `${compositionWidth} / ${compositionHeight}`,
          maxHeight: '100%',
          maxWidth: '100%',
        }}
      >
        <Player
          controls
          component={MyComp}
          compositionWidth={compositionWidth}
          compositionHeight={compositionHeight}
          durationInFrames={200}
          fps={30}
          style={{
            width: '100%',
          }}
        />
      </div>
    </div>
  );
};
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/scaling.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Examples](/docs/player/examples) [Next\
\
Code sharing](/docs/player/integration)

- [Full width](#full-width)
- [Fitting to a container](#fitting-to-a-container)
