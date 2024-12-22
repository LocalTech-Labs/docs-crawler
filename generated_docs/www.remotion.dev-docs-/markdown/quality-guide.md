---
title: Quality Guide
source: Remotion Documentation
last_updated: 2024-12-22
---

# Quality Guide

- [Home page](/)
- [Rendering](/docs/render)
- Quality Guide

On this page

# Quality Guide

Video encoding involves heavily compressing the video data to reduce the file size, which is a lossy process.

This document explains a few factors and settings that can be used to control the quality of the video.

## CRF [​](\#crf "Direct link to CRF")

The main setting to control the quality of the video is the [Constant Rate Factor (CRF)](/docs/encoding/#controlling-quality-using-the-crf-setting).

All ways to render a video in Remotion allow passing a CRF value. Allowed values depend on the codec.

## Resolution (Unsharp text) [​](\#resolution-unsharp-text "Direct link to Resolution (Unsharp text)")

A somewhat surprising way to lose quality is the fact that many computers have a high-density display, but the dimensions of a video don't consider that.

Consider a 1920x1080 video, in which you render text:

```

tsx

import { Composition } from "remotion";

const MyComponent = () => {
  return <div>Hello World</div>;
};

const Root: React.FC = () => {
  return (
    <Composition
      width={1920}
      height={1080}
      fps={30}
      durationInFrames={100}
      component={MyComponent}
      id="MyComp"
    />
  );
};
```

```

tsx

import { Composition } from "remotion";

const MyComponent = () => {
  return <div>Hello World</div>;
};

const Root: React.FC = () => {
  return (
    <Composition
      width={1920}
      height={1080}
      fps={30}
      durationInFrames={100}
      component={MyComponent}
      id="MyComp"
    />
  );
};
```

If you render it out and embed the video on a webpage:

```

html

<video src="video.mp4" width="1920" height="1080"></video>
```

```

html

<video src="video.mp4" width="1920" height="1080"></video>
```

You will have a less sharp text if the video is displayed on a high-density display (like a Macbook).

This is because the device has a 2x pixel density, but the video is 1920x1080, which means that each pixel is 2x bigger than the device's pixel.

To display the font with the same sharpness as on a low-density display, you would need to render a 4K video.

Use [Output Scaling](/docs/scaling) to render the video with a higher pixel density.

## JPEG Quality [​](\#jpeg-quality "Direct link to JPEG Quality")

Remotion makes screenshots of the browser page by default in JPEG format.

The JPEG quality by default is 80 (on a scale of 0 to 100).

You may tweak it using the [`--jpeg-quality`](/docs/cli/render#--jpeg-quality) flag or switch to PNG as the image format using [`--image-format=png`](/docs/cli/render#--image-format) (slower).

## Color accuracy [​](\#color-accuracy "Direct link to Color accuracy")

Export your video in [`bt709`](/docs/cli/render#--color-space) for more accurate colors.

note

This will become the default in Remotion 5.0.

## GIF Colors [​](\#gif-colors "Direct link to GIF Colors")

Ensure you are on at least Remotion version v4.0.138 to get the best color accuracy.

GIFs only support a 256 colors palette, which leads to a big quality loss. Embrace it, or choose a different format if it is not acceptable.

## Video bitrate [​](\#video-bitrate "Direct link to Video bitrate")

You can use the [`--video-bitrate`](/docs/cli/render#--video-bitrate) flag to control the bitrate of the video.

It is mutually exclusive with [`--crf`](/docs/cli/render#--crf) option, which means you can not use both at the same time.

## x264 Preset [​](\#x264-preset "Direct link to x264 Preset")

When exporting a H.264 video, you can use the [`--x264-preset`](/docs/cli/render#--x264-preset) flag to control the quality/speed/compression tradeoff.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/quality.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Creating overlays](/docs/overlay) [Next\
\
Render a dataset](/docs/dataset-render)

- [CRF](#crf)
- [Resolution (Unsharp text)](#resolution-unsharp-text)
- [JPEG Quality](#jpeg-quality)
- [Color accuracy](#color-accuracy)
- [GIF Colors](#gif-colors)
- [Video bitrate](#video-bitrate)
- [x264 Preset](#x264-preset)
