---
title: Custom presentations
source: Remotion Documentation
last_updated: 2024-12-22
---

# Custom presentations

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- Custom presentations

On this page

# Custom presentations

This page describes how to create your own custom effect for [`<TransitionSeries>`](/docs/transitions/transitionseries).

## Concept [​](\#concept "Direct link to Concept")

A presentation is a higher order component which wraps around the entering slide as well as the exiting slide and which implements an effect based on three parameters:

[1](#1)

The current `presentationProgress`, influenced by the
timing of the transition

[2](#2)

The `presentationDirection`, either `entering` or `exiting`

[3](#3)

The `presentationDurationInFrames` ( _available from v4.0.153_)
[4](#4) Additional developer-defined `passedProps`

## Boilerplate [​](\#boilerplate "Direct link to Boilerplate")

A custom presentation is a function which returns an object of type `TransitionPresentation`.

It is an object with a React component ( `component`) and React props which are passed to the component as `passedProps`.

```

custom-presentation.tsx
tsx

import type {TransitionPresentation} from '@remotion/transitions';

type CustomPresentationProps = {
  width: number;
  height: number;
};

export const customPresentation = (
  props: CustomPresentationProps,
): TransitionPresentation<CustomPresentationProps> => {
  return {component: StarPresentation, props};
};
```

```

custom-presentation.tsx
tsx

import type {TransitionPresentation} from '@remotion/transitions';

type CustomPresentationProps = {
  width: number;
  height: number;
};

export const customPresentation = (
  props: CustomPresentationProps,
): TransitionPresentation<CustomPresentationProps> => {
  return {component: StarPresentation, props};
};
```

The `component` is a React component which receives the following props:

[1](#1)

`children`: The markup to wrap around

[2](#2)

`presentationDirection`: Either `
"entering"
` or `"exiting"`

[3](#3)

`presentationProgress`: A number between `
0
` and `1` which represents the progress of the transition.

[4](#4)

`passedProps`: The custom props passed to the presentation

```

StarPresentation.tsx
tsx

import type {TransitionPresentationComponentProps} from '@remotion/transitions';
import {AbsoluteFill} from 'remotion';

const StarPresentation: React.FC<
  TransitionPresentationComponentProps<CustomPresentationProps>
> = ({children, presentationDirection, presentationProgress, passedProps}) => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>{children}</AbsoluteFill>
    </AbsoluteFill>
  );
};
```

```

StarPresentation.tsx
tsx

import type {TransitionPresentationComponentProps} from '@remotion/transitions';
import {AbsoluteFill} from 'remotion';

const StarPresentation: React.FC<
  TransitionPresentationComponentProps<CustomPresentationProps>
> = ({children, presentationDirection, presentationProgress, passedProps}) => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>{children}</AbsoluteFill>
    </AbsoluteFill>
  );
};
```

## Example [​](\#example "Direct link to Example")

A

B

The following example implements a star mask transition:

[1](#1)

Based on the `passedProps` `height` and `width`
, the inner radius of the star is calculated that will completely fill the canvas.

[2](#2)

Using [`@remotion/shapes`](/docs/shapes), an SVG path is calculated, and grown from zero to the full size of the canvas.

[3](#3) The `presentationProgress` is used to interpolate the shape
size.

[4](#4) [`@remotion/paths`](/docs/paths) is used to center the star in the middle of the canvas.

[5](#5)

A `clipPath` is used to clip the entering slide.
Inside the container, the `children` get rendered.

[6](#6)

The effect is disabled if `presentationDirection` is set
to `"exiting"`

```

StarPresentation.tsx
tsx

import {getBoundingBox, translatePath} from '@remotion/paths';
import {makeStar} from '@remotion/shapes';
import type {TransitionPresentationComponentProps} from '@remotion/transitions';
import React, {useMemo, useState} from 'react';
import {AbsoluteFill, random} from 'remotion';

export type CustomPresentationProps = {
  width: number;
  height: number;
};

const StarPresentation: React.FC<
  TransitionPresentationComponentProps<CustomPresentationProps>
> = ({children, presentationDirection, presentationProgress, passedProps}) => {
  const finishedRadius =
    Math.sqrt(passedProps.width ** 2 + passedProps.height ** 2) / 2;
  const innerRadius = finishedRadius * presentationProgress;
  const outerRadius = finishedRadius * 2 * presentationProgress;

  const {path} = makeStar({
    innerRadius,
    outerRadius,
    points: 5,
  });

  const boundingBox = getBoundingBox(path);

  const translatedPath = translatePath(
    path,
    passedProps.width / 2 - boundingBox.width / 2,
    passedProps.height / 2 - boundingBox.height / 2,
  );

  const [clipId] = useState(() => String(random(null)));

  const style: React.CSSProperties = useMemo(() => {
    return {
      width: '100%',
      height: '100%',
      clipPath:
        presentationDirection === 'exiting' ? undefined : `url(#${clipId})`,
    };
  }, [clipId, presentationDirection]);

  return (
    <AbsoluteFill>
      <AbsoluteFill style={style}>{children}</AbsoluteFill>
      {presentationDirection === 'exiting' ? null : (
        <AbsoluteFill>
          <svg>
            <defs>
              <clipPath id={clipId}>
                <path d={translatedPath} fill="black" />
              </clipPath>
            </defs>
          </svg>
        </AbsoluteFill>
      )}
    </AbsoluteFill>
  );
};
```

```

StarPresentation.tsx
tsx

import {getBoundingBox, translatePath} from '@remotion/paths';
import {makeStar} from '@remotion/shapes';
import type {TransitionPresentationComponentProps} from '@remotion/transitions';
import React, {useMemo, useState} from 'react';
import {AbsoluteFill, random} from 'remotion';

export type CustomPresentationProps = {
  width: number;
  height: number;
};

const StarPresentation: React.FC<
  TransitionPresentationComponentProps<CustomPresentationProps>
> = ({children, presentationDirection, presentationProgress, passedProps}) => {
  const finishedRadius =
    Math.sqrt(passedProps.width ** 2 + passedProps.height ** 2) / 2;
  const innerRadius = finishedRadius * presentationProgress;
  const outerRadius = finishedRadius * 2 * presentationProgress;

  const {path} = makeStar({
    innerRadius,
    outerRadius,
    points: 5,
  });

  const boundingBox = getBoundingBox(path);

  const translatedPath = translatePath(
    path,
    passedProps.width / 2 - boundingBox.width / 2,
    passedProps.height / 2 - boundingBox.height / 2,
  );

  const [clipId] = useState(() => String(random(null)));

  const style: React.CSSProperties = useMemo(() => {
    return {
      width: '100%',
      height: '100%',
      clipPath:
        presentationDirection === 'exiting' ? undefined : `url(#${clipId})`,
    };
  }, [clipId, presentationDirection]);

  return (
    <AbsoluteFill>
      <AbsoluteFill style={style}>{children}</AbsoluteFill>
      {presentationDirection === 'exiting' ? null : (
        <AbsoluteFill>
          <svg>
            <defs>
              <clipPath id={clipId}>
                <path d={translatedPath} fill="black" />
              </clipPath>
            </defs>
          </svg>
        </AbsoluteFill>
      )}
    </AbsoluteFill>
  );
};
```

Example usage:

```

MyComp.tsx
tsx

export const MyComp: React.FC = () => {
  const {width, height} = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={70}>
        <Letter color="orange">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={customPresentation({width, height})}
        timing={springTiming({
          durationInFrames: 45,
          config: {
            damping: 200,
          },
          durationRestThreshold: 0.0001,
        })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

```

MyComp.tsx
tsx

export const MyComp: React.FC = () => {
  const {width, height} = useVideoConfig();

  return (
    <TransitionSeries>
      <TransitionSeries.Sequence durationInFrames={70}>
        <Letter color="orange">A</Letter>
      </TransitionSeries.Sequence>
      <TransitionSeries.Transition
        presentation={customPresentation({width, height})}
        timing={springTiming({
          durationInFrames: 45,
          config: {
            damping: 200,
          },
          durationRestThreshold: 0.0001,
        })}
      />
      <TransitionSeries.Sequence durationInFrames={60}>
        <Letter color="pink">B</Letter>
      </TransitionSeries.Sequence>
    </TransitionSeries>
  );
};
```

## References [​](\#references "Direct link to References")

See the source code for [already implemented presentations](https://github.com/remotion-dev/remotion/blob/main/packages/transitions/src/presentations) for a useful reference.

## See also [​](\#see-also "Direct link to See also")

- [Custom timings](/docs/transitions/timings/custom)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/presentations/custom.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
cube()](/docs/transitions/presentations/cube) [Next\
\
Audio transitions](/docs/transitions/audio-transitions)

- [Concept](#concept)
- [Boilerplate](#boilerplate)
- [Example](#example)
- [References](#references)
- [See also](#see-also)
