---
title: @remotion/enable-scss
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/enable-scss

- [Home page](/)
- @remotion/enable-scss

On this page

# @remotion/enable-scss

This package provides a Webpack override for enabling [SCSS/SASS](https://sass-lang.com/) with Remotion..

## Installation [​](\#installation "Direct link to Installation")

Install `@remotion/enable-scss` as well as TailwindCSS dependencies.

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

npm i --save-exact @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

pnpm i @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

pnpm i @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

bun i @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

bun i @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

yarn --exact add @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

```

yarn --exact add @remotion/enable-scss@4.0.242 sass@1.77.2 sass-loader@14.2.1 css-loader@5.2.7Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

warning

Pay attention to install exactly these versions.

Newer versions may not work.

## Usage [​](\#usage "Direct link to Usage")

[Override the Webpack config](/docs/webpack) by using [`enableScss()`](/docs/enable-scss/enable-scss).

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss(currentConfiguration);
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss(currentConfiguration);
});
```

## APIs [​](\#apis "Direct link to APIs")

[**enableScss()** \
\
Override the Webpack config to enable SCSS](/docs/enable-scss/enable-scss)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/enable-scss/overview.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
enableTailwind()](/docs/tailwind/enable-tailwind) [Next\
\
enableScss()](/docs/enable-scss/enable-scss)

- [Installation](#installation)
- [Usage](#usage)
- [APIs](#apis)
