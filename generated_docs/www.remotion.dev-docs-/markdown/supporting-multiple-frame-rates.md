---
title: Supporting multiple frame rates
source: Remotion Documentation
last_updated: 2024-12-22
---

# Supporting multiple frame rates

- [Home page](/)
- Building apps
- Supporting multiple frame rates

On this page

# Supporting multiple frame rates

You may want to support multiple frame rates for a composition.

For example, to build an option allow the user to export the video at either 30 FPS or 60 FPS.

This document outlines the best practices for doing so.

## Write animations frame-rate-independently [​](\#write-animations-frame-rate-independently "Direct link to Write animations frame-rate-independently")

When animating over time, use the `fps` value from [`useVideoConfig()`](/docs/use-video-config) to calculate the current time.

The following animation will change the speed if the frame rate is changed:

```

ts

// Animate from second 1 to second 2
const animationProgress = interpolate(frame, [30, 60], [0, 1], {
  extrapolateLeft: 'clamp',
  extrapolateRight: 'clamp',
});
```

```

ts

// Animate from second 1 to second 2
const animationProgress = interpolate(frame, [30, 60], [0, 1], {
  extrapolateLeft: 'clamp',
  extrapolateRight: 'clamp',
});
```

It is better to make it frame-rate-dependnt:

```

ts

// Animate from second 1 to second 2
const {fps} = useVideoConfig();
const animationProgress = interpolate(frame, [1 * fps, 2 * fps], [0, 1], {
  extrapolateLeft: 'clamp',
  extrapolateRight: 'clamp',
});
```

```

ts

// Animate from second 1 to second 2
const {fps} = useVideoConfig();
const animationProgress = interpolate(frame, [1 * fps, 2 * fps], [0, 1], {
  extrapolateLeft: 'clamp',
  extrapolateRight: 'clamp',
});
```

Also use the `fps` value when specifying `from` and `durationInFrames` of a `<Sequence>`:

```

tsx

// Show for 3 seconds
<Sequence durationInFrames={3 * fps}>
  <div />
</Sequence>;
```

```

tsx

// Show for 3 seconds
<Sequence durationInFrames={3 * fps}>
  <div />
</Sequence>;
```

And when passing `fps`, `delay` and `durationInFrames` of a `spring()`:

```

ts

const animationProgress = spring({
  frame,
  fps,
  durationInFrames: 2 * fps,
  delay: 1 * fps,
});
```

```

ts

const animationProgress = spring({
  frame,
  fps,
  durationInFrames: 2 * fps,
  delay: 1 * fps,
});
```

## Change frame rate dynamically [​](\#change-frame-rate-dynamically "Direct link to Change frame rate dynamically")

Here is how you could switch the `fps` of a [`<Composition>`](/docs/composition) based on an [input prop](/docs/terminology/input-props):

```

tsx

import {Composition, useCurrentFrame, useVideoConfig} from 'remotion';

const VariableFps: React.FC<{
  frameRate: '30fps' | '60fps';
}> = () => {
  const {fps} = useVideoConfig();
  return <div>{fps} FPS</div>;
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="VariableFps"
      component={VariableFps}
      width={1080}
      height={1080}
      durationInFrames={100}
      calculateMetadata={({props}) => {
        return {
          fps: props.frameRate === '60fps' ? 60 : 30,
        };
      }}
      defaultProps={{
        frameRate: '30fps',
      }}
    />
  );
};
```

```

tsx

import {Composition, useCurrentFrame, useVideoConfig} from 'remotion';

const VariableFps: React.FC<{
  frameRate: '30fps' | '60fps';
}> = () => {
  const {fps} = useVideoConfig();
  return <div>{fps} FPS</div>;
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="VariableFps"
      component={VariableFps}
      width={1080}
      height={1080}
      durationInFrames={100}
      calculateMetadata={({props}) => {
        return {
          fps: props.frameRate === '60fps' ? 60 : 30,
        };
      }}
      defaultProps={{
        frameRate: '30fps',
      }}
    />
  );
};
```

## FPS Convert is discouraged [​](\#fps-convert-is-discouraged "Direct link to FPS Convert is discouraged")

Previously, this page showed a FPS converter snippet.

It's usage is not recommended because it does not work with media tags ( `<Video>`, `<Audio>`, `<OffthreadVideo>`, etc.).

## See also [​](\#see-also "Direct link to See also")

- [Dynamic FPS, duration and dimensions](/docs/dynamic-metadata)
- [`<Composition>`](/docs/composition)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/multiple-fps.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Build a timeline-based video editor](/docs/building-a-timeline) [Next\
\
Integration into Angular](/docs/angular)

- [Write animations frame-rate-independently](#write-animations-frame-rate-independently)
- [Change frame rate dynamically](#change-frame-rate-dynamically)
- [FPS Convert is discouraged](#fps-convert-is-discouraged)
- [See also](#see-also)
