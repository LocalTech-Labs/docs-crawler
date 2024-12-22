---
title: Transparent videos
source: Remotion Documentation
last_updated: 2024-12-22
---

# Transparent videos

- [Home page](/)
- [Rendering](/docs/render)
- Transparent videos

On this page

# Transparent videos

Chrome and Firefox support WebM videos with alpha channels. That means that on these browsers, you can embed videos with transparency.

If you are lucky enough to visit the website in one of those supported web browsers, check it out:

Toggle transparency

## Creating a WebM video with Alpha channel [​](\#creating-a-webm-video-with-alpha-channel "Direct link to Creating a WebM video with Alpha channel")

In order to create a transparent video, you need at least version 1.4 of Remotion. Make sure to not set any background (use the checkerboard button to ensure your video is transparent). In order to render your videos transparent, you need to change 3 settings:

- Render each frame as PNG.
- Use the VP8 or VP9 codec
- Use the `yuva420p` pixel format.

If you want to set these options and persist them, add this to your `remotion.config.ts` file (create it if you don't yet have one);

```

tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva420p");
Config.setCodec("vp8");
```

```

tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva420p");
Config.setCodec("vp8");
```

You can also set the settings on the command line:

```

bash

--image-format=png --pixel-format=yuva420p --codec=vp8
```

```

bash

--image-format=png --pixel-format=yuva420p --codec=vp8
```

## Creating a ProRes video with Alpha channels [​](\#creating-a-prores-video-with-alpha-channels "Direct link to Creating a ProRes video with Alpha channels")

If you want to export a transparent video for use in another video editing program, Apple ProRes is a more suitable option.
ProRes is supported by Final Cut Pro, Adobe Premiere and Davinci Resolve.

Supported since v2.1.7, you can set the codec to `prores` and choose a ProRes profile with alpha support: Either: `4444` or `hq`. The pixel format must be `yuva444p10le`.

```

tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva444p10le");
Config.setCodec("prores");
Config.setProResProfile("4444");
```

```

tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva444p10le");
Config.setCodec("prores");
Config.setProResProfile("4444");
```

You can also set the settings on the command line:

```

bash

--image-format=png --pixel-format=yuva444p10le --codec=prores --prores-profile=4444
```

```

bash

--image-format=png --pixel-format=yuva444p10le --codec=prores --prores-profile=4444
```

## Enabling transparency when using `<OffthreadVideo>` [​](\#enabling-transparency-when-using-offthreadvideo "Direct link to enabling-transparency-when-using-offthreadvideo")

By default, the [`<OffthreadVideo>`](/docs/offthreadvideo) does not add an alpha layer during export.

Add the [`transparent`](/docs/offthreadvideo#transparent) prop to flag that the video has an alpha channel.

This slightly reduces the performance of the render.

## Creating a fallback version [​](\#creating-a-fallback-version "Direct link to Creating a fallback version")

Given the poor browser support, consider making two versions of a video, one with alpha channel, and an opaque video as a fallback. You can achieve this in Remotion using standard React props:

```

tsx

const MyVideo: React.FC<{
  transparent: boolean;
}> = ({transparent}) => {
  return (
    <div style={{backgroundColor: transparent ? undefined : 'white', flex: 1}}>
      {/* Omit opaque background based on `transparent` prop */}
    </div>
  )
}
```

```

tsx

const MyVideo: React.FC<{
  transparent: boolean;
}> = ({transparent}) => {
  return (
    <div style={{backgroundColor: transparent ? undefined : 'white', flex: 1}}>
      {/* Omit opaque background based on `transparent` prop */}
    </div>
  )
}
```

It's a good practice to set a default when defining the composition:

```

tsx

<Composition
  id="my-video"
  component={MyVideo}
  width={1920}
  height={1080}
  fps={30}
  durationInFrames={150}
  defaultProps={{
    transparent: true,
  }}
/>;
```

```

tsx

<Composition
  id="my-video"
  component={MyVideo}
  width={1920}
  height={1080}
  fps={30}
  durationInFrames={150}
  defaultProps={{
    transparent: true,
  }}
/>;
```

You can then have separate render commands in your package.json:

```

json

"scripts": {
  "render": "remotion render my-video video.mp4",
  "render-transparent": "remotion render --image-format=png --pixel-format=yuva420p --codec=vp8 my-video video-transparent.webm"
}
```

```

json

"scripts": {
  "render": "remotion render my-video video.mp4",
  "render-transparent": "remotion render --image-format=png --pixel-format=yuva420p --codec=vp8 my-video video-transparent.webm"
}
```

Now you can render two variants of the same video and serve a different one
based on browser support. You can use APIs like [`<source>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source) or the [`canplay`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/canplay_event) event to determine programmatically if a browser is able to play a video.

## See also [​](\#see-also "Direct link to See also")

- [CLI options](/docs/cli)
- [Configuration option](/docs/config)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transparent-videos.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Output scaling](/docs/scaling) [Next\
\
Rendering GIFs](/docs/render-as-gif)

- [Creating a WebM video with Alpha channel](#creating-a-webm-video-with-alpha-channel)
- [Creating a ProRes video with Alpha channels](#creating-a-prores-video-with-alpha-channels)
- [Enabling transparency when using `<OffthreadVideo>`](#enabling-transparency-when-using-offthreadvideo)
- [Creating a fallback version](#creating-a-fallback-version)
- [See also](#see-also)
