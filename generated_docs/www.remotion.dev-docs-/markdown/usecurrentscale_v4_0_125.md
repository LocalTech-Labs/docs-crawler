---
title: useCurrentScale()v4.0.125
source: Remotion Documentation
last_updated: 2024-12-22
---

# useCurrentScale()v4.0.125

- [Home page](/)
- [remotion](/docs/remotion)
- useCurrentScale()

On this page

# useCurrentScale() [v4.0.125](https://github.com/remotion-dev/remotion/releases/v4.0.125)

With this hook, you can retrieve the scale factor of the canvas.

Useful for if you want to [measure DOM nodes](/docs/measuring).

In the [Studio](/docs/terminology/studio), it will correspond to the zoom level - the value is `1` if the zoom is at 100%.

In the [Player](/docs/terminology/player), it will correspond to the scale that is needed to fit the video canvas in the Player.

```

MyComp.tsx
tsx

import { Sequence, useCurrentScale } from "remotion";

const MyVideo = () => {
  const scale = useCurrentScale();

  return <div>The current scale is {scale}</div>;
};
```

```

MyComp.tsx
tsx

import { Sequence, useCurrentScale } from "remotion";

const MyVideo = () => {
  const scale = useCurrentScale();

  return <div>The current scale is {scale}</div>;
};
```

If you are outside of a Remotion context, the hook will throw an error.

If you want to avoid this error and return a default scale, you can pass an options object with the property `dontThrowIfOutsideOfRemotion` set to `true`.

In this case, the hook will return `1`.

```

MyComp.tsx
tsx

import { useCurrentScale } from "remotion";

const MyRegularReactComponent = () => {
  const scale = useCurrentScale({ dontThrowIfOutsideOfRemotion: true });

  return <div>The current scale is {scale}</div>;
};
```

```

MyComp.tsx
tsx

import { useCurrentScale } from "remotion";

const MyRegularReactComponent = () => {
  const scale = useCurrentScale({ dontThrowIfOutsideOfRemotion: true });

  return <div>The current scale is {scale}</div>;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/use-current-scale.ts)
- [Measuring DOM nodes](/docs/measuring)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-current-scale.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useCurrentFrame()](/docs/use-current-frame) [Next\
\
useVideoConfig()](/docs/use-video-config)

- [See also](#see-also)
