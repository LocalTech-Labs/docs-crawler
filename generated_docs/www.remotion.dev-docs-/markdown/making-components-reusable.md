---
title: Making components reusable
source: Remotion Documentation
last_updated: 2024-12-22
---

# Making components reusable

- [Home page](/)
- Getting started
- Reuse components

On this page

# Making components reusable

React components allow us to encapsulate video logic and reuse the same visuals multiple times.

Consider a title - to make it reusable, factor it out into its own component:

```

MyComposition.tsx
tsx

import {AbsoluteFill, interpolate, useCurrentFrame} from 'remotion';

const Title: React.FC<{title: string}> = ({title}) => {
  const frame = useCurrentFrame();
  const opacity = interpolate(frame, [0, 20], [0, 1], {
    extrapolateRight: 'clamp',
  });

  return (
    <div style={{opacity, textAlign: 'center', fontSize: '7em'}}>{title}</div>
  );
};

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <Title title="Hello World" />
    </AbsoluteFill>
  );
};
```

```

MyComposition.tsx
tsx

import {AbsoluteFill, interpolate, useCurrentFrame} from 'remotion';

const Title: React.FC<{title: string}> = ({title}) => {
  const frame = useCurrentFrame();
  const opacity = interpolate(frame, [0, 20], [0, 1], {
    extrapolateRight: 'clamp',
  });

  return (
    <div style={{opacity, textAlign: 'center', fontSize: '7em'}}>{title}</div>
  );
};

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <Title title="Hello World" />
    </AbsoluteFill>
  );
};
```

To render multiple instances of the title, duplicate the `<Title>` component.

You can also use the [`<Sequence>`](/docs/sequence) component to limit the duration of the first title and time-shift the appearance of the second title.

```

tsx

import {Sequence} from 'remotion';

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <Sequence durationInFrames={40}>
        <Title title="Hello" />
      </Sequence>
      <Sequence from={40}>
        <Title title="World" />
      </Sequence>
    </AbsoluteFill>
  );
};
```

```

tsx

import {Sequence} from 'remotion';

export const MyVideo = () => {
  return (
    <AbsoluteFill>
      <Sequence durationInFrames={40}>
        <Title title="Hello" />
      </Sequence>
      <Sequence from={40}>
        <Title title="World" />
      </Sequence>
    </AbsoluteFill>
  );
};
```

You should see two titles appearing after each other.

Note that the value of [`useCurrentFrame()`](/docs/use-current-frame) has been shifted in the second instance, so that it returns `0` only when the absolute time is `40`. Before that, the sequence was not mounted at all.

Sequences by default are absolutely positioned - you can use [`layout="none"`](/docs/sequence#layout) to make them render like a regular `<div>`.

## See also [​](\#see-also "Direct link to See also")

- [`<Sequence>` component](/docs/sequence)
- [How do I combine compositions?](/docs/miscellaneous/snippets/combine-compositions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/sequences.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Animating properties](/docs/animating-properties) [Next\
\
Preview your video](/docs/preview)

- [See also](#see-also)