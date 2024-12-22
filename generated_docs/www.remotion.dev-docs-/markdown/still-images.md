---
title: Still images
source: Remotion Documentation
last_updated: 2024-12-22
---

# Still images

- [Home page](/)
- [Rendering](/docs/render)
- Still images

On this page

# Still images

_Available from v2.3_

Remotion is a great solution for rendering thumbnails of videos or dynamic still images too. See our [Still template](https://github.com/remotion-dev/template-still) for an easy way to get started.

If you already have a Remotion project, read on how you can render stills.

## Defining a still [​](\#defining-a-still "Direct link to Defining a still")

Use the [`<Still />`](/docs/still) component instead of the [`<Composition />`](/docs/composition) one to define a still. The timeline will disappear, and you will not have to define a duration or FPS value.

## Rendering via CLI [​](\#rendering-via-cli "Direct link to Rendering via CLI")

You can use the [`npx remotion still`](/docs/cli/still) command to render a still image. Example command:

```

bash

npx remotion still --props='{"custom": "data"}' my-comp out.png
```

```

bash

npx remotion still --props='{"custom": "data"}' my-comp out.png
```

You can use the `--image-format` flag to determine the output format. The default format is `png`, with `jpeg`, `webp` and `pdf` being the other options.

By default the frame with number of a composition is being rendered, you can control it using the `--frame` flag.

## Rendering using Node.JS [​](\#rendering-using-nodejs "Direct link to Rendering using Node.JS")

You can use the [`renderStill()`](/docs/renderer/render-still) Node.JS API to render a still frame programmatically.

## Rendering using serverless [​](\#rendering-using-serverless "Direct link to Rendering using serverless")

You can use [Remotion Lambda](/lambda) to render stills:

- Via [CLI](/docs/lambda/cli/still)
- Via the [renderStillOnLambda](/docs/lambda/renderstillonlambda) API.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/stills.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Encoding Guide](/docs/encoding) [Next\
\
Output scaling](/docs/scaling)

- [Defining a still](#defining-a-still)
- [Rendering via CLI](#rendering-via-cli)
- [Rendering using Node.JS](#rendering-using-nodejs)
- [Rendering using serverless](#rendering-using-serverless)
