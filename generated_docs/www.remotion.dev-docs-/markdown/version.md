---
title: VERSION
source: Remotion Documentation
last_updated: 2024-12-22
---

# VERSION

- [Home page](/)
- [remotion](/docs/remotion)
- VERSION

On this page

# VERSION

You may import this constant to get the current version of Remotion.

Only the version of the [`remotion`](/docs/remotion) package will be reported.

A [version conflict](/docs/cli/versions) with other Remotion packages cannot be ruled out.

```

Importing VERSION from remotion
ts

import { VERSION } from "remotion";

console.log(VERSION); // "4.0.57";
```

```

Importing VERSION from remotion
ts

import { VERSION } from "remotion";

console.log(VERSION); // "4.0.57";
```

You can also import it from `remotion/version` to avoid importing Remotion and its dependencies (i.e, `react` and `react-dom`):

```

Importing VERSION from remotion/version
ts

import { VERSION } from "remotion/version";

console.log(VERSION); // "4.0.57";
```

```

Importing VERSION from remotion/version
ts

import { VERSION } from "remotion/version";

console.log(VERSION); // "4.0.57";
```

## See also [​](\#see-also "Direct link to See also")

- [Source code for this constant](https://github.com/remotion-dev/remotion/blob/main/packages/core/src/version.ts)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/version.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
useVideoConfig()](/docs/use-video-config) [Next\
\
<Video>](/docs/video)

- [See also](#see-also)
