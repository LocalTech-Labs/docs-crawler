---
title: How props get resolved
source: Remotion Documentation
last_updated: 2024-12-22
---

# How props get resolved

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- How props get resolved

On this page

# How props get resolved

## During rendering [​](\#during-rendering "Direct link to During rendering")

Remotion performs an algorithm for determining the props that are being passed to your component when rendering the video.
Three factors play a role:

[1](#1)

**Default props** is the fallback data if no props are passed to the render. You can specify them using the [`defaultProps`](/docs/composition#defaultprops) property of your [`<Composition />`](/docs/composition#defaultprops).

[2](#2)

**Input props** is the data being passed while invoking a render, either via the [`inputProps`](/docs/renderer/render-media#inputprops) option, the [`--props` flag](/docs/cli/render#--props) or using the render dialog in the Remotion Studio.

[3](#3)

[`calculateMetadata()`](/docs/composition#calculatemetadata) may be used to dynamically transform the props, as well as the metadata of the composition.

The following diagram shows how the props get resolved:

## In the Remotion Studio [​](\#in-the-remotion-studio "Direct link to In the Remotion Studio")

In the Remotion Studio, the props are resolved in a similar way, but with a **few differences**:

[1](#1)

The default props can be **edited in the right sidebar**. Invalid modifications will be marked with a red outline and do not apply.

[2](#2) If you render a video using the Render button, the **input props form gets pre-propagated** with the default props, including modifications in the right sidebar.

The following rules **stay the same**, which you should be aware of:

[1](#1)

If you start the Studio with the [`--props`](/docs/cli/studio#--props) flag, this data will take priority over the default props, including modifications in the sidebar. It is not recommended to pass input props to the Studio.

[2](#2)

The passed input props may get transformed by [`calculateMetadata()`](/docs/composition#calculatemetadata).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/how-props-flow.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Variable duration](/docs/dynamic-metadata) [Next\
\
Render your video](/docs/render)

- [During rendering](#during-rendering)
- [In the Remotion Studio](#in-the-remotion-studio)
