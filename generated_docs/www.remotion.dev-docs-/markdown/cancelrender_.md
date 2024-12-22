---
title: cancelRender()
source: Remotion Documentation
last_updated: 2024-12-22
---

# cancelRender()

- [Home page](/)
- [remotion](/docs/remotion)
- cancelRender()

On this page

# cancelRender()

_Available from v3.3.44_

By invoking `cancelRender()`, Remotion will stop rendering all the frames, and not do any retries.

Pass a `string` or an `Error` (for best stack traces) to `cancelRender()` so you can identify the error when your render failed.

## Example [​](\#example "Direct link to Example")

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

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/cancel-render.ts)
- [delayRender()](/docs/delay-render)
- [continueRender()](/docs/continue-render)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/cancel-render.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
calculateMetadata()](/docs/calculate-metadata) [Next\
\
<Composition>](/docs/composition)

- [Example](#example)
- [See also](#see-also)
