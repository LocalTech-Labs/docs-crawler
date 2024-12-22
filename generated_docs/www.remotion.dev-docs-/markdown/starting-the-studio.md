---
title: Starting the Studio
source: Remotion Documentation
last_updated: 2024-12-22
---

# Starting the Studio

- [Home page](/)
- Studio
- Starting the Studio

On this page

# Starting the Studio

Using the Remotion Studio, you can preview your video, and if a server is connected, even render the video.

## Prerequisites [​](\#prerequisites "Direct link to Prerequisites")

The Remotion CLI is required for this guide.

Most templates have it out of the box, but you can install it by running the following command in your terminal:

- npm
- pnpm
- yarn
- bun

```

bash

npm i @remotion/cli
```

```

bash

npm i @remotion/cli
```

```

bash

pnpm i @remotion/cli
```

```

bash

pnpm i @remotion/cli
```

```

bash

yarn add @remotion/cli
```

```

bash

yarn add @remotion/cli
```

```

bash

bun i @remotion/cli
```

```

bash

bun i @remotion/cli
```

## Launching the Studio [​](\#launching-the-studio "Direct link to Launching the Studio")

You can start the Remotion Studio by running the following command in your terminal:

- Regular templates
- Next.js and Remix templates

```

bash

npm start
```

```

bash

npm start
```

```

bash

npm run remotion
```

```

bash

npm run remotion
```

This is a shorthand for the [`studio`](/docs/cli/studio) command of the [Remotion CLI](/docs/cli):

```

bash

npx remotion studio
```

```

bash

npx remotion studio
```

See the available [options here](/docs/cli/studio).

A server will be started on port `3000` (or a higher port if it isn't available) and the Remotion Studio should open in the browser.

![](/img/timeline.png)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/studio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Hardware acceleration](/docs/hardware-acceleration) [Next\
\
Keyboard navigation](/docs/studio/shortcuts)

- [Prerequisites](#prerequisites)
- [Launching the Studio](#launching-the-studio)
