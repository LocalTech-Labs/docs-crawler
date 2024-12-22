---
title: staticFile() does not support relative paths
source: Remotion Documentation
last_updated: 2024-12-22
---

# staticFile() does not support relative paths

- [Home page](/)
- Troubleshooting
- staticFile() relative paths

On this page

# staticFile() does not support relative paths

If you got the following error message:

```

staticFile() does not support relative paths. Instead, pass the name of a file that is inside the public/ folder.
```

```

staticFile() does not support relative paths. Instead, pass the name of a file that is inside the public/ folder.
```

You have tried to pass a relative path to `staticFile()`:

```

❌ Relative path
tsx

import { staticFile } from "remotion";
staticFile("../public/image.png");
```

```

❌ Relative path
tsx

import { staticFile } from "remotion";
staticFile("../public/image.png");
```

```

❌ File should not have ./ prefix
tsx

import { staticFile } from "remotion";
staticFile("./image.png");
```

```

❌ File should not have ./ prefix
tsx

import { staticFile } from "remotion";
staticFile("./image.png");
```

Or you tried to pass an absolute path:

```

❌ File should not use absolute paths
tsx

import { staticFile } from "remotion";
staticFile("/Users/bob/remotion-project/public/image.png");
```

```

❌ File should not use absolute paths
tsx

import { staticFile } from "remotion";
staticFile("/Users/bob/remotion-project/public/image.png");
```

Or you tried to add a public/ fix which is unnecessary:

```

❌ File should not include the public/ prefix
tsx

import { staticFile } from "remotion";
staticFile("public/image.png");
```

```

❌ File should not include the public/ prefix
tsx

import { staticFile } from "remotion";
staticFile("public/image.png");
```

Instead, pass the name of the file that is inside the public folder directly:

```

✅ Filename
tsx

import { staticFile } from "remotion";
staticFile("image.png");
```

```

✅ Filename
tsx

import { staticFile } from "remotion";
staticFile("image.png");
```

## See also [​](\#see-also "Direct link to See also")

- [`staticFile()`](/docs/staticfile)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/static-file-relative-paths.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Wrongly mounted <Composition>](/docs/wrong-composition-mount) [Next\
\
staticFile() remote URLs](/docs/staticfile-remote-urls)

- [See also](#see-also)
