---
title: enableTailwind()
source: Remotion Documentation
last_updated: 2024-12-22
---

# enableTailwind()

- [Home page](/)
- [@remotion/tailwind](/docs/tailwind/tailwind)
- enableTailwind()

On this page

# enableTailwind()

_available from v3.3.95_

A function that modifies the default Webpack configuration to make the necessary changes to support TailwindCSS.
See the [setup](/docs/tailwind) to see full instructions on how to setup TailwindCSS in Remotion.

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

## Multiple Webpack changes [​](\#multiple-webpack-changes "Direct link to Multiple Webpack changes")

If you want to make other configuration changes, you can do so by doing them reducer-style:

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind({
    ...currentConfiguration,

    // Make other changes
  });
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind({
    ...currentConfiguration,

    // Make other changes
  });
});
```

## Custom Tailwind config location [v4.0.187](https://github.com/remotion-dev/remotion/releases/v4.0.187) [​](\#custom-tailwind-config-location "Direct link to custom-tailwind-config-location")

By default, TailwindCSS will search for a file called `tailwind.config.js` in the current working directory where you executed the Remotion CLI.

This is not in line with Remotion which resolves files relative to the [Remotion Root](/docs/terminology/remotion-root).

This mean if you execute the Remotion CLI from a different directory, Tailwind will not be able to find the config file.

To fix this, you can pass the path to the config file as the second argument to `enableTailwind()`:

```

remotion.config.ts
ts

import path from 'node:path';
import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind(currentConfiguration, {
    configLocation: path.join(__dirname, 'tailwind.config.js'),
  });
});
```

```

remotion.config.ts
ts

import path from 'node:path';
import {Config} from '@remotion/cli/config';
import {enableTailwind} from '@remotion/tailwind';

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableTailwind(currentConfiguration, {
    configLocation: path.join(__dirname, 'tailwind.config.js'),
  });
});
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/tailwind/enable-tailwind.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/tailwind/tailwind) [Next\
\
Overview](/docs/enable-scss/overview)

- [Multiple Webpack changes](#multiple-webpack-changes)
- [Custom Tailwind config location](#custom-tailwind-config-location)
