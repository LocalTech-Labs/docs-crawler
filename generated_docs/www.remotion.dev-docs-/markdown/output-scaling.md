---
title: Output scaling
source: Remotion Documentation
last_updated: 2024-12-22
---

# Output scaling

- [Home page](/)
- [Rendering](/docs/render)
- Output scaling

On this page

# Output scaling

_Available from v2.6.7._

Output scaling is useful if you would like to render the video in multiple resolutions in the same aspect ratio.

**Example**: Your video canvas is in Full HD ( `1920x1080`), but would like to render your video to be in 4k ( `3840x2160` or `2x`).

Remotion can support this higher resolution by setting the [`deviceScaleFactor`](https://puppeteer.github.io/puppeteer/docs/puppeteer.viewport.devicescalefactor) of Puppeteer and upscale certain elements.

## How to scale [​](\#how-to-scale "Direct link to How to scale")

- In the CLI, during a render of a video or a still, pass the [`--scale`](/docs/cli/render#--scale) flag. For example: `--scale=2`

- In the Node.JS functions [`renderStill()`](/docs/renderer/render-still#scale), [`renderFrames()`](/docs/renderer/render-frames#scale), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda), you can pass a `scale` option.

- In the [config file](/docs/config), you can pass the scale using the following statement:



```

ts

Config.setScale(2);
```


```

ts

Config.setScale(2);
```

## Allowed values [​](\#allowed-values "Direct link to Allowed values")

The highest scale possible is `16` (sixteen times higher dimensions on each size or 256 times more pixels).

Positive values below 1 are allowed. For example, `0.5` will half each dimension.

The scale must result in a value that will result in integer pixels. A value of `1.00000001` for a composition with a width of `1920` pixels is not allowed.

For MP4 videos, the scale must resolve in a value where both dimensions are divisible by 2, since the codec does not support odd numbers.

If you would like to downscale a composition from `1920` to `1280` pixels, pass a scale of `2/3` to avoid rounding errors. This does not currently work as a CLI flag.

## Scalable elements [​](\#scalable-elements "Direct link to Scalable elements")

Elements that can be upscaled and that will enhance increased resolution are:

- Text elements
- SVG elements
- Images (if their resolution is sufficient to display in a higher resolution)

Elements that **cannot** be upscaled for increased resolution are:

- Videos
- Canvas and WebGL elements

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/scaling.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Still images](/docs/stills) [Next\
\
Transparent videos](/docs/transparent-videos)

- [How to scale](#how-to-scale)
- [Allowed values](#allowed-values)
- [Scalable elements](#scalable-elements)
