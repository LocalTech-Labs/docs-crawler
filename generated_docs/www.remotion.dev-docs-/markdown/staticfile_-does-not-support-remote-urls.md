---
title: staticFile() does not support remote URLs
source: Remotion Documentation
last_updated: 2024-12-22
---

# staticFile() does not support remote URLs

- [Home page](/)
- Troubleshooting
- staticFile() remote URLs

On this page

# staticFile() does not support remote URLs

If you got the following error message:

```

staticFile() does not support remote URLs. Instead, pass the URL without wrapping it in staticFile().
```

```

staticFile() does not support remote URLs. Instead, pass the URL without wrapping it in staticFile().
```

You have tried to pass a remote URL to `staticFile()`. You don't need to wrap the path in `staticFile`, you can pass it directly:

```

❌ Remote URL inside staticFile()
tsx

import { Img, staticFile } from "remotion";

const MyComp = () => {
  return <Img src={staticFile("https://example.com/image.png")} />;
};
```

```

❌ Remote URL inside staticFile()
tsx

import { Img, staticFile } from "remotion";

const MyComp = () => {
  return <Img src={staticFile("https://example.com/image.png")} />;
};
```

Instead, :

```

✅ Remote URL passed directly
tsx

import { Img } from "remotion";

const MyComp = () => {
  return <Img src={"https://example.com/image.png"} />;
};
```

```

✅ Remote URL passed directly
tsx

import { Img } from "remotion";

const MyComp = () => {
  return <Img src={"https://example.com/image.png"} />;
};
```

## See also [​](\#see-also "Direct link to See also")

- [`staticFile()`](/docs/staticfile)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/static-file-remote-urls.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
staticFile() relative paths](/docs/staticfile-relative-paths) [Next\
\
Flickering with background-image](/docs/troubleshooting/background-image)

- [See also](#see-also)
