---
title: B-Roll
source: Remotion Documentation
last_updated: 2024-12-22
---

# B-Roll

- [Home page](/)
- Recorder
- Editing
- B-Roll

On this page

# B-Roll

You can overlay your recorded scenes with additional video or photo footage by adding `bRoll` elements in the right sidebar.

Add the footage you want to use to the `public` folder and select it from the `bRoll` drop-down menu for the scene you want to overlay.

Configure when and for how long to play an asset by adjusting the `from` and `durationInFrames` values.
The `from` value is relative to the scene to which you are adding the B-roll and defines the frame from which the asset is to be displayed.

## Constraints [​](\#constraints "Direct link to Constraints")

The data that you add in bRoll is postprocessed to enforce the following rules:

1. If a second B-Roll overlays the first B-Roll, the first B-Roll may not disappear before the second B-Roll disappears.
2. A B-Roll must have disappeared before the transition to the next scene begins.

These constraints are defined in [`remotion/scenes/BRoll/apply-b-roll-rules.ts`](https://github.com/remotion-dev/recorder/tree/main/remotion/scenes/BRoll/apply-b-roll-rules.ts).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/b-roll.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Adding music](/docs/recorder/editing/music) [Next\
\
Cutting clips](/docs/recorder/editing/cutting-clips)

- [Constraints](#constraints)
