---
title: deleteStaticFile()v4.0.154
source: Remotion Documentation
last_updated: 2024-12-22
---

# deleteStaticFile()v4.0.154

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- deleteStaticFile()

On this page

# deleteStaticFile() [v4.0.154](https://github.com/remotion-dev/remotion/releases/v4.0.154)

Deletes a file from the [`public` directory](/docs/terminology/public-dir).

This API is useful for building interactive experiences in the [Remotion Studio](/docs/terminology/studio).

## Examples [​](\#examples "Direct link to Examples")

```

Delete 'video.webm'
tsx

import React, { useCallback } from "react";
import { deleteStaticFile } from "@remotion/studio";

export const DeleteStaticFileComp: React.FC = () => {
  const deleteFile = useCallback(async () => {
    const { existed } = await deleteStaticFile("video.webm");

    console.log(`Deleted file (${existed ? "existed" : "did not exist"})`);
  }, []);

  return <button onClick={deleteFile}>Delete</button>;
};
```

```

Delete 'video.webm'
tsx

import React, { useCallback } from "react";
import { deleteStaticFile } from "@remotion/studio";

export const DeleteStaticFileComp: React.FC = () => {
  const deleteFile = useCallback(async () => {
    const { existed } = await deleteStaticFile("video.webm");

    console.log(`Deleted file (${existed ? "existed" : "did not exist"})`);
  }, []);

  return <button onClick={deleteFile}>Delete</button>;
};
```

## Rules [​](\#rules "Direct link to Rules")

[1](#1)

This API can only be used while in the Remotion Studio.

[2](#2)

The file path must be relative to the [`public` directory](/docs/terminology/public-dir).

[3](#3)

It's not allowed to delete a file outside the [`public` directory](/docs/terminology/public-dir).

[4](#4) To delete a file in a subfolder, use forward slashes `
/
` even on Windows.
[5](#5) You can, but don't have to wrap the file path in a [`staticFile()`](/docs/staticfile) function.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/delete-static-file.ts)
- [`writeStaticFile()`](/docs/studio/write-static-file)
- [`getStaticFiles()`](/docs/studio/get-static-files)
- [`watchStaticFile()`](/docs/studio/watch-static-file)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/delete-static-file.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
writeStaticFile()](/docs/studio/write-static-file) [Next\
\
restartStudio()](/docs/studio/restart-studio)

- [Examples](#examples)
- [Rules](#rules)
- [See also](#see-also)
