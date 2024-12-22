---
title: Start editing
source: Remotion Documentation
last_updated: 2024-12-22
---

# Start editing

- [Home page](/)
- Recorder
- Editing
- Start editing

On this page

# Start editing

To start the editing environment, run

```

bun start
```

```

bun start
```

This will start the development server and open a Remotion Studio instance in a new window.

## Create a new composition [​](\#create-a-new-composition "Direct link to Create a new composition")

A composition is a video that you can edit and render.

In the left sidebar, create a new composition by right-clicking the `empty` composition and selecting `Duplicate` and give it an ID.

The ID should be the same as the name of the subfolder of `public` where you store your videos.
**Example**: If you store your recordings in `public/hello-world`, the ID should be `hello-world`.

## Add new scenes to the composition [​](\#add-new-scenes-to-the-composition "Direct link to Add new scenes to the composition")

Open the right sidebar. By default there are no scenes.

Click the `+` button to add a new scene.

![](/recorder/add-scene.png)

By default, it will add a video scene, which will consume the first recording in the `public/[composition-id]` folder.

You can also switch the scene type to something different like `endcard` or `title`.

Each scene type has its own settings which you can configure in the right sidebar.

To save the scenes, click the disk icon or use `Ctrl+S` / `Cmd+S`.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/editing.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Generate captions](/docs/recorder/captions) [Next\
\
Silence removal](/docs/recorder/editing/silence-removal)

- [Create a new composition](#create-a-new-composition)
- [Add new scenes to the composition](#add-new-scenes-to-the-composition)
