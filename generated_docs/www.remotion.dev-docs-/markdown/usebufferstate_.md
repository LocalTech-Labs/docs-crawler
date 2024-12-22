---
title: useBufferState()
source: Remotion Documentation
last_updated: 2024-12-22
---

# useBufferState()

- [Home page](/)
- [remotion](/docs/remotion)
- useBufferState()

On this page

# useBufferState()

_available from 4.0.111_

You can use this hook inside your [composition](/docs/terminology/composition) to retrieve functions that can inform the Player that your component is currently loading data.

```

MyComp.tsx
tsx

import React from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

```

MyComp.tsx
tsx

import React from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

## API [​](\#api "Direct link to API")

Returns an object with one method:

### `delayPlayback()` [​](\#delayplayback "Direct link to delayplayback")

Tells the Player to delay playback until you call `unblock()`.

## Usage notes [​](\#usage-notes "Direct link to Usage notes")

### Handle unmounting [​](\#handle-unmounting "Direct link to Handle unmounting")

The user might seek to a different portion of the video which is immediately available.

Use the cleanup function of `useEffect()` to clear the handle when your component is unmounted.

```

❌ Causes problems with React strict mode
tsx

import React, { useState } from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);
  }, []);

  return <></>;
};
```

```

❌ Causes problems with React strict mode
tsx

import React, { useState } from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);
  }, []);

  return <></>;
};
```

### Avoid calling inside `useState()` [​](\#avoid-calling-inside-usestate "Direct link to avoid-calling-inside-usestate")

While the following implementation works in production, it fails in development if you are inside React Strict mode, because the `useState()` hook is called twice, which causes the first invocation of the buffer to never be cleared and the video to buffer forever.

```

❌ Doesn't clear the buffer handle when seeking to a different portion of a video
tsx

import React, { useState } from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

```

❌ Doesn't clear the buffer handle when seeking to a different portion of a video
tsx

import React, { useState } from "react";
import { useBufferState } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

### Together with `delayRender()` [​](\#together-with-delayrender "Direct link to together-with-delayrender")

[`delayRender()`](/docs/delay-render) is a different API that comes into play during rendering.

If you are loading data, you might want to both delay the screenshotting of your component during rendering and start a buffering state during Preview, in which case you need to use both APIs together.

```

Using delayRender() and delayPlayback() together
tsx

import React from "react";
import { useBufferState, delayRender, continueRender } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [handle] = React.useState(() => delayRender());

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
      continueRender(handle);
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

```

Using delayRender() and delayPlayback() together
tsx

import React from "react";
import { useBufferState, delayRender, continueRender } from "remotion";

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [handle] = React.useState(() => delayRender());

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
      continueRender(handle);
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Buffer state](/docs/player/buffer-state)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/use-buffer-state.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
<Still>](/docs/still) [Next\
\
useCurrentFrame()](/docs/use-current-frame)

- [API](#api)
  - [`delayPlayback()`](#delayplayback)
- [Usage notes](#usage-notes)
  - [Handle unmounting](#handle-unmounting)
  - [Avoid calling inside `useState()`](#avoid-calling-inside-usestate)
  - [Together with `delayRender()`](#together-with-delayrender)
- [See also](#see-also)
