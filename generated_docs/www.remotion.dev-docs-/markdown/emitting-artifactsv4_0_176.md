---
title: Emitting Artifactsv4.0.176
source: Remotion Documentation
last_updated: 2024-12-22
---

# Emitting Artifactsv4.0.176

- [Home page](/)
- [Rendering](/docs/render)
- Emitting Artifacts

On this page

# Emitting Artifacts [v4.0.176](https://github.com/remotion-dev/remotion/releases/v4.0.176)

Sometimes you wish to generate additional files when rendering your video. For example:

- A `.srt` subtitle file
- A `.txt` containing chapters of the video
- A `CREDITS` file for the assets used in the video
- Debug information from the render.

You can use the [`<Artifact>`](/docs/artifact) component to emit arbitrary files from your video.

note

Emitting artifacts is not currently supported by `@remotion/cloudrun`.

## Example [​](\#example "Direct link to Example")

```

MyComp.tsx
tsx

import React from 'react';
import {Artifact, useCurrentFrame} from 'remotion';
import {generateSubtitles} from './subtitles';

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();
  return (
    <>
      {frame === 0 ? (
        <Artifact filename="captions.srt" content={generateSubtitles()} />
      ) : null}
    </>
  );
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {Artifact, useCurrentFrame} from 'remotion';
import {generateSubtitles} from './subtitles';

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();
  return (
    <>
      {frame === 0 ? (
        <Artifact filename="captions.srt" content={generateSubtitles()} />
      ) : null}
    </>
  );
};
```

## Rules of artifacts [​](\#rules-of-artifacts "Direct link to Rules of artifacts")

[1](#1)

The asset should only be rendered for one single frame of the video.
Otherwise, the asset will get emitted multiple times.

[2](#2)

It is possible to emit multiple assets, but they may not have the
same filename.

[3](#3)

For the `content` prop it is possible to pass a `string`, or a `Uint8Array` for binary data. Passing an `
Uint8Array
` should not be considered faster due to it having to be serialized.

## Receiving artifacts [​](\#receiving-artifacts "Direct link to Receiving artifacts")

### In the CLI or Studio [​](\#in-the-cli-or-studio "Direct link to In the CLI or Studio")

Artifacts get saved to `out/[composition-id]/[filename]` when rendering a video.

### Using `renderMedia()`, `renderStill()` or `renderFrames()` [​](\#using-rendermedia-renderstill-or-renderframes "Direct link to using-rendermedia-renderstill-or-renderframes")

Use the [`onArtifact`](/docs/renderer/render-media#onartifact) callback to receive the artifacts.

```

render.mjs
tsx

import {renderMedia, OnArtifact} from '@remotion/renderer';

const onArtifact: OnArtifact = (artifact) => {
  console.log(artifact.filename); // string
  console.log(artifact.content); // string | Uint8Array
  console.log(artifact.frame); // number, frame in the composition which emitted this

  // Example action: Write the artifact to disk
  fs.writeFileSync(artifact.filename, artifact.content);
};

await renderMedia({
  composition,
  serveUrl,
  onArtifact,
  codec: 'h264',
  inputProps,
});
```

```

render.mjs
tsx

import {renderMedia, OnArtifact} from '@remotion/renderer';

const onArtifact: OnArtifact = (artifact) => {
  console.log(artifact.filename); // string
  console.log(artifact.content); // string | Uint8Array
  console.log(artifact.frame); // number, frame in the composition which emitted this

  // Example action: Write the artifact to disk
  fs.writeFileSync(artifact.filename, artifact.content);
};

await renderMedia({
  composition,
  serveUrl,
  onArtifact,
  codec: 'h264',
  inputProps,
});
```

### Using the Remotion Lambda CLI [​](\#using-the-remotion-lambda-cli "Direct link to Using the Remotion Lambda CLI")

When using [`npx remotion lambda render`](/docs/lambda/cli/render) or [`npx remotion lambda still`](/docs/lambda/cli/still), artifacts get saved to the S3 bucket under the key `renders/[render-id]/artifacts/[filename]`.

They will get logged to the console and you can click them to download them.

The `--privacy` option also applies to artifacts.

### Using `renderMediaOnLambda()` [​](\#using-rendermediaonlambda "Direct link to using-rendermediaonlambda")

When using [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda), artifacts get saved to the S3 bucket under the key `renders/[render-id]/artifacts/[filename]`.

You can obtain a list of currently received assets from [`getRenderProgress()`](/docs/lambda/getrenderprogress#artifacts).

```

progress.ts
tsx

import {getRenderProgress} from '@remotion/lambda/client';

const renderProgress = await getRenderProgress({
  renderId: 'hi',
  functionName: 'hi',
  bucketName: 'hi',
  region: 'eu-central-1',
});

for (const artifact of renderProgress.artifacts) {
  console.log(artifact.filename); // "hello-world.txt"
  console.log(artifact.sizeInBytes); // 12
  console.log(artifact.s3Url); // "https://s3.eu-central-1.amazonaws.com/remotion-lambda-abcdef/renders/abcdef/artifacts/hello-world.txt"
  console.log(artifact.s3Key); // "renders/abcdef/artifacts/hello-world.txt"
}
```

```

progress.ts
tsx

import {getRenderProgress} from '@remotion/lambda/client';

const renderProgress = await getRenderProgress({
  renderId: 'hi',
  functionName: 'hi',
  bucketName: 'hi',
  region: 'eu-central-1',
});

for (const artifact of renderProgress.artifacts) {
  console.log(artifact.filename); // "hello-world.txt"
  console.log(artifact.sizeInBytes); // 12
  console.log(artifact.s3Url); // "https://s3.eu-central-1.amazonaws.com/remotion-lambda-abcdef/renders/abcdef/artifacts/hello-world.txt"
  console.log(artifact.s3Key); // "renders/abcdef/artifacts/hello-world.txt"
}
```

### Using `renderStillOnLambda()` [​](\#using-renderstillonlambda "Direct link to using-renderstillonlambda")

When using [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda), artifacts get saved to the S3 bucket under the key `renders/[render-id]/artifacts/[filename]`.

You can obtain a list of received assets from [`artifacts`](/docs/lambda/renderstillonlambda#artifacts) field of `renderStillOnLambda()`.

```

still.ts
tsx

import {renderStillOnLambda} from '@remotion/lambda/client';

const stillResponse = await renderStillOnLambda({
  functionName,
  region,
  serveUrl,
  composition,
  inputProps,
  imageFormat,
  privacy,
});

for (const artifact of stillResponse.artifacts) {
  console.log(artifact.filename); // "hello-world.txt"
  console.log(artifact.sizeInBytes); // 12
  console.log(artifact.s3Url); // "https://s3.eu-central-1.amazonaws.com/remotion-lambda-abcdef/renders/abcdef/artifacts/hello-world.txt"
  console.log(artifact.s3Key); // "renders/abcdef/artifacts/hello-world.txt"
}
```

```

still.ts
tsx

import {renderStillOnLambda} from '@remotion/lambda/client';

const stillResponse = await renderStillOnLambda({
  functionName,
  region,
  serveUrl,
  composition,
  inputProps,
  imageFormat,
  privacy,
});

for (const artifact of stillResponse.artifacts) {
  console.log(artifact.filename); // "hello-world.txt"
  console.log(artifact.sizeInBytes); // 12
  console.log(artifact.s3Url); // "https://s3.eu-central-1.amazonaws.com/remotion-lambda-abcdef/renders/abcdef/artifacts/hello-world.txt"
  console.log(artifact.s3Key); // "renders/abcdef/artifacts/hello-world.txt"
}
```

### Using Cloud Run [​](\#using-cloud-run "Direct link to Using Cloud Run")

In the Cloud Run Alpha, emitting artifacts is not supported and will throw an error.

We plan on revising Cloud Run to use the same runtime as Lambda in the future and will bring this feature along.

## See also [​](\#see-also "Direct link to See also")

- [`<Artifact>`](/docs/artifact)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/artifacts.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Video formats](/docs/miscellaneous/video-formats) [Next\
\
Setting video metadata](/docs/metadata)

- [Example](#example)
- [Rules of artifacts](#rules-of-artifacts)
- [Receiving artifacts](#receiving-artifacts)
  - [In the CLI or Studio](#in-the-cli-or-studio)
  - [Using `renderMedia()`, `renderStill()` or `renderFrames()`](#using-rendermedia-renderstill-or-renderframes)
  - [Using the Remotion Lambda CLI](#using-the-remotion-lambda-cli)
  - [Using `renderMediaOnLambda()`](#using-rendermediaonlambda)
  - [Using `renderStillOnLambda()`](#using-renderstillonlambda)
  - [Using Cloud Run](#using-cloud-run)
- [See also](#see-also)
