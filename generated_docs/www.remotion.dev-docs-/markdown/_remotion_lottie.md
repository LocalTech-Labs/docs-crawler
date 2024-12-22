---
title: @remotion/lottie
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/lottie

- [Home page](/)
- @remotion/lottie

On this page

# @remotion/lottie

This package provides a component for displaying [Lottie](http://airbnb.io/lottie/) animations in Remotion.

## Installation [​](\#installation "Direct link to Installation")

Install, `@remotion/lottie` as well as `lottie-web`.

- npm
- yarn
- pnpm

```

bash

npm i @remotion/lottie lottie-web
```

```

bash

npm i @remotion/lottie lottie-web
```

```

bash

yarn add @remotion/lottie lottie-web
```

```

bash

yarn add @remotion/lottie lottie-web
```

```

bash

pnpm i @remotion/lottie lottie-web
```

```

bash

pnpm i @remotion/lottie lottie-web
```

You can now start using the [`<Lottie>`](/docs/lottie/lottie) component in your Remotion project.

## Supported features [​](\#supported-features "Direct link to Supported features")

- Playing Lottie animations using `lottie-web`
- Change the speed of the animation
- Playing animation forwards and backwards
- Playing remote files
- Determining dimensions and duration of a Lottie animation

## Unsupported features [​](\#unsupported-features "Direct link to Unsupported features")

- Rendering on other renderers as `svg`
- `setSubFrame()`, `setLocationHref()`
- Limited expression support: Remotion uses the `.goToAndStop()` method from `lottie-web` to seek through the Lottie file. Depending on the expression, the frame might not render deterministally, leading to [flickering](/docs/flickering) in the Remotion output. Remotion cannot fix this without a change in `lottie-web` upstream. You need to evaluate on a case-by-case basis whether the expression you are using is supported by Remotion.

note

[Open an issue](https://remotion.dev/issue) if you want to request a currently unsupported feature.

## Table of contents [​](\#table-of-contents "Direct link to Table of contents")

[**<Lottie>** \
\
Embed a Lottie animation in Remotion](/docs/lottie/lottie) [**getLottieMetadata()** \
\
Get metadata of a Lottie animation](/docs/lottie/getlottiemetadata) [**staticFile()** \
\
Load Lottie animations from a static file](/docs/lottie/staticfile) [Loading Lottie animations from a remote URL](/docs/lottie/remote) [Where to find Lottie files](/docs/lottie/lottiefiles)

## License [​](\#license "Direct link to License")

[Remotion License](https://remotion.dev/license)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lottie/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<SkiaCanvas />](/docs/skia/skia-canvas) [Next\
\
<Lottie>](/docs/lottie/lottie)

- [Installation](#installation)
- [Supported features](#supported-features)
- [Unsupported features](#unsupported-features)
- [Table of contents](#table-of-contents)
- [License](#license)
