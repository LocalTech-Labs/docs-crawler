---
title: Examples for @remotion/player
source: Remotion Documentation
last_updated: 2024-12-22
---

# Examples for @remotion/player

- [Home page](/)
- [Player](/docs/player/)
- Examples

On this page

# Examples for @remotion/player

## Bare example [​](\#bare-example "Direct link to Bare example")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

## Adding controls [​](\#adding-controls "Direct link to Adding controls")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
    />
  );
};
```

## Loop video [​](\#loop-video "Direct link to Loop video")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
    />
  );
};
```

## Changing size [​](\#changing-size "Direct link to Changing size")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
      style={{
        width: 1280,
        height: 720,
      }}
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
      style={{
        width: 1280,
        height: 720,
      }}
    />
  );
};
```

note

See [Scaling](/docs/player/scaling) for more ways to scale the Player.

## Adding autoplay [​](\#adding-autoplay "Direct link to Adding autoplay")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
      style={{
        width: 1280,
        height: 720,
      }}
      autoPlay
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      controls
      loop
      style={{
        width: 1280,
        height: 720,
      }}
      autoPlay
    />
  );
};
```

tip

You need to be wary of the browser's autoplay policy. In most browsers, you cannot autoplay an audio file or a video with audio.

## Programmatically control the player [​](\#programmatically-control-the-player "Direct link to Programmatically control the player")

```

tsx

import { Player, PlayerRef } from "@remotion/player";
import { useCallback, useRef } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  const seekToMiddle = useCallback(() => {
    const { current } = playerRef;
    if (!current) {
      return;
    }
    current.seekTo(60);
  }, []);

  return (
    <Player
      ref={playerRef}
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

```

tsx

import { Player, PlayerRef } from "@remotion/player";
import { useCallback, useRef } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  const seekToMiddle = useCallback(() => {
    const { current } = playerRef;
    if (!current) {
      return;
    }
    current.seekTo(60);
  }, []);

  return (
    <Player
      ref={playerRef}
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

## Listen to events [​](\#listen-to-events "Direct link to Listen to events")

```

tsx

import { Player, PlayerRef } from "@remotion/player";
import { useEffect, useRef } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  useEffect(() => {
    const { current } = playerRef;
    if (!current) {
      return;
    }

    const listener = () => {
      console.log("paused");
    };
    current.addEventListener("pause", listener);
    return () => {
      current.removeEventListener("pause", listener);
    };
  }, []);

  return (
    <Player
      ref={playerRef}
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

```

tsx

import { Player, PlayerRef } from "@remotion/player";
import { useEffect, useRef } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);

  useEffect(() => {
    const { current } = playerRef;
    if (!current) {
      return;
    }

    const listener = () => {
      console.log("paused");
    };
    current.addEventListener("pause", listener);
    return () => {
      current.removeEventListener("pause", listener);
    };
  }, []);

  return (
    <Player
      ref={playerRef}
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
    />
  );
};
```

## Interactive player [​](\#interactive-player "Direct link to Interactive player")

```

tsx

import { Player } from "@remotion/player";
import { useState, useMemo } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  // Connect the state to a text field
  const [text, setText] = useState("world");
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

tsx

import { Player } from "@remotion/player";
import { useState, useMemo } from "react";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  // Connect the state to a text field
  const [text, setText] = useState("world");
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

## Only play a portion of a video [​](\#only-play-a-portion-of-a-video "Direct link to Only play a portion of a video")

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      loop
      inFrame={30}
      outFrame={60}
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { MyVideo } from "./remotion/MyVideo";

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      loop
      inFrame={30}
      outFrame={60}
    />
  );
};
```

## Loading a component lazily [​](\#loading-a-component-lazily "Direct link to Loading a component lazily")

```

tsx

import { Player } from "@remotion/player";
import { useCallback } from "react";

export const App: React.FC = () => {
  const lazyComponent = useCallback(() => {
    return import("./remotion/MyVideo");
  }, []);

  return (
    <Player
      lazyComponent={lazyComponent}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      loop
    />
  );
};
```

```

tsx

import { Player } from "@remotion/player";
import { useCallback } from "react";

export const App: React.FC = () => {
  const lazyComponent = useCallback(() => {
    return import("./remotion/MyVideo");
  }, []);

  return (
    <Player
      lazyComponent={lazyComponent}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      loop
    />
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [Snippet: Embedding a `<Player>` in an iframe](/docs/miscellaneous/snippets/player-in-iframe)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/player-examples.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
@remotion/player](/docs/player/) [Next\
\
Sizing](/docs/player/scaling)

- [Bare example](#bare-example)
- [Adding controls](#adding-controls)
- [Loop video](#loop-video)
- [Changing size](#changing-size)
- [Adding autoplay](#adding-autoplay)
- [Programmatically control the player](#programmatically-control-the-player)
- [Listen to events](#listen-to-events)
- [Interactive player](#interactive-player)
- [Only play a portion of a video](#only-play-a-portion-of-a-video)
- [Loading a component lazily](#loading-a-component-lazily)
- [See also](#see-also)
