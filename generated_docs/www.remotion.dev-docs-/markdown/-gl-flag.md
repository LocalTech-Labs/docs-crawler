---
title: --gl flag
source: Remotion Documentation
last_updated: 2024-12-22
---

# --gl flag

- [Home page](/)
- Miscellaneous
- --gl flag

On this page

# --gl flag

When rendering a video in Remotion, different [GL](https://en.wikipedia.org/wiki/OpenGL) renderer backends can be selected.

The following renderer backends are supported in Remotion:

- `null` \- default, lets Chrome decide
- `angle`
- `egl`
- `swiftshader`
- `vulkan` (from Remotion v4.0.41)
- `angle-egl` (from Remotion v4.0.52)
- `swangle` \- default on Lambda

## Recommended renderers [​](\#recommended-renderers "Direct link to Recommended renderers")

[1](#1)

If you use WebGL/Three.js:

- On a desktop, `angle` is recommended
- On a [cloud instance with a GPU](/docs/miscellaneous/cloud-gpu), `angle-egl` is recommended
- On Lambda, use `swangle` (default on Lambda)
- On a machine with no GPU, `swangle` is recommended. Rendering might be slow.

[2](#2)

If you don't use WebGL/Three.js, the default renderer ( `null` on local, `swangle` on Lambda) are the best choice.

## Using the GPU [​](\#using-the-gpu "Direct link to Using the GPU")

In cases where a GPU could be beneficial for rendering, it can often make sense to use the `angle` renderer ( `angle-egl` on Linux). An in-depth explanation when and how to use it is given in [this article](/docs/gpu).

⚠️ Memory leaks are a known problem with `angle`. We recommend to split up long renders into multiple parts when rendering large videos, since sometimes renders can fail due to memory leaks.

Currently, GitHub Actions will fail when using `angle`, since Actions runners don't have a GPU.

## Selecting the renderer backend [​](\#selecting-the-renderer-backend "Direct link to Selecting the renderer backend")

The renderer backend can be set in different ways:

### Via Node.JS APIs [​](\#via-nodejs-apis "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions#chromiumoptions), [`renderStill()`](/docs/renderer/render-still#gl), [`renderMedia()`](/docs/renderer/render-media#gl), [`renderFrames()`](/docs/renderer/render-frames#gl), [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda#gl), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#gl) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#gl), you can pass [`chromiumOptions.gl`](/docs/renderer/render-still#gl).

### Via Config file [​](\#via-config-file "Direct link to Via Config file")

```

tsx

Config.setChromiumOpenGlRenderer("angle");
```

```

tsx

Config.setChromiumOpenGlRenderer("angle");
```

note

The config file only applies to CLI commands.

note

Prior to `v3.3.39`, the option was called `Config.Puppeteer.setChromiumOpenGlRenderer()`.

### Via CLI flag [​](\#via-cli-flag "Direct link to Via CLI flag")

Pass [`--gl=[angle,swangle,...]`](/docs/cli) in one of the following commands: [`remotion render`](/docs/cli/render), [`remotion compositions`](/docs/cli/compositions), [`remotion still`](/docs/cli/still), [`remotion lambda render`](/docs/lambda/cli/render), [`remotion lambda still`](/docs/lambda/cli/still), [`remotion lambda compositions`](/docs/lambda/cli/compositions).

## See also [​](\#see-also "Direct link to See also")

- [Using the GPU](/docs/gpu)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/open-gl.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Linux Dependencies](/docs/miscellaneous/linux-dependencies) [Next\
\
Bun support](/docs/bun)

- [Recommended renderers](#recommended-renderers)
- [Using the GPU](#using-the-gpu)
- [Selecting the renderer backend](#selecting-the-renderer-backend)
  - [Via Node.JS APIs](#via-nodejs-apis)
  - [Via Config file](#via-config-file)
  - [Via CLI flag](#via-cli-flag)
- [See also](#see-also)
