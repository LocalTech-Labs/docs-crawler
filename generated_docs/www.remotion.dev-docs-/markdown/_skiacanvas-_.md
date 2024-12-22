---
title: <SkiaCanvas />
source: Remotion Documentation
last_updated: 2024-12-22
---

# <SkiaCanvas />

- [Home page](/)
- [@remotion/skia](/docs/skia/)
- <SkiaCanvas />

On this page

# <SkiaCanvas />

A [React Native Skia `<Canvas />` component](https://shopify.github.io/react-native-skia/docs/canvas/overview) that wraps Remotion contexts.

You can place elements from `@shopify/react-native-skia` in it!

```

tsx

import { SkiaCanvas } from "@remotion/skia";
import { Fill } from "@shopify/react-native-skia";
import React from "react";
import { useVideoConfig } from "remotion";

const MySkiaVideo: React.FC = () => {
  const { width, height } = useVideoConfig();
  return (
    <SkiaCanvas width={width} height={height}>
      <Fill color="black" />
    </SkiaCanvas>
  );
};
```

```

tsx

import { SkiaCanvas } from "@remotion/skia";
import { Fill } from "@shopify/react-native-skia";
import React from "react";
import { useVideoConfig } from "remotion";

const MySkiaVideo: React.FC = () => {
  const { width, height } = useVideoConfig();
  return (
    <SkiaCanvas width={width} height={height}>
      <Fill color="black" />
    </SkiaCanvas>
  );
};
```

## Props [​](\#props "Direct link to Props")

### `width` [​](\#width "Direct link to width")

The width of the canvas in pixels.

### `height` [​](\#height "Direct link to height")

The height of the canvas in pixels.

### Inherited props [​](\#inherited-props "Direct link to Inherited props")

All the props that are accepted by [`<Canvas>`](https://shopify.github.io/react-native-skia/docs/canvas/overview) are accepted as well.

## See also [​](\#see-also "Direct link to See also")

- [Installation](/docs/skia)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/skia/skia-canvas.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
enableSkia()](/docs/skia/enable-skia) [Next\
\
Overview](/docs/lottie/)

- [Props](#props)
  - [`width`](#width)
  - [`height`](#height)
  - [Inherited props](#inherited-props)
- [See also](#see-also)
