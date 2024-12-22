---
title: Flickering when using Next.js <Image> tag
source: Remotion Documentation
last_updated: 2024-12-22
---

# Flickering when using Next.js <Image> tag

- [Home page](/)
- Troubleshooting
- Flickering with Next.js Image

On this page

# Flickering when using Next.js <Image> tag

The following code is discouraged in Remotion:

```

❌ Don't do this
tsx

import Image from "next/image";

const myMarkup = <Image src="https://picsum.photos/200/300"></Image>;
```

```

❌ Don't do this
tsx

import Image from "next/image";

const myMarkup = <Image src="https://picsum.photos/200/300"></Image>;
```

## Problem [​](\#problem "Direct link to Problem")

Remotion has no way of knowing when the image is finished loading because it is impossible to determine so when loading an image through `<Image>`.

This will lead to Remotion not waiting for the image to be loaded during rendering and cause visible flickers.

## Solution [​](\#solution "Direct link to Solution")

Use the [`<Img>`](/docs/img) tag instead:

```

✅ Do this
tsx

import { AbsoluteFill, Img } from "remotion";

const myMarkup = <Img src="https://picsum.photos/200/300" />;
```

```

✅ Do this
tsx

import { AbsoluteFill, Img } from "remotion";

const myMarkup = <Img src="https://picsum.photos/200/300" />;
```

## See also [​](\#see-also "Direct link to See also")

- [`<Img>`](/docs/img)
- [Flickering](/docs/flickering)
- [Flickering with background-image](/docs/troubleshooting/background-image)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/troubleshooting/nextjs-image.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Flickering with background-image](/docs/troubleshooting/background-image) [Next\
\
Apple Silicon detected](/docs/troubleshooting/rosetta)

- [Problem](#problem)
- [Solution](#solution)
- [See also](#see-also)
