---
title: Defining a schema for your props
source: Remotion Documentation
last_updated: 2024-12-22
---

# Defining a schema for your props

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- Defining a schema

On this page

# Defining a schema for your props

As an alternative to [using TypeScript types](/docs/parameterized-rendering) to define the shape of the props your component accepts, you may use [Zod](https://zod.dev/) to define a schema for your props. You may do so if you want to edit the props visually in the Remotion Studio.

## Prerequisites [​](\#prerequisites "Direct link to Prerequisites")

If you want to use this feature, install `zod@3.22.3` and [`@remotion/zod-types`](/docs/zod-types/) for Remotion-specific types:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

npm i --save-exact @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

pnpm i @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

pnpm i @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

bun i @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

bun i @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

yarn --exact add @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

```

yarn --exact add @remotion/zod-types@4.0.242 zod@3.22.3Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## Defining a schema [​](\#defining-a-schema "Direct link to Defining a schema")

To define a schema for your props, use [`z.object()`](https://zod.dev/?id=objects):

```

tsx

import { z } from "zod";

export const myCompSchema = z.object({
  propOne: z.string(),
  propTwo: z.string(),
});
```

```

tsx

import { z } from "zod";

export const myCompSchema = z.object({
  propOne: z.string(),
  propTwo: z.string(),
});
```

Using `z.infer()`, you can turn the schema into a type:

```

tsx

export const MyComp: React.FC<z.infer<typeof myCompSchema>> = ({
  propOne,
  propTwo,
}) => {
  return (
    <div>
      props: {propOne}, {propTwo}
    </div>
  );
};
```

```

tsx

export const MyComp: React.FC<z.infer<typeof myCompSchema>> = ({
  propOne,
  propTwo,
}) => {
  return (
    <div>
      props: {propOne}, {propTwo}
    </div>
  );
};
```

## Adding a schema to your composition [​](\#adding-a-schema-to-your-composition "Direct link to Adding a schema to your composition")

Use the [`schema`](/docs/composition#schema) prop to attach the schema to your [`<Composition>`](/docs/composition). Remotion will require you to specify matching [`defaultProps`](/docs/composition#schema).

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComponent, myCompSchema } from "./MyComponent";

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
        propOne: "Hello World",
        propTwo: "Welcome to Remotion",
      }}
    />
  );
};
```

```

src/Root.tsx
tsx

import React from "react";
import { Composition } from "remotion";
import { MyComponent, myCompSchema } from "./MyComponent";

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
        propOne: "Hello World",
        propTwo: "Welcome to Remotion",
      }}
    />
  );
};
```

## Editing props visually [​](\#editing-props-visually "Direct link to Editing props visually")

When you have defined a schema for your props, you can [edit them visually in the Remotion Studio](/docs/visual-editing). This is useful if you want to quickly try out different values for your props.

## Supported types [​](\#supported-types "Direct link to Supported types")

All schemas that are supported by [Zod](https://zod.dev/) are supported by Remotion.

Remotion requires that the top-level type is a `z.object()`, because the collection of props of a React component is always an object.

In addition to the built in types, the `@remotion/zod-types` package also provides [`zColor()`](/docs/zod-types/z-color), including a color picker interface for the Remotion Studio.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/schemas.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Passing props](/docs/passing-props) [Next\
\
Visual editing](/docs/visual-editing)

- [Prerequisites](#prerequisites)
- [Defining a schema](#defining-a-schema)
- [Adding a schema to your composition](#adding-a-schema-to-your-composition)
- [Editing props visually](#editing-props-visually)
- [Supported types](#supported-types)
