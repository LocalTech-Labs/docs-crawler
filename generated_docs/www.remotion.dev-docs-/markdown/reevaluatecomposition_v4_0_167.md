---
title: reevaluateComposition()v4.0.167
source: Remotion Documentation
last_updated: 2024-12-22
---

# reevaluateComposition()v4.0.167

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- reevaluateComposition()

On this page

# reevaluateComposition() [v4.0.167](https://github.com/remotion-dev/remotion/releases/v4.0.167)

Re-runs [`calculateMetadata()`](/docs/calculate-metadata) on the currently selected composition.

This is useful if the function relies on:

- The `public/` folder
- Randomness
- Changing network resources
- Time

## Example [​](\#example "Direct link to Example")

```

Re-run calculateMetadata() on the currently selected composition
tsx

import React, { useCallback } from "react";
import { reevaluateComposition } from "@remotion/studio";

export const ReevaluateCompositionComp: React.FC = () => {
  const reevaluate = useCallback(async () => {
    reevaluateComposition();

    console.log("Re-evaluated!");
  }, []);

  return <button onClick={reevaluate}>Re-evaluate</button>;
};
```

```

Re-run calculateMetadata() on the currently selected composition
tsx

import React, { useCallback } from "react";
import { reevaluateComposition } from "@remotion/studio";

export const ReevaluateCompositionComp: React.FC = () => {
  const reevaluate = useCallback(async () => {
    reevaluateComposition();

    console.log("Re-evaluated!");
  }, []);

  return <button onClick={reevaluate}>Re-evaluate</button>;
};
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/reevaluate-composition.ts)
- [`calculateMetadata()`](/docs/calculate-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/reevaluate-composition.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
focusDefaultPropsPath()](/docs/studio/focus-default-props-path) [Next\
\
Overview](/docs/transitions/)

- [Example](#example)
- [See also](#see-also)
