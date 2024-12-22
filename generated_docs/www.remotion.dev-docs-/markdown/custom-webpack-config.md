---
title: Custom Webpack config
source: Remotion Documentation
last_updated: 2024-12-22
---

# Custom Webpack config

- [Home page](/)
- Tooling
- Custom Webpack config

On this page

# Custom Webpack config

Remotion ships with [it's own Webpack configuration](https://github.com/remotion-dev/remotion/blob/main/packages/bundler/src/webpack-config.ts).

You can override it reducer-style by creating a function that takes the previous Webpack configuration and returns the the new one.

## When rendering using the command line [​](\#when-rendering-using-the-command-line "Direct link to When rendering using the command line")

In your [`remotion.config.ts`](/docs/config) file, you can call [`Config.overrideWebpackConfig()`](/docs/config#overridewebpackconfig) from `@remotion/cli/config`.

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules ?? []),
        // Add more loaders here
      ],
    },
  };
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules ?? []),
        // Add more loaders here
      ],
    },
  };
});
```

info

Using the reducer pattern will help with type safety, give you auto-complete, ensure forwards-compatibility and keep it completely flexible - you can override just one property or pass in a completely new Webpack configuration.

## When using `bundle()` and `deploySite()` [​](\#when-using-bundle-and-deploysite "Direct link to when-using-bundle-and-deploysite")

When using the Node.JS APIs - [`bundle()`](/docs/bundle) for SSR or [`deploySite()`](/docs/lambda/deploysite) for Lambda, you also need to provide the Webpack override, since the Node.JS APIs do not read from the config file. We recommend you put the webpack override in a separate file so you can read it from both the command line and your Node.JS script.

```

src/webpack-override.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const webpackOverride: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    // Your override here
  };
};
```

```

src/webpack-override.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const webpackOverride: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    // Your override here
  };
};
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {webpackOverride} from './src/webpack-override';

Config.overrideWebpackConfig(webpackOverride);
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {webpackOverride} from './src/webpack-override';

Config.overrideWebpackConfig(webpackOverride);
```

With `bundle`:

```

my-script.js
ts

import {bundle} from '@remotion/bundler';
import {webpackOverride} from './src/webpack-override';

await bundle({
  entryPoint: require.resolve('./src/index.ts'),
  webpackOverride,
});
```

```

my-script.js
ts

import {bundle} from '@remotion/bundler';
import {webpackOverride} from './src/webpack-override';

await bundle({
  entryPoint: require.resolve('./src/index.ts'),
  webpackOverride,
});
```

Or while using with `deploySite`:

```

my-script.js
ts

import {deploySite} from '@remotion/lambda';
import {webpackOverride} from './src/webpack-override';

await deploySite({
  entryPoint: require.resolve('./src/index.ts'),
  region: 'us-east-1',
  bucketName: 'remotionlambda-c7fsl3d',
  options: {
    webpackOverride,
  },
  // ...other parameters
});
```

```

my-script.js
ts

import {deploySite} from '@remotion/lambda';
import {webpackOverride} from './src/webpack-override';

await deploySite({
  entryPoint: require.resolve('./src/index.ts'),
  region: 'us-east-1',
  bucketName: 'remotionlambda-c7fsl3d',
  options: {
    webpackOverride,
  },
  // ...other parameters
});
```

## Multiple Webpack overrides [​](\#multiple-webpack-overrides "Direct link to Multiple Webpack overrides")

If you have multiple overrides, you should curry them:

```

tsx

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((c) => enableScss(enableTailwind(c)));
```

```

tsx

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((c) => enableScss(enableTailwind(c)));
```

From Remotion v4.0.229, you can also use `Config.overrideWebpackConfig` multiple times, but this only works in the config file:

```

tsx

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig(enableScss);
Config.overrideWebpackConfig(enableTailwind);
```

```

tsx

import {Config} from '@remotion/cli/config';
import {enableScss} from '@remotion/enable-scss';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig(enableScss);
Config.overrideWebpackConfig(enableTailwind);
```

## Snippets [​](\#snippets "Direct link to Snippets")

### Enabling MDX support [​](\#enabling-mdx-support "Direct link to Enabling MDX support")

1. Install the following dependencies:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @mdx-js/loader @mdx-js/reactCopy
```

```

npm i --save-exact @mdx-js/loader @mdx-js/reactCopy
```

```

pnpm i @mdx-js/loader @mdx-js/reactCopy
```

```

pnpm i @mdx-js/loader @mdx-js/reactCopy
```

```

bun i @mdx-js/loader @mdx-js/reactCopy
```

```

bun i @mdx-js/loader @mdx-js/reactCopy
```

```

yarn --exact add @mdx-js/loader @mdx-js/reactCopy
```

```

yarn --exact add @mdx-js/loader @mdx-js/reactCopy
```

1. Create a file with the Webpack override:

```

enable-mdx.ts
ts

export const enableMdx: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []),
        {
          test: /\.mdx?$/,
          use: [
            {
              loader: '@mdx-js/loader',
              options: {},
            },
          ],
        },
      ],
    },
  };
};
```

```

enable-mdx.ts
ts

export const enableMdx: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []),
        {
          test: /\.mdx?$/,
          use: [
            {
              loader: '@mdx-js/loader',
              options: {},
            },
          ],
        },
      ],
    },
  };
};
```

1. Add it to the config file:

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableMdx} from './src/enable-mdx';

Config.overrideWebpackConfig(enableMdx);
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableMdx} from './src/enable-mdx';

Config.overrideWebpackConfig(enableMdx);
```

1. Add it to your [Node.JS API calls as well if necessary](#when-using-bundle-and-deploysite).

2. Create a file which contains `declare module '*.mdx';` in your project to fix a TypeScript error showing up.

### Enable TailwindCSS support [​](\#enable-tailwindcss-support "Direct link to Enable TailwindCSS support")

- [TailwindCSS V3](/docs/tailwind)
- [TailwindCSS V2 (Legacy)](/docs/tailwind-legacy)

### Enable SASS/SCSS support [​](\#enable-sassscss-support "Direct link to Enable SASS/SCSS support")

The easiest way is to use the `@remotion/enable-scss`.

Follow these [instructions](/docs/enable-scss/overview) to enable it.

### Enable SVGR support [​](\#enable-svgr-support "Direct link to Enable SVGR support")

This allows you to enable `import` SVG files as React components.

1. Install the following dependency:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @svgr/webpackCopy
```

```

npm i --save-exact @svgr/webpackCopy
```

```

pnpm i @svgr/webpackCopy
```

```

pnpm i @svgr/webpackCopy
```

```

bun i @svgr/webpackCopy
```

```

bun i @svgr/webpackCopy
```

```

yarn --exact add @svgr/webpackCopy
```

```

yarn --exact add @svgr/webpackCopy
```

1. Declare an override function:

```

src/enable-svgr.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const enableSvgr: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        {
          test: /\.svg$/i,
          issuer: /\.[jt]sx?$/,
          resourceQuery: {not: [/url/]}, // Exclude react component if *.svg?url
          use: ['@svgr/webpack'],
        },
        ...(currentConfiguration.module?.rules ?? []).map((r) => {
          if (!r) {
            return r;
          }
          if (r === '...') {
            return r;
          }
          if (!r.test?.toString().includes('svg')) {
            return r;
          }
          return {
            ...r,
            // Remove Remotion loading SVGs as a URL
            test: new RegExp(
              r.test.toString().replace(/svg\|/g, '').slice(1, -1),
            ),
          };
        }),
      ],
    },
  };
};
```

```

src/enable-svgr.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const enableSvgr: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        {
          test: /\.svg$/i,
          issuer: /\.[jt]sx?$/,
          resourceQuery: {not: [/url/]}, // Exclude react component if *.svg?url
          use: ['@svgr/webpack'],
        },
        ...(currentConfiguration.module?.rules ?? []).map((r) => {
          if (!r) {
            return r;
          }
          if (r === '...') {
            return r;
          }
          if (!r.test?.toString().includes('svg')) {
            return r;
          }
          return {
            ...r,
            // Remove Remotion loading SVGs as a URL
            test: new RegExp(
              r.test.toString().replace(/svg\|/g, '').slice(1, -1),
            ),
          };
        }),
      ],
    },
  };
};
```

1. Add the override function to your [`remotion.config.ts`](/docs/config) file:

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableSvgr} from './src/enable-svgr';

Config.overrideWebpackConfig(enableSvgr);
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableSvgr} from './src/enable-svgr';

Config.overrideWebpackConfig(enableSvgr);
```

1. Add it to your [Node.JS API calls as well if necessary](#when-using-bundle-and-deploysite).

2. Restart the Remotion Studio.

### Enable support for GLSL imports [​](\#enable-support-for-glsl-imports "Direct link to Enable support for GLSL imports")

1. Install the following dependencies:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

npm i --save-exact glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

pnpm i glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

pnpm i glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

bun i glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

bun i glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

yarn --exact add glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

```

yarn --exact add glsl-shader-loader glslify glslify-import-loader raw-loaderCopy
```

1. Declare a webpack override:

```

src/enable.glsl.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const enableGlsl: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []),
        {
          test: /\.(glsl|vs|fs|vert|frag)$/,
          exclude: /node_modules/,
          use: ['glslify-import-loader', 'raw-loader', 'glslify-loader'],
        },
      ],
    },
  };
};
```

```

src/enable.glsl.ts
ts

import {WebpackOverrideFn} from '@remotion/bundler';

export const enableGlsl: WebpackOverrideFn = (currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []),
        {
          test: /\.(glsl|vs|fs|vert|frag)$/,
          exclude: /node_modules/,
          use: ['glslify-import-loader', 'raw-loader', 'glslify-loader'],
        },
      ],
    },
  };
};
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableGlsl} from './src/enable-glsl';

Config.overrideWebpackConfig(enableGlsl);
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableGlsl} from './src/enable-glsl';

Config.overrideWebpackConfig(enableGlsl);
```

1. Add the following to your [entry point](/docs/terminology/entry-point) (e.g. `src/index.ts`):

```

ts

declare module '*.glsl' {
  const value: string;
  export default value;
}
```

```

ts

declare module '*.glsl' {
  const value: string;
  export default value;
}
```

1. Add it to your [Node.JS API calls as well if necessary](#when-using-bundle-and-deploysite).

2. Reset the webpack cache by deleting the `node_modules/.cache` folder.

3. Restart the Remotion Studio.

### Enable WebAssembly [​](\#enable-webassembly "Direct link to Enable WebAssembly")

There are two WebAssembly modes: asynchronous and synchronous. We recommend testing both and seeing which one works for the WASM library you are trying to use.

```

remotion.config.ts - synchronous
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((conf) => {
  return {
    ...conf,
    experiments: {
      syncWebAssembly: true,
    },
  };
});
```

```

remotion.config.ts - synchronous
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((conf) => {
  return {
    ...conf,
    experiments: {
      syncWebAssembly: true,
    },
  };
});
```

note

Since Webpack does not allow synchronous WebAssembly code in the main chunk, you most likely need to declare your composition using [`lazyComponent`](/docs/composition#example-using-lazycomponent) instead of `component`. Check out a [demo project](https://github.com/remotion-dev/id3-tags) for an example.

```

remotion.config.ts - asynchronous
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((conf) => {
  return {
    ...conf,
    experiments: {
      asyncWebAssembly: true,
    },
  };
});
```

```

remotion.config.ts - asynchronous
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig((conf) => {
  return {
    ...conf,
    experiments: {
      asyncWebAssembly: true,
    },
  };
});
```

After you've done that, clear the Webpack cache:

```

bash

rm -rf node_modules/.cache
```

```

bash

rm -rf node_modules/.cache
```

After restarting, you can import `.wasm` files using an import statement.

Add the Webpack override to your [Node.JS API calls as well if necessary](#when-using-bundle-and-deploysite).

### Use legacy babel loader [​](\#use-legacy-babel-loader "Direct link to Use legacy babel loader")

See [Using legacy Babel transpilation](/docs/legacy-babel).

## Enable TypeScript aliases [​](\#enable-typescript-aliases "Direct link to Enable TypeScript aliases")

See [TypeScript aliases](/docs/typescript-aliases).

## Customizing configuration file location [​](\#customizing-configuration-file-location "Direct link to Customizing configuration file location")

You can pass a `--config` option to the command line to specify a custom location for your configuration file.

## Importing ES Modules in remotion.config.ts [v4.0.117](https://github.com/remotion-dev/remotion/releases/v4.0.117) [​](\#importing-es-modules-in-remotionconfigts "Direct link to importing-es-modules-in-remotionconfigts")

The [config file](/docs/config) gets executed in a CommonJS environment. If you want to import ES modules, you can pass an async function to `Config.overrideWebpackConfig`:

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig(async (currentConfiguration) => {
  const {enableSass} = await import('./src/enable-sass');
  return enableSass(currentConfiguration);
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig(async (currentConfiguration) => {
  const {enableSass} = await import('./src/enable-sass');
  return enableSass(currentConfiguration);
});
```

## See also [​](\#see-also "Direct link to See also")

- [Configuration file](/docs/config)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/overwriting-webpack-config.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Environment variables](/docs/env-variables) [Next\
\
Using legacy Babel transpilation](/docs/legacy-babel)

- [When rendering using the command line](#when-rendering-using-the-command-line)
- [When using `bundle()` and `deploySite()`](#when-using-bundle-and-deploysite)
- [Multiple Webpack overrides](#multiple-webpack-overrides)
- [Snippets](#snippets)
  - [Enabling MDX support](#enabling-mdx-support)
  - [Enable TailwindCSS support](#enable-tailwindcss-support)
  - [Enable SASS/SCSS support](#enable-sassscss-support)
  - [Enable SVGR support](#enable-svgr-support)
  - [Enable support for GLSL imports](#enable-support-for-glsl-imports)
  - [Enable WebAssembly](#enable-webassembly)
  - [Use legacy babel loader](#use-legacy-babel-loader)
- [Enable TypeScript aliases](#enable-typescript-aliases)
- [Customizing configuration file location](#customizing-configuration-file-location)
- [Importing ES Modules in remotion.config.ts](#importing-es-modules-in-remotionconfigts)
- [See also](#see-also)
