---
title: registerRoot()
source: Remotion Documentation
last_updated: 2024-12-22
---

# registerRoot()

- [Home page](/)
- [remotion](/docs/remotion)
- registerRoot()

On this page

# registerRoot()

`registerRoot` is a function that registers the root component of the Remotion project. In the root component, one or multiple compositions should be returned (in the case of multiple compositions, they should be wrapped in a React Fragment).

info

`registerRoot()` should live in a file that is separate from the list of compositions. This is because when using React Fast Refresh, all the code in the file that gets refreshed gets executed again, however, this function should only be called once.

## Example [​](\#example "Direct link to Example")

```

src/index.ts
tsx

import { registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

registerRoot(RemotionRoot);
```

```

src/index.ts
tsx

import { registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

registerRoot(RemotionRoot);
```

```

src/Root.tsx
tsx

import MyComponent from "./MyComponent";
import MyOtherComponent from "./MyOtherComponent";

export const RemotionRoot = () => {
  return (
    <>
      <Composition
        id="comp"
        fps={30}
        height={1080}
        width={1920}
        durationInFrames={90}
        component={MyComponent}
      />
      <Composition
        id="anothercomp"
        fps={30}
        height={1080}
        width={1920}
        durationInFrames={90}
        component={MyOtherComponent}
      />
    </>
  );
};
```

```

src/Root.tsx
tsx

import MyComponent from "./MyComponent";
import MyOtherComponent from "./MyOtherComponent";

export const RemotionRoot = () => {
  return (
    <>
      <Composition
        id="comp"
        fps={30}
        height={1080}
        width={1920}
        durationInFrames={90}
        component={MyComponent}
      />
      <Composition
        id="anothercomp"
        fps={30}
        height={1080}
        width={1920}
        durationInFrames={90}
        component={MyOtherComponent}
      />
    </>
  );
};
```

## Defer registerRoot() [​](\#defer-registerroot "Direct link to Defer registerRoot()")

In some cases, such as dynamically importing roots or loading WebAssembly, you might want to defer the loading of registerRoot(). In newer versions of Remotion, you may do so, without having to invoke `delayRender()`.

```

tsx

import { continueRender, delayRender, registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

loadWebAssembly().then(() => {
  registerRoot(RemotionRoot);
});
```

```

tsx

import { continueRender, delayRender, registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

loadWebAssembly().then(() => {
  registerRoot(RemotionRoot);
});
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/register-root.ts)
- [`<Composition />` component](/docs/composition)
- [The fundamentals](/docs/the-fundamentals)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/register-root.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
random()](/docs/random) [Next\
\
<Sequence>](/docs/sequence)

- [Example](#example)
- [Defer registerRoot()](#defer-registerroot)
- [See also](#see-also)
