---
title: Flickering when using background-image or mask-image
source: Remotion Documentation
last_updated: 2024-12-22
---

# Flickering when using background-image or mask-image

- [Home page](/)
- Troubleshooting
- Flickering with background-image

On this page

# Flickering when using background-image or mask-image

The following code is discouraged in Remotion:

```

❌ Don't do this
tsx

const myMarkup = (
  <div
    style={{
      backgroundImage: `url(${src})`,
    }}
  >
    <p>Hello World</p>
  </div>
);
```

```

❌ Don't do this
tsx

const myMarkup = (
  <div
    style={{
      backgroundImage: `url(${src})`,
    }}
  >
    <p>Hello World</p>
  </div>
);
```

## Problem [​](\#problem "Direct link to Problem")

Remotion has no way of knowing when the image is finished loading because it is impossible to determine so when loading an image through `background-image`, `mask-image`, or other CSS properties. This will lead to Remotion not waiting for the image to be loaded during rendering and cause visible flickers.

## Solution [​](\#solution "Direct link to Solution")

Use the [`<Img>`](/docs/img) tag instead and place it in an [`<AbsoluteFill>`](/docs/absolute-fill):

```

✅ Do this
tsx

import { AbsoluteFill, Img } from "remotion";

const myMarkup = (
  <AbsoluteFill>
    <AbsoluteFill>
      <Img
        style={{
          width: "100%",
        }}
        src={src}
      />
    </AbsoluteFill>
    <AbsoluteFill>
      <p>Hello World</p>
    </AbsoluteFill>
  </AbsoluteFill>
);
```

```

✅ Do this
tsx

import { AbsoluteFill, Img } from "remotion";

const myMarkup = (
  <AbsoluteFill>
    <AbsoluteFill>
      <Img
        style={{
          width: "100%",
        }}
        src={src}
      />
    </AbsoluteFill>
    <AbsoluteFill>
      <p>Hello World</p>
    </AbsoluteFill>
  </AbsoluteFill>
);
```

The next will be placed on top of the image and will not flicker.

## Workaround [​](\#workaround "Direct link to Workaround")

If you cannot use an [`<Img>`](/docs/img) tag, for example because you need to use `mask-image()`, render an adjacent `<Img>` tag that renders the same src and place it outside the viewport:

```

✅ Do this
tsx

import { Img } from "remotion";

const myMarkup = (
  <>
    <Img
      src={src}
      style={{
        opacity: 0,
        position: "absolute",
        left: "-100%",
      }}
    />
    <div
      style={{
        maskImage: `url(${src})`,
      }}
    >
      <p>Hello World</p>
    </div>
  </>
);
```

```

✅ Do this
tsx

import { Img } from "remotion";

const myMarkup = (
  <>
    <Img
      src={src}
      style={{
        opacity: 0,
        position: "absolute",
        left: "-100%",
      }}
    />
    <div
      style={{
        maskImage: `url(${src})`,
      }}
    >
      <p>Hello World</p>
    </div>
  </>
);
```

The hidden `<Img>` tag ensures the image will download completely which allows the `background-image` will fully render.

## See also [​](\#see-also "Direct link to See also")

- [`<Img>`](/docs/img)
- [Flickering](/docs/flickering)
- [Flickering with Next.js' Image tag](/docs/troubleshooting/nextjs-image)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/background-image.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
staticFile() remote URLs](/docs/staticfile-remote-urls) [Next\
\
Flickering with Next.js Image](/docs/troubleshooting/nextjs-image)

- [Problem](#problem)
- [Solution](#solution)
- [Workaround](#workaround)
- [See also](#see-also)
