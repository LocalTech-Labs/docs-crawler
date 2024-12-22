---
title: <Artifact>v4.0.176
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Artifact>v4.0.176

- [Home page](/)
- [remotion](/docs/remotion)
- <Artifact>

On this page

# `<Artifact>` [v4.0.176](https://github.com/remotion-dev/remotion/releases/v4.0.176)

By rendering an `<Artifact>` tag in your Remotion markup, [an extra file will get emitted during rendering](/docs/artifacts).

```

MyComp.tsx
tsx

import { Artifact, useCurrentFrame } from "remotion";

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return frame === 0 ? (
    <Artifact filename="my-file.txt" content="Hello World!" />
  ) : null;
};
```

```

MyComp.tsx
tsx

import { Artifact, useCurrentFrame } from "remotion";

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return frame === 0 ? (
    <Artifact filename="my-file.txt" content="Hello World!" />
  ) : null;
};
```

If rendered on the CLI or via the Studio, this will emit an additional file:

```

$ npx remotion render MyComp
+ out/MyComp.mp4
+ my-file.txt (12B)
```

```

$ npx remotion render MyComp
+ out/MyComp.mp4
+ my-file.txt (12B)
```

It is allowed for a composition to emit multiple files.

However, the file names must be unique.

The component will get evaluated on every frame, which means if you want to emit just one file, only render it on one frame.

```

❌ Will generate an asset on every frame and throw an error because the file name is not unique
tsx

import { Artifact, useCurrentFrame } from "remotion";

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return frame === 0 ? (
    <Artifact filename="my-file.txt" content="Hello World!" />
  ) : null;
};
```

```

❌ Will generate an asset on every frame and throw an error because the file name is not unique
tsx

import { Artifact, useCurrentFrame } from "remotion";

export const MyComp: React.FC = () => {
  const frame = useCurrentFrame();

  return frame === 0 ? (
    <Artifact filename="my-file.txt" content="Hello World!" />
  ) : null;
};
```

## API [​](\#api "Direct link to API")

### `filename` [​](\#filename "Direct link to filename")

A string that is the name of the file that will be emitted.

Use forward slashes only, even on Windows.

Must match the regex `/^([0-9a-zA-Z-!_.*'()/:&$@=;+,?]+)/g`.

### `content` [​](\#content "Direct link to content")

A `string` or `Uint8Array` that is the content of the file that will be emitted. Don't consider an `Uint8Array` to be faster, because it needs to be serialized.

## See also [​](\#see-also "Direct link to See also")

- [Emitting Artifacts](/docs/artifacts)
- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/components/Artifact.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/artifact.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Audio>](/docs/audio) [Next\
\
calculateMetadata()](/docs/calculate-metadata)

- [API](#api)
  - [`filename`](#filename)
  - [`content`](#content)
- [See also](#see-also)
