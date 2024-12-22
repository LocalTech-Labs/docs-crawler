---
title: #t= Media Fragments
source: Remotion Documentation
last_updated: 2024-12-22
---

# #t= Media Fragments

- [Home page](/)
- Miscellaneous
- #t= Media Fragments

On this page

# \#t= Media Fragments

Remotion by default adds [Media fragment](https://www.w3.org/TR/media-frags/) hashes to video URLs.

## Behavior [​](\#behavior "Direct link to Behavior")

For example, the following Remotion code:

```

MyComp.tsx
tsx

import {Sequence, OffthreadVideo, useVideoConfig} from 'remotion';

export const MyComp: React.FC = () => {
  const {fps} = useVideoConfig();

  return (
    <Sequence from={2 * fps} durationInFrames={4 * fps}>
      <OffthreadVideo src="https://example.com/bbb.mp4" />
    </Sequence>
  );
};
```

```

MyComp.tsx
tsx

import {Sequence, OffthreadVideo, useVideoConfig} from 'remotion';

export const MyComp: React.FC = () => {
  const {fps} = useVideoConfig();

  return (
    <Sequence from={2 * fps} durationInFrames={4 * fps}>
      <OffthreadVideo src="https://example.com/bbb.mp4" />
    </Sequence>
  );
};
```

results in the following `<video>` tag being mounted:

```

tsx

<video src="https://example.com/bbb.mp4#t=2.0,4.0" />
```

```

tsx

<video src="https://example.com/bbb.mp4#t=2.0,4.0" />
```

This instructs browsers to only load the section of 00:02 to 00:04 of the video.

## Upsides [​](\#upsides "Direct link to Upsides")

It improves playback performance and reduces bandwidth usage.

On Safari, these hints have higher importance. We find that without these hints, Safari videos behave badly on mobile.

## How to disable media fragment hints [​](\#how-to-disable-media-fragment-hints "Direct link to How to disable media fragment hints")

You can disable Remotion appending media fragments by appending your own hash:

```

MyComp.tsx
tsx

import {Sequence, OffthreadVideo, useVideoConfig} from 'remotion';

export const MyComp: React.FC = () => {
  const {fps} = useVideoConfig();

  return (
    <Sequence from={2 * fps} durationInFrames={4 * fps}>
      <OffthreadVideo src="https://example.com/bbb.mp4#disable" />
    </Sequence>
  );
};
```

```

MyComp.tsx
tsx

import {Sequence, OffthreadVideo, useVideoConfig} from 'remotion';

export const MyComp: React.FC = () => {
  const {fps} = useVideoConfig();

  return (
    <Sequence from={2 * fps} durationInFrames={4 * fps}>
      <OffthreadVideo src="https://example.com/bbb.mp4#disable" />
    </Sequence>
  );
};
```

## When to disable media fragment hints [​](\#when-to-disable-media-fragment-hints "Direct link to When to disable media fragment hints")

If your `<Sequence>`'s change value over time, the media fragments will change too.

Anytime they change, the browser considers this a new source and re-loads the video.

There are valid scenarios where we recommend disabling the media fragment hints:

- Changing the `from` and `durationInFrames` values of a `<Sequence>` dynamically
- Changing the `startFrom` and `endAt` values of a `<Video>` or `<OffthreadVideo>` dynamically
- Changing the `playbackRate` of a `<Video>` over time.

Normally you would not do this, but there are exceptions, like: [Changing the speed of a video over time](/docs/miscellaneous/snippets/accelerated-video)

## Recommendation [​](\#recommendation "Direct link to Recommendation")

Use the automatic media fragment hints as a default because they are helpful with playback performance.

If you find a valid scenario where you need to disable the media fragment hints, you can do so.

If the playback performance suffers, you can also provide your own media fragment hint to the URL.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-fragments.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Emojis](/docs/miscellaneous/emojis) [Next\
\
AI System Prompt](/docs/system-prompt)

- [Behavior](#behavior)
- [Upsides](#upsides)
- [How to disable media fragment hints](#how-to-disable-media-fragment-hints)
- [When to disable media fragment hints](#when-to-disable-media-fragment-hints)
- [Recommendation](#recommendation)
