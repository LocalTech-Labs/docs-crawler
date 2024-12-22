---
title: Render videos programmatically from a dataset
source: Remotion Documentation
last_updated: 2024-12-22
---

# Render videos programmatically from a dataset

- [Home page](/)
- [Rendering](/docs/render)
- Render a dataset

On this page

# Render videos programmatically from a dataset

You can use Remotion to do a batch render to create many videos based on a dataset. In the following example, we are going to turn a JSON dataset into a series of videos.

We'll start by creating a blank Remotion project:

- npm
- pnpm
- bun
- yarn

```

bash

npm init video --blank
```

```

bash

npm init video --blank
```

```

bash

pnpm create video --blank
```

```

bash

pnpm create video --blank
```

```

bash

bun create video --blank
```

```

bash

bun create video --blank
```

```

bash

yarn create video --blank
```

```

bash

yarn create video --blank
```

## Sample dataset [​](\#sample-dataset "Direct link to Sample dataset")

JSON is the most convenient format to import in Remotion. If your dataset is in a different format, you can convert it using one of many available libraries on NPM.

```

my-data.ts
ts

export const data = [
  {
    name: "React",
    repo: "facebook/react",
    logo: "https://upload.wikimedia.org/wikipedia/commons/a/a7/React-icon.svg",
  },
  {
    name: "Remotion",
    repo: "remotion-dev/remotion",
    logo: "https://github.com/remotion-dev/logo/raw/main/withouttitle/element-0.png",
  },
];
```

```

my-data.ts
ts

export const data = [
  {
    name: "React",
    repo: "facebook/react",
    logo: "https://upload.wikimedia.org/wikipedia/commons/a/a7/React-icon.svg",
  },
  {
    name: "Remotion",
    repo: "remotion-dev/remotion",
    logo: "https://github.com/remotion-dev/logo/raw/main/withouttitle/element-0.png",
  },
];
```

## Sample component [​](\#sample-component "Direct link to Sample component")

This component will animate a title, subtitle and image using Remotion. Replace the contents of the `src/Composition.tsx` file with the following:

```

Composition.tsx
tsx

import React from "react";
import {
  AbsoluteFill,
  Img,
  interpolate,
  spring,
  useCurrentFrame,
  useVideoConfig,
} from "remotion";

type Props = {
  name: string;
  logo: string;
  repo: string;
};

export const MyComposition: React.FC<Props> = ({ name, repo, logo }) => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  const scale = spring({
    fps,
    frame: frame - 10,
    config: {
      damping: 100,
    },
  });

  const opacity = interpolate(frame, [30, 40], [0, 1], {
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });
  const moveY = interpolate(frame, [20, 30], [10, 0], {
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });

  return (
    <AbsoluteFill
      style={{
        scale: String(scale),
        backgroundColor: "white",
        fontWeight: "bold",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          gap: 20,
        }}
      >
        <Img
          src={logo}
          style={{
            height: 80,
          }}
        />
        <div
          style={{
            display: "flex",
            flexDirection: "column",
          }}
        >
          <div
            style={{
              fontSize: 40,
              transform: `translateY(${moveY}px)`,
              lineHeight: 1,
            }}
          >
            {name}
          </div>
          <div
            style={{
              fontSize: 20,
              opacity,
              lineHeight: 1.25,
            }}
          >
            {repo}
          </div>
        </div>
      </div>
    </AbsoluteFill>
  );
};
```

```

Composition.tsx
tsx

import React from "react";
import {
  AbsoluteFill,
  Img,
  interpolate,
  spring,
  useCurrentFrame,
  useVideoConfig,
} from "remotion";

type Props = {
  name: string;
  logo: string;
  repo: string;
};

export const MyComposition: React.FC<Props> = ({ name, repo, logo }) => {
  const frame = useCurrentFrame();
  const { fps } = useVideoConfig();

  const scale = spring({
    fps,
    frame: frame - 10,
    config: {
      damping: 100,
    },
  });

  const opacity = interpolate(frame, [30, 40], [0, 1], {
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });
  const moveY = interpolate(frame, [20, 30], [10, 0], {
    extrapolateLeft: "clamp",
    extrapolateRight: "clamp",
  });

  return (
    <AbsoluteFill
      style={{
        scale: String(scale),
        backgroundColor: "white",
        fontWeight: "bold",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          gap: 20,
        }}
      >
        <Img
          src={logo}
          style={{
            height: 80,
          }}
        />
        <div
          style={{
            display: "flex",
            flexDirection: "column",
          }}
        >
          <div
            style={{
              fontSize: 40,
              transform: `translateY(${moveY}px)`,
              lineHeight: 1,
            }}
          >
            {name}
          </div>
          <div
            style={{
              fontSize: 20,
              opacity,
              lineHeight: 1.25,
            }}
          >
            {repo}
          </div>
        </div>
      </div>
    </AbsoluteFill>
  );
};
```

Remotion

remotion-dev/remotion

0:00 / 0:02

## Writing the script [​](\#writing-the-script "Direct link to Writing the script")

In order to render our videos, we'll first need to bundle our project using Webpack and prepare it for rendering.
This can be done by using the [`bundle()`](/docs/bundle) function from the [`@remotion/bundler`](/docs/bundler) package. Make sure to include the webpack override in the bundle if you have one.

```

ts

import { bundle } from "@remotion/bundler";
import { webpackOverride } from "./webpack-override";

const bundleLocation = await bundle({
  entryPoint: "./src/index.ts",
  // If you have a webpack override, don't forget to add it
  webpackOverride: webpackOverride,
});
```

```

ts

import { bundle } from "@remotion/bundler";
import { webpackOverride } from "./webpack-override";

const bundleLocation = await bundle({
  entryPoint: "./src/index.ts",
  // If you have a webpack override, don't forget to add it
  webpackOverride: webpackOverride,
});
```

## Rendering videos [​](\#rendering-videos "Direct link to Rendering videos")

Import the dataset and loop over each entry.

Fetch the composition metadata for each render using [`selectComposition()`](/docs/renderer/select-composition). You have the opportunity to parametrize the duration, width and height of the composition based on the data entry with the [`calculateMetadata()`](/docs/dynamic-metadata) function.

Trigger a render using [`renderMedia()`](/docs/renderer/render-media) and pass the data entry as [`inputProps`](/docs/renderer/render-media#inputprops). This will pass the object as React props to the component above.

```

ts

import { renderMedia, selectComposition } from "@remotion/renderer";
import { data } from "./dataset";

for (const entry of data) {
  const composition = await selectComposition({
    serveUrl: bundleLocation,
    id: compositionId,
    inputProps: entry,
  });

  await renderMedia({
    composition,
    serveUrl: bundleLocation,
    codec: "h264",
    outputLocation: `out/${entry.name}.mp4`,
    inputProps: entry,
  });
}
```

```

ts

import { renderMedia, selectComposition } from "@remotion/renderer";
import { data } from "./dataset";

for (const entry of data) {
  const composition = await selectComposition({
    serveUrl: bundleLocation,
    id: compositionId,
    inputProps: entry,
  });

  await renderMedia({
    composition,
    serveUrl: bundleLocation,
    codec: "h264",
    outputLocation: `out/${entry.name}.mp4`,
    inputProps: entry,
  });
}
```

note

It is not recommended to render more than one video at once.

## Full script [​](\#full-script "Direct link to Full script")

```

render.mjs
ts

import { selectComposition, renderMedia } from "@remotion/renderer";
import { webpackOverride } from "./webpack-override";
import { bundle } from "@remotion/bundler";
import { data } from "./dataset";

const compositionId = "MyComp";

const bundleLocation = await bundle({
  entryPoint: "./src/index.ts",
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: webpackOverride,
});

for (const entry of data) {
  const composition = await selectComposition({
    serveUrl: bundleLocation,
    id: compositionId,
    inputProps: entry,
  });

  await renderMedia({
    composition,
    serveUrl: bundleLocation,
    codec: "h264",
    outputLocation: `out/${entry.name}.mp4`,
    inputProps: entry,
  });
}
```

```

render.mjs
ts

import { selectComposition, renderMedia } from "@remotion/renderer";
import { webpackOverride } from "./webpack-override";
import { bundle } from "@remotion/bundler";
import { data } from "./dataset";

const compositionId = "MyComp";

const bundleLocation = await bundle({
  entryPoint: "./src/index.ts",
  // If you have a webpack override in remotion.config.ts, pass it here as well.
  webpackOverride: webpackOverride,
});

for (const entry of data) {
  const composition = await selectComposition({
    serveUrl: bundleLocation,
    id: compositionId,
    inputProps: entry,
  });

  await renderMedia({
    composition,
    serveUrl: bundleLocation,
    codec: "h264",
    outputLocation: `out/${entry.name}.mp4`,
    inputProps: entry,
  });
}
```

## Running the script [​](\#running-the-script "Direct link to Running the script")

- Node
- Bun

```

bash

node render.mjs
```

```

bash

node render.mjs
```

To use TypeScript, rename the file to `render.ts`, install `ts-node` from npm and run `ts-node render.ts`. If you get errors, wrap the asynchronous code in an async function and call it.

```

bash

bun render.mjs
```

```

bash

bun render.mjs
```

TypeScript should work out of the box if you rename your file to `render.ts`.

## Rendering videos from a CSV dataset [​](\#rendering-videos-from-a-csv-dataset "Direct link to Rendering videos from a CSV dataset")

Use a package like [`csv2json`](https://www.npmjs.com/package/csv2json) to convert your dataset into a JSON.

## See also [​](\#see-also "Direct link to See also")

- [Example repository using a dataset](https://github.com/alexfernandez803/remotion-dataset)

CONTRIBUTORS

[![alexfernandez803](https://github.com/alexfernandez803.png)\
\
**alexfernandez803** \
\
Author](https://github.com/alexfernandez803) [![ThePerfectSystem](https://github.com/ThePerfectSystem.png)\
\
**ThePerfectSystem** \
\
Author](https://github.com/ThePerfectSystem) [![JonnyBurger](https://github.com/JonnyBurger.png)\
\
**JonnyBurger** \
\
Editor](https://github.com/JonnyBurger)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/dataset-render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Quality Guide](/docs/quality) [Next\
\
Render all compositions](/docs/render-all)

- [Sample dataset](#sample-dataset)
- [Sample component](#sample-component)
- [Writing the script](#writing-the-script)
- [Rendering videos](#rendering-videos)
- [Full script](#full-script)
- [Running the script](#running-the-script)
- [Rendering videos from a CSV dataset](#rendering-videos-from-a-csv-dataset)
- [See also](#see-also)
