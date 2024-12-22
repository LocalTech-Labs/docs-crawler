---
title: Composition
source: Remotion Documentation
last_updated: 2024-12-22
---

# Composition

- [Home page](/)
- [Terminology](/docs/terminology)
- Composition

On this page

# Composition

A composition is something you can render. It consists of:

[1](#1)

a React component

[2](#2) the canvas width and height

[3](#3) a "frames per second" (FPS) value

[4](#4) a duration value

[5](#5) an identifier `id`

It can be registered in the [Remotion Studio](/docs/terminology/studio) by rendering a [`<Composition>`](/docs/composition) component.

A composition with a duration of 1 frame is also called a [`<Still>`](/docs/still).

In the [Remotion Player](/docs/terminology/player), you don't use the `<Composition>` component, rather you pass the component and metadata directly to the [`<Player>`](/docs/player/player).

Not to be confused with [`<Sequence>`](/docs/terminology/sequence).

## Composition ID [​](\#composition-id "Direct link to Composition ID")

The string that you pass as the `id` prop to the [`<Composition>` component](/docs/composition).

You need the composition ID to reference what you would like to render, for example: `npx remotion render src/index <composition-id>`.

## See [​](\#see "Direct link to See")

- [Defining compositions](/docs/the-fundamentals#compositions)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/composition.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Sequence](/docs/terminology/sequence) [Next\
\
List of resources](/docs/resources)

- [Composition ID](#composition-id)
- [See](#see)
