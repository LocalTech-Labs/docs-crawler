---
title: Sequence
source: Remotion Documentation
last_updated: 2024-12-22
---

# Sequence

- [Home page](/)
- [Terminology](/docs/terminology)
- Sequence

On this page

# Sequence

[`<Sequence>`](/docs/sequence) is a built-in component in Remotion that:

- Absolutely positions content
- Allows you to time-shift elements

Think of it as the equivalent of a "layer" in After Effects or Photoshop.

Not to be confused with [`<Composition>`](/docs/terminology/composition).

## Layouting [​](\#layouting "Direct link to Layouting")

A [`<Sequence>`](/docs/sequence) by default will also be absolutely positioned in the DOM, so you can use it to overlay elements on top of each other.

## Time-shifting [​](\#time-shifting "Direct link to Time-shifting")

Consider an element which starts animating at frame 0.

If you would like to delay the animation, rather than refactoring the animation, you can wrap it in a `<Sequence>` and define a delay using the [`from`](/docs/sequence#from) prop.

Using a sequence, you can also trim the start and end of an animation by passing the [`durationInFrames`](/docs/sequence#durationinframes) prop.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/sequence.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Remotion Bundle](/docs/terminology/bundle) [Next\
\
Composition](/docs/terminology/composition)

- [Layouting](#layouting)
- [Time-shifting](#time-shifting)
