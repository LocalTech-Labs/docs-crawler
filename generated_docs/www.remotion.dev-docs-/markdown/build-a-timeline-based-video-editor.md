---
title: Build a timeline-based video editor
source: Remotion Documentation
last_updated: 2024-12-22
---

# Build a timeline-based video editor

- [Home page](/)
- Building apps
- Build a timeline-based video editor

On this page

# Build a timeline-based video editor

This document describes on a high-level how the [Remotion Player](/player) can be synchronized with a timeline.

Read this document for guidance on building a video editor with the following characteristics:

- Multiple tracks that overlay each other
- Items can be arbitrarily placed on a track
- Items can be of different types (e.g. video, audio, text, etc.)

![](/img/timelineitems.png)

[1](#1)

Define a TypeScript type `Item` defining the different
item types. Create another one for defining the shape of a `Track`:

```

types.ts
tsx

type BaseItem = {
  from: number;
  durationInFrames: number;
  id: string;
};

export type SolidItem = BaseItem & {
  type: 'solid';
  color: string;
};

export type TextItem = BaseItem & {
  type: 'text';
  text: string;
  color: string;
};

export type VideoItem = BaseItem & {
  type: 'video';
  src: string;
};

export type Item = SolidItem | TextItem | VideoItem;

export type Track = {
  name: string;
  items: Item[];
};
```

```

types.ts
tsx

type BaseItem = {
  from: number;
  durationInFrames: number;
  id: string;
};

export type SolidItem = BaseItem & {
  type: 'solid';
  color: string;
};

export type TextItem = BaseItem & {
  type: 'text';
  text: string;
  color: string;
};

export type VideoItem = BaseItem & {
  type: 'video';
  src: string;
};

export type Item = SolidItem | TextItem | VideoItem;

export type Track = {
  name: string;
  items: Item[];
};
```

[2](#2)

Create a component that can render a list of tracks.

```

remotion/Main.tsx
tsx

import type {Track, Item} from './types';
import React from 'react';
import {AbsoluteFill, Sequence, OffthreadVideo} from 'remotion';

const ItemComp: React.FC<{
  item: Item;
}> = ({item}) => {
  if (item.type === 'solid') {
    return <AbsoluteFill style={{backgroundColor: item.color}} />;
  }

  if (item.type === 'text') {
    return <h1>{item.text}</h1>;
  }

  if (item.type === 'video') {
    return <OffthreadVideo src={item.src} />;
  }

  throw new Error(`Unknown item type: ${JSON.stringify(item)}`);
};

const Track: React.FC<{
  track: Track;
}> = ({track}) => {
  return (
    <AbsoluteFill>
      {track.items.map((item) => {
        return (
          <Sequence
            key={item.id}
            from={item.from}
            durationInFrames={item.durationInFrames}
          >
            <ItemComp item={item} />
          </Sequence>
        );
      })}
    </AbsoluteFill>
  );
};

export const Main: React.FC<{
  tracks: Track[];
}> = ({tracks}) => {
  return (
    <AbsoluteFill>
      {tracks.map((track) => {
        return <Track track={track} key={track.name} />;
      })}
    </AbsoluteFill>
  );
};
```

```

remotion/Main.tsx
tsx

import type {Track, Item} from './types';
import React from 'react';
import {AbsoluteFill, Sequence, OffthreadVideo} from 'remotion';

const ItemComp: React.FC<{
  item: Item;
}> = ({item}) => {
  if (item.type === 'solid') {
    return <AbsoluteFill style={{backgroundColor: item.color}} />;
  }

  if (item.type === 'text') {
    return <h1>{item.text}</h1>;
  }

  if (item.type === 'video') {
    return <OffthreadVideo src={item.src} />;
  }

  throw new Error(`Unknown item type: ${JSON.stringify(item)}`);
};

const Track: React.FC<{
  track: Track;
}> = ({track}) => {
  return (
    <AbsoluteFill>
      {track.items.map((item) => {
        return (
          <Sequence
            key={item.id}
            from={item.from}
            durationInFrames={item.durationInFrames}
          >
            <ItemComp item={item} />
          </Sequence>
        );
      })}
    </AbsoluteFill>
  );
};

export const Main: React.FC<{
  tracks: Track[];
}> = ({tracks}) => {
  return (
    <AbsoluteFill>
      {tracks.map((track) => {
        return <Track track={track} key={track.name} />;
      })}
    </AbsoluteFill>
  );
};
```

tip

In CSS, the elements that are rendered at the bottom appear at the top. See: [Layers](/docs/layers)

[3](#3)

Keep a state of tracks each containing an array of items.

Render
a [`<Player />`](/docs/player/player) component and pass the `tracks` as [`inputProps`](/docs/player/player#inputprops).

```

Editor.tsx
tsx

import React, {useMemo, useState} from 'react';
import {Player} from '@remotion/player';
import type {Item} from './types';
import {Main} from './remotion/Main';

type Track = {
  name: string;
  items: Item[];
};

export const Editor = () => {
  const [tracks, setTracks] = useState<Track[]>([
    {name: 'Track 1', items: []},
    {name: 'Track 2', items: []},
  ]);

  const inputProps = useMemo(() => {
    return {
      tracks,
    };
  }, [tracks]);

  return (
    <>
      <Player
        component={Main}
        fps={30}
        inputProps={inputProps}
        durationInFrames={600}
        compositionWidth={1280}
        compositionHeight={720}
      />
    </>
  );
};
```

```

Editor.tsx
tsx

import React, {useMemo, useState} from 'react';
import {Player} from '@remotion/player';
import type {Item} from './types';
import {Main} from './remotion/Main';

type Track = {
  name: string;
  items: Item[];
};

export const Editor = () => {
  const [tracks, setTracks] = useState<Track[]>([
    {name: 'Track 1', items: []},
    {name: 'Track 2', items: []},
  ]);

  const inputProps = useMemo(() => {
    return {
      tracks,
    };
  }, [tracks]);

  return (
    <>
      <Player
        component={Main}
        fps={30}
        inputProps={inputProps}
        durationInFrames={600}
        compositionWidth={1280}
        compositionHeight={720}
      />
    </>
  );
};
```

[4](#4)

Build a timeline component: You now have access to the `
tracks
` state and can update it using the `setTracks` function.

We do not currently provide samples how to build a timeline component, since everybody has different needs and styling preferences.

An opinionated sample implementation is [available for purchase in the Remotion Store](https://www.remotion.pro/timeline).

```

remotion/Timeline.tsx
tsx

const Editor: React.FC = () => {
  const [tracks, setTracks] = useState<Track[]>([
    {name: 'Track 1', items: []},
    {name: 'Track 2', items: []},
  ]);

  const inputProps = useMemo(() => {
    return {
      tracks,
    };
  }, [tracks]);

  return (
    <>
      <Player
        component={Main}
        fps={30}
        inputProps={inputProps}
        durationInFrames={600}
        compositionWidth={1280}
        compositionHeight={720}
      />
      <Timeline tracks={tracks} setTracks={setTracks} />
    </>
  );
};
```

```

remotion/Timeline.tsx
tsx

const Editor: React.FC = () => {
  const [tracks, setTracks] = useState<Track[]>([
    {name: 'Track 1', items: []},
    {name: 'Track 2', items: []},
  ]);

  const inputProps = useMemo(() => {
    return {
      tracks,
    };
  }, [tracks]);

  return (
    <>
      <Player
        component={Main}
        fps={30}
        inputProps={inputProps}
        durationInFrames={600}
        compositionWidth={1280}
        compositionHeight={720}
      />
      <Timeline tracks={tracks} setTracks={setTracks} />
    </>
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Layers](/docs/layers)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/building-a-timeline.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Build a Google Font picker](/docs/font-picker) [Next\
\
Supporting multiple frame rates](/docs/multiple-fps)

- [See also](#see-also)
