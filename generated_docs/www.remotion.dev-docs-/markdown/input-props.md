---
title: Input Props
source: Remotion Documentation
last_updated: 2024-12-22
---

# Input Props

- [Home page](/)
- [Terminology](/docs/terminology)
- Input Props

# Input Props

Input props are data that can be passed to a render in order to parametrize the video. They can be obtained in two ways:

- This data is passed as actual React props to the component that you defined in your [composition](/docs/terminology/composition)
- Using the [`getInputProps()`](/docs/get-input-props) function, you can retrieve the props even outside your component, for example to dynamically change the [duration or dimensions](/docs/dynamic-metadata).

In the [Remotion Studio](/docs/terminology/studio), you can set [default props](/docs/composition#defaultprops) to serve as placeholder data for designing your video. If your input props don't override the default props, the default props will be used. See: [How props get resolved](/docs/props-resolution)

In the [Remotion Player](/docs/terminology/player), there are no default props, but you can pass [`inputProps`](/docs/player/player#inputprops) directly to the [`<Player>`](/docs/player).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/terminology/input-props.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Concurrency](/docs/terminology/concurrency) [Next\
\
Cloud Run URL](/docs/terminology/cloud-run-url)
