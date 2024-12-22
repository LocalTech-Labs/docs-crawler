---
title: useDelayRender()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useDelayRender()

- [Home page](/)
- Snippets
- useDelayRender()

On this page

# useDelayRender()

Same as [`delayRender()`](/docs/delay-render), but as a hook.

## Snippet [​](\#snippet "Direct link to Snippet")

```

use-delay-render.ts
tsx

import { useCallback, useState } from "react";
import { continueRender, delayRender } from "remotion";

type ContinueRenderFnBound = () => void;

export const useDelayRender = (label?: string): ContinueRenderFnBound => {
  const [handle] = useState(() => delayRender(label));

  return useCallback(() => {
    continueRender(handle);
  }, [handle]);
};
```

```

use-delay-render.ts
tsx

import { useCallback, useState } from "react";
import { continueRender, delayRender } from "remotion";

type ContinueRenderFnBound = () => void;

export const useDelayRender = (label?: string): ContinueRenderFnBound => {
  const [handle] = useState(() => delayRender(label));

  return useCallback(() => {
    continueRender(handle);
  }, [handle]);
};
```

## Example [​](\#example "Direct link to Example")

```

MyComp.tsx
tsx

import { useCallback, useEffect, useState } from "react";

export const MyVideo = () => {
  const [data, setData] = useState(null);
  const continueRender = useDelayRender();

  const fetchData = useCallback(async () => {
    const response = await fetch("http://example.com/api");
    const json = await response.json();
    setData(json);

    continueRender();
  }, [continueRender]);

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      {data ? (
        <div>This video has data from an API! {JSON.stringify(data)}</div>
      ) : null}
    </div>
  );
};
```

```

MyComp.tsx
tsx

import { useCallback, useEffect, useState } from "react";

export const MyVideo = () => {
  const [data, setData] = useState(null);
  const continueRender = useDelayRender();

  const fetchData = useCallback(async () => {
    const response = await fetch("http://example.com/api");
    const json = await response.json();
    setData(json);

    continueRender();
  }, [continueRender]);

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      {data ? (
        <div>This video has data from an API! {JSON.stringify(data)}</div>
      ) : null}
    </div>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Data fetching](/docs/data-fetching)
- [delayRender()](/docs/delay-render)
- [continueRender()](/docs/continue-render)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/snippets/use-delay-render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Embedding a <Player> into an <iframe>](/docs/miscellaneous/snippets/player-in-iframe) [Next\
\
<OffthreadVideo /> while rendering](/docs/miscellaneous/snippets/offthread-video-while-rendering)

- [Snippet](#snippet)
- [Example](#example)
- [See also](#see-also)
