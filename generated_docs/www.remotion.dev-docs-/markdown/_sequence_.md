---
title: <Sequence>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Sequence>

- [Home page](/)
- [remotion](/docs/remotion)
- <Sequence>

On this page

# <Sequence>

By using a sequence, you can time-shift the display of your components or parts of your animation in the video.

```

MyTrailer.tsx
tsx

const MyTrailer = () => {
  return (
    <>
      <Sequence durationInFrames={30}>
        <Intro />
      </Sequence>
      <Sequence from={30} durationInFrames={30}>
        <Clip />
      </Sequence>
      <Sequence from={60}>
        <Outro />
      </Sequence>
    </>
  );
};
```

```

MyTrailer.tsx
tsx

const MyTrailer = () => {
  return (
    <>
      <Sequence durationInFrames={30}>
        <Intro />
      </Sequence>
      <Sequence from={30} durationInFrames={30}>
        <Clip />
      </Sequence>
      <Sequence from={60}>
        <Outro />
      </Sequence>
    </>
  );
};
```

_`<Intro>`_
will show from frame 0-29.

_`<Clip>`_

will show from frame 30 on until frame 59.

_`<Outro>`_

will show from frame 60 on until the end of the composition.

All children of a `<Sequence>` that call [`useCurrentFrame()`](/docs/use-current-frame) will receive a value that is shifted by [`from`](#from).

```

MyTrailer.tsx
tsx

import { Sequence, useCurrentFrame } from "remotion";

const Intro = () => <div>{useCurrentFrame()}</div>;

const MyTrailer = () => {
  return (
    <>
      <Intro />
      <Sequence from={30}>
        <Intro />
      </Sequence>
    </>
  );
};
```

```

MyTrailer.tsx
tsx

import { Sequence, useCurrentFrame } from "remotion";

const Intro = () => <div>{useCurrentFrame()}</div>;

const MyTrailer = () => {
  return (
    <>
      <Intro />
      <Sequence from={30}>
        <Intro />
      </Sequence>
    </>
  );
};
```

At frame `0`, this would render `<div>0</div>`.

At frame `30`, this would render `<div>30</div><div>0</div>`.

Using the [`durationInFrames`](#durationinframes) prop, you can define for how long the children of a `<Sequence>` should be mounted.

By default, the children of a `<Sequence>` are wrapped in an [`<AbsoluteFill>`](/docs/absolute-fill) component. If you don't want this behavior, add [`layout="none"`](#layout) as a prop.

## Cascading [​](\#cascading "Direct link to Cascading")

You can nest sequences within each other and they will cascade.

For example, a sequence that starts at frame 60 which is inside a sequence that starts at frame 30 will have it's children start at frame 90.

## Examples [​](\#examples "Direct link to Examples")

All the examples below are based on the following animation of a blue square:

0

0:00 / 0:05

```

MyVideo.tsx
tsx

const MyVideo = () => {
  return <BlueSquare />;
};
```

```

MyVideo.tsx
tsx

const MyVideo = () => {
  return <BlueSquare />;
};
```

### Delay [​](\#delay "Direct link to Delay")

If you would like to delay the content by say 30 frames, you can wrap it in

`<Sequence from={30}>`.

0:00 / 0:05

```

delay.tsx
tsx

const MyVideo = () => {
  return (
    <Sequence from={30}>
      <BlueSquare />
    </Sequence>
  );
};
```

```

delay.tsx
tsx

const MyVideo = () => {
  return (
    <Sequence from={30}>
      <BlueSquare />
    </Sequence>
  );
};
```

### Trim end [​](\#trim-end "Direct link to Trim end")

Wrap your component in a `<Sequence>` with a finite `durationInFrames` prop to make it unmount after the duration has passed.

0

0:00 / 0:05

```

trim-end.tsx
tsx

const ClipExample: React.FC = () => {
  return (
    <Sequence durationInFrames={45}>
      <BlueSquare />
    </Sequence>
  );
};
```

```

trim-end.tsx
tsx

const ClipExample: React.FC = () => {
  return (
    <Sequence durationInFrames={45}>
      <BlueSquare />
    </Sequence>
  );
};
```

### Trim start [​](\#trim-start "Direct link to Trim start")

Wrap the square in `<Sequence>` with a negative `from` value to trim the beginning of the content.

By shifting the time backwards, the animation has already progressed by 15 frames when the content appears.

15

0:00 / 0:05

```

trim-start.tsx
tsx

const TrimStartExample: React.FC = () => {
  return (
    <Sequence from={-15}>
      <BlueSquare />
    </Sequence>
  );
};
```

```

trim-start.tsx
tsx

const TrimStartExample: React.FC = () => {
  return (
    <Sequence from={-15}>
      <BlueSquare />
    </Sequence>
  );
};
```

### Trim and delay [​](\#trim-and-delay "Direct link to Trim and delay")

Wrap the content in two `<Sequence>`'s.

To the inner one, pass a negative start value `from={-15}` to trim away the first 15 frames of the content.

To the outer one we pass a positive value `from={30}` to then shift it forwards by 30 frames.

0:00 / 0:05

```

trim-and-delay.tsx
tsx

const TrimAndDelayExample: React.FC = () => {
  return (
    <Sequence from={30}>
      <Sequence from={-15}>
        <BlueSquare />
      </Sequence>
    </Sequence>
  );
};
```

```

trim-and-delay.tsx
tsx

const TrimAndDelayExample: React.FC = () => {
  return (
    <Sequence from={30}>
      <Sequence from={-15}>
        <BlueSquare />
      </Sequence>
    </Sequence>
  );
};
```

## Play Sequences sequentially [​](\#play-sequences-sequentially "Direct link to Play Sequences sequentially")

See the [`<Series />`](/docs/series) helper component, which helps you calculate markup that makes sequences play after each other.

## Props [​](\#props "Direct link to Props")

The Sequence component is a high order component and accepts, besides children, the following props:

### `from` [​](\#from "Direct link to from")

_optional_ (From v3.2.36, _required_ in previous versions)

At which frame it's children should assume the video starts. When the sequence is at `frame`, it's children are at frame `0`.
From v3.2.36 onwards, this prop will be optional; by default, it will be 0.

### `durationInFrames` [​](\#durationinframes "Direct link to durationinframes")

_optional_

For how many frames the sequence should be displayed. Children are unmounted if they are not within the time range of display. By default it will be `Infinity` to avoid limit the duration of the sequence.

### `height` [v4.0.80](https://github.com/remotion-dev/remotion/releases/v4.0.80) [​](\#height "Direct link to height")

_optional_

Gives the sequence a specific `style={{height: height}}` style and overrides `height` that is returned by the [`useVideoConfig()`](/docs/use-video-config) hook in child components. Useful for including a component that was designed for a specific height.

### `width` [v4.0.80](https://github.com/remotion-dev/remotion/releases/v4.0.80) [​](\#width "Direct link to width")

_optional_

Gives the sequence a specific `style={{width: width}}` style and overrides `width` that is returned by the [`useVideoConfig()`](/docs/use-video-config) hook in child components. Useful for including a component that was designed for a specific width.

### `name` [​](\#name "Direct link to name")

_optional_

You can give your sequence a name and it will be shown as the label of the sequence in the timeline of the Remotion Studio. This property is purely for helping you keep track of sequences in the timeline.

### `layout` [​](\#layout "Direct link to layout")

_optional_

Either `"absolute-fill"` _(default)_ or `"none"`. By default, your sequences will be absolutely positioned, so they will overlay each other. If you would like to opt out of it and handle layouting yourself, pass `layout="none"`. Available since v1.4.

### `style` [v3.0.27](https://github.com/remotion-dev/remotion/releases/v3.0.27) [​](\#style "Direct link to style")

_optional_

CSS styles to be applied to the container. If `layout` is set to `none`, there is no container and setting this style is not allowed.

### `className` [v3.3.45](https://github.com/remotion-dev/remotion/releases/v3.3.45) [​](\#classname "Direct link to classname")

_optional_

A class name to be applied to the container. If `layout` is set to `none`, there is no container and setting this style is not allowed.

### `premountFor` [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#premountfor "Direct link to premountfor")

_optional_

[Premount](/docs/player/premounting) the sequence for a set number of frames.

### `showInTimeline` [v4.0.110](https://github.com/remotion-dev/remotion/releases/v4.0.110) [​](\#showintimeline "Direct link to showintimeline")

_optional_

If set to `false`, the track will not be shown in the Studio's timeline.

Child `<Sequence>`'s will show by default, unless `showInTimeline` is also set to false.

This behavior is stable as of v4.0.110, previously the behavior was different, but this prop not documented.

## Adding a ref [​](\#adding-a-ref "Direct link to Adding a ref")

You can add a [React ref](https://react.dev/learn/manipulating-the-dom-with-refs) to an `<Sequence>` from version `v3.2.13` on. If you use TypeScript, you need to type it with `HTMLDivElement`:

```

tsx

const MyComp = () => {
  const ref = useRef<HTMLDivElement>(null);
  return (
    <Sequence from={10} ref={ref}>
      {content}
    </Sequence>
  );
};
```

```

tsx

const MyComp = () => {
  const ref = useRef<HTMLDivElement>(null);
  return (
    <Sequence from={10} ref={ref}>
      {content}
    </Sequence>
  );
};
```

## Note for `@remotion/three` [​](\#note-for-remotionthree "Direct link to note-for-remotionthree")

A `<Sequence>` by default will return a `<div>` component which is not allowed inside a [`<ThreeCanvas>`](/docs/three-canvas). Avoid an error by passing `layout="none"` to `<Sequence>`. Example: `<Sequence layout="none">`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Sequence.tsx)
- [Reuse components using Sequences](/docs/reusability)
- [`<Composition />`](/docs/composition)
- [`<Series />`](/docs/series)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/sequence.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
registerRoot()](/docs/register-root) [Next\
\
<Series>](/docs/series)

- [Cascading](#cascading)
- [Examples](#examples)
  - [Delay](#delay)
  - [Trim end](#trim-end)
  - [Trim start](#trim-start)
  - [Trim and delay](#trim-and-delay)
- [Play Sequences sequentially](#play-sequences-sequentially)
- [Props](#props)
  - [`from`](#from)
  - [`durationInFrames`](#durationinframes)
  - [`height`](#height)
  - [`width`](#width)
  - [`name`](#name)
  - [`layout`](#layout)
  - [`style`](#style)
  - [`className`](#classname)
  - [`premountFor`](#premountfor)
  - [`showInTimeline`](#showintimeline)
- [Adding a ref](#adding-a-ref)
- [Note for `@remotion/three`](#note-for-remotionthree)
- [See also](#see-also)
