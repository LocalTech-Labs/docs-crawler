---
title: restartStudio()v4.0.162
source: Remotion Documentation
last_updated: 2024-12-22
---

# restartStudio()v4.0.162

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- restartStudio()

On this page

# restartStudio() [v4.0.162](https://github.com/remotion-dev/remotion/releases/v4.0.162)

```

MyComp.tsx
tsx

import { restartStudio } from "@remotion/studio";
import { useCallback } from "react";

const MyComp: React.FC = () => {
  const onClick = useCallback(async () => {
    try {
      await restartStudio();
      console.log("Studio will restart now");
    } catch (err) {
      console.error(err);
    }
  }, []);

  return (
    <button type="button" onClick={onClick}>
      Hello World
    </button>
  );
};
```

```

MyComp.tsx
tsx

import { restartStudio } from "@remotion/studio";
import { useCallback } from "react";

const MyComp: React.FC = () => {
  const onClick = useCallback(async () => {
    try {
      await restartStudio();
      console.log("Studio will restart now");
    } catch (err) {
      console.error(err);
    }
  }, []);

  return (
    <button type="button" onClick={onClick}>
      Hello World
    </button>
  );
};
```

## Requirements [​](\#requirements "Direct link to Requirements")

In order to use this function:

[1](#1)

You need to be inside the Remotion Studio.

[2](#2)

The Studio must be running (no static deployment)

Otherwise, the function will throw.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/restart-studio.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/restart-studio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
deleteStaticFile()](/docs/studio/delete-static-file) [Next\
\
saveDefaultProps()](/docs/studio/save-default-props)

- [Requirements](#requirements)
- [See also](#see-also)
