---
title: Displaying the current time
source: Remotion Documentation
last_updated: 2024-12-22
---

# Displaying the current time

- [Home page](/)
- [Player](/docs/player/)
- Getting the current time

On this page

# Displaying the current time

When rendering a [`<Player>`](/docs/player) in your app, special considerations must be taken to prevent constant re-renders of the app or [`<Player>`](/docs/player) if the time changes.

This is why the [`useCurrentFrame()`](/docs/use-current-frame) hook does not work outside a composition.

warning

Do not put this hook into the same component in which the `<Player>` is rendered, otherwise you'll see constant re-renders. Instead, put it inside a component that is rendered adjacent to the component in which the Player is rendered.

## Synchronizing a component with the Player time [​](\#synchronizing-a-component-with-the-player-time "Direct link to Synchronizing a component with the Player time")

If you want to display a component that synchronizes with the time of the player, for example the cursor of a timeline component or a custom time display, you can use the following hook:

```

use-current-player-frame.ts
tsx

import {CallbackListener, PlayerRef} from '@remotion/player';
import {useCallback, useSyncExternalStore} from 'react';

export const useCurrentPlayerFrame = (
  ref: React.RefObject<PlayerRef | null>,
) => {
  const subscribe = useCallback(
    (onStoreChange: () => void) => {
      const {current} = ref;
      if (!current) {
        return () => undefined;
      }
      const updater: CallbackListener<'frameupdate'> = ({detail}) => {
        onStoreChange();
      };
      current.addEventListener('frameupdate', updater);
      return () => {
        current.removeEventListener('frameupdate', updater);
      };
    },
    [ref],
  );

  const data = useSyncExternalStore<number>(
    subscribe,
    () => ref.current?.getCurrentFrame() ?? 0,
    () => 0,
  );

  return data;
};
```

```

use-current-player-frame.ts
tsx

import {CallbackListener, PlayerRef} from '@remotion/player';
import {useCallback, useSyncExternalStore} from 'react';

export const useCurrentPlayerFrame = (
  ref: React.RefObject<PlayerRef | null>,
) => {
  const subscribe = useCallback(
    (onStoreChange: () => void) => {
      const {current} = ref;
      if (!current) {
        return () => undefined;
      }
      const updater: CallbackListener<'frameupdate'> = ({detail}) => {
        onStoreChange();
      };
      current.addEventListener('frameupdate', updater);
      return () => {
        current.removeEventListener('frameupdate', updater);
      };
    },
    [ref],
  );

  const data = useSyncExternalStore<number>(
    subscribe,
    () => ref.current?.getCurrentFrame() ?? 0,
    () => 0,
  );

  return data;
};
```

## Usage example [​](\#usage-example "Direct link to Usage example")

Add a ref to a React Player and pass it to another component:

```

tsx

import {Player, PlayerRef} from '@remotion/player';
import {useRef} from 'react';
import {MyVideo} from './remotion/MyVideo';
import {TimeDisplay} from './remotion/TimeDisplay';

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  return (
    <>
      <Player
        ref={playerRef}
        component={MyVideo}
        durationInFrames={120}
        compositionWidth={1920}
        compositionHeight={1080}
        fps={30}
      />
      <TimeDisplay playerRef={playerRef} />
    </>
  );
};
```

```

tsx

import {Player, PlayerRef} from '@remotion/player';
import {useRef} from 'react';
import {MyVideo} from './remotion/MyVideo';
import {TimeDisplay} from './remotion/TimeDisplay';

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  return (
    <>
      <Player
        ref={playerRef}
        component={MyVideo}
        durationInFrames={120}
        compositionWidth={1920}
        compositionHeight={1080}
        fps={30}
      />
      <TimeDisplay playerRef={playerRef} />
    </>
  );
};
```

This is how a component could access the current time:

```

TimeDisplay.tsx
tsx

import React from 'react';
import {PlayerRef} from '@remotion/player';
import {useCurrentPlayerFrame} from './use-current-player-frame';

export const TimeDisplay: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  const frame = useCurrentPlayerFrame(playerRef);

  return <div>current frame: {frame}</div>;
};
```

```

TimeDisplay.tsx
tsx

import React from 'react';
import {PlayerRef} from '@remotion/player';
import {useCurrentPlayerFrame} from './use-current-player-frame';

export const TimeDisplay: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  const frame = useCurrentPlayerFrame(playerRef);

  return <div>current frame: {frame}</div>;
};
```

This approach is efficient, because only the video itself and the component relying on the time are re-rendering, but the `<App>` component is not.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/current-time.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Autoplay](/docs/player/autoplay) [Next\
\
Avoiding flickers](/docs/troubleshooting/player-flicker)

- [Synchronizing a component with the Player time](#synchronizing-a-component-with-the-player-time)
- [Usage example](#usage-example)
