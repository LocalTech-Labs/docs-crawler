---
title: TailwindCSS v2 (Legacy)
source: Remotion Documentation
last_updated: 2024-12-22
---

# TailwindCSS v2 (Legacy)

info

This documentation concerns TailwindCSS v2. [See here for V3!](/docs/tailwind)

1. Install the following dependencies:

- npm
- yarn
- pnpm

```

bash

npm i postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

```

bash

npm i postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

```

bash

pnpm i postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

```

bash

pnpm i postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

```

bash

yarn add postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

```

bash

yarn add postcss-loader postcss postcss-preset-env tailwindcss@2 autoprefixer
```

1. Add the following to your [`remotion.config.ts`](/docs/config) file:

```

ts

Config.overrideWebpackConfig((currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []
        ).filter((rule) => {
          if (!rule) {
            return false;
          }
          if (rule === "...") {
            return false;
          }
          if (rule.test?.toString().includes(".css")) {
            return false;
          }
          return true;
        }),
        {
          test: /\.css$/i,
          use: [
            "style-loader",
            "css-loader",
            {
              loader: "postcss-loader",
              options: {
                postcssOptions: {
                  plugins: [
                    "postcss-preset-env",
                    "tailwindcss",
                    "autoprefixer",
                  ],
                },
              },
            },
          ],
        },
      ],
    },
  };
});
```

```

ts

Config.overrideWebpackConfig((currentConfiguration) => {
  return {
    ...currentConfiguration,
    module: {
      ...currentConfiguration.module,
      rules: [
        ...(currentConfiguration.module?.rules
          ? currentConfiguration.module.rules
          : []
        ).filter((rule) => {
          if (!rule) {
            return false;
          }
          if (rule === "...") {
            return false;
          }
          if (rule.test?.toString().includes(".css")) {
            return false;
          }
          return true;
        }),
        {
          test: /\.css$/i,
          use: [
            "style-loader",
            "css-loader",
            {
              loader: "postcss-loader",
              options: {
                postcssOptions: {
                  plugins: [
                    "postcss-preset-env",
                    "tailwindcss",
                    "autoprefixer",
                  ],
                },
              },
            },
          ],
        },
      ],
    },
  };
});
```

1. Create a file `src/style.css` with the following content:

```

css

@tailwind base;
@tailwind components;
@tailwind utilities;
```

```

css

@tailwind base;
@tailwind components;
@tailwind utilities;
```

1. Import the stylesheet in your `src/Root.tsx` file. Add to the top of the file:

```

js

import "./style.css";
```

```

js

import "./style.css";
```

1. Start using TailwindCSS! You can verify that it's working by adding `className="bg-red-900"` to any element.

2. _Optional_: Add a `tailwind.config.js` file to the root of your project. Add `/* eslint-env node */` to the top of the file to get rid of an ESLint rule complaining that `module` is not defined.

warning

Due to a caching bug, the config file might not be picked up until you remove the `node_modules/.cache` folder - watch this issue: [https://github.com/remotion-dev/remotion/issues/315](https://github.com/remotion-dev/remotion/issues/315)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/tailwind-2.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**
