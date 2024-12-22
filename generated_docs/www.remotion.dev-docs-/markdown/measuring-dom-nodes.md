---
title: Measuring DOM nodes
source: Remotion Documentation
last_updated: 2024-12-22
---

# Measuring DOM nodes

- [Home page](/)
- Designing visuals
- Measuring items

On this page

# Measuring DOM nodes

If you want to measure a DOM node, you can assign a [React Ref](https://react.dev/learn/manipulating-the-dom-with-refs) to it and then use the [`getBoundingClientRect()`](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect?retiredLocale=de) browser API to get the position and size of the node.

In Remotion, it is not that easy because the `<div>` element in which the video is rendered into has a `scale()` transform applied to it, which affects the value that you get from `getBoundingClientRect()`.

## Measure using the `useCurrentScale()` hook [​](\#measure-using-the-usecurrentscale-hook "Direct link to measure-using-the-usecurrentscale-hook")

From v4.0.111 on, you can use the [`useCurrentScale()`](/docs/use-current-scale) hook to correct the dimensions of the element.

```

MyComponent.tsx
tsx

import { useCallback, useEffect, useState, useRef } from "react";
import { useCurrentScale } from "remotion";

export const MyComponent = () => {
  const ref = useRef<HTMLDivElement>(null);

  const [dimensions, setDimensions] = useState<{
    correctedHeight: number;
    correctedWidth: number;
  } | null>(null);

  const scale = useCurrentScale();

  useEffect(() => {
    if (!ref.current) {
      return;
    }

    const rect = ref.current.getBoundingClientRect();

    setDimensions({
      correctedHeight: rect.height / scale,
      correctedWidth: rect.width / scale,
    });
  }, [scale]);

  return (
    <div>
      <div ref={ref}>Hello World!</div>
    </div>
  );
};
```

```

MyComponent.tsx
tsx

import { useCallback, useEffect, useState, useRef } from "react";
import { useCurrentScale } from "remotion";

export const MyComponent = () => {
  const ref = useRef<HTMLDivElement>(null);

  const [dimensions, setDimensions] = useState<{
    correctedHeight: number;
    correctedWidth: number;
  } | null>(null);

  const scale = useCurrentScale();

  useEffect(() => {
    if (!ref.current) {
      return;
    }

    const rect = ref.current.getBoundingClientRect();

    setDimensions({
      correctedHeight: rect.height / scale,
      correctedWidth: rect.width / scale,
    });
  }, [scale]);

  return (
    <div>
      <div ref={ref}>Hello World!</div>
    </div>
  );
};
```

## Versions prior to v4.0.110 [​](\#versions-prior-to-v40110 "Direct link to Versions prior to v4.0.110")

To get an accurate measurement, you can render an additional element with a fixed width (say `10px`) and measure it too. Then, you can divide the width of the element by `10` to get the scale factor.

```

MyComponent.tsx
tsx

import { useCallback, useEffect, useState, useRef } from "react";

const MEASURER_SIZE = 10;

export const MyComponent = () => {
  const ref = useRef<HTMLDivElement>(null);
  const measurer = useRef<HTMLDivElement>(null);

  const [dimensions, setDimensions] = useState<{
    correctedHeight: number;
    correctedWidth: number;
  } | null>(null);

  useEffect(() => {
    if (!ref.current || !measurer.current) {
      return;
    }

    const rect = ref.current.getBoundingClientRect();
    const measurerRect = measurer.current.getBoundingClientRect();
    const scale = measurerRect.width / MEASURER_SIZE;

    setDimensions({
      correctedHeight: rect.height * scale,
      correctedWidth: rect.width * scale,
    });
  }, []);

  return (
    <div>
      <div ref={ref}>Hello World!</div>
      <div
        ref={measurer}
        style={{
          width: MEASURER_SIZE,
          position: "fixed",
          top: -99999,
        }}
      />
    </div>
  );
};
```

```

MyComponent.tsx
tsx

import { useCallback, useEffect, useState, useRef } from "react";

const MEASURER_SIZE = 10;

export const MyComponent = () => {
  const ref = useRef<HTMLDivElement>(null);
  const measurer = useRef<HTMLDivElement>(null);

  const [dimensions, setDimensions] = useState<{
    correctedHeight: number;
    correctedWidth: number;
  } | null>(null);

  useEffect(() => {
    if (!ref.current || !measurer.current) {
      return;
    }

    const rect = ref.current.getBoundingClientRect();
    const measurerRect = measurer.current.getBoundingClientRect();
    const scale = measurerRect.width / MEASURER_SIZE;

    setDimensions({
      correctedHeight: rect.height * scale,
      correctedWidth: rect.width * scale,
    });
  }, []);

  return (
    <div>
      <div ref={ref}>Hello World!</div>
      <div
        ref={measurer}
        style={{
          width: MEASURER_SIZE,
          position: "fixed",
          top: -99999,
        }}
      />
    </div>
  );
};
```

### Example project [​](\#example-project "Direct link to Example project")

- [Source code](https://github.com/remotion-dev/measure-item)
- [Preview](https://measure-item.vercel.app)

## Versions prior to v4.0.103 [​](\#versions-prior-to-v40103 "Direct link to Versions prior to v4.0.103")

In previous versions of Remotion, `getBoundingClientRect()` could return dimensions with all values being `0` in the first `useEffect()` due to mounting your component but not showing it.

Going forward, you can rely on the dimensions being non-zero.

## See also [​](\#see-also "Direct link to See also")

- [react-use-measure](https://github.com/pmndrs/react-use-measure) \- Measure elements correctly inside a scroll container

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/measuring.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Fonts](/docs/fonts) [Next\
\
Randomness](/docs/using-randomness)

- [Measure using the `useCurrentScale()` hook](#measure-using-the-usecurrentscale-hook)
- [Versions prior to v4.0.110](#versions-prior-to-v40110)
  - [Example project](#example-project)
- [Versions prior to v4.0.103](#versions-prior-to-v40103)
- [See also](#see-also)
