---
title: Creating overlays
source: Remotion Documentation
last_updated: 2024-12-22
---

# Creating overlays

- [Home page](/)
- [Rendering](/docs/render)
- Creating overlays

On this page

# Creating overlays

If you want to export a Remotion video to use it as an overlay or transition in a conventional video editing software, you can export it as a transparent Apple ProRes file, which is supported by Final Cut Pro, Adobe Premiere and Davinci Resolve.

## Exporting a video as ProRes file [​](\#exporting-a-video-as-prores-file "Direct link to Exporting a video as ProRes file")

1

Make sure your composition has no background color. If you toggle the transparency option on in the editor, you should see a checkerboard background.

2

To export the video as ProRes file, setup the `remotion.config.ts` file as follows:

```

remotion.config.ts
tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva444p10le");
Config.setCodec("prores");
Config.setProResProfile("4444");
```

```

remotion.config.ts
tsx

import { Config } from "@remotion/cli/config";

Config.setVideoImageFormat("png");
Config.setPixelFormat("yuva444p10le");
Config.setCodec("prores");
Config.setProResProfile("4444");
```

and trigger the render with

```

bash

npx remotion render
```

```

bash

npx remotion render
```

Alternatively, you can set the settings directly on the command line:

```

bash

npx remotion render --image-format=png --pixel-format=yuva444p10le --codec=prores --prores-profile=4444
```

```

bash

npx remotion render --image-format=png --pixel-format=yuva444p10le --codec=prores --prores-profile=4444
```

## Use a template [​](\#use-a-template "Direct link to Use a template")

Toggle transparency

Use our template by cloning the [GitHub repo](https://github.com/remotion-dev/template-overlay) or running `npx create-video --overlay`.

## Use it in your Video Editor [​](\#use-it-in-your-video-editor "Direct link to Use it in your Video Editor")

Now, you can simply import the video into your video editing software. You can find the video in the out folder of your Remotion project.

## See also [​](\#see-also "Direct link to See also")

- [Transparent videos](/docs/transparent-videos)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/overlay.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Rendering GIFs](/docs/render-as-gif) [Next\
\
Quality Guide](/docs/quality)

- [Exporting a video as ProRes file](#exporting-a-video-as-prores-file)
- [Use a template](#use-a-template)
- [Use it in your Video Editor](#use-it-in-your-video-editor)
- [See also](#see-also)
