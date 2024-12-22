---
title: enableSkia()
source: Remotion Documentation
last_updated: 2024-12-22
---

# enableSkia()

- [Home page](/)
- [@remotion/skia](/docs/skia/)
- enableSkia()

# enableSkia()

A function that modifies the default Webpack configuration to make the necessary changes to support Skia.

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableSkia } from "@remotion/skia/enable";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia(currentConfiguration);
});
```

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableSkia } from "@remotion/skia/enable";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia(currentConfiguration);
});
```

note

Prior to `v3.3.39`, the option was called `Config.Bundling.overrideWebpackConfig()`.

If you want to make other configuration changes, you can do so by doing them reducer-style:

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableSkia } from "@remotion/skia/enable";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia({
    ...currentConfiguration,

    // Make other changes
  });
});
```

```

remotion.config.ts
ts

import { Config } from "@remotion/cli/config";
import { enableSkia } from "@remotion/skia/enable";

Config.overrideWebpackConfig((currentConfiguration) => {
  return enableSkia({
    ...currentConfiguration,

    // Make other changes
  });
});
```

note

Prior to `v3.3.39`, the option was called `Config.Bundling.overrideWebpackConfig()`.

See the [setup](/docs/skia) to see full instructions on how to setup React Native Skia in Remotion.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/skia/enable-skia.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Overview](/docs/skia/) [Next\
\
<SkiaCanvas />](/docs/skia/skia-canvas)
