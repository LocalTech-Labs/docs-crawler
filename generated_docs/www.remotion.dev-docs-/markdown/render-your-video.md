---
title: Render your video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Render your video

- [Home page](/)
- Rendering

On this page

# Render your video

There are various ways to render your video.

## Remotion Studio [​](\#remotion-studio "Direct link to Remotion Studio")

To render a video, click the " Render" button.

Choose your preferred settings, then confirm using the `Render video` button.

### Remotion Studio deployment [​](\#remotion-studio-deployment "Direct link to Remotion Studio deployment")

It’s possible to deploy the Remotion Studio onto a long-running server in the cloud, which can then be accessed by your non-technical team members using just a URL. Check out the [Deploy the Remotion Studio](/docs/studio/deploy-server) guide to learn how to do this.

## CLI [​](\#cli "Direct link to CLI")

Render a video using [`render` CLI command](/docs/cli/render):

```

bash

npx remotion render HelloWorld
```

```

bash

npx remotion render HelloWorld
```

Modify the composition ID to select a different video to render, or add an output path at the end if you want to override the default.

You can leave out the composition name and a picker will be shown:

```

bash

npx remotion render
```

```

bash

npx remotion render
```

## SSR [​](\#ssr "Direct link to SSR")

Remotion has a full-featured server-side rendering API. Read more about it on the [server-side rendering API](/docs/ssr).

## AWS Lambda [​](\#aws-lambda "Direct link to AWS Lambda")

Check out [Remotion Lambda](/docs/lambda).

## GitHub Actions [​](\#github-actions "Direct link to GitHub Actions")

You can also render a [video using a GitHub Action.](/docs/ssr#render-using-github-actions)

## Google Cloud Run [​](\#google-cloud-run "Direct link to Google Cloud Run")

An alpha version of Remotion for Cloud Run is [available](/docs/cloudrun).

We plan to change it in the future so it shares a runtime with Remotion Lambda.

## Variants [​](\#variants "Direct link to Variants")

### Audio-only [​](\#audio-only "Direct link to Audio-only")

Instead of rendering a video, [you can also just export the audio.](/docs/encoding#audio-only-export)

### Image Sequence [​](\#image-sequence "Direct link to Image Sequence")

Instead of encoding as a video, you can use the [`--sequence`](/docs/cli/render#--sequence) command to output a series of image.

### Still images [​](\#still-images "Direct link to Still images")

If you want a single image, you can do so using [the CLI or Node.JS API](/docs/stills).

### GIF [​](\#gif "Direct link to GIF")

See: [Render as GIF](/docs/render-as-gif)

### Transparent videos [​](\#transparent-videos "Direct link to Transparent videos")

See: [Creating overlays](/docs/overlay)

### See also [​](\#see-also "Direct link to See also")

- [CLI options](/docs/cli)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
How props get resolved](/docs/props-resolution) [Next\
\
Encoding Guide](/docs/encoding)

- [Remotion Studio](#remotion-studio)
  - [Remotion Studio deployment](#remotion-studio-deployment)
- [CLI](#cli)
- [SSR](#ssr)
- [AWS Lambda](#aws-lambda)
- [GitHub Actions](#github-actions)
- [Google Cloud Run](#google-cloud-run)
- [Variants](#variants)
  - [Audio-only](#audio-only)
  - [Image Sequence](#image-sequence)
  - [Still images](#still-images)
  - [GIF](#gif)
  - [Transparent videos](#transparent-videos)
  - [See also](#see-also)
