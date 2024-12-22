---
title: Webpack dynamic imports
source: Remotion Documentation
last_updated: 2024-12-22
---

# Webpack dynamic imports

- [Home page](/)
- Troubleshooting
- Webpack dynamic imports

On this page

# Webpack dynamic imports

A common need in Remotion is to import assets from dynamic paths. This means that sometimes you don't know the exact path of an asset that should be imported upfront and you want to calculate it during runtime.

This can become an unexpected hurdle in Webpack. On this page we collect some tips how to handle dynamic assets.

Take this scenario for an example:

```

tsx

import { Img, useCurrentFrame } from "remotion";

export const DynamicImports: React.FC = () => {
  const frame = useCurrentFrame();
  const img = "./assets/image" + frame + ".png";
  return <Img src={require(img)} />;
};
```

```

tsx

import { Img, useCurrentFrame } from "remotion";

export const DynamicImports: React.FC = () => {
  const frame = useCurrentFrame();
  const img = "./assets/image" + frame + ".png";
  return <Img src={require(img)} />;
};
```

may result in:

```

bash

Error: Cannot find module './image0.png'
```

```

bash

Error: Cannot find module './image0.png'
```

even if the file exists. This is because Webpack needs to figure out using static code analysis which assets it should bundle and cannot do so.

## Recommended: Use `staticFile()` instead [​](\#recommended-use-staticfile-instead "Direct link to recommended-use-staticfile-instead")

We recommend that you put the asset inside the `public/` folder and use `staticFile()` to reference it. This new way does not suffer from the underlying problem.

- [Importing assets](/docs/assets)
- [`staticFile()`](/docs/staticfile)

## Write dynamic expressions correctly [​](\#write-dynamic-expressions-correctly "Direct link to Write dynamic expressions correctly")

While the example at the top did not work, Webpack is smart enough to do so if you place your expression inside the `require()` or `import()` statement. In this case, Webpack will automatically bundle all `.png` files in the `assets/image` folder even if the asset is never used.

The following **does** work:

```

tsx

import { Img, useCurrentFrame } from "remotion";

export const DynamicImports: React.FC = () => {
  const frame = useCurrentFrame();
  return <Img src={require("./assets/image" + frame + ".png")} />;
};
```

```

tsx

import { Img, useCurrentFrame } from "remotion";

export const DynamicImports: React.FC = () => {
  const frame = useCurrentFrame();
  return <Img src={require("./assets/image" + frame + ".png")} />;
};
```

Please read [the Webpack documentation page](https://webpack.js.org/api/module-methods/#dynamic-expressions-in-import) about this behavior if you would like to learn more.

## Import assets at runtime [​](\#import-assets-at-runtime "Direct link to Import assets at runtime")

Let's imagine a scenario where the asset that should be imported is completely unknown and will be read at runtime, for example through an [input prop](/docs/get-input-props):

```

tsx

import { getInputProps, Img } from "remotion";

const DynamicAsset: React.FC = () => {
  const inputProps = getInputProps(); // {"imageSrc": "./assets/img0.png"}
  return <Img src={require(inputProps.imageSrc as string)} />;
};
```

```

tsx

import { getInputProps, Img } from "remotion";

const DynamicAsset: React.FC = () => {
  const inputProps = getInputProps(); // {"imageSrc": "./assets/img0.png"}
  return <Img src={require(inputProps.imageSrc as string)} />;
};
```

This cannot work because Webpack has no idea which assets it has to bundle. Therefore the import has to fail.
Like above, you can force Webpack to bundle the whole assets folder by putting an expression inside the `require()` statement:

```

tsx

import { getInputProps, Img } from "remotion";

const DynamicAsset: React.FC = () => {
  const inputProps = getInputProps(); // {"imageSrc": "img0.png"}
  // Works!
  return <Img src={require(("./assets/" + inputProps.imageSrc) as string)} />;
};
```

```

tsx

import { getInputProps, Img } from "remotion";

const DynamicAsset: React.FC = () => {
  const inputProps = getInputProps(); // {"imageSrc": "img0.png"}
  // Works!
  return <Img src={require(("./assets/" + inputProps.imageSrc) as string)} />;
};
```

## Struggling? [​](\#struggling "Direct link to Struggling?")

If you still have troubles importing your assets, hit us up on [Discord](https://remotion.dev/discord) or file an issue. We'd love to hear your input about how we can improve and will help you out.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/dynamic-import.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Performance](/docs/performance) [Next\
\
Non-seekable media](/docs/non-seekable-media)

- [Recommended: Use `staticFile()` instead](#recommended-use-staticfile-instead)
- [Write dynamic expressions correctly](#write-dynamic-expressions-correctly)
- [Import assets at runtime](#import-assets-at-runtime)
- [Struggling?](#struggling)
