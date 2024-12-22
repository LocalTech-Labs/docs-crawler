---
title: TailwindCSS
source: Remotion Documentation
last_updated: 2024-12-21
---

# TailwindCSS

- [Home page](/)
- Tooling
- TailwindCSS

On this page

# TailwindCSS

## Using the CLI [​](\#using-the-cli "Direct link to Using the CLI")

The easiest way to get started with Tailwind is by scaffolding a new video using the CLI and selecting a template that supports adding Tailwind automatically.

- npm
- bun
- pnpm
- yarn

```

bash

npx create-video@latest
```

```

bash

npx create-video@latest
```

```

bash

pnpm create video
```

```

bash

pnpm create video
```

```

bash

bun create video
```

```

bash

bun create video
```

```

bash

yarn create video
```

```

bash

yarn create video
```

The following templates support scaffolding with TailwindCSS:

- [Hello World](/templates/hello-world)
- [Blank](/templates/blank)
- [Hello World (JavaScript)](/templates/javascript)
- [Audiogram](/templates/audiogram)
- [Overlay](/templates/overlay)
- [Stargazer](/templates/stargazer)
- [TikTok](/templates/tiktok)

## Install in existing project [​](\#install-in-existing-project "Direct link to Install in existing project")

_from v3.3.95, see [instructions for older versions](https://github.com/remotion-dev/remotion/blob/88a5d8d17f50d6ab2b8a408556118d15a3686343/packages/docs/docs/tailwind.md)_

1. Install [`@remotion/tailwind`](/docs/tailwind/tailwind) package and TailwindCSS dependencies.

- npm
- yarn
- pnpm

```

bash

npm i -D @remotion/tailwind
```

```

bash

npm i -D @remotion/tailwind
```

```

bash

yarn add -D @remotion/tailwind
```

```

bash

yarn add -D @remotion/tailwind
```

```

bash

pnpm i -D @remotion/tailwind
```

```

bash

pnpm i -D @remotion/tailwind
```

2. Add the Webpack override from `@remotion/tailwind` to your config file:

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

note

Prior to `v3.3.39`, the option was called `Config.Bundling.overrideWebpackConfig()`.

3. If you use the [`bundle()` or `deploySite()` Node.JS API, add the Webpack override to it as well](/docs/webpack#when-using-bundle-and-deploysite).

4. Create a file `src/style.css` with the following content:

```

src/style.css
css

@tailwind base;
@tailwind components;
@tailwind utilities;
```

```

src/style.css
css

@tailwind base;
@tailwind components;
@tailwind utilities;
```

5. Import the stylesheet in your `src/Root.tsx` file. Add to the top of the file:

```

src/Root.tsx
js

import './style.css';
```

```

src/Root.tsx
js

import './style.css';
```

6. Add a `tailwind.config.js` file to the root of your project:

```

tailwind.config.js
js

/* eslint-env node */
module.exports = {
  content: ['./src/**/*.{ts,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

```

tailwind.config.js
js

/* eslint-env node */
module.exports = {
  content: ['./src/**/*.{ts,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

7. Ensure your `package.json` does not have `"sideEffects": false` set. If it has, declare that CSS files have a side effect:

```

package.json
diff

{
// Only if `"sideEffects": false` exists in your package.json.
-  "sideEffects": false
+  "sideEffects": ["*.css"]
}
```

```

package.json
diff

{
// Only if `"sideEffects": false` exists in your package.json.
-  "sideEffects": false
+  "sideEffects": ["*.css"]
}
```

8. Start using TailwindCSS! You can verify that it's working by adding `className="bg-red-900"` to any element.

note

Your package manager might show a peer dependency warning. You can safely ignore it or add `strict-peer-dependencies=false` to your `.npmrc` to suppress it.

## See also [​](\#see-also "Direct link to See also")

- [TailwindCSS v2 (legacy)](/docs/tailwind-legacy)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/tailwind.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Integration into Vue](/docs/vue) [Next\
\
Environment variables](/docs/env-variables)

- [Using the CLI](#using-the-cli)
- [Install in existing project](#install-in-existing-project)
- [See also](#see-also)
