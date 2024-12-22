---
title: getRemotionEnvironment()v4.0.25
source: Remotion Documentation
last_updated: 2024-12-22
---

# getRemotionEnvironment()v4.0.25

- [Home page](/)
- [remotion](/docs/remotion)
- getRemotionEnvironment()

On this page

# getRemotionEnvironment() [v4.0.25](https://github.com/remotion-dev/remotion/releases/v4.0.25)

With the `getRemotionEnvironment()` function, you can retrieve information about the current Remotion Environment.
It returns an object with the following properties:

- `isStudio`: Whether the function got called in the [Remotion Studio](/docs/cli/studio).
- `isRendering`: Whether the function got called in a render.
- `isPlayer`: Whether a [`<Player>`](/docs/player) is mounted on the current page.
- `isReadOnlyStudio`: Whether in a [statically deployed studio](https://www.remotion.dev/docs/studio/deploy-static), where the [`@remotion/studio`](/docs/studio/api) APIs cannot be used ( _available from v4.0.238_)

This can be useful if you want components or functions to behave differently depending on the environment.

### Example [​](\#example "Direct link to Example")

```

tsx

import React from 'react';
import {getRemotionEnvironment} from 'remotion';

export const MyComp: React.FC = () => {
  const {isStudio, isPlayer, isRendering} = getRemotionEnvironment();

  if (isStudio) {
    return <div>I'm in the Studio!</div>;
  }

  if (isPlayer) {
    return <div>I'm in the Player!</div>;
  }

  if (isRendering) {
    return <div>I'm Rendering</div>;
  }

  return <div>Hello World!</div>;
};
```

```

tsx

import React from 'react';
import {getRemotionEnvironment} from 'remotion';

export const MyComp: React.FC = () => {
  const {isStudio, isPlayer, isRendering} = getRemotionEnvironment();

  if (isStudio) {
    return <div>I'm in the Studio!</div>;
  }

  if (isPlayer) {
    return <div>I'm in the Player!</div>;
  }

  if (isRendering) {
    return <div>I'm Rendering</div>;
  }

  return <div>Hello World!</div>;
};
```

A more realistic use case: You might want to debounce a request during editing in the Remotion Studio, but not during rendering. See: [debouncing requests](/docs/data-fetching#debouncing-requests)

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/get-remotion-environment.ts)
- [useVideoConfig()](/docs/use-video-config)
- [`<OffthreadVideo/> while rendering`](/docs/miscellaneous/snippets/offthread-video-while-rendering)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/get-remotion-environment.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
getInputProps()](/docs/get-input-props) [Next\
\
getStaticFiles()](/docs/getstaticfiles)

- [Example](#example)
- [See also](#see-also)
