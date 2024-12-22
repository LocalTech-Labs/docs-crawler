---
title: Using legacy Babel transpilation
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using legacy Babel transpilation

- [Home page](/)
- Tooling
- Using legacy Babel transpilation

On this page

# Using legacy Babel transpilation

In Remotion 2.0, the traditional transpilation of Javascript and Typescript using the `babel-loader` has been replaced by the faster `esbuild-loader` by default.

If you for some reason need to go back to the previous behavior, you may [override the Webpack configuration](/docs/webpack). Remember that overriding the Webpack configuration works reducer-style, where you get the default configuration in a function argument and you return the modified version of your config.

We provide a compatibility package `@remotion/babel-loader` that you can install into your Remotion project and use the function `replaceLoadersWithBabel()` to swap out the ESBuild loader with the old Babel one that was in Remotion 1.0

This should not be necessary in general, it is encouraged to [report issues](https://github.com/remotion-dev/remotion/issues/new) regarding the new ESBuild loader.

## Example [​](\#example "Direct link to Example")

- npm
- yarn
- pnpm

```

bash

npm i babel-loader @babel/preset-env @babel/preset-react
```

```

bash

npm i babel-loader @babel/preset-env @babel/preset-react
```

```

bash

pnpm i babel-loader @babel/preset-env @babel/preset-react
```

```

bash

pnpm i babel-loader @babel/preset-env @babel/preset-react
```

```

bash

yarn add babel-loader @babel/preset-env @babel/preset-react
```

```

bash

yarn add babel-loader @babel/preset-env @babel/preset-react
```

```

remotion.config.ts
ts

import { replaceLoadersWithBabel } from "@remotion/babel-loader";

Config.overrideWebpackConfig((currentConfiguration) => {
  return replaceLoadersWithBabel(currentConfiguration);
});
```

```

remotion.config.ts
ts

import { replaceLoadersWithBabel } from "@remotion/babel-loader";

Config.overrideWebpackConfig((currentConfiguration) => {
  return replaceLoadersWithBabel(currentConfiguration);
});
```

## When using `bundle` or `deploySite` [​](\#when-using-bundle-or-deploysite "Direct link to when-using-bundle-or-deploysite")

When using the Node.JS APIs - [`bundle()`](/docs/bundle) for SSR or [`deploySite()`](/docs/lambda/deploysite) for Lambda, you also need to provide the Webpack override, since the Node.JS APIs do not read from the config file.

```

my-script.js
ts

import { bundle } from "@remotion/bundler";
import { replaceLoadersWithBabel } from "@remotion/babel-loader";

await bundle({
  entryPoint: require.resolve("./src/index.ts"),
  webpackOverride: (config) => replaceLoadersWithBabel(config),
});
```

```

my-script.js
ts

import { bundle } from "@remotion/bundler";
import { replaceLoadersWithBabel } from "@remotion/babel-loader";

await bundle({
  entryPoint: require.resolve("./src/index.ts"),
  webpackOverride: (config) => replaceLoadersWithBabel(config),
});
```

## See also [​](\#see-also "Direct link to See also")

- [Custom Webpack config](/docs/webpack)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/legacy-babel-loader.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Custom Webpack config](/docs/webpack) [Next\
\
Plain JavaScript](/docs/javascript)

- [Example](#example)
- [When using `bundle` or `deploySite`](#when-using-bundle-or-deploysite)
- [See also](#see-also)
