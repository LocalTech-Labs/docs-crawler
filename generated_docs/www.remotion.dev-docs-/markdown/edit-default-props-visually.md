---
title: Edit default props visually
source: Remotion Documentation
last_updated: 2024-12-22
---

# Edit default props visually

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- Visual editing

On this page

# Edit default props visually

[![](https://i.ytimg.com/vi/NX9YTOsLGpQ/hq720.jpg)\
\
Also available as a 6min video\
\
**Visual editing**](https://www.youtube.com/watch?v=NX9YTOsLGpQ)

You can edit the default props of a component visually, first register a [`schema`](/docs/composition#schema) with your [`<Composition>`](/docs/composition).

## Opening the default props editor [​](\#opening-the-default-props-editor "Direct link to Opening the default props editor")

To open the default props editor, press the
icon in the top right corner of the Remotion Studio to open the right sidebar. Alternatively, press `Cmd/Ctrl` \+ `J` to toggle the sidebar. Select the `Props` tab.

## Editing default props [​](\#editing-default-props "Direct link to Editing default props")

Use the controls in the props editor to live-update the default props of your composition. If you have changes, an undo ↩️ button will appear which you can use to revert to your changes in the code.

## Supported controls [​](\#supported-controls "Direct link to Supported controls")

Controls are implemented for:

- `z.object()`
- `z.string()`
- `z.date()`
- `z.number()`
- `z.boolean()`
- `z.array()`
- `z.union()` (only unions of two types, one of them being `z.null()` or `z.undefined()`)
- `z.optional()`
- `z.nullable()`
- `z.enum()`
- [`zColor()`](/docs/zod-types/z-color) (from `@remotion/zod-types`)
- [`zTextarea()`](/docs/zod-types/z-textarea) (from `@remotion/zod-types`)
- [`staticFile()`](/docs/staticfile) assets by typing as `z.string()` and using `staticFile()` in your code
- `.min()`
- `.max()`
- `.step()`

## Editing JSON [​](\#editing-json "Direct link to Editing JSON")

Alternatively, you can edit the JSON schema directly by pressing `JSON` in the props editor. Changes will not apply if the schema is invalid.

## Saving default props to your code [​](\#saving-default-props-to-your-code "Direct link to Saving default props to your code")

Remotion can statically analyze your root file and allow you to save props back to the code via the 💾 button. For this to work, your default props must be inlined into your [`<Composition>`](/docs/composition):

```

Inlined defaultProps
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComponent, myCompSchema} from './MyComponent';

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="my-video"
      component={MyComponent}
      durationInFrames={100}
      fps={30}
      width={1920}
      height={1080}
      schema={myCompSchema}
      defaultProps={{
        propOne: 'Hello World',
        propTwo: 'Welcome to Remotion',
      }}
    />
  );
};
```

```

Inlined defaultProps
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComponent, myCompSchema} from './MyComponent';

export const RemotionRoot: React.FC = () => {
  return (
    <Composition
      id="my-video"
      component={MyComponent}
      durationInFrames={100}
      fps={30}
      width={1920}
      height={1080}
      schema={myCompSchema}
      defaultProps={{
        propOne: 'Hello World',
        propTwo: 'Welcome to Remotion',
      }}
    />
  );
};
```

## Rendering videos from your input [​](\#rendering-videos-from-your-input "Direct link to Rendering videos from your input")

If you would like to render a video based on your input in the controls, then you don't need to save the props back to the code. Use the `Render` button to open a render dialog where your modified [default props are filled in as input props](/docs/props-resolution).

## Changing input props visually [​](\#changing-input-props-visually "Direct link to Changing input props visually")

In the Props editor, you can only edit the [default props](/docs/props-resolution) visually. You then have the chance to override them using the [`calculateMetadata()`](/docs/calculate-metadata) function.

In the [Render dialog](/docs/studio), you have the chance to change the input props.\_\_
They can also be overridden using [`calculateMetadata()`](/docs/calculate-metadata).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/visual-editing.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Defining a schema](/docs/schemas) [Next\
\
Data fetching](/docs/data-fetching)

- [Opening the default props editor](#opening-the-default-props-editor)
- [Editing default props](#editing-default-props)
- [Supported controls](#supported-controls)
- [Editing JSON](#editing-json)
- [Saving default props to your code](#saving-default-props-to-your-code)
- [Rendering videos from your input](#rendering-videos-from-your-input)
- [Changing input props visually](#changing-input-props-visually)
