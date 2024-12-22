---
title: Passing props to a composition
source: Remotion Documentation
last_updated: 2024-12-22
---

# Passing props to a composition

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- Passing props

On this page

# Passing props to a composition

You can parametrize the content of the videos using [React properties ("props")](https://react.dev/learn/passing-props-to-a-component).

## Defining accepted props [​](\#defining-accepted-props "Direct link to Defining accepted props")

To define which props your video accepts, give your component the `React.FC` type and pass in a generic argument describing the shape of the props you want to accept:

```

src/MyComponent.tsx
tsx

type Props = {
  propOne: string;
  propTwo: number;
}

export const MyComponent: React.FC<Props> = ({propOne, propTwo}) => {
  return (
    <div>props: {propOne}, {propTwo}</div>
  );
}
```

```

src/MyComponent.tsx
tsx

type Props = {
  propOne: string;
  propTwo: number;
}

export const MyComponent: React.FC<Props> = ({propOne, propTwo}) => {
  return (
    <div>props: {propOne}, {propTwo}</div>
  );
}
```

## Define default props [​](\#define-default-props "Direct link to Define default props")

When registering a component that takes props as a composition, you must define default props:

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComponent } from "./MyComponent";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="my-video"
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={30}
        component={MyComponent}
        defaultProps={{
          propOne: "Hi",
          propTwo: 10,
        }}
      />
    </>
  );
};
```

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComponent } from "./MyComponent";

export const Root: React.FC = () => {
  return (
    <>
      <Composition
        id="my-video"
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={30}
        component={MyComponent}
        defaultProps={{
          propOne: "Hi",
          propTwo: 10,
        }}
      />
    </>
  );
};
```

Default props are useful so you don't preview your video with no data. [Default props will be overriden by input props](/docs/props-resolution).

## Define a schema [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#define-a-schema "Direct link to define-a-schema")

You can use [Zod](https://github.com/colinhacks/zod) to [define a typesafe schema for your composition](/docs/schemas).

## Input props [​](\#input-props "Direct link to Input props")

Input props are props that are passed in while invoking a render that can replace or override the default props.

note

Input props must be an object and serializable to JSON.

### Passing input props in the CLI [​](\#passing-input-props-in-the-cli "Direct link to Passing input props in the CLI")

When rendering, you can override default props by passing a [CLI](/docs/cli/render) flag. It must be either valid JSON or a path to a file that contains valid JSON.

```

Using inline JSON
bash

npx remotion render HelloWorld out/helloworld.mp4 --props='{"propOne": "Hi", "propTwo": 10}'
```

```

Using inline JSON
bash

npx remotion render HelloWorld out/helloworld.mp4 --props='{"propOne": "Hi", "propTwo": 10}'
```

```

Using a file path
bash

npx remotion render HelloWorld out/helloworld.mp4 --props=./path/to/props.json
```

```

Using a file path
bash

npx remotion render HelloWorld out/helloworld.mp4 --props=./path/to/props.json
```

### Passing input props when using server-side rendering [​](\#passing-input-props-when-using-server-side-rendering "Direct link to Passing input props when using server-side rendering")

When server-rendering using [`renderMedia()`](/docs/renderer/render-media) or [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda), you can pass props using the [`inputProps`](/docs/renderer/render-media#inputprops) option:

```

tsx

import { renderMedia, selectComposition } from "@remotion/renderer";

const inputProps = {
  titleText: "Hello World",
};

const composition = await selectComposition({
  serveUrl,
  id: "my-video",
  inputProps,
});

await renderMedia({
  composition,
  serveUrl,
  codec: "h264",
  outputLocation,
  inputProps,
});
```

```

tsx

import { renderMedia, selectComposition } from "@remotion/renderer";

const inputProps = {
  titleText: "Hello World",
};

const composition = await selectComposition({
  serveUrl,
  id: "my-video",
  inputProps,
});

await renderMedia({
  composition,
  serveUrl,
  codec: "h264",
  outputLocation,
  inputProps,
});
```

You should pass your `inputProps` to both [`selectComposition()`](/docs/renderer/select-composition) and [`renderMedia()`](/docs/renderer/render-media).

### Passing input props in GitHub Actions [​](\#passing-input-props-in-github-actions "Direct link to Passing input props in GitHub Actions")

[See: Render using GitHub Actions](/docs/ssr#render-using-github-actions)

When using GitHub Actions, you need to adjust the file at `.github/workflows/render-video.yml` to make the inputs in the `workflow_dispatch` section manually match the shape of the props your root component accepts.

```

yaml

workflow_dispatch:
  inputs:
    titleText:
      description: "Which text should it say?"
      required: true
      default: "Welcome to Remotion"
    titleColor:
      description: "Which color should it be in?"
      required: true
      default: "black"
```

```

yaml

workflow_dispatch:
  inputs:
    titleText:
      description: "Which text should it say?"
      required: true
      default: "Welcome to Remotion"
    titleColor:
      description: "Which color should it be in?"
      required: true
      default: "black"
```

### Retrieve input props [​](\#retrieve-input-props "Direct link to Retrieve input props")

Input props are passed to the [`component`](/docs/composition) of your [`<Composition>`](/docs/composition) directly and you can access them like regular React component props.

If you need the input props in your root component, use the [`getInputProps()`](/docs/get-input-props) function to retrieve input props.

## You can still use components normally [​](\#you-can-still-use-components-normally "Direct link to You can still use components normally")

Even if a component is registered as a composition, you can still use it like a regular React component and pass the props directly:

```

tsx

<MyComponent propOne="hi" propTwo={10} />
```

```

tsx

<MyComponent propOne="hi" propTwo={10} />
```

This is useful if you want to concatenate multiple scenes together. You can use a [`<Series>`](/docs/series) to play two components after each other:

```

ChainedScenes.tsx
tsx

import { Series } from "remotion";

const ChainedScenes = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={90}>
        <MyComponent propOne="hi" propTwo={10} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={90}>
        <AnotherComponent />
      </Series.Sequence>
    </Series>
  );
};
```

```

ChainedScenes.tsx
tsx

import { Series } from "remotion";

const ChainedScenes = () => {
  return (
    <Series>
      <Series.Sequence durationInFrames={90}>
        <MyComponent propOne="hi" propTwo={10} />
      </Series.Sequence>
      <Series.Sequence durationInFrames={90}>
        <AnotherComponent />
      </Series.Sequence>
    </Series>
  );
};
```

You may then register this "Master" component as an additional [`<Composition>`](/docs/the-fundamentals#compositions).

## See also [​](\#see-also "Direct link to See also")

- [Avoid huge payloads for `defaultProps`](/docs/troubleshooting/defaultprops-too-big)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/passing-props.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Passing props](/docs/parameterized-rendering) [Next\
\
Defining a schema](/docs/schemas)

- [Defining accepted props](#defining-accepted-props)
- [Define default props](#define-default-props)
- [Define a schema](#define-a-schema)
- [Input props](#input-props)
  - [Passing input props in the CLI](#passing-input-props-in-the-cli)
  - [Passing input props when using server-side rendering](#passing-input-props-when-using-server-side-rendering)
  - [Passing input props in GitHub Actions](#passing-input-props-in-github-actions)
  - [Retrieve input props](#retrieve-input-props)
- [You can still use components normally](#you-can-still-use-components-normally)
- [See also](#see-also)
