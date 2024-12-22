---
title: Importing assets
source: Remotion Documentation
last_updated: 2024-12-22
---

# Importing assets

- [Home page](/)
- Designing visuals
- Assets

On this page

# Importing assets

To import assets in Remotion, create a `public/` folder in your project and use [`staticFile()`](/docs/staticfile) to import it.

```

txt

my-video/
├─ node_modules/
├─ public/
│  ├─ logo.png
├─ src/
│  ├─ MyComp.tsx
│  ├─ Root.tsx
│  ├─ index.ts
├─ package.json
```

```

txt

my-video/
├─ node_modules/
├─ public/
│  ├─ logo.png
├─ src/
│  ├─ MyComp.tsx
│  ├─ Root.tsx
│  ├─ index.ts
├─ package.json
```

```

src/MyComp.tsx
tsx

import { Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src={staticFile("logo.png")} />;
};
```

```

src/MyComp.tsx
tsx

import { Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src={staticFile("logo.png")} />;
};
```

## Using images [​](\#using-images "Direct link to Using images")

Use the [`<Img/>`](/docs/img) tag from Remotion.

```

MyComp.tsx
tsx

import { Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src={staticFile("logo.png")} />;
};
```

```

MyComp.tsx
tsx

import { Img, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src={staticFile("logo.png")} />;
};
```

You can also pass a URL:

```

MyComp.tsx
tsx

import { Img } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src="https://picsum.photos/id/237/200/300" />;
};
```

```

MyComp.tsx
tsx

import { Img } from "remotion";

export const MyComp: React.FC = () => {
  return <Img src="https://picsum.photos/id/237/200/300" />;
};
```

## Using image sequences [​](\#using-image-sequences "Direct link to Using image sequences")

If you have a series of images, for example exported from another program like After Effects or Rotato, you can interpolate the path to create a dynamic import.

```

txt

my-video/
├─ public/
│  ├─ frame1.png
│  ├─ frame2.png
│  ├─ frame3.png
├─ package.json
```

```

txt

my-video/
├─ public/
│  ├─ frame1.png
│  ├─ frame2.png
│  ├─ frame3.png
├─ package.json
```

```

tsx

import { Img, staticFile, useCurrentFrame } from "remotion";

const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return <Img src={staticFile(`/frame${frame}.png`)} />;
};
```

```

tsx

import { Img, staticFile, useCurrentFrame } from "remotion";

const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return <Img src={staticFile(`/frame${frame}.png`)} />;
};
```

## Using videos [​](\#using-videos "Direct link to Using videos")

Use the [`<OffthreadVideo />`](/docs/offthreadvideo) or [`<Video />`](/docs/video) component to keep the timeline and your video in sync.

```

tsx

import { OffthreadVideo, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile("vid.webm")} />;
};
```

```

tsx

import { OffthreadVideo, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <OffthreadVideo src={staticFile("vid.webm")} />;
};
```

Loading videos via URL is also possible:

```

tsx

import { OffthreadVideo } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

```

tsx

import { OffthreadVideo } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <OffthreadVideo src="https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4" />
  );
};
```

See also: [Which video formats does Remotion support?](/docs/miscellaneous/video-formats)

## Using Audio [​](\#using-audio "Direct link to Using Audio")

Use the [`<Audio/ >`](/docs/audio) component.

```

tsx

import { Audio, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Audio src={staticFile("tune.mp3")} />;
};
```

```

tsx

import { Audio, staticFile } from "remotion";

export const MyComp: React.FC = () => {
  return <Audio src={staticFile("tune.mp3")} />;
};
```

Loading audio from an URL is also possible:

```

tsx

import { Audio } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <Audio src="https://file-examples.com/storage/fe48a63c5264cbd519788b3/2017/11/file_example_MP3_700KB.mp3" />
  );
};
```

```

tsx

import { Audio } from "remotion";

export const MyComp: React.FC = () => {
  return (
    <Audio src="https://file-examples.com/storage/fe48a63c5264cbd519788b3/2017/11/file_example_MP3_700KB.mp3" />
  );
};
```

See the [audio guide](/docs/using-audio) for guidance on including audio.

## Using CSS [​](\#using-css "Direct link to Using CSS")

Put the .css file alongside your JavaScript source files and use an `import` statement.

```

txt

my-video/
├─ node_modules/
├─ src/
│  ├─ style.css
│  ├─ MyComp.tsx
│  ├─ Root.tsx
│  ├─ index.ts
├─ package.json
```

```

txt

my-video/
├─ node_modules/
├─ src/
│  ├─ style.css
│  ├─ MyComp.tsx
│  ├─ Root.tsx
│  ├─ index.ts
├─ package.json
```

```

MyComp.tsx
tsx

import "./style.css";
```

```

MyComp.tsx
tsx

import "./style.css";
```

note

Want to use SASS, Tailwind or similar? [See examples on how to override the Webpack configuration](/docs/webpack).

## Using Fonts [​](\#using-fonts "Direct link to Using Fonts")

[Read the separate page for fonts.](/docs/fonts)

## `import` statements [​](\#import-statements "Direct link to import-statements")

As an alternative way to import files, Remotion allows you to `import` or `require()` several types of files in your project:

- Images ( `.png`, `.svg`, `.jpg`, `.jpeg`, `.webp`, `.gif`, `.bmp`)
- Videos ( `.webm`, `.mov`, `.mp4`)
- Audio ( `.mp3`, `.wav`, `.aac`, `.m4a`)
- Fonts ( `.woff`, `.woff2`, `.otf`, `.ttf`, `.eot`)

For example:

```

MyComp.tsx
tsx

import { Img } from "remotion";
import logo from "./logo.png";

export const MyComp: React.FC = () => {
  return <Img src={logo} />;
};
```

```

MyComp.tsx
tsx

import { Img } from "remotion";
import logo from "./logo.png";

export const MyComp: React.FC = () => {
  return <Img src={logo} />;
};
```

### Caveats [​](\#caveats "Direct link to Caveats")

While this was previously the main way of importing files, we now recommend against it because of the following reasons:

- Only the above listed file extensions are supported.
- The maximum file size is 2GB.
- Dynamic imports such as `require('img' + frame + '.png')` are [funky](/docs/webpack-dynamic-imports).

Prefer importing using [`staticFile()`](/docs/staticfile) if possible.

## Dynamic duration based on assets [​](\#dynamic-duration-based-on-assets "Direct link to Dynamic duration based on assets")

To make your videos duration dependent based on your assets, see: [Dynamic duration, FPS & dimensions](/docs/dynamic-metadata)

## Files outside of the project [​](\#files-outside-of-the-project "Direct link to Files outside of the project")

Remotion runs in the browser, so it does not have access to arbitrary files on your computer.

It is also not possible to use the `fs` module from Node.js in the browser.

Instead, put assets in the `public/` folder and use [`getStaticFiles()`](/docs/getstaticfiles) to enumerate them.

See [why does Remotion does not support absolute paths](/docs/miscellaneous/absolute-paths).

## Adding assets after bundling [​](\#adding-assets-after-bundling "Direct link to Adding assets after bundling")

Before rendering, the code gets bundled using Webpack, and only bundled assets can be accessed afterwards.

For this reason, assets that are being added to the public folder after [`bundle()`](/docs/bundle) is called will not be accessible during render.

However, if you use the [server-side rendering APIs](/docs/ssr-node), you can add assets to the `public` folder that is inside the bundle after the fact.

## Use `<Img>`, `<Video>` and `<Audio>` [​](\#use-img-video-and-audio "Direct link to use-img-video-and-audio")

**Prefer [`<Img />`](/docs/img) or [`<Gif />`](/docs/gif)** over the native `<img>` tag, `<Image>` from Next.js and CSS `background-image`.

**Prefer [`<OffthreadVideo />`](/docs/offthreadvideo) or [`<Video />`](/docs/video)** over the native `<video>` tag.

**Prefer [`<Audio />`](/docs/audio)** over the native `<audio>` tag.

**Prefer [`<IFrame />`](/docs/iframe)** over the native `<iframe>` tag.

By using the components from Remotion, you ensure that:

[1](#1)

The assets are fully loaded before the the frame is rendered

[2](#2) The images and videos are synchronized with Remotion's timeline.

## See also [​](\#see-also "Direct link to See also")

- [staticFile()](/docs/staticfile)
- [getStaticFiles()](/docs/getstaticfiles)
- [watchStaticFile()](/docs/watchstaticfile)
- [Why Remotion does not support absolute paths](/docs/miscellaneous/absolute-paths)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/importing-assets.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Transforms](/docs/transforms) [Next\
\
Layers](/docs/layers)

- [Using images](#using-images)
- [Using image sequences](#using-image-sequences)
- [Using videos](#using-videos)
- [Using Audio](#using-audio)
- [Using CSS](#using-css)
- [Using Fonts](#using-fonts)
- [`import` statements](#import-statements)
  - [Caveats](#caveats)
- [Dynamic duration based on assets](#dynamic-duration-based-on-assets)
- [Files outside of the project](#files-outside-of-the-project)
- [Adding assets after bundling](#adding-assets-after-bundling)
- [Use `<Img>`, `<Video>` and `<Audio>`](#use-img-video-and-audio)
- [See also](#see-also)
