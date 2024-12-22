---
title: <CameraMotionBlur>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <CameraMotionBlur>

- [Home page](/)
- [@remotion/motion-blur](/docs/motion-blur/)
- <CameraMotionBlur>

On this page

# <CameraMotionBlur>

_available from v3.2.39_

The `<CameraMotionBlur>` produces natural looking motion blur similar to what would be produced by
a film camera.

For this technique to work, the children must be absolutely positioned so many layers can be created without influencing the layout.

You can use the [`<AbsoluteFill>`](/docs/absolute-fill) component to absolutely position content.

note

The technique is destructive to colors. It is recommended to keep the `samples` property as low as
possible and carefully inspect that the output is of acceptable quality.

## API [​](\#api "Direct link to API")

Wrap your content in `<CameraMotionBlur>` and optionally add the following props in addition.

### `shutterAngle` [​](\#shutterangle "Direct link to shutterangle")

_optional - default: `180`_

Controls the amount of blur.

A lower value will result in less blur and a higher value will result in more.

The blur also depends on the frames per second (FPS). Higher FPS will naturally have less blur and
lower FPS will have more blur.

In movies and TV common values are (FPS/shutter angle):

- 24 fps / 180° or 90°
- 25 fps / 180° or 90°
- 30 fps / 180° or 90°
- 50 fps / 180° or 90°
- 60 fps / 180° or 90°

Details

What is "shutter angle"?
Many analog film cameras use rotating discs with partial cut-outs to block or let light through to
expose the analog film. Zero degrees is equal to completely blocking the light, and 360 degrees is
the same as not blocking any light at all.

The most common values used in the film industry are 90 and 180 degrees. These values are the same
as what you've experienced in most movies.

Read more here: [Rotary disc shutter on Wikipedia](https://en.wikipedia.org/wiki/Rotary_disc_shutter)

### `samples` [​](\#samples "Direct link to samples")

_optional - default: `10`_

The final image is an average of the samples. For a value of `10` the component will render ten
frames with different time offsets and combine them into a final image.

caution

A high number will produce a higher quality blur at the cost of image quality. See example below.

Recommended values: 5-10.

## Example usage [​](\#example-usage "Direct link to Example usage")

```

tsx

import {CameraMotionBlur} from '@remotion/motion-blur';
import {AbsoluteFill} from 'remotion';

export const MyComposition = () => {
  return (
    <CameraMotionBlur shutterAngle={180} samples={10}>
      <AbsoluteFill
        style={{
          backgroundColor: 'white',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <RainbowSquare />
      </AbsoluteFill>
    </CameraMotionBlur>
  );
};
```

```

tsx

import {CameraMotionBlur} from '@remotion/motion-blur';
import {AbsoluteFill} from 'remotion';

export const MyComposition = () => {
  return (
    <CameraMotionBlur shutterAngle={180} samples={10}>
      <AbsoluteFill
        style={{
          backgroundColor: 'white',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <RainbowSquare />
      </AbsoluteFill>
    </CameraMotionBlur>
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

# Animation

A

`shutterAngle={180}`

`samples={10}`

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/motion-blur/src/CameraMotionBlur.tsx)
- [Common mistake with `<Trail>` and `<CameraMotionBlur>`](/docs/motion-blur/common-mistake)

CONTRIBUTORS

[![UmungoBungo](https://github.com/UmungoBungo.png)\
\
**UmungoBungo** \
\
Idea](https://github.com/UmungoBungo) [![marcusstenbeck](https://github.com/marcusstenbeck.png)\
\
**marcusstenbeck** \
\
Implementation](https://github.com/marcusstenbeck)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/motion-blur/camera-motion-blur.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Trail>](/docs/motion-blur/trail) [Next\
\
Common mistake](/docs/motion-blur/common-mistake)

- [API](#api)
  - [`shutterAngle`](#shutterangle)
  - [`samples`](#samples)
- [Example usage](#example-usage)
- [Demo](#demo)
- [See also](#see-also)
