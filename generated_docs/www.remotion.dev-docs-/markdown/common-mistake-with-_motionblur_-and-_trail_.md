---
title: Common mistake with <MotionBlur> and <Trail>
source: Remotion Documentation
last_updated: 2024-12-22
---

# Common mistake with <MotionBlur> and <Trail>

- [Home page](/)
- [@remotion/motion-blur](/docs/motion-blur/)
- Common mistake

On this page

# Common mistake with <MotionBlur> and <Trail>

The [`<Trail>`](/docs/motion-blur/trail) and [`<CameraMotionBlur>`](/docs/motion-blur/camera-motion-blur) components manipulate the React context that holds the current time.

This means that the motion blur effect doesn't work if you use the [`useCurrentFrame()`](/docs/use-current-frame) hook outside of a motion blur component.

```

❌ Wrong - useCurrentFrame() outside of CameraMotionBlur
tsx

import {AbsoluteFill, useCurrentFrame} from 'remotion';
import {CameraMotionBlur} from '@remotion/motion-blur';

export const MyComp = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill>
      <CameraMotionBlur>
        <AbsoluteFill
          style={{
            backgroundColor: 'red',
            justifyContent: 'center',
            alignItems: 'center',
            color: 'white',
            fontSize: frame,
          }}
        >
          A
        </AbsoluteFill>
      </CameraMotionBlur>
    </AbsoluteFill>
  );
};
```

```

❌ Wrong - useCurrentFrame() outside of CameraMotionBlur
tsx

import {AbsoluteFill, useCurrentFrame} from 'remotion';
import {CameraMotionBlur} from '@remotion/motion-blur';

export const MyComp = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill>
      <CameraMotionBlur>
        <AbsoluteFill
          style={{
            backgroundColor: 'red',
            justifyContent: 'center',
            alignItems: 'center',
            color: 'white',
            fontSize: frame,
          }}
        >
          A
        </AbsoluteFill>
      </CameraMotionBlur>
    </AbsoluteFill>
  );
};
```

You can fix this by extracting the animation into a separate component:

```

✅ Correct - useCurrentFrame() inside the child component
tsx

import {AbsoluteFill, useCurrentFrame} from 'remotion';
import {CameraMotionBlur} from '@remotion/motion-blur';

const A: React.FC = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill
      style={{
        backgroundColor: 'red',
        justifyContent: 'center',
        alignItems: 'center',
        color: 'white',
        fontSize: frame,
      }}
    >
      A
    </AbsoluteFill>
  );
};

export const MyComp = () => {
  return (
    <AbsoluteFill>
      <CameraMotionBlur>
        <A />
      </CameraMotionBlur>
    </AbsoluteFill>
  );
};
```

```

✅ Correct - useCurrentFrame() inside the child component
tsx

import {AbsoluteFill, useCurrentFrame} from 'remotion';
import {CameraMotionBlur} from '@remotion/motion-blur';

const A: React.FC = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill
      style={{
        backgroundColor: 'red',
        justifyContent: 'center',
        alignItems: 'center',
        color: 'white',
        fontSize: frame,
      }}
    >
      A
    </AbsoluteFill>
  );
};

export const MyComp = () => {
  return (
    <AbsoluteFill>
      <CameraMotionBlur>
        <A />
      </CameraMotionBlur>
    </AbsoluteFill>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [`<Trail>`](/docs/motion-blur/trail)
- [`<CameraMotionBlur>`](/docs/motion-blur/camera-motion-blur)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/motion-blur/common-mistake.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<CameraMotionBlur>](/docs/motion-blur/camera-motion-blur) [Next\
\
@remotion/lambda](/docs/lambda/api)

- [See also](#see-also)
