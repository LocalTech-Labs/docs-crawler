---
title: enableScss()
source: Remotion Documentation
last_updated: 2024-12-22
---

# enableScss()

- [Home page](/)
- [@remotion/enable-scss](/docs/enable-scss/overview)
- enableScss()

On this page

# enableScss()

_available from v4.0.162_

A function that modifies the default Webpack configuration to make the necessary changes to support SASS/SCSS.

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableScss } from "@remotion/enable-scss";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss(currentConfiguration);
});
```

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableScss } from "@remotion/enable-scss";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss(currentConfiguration);
});
```

If you want to make other configuration changes, you can do so by doing them reducer-style:

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableScss } from "@remotion/enable-scss";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss({
    ...currentConfiguration,
    // Make other changes
  });
});
```

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableScss } from "@remotion/enable-scss";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableScss({
    ...currentConfiguration,
    // Make other changes
  });
});
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/enable-scss/src/enable-scss.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/enable-scss/enable-scss.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/enable-scss/overview) [Next\
\
Overview](/docs/three)

- [See also](#see-also)
