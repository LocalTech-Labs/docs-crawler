---
title: <Freeze>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Freeze>

- [Home page](/)
- [remotion](/docs/remotion)
- <Freeze>

On this page

# <Freeze>

_Available from v2.2.0._

When using the `<Freeze/>` component, all of its children will freeze to the frame that you specify as a prop.

If a component is a child of `<Freeze/>`, calling the [`useCurrentFrame()`](/docs/use-current-frame) hook will always return the frame number you specify, irrespective of any [`<Sequence>`](/docs/sequence).

[`<Video />`](/docs/video) and [`<OffthreadVideo />`](/docs/video) elements will be paused and [`<Audio />`](/docs/audio) elements will render muted.

## Example [​](\#example "Direct link to Example")

```

MyComp.tsx
tsx

import { Freeze } from "remotion";

const MyVideo = () => {
  return (
    <Freeze frame={30}>
      <BlueSquare />
    </Freeze>
  );
};
```

```

MyComp.tsx
tsx

import { Freeze } from "remotion";

const MyVideo = () => {
  return (
    <Freeze frame={30}>
      <BlueSquare />
    </Freeze>
  );
};
```

## API [​](\#api "Direct link to API")

The Freeze component is a high order component and accepts, besides it's children, the following props:

### `frame` [​](\#frame "Direct link to frame")

At which frame it's children should freeze.

### `active` [v4.0.127](https://github.com/remotion-dev/remotion/releases/v4.0.127) [​](\#active "Direct link to active")

Deactivate freezing component by passing `false`.

You may also pass a callback function and accept the current frame and return a boolean.

```

From frame 30 on
tsx

import { Freeze } from "remotion";

const MyVideo = () => {
  return (
    <Freeze frame={30} active={(f) => f < 30}>
      <BlueSquare />
    </Freeze>
  );
};
```

```

From frame 30 on
tsx

import { Freeze } from "remotion";

const MyVideo = () => {
  return (
    <Freeze frame={30} active={(f) => f < 30}>
      <BlueSquare />
    </Freeze>
  );
};
```

## Demo [​](\#demo "Direct link to Demo")

0

0:00 / 0:05

Use Freeze component

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/freeze.tsx)
- [`<Video/>` Playback speed](/docs/video#playbackrate)
- [`useCurrentFrame()`](/docs/use-current-frame)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/freeze.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Folder>](/docs/folder) [Next\
\
getInputProps()](/docs/get-input-props)

- [Example](#example)
- [API](#api)
  - [`frame`](#frame)
  - [`active`](#active)
- [Demo](#demo)
- [See also](#see-also)
