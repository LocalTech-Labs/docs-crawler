---
title: Transitioning with audio
source: Remotion Documentation
last_updated: 2024-12-22
---

# Transitioning with audio

- [Home page](/)
- [@remotion/transitions](/docs/transitions/)
- [Presentations](/docs/transitions/presentations/)
- Audio transitions

On this page

# Transitioning with audio

_working from v4.0.58_

To add sound to a transition, you may use this function to wrap any [presentation](/docs/transitions/presentations/):

```

add-sound.tsx
tsx

import {
  TransitionPresentation,
  TransitionPresentationComponentProps,
} from "@remotion/transitions";
import { Audio } from "remotion";

export function addSound<T extends Record<string, unknown>>(
  transition: TransitionPresentation<T>,
  src: string,
): TransitionPresentation<T> {
  const { component: Component, ...other } = transition;

  const C = Component as React.FC<TransitionPresentationComponentProps<T>>;

  const NewComponent: React.FC<TransitionPresentationComponentProps<T>> = (
    p,
  ) => {
    return (
      <>
        {p.presentationDirection === "entering" ? <Audio src={src} /> : null}
        <C {...p} />
      </>
    );
  };

  return {
    component: NewComponent,
    ...other,
  };
}
```

```

add-sound.tsx
tsx

import {
  TransitionPresentation,
  TransitionPresentationComponentProps,
} from "@remotion/transitions";
import { Audio } from "remotion";

export function addSound<T extends Record<string, unknown>>(
  transition: TransitionPresentation<T>,
  src: string,
): TransitionPresentation<T> {
  const { component: Component, ...other } = transition;

  const C = Component as React.FC<TransitionPresentationComponentProps<T>>;

  const NewComponent: React.FC<TransitionPresentationComponentProps<T>> = (
    p,
  ) => {
    return (
      <>
        {p.presentationDirection === "entering" ? <Audio src={src} /> : null}
        <C {...p} />
      </>
    );
  };

  return {
    component: NewComponent,
    ...other,
  };
}
```

Customize the volume, offset, playback rate, and other properties of the [`<Audio>`](/docs/audio) component as you wish.

You may use it like this:

```

example.tsx
tsx

import { slide } from "@remotion/transitions/slide";
import { staticFile } from "remotion";

const presentation = slide();
const withSound = addSound(presentation, staticFile("whoosh.mp3"));
```

```

example.tsx
tsx

import { slide } from "@remotion/transitions/slide";
import { staticFile } from "remotion";

const presentation = slide();
const withSound = addSound(presentation, staticFile("whoosh.mp3"));
```

Now you can use `withSound` in place of `presentation`.

## See also [​](\#see-also "Direct link to See also")

- [`<Audio>`](/docs/audio)
- [Using audio](/docs/using-audio)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/transitions/audio-transitions.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Custom presentations](/docs/transitions/presentations/custom) [Next\
\
Overview](/docs/zod-types/)

- [See also](#see-also)
