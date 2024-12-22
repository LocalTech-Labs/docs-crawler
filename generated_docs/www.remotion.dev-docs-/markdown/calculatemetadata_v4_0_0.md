---
title: calculateMetadata()v4.0.0
source: Remotion Documentation
last_updated: 2024-12-22
---

# calculateMetadata()v4.0.0

- [Home page](/)
- [remotion](/docs/remotion)
- calculateMetadata()

On this page

# calculateMetadata() [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0)

`calculateMetadata` is a prop that gets passed to [`<Composition>`](/docs/composition) and takes a callback function which may transform metadata.

Use it if you:

[1](#1)

need to make the `durationInFrames`, `width`, `height`, or `fps` [dynamic based on some data](/docs/dynamic-metadata)

[2](#2) want to transform the props passed to the composition, for example
to do [data fetching](/docs/data-fetching)

[3](#3) want to add a per-composition default codec

[4](#4) want to run code once, before the actual render starts.

The `calculateMetadata()` function is called a single time, independently from the concurrency of the render.

It runs in a separate tab, as part of the render process calling [`selectComposition()`](/docs/renderer/select-composition).

## Usage / API [​](\#usage--api "Direct link to Usage / API")

Define a function, you may type it using `CalculateMetadataFunction` \- the generic argument takes the props of the component of your composition:

```

src/Root.tsx
tsx

import React from "react";
import { CalculateMetadataFunction, Composition } from "remotion";
import { MyComponent, MyComponentProps } from "./MyComp";

const calculateMetadata: CalculateMetadataFunction<MyComponentProps> = ({
  props,
  defaultProps,
  abortSignal,
}) => {
  return {
    // Change the metadata
    durationInFrames: props.duration,
    // or transform some props
    props,
    // or add per-composition default codec
    defaultCodec: "h264",
  };
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComponent}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        text: "Hello World",
        duration: 1,
      }}
      calculateMetadata={calculateMetadata}
    />
  );
};
```

```

src/Root.tsx
tsx

import React from "react";
import { CalculateMetadataFunction, Composition } from "remotion";
import { MyComponent, MyComponentProps } from "./MyComp";

const calculateMetadata: CalculateMetadataFunction<MyComponentProps> = ({
  props,
  defaultProps,
  abortSignal,
}) => {
  return {
    // Change the metadata
    durationInFrames: props.duration,
    // or transform some props
    props,
    // or add per-composition default codec
    defaultCodec: "h264",
  };
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComponent}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        text: "Hello World",
        duration: 1,
      }}
      calculateMetadata={calculateMetadata}
    />
  );
};
```

As argument, you get an object with the following properties:

- `defaultProps`: Only the default props, taken from the [`defaultProps`](/docs/composition#defaultprops) prop or the Remotion Studio Data sidebar.
- `props`: The [resolved props](/docs/props-resolution), taking input props into account.
- `abortSignal`: An [`AbortSignal`](https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal) which you can use to abort network requests if [the default props have been changed](/docs/data-fetching#aborting-stale-requests) in the meanwhile.
- `compositionId` ( _available from v4.0.98_): The ID of the current composition

The function must return an pure JSON-serializable object, which can contain the following properties:

- `props` _optional_: The [final props](/docs/props-resolution) the component receives. It must have the [same shape](/docs/data-fetching#usage) as the `props` received by this function. The props must be a plain object and must not contain any non-JSON-serializable values, except `Date`, `Map`, `Set` and [`staticFile()`](/docs/staticfile).
- `durationInFrames` _optional_: The duration of the composition in frames
- `width` _optional_: The width of the composition in pixels
- `height` _optional_: The height of the composition in pixels
- `fps` _optional_: The frames per second of the composition
- `defaultCodec` _optional_: The default codec to use for the composition.

If you return a field, it will take precendence over the props directly passed to the composition. The defaultCodec returned will have a higher priority than the config file, but less priority than explicitly passing the option to renderMedia() i.e. the composition will be rendered with this codec if nothing with higher priority was specified.

The function may be `async`.

The function must resolve within the [timeout](/docs/delay-render#timeout).

The function will get executed every time the props in the [Remotion Studio](/docs/visual-editing) changes.

## See also [​](\#see-also "Direct link to See also")

- [`<Composition>`](/docs/composition)
- [Data fetching](/docs/data-fetching)
- [Dynamic metadata](/docs/dynamic-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/calculate-metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Artifact>](/docs/artifact) [Next\
\
cancelRender()](/docs/cancel-render)

- [Usage / API](#usage--api)
- [See also](#see-also)