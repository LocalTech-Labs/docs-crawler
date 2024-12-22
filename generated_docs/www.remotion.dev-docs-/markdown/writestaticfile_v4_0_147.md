---
title: writeStaticFile()v4.0.147
source: Remotion Documentation
last_updated: 2024-12-22
---

# writeStaticFile()v4.0.147

- [Home page](/)
- [@remotion/studio](/docs/studio/api)
- writeStaticFile()

On this page

# writeStaticFile() [v4.0.147](https://github.com/remotion-dev/remotion/releases/v4.0.147)

Saves some content into a file in the [`public` directory](/docs/).

This API is useful for building interactive experiences in the [Remotion Studio](/docs/terminology/studio).

## Examples [​](\#examples "Direct link to Examples")

```

Write 'Hello world' to public/file.txt
tsx

import React, { useCallback } from "react";
import { writeStaticFile } from "@remotion/studio";

export const WriteStaticFileComp: React.FC = () => {
  const saveFile = useCallback(async () => {
    await writeStaticFile({
      filePath: "file.txt",
      contents: "Hello world",
    });

    console.log("Saved!");
  }, []);

  return <button onClick={saveFile}>Save</button>;
};
```

```

Write 'Hello world' to public/file.txt
tsx

import React, { useCallback } from "react";
import { writeStaticFile } from "@remotion/studio";

export const WriteStaticFileComp: React.FC = () => {
  const saveFile = useCallback(async () => {
    await writeStaticFile({
      filePath: "file.txt",
      contents: "Hello world",
    });

    console.log("Saved!");
  }, []);

  return <button onClick={saveFile}>Save</button>;
};
```

```

Allow a file upload
tsx

import React, { useCallback } from "react";
import { writeStaticFile } from "@remotion/studio";

export const WriteStaticFileComp: React.FC = () => {
  const saveFile = useCallback(
    async (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files![0];

      await writeStaticFile({
        filePath: file.name,
        contents: await file.arrayBuffer(),
      });

      console.log("Saved!");
    },
    [],
  );

  return <input type="file" onChange={saveFile} />;
};
```

```

Allow a file upload
tsx

import React, { useCallback } from "react";
import { writeStaticFile } from "@remotion/studio";

export const WriteStaticFileComp: React.FC = () => {
  const saveFile = useCallback(
    async (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files![0];

      await writeStaticFile({
        filePath: file.name,
        contents: await file.arrayBuffer(),
      });

      console.log("Saved!");
    },
    [],
  );

  return <input type="file" onChange={saveFile} />;
};
```

## Rules [​](\#rules "Direct link to Rules")

[1](#1)

This API can only be used while in the Remotion Studio.

[2](#2)

The file path must be relative to the [`public` directory](/docs/terminology/public-dir).

[3](#3)

It's not allowed to write outside the [`public` directory](/docs/terminology/public-dir).

[4](#4) To write into subfolders, use forward slashes `/` even
on Windows.
[5](#5) You can pass a `string` or `ArrayBuffer`.

## See also [​](\#see-also "Direct link to See also")

- [Source code for this function](https://github.com/remotion-dev/remotion/blob/main/packages/studio/src/api/write-static-file.ts)
- [`staticFile()`](/docs/staticfile)
- [`getStaticFiles()`](/docs/studio/get-static-files)
- [`watchStaticFile()`](/docs/studio/watch-static-file)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/write-static-file.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
watchStaticFile()](/docs/studio/watch-static-file) [Next\
\
deleteStaticFile()](/docs/studio/delete-static-file)

- [Examples](#examples)
- [Rules](#rules)
- [See also](#see-also)
