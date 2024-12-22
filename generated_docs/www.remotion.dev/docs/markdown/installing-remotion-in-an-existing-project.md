---
title: Installing Remotion in an existing project
source: Remotion Documentation
last_updated: 2024-12-21
---

# Installing Remotion in an existing project

- [Home page](/)
- Building apps
- Installation in existing project

On this page

# Installing Remotion in an existing project

Remotion can be installed into existing projects, such as [Next.JS](https://nextjs.org/), [Remix](https://remix.run/), [Vite](https://vitejs.dev/guide/) or [Create React App](https://create-react-app.dev/), as well as server-only projects that run on Node.JS. Get started by adding the following packages:

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

npm i --save-exact remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

pnpm i remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

pnpm i remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

bun i remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

bun i remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

yarn --exact add remotion@4.0.241 @remotion/cli@4.0.241Copy
```

```

yarn --exact add remotion@4.0.241 @remotion/cli@4.0.241Copy
```

This assumes you are currently using v4.0.241 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

- If you'd like to embed a Remotion video in your existing React app, install [`@remotion/player`](/docs/player/installation) as well.
- If you'd like to render a video using the Node.JS APIs, install [`@remotion/renderer`](/docs/renderer) as well.
- If you'd like to trigger a render on Remotion Lambda, install [`@remotion/lambda`](/docs/lambda/setup) as well.

## Setting up the folder structure [​](\#setting-up-the-folder-structure "Direct link to Setting up the folder structure")

Create a new folder for the Remotion files. It can be anywhere and assume any name, in this example we name it `remotion` and put it in the root of our project. Inside the folder you created, create 3 files:

```

remotion/Composition.tsx
tsx

export const MyComposition = () => {
  return null;
};
```

```

remotion/Composition.tsx
tsx

export const MyComposition = () => {
  return null;
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

The file that calls [`registerRoot()`](/docs/register-root) is now your Remotion [**entry point**](/docs/terminology/entry-point).

note

Watch out for import aliases in your `tsconfig.json` that will resolve `import {...} from "remotion"` to the `remotion` folder. We recommend to not use `paths` without a prefix.

## Starting the Studio [​](\#starting-the-studio "Direct link to Starting the Studio")

Start the Remotion Studio using the following command:

```

npx remotion studio remotion/index.ts
```

```

npx remotion studio remotion/index.ts
```

Replace `remotion/index.ts` with your entrypoint if necessary.

## Render a video [​](\#render-a-video "Direct link to Render a video")

Render our a video using

```

npx remotion render remotion/index.ts MyComp out.mp4
```

```

npx remotion render remotion/index.ts MyComp out.mp4
```

Replace the entrypoint, composition name and output filename with the values that correspond to your usecase.

## Install the ESLint Plugin [​](\#install-the-eslint-plugin "Direct link to Install the ESLint Plugin")

Remotion has an ESLint plugin that warns about improper usage of Remotion APIs. To add it to your existing project, install it:

- npm
- yarn
- pnpm

```

bash

npm i @remotion/eslint-plugin
```

```

bash

npm i @remotion/eslint-plugin
```

```

bash

yarn add @remotion/eslint-plugin
```

```

bash

yarn add @remotion/eslint-plugin
```

```

bash

pnpm i @remotion/eslint-plugin
```

```

bash

pnpm i @remotion/eslint-plugin
```

This snippet will enable the recommended rules only for the Remotion files:

```

.eslintrc
json

{
  "plugins": ["@remotion"],
  "overrides": [
    {
      "files": ["remotion/*.{ts,tsx}"],
      "extends": ["plugin:@remotion/recommended"]
    }
  ]
}
```

```

.eslintrc
json

{
  "plugins": ["@remotion"],
  "overrides": [
    {
      "files": ["remotion/*.{ts,tsx}"],
      "extends": ["plugin:@remotion/recommended"]
    }
  ]
}
```

## Embed a Remotion video into your React app [​](\#embed-a-remotion-video-into-your-react-app "Direct link to Embed a Remotion video into your React app")

You can use the `<Player>` component to display a Remotion video in your React project. Read the [separate page](/docs/player/integration) about it for instructions.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/brownfield-installation.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Previous\
\
Telemetry in @remotion/webcodecs](/docs/webcodecs/telemetry) [Next\
\
Convert Studio to an app](/docs/studio-into-app)

- [Setting up the folder structure](#setting-up-the-folder-structure)
- [Starting the Studio](#starting-the-studio)
- [Render a video](#render-a-video)
- [Install the ESLint Plugin](#install-the-eslint-plugin)
- [Embed a Remotion video into your React app](#embed-a-remotion-video-into-your-react-app)
