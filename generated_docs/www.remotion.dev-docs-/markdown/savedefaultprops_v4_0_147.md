---
title: saveDefaultProps()v4.0.147
source: Remotion Documentation
last_updated: 2024-12-22
---

# saveDefaultProps()v4.0.147

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- saveDefaultProps()

On this page

# saveDefaultProps() [v4.0.147](https://github.com/remotion-dev/remotion/releases/v4.0.147)

Saves the [`defaultProps`](/docs/composition) for a [composition](/docs/terminology/composition) back to the [root file](/docs/terminology/root-file).

If you just want to update the default props in the Props Editor (right sidebar in the Studio) without saving them to the root file, use [`updateDefaultProps()`](/docs/studio/update-default-props).

## Examples [​](\#examples "Direct link to Examples")

```

Saving {color: 'green'} to Root.tsx
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: () => {
    return {
      color: "green",
    };
  },
});
```

```

Saving {color: 'green'} to Root.tsx
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: () => {
    return {
      color: "green",
    };
  },
});
```

You can access the saved default props to only override part of it (reducer-style):

```

Accessing the saved default props
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ savedDefaultProps }) => {
    return {
      ...savedDefaultProps,
      color: "green",
    };
  },
});
```

```

Accessing the saved default props
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ savedDefaultProps }) => {
    return {
      ...savedDefaultProps,
      color: "green",
    };
  },
});
```

If you modified props in the Props Editor (right sidebar in the Studio), you can also access the unsaved props from there, and for example save them:

```

Save props from the Props Editor
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ unsavedDefaultProps }) => {
    return unsavedDefaultProps;
  },
});
```

```

Save props from the Props Editor
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ unsavedDefaultProps }) => {
    return unsavedDefaultProps;
  },
});
```

If you have a Zod schema, you can also access its runtime value:

```

Save props from the Props Editor
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ schema, unsavedDefaultProps }) => {
    // Do something with the Zod schema

    return {
      ...unsavedDefaultProps,
      color: "red",
    };
  },
});
```

```

Save props from the Props Editor
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ schema, unsavedDefaultProps }) => {
    // Do something with the Zod schema

    return {
      ...unsavedDefaultProps,
      color: "red",
    };
  },
});
```

## Requirements [​](\#requirements "Direct link to Requirements")

In order to use this function:

[1](#1)

You need to be inside the Remotion Studio.

[2](#2)

The Studio must be running (no static deployment)

[3](#3)

`zod` needs to be installed.

Otherwise, the function will throw.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/save-default-props.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/save-default-props.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
restartStudio()](/docs/studio/restart-studio) [Next\
\
updateDefaultProps()](/docs/studio/update-default-props)

- [Examples](#examples)
- [Requirements](#requirements)
- [See also](#see-also)
