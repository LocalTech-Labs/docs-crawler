---
title: Player - Best practices
source: Remotion Documentation
last_updated: 2024-12-22
---

# Player - Best practices

- [Home page](/)
- [Player](/docs/player/)
- Best practices

On this page

# Player - Best practices

## Avoid re-renders of the `<Player>` [​](\#avoid-re-renders-of-the-player "Direct link to avoid-re-renders-of-the-player")

The following pattern is not ideal because every time the time updates, the `<Player>` is being re-rendered:

```

❌ Problematic
tsx

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);
  const [currentTime, setCurrentTime] = useState(0);

  useEffect(() => {
    playerRef.current?.addEventListener('timeupdate', (e) => {
      setCurrentTime(e.detail.frame);
    });
  }, []);

  return (
    <div>
      <Player ref={playerRef} component={MyVideo} {...otherProps} />
      <div>Current time: {currentTime}</div>
    </div>
  );
};
```

```

❌ Problematic
tsx

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);
  const [currentTime, setCurrentTime] = useState(0);

  useEffect(() => {
    playerRef.current?.addEventListener('timeupdate', (e) => {
      setCurrentTime(e.detail.frame);
    });
  }, []);

  return (
    <div>
      <Player ref={playerRef} component={MyVideo} {...otherProps} />
      <div>Current time: {currentTime}</div>
    </div>
  );
};
```

We advise to render your controls and UI as a sibling to the component that renders the `<Player>` and pass a ref to the player as a prop:

```

✅ Better
tsx

const PlayerOnly: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  return <Player ref={playerRef} component={MyVideo} {...otherProps} />;
};

const ControlsOnly: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  const [currentTime, setCurrentTime] = useState(0);

  useEffect(() => {
    playerRef.current?.addEventListener('timeupdate', (e) => {
      setCurrentTime(e.detail.frame);
    });
  }, []);

  return <div>Current time: {currentTime}</div>;
};

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  return (
    <>
      <PlayerOnly playerRef={playerRef} />
      <ControlsOnly playerRef={playerRef} />
    </>
  );
};
```

```

✅ Better
tsx

const PlayerOnly: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  return <Player ref={playerRef} component={MyVideo} {...otherProps} />;
};

const ControlsOnly: React.FC<{
  playerRef: React.RefObject<PlayerRef | null>;
}> = ({playerRef}) => {
  const [currentTime, setCurrentTime] = useState(0);

  useEffect(() => {
    playerRef.current?.addEventListener('timeupdate', (e) => {
      setCurrentTime(e.detail.frame);
    });
  }, []);

  return <div>Current time: {currentTime}</div>;
};

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  return (
    <>
      <PlayerOnly playerRef={playerRef} />
      <ControlsOnly playerRef={playerRef} />
    </>
  );
};
```

This is much more performant because the `<Player>` gets rerendered less often.

This advice is mainly for frequently updated state like the current time. Keeping state like a `loop` toggle in the parent component is fine, because it is not expected to change frequently.

[A more complex example can be found here.](https://github.com/remotion-dev/remotion/blob/main/packages/player-example/src/App.tsx)

## Pass the user interaction event to `play()` [​](\#pass-the-user-interaction-event-to-play "Direct link to pass-the-user-interaction-event-to-play")

When you listen to an `onClick()` event, the browser will give you an `event` argument.

Pass it on to [`.play()`](/docs/player/player#play) and [`.toggle()`](/docs/player/player#toggle) to have the lowest chance of hitting autoplay restrictions that the browser may impose.

[Read here for more details.](/docs/player/autoplay#pass-the-event-to-the-play-or-toggle-method)

## Memoize the `inputProps` [​](\#memoize-the-inputprops "Direct link to memoize-the-inputprops")

Not memoizing the `inputProps` can cause the whole tree to re-render too many times and create a bottleneck.

```

Player.tsx
tsx

import {Player} from '@remotion/player';
import {useState, useMemo} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  const [text, setText] = useState('world');
  const inputProps = useMemo(() => {
    return {
      text,
    };
  }, [text]);

  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      inputProps={inputProps}
    />
  );
};
```

```

Player.tsx
tsx

import {Player} from '@remotion/player';
import {useState, useMemo} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  const [text, setText] = useState('world');
  const inputProps = useMemo(() => {
    return {
      text,
    };
  }, [text]);

  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      inputProps={inputProps}
    />
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Handling autoplay in the Player](/docs/player/autoplay)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/best-practices.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Premounting](/docs/player/premounting) [Next\
\
Drag & Drop](/docs/player/drag-and-drop/)

- [Avoid re-renders of the `<Player>`](#avoid-re-renders-of-the-player)
- [Pass the user interaction event to `play()`](#pass-the-user-interaction-event-to-play)
- [Memoize the `inputProps`](#memoize-the-inputprops)
- [See also](#see-also)
