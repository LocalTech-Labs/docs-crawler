---
title: remotion
source: Remotion Documentation
last_updated: 2024-12-22
---

# remotion

- [Home page](/)
- remotion

On this page

# remotion

A package containing the essential building blocks of expressing videos in Remotion.

Some pure functions such as [`interpolate()`](/docs/interpolate) and [`interpolateColors()`](/docs/interpolate-colors) can also be used outside of Remotion.

## Installation [​](\#installation "Direct link to Installation")

- npm
- bun
- pnpm
- yarn

```

bash

npm i remotion
```

```

bash

npm i remotion
```

```

bash

bun i remotion
```

```

bash

bun i remotion
```

```

bash

pnpm i remotion
```

```

bash

pnpm i remotion
```

```

bash

yarn add remotion
```

```

bash

yarn add remotion
```

## API [​](\#api "Direct link to API")

The following functions and components are exposed:

[**<Composition>** \
\
Define a video](/docs/composition) [**<Still>** \
\
Define a still](/docs/still) [**<Folder>** \
\
Organize compositions in the Studio sidebar](/docs/folder) [**registerRoot()** \
\
Initialize a Remotion project](/docs/register-root) [**useCurrentFrame()** \
\
Obtain the current time](/docs/use-current-frame) [**useVideoConfig()** \
\
Get the duration, dimensions and FPS of a composition](/docs/use-video-config) [**interpolate()** \
\
Map a range of values to another](/docs/interpolate) [**spring()** \
\
Physics-based animation primitive](/docs/spring) [**interpolateColors()** \
\
Map a range of values to colors](/docs/interpolate-colors) [**measureSpring()** \
\
Determine the duration of a spring](/docs/measure-spring) [**Easing** \
\
Customize animation curve of `interpolate()`](/docs/easing) [**<Img>** \
\
Render an `<img>` tag and wait for it to load](/docs/img) [**<Video>** \
\
Synchronize a `<video>` with Remotion's time](/docs/video) [**<Audio>** \
\
Synchronize `<audio>` with Remotion's time](/docs/audio) [**<OffthreadVideo>** \
\
Alternative to `<Video>`](/docs/offthreadvideo) [**<IFrame>** \
\
Render an `<iframe>` tag and wait for it to load](/docs/iframe) [**<Sequence>** \
\
Time-shifts it's children](/docs/sequence) [**<Series>** \
\
Display contents after another](/docs/series) [**<Freeze>** \
\
Freeze some content in time](/docs/freeze) [**<Loop>** \
\
Play some content repeatedly](/docs/loop) [**delayRender()** \
\
Block a render from continuing](/docs/delay-render) [**continueRender()** \
\
Unblock a render](/docs/continue-render) [**cancelRender()** \
\
Abort an error](/docs/cancel-render) [**getInputProps()** \
\
Receive the user-defined input data](/docs/get-input-props) [**getRemotionEnvironment()** \
\
Determine if you are currently previewing or rendering](/docs/get-remotion-environment) [**staticFile()** \
\
Access file from `public/` folder](/docs/staticfile) [**<AbsoluteFill>** \
\
Position content absolutely and in full size](/docs/absolute-fill) [**VERSION** \
\
Get the current version of Remotion](/docs/version)

## License [​](\#license "Direct link to License")

[Remotion License](https://remotion.dev/license)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/remotion.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Configuration file](/docs/config) [Next\
\
<AbsoluteFill>](/docs/absolute-fill)

- [Installation](#installation)
- [API](#api)
- [License](#license)
