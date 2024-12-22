---
title: focusDefaultPropsPath()v4.0.165
source: Remotion Documentation
last_updated: 2024-12-22
---

# focusDefaultPropsPath()v4.0.165

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- focusDefaultPropsPath()

On this page

# focusDefaultPropsPath() [v4.0.165](https://github.com/remotion-dev/remotion/releases/v4.0.165)

Scrolls to a specific field in the default props editor.

## Example [​](\#example "Direct link to Example")

For the following Zod schema:

```

schema.ts
tsx

import { z } from "zod";

const MySchema = z.object({
  array: z.array(
    z.object({
      subfield: z.string(),
    }),
  ),
});
```

```

schema.ts
tsx

import { z } from "zod";

const MySchema = z.object({
  array: z.array(
    z.object({
      subfield: z.string(),
    }),
  ),
});
```

Call `focusDefaultPropsPath()` with the path to the field you want to focus on:

```

MyComp.tsx
tsx

import { focusDefaultPropsPath } from "@remotion/studio";

focusDefaultPropsPath({
  path: ["array", 0, "subfield"],
});
```

```

MyComp.tsx
tsx

import { focusDefaultPropsPath } from "@remotion/studio";

focusDefaultPropsPath({
  path: ["array", 0, "subfield"],
});
```

## API [​](\#api "Direct link to API")

### `path` [​](\#path "Direct link to path")

The path to the field you want to focus on. An array containing numbers and strings.

### `scrollBehavior` [​](\#scrollbehavior "Direct link to scrollbehavior")

The behavior of the scrolling.

One of `"auto" | "instant" | "smooth"`.

Defaults to the [default scroll behavior](https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView).

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/focus-default-props-path.ts)
- [Visual editing](/docs/visual-editing)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/focus-default-props-path.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
updateDefaultProps()](/docs/studio/update-default-props) [Next\
\
reevaluateComposition()](/docs/studio/reevaluate-composition)

- [Example](#example)
- [API](#api)
  - [`path`](#path)
  - [`scrollBehavior`](#scrollbehavior)
- [See also](#see-also)
