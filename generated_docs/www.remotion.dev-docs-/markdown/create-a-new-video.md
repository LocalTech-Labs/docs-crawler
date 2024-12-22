---
title: Create a new video
source: Remotion Documentation
last_updated: 2024-12-22
---

# Create a new video

- [Home page](/)
- Recorder
- Create a new video

On this page

# Create a new video

If you haven't done so yet, start the Recorder:

```

bash

bun start
```

```

bash

bun start
```

Remotion calls each output video a "Composition". There are two ways to generate a new composition:

## 1\. Visual way [​](\#1-visual-way "Direct link to 1. Visual way")

Open [`http://localhost:3000`](https://localhost:3000) and in the left sidebar, right-click on the "empty" composition. Select "Duplicate" and give your composition an ID. Confirm.

## 2\. Manual way [​](\#2-manual-way "Direct link to 2. Manual way")

Open the project in a code editor. You will see an empty [`<Composition>`](/docs/terminology/composition) under `remotion/Root.tsx`:

```

tsx

<Composition
  component={Main}
  id="empty"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
    platform: "x" as const,
  }}
  calculateMetadata={calcMetadata}
/>
```

```

tsx

<Composition
  component={Main}
  id="empty"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
    platform: "x" as const,
  }}
  calculateMetadata={calcMetadata}
/>
```

Copy and paste the [`<Composition />`](/docs/composition) component to duplicate it and give it an [`id`](/docs/terminology/composition#composition-id) that is different than `"empty"`. For example:

```

remotion/Root.tsx
tsx

<Composition
  component={Main}
  id="empty"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
  }}
  calculateMetadata={calcMetadata}
/>
<Composition
  component={Main}
  id="my-video"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
  }}
  calculateMetadata={calcMetadata}
/>
```

```

remotion/Root.tsx
tsx

<Composition
  component={Main}
  id="empty"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
  }}
  calculateMetadata={calcMetadata}
/>
<Composition
  component={Main}
  id="my-video"
  schema={videoConf}
  defaultProps={{
    theme: "light" as const,
    canvasLayout: "square" as const,
    scenes: [],
    scenesAndMetadata: [],
  }}
  calculateMetadata={calcMetadata}
/>
```

Save, and you should see the new composition when you look at the Remotion Studio on [`http://localhost:3000`](https://localhost:3000).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/create.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Setup](/docs/recorder/setup) [Next\
\
Using the interface](/docs/recorder/record/)

- [1\. Visual way](#1-visual-way)
- [2\. Manual way](#2-manual-way)
