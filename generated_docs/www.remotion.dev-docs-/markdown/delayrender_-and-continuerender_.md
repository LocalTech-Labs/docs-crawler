---
title: delayRender() and continueRender()
source: Remotion Documentation
last_updated: 2024-12-22
---

# delayRender() and continueRender()

- [Home page](/)
- [remotion](/docs/remotion)
- delayRender()

On this page

# delayRender() and continueRender()

By calling `delayRender()`, you are signaling that a frame should not be immediately rendered and instead should wait on an asynchronous task to complete.

This method is useful if you want to call an API to fetch data before you render.

`delayRender()` returns a handle. Once you have fetched data or finished the asynchronous task, you should call `continueRender(handle)` to let Remotion know that you are now ready to render.

## Example [​](\#example "Direct link to Example")

```

tsx

import { useCallback, useEffect, useState } from "react";
import { continueRender, delayRender } from "remotion";

export const MyVideo = () => {
  const [data, setData] = useState(null);
  const [handle] = useState(() => delayRender());

  const fetchData = useCallback(async () => {
    const response = await fetch("http://example.com/api");
    const json = await response.json();
    setData(json);

    continueRender(handle);
  }, []);

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

tsx

import { useCallback, useEffect, useState } from "react";
import { continueRender, delayRender } from "remotion";

export const MyVideo = () => {
  const [data, setData] = useState(null);
  const [handle] = useState(() => delayRender());

  const fetchData = useCallback(async () => {
    const response = await fetch("http://example.com/api");
    const json = await response.json();
    setData(json);

    continueRender(handle);
  }, []);

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

## Timeout [​](\#timeout "Direct link to Timeout")

You need to call `continueRender()` within 30 seconds of page load. This is the default timeout of puppeteer and it will throw if you miss to call `continueRender()`. You can [customize the timeout](/docs/timeout#increase-timeout).

If `continueRender()` is not called within the timeout frame, the render will fail with an exception similarly to this:

```

A delayRender() was called but not cleared after 28000ms. See https://remotion.dev/docs/timeout for help. The delayRender was called
```

```

A delayRender() was called but not cleared after 28000ms. See https://remotion.dev/docs/timeout for help. The delayRender was called
```

See the [Timeout](/docs/timeout) page to troubleshoot timeouts.

## Adding a label [v2.6.13](https://github.com/remotion-dev/remotion/releases/v2.6.13) [​](\#adding-a-label "Direct link to adding-a-label")

If you encounter a timeout and don't know where it came from, you can add a label as a parameter:

```

tsx

delayRender("Fetching data from API...");
```

```

tsx

delayRender("Fetching data from API...");
```

If the call times out, the label will be referenced in the error message:

```

Uncaught Error: A delayRender() "Fetching data from API..." was called but not cleared after 28000ms. See https://remotion.dev/docs/timeout for help. The delayRender was called
```

```

Uncaught Error: A delayRender() "Fetching data from API..." was called but not cleared after 28000ms. See https://remotion.dev/docs/timeout for help. The delayRender was called
```

## Concurrency [​](\#concurrency "Direct link to Concurrency")

Multiple pages are used for rendering, so delayRender() can be called multiple times for a render. If you are doing an API request, you can speed up the render and avoid rate limits by caching the request, for example by storing the data in `localStorage`.

## Multiple calls [​](\#multiple-calls "Direct link to Multiple calls")

You can call `delayRender()` multiple times. The render will be blocked for as long as at least one blocking handle exists and that has not been cleared by `continueRender()`.

```

tsx

import { useEffect, useState } from "react";
import { continueRender, delayRender } from "remotion";

const MyComp: React.FC = () => {
  const [handle1] = useState(() => delayRender());
  const [handle2] = useState(() => delayRender());

  useEffect(() => {
    // You need to clear all handles before the render continues
    continueRender(handle1);
    continueRender(handle2);
  }, []);

  return null;
};
```

```

tsx

import { useEffect, useState } from "react";
import { continueRender, delayRender } from "remotion";

const MyComp: React.FC = () => {
  const [handle1] = useState(() => delayRender());
  const [handle2] = useState(() => delayRender());

  useEffect(() => {
    // You need to clear all handles before the render continues
    continueRender(handle1);
    continueRender(handle2);
  }, []);

  return null;
};
```

## Encapsulation [​](\#encapsulation "Direct link to Encapsulation")

You should put `delayRender()` calls inside your components rather than placing them as a top-level statement, to avoid blocking a render if a different composition is rendered. Also, in the example below the call is wrapped in a `useState()` to avoid creating multiple blocking calls when the component rerenders.

```

❌ Don't do this
tsx

import { useEffect } from "react";
import { continueRender, delayRender } from "remotion";

// Don't call a delayRender() call outside a component -
// it will block the render if a different composition is rendered
// as well as block the fetching of the list of compositions.
const handle = delayRender();

const MyComp: React.FC = () => {
  useEffect(() => {
    continueRender(handle);
  }, []);

  return null;
};
```

```

❌ Don't do this
tsx

import { useEffect } from "react";
import { continueRender, delayRender } from "remotion";

// Don't call a delayRender() call outside a component -
// it will block the render if a different composition is rendered
// as well as block the fetching of the list of compositions.
const handle = delayRender();

const MyComp: React.FC = () => {
  useEffect(() => {
    continueRender(handle);
  }, []);

  return null;
};
```

## Failing with an error [v3.3.44](https://github.com/remotion-dev/remotion/releases/v3.3.44) [​](\#failing-with-an-error "Direct link to failing-with-an-error")

If your code fails to do an asynchronous operation and you want to cancel the render, you can call [`cancelRender()`](/docs/cancel-render) with an error message. This will automatically cancel all `delayRender()` calls to not further delay the render.

```

MyComposition.tsx
tsx

import React, { useEffect, useState } from "react";
import { cancelRender, continueRender, delayRender } from "remotion";

export const MyComp: React.FC = () => {
  const [handle] = useState(() => delayRender("Fetching data..."));

  useEffect(() => {
    fetch("https://example.com")
      .then(() => {
        continueRender(handle);
      })
      .catch((err) => cancelRender(err));
  }, []);

  return null;
};
```

```

MyComposition.tsx
tsx

import React, { useEffect, useState } from "react";
import { cancelRender, continueRender, delayRender } from "remotion";

export const MyComp: React.FC = () => {
  const [handle] = useState(() => delayRender("Fetching data..."));

  useEffect(() => {
    fetch("https://example.com")
      .then(() => {
        continueRender(handle);
      })
      .catch((err) => cancelRender(err));
  }, []);

  return null;
};
```

## Retrying [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#retrying "Direct link to retrying")

If an operation is flaky (for example, if loading an asset from a CDN does sometimes give 5xx errors), you can pass an object with a `retries` value as a second argument.

If a `delayRender()` call is not cleared within the timeout, the whole browser tab will be closed and the frame will be retried from scratch.

```

Retrying a delayRender()
tsx

import { delayRender } from "remotion";

delayRender("Loading asset...", {
  retries: 1, // default: 0
});
```

```

Retrying a delayRender()
tsx

import { delayRender } from "remotion";

delayRender("Loading asset...", {
  retries: 1, // default: 0
});
```

The [`<Img>`](/docs/img#delayrenderretries), [`<Audio>`](/docs/audio#delayrenderretries), [`<Video>`](/docs/video#delayrenderretries) and [`<IFrame>`](/docs/iframe#delayrenderretries) tags support a `delayRenderRetries` prop to control the value of `retries` for the `delayRender()` call that those components make.

## Modifying the timeout [v4.0.140](https://github.com/remotion-dev/remotion/releases/v4.0.140) [​](\#modifying-the-timeout "Direct link to modifying-the-timeout")

In addition to the [global timeout](/docs/timeout#increase-timeout) that can be set, the timeout can be modified on a per- `delayRender()` level.

```

Modifying the timeout of a delayRender()
tsx

import { delayRender } from "remotion";

delayRender("Loading asset...", {
  timeoutInMilliseconds: 7000,
});
```

```

Modifying the timeout of a delayRender()
tsx

import { delayRender } from "remotion";

delayRender("Loading asset...", {
  timeoutInMilliseconds: 7000,
});
```

The [`<Img>`](/docs/img#delayrendertimeoutinmilliseconds), [`<Audio>`](/docs/audio#delayrendertimeoutinmilliseconds), [`<Video>`](/docs/video#delayrendertimeoutinmilliseconds) and [`<IFrame>`](/docs/iframe#delayrendertimeoutinmilliseconds) tags support a `delayRenderTimeoutInMilliseconds` prop to control the value of `timeoutInMilliseconds` for the `delayRender()` call that those components make.

## Difference to `useBufferState().delayPlayback()` [​](\#difference-to-usebufferstatedelayplayback "Direct link to difference-to-usebufferstatedelayplayback")

[`useBufferState()`](/docs/use-buffer-state) is a different API that allows pausing playback in the [Studio](/docs/terminology/studio) and in the [Player](/docs/terminology/player).

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

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/delay-render.ts)
- [Data fetching](/docs/data-fetching)
- [continueRender()](/docs/continue-render)
- [cancelRender()](/docs/cancel-render)
- [useDelayRender()](/docs/miscellaneous/snippets/use-delay-render)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/delay-render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
continueRender()](/docs/continue-render) [Next\
\
Easing](/docs/easing)

- [Example](#example)
- [Timeout](#timeout)
- [Adding a label](#adding-a-label)
- [Concurrency](#concurrency)
- [Multiple calls](#multiple-calls)
- [Encapsulation](#encapsulation)
- [Failing with an error](#failing-with-an-error)
- [Retrying](#retrying)
- [Modifying the timeout](#modifying-the-timeout)
- [Difference to `useBufferState().delayPlayback()`](#difference-to-usebufferstatedelayplayback)
- [See also](#see-also)
