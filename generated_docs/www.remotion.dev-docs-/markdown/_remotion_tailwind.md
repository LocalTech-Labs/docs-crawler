---
title: @remotion/tailwind
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/tailwind

- [Home page](/)
- @remotion/tailwind

On this page

# @remotion/tailwind

This package provides utilities useful for integrating [TailwindCSS](https://tailwindcss.com/) with Remotion, as documented on our [TailwindCSS](/docs/tailwind) page.

## Installation [​](\#installation "Direct link to Installation")

Install `@remotion/tailwind` as well as TailwindCSS dependencies.

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/tailwind@4.0.242Copy
```

```

npm i --save-exact @remotion/tailwind@4.0.242Copy
```

```

pnpm i @remotion/tailwind@4.0.242Copy
```

```

pnpm i @remotion/tailwind@4.0.242Copy
```

```

bun i @remotion/tailwind@4.0.242Copy
```

```

bun i @remotion/tailwind@4.0.242Copy
```

```

yarn --exact add @remotion/tailwind@4.0.242Copy
```

```

yarn --exact add @remotion/tailwind@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

[Override the Webpack config](/docs/webpack) by using [`enableTailwind()`](/docs/tailwind/enable-tailwind).

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind(currentConfiguration);
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind(currentConfiguration);
});
```

Then follow the remaining set up steps from the [TailwindCSS](/docs/tailwind) page.

## APIs [​](\#apis "Direct link to APIs")

[**enableTailwind()** \
\
Override the Webpack config to enable TailwindCSS](/docs/tailwind/enable-tailwind)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/tailwind/overview.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
testPermissions()](/docs/cloudrun/testpermissions) [Next\
\
enableTailwind()](/docs/tailwind/enable-tailwind)

- [Installation](#installation)
- [APIs](#apis)
