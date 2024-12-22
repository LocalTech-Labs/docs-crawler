---
title: Turn a <Player> into a Remotion project
source: Remotion Documentation
last_updated: 2024-12-22
---

# Turn a <Player> into a Remotion project

- [Home page](/)
- Building apps
- Turn a <Player> into a Remotion project

On this page

# Turn a <Player> into a Remotion project

If you have a React app that uses the [`<Player>`](/docs/player/player) component, you can turn it also into a Remotion project without having to create a new repository. This will enable you to:

[1](#1)

Use the [Remotion Studio](/docs/studio/)

[2](#2) Render videos locally

[2](#2) Create a [Remotion Bundle](/docs/terminology/bundle) that allows rendering in the cloud

## Instructions [​](\#instructions "Direct link to Instructions")

[1](#1)

Ensure that you don't already have a Studio installed - in the following
templates, the Studio is already installed and can be started using the command [`npx remotion studio`](/docs/cli/studio)
:

- [Next.js (App dir)](/templates/next)
- [Next.js (App dir + TailwindCSS)](/templates/next-tailwind)
- [Next.js (Pages dir)](/templates/next-pages-dir)
- [Remix](/templates/remix)

[2](#2)

Install the Remotion CLI:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/cli@4.0.242Copy
```

```

npm i --save-exact @remotion/cli@4.0.242Copy
```

```

pnpm i @remotion/cli@4.0.242Copy
```

```

pnpm i @remotion/cli@4.0.242Copy
```

```

bun i @remotion/cli@4.0.242Copy
```

```

bun i @remotion/cli@4.0.242Copy
```

```

yarn --exact add @remotion/cli@4.0.242Copy
```

```

yarn --exact add @remotion/cli@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

[3](#3)

Make a new folder `remotion` and add these two files in
it:

```

remotion/Root.tsx
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComposition} from './Composition';

export const RemotionRoot: React.FC = () => {
  return (
    <>
      <Composition
        id="Empty"
        // Import the component and add the properties you had in the `<Player>` before
        component={MyComposition}
        durationInFrames={60}
        fps={30}
        width={1280}
        height={720}
      />
    </>
  );
};
```

```

remotion/Root.tsx
tsx

import React from 'react';
import {Composition} from 'remotion';
import {MyComposition} from './Composition';

export const RemotionRoot: React.FC = () => {
  return (
    <>
      <Composition
        id="Empty"
        // Import the component and add the properties you had in the `<Player>` before
        component={MyComposition}
        durationInFrames={60}
        fps={30}
        width={1280}
        height={720}
      />
    </>
  );
};
```

```

remotion/index.ts
ts

import { registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

registerRoot(RemotionRoot);
```

```

remotion/index.ts
ts

import { registerRoot } from "remotion";
import { RemotionRoot } from "./Root";

registerRoot(RemotionRoot);
```

Add the component you previously registered in the `<Player>` also to the `<Composition>`.

If necessary, move the components into the `remotion` directory.

The file that calls [`registerRoot()`](/docs/register-root) is now your Remotion [**entry point**](/docs/terminology/entry-point).

[4](#4)

Apply common Webpack overrides that are commonly enabled in React
apps:

- [TailwindCSS](/docs/tailwind/tailwind)
- [SCSS](/docs/enable-scss/overview)
- [TypeScript Aliases](/docs/typescript-aliases#automatically-syncing-webpack-and-typescript-aliases)

## Start the Remotion Studio: [​](\#start-the-remotion-studio "Direct link to Start the Remotion Studio:")

```

npx remotion studio
```

```

npx remotion studio
```

## Render a video [​](\#render-a-video "Direct link to Render a video")

Render our a video using

```

npx remotion render remotion/index.ts
```

```

npx remotion render remotion/index.ts
```

## Set up server-side rendering [​](\#set-up-server-side-rendering "Direct link to Set up server-side rendering")

See [server-side rendering](/docs/ssr) for more information.

## How do I download a video from the Remotion Player? [​](\#how-do-i-download-a-video-from-the-remotion-player "Direct link to How do I download a video from the Remotion Player?")

The video first needs to get rendered - this can only be done on the server. See [ways to render](/docs/render) and [server-side rendering](/docs/ssr) for more information.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player-into-remotion-project.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Convert Studio to an app](/docs/studio-into-app) [Next\
\
Handling user video uploads](/docs/video-uploads)

- [Instructions](#instructions)
- [Start the Remotion Studio:](#start-the-remotion-studio)
- [Render a video](#render-a-video)
- [Set up server-side rendering](#set-up-server-side-rendering)
- [How do I download a video from the Remotion Player?](#how-do-i-download-a-video-from-the-remotion-player)
