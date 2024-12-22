---
title: <AbsoluteFill>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <AbsoluteFill>

- [Home page](/)
- [remotion](/docs/remotion)
- <AbsoluteFill>

On this page

# <AbsoluteFill>

A helper component - it is an absolutely positioned `<div>` with the following styles:

```

Styles of AbsoluteFill
ts

const style: React.CSSProperties = {
  position: "absolute",
  top: 0,
  left: 0,
  right: 0,
  bottom: 0,
  width: "100%",
  height: "100%",
  display: "flex",
  flexDirection: "column",
};
```

```

Styles of AbsoluteFill
ts

const style: React.CSSProperties = {
  position: "absolute",
  top: 0,
  left: 0,
  right: 0,
  bottom: 0,
  width: "100%",
  height: "100%",
  display: "flex",
  flexDirection: "column",
};
```

This component is useful for layering content on top of each other. For example, you can use it to create a full-screen video background:

```

Layer example
tsx

import { AbsoluteFill, OffthreadVideo } from "remotion";

const MyComp = () => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo src="https://example.com/video.mp4" />
      </AbsoluteFill>
      <AbsoluteFill>
        <h1>This text is written on top!</h1>
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

```

Layer example
tsx

import { AbsoluteFill, OffthreadVideo } from "remotion";

const MyComp = () => {
  return (
    <AbsoluteFill>
      <AbsoluteFill>
        <OffthreadVideo src="https://example.com/video.mp4" />
      </AbsoluteFill>
      <AbsoluteFill>
        <h1>This text is written on top!</h1>
      </AbsoluteFill>
    </AbsoluteFill>
  );
};
```

The layers that get rendered last appear on top - this is because of how HTML works.

## Adding a ref [​](\#adding-a-ref "Direct link to Adding a ref")

You can add a [React ref](https://react.dev/learn/manipulating-the-dom-with-refs) to an `<AbsoluteFill>` from version `v3.2.13` on. If you use TypeScript, you need to type it with `HTMLDivElement`:

```

tsx

const MyComp = () => {
  const ref = useRef<HTMLDivElement>(null);
  return <AbsoluteFill ref={ref}>{content}</AbsoluteFill>;
};
```

```

tsx

const MyComp = () => {
  const ref = useRef<HTMLDivElement>(null);
  return <AbsoluteFill ref={ref}>{content}</AbsoluteFill>;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/AbsoluteFill.tsx)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/absolute-fill.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
remotion](/docs/remotion) [Next\
\
<Audio>](/docs/audio)

- [Adding a ref](#adding-a-ref)
- [See also](#see-also)
