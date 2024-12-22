---
title: <Folder>
source: Remotion Documentation
last_updated: 2024-12-22
---

# <Folder>

- [Home page](/)
- [remotion](/docs/remotion)
- <Folder>

On this page

# <Folder>

_Available from v3.0.1_

By wrapping a [`<Composition />`](/docs/composition) inside a `<Folder />`, you can visually categorize it in your sidebar, should you have many compositions.

## Example [​](\#example "Direct link to Example")

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

will create a tree structure in the sidebar that looks like this:

![](/img/folders.png)

## Folder behavior [​](\#folder-behavior "Direct link to Folder behavior")

- They only visually change the sidebar in the Remotion Studio and have no further behavior.
- Folder names can only contain `a-z`, `A-Z`, `0-9` and `-`.
- Folders can be nested.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this component](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/Folder.tsx)
- [`<Composition />`](/docs/composition)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/folder.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Easing](/docs/easing) [Next\
\
<Freeze>](/docs/freeze)

- [Example](#example)
- [Folder behavior](#folder-behavior)
- [See also](#see-also)
