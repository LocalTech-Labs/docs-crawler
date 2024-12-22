---
title: updateDefaultProps()v4.0.154
source: Remotion Documentation
last_updated: 2024-12-22
---

# updateDefaultProps()v4.0.154

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- updateDefaultProps()

On this page

# updateDefaultProps() [v4.0.154](https://github.com/remotion-dev/remotion/releases/v4.0.154)

Updates the default props in the Props Editor (in the right sidebar in the Studio).

Your component will be re-rendered with the new props.

The props will **not** be saved to the [Root file](/docs/terminology/root-file) \- use [`saveDefaultProps()`](/docs/studio/save-default-props) for that.

## Examples [​](\#examples "Direct link to Examples")

```

Setting {color: 'green'} as the default props
tsx

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
  compositionId: "my-composition",
  defaultProps: () => {
    return {
      color: "green",
    };
  },
});
```

```

Setting {color: 'green'} as the default props
tsx

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
  compositionId: "my-composition",
  defaultProps: () => {
    return {
      color: "green",
    };
  },
});
```

You can access the current unsaved default props to only override part of it (reducer-style):

```

Accessing the current props
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ unsavedDefaultProps }) => {
    return { ...unsavedDefaultProps, color: "green" };
  },
});
```

```

Accessing the current props
tsx

import { saveDefaultProps } from "@remotion/studio";

await saveDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ unsavedDefaultProps }) => {
    return { ...unsavedDefaultProps, color: "green" };
  },
});
```

If you only want to override on top of the saved changes:

```

Accessing the saved props
tsx

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
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

Accessing the saved props
tsx

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
  compositionId: "my-composition",
  defaultProps: ({ savedDefaultProps }) => {
    return {
      ...savedDefaultProps,
      color: "green",
    };
  },
});
```

If you have a Zod schema, you can also access its runtime value:

```

Save props from the Props Editor
tsx

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
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

import { updateDefaultProps } from "@remotion/studio";

updateDefaultProps({
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

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/update-default-props.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/update-default-props.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
saveDefaultProps()](/docs/studio/save-default-props) [Next\
\
focusDefaultPropsPath()](/docs/studio/focus-default-props-path)

- [Examples](#examples)
- [Requirements](#requirements)
- [See also](#see-also)
