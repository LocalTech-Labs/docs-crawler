---
title: The fundamentals
source: Remotion Documentation
last_updated: 2024-12-22
---

# The fundamentals

- [Home page](/)
- Getting started
- The fundamentals

On this page

# The fundamentals

## React components [​](\#react-components "Direct link to React components")

The idea of Remotion is to give you a frame number and a blank canvas, to which you can render anything you want using [React](https://react.dev). This is an example React component that renders the current frame as text:

```

MyComposition.tsx
tsx

import { AbsoluteFill, useCurrentFrame } from "remotion";

export const MyComposition = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        fontSize: 100,
        backgroundColor: "white",
      }}
    >
      The current frame is {frame}.
    </AbsoluteFill>
  );
};
```

```

MyComposition.tsx
tsx

import { AbsoluteFill, useCurrentFrame } from "remotion";

export const MyComposition = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill
      style={{
        justifyContent: "center",
        alignItems: "center",
        fontSize: 100,
        backgroundColor: "white",
      }}
    >
      The current frame is {frame}.
    </AbsoluteFill>
  );
};
```

A video is a function of images over time. If you change content every frame, you'll end up with an animation.

## Video properties [​](\#video-properties "Direct link to Video properties")

A video has 4 properties:

- `width` in pixels.
- `height` in pixels.
- `durationInFrames`: the number of frames which the video is long.
- `fps`: framerate of the video.

You can obtain these values from the [`useVideoConfig()`](/docs/use-video-config) hook:

```

tsx

import {AbsoluteFill, useVideoConfig} from 'remotion';

export const MyComposition = () => {
  const {fps, durationInFrames, width, height} = useVideoConfig();

  return (
    <AbsoluteFill
      style={{
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: 60,
        backgroundColor: 'white',
      }}
    >
      This {width}x{height}px video is {durationInFrames / fps} seconds long.
    </AbsoluteFill>
  );
};
```

```

tsx

import {AbsoluteFill, useVideoConfig} from 'remotion';

export const MyComposition = () => {
  const {fps, durationInFrames, width, height} = useVideoConfig();

  return (
    <AbsoluteFill
      style={{
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: 60,
        backgroundColor: 'white',
      }}
    >
      This {width}x{height}px video is {durationInFrames / fps} seconds long.
    </AbsoluteFill>
  );
};
```

note

A video's first frame is `0` and its last frame is `durationInFrames - 1`.

## Compositions [​](\#compositions "Direct link to Compositions")

A composition is the combination of a React component and video metadata.

By rendering a [`<Composition>`](/docs/composition) component in `src/Root.tsx`, you can register a renderable video and make it visible in the Remotion sidebar.

```

src/Root.tsx
tsx

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="MyComposition"
      durationInFrames={150}
      fps={30}
      width={1920}
      height={1080}
      component={MyComposition}
    />
  );
};
```

```

src/Root.tsx
tsx

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="MyComposition"
      durationInFrames={150}
      fps={30}
      width={1920}
      height={1080}
      component={MyComposition}
    />
  );
};
```

note

You can register multiple compositions in `src/Root.tsx` by wrapping them in a React Fragment:

`<><Composition/><Composition/></>`. See also: [How to combine compositions?](/docs/miscellaneous/snippets/combine-compositions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/the-fundamentals.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Installation](/docs/) [Next\
\
Animating properties](/docs/animating-properties)

- [React components](#react-components)
- [Video properties](#video-properties)
- [Compositions](#compositions)
