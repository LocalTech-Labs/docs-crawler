---
title: Render all compositions
source: Remotion Documentation
last_updated: 2024-12-22
---

# Render all compositions

- [Home page](/)
- [Rendering](/docs/render)
- Render all compositions

On this page

# Render all compositions

In some scenarios, you might find it useful to render all compositions.

## Via CLI [​](\#via-cli "Direct link to Via CLI")

You can combine the [`npx remotion compositions`](/docs/cli/compositions) command with a bash loop:

```

render-all.sh
sh

compositions=($(npx remotion compositions src/index.ts -q))

for composition in "${compositions[@]}"
do
  npx remotion render src/index.ts $composition $composition.mp4
done
```

```

render-all.sh
sh

compositions=($(npx remotion compositions src/index.ts -q))

for composition in "${compositions[@]}"
do
  npx remotion render src/index.ts $composition $composition.mp4
done
```

You can execute it using:

```

sh

sh ./render-all.sh
```

```

sh

sh ./render-all.sh
```

note

This only works on UNIX-based systems (Linux, macOS) and on WSL in Windows.

## Via Node.JS script [​](\#via-nodejs-script "Direct link to Via Node.JS script")

You can create a script that fits you using the [Node.JS APIs](/docs/renderer). Below is an example

```

render-all.mjs
ts

import { bundle } from "@remotion/bundler";
import { getCompositions, renderMedia } from "@remotion/renderer";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

const bundled = await bundle({
  entryPoint: require.resolve("./src/index.ts"),
  // If you have a Webpack override, make sure to add it here
  webpackOverride: (config) => config,
});

const compositions = await getCompositions(bundled);

for (const composition of compositions) {
  console.log(`Rendering ${composition.id}...`);
  await renderMedia({
    codec: "h264",
    composition,
    serveUrl: bundled,
    outputLocation: `out/${composition.id}.mp4`,
  });
}
```

```

render-all.mjs
ts

import { bundle } from "@remotion/bundler";
import { getCompositions, renderMedia } from "@remotion/renderer";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

const bundled = await bundle({
  entryPoint: require.resolve("./src/index.ts"),
  // If you have a Webpack override, make sure to add it here
  webpackOverride: (config) => config,
});

const compositions = await getCompositions(bundled);

for (const composition of compositions) {
  console.log(`Rendering ${composition.id}...`);
  await renderMedia({
    codec: "h264",
    composition,
    serveUrl: bundled,
    outputLocation: `out/${composition.id}.mp4`,
  });
}
```

```

Execute
bash

node render-all.mjs
```

```

Execute
bash

node render-all.mjs
```

## See also [​](\#see-also "Direct link to See also")

- [Render your video](/docs/render)
- [@remotion/renderer](/docs/renderer)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/render-all.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Render a dataset](/docs/dataset-render) [Next\
\
Video formats](/docs/miscellaneous/video-formats)

- [Via CLI](#via-cli)
- [Via Node.JS script](#via-nodejs-script)
- [See also](#see-also)
