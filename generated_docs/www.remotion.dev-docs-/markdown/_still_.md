---
title: <Still>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Still>

- [Home page](/)
- [remotion](/docs/remotion)
- <Still>

On this page

# <Still>

A `<Still />` is a single-frame [`<Composition />`](/docs/composition). It is a convenient way to define a composition that renders an image rather than a video.

## Example [​](\#example "Direct link to Example")

The `<Still />` component has the same API as the [`<Composition />`](/docs/composition) component, except that it's not necessary to pass `durationInFrames` and `fps`.

```

tsx

import { Composition, Still } from "remotion";
import { MyComp } from "./MyComp";

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-video"
        component={MyComp}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
      <Still id="my-image" component={MyComp} width={1080} height={1080} />
    </>
  );
};
```

```

tsx

import { Composition, Still } from "remotion";
import { MyComp } from "./MyComp";

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-video"
        component={MyComp}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
      <Still id="my-image" component={MyComp} width={1080} height={1080} />
    </>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Still.tsx)
- [`<Composition />`](/docs/composition)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/still.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
staticFile()](/docs/staticfile) [Next\
\
useBufferState()](/docs/use-buffer-state)

- [Example](#example)
- [See also](#see-also)
