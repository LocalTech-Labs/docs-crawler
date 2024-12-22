---
title: <OffthreadVideo /> while rendering
source: Remotion Documentation
last_updated: 2024-12-22
---

# <OffthreadVideo /> while rendering

- [Home page](/)
- Snippets
- <OffthreadVideo /> while rendering

# <OffthreadVideo /> while rendering

The following component will only use [`<OffthreadVideo />`](/docs/offthreadvideo) while rendering, but [`<Video />`](/docs/video) in the Player.
This is useful for attaching a `ref` to the [`<Video />`](/docs/video) tag.

```

tsx

import { forwardRef } from "react";
import {
  getRemotionEnvironment,
  OffthreadVideo,
  RemotionOffthreadVideoProps,
  Video,
} from "remotion";

const OffthreadWhileRenderingRefForwardingFunction: React.ForwardRefRenderFunction<
  HTMLVideoElement,
  RemotionOffthreadVideoProps
> = (props, ref) => {
  const { imageFormat, ...otherProps } = props;
  const isPreview = !getRemotionEnvironment().isRendering;

  if (isPreview) {
    return <Video ref={ref} {...otherProps}></Video>;
  }

  return <OffthreadVideo {...props}></OffthreadVideo>;
};

export const OffthreadVideoWhileRendering = forwardRef(
  OffthreadWhileRenderingRefForwardingFunction
);
```

```

tsx

import { forwardRef } from "react";
import {
  getRemotionEnvironment,
  OffthreadVideo,
  RemotionOffthreadVideoProps,
  Video,
} from "remotion";

const OffthreadWhileRenderingRefForwardingFunction: React.ForwardRefRenderFunction<
  HTMLVideoElement,
  RemotionOffthreadVideoProps
> = (props, ref) => {
  const { imageFormat, ...otherProps } = props;
  const isPreview = !getRemotionEnvironment().isRendering;

  if (isPreview) {
    return <Video ref={ref} {...otherProps}></Video>;
  }

  return <OffthreadVideo {...props}></OffthreadVideo>;
};

export const OffthreadVideoWhileRendering = forwardRef(
  OffthreadWhileRenderingRefForwardingFunction
);
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/offthread-video-while-rendering.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useDelayRender()](/docs/miscellaneous/snippets/use-delay-render) [Next\
\
Combining compositions](/docs/miscellaneous/snippets/combine-compositions)
