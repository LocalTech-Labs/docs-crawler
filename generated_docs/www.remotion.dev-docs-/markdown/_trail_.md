---
title: <Trail>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Trail>

- [Home page](/)
- [@remotion/motion-blur](/docs/motion-blur/)
- <Trail>

On this page

# <Trail>

_available from v3.2.39, previously called `<MotionBlur>`_

The `<Trail>` component duplicates it's children and adds a time offset to each layer in order to create a trail effect.

For this technique to work, the children must be absolutely positioned so many layers can be created without influencing the layout.

You can use the [`<AbsoluteFill>`](/docs/absolute-fill) component to absolutely position content.

## API [​](\#api "Direct link to API")

Wrap your content in `<Trail>` and add the following props in addition.

### `layers` [​](\#layers "Direct link to layers")

How many layers are added below the content. Must be an integer

### `lagInFrames` [​](\#laginframes "Direct link to laginframes")

How many frames each layer is lagging behind the last one. Can also a floating point number.

### `trailOpacity` [​](\#trailopacity "Direct link to trailopacity")

_previously called `blurOpacity`_

The highest opacity of a layer. The lowest opacity is 0 and layers intbetween get interpolated.

## Example usage [​](\#example-usage "Direct link to Example usage")

```

tsx

import {Trail} from '@remotion/motion-blur';
import {AbsoluteFill} from 'remotion';

export const MyComposition = () => {
  return (
    <Trail layers={50} lagInFrames={0.1} trailOpacity={1}>
      <AbsoluteFill
        style={{
          backgroundColor: 'white',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <BlueSquare />
      </AbsoluteFill>
    </Trail>
  );
};
```

```

tsx

import {Trail} from '@remotion/motion-blur';
import {AbsoluteFill} from 'remotion';

export const MyComposition = () => {
  return (
    <Trail layers={50} lagInFrames={0.1} trailOpacity={1}>
      <AbsoluteFill
        style={{
          backgroundColor: 'white',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <BlueSquare />
      </AbsoluteFill>
    </Trail>
  );
};
```

## Demo [​](\#demo "Direct link to Demo")

# Still

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

# Animation

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

A

`layers={50}`

`trailOpacity={1}`

`lagInFrames={0.3}`

CONTRIBUTORS

[![UmungoBungo](https://github.com/UmungoBungo.png)\
\
**UmungoBungo** \
\
Implementation](https://github.com/UmungoBungo)

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/motion-blur/src/Trail.tsx)
- [Common mistake with `<Trail>` and `<CameraMotionBlur>`](/docs/motion-blur/common-mistake)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/motion-blur/trail.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/motion-blur](/docs/motion-blur/) [Next\
\
<CameraMotionBlur>](/docs/motion-blur/camera-motion-blur)

- [API](#api)
  - [`layers`](#layers)
  - [`lagInFrames`](#laginframes)
  - [`trailOpacity`](#trailopacity)
- [Example usage](#example-usage)
- [Demo](#demo)
- [See also](#see-also)
