---
title: Using Remotion in Vue
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using Remotion in Vue

- [Home page](/)
- Building apps
- Integration into Vue

On this page

# Using Remotion in Vue

This guide explains how to integrate Remotion into a **Vue.js** project.

## Install required packages [​](\#install-required-packages "Direct link to Install required packages")

Install Remotion and necessary dependencies:

- npm
- pnpm
- yarn
- bun

```

Use npm as the package manager
bash

npm i remotion @remotion/player @remotion/cli react react-dom
npm i --save-dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use npm as the package manager
bash

npm i remotion @remotion/player @remotion/cli react react-dom
npm i --save-dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use pnpm as the package manager
bash

pnpm i remotion @remotion/player @remotion/cli react react-dom
pnpm i --dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use pnpm as the package manager
bash

pnpm i remotion @remotion/player @remotion/cli react react-dom
pnpm i --dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use Yarn as the package manager
bash

yarn add remotion @remotion/player @remotion/cli react react-dom
yarn add --dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use Yarn as the package manager
bash

yarn add remotion @remotion/player @remotion/cli react react-dom
yarn add --dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use Bun as the package manager and runtime
bash

bun i remotion @remotion/player @remotion/cli react react-dom
bun --dev @types/react @types/react-dom @vitejs/plugin-react
```

```

Use Bun as the package manager and runtime
bash

bun i remotion @remotion/player @remotion/cli react react-dom
bun --dev @types/react @types/react-dom @vitejs/plugin-react
```

note

Bun as a runtime is mostly supported. [Read more here](/docs/bun).

## Create a Remotion folder [​](\#create-a-remotion-folder "Direct link to Create a Remotion folder")

For better separation, create a folder to hold your Remotion files:

```

plaintext

src/remotion
```

```

plaintext

src/remotion
```

Copy the contents of your Remotion project or a starter template (e.g., HelloWorld) into this new folder. This will help separate Remotion related files from the rest of your VueJS codebase.

## Copy `remotion.config.ts` [​](\#copy-remotionconfigts "Direct link to copy-remotionconfigts")

Copy the `remotion.config.ts` file to the root directory of your VueJS project, placing it at the same level as `package.json`.

This configuration file is necessary for Remotion to recognize and compile your project settings.

## Configure TypeScript for JSX [​](\#configure-typescript-for-jsx "Direct link to Configure TypeScript for JSX")

To enable JSX support in VueJS, update the `tsconfig.app.json` file by setting `"jsx": "react"` under `compilerOptions`. This configuration allows VueJS to interpret JSX syntax used in Remotion's React components.

```

tsconfig.json
json

{
  "compilerOptions": {
    "jsx": "react",
    "jsxImportSource": "" // Recommended to avoid errors while building
    // other options
  }
}
```

```

tsconfig.json
json

{
  "compilerOptions": {
    "jsx": "react",
    "jsxImportSource": "" // Recommended to avoid errors while building
    // other options
  }
}
```

note

The "jsxImportSource": "" setting is recommended to prevent build issues.

## Configure Vite for React [​](\#configure-vite-for-react "Direct link to Configure Vite for React")

To enable React support in Vite, update the `vite.config.ts` file by adding `react` to the plugins.

```

vite.config.ts
ts

import {fileURLToPath, URL} from 'node:url';
import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [
    vue(),
    react({
      include: /\.(jsx|tsx)$/,
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
});
```

```

vite.config.ts
ts

import {fileURLToPath, URL} from 'node:url';
import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [
    vue(),
    react({
      include: /\.(jsx|tsx)$/,
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
});
```

## Create a React wrapper component for VueJS [​](\#create-a-react-wrapper-component-for-vuejs "Direct link to Create a React wrapper component for VueJS")

To embed Remotion components in VueJS, create a **wrapper component**:

1. In your `remotion` folder, create a file named `PlayerViewWrapper.vue`.

```

PlayerViewWrapper.vue
html

<script lang="ts" setup>
  import {onBeforeUnmount, onMounted, ref, watch} from 'vue';
  import {type PlayerSchema, PlayerView} from '@/remotion/PlayerView';
  import {createRoot, type Root} from 'react-dom/client';
  import * as React from 'react';

  interface PlayerData {
    data: PlayerSchema;
  }

  const playerRef = ref();
  const reactRoot = ref();
  let root: Root;
  const emit = defineEmits(['paused']);

  const props = defineProps<PlayerData>();

  onMounted(() => {
    root = createRoot(reactRoot.value);
    render();
  });

  function render() {
    root.render(
      React.createElement(PlayerView, {
        data: props.data,
        playerRefInstance: playerRef,
        onPaused: () => emit('paused'),
      }),
    );
  }

  watch(
    () => playerRef,
    (newData, oldData) => {
      playerRef.value.play();
    },
    {deep: true},
  );

  watch(
    () => props.data,
    (newData, oldData) => {
      render();
    },
    {deep: true},
  );

  onBeforeUnmount(() => {
    if (root) {
      root.unmount();
    }
  });
</script>

<template>
  <div ref="reactRoot"></div>
</template>
```

```

PlayerViewWrapper.vue
html

<script lang="ts" setup>
  import {onBeforeUnmount, onMounted, ref, watch} from 'vue';
  import {type PlayerSchema, PlayerView} from '@/remotion/PlayerView';
  import {createRoot, type Root} from 'react-dom/client';
  import * as React from 'react';

  interface PlayerData {
    data: PlayerSchema;
  }

  const playerRef = ref();
  const reactRoot = ref();
  let root: Root;
  const emit = defineEmits(['paused']);

  const props = defineProps<PlayerData>();

  onMounted(() => {
    root = createRoot(reactRoot.value);
    render();
  });

  function render() {
    root.render(
      React.createElement(PlayerView, {
        data: props.data,
        playerRefInstance: playerRef,
        onPaused: () => emit('paused'),
      }),
    );
  }

  watch(
    () => playerRef,
    (newData, oldData) => {
      playerRef.value.play();
    },
    {deep: true},
  );

  watch(
    () => props.data,
    (newData, oldData) => {
      render();
    },
    {deep: true},
  );

  onBeforeUnmount(() => {
    if (root) {
      root.unmount();
    }
  });
</script>

<template>
  <div ref="reactRoot"></div>
</template>
```

This wrapper component will serve as the bridge between VueJs and Remotion’s React components.

## Create a wrapper for the Remotion Player [​](\#create-a-wrapper-for-the-remotion-player "Direct link to Create a wrapper for the Remotion Player")

1. In your `remotion` folder, create a file named `PlayerView.tsx`.
2. Ensure each `.tsx` file imports React explicitly at the top of the file:

This will get the reference to the player with `createRef`.

```

PlayerView.tsx
tsx

import React, {useEffect} from 'react';
import {Player, type PlayerRef} from '@remotion/player';
import {HelloWorld} from '@/remotion/HelloWorld';
import type {Ref} from 'vue';

export interface PlayerSchema {
  titleText: string;
  titleColor: string;
  logoColor1: string;
  logoColor2: string;
}

export const PlayerView = ({
  data,
  playerRefInstance,
  onPaused,
}: {
  data: PlayerSchema;
  playerRefInstance: Ref<PlayerRef>;
  onPaused?: () => void;
}): React.ReactElement => {
  const playerRef: React.RefObject<PlayerRef> = React.createRef();

  useEffect(() => {
    if (playerRef.current) {
      playerRefInstance.value = playerRef.current;

      // add callback when player pauses
      playerRef.current.addEventListener('pause', () => {
        onPaused?.();
      });
    }
  }, []);

  return (
    <Player
      ref={playerRef}
      component={HelloWorld}
      durationInFrames={150}
      fps={30}
      compositionHeight={1080}
      compositionWidth={1920}
      inputProps={data}
      controls
    />
  );
};
```

```

PlayerView.tsx
tsx

import React, {useEffect} from 'react';
import {Player, type PlayerRef} from '@remotion/player';
import {HelloWorld} from '@/remotion/HelloWorld';
import type {Ref} from 'vue';

export interface PlayerSchema {
  titleText: string;
  titleColor: string;
  logoColor1: string;
  logoColor2: string;
}

export const PlayerView = ({
  data,
  playerRefInstance,
  onPaused,
}: {
  data: PlayerSchema;
  playerRefInstance: Ref<PlayerRef>;
  onPaused?: () => void;
}): React.ReactElement => {
  const playerRef: React.RefObject<PlayerRef> = React.createRef();

  useEffect(() => {
    if (playerRef.current) {
      playerRefInstance.value = playerRef.current;

      // add callback when player pauses
      playerRef.current.addEventListener('pause', () => {
        onPaused?.();
      });
    }
  }, []);

  return (
    <Player
      ref={playerRef}
      component={HelloWorld}
      durationInFrames={150}
      fps={30}
      compositionHeight={1080}
      compositionWidth={1920}
      inputProps={data}
      controls
    />
  );
};
```

## Use the component in VueJS [​](\#use-the-component-in-vuejs "Direct link to Use the component in VueJS")

To display the Remotion player within an VueJS template, add your new wrapper component wherever you'd like the Player to appear:

```

html

<PlayerWrapper @paused="onPaused" :data="data" />
```

```

html

<PlayerWrapper @paused="onPaused" :data="data" />
```

This VueJS component tag will render the Remotion Player, allowing you to pass data or configuration as needed through VueJS's data binding.

You're now able to use the [`API`](/docs/player/player) of the player via the `player` ref binding

## See also [​](\#see-also "Direct link to See also")

- [Source code for this documentation](https://github.com/remotion-dev/vue-starter)
- [Integration into Angular](/docs/angular)
- [Integration into Svelte](/docs/svelte)

CONTRIBUTORS

[![dothem1337](https://github.com/dothem1337.png)\
\
**dothem1337** \
\
Author](https://github.com/dothem1337)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/vue.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Integration into Svelte](/docs/svelte) [Next\
\
TailwindCSS](/docs/tailwind)

- [Install required packages](#install-required-packages)
- [Create a Remotion folder](#create-a-remotion-folder)
- [Copy `remotion.config.ts`](#copy-remotionconfigts)
- [Configure TypeScript for JSX](#configure-typescript-for-jsx)
- [Configure Vite for React](#configure-vite-for-react)
- [Create a React wrapper component for VueJS](#create-a-react-wrapper-component-for-vuejs)
- [Create a wrapper for the Remotion Player](#create-a-wrapper-for-the-remotion-player)
- [Use the component in VueJS](#use-the-component-in-vuejs)
- [See also](#see-also)
