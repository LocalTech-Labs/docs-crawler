---
title: Deploy Remotion Studio as static site
source: Remotion Documentation
last_updated: 2024-12-22
---

# Deploy Remotion Studio as static site

- [Home page](/)
- Studio
- Deploy as static site

On this page

# Deploy Remotion Studio as static site

_available from v4.0.97_

You can deploy the Remotion Studio as a static site, for example to Vercel or Netlify.

While the Render button will be disabled, it may be used as a [Serve URL](/docs/terminology/serve-url) to pass to rendering APIs.

Make sure you are on at least v4.0.97 to use this feature - use `npx remotion upgrade` to upgrade.

## Export the Remotion Studio as a static site [​](\#export-the-remotion-studio-as-a-static-site "Direct link to Export the Remotion Studio as a static site")

To export the Remotion Studio as a static site, run the [`remotion bundle`](/docs/cli/bundle) command:

```

bash

npx remotion bundle
```

```

bash

npx remotion bundle
```

The output will be in the `build` folder.

We recommend to add `build` to your `.gitignore` file. The command will offer to do it for you when you run it.

## Deploy to Vercel [​](\#deploy-to-vercel "Direct link to Deploy to Vercel")

To deploy to Vercel, connect your repository to Vercel.

In the "Build and Output" settings, enable "OVERRIDE" and set the following:

- **Build Command**: `bunx remotion bundle`
- **Output Directory**: `build`
- **Installation Command**: _Leave default_

![Alt text](https://github.com/remotion-dev/remotion/assets/86873911/ac38e1fb-0c95-4f64-b2b9-b72b90d69f22)

note

Using bunx as a script runner is slightly faster than using `npx`.

## Deploy to Netlify [​](\#deploy-to-netlify "Direct link to Deploy to Netlify")

Connect your repository to Netlify.

- **Base directory**: _Leave default_
- **Build command**: `npx remotion bundle`
- **Publish directory**: `build`
- **Functions directory**: _Leave default_

## Deploy to GitHub Pages [​](\#deploy-to-github-pages "Direct link to Deploy to GitHub Pages")

[1](#1)

Create the following file in your repo:

```

.github/workflows/deploy-studio.yml
yaml

name: Deploy Remotion studio
on:
  workflow_dispatch:
  push:
    branches:
      - "main"
permissions:
  contents: write
jobs:
  render:
    name: Render video
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@main
      - name: install packages
        run: npm i
      - name: Bundle Studio
        run: npx remotion bundle --public-path="./"
      - name: Deploy Studio
        uses: JamesIves/github-pages-deploy-action@v4
        with:
          folder: build
```

```

.github/workflows/deploy-studio.yml
yaml

name: Deploy Remotion studio
on:
  workflow_dispatch:
  push:
    branches:
      - "main"
permissions:
  contents: write
jobs:
  render:
    name: Render video
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@main
      - name: install packages
        run: npm i
      - name: Bundle Studio
        run: npx remotion bundle --public-path="./"
      - name: Deploy Studio
        uses: JamesIves/github-pages-deploy-action@v4
        with:
          folder: build
```

The above code will deploy the bundled Remotion Studio to a branch `gh-pages`.

note

The --public-path flag for `npx remotion bundle` is available only from remotion version `4.0.127`

[3](#3)

Go to the settings of your repository (github.com/\[username\]/\[repo\]/settings/pages)

[4](#4)

In the Branch section, select the `gh-pages` branch and
click save.

## Rendering from a URL [​](\#rendering-from-a-url "Direct link to Rendering from a URL")

The deployed URL is a [Serve URL](/docs/terminology/serve-url) can also be used to render a video:

- CLI
- Node.js/Bun
- Lambda
- Cloud Run

Minimal example:

```

bash

npx remotion render https://remotion-helloworld.vercel.app
```

```

bash

npx remotion render https://remotion-helloworld.vercel.app
```

Specify "HelloWorld" [composition ID](/docs/terminology/composition#composition-id) and [input props](/docs/terminology/input-props):

```

bash

npx remotion render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

bash

npx remotion render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

render.mjs
tsx

const inputProps = {
  titleText: "Hello World",
};

const serveUrl = "https://remotion-helloworld.vercel.app";

const composition = await selectComposition({
  serveUrl,
  id: "HelloWorld",
  inputProps,
});

await renderMedia({
  composition,
  serveUrl,
  codec: "h264",
  inputProps,
});
```

```

render.mjs
tsx

const inputProps = {
  titleText: "Hello World",
};

const serveUrl = "https://remotion-helloworld.vercel.app";

const composition = await selectComposition({
  serveUrl,
  id: "HelloWorld",
  inputProps,
});

await renderMedia({
  composition,
  serveUrl,
  codec: "h264",
  inputProps,
});
```

Refer to [`renderMedia()`](/docs/renderer/render-media) to see all available options.

```

CLI
bash

npx remotion lambda render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

CLI
bash

npx remotion lambda render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

Node.js/Bun
tsx

import { renderMediaOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "HelloWorld",
  serveUrl: "https://remotion-helloworld.vercel.app",
  codec: "h264",
  inputProps: {
    titleText: "Hello World",
  },
});
```

```

Node.js/Bun
tsx

import { renderMediaOnLambda } from "@remotion/lambda/client";

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "HelloWorld",
  serveUrl: "https://remotion-helloworld.vercel.app",
  codec: "h264",
  inputProps: {
    titleText: "Hello World",
  },
});
```

You need to complete the [Remotion Lambda Setup](/docs/lambda/setup) first.

```

CLI
bash

npx remotion cloudrun render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

CLI
bash

npx remotion cloudrun render https://remotion-helloworld.vercel.app HelloWorld --props '{"titleText":"Hello World"}'
```

```

Node.js/Bun
tsx

import { renderMediaOnCloudrun } from "@remotion/cloudrun/client";

const result = await renderMediaOnCloudrun({
  region: "us-east1",
  serviceName: "remotion-render-bds9aab",
  composition: "HelloWorld",
  serveUrl: "https://remotion-helloworld.vercel.app",
  codec: "h264",
  inputProps: {
    titleText: "Hello World",
  },
});

if (result.type === "success") {
  console.log(result.bucketName);
  console.log(result.renderId);
}
```

```

Node.js/Bun
tsx

import { renderMediaOnCloudrun } from "@remotion/cloudrun/client";

const result = await renderMediaOnCloudrun({
  region: "us-east1",
  serviceName: "remotion-render-bds9aab",
  composition: "HelloWorld",
  serveUrl: "https://remotion-helloworld.vercel.app",
  codec: "h264",
  inputProps: {
    titleText: "Hello World",
  },
});

if (result.type === "success") {
  console.log(result.bucketName);
  console.log(result.renderId);
}
```

You need to complete the [Remotion Cloud Run Setup](/docs/cloudrun-alpha) first.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/deploy-static.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Quick Switcher](/docs/studio/quick-switcher) [Next\
\
Deploy to a VPS](/docs/studio/deploy-server)

- [Export the Remotion Studio as a static site](#export-the-remotion-studio-as-a-static-site)
- [Deploy to Vercel](#deploy-to-vercel)
- [Deploy to Netlify](#deploy-to-netlify)
- [Deploy to GitHub Pages](#deploy-to-github-pages)
- [Rendering from a URL](#rendering-from-a-url)
