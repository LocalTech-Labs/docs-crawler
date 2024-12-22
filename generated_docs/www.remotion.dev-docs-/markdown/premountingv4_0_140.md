---
title: Premountingv4.0.140
source: Remotion Documentation
last_updated: 2024-12-22
---

# Premountingv4.0.140

- [Home page](/)
- [Player](/docs/player/)
- Premounting

On this page

# Premounting [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140)

Remotion only ever renders the current frame.

This means when a video or another asset is about to come up, it is not loaded by default.

note

Even though the video is soon about to appear, it is not yet loaded
because Remotion only ever renders the current time.

## What is premounting? [​](\#what-is-premounting "Direct link to What is premounting?")

Premounting is the practice of mounting a component containing assets earlier to allow assets some time to load before they appear.

note

The video is mounted earlier to give it some time to load.

It carries the style `opacity: 0` to make it invisible.
It's time defined by [`useCurrentFrame()`](/docs/use-current-frame) is frozen at `0`.

## Preloading a `<Sequence>` [​](\#preloading-a-sequence "Direct link to preloading-a-sequence")

Use premounting sparingly since having more elements mounted will slow down the page.

Add the [`premountFor`](/docs/sequence#premountfor) prop to the [`<Sequence>`](/docs/sequence) component to enable premounting.

The number you pass is the number in frames you premount the component for.

```

Premounted.tsx
tsx

const MyComp: React.FC = () => {
  return (
    <Sequence premountFor={100}>
      <Video src={staticFile('bigbuckbunny.mp4')}></Video>
    </Sequence>
  );
};
```

```

Premounted.tsx
tsx

const MyComp: React.FC = () => {
  return (
    <Sequence premountFor={100}>
      <Video src={staticFile('bigbuckbunny.mp4')}></Video>
    </Sequence>
  );
};
```

In the [Remotion Studio](/docs/studio), a premounted component is indicated with diagonal stripes:

![](/img/premount.png)

A premounted [`<Sequence>`](/docs/sequence) does carry the styles `opacity: 0` and `pointer-events: none`.

It is not possible to premount a [`<Sequence>`](/docs/sequence) with `layout="none"` because such a sequence does not have a container to apply the styles to.

## With `<Series>` and `<TransitionSeries>` [​](\#with-series-and-transitionseries "Direct link to with-series-and-transitionseries")

`premountFor` is also available on [`<Series.Sequence>`](/docs/series) and [`<TransitionSeries.Sequence>`](/docs/transitions/transitionseries).

It is not yet possible to premount a whole [`<Series>`](/docs/series) or [`<TransitionSeries>`](/docs/transitions/transitionseries).

## Custom premount component [​](\#custom-premount-component "Direct link to Custom premount component")

Premounting can also be implemented yourself with the public Remotion APIs.

Use the following component like you would `<Sequence>`:

```

PremountedSequence.tsx
tsx

import React, {forwardRef, useMemo} from 'react';
import {
  Freeze,
  getRemotionEnvironment,
  Sequence,
  SequenceProps,
  useCurrentFrame,
} from 'remotion';

export type PremountedSequenceProps = SequenceProps & {
  premountFor: number;
};

const PremountedSequenceRefForwardingFunction: React.ForwardRefRenderFunction<
  HTMLDivElement,
  {
    premountFor: number;
  } & SequenceProps
> = ({premountFor, ...props}, ref) => {
  const frame = useCurrentFrame();

  if (props.layout === 'none') {
    throw new Error('`<Premount>` does not support layout="none"');
  }

  const {style: passedStyle, from = 0, ...otherProps} = props;
  const active =
    frame < from &&
    frame >= from - premountFor &&
    !getRemotionEnvironment().isRendering;

  const style: React.CSSProperties = useMemo(() => {
    return {
      ...passedStyle,
      opacity: active ? 0 : 1,
      // @ts-expect-error Only in the docs - it will not give a type error in a Remotion project
      pointerEvents: active ? 'none' : (passedStyle?.pointerEvents ?? 'auto'),
    };
  }, [active, passedStyle]);

  return (
    <Freeze frame={from} active={active}>
      <Sequence
        ref={ref}
        name={`<PremountedSequence premountFor={${premountFor}}>`}
        from={from}
        style={style}
        {...otherProps}
      />
    </Freeze>
  );
};

export const PremountedSequence = forwardRef(
  PremountedSequenceRefForwardingFunction,
);
```

```

PremountedSequence.tsx
tsx

import React, {forwardRef, useMemo} from 'react';
import {
  Freeze,
  getRemotionEnvironment,
  Sequence,
  SequenceProps,
  useCurrentFrame,
} from 'remotion';

export type PremountedSequenceProps = SequenceProps & {
  premountFor: number;
};

const PremountedSequenceRefForwardingFunction: React.ForwardRefRenderFunction<
  HTMLDivElement,
  {
    premountFor: number;
  } & SequenceProps
> = ({premountFor, ...props}, ref) => {
  const frame = useCurrentFrame();

  if (props.layout === 'none') {
    throw new Error('`<Premount>` does not support layout="none"');
  }

  const {style: passedStyle, from = 0, ...otherProps} = props;
  const active =
    frame < from &&
    frame >= from - premountFor &&
    !getRemotionEnvironment().isRendering;

  const style: React.CSSProperties = useMemo(() => {
    return {
      ...passedStyle,
      opacity: active ? 0 : 1,
      // @ts-expect-error Only in the docs - it will not give a type error in a Remotion project
      pointerEvents: active ? 'none' : (passedStyle?.pointerEvents ?? 'auto'),
    };
  }, [active, passedStyle]);

  return (
    <Freeze frame={from} active={active}>
      <Sequence
        ref={ref}
        name={`<PremountedSequence premountFor={${premountFor}}>`}
        from={from}
        style={style}
        {...otherProps}
      />
    </Freeze>
  );
};

export const PremountedSequence = forwardRef(
  PremountedSequenceRefForwardingFunction,
);
```

note

One caveat of this custom component: In a premounted sequence, the native buffer state can still be triggered.

It is your task to disable triggering the native buffer state (e.g set [`pauseWhenBuffering`](/docs/offthreadvideo#pausewhenbuffering) to false) when in a premounted sequence.

You can achieve this for example by using a React context.

## Usage together with the buffer state [​](\#usage-together-with-the-buffer-state "Direct link to Usage together with the buffer state")

If you also use the [buffer state](/docs/player/buffer-state), the [`<Audio>`](/docs/audio#pausewhenbuffering), [`<Video>`](/docs/video#pausewhenbuffering), [`<OffthreadVideo>`](/docs/offthreadvideo#pausewhenbuffering) and [`<Img>`](/docs/img#pausewhenloading) tags are aware of premounting and don't trigger the buffer state while in a premounted [`<Sequence>`](/docs/sequence).

Otherwise it would lead to pausing playback while a scene is not even visible yet.

## See also [​](\#see-also "Direct link to See also")

- [Buffer state](/docs/player/buffer-state)
- [Preloading](/docs/player/preloading)
- [Avoiding flickering](/docs/troubleshooting/player-flicker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/premounting.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Preloading](/docs/player/preloading) [Next\
\
Best practices](/docs/player/best-practices)

- [What is premounting?](#what-is-premounting)
- [Preloading a `<Sequence>`](#preloading-a-sequence)
- [With `<Series>` and `<TransitionSeries>`](#with-series-and-transitionseries)
- [Custom premount component](#custom-premount-component)
- [Usage together with the buffer state](#usage-together-with-the-buffer-state)
- [See also](#see-also)