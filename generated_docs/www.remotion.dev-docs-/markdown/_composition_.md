---
title: <Composition>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Composition>

- [Home page](/)
- [remotion](/docs/remotion)
- <Composition>

On this page

# <Composition>

This is the component to use to register a video to make it renderable and make it show up in the sidebar of the Remotion development interface.

A composition represents the video you want to create, as a collection of clips (for example, several `<Sequence>`) that will play back to back to form your video.

```

src/Root.tsx
tsx

import { Composition } from "remotion";

export const RemotionRoot: React.FC = () => {
  return (
    <>
      <Composition
        component={Component}
        durationInFrames={300}
        width={1080}
        height={1080}
        fps={30}
        id="test-render"
        defaultProps={{}}
      />
      {/* Additional compositions can be rendered */}
    </>
  );
};
```

```

src/Root.tsx
tsx

import { Composition } from "remotion";

export const RemotionRoot: React.FC = () => {
  return (
    <>
      <Composition
        component={Component}
        durationInFrames={300}
        width={1080}
        height={1080}
        fps={30}
        id="test-render"
        defaultProps={{}}
      />
      {/* Additional compositions can be rendered */}
    </>
  );
};
```

## API [​](\#api "Direct link to API")

A `<Composition />` should be placed within a fragment of the root component (which is registered using [`registerRoot()`](/docs/register-root)).

The component takes the following props:

### `id` [​](\#id "Direct link to id")

ID of the composition, as shown in the sidebar and also a unique identifier of the composition that you need to specify if you want to render it. The ID can only contain letters, numbers and `-`.

### `fps` [​](\#fps "Direct link to fps")

At how many frames the composition should be rendered.

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

How many frames the composition should be long.

### `height` [​](\#height "Direct link to height")

Height of the composition in pixels.

### `width` [​](\#width "Direct link to width")

Width of the composition in pixels.

### `component` **or** `lazyComponent` [​](\#component-or-lazycomponent "Direct link to component-or-lazycomponent")

Pass the component in directly **or** pass a function that returns a dynamic import. Passing neither or both of the props is an error.

note

If you use `lazyComponent`, Remotion will use React Suspense to load the component. Components will be compiled by Webpack as they are needed, which will reduce startup time of Remotion. If you use `lazyComponent`, you need to use a default export for your component. This is a restriction of React Suspense.

### `defaultProps` [​](\#defaultprops "Direct link to defaultprops")

_optional_

Give your component default props. Default props can be overriden using the [Props editor](/docs/visual-editing) and with [Input Props](/docs/props-resolution).

Props must be an object that contains only pure JSON-serializable values.

Functions, classes, and other non-serializable values will be lost while rendering.

(This restriction does not apply to the [`<Player>`](/docs/player/player), where you are allowed to pass functions as props.)

You may in addition use `Date`, `Map`, `Set` and [`staticFile()`](/docs/staticfile) in your default props and Remotion will guarantee the proper serialization.

note

Type your components using the `React.FC<{}>` type and the `defaultProps` prop will be typesafe.

note

Passing huge objects to `defaultProps` can be slow. [Learn how to avoid it.](/docs/troubleshooting/defaultprops-too-big)

note

Use a `type`, not an `interface` to type your `defaultProps`. [Learn why](/docs/4-0-migration#cannot-use-an-interface-for-props)

### `calculateMetadata` [​](\#calculatemetadata "Direct link to calculatemetadata")

See: [`calculateMetadata()`](/docs/calculate-metadata)

### `schema` [​](\#schema "Direct link to schema")

Pass a Zod schema which your default props must satisfy. By doing so, you enable [visual editing](/docs/visual-editing). See: [Define a schema](/docs/schemas)

## Example using `component` [​](\#example-using-component "Direct link to example-using-component")

```

tsx

import { Composition } from "remotion";
import { MyComp } from "./MyComp";

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        component={MyComp}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

```

tsx

import { Composition } from "remotion";
import { MyComp } from "./MyComp";

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        component={MyComp}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

## Example using `lazyComponent` [​](\#example-using-lazycomponent "Direct link to example-using-lazycomponent")

```

tsx

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        lazyComponent={() => import("./LazyComponent")}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

```

tsx

export const MyVideo = () => {
  return (
    <>
      <Composition
        id="my-comp"
        lazyComponent={() => import("./LazyComponent")}
        width={1080}
        height={1080}
        fps={30}
        durationInFrames={3 * 30}
      />
    </>
  );
};
```

## Organize compositions using folders [​](\#organize-compositions-using-folders "Direct link to Organize compositions using folders")

You can use the [`<Folder />`](/docs/folder) component to organize your compositions in the sidebar.

```

tsx

import { Composition, Folder } from "remotion";

export const Video = () => {
  return (
    <>
      <Folder name="Visuals">
        <Composition
          id="CompInFolder"
          durationInFrames={100}
          fps={30}
          width={1080}
          height={1080}
          component={Component}
        />
      </Folder>
      <Composition
        id="CompOutsideFolder"
        durationInFrames={100}
        fps={30}
        width={1080}
        height={1080}
        component={Component}
      />
    </>
  );
};
```

```

tsx

import { Composition, Folder } from "remotion";

export const Video = () => {
  return (
    <>
      <Folder name="Visuals">
        <Composition
          id="CompInFolder"
          durationInFrames={100}
          fps={30}
          width={1080}
          height={1080}
          component={Component}
        />
      </Folder>
      <Composition
        id="CompOutsideFolder"
        durationInFrames={100}
        fps={30}
        width={1080}
        height={1080}
        component={Component}
      />
    </>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Composition.tsx)
- [registerRoot()](/docs/register-root)
- [The fundamentals](/docs/the-fundamentals)
- [CLI options](/docs/cli)
- [`<Sequence />`](/docs/sequence)
- [`<Still />`](/docs/still)
- [`<Folder />`](/docs/folder)
- [Avoid huge payloads for `defaultProps`](/docs/troubleshooting/defaultprops-too-big)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/composition.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
cancelRender()](/docs/cancel-render) [Next\
\
continueRender()](/docs/continue-render)

- [API](#api)
  - [`id`](#id)
  - [`fps`](#fps)
  - [`durationInFrames`](#durationinframes)
  - [`height`](#height)
  - [`width`](#width)
  - [`component` **or** `lazyComponent`](#component-or-lazycomponent)
  - [`defaultProps`](#defaultprops)
  - [`calculateMetadata`](#calculatemetadata)
  - [`schema`](#schema)
- [Example using `component`](#example-using-component)
- [Example using `lazyComponent`](#example-using-lazycomponent)
- [Organize compositions using folders](#organize-compositions-using-folders)
- [See also](#see-also)
