---
title: TypeScript aliases
source: Remotion Documentation
last_updated: 2024-12-22
---

# TypeScript aliases

- [Home page](/)
- Tooling
- TypeScript aliases

On this page

# TypeScript aliases

Typescript aliases are not supported by default, since the ESBuild Webpack loader we have does not support them.
You can however patch the [Webpack config](/docs/webpack) to make them resolve.

Assuming you have a file:

```

 └── src/
   ├── lib/
   │   ├── one.ts
   │   ├── two.ts
   ├── Root.tsx
   └── index.ts
```

```

 └── src/
   ├── lib/
   │   ├── one.ts
   │   ├── two.ts
   ├── Root.tsx
   └── index.ts
```

and your tsconfig.json has the following `paths`:

```

json

{
  "compilerOptions": {
    "paths": {
      "lib/*": ["./src/lib/*"]
    }
  }
}
```

```

json

{
  "compilerOptions": {
    "paths": {
      "lib/*": ["./src/lib/*"]
    }
  }
}
```

you can add the aliases to Webpack, however you need to add each of them manually:

```

ts

import path from 'path';
import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((config) => {
  return {
    ...config,
    resolve: {
      ...config.resolve,
      alias: {
        ...(config.resolve?.alias ?? {}),
        lib: path.join(process.cwd(), 'src', 'lib'),
      },
    },
  };
});
```

```

ts

import path from 'path';
import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((config) => {
  return {
    ...config,
    resolve: {
      ...config.resolve,
      alias: {
        ...(config.resolve?.alias ?? {}),
        lib: path.join(process.cwd(), 'src', 'lib'),
      },
    },
  };
});
```

Remember that in Node.JS APIs, the config file does not apply, so you need to add the Webpack override also to the [`bundle()`](/docs/bundle) and [`deploySite()`](/docs/lambda/deploysite) functions.

## Automatically syncing Webpack and TypeScript aliases [​](\#automatically-syncing-webpack-and-typescript-aliases "Direct link to Automatically syncing Webpack and TypeScript aliases")

To not duplicate the aliases in your Webpack override and in your `tsconfig.json`, you can install `tsconfig-paths-webpack-plugin` and use it:

```

ts

import {Config} from '@remotion/cli/config';
import TsconfigPathsPlugin from 'tsconfig-paths-webpack-plugin';

Config.overrideWebpackConfig((config) => {
  return {
    ...config,
    resolve: {
      ...config.resolve,
      plugins: [...(config.resolve?.plugins ?? []), new TsconfigPathsPlugin()],
    },
  };
});
```

```

ts

import {Config} from '@remotion/cli/config';
import TsconfigPathsPlugin from 'tsconfig-paths-webpack-plugin';

Config.overrideWebpackConfig((config) => {
  return {
    ...config,
    resolve: {
      ...config.resolve,
      plugins: [...(config.resolve?.plugins ?? []), new TsconfigPathsPlugin()],
    },
  };
});
```

## See also [​](\#see-also "Direct link to See also")

- [Overriding webpack config](/docs/webpack)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/ts-aliases.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Third party integrations](/docs/third-party) [Next\
\
Import from Figma](/docs/figma)

- [Automatically syncing Webpack and TypeScript aliases](#automatically-syncing-webpack-and-typescript-aliases)
- [See also](#see-also)
