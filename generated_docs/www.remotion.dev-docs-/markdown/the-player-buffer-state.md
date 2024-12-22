---
title: The Player buffer state
source: Remotion Documentation
last_updated: 2024-12-22
---

# The Player buffer state

- [Home page](/)
- [Player](/docs/player/)
- Buffer state

On this page

# The Player buffer state

_available from 4.0.111_

Just like regular video players, it is possible that the content being displayed in the Player is not yet fully loaded.

In this case, the best practice is to briefly pause the video, let the content load and then resume playback.

Remotion has a native buffer state, which can be used to pause the video when the buffer is empty.

## TL;DR [​](\#tldr "Direct link to TL;DR")

You can add a `pauseWhenBuffering` prop to your [`<Video>`](/docs/video/#pausewhenbuffering), [`<OffthreadVideo>`](/docs/offthreadvideo/#pausewhenbuffering), [`<Audio>`](/docs/audio/#pausewhenbuffering) tags.

The prop is called `pauseWhenLoading` for [`<Img>`](/docs/img/#pausewhenloading) tags.

By doing so, the Player will briefly pause until your media is loaded.

## Mechanism [​](\#mechanism "Direct link to Mechanism")

### Activating the buffer state [​](\#activating-the-buffer-state "Direct link to Activating the buffer state")

A component can tell the player to switch into a buffer state by first using the [`useBufferState()`](/docs/use-buffer-state) hook and then calling `buffer.delayPlayback()`:

```

MyComp.tsx
tsx

import React from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return null;
};
```

```

MyComp.tsx
tsx

import React from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return null;
};
```

To clear the handle, call `.unblock()` on the return value of `delayPlayback()`.

When activating the buffer state, pay attention to the following:

Clear the handle when the component unmounts

The user might seek to a different portion of the video which is immediately available.
Use the cleanup function of `useEffect()` to clear the handle when your component is unmounted.

```

❌ Causes problems with React strict mode
tsx

import React, {useState} from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);
  }, []);

  return <></>;
};
```

```

❌ Causes problems with React strict mode
tsx

import React, {useState} from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);
  }, []);

  return <></>;
};
```

Don't use `delayPlayback()` inside a `useState()`

While the following implementation works in production, it fails in React Strict mode, because the `useState()` hook is called twice, which causes the first invocation of the buffer to never be cleared.

```

❌ Doesn't clear the buffer handle when seeking to a different portion of a video
tsx

import React, {useState} from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

```

❌ Doesn't clear the buffer handle when seeking to a different portion of a video
tsx

import React, {useState} from 'react';
import {useBufferState} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [delayHandle] = useState(() => buffer.delayPlayback()); // 💥

  React.useEffect(() => {
    setTimeout(() => {
      delayHandle.unblock();
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

Details

It doesn't replace `delayRender()` [delayRender()](/docs/delay-render) is a different API which controls when a screenshot is taken during rendering.

If you are loading data, you might want to delay the screenshotting of your component during rendering and delay the playback of the video in preview, in which case you need to use both APIs together.

```

Using delayRender() and delayPlayback() together
tsx

import React from 'react';
import {useBufferState, delayRender, continueRender} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [handle] = React.useState(() => delayRender());

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
      continueRender(handle);
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

```

Using delayRender() and delayPlayback() together
tsx

import React from 'react';
import {useBufferState, delayRender, continueRender} from 'remotion';

const MyComp: React.FC = () => {
  const buffer = useBufferState();
  const [handle] = React.useState(() => delayRender());

  React.useEffect(() => {
    const delayHandle = buffer.delayPlayback();

    setTimeout(() => {
      delayHandle.unblock();
      continueRender(handle);
    }, 5000);

    return () => {
      delayHandle.unblock();
    };
  }, []);

  return <></>;
};
```

### Possible states [​](\#possible-states "Direct link to Possible states")

Whether a player is buffering does not internally change the `playing` / `paused` state.

Therefore, a player can be in four playback states:

[1](#1) `playing && !buffering`

[2](#2) `playing && buffering`

[3](#3) `paused && !buffering`

[4](#4) `paused && buffering`

Only in state

1

the time moves forward.

* * *

By default, Remotion will display the following UI based on the state of the player:

When in State

1

a Pause button is shown.

When in State

2

at first a Pause button, then after [a delay](#indicating-buffering-in-the-ui), a [customizable spinner](/docs/player/buffer-state#indicating-buffering-in-the-ui)
is shown.

Otherwise, the Play button is shown.

You may add additional UI to this, for example by overlaying the Player with a spinner when the Player is buffering.

### Listening to buffer events [​](\#listening-to-buffer-events "Direct link to Listening to buffer events")

If the `<Player />` is entering a buffer state, it will emit the `waiting` event.

Once it resumes, it emits the `resume` event.

```

Listening to waiting and resume events
tsx

import {Player, PlayerRef} from '@remotion/player';
import {useEffect, useRef, useState} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);
  const [buffering, setBuffering] = useState(false);

  useEffect(() => {
    const {current} = playerRef;
    if (!current) {
      return;
    }

    const onBuffering = () => {
      setBuffering(true);
    };
    const onResume = () => {
      setBuffering(false);
    };

    current.addEventListener('waiting', onBuffering);
    current.addEventListener('resume', onResume);
    return () => {
      current.removeEventListener('waiting', onBuffering);
      current.removeEventListener('resume', onResume);
    };
  }, [setBuffering]);

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

Listening to waiting and resume events
tsx

import {Player, PlayerRef} from '@remotion/player';
import {useEffect, useRef, useState} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  const playerRef = useRef<PlayerRef>(null);
  const [buffering, setBuffering] = useState(false);

  useEffect(() => {
    const {current} = playerRef;
    if (!current) {
      return;
    }

    const onBuffering = () => {
      setBuffering(true);
    };
    const onResume = () => {
      setBuffering(false);
    };

    current.addEventListener('waiting', onBuffering);
    current.addEventListener('resume', onResume);
    return () => {
      current.removeEventListener('waiting', onBuffering);
      current.removeEventListener('resume', onResume);
    };
  }, [setBuffering]);

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

## Components with built-in buffering [​](\#components-with-built-in-buffering "Direct link to Components with built-in buffering")

You can enable buffering on the following components:

- [`<Audio>`](/docs/audio#pausewhenbuffering) \- add the `pauseWhenBuffering` prop
- [`<Video>`](/docs/video#pausewhenbuffering) \- add the `pauseWhenBuffering` prop
- [`<OffthreadVideo>`](/docs/offthreadvideo#pausewhenbuffering) \- add the `pauseWhenBuffering` prop
- [`<Img>`](/docs/img#pausewhenloading) \- add the `pauseWhenLoading` prop

## Indicating buffering in the UI [​](\#indicating-buffering-in-the-ui "Direct link to Indicating buffering in the UI")

When the Player is buffering, by default the Play button will be replaced with a spinner.

To prevent a janky UI, this spinner will only be shown after the Player has been in a buffering state for 300ms.

You may customize the timeout of `300` milliseconds by passing the [`bufferStateDelayInMilliseconds`](/docs/player/player#bufferstatedelayinmilliseconds) prop to the `<Player />` component.

```

Setting the delay until the spinner is shown
tsx

import {Player, PlayerRef} from '@remotion/player';
import {useEffect, useRef, useState} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      bufferStateDelayInMilliseconds={1000} // Or set to `0` to immediately show the spinner
    />
  );
};
```

```

Setting the delay until the spinner is shown
tsx

import {Player, PlayerRef} from '@remotion/player';
import {useEffect, useRef, useState} from 'react';
import {MyVideo} from './remotion/MyVideo';

export const App: React.FC = () => {
  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      bufferStateDelayInMilliseconds={1000} // Or set to `0` to immediately show the spinner
    />
  );
};
```

In the Studio, you can change the delay in the [config file](/docs/config#setbufferstatedelayinmilliseconds):

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.setBufferStateDelayInMilliseconds(0);
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.setBufferStateDelayInMilliseconds(0);
```

To customize the spinner that is shown in place of the Play button, you can pass a [`renderPlayPauseButton()`](/docs/player/player#renderplaypausebutton) prop:

```

Rendering a custom spinner inside the Play button
tsx

import {Player, RenderPlayPauseButton} from '@remotion/player';
import {useCallback} from 'react';

export const App: React.FC = () => {
  const renderPlayPauseButton: RenderPlayPauseButton = useCallback(
    ({playing, isBuffering}) => {
      if (playing && isBuffering) {
        return <MySpinner />;
      }

      return null;
    },
    [],
  );

  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      renderPlayPauseButton={renderPlayPauseButton}
    />
  );
};
```

```

Rendering a custom spinner inside the Play button
tsx

import {Player, RenderPlayPauseButton} from '@remotion/player';
import {useCallback} from 'react';

export const App: React.FC = () => {
  const renderPlayPauseButton: RenderPlayPauseButton = useCallback(
    ({playing, isBuffering}) => {
      if (playing && isBuffering) {
        return <MySpinner />;
      }

      return null;
    },
    [],
  );

  return (
    <Player
      component={MyVideo}
      durationInFrames={120}
      compositionWidth={1920}
      compositionHeight={1080}
      fps={30}
      renderPlayPauseButton={renderPlayPauseButton}
    />
  );
};
```

To display a loading UI layered on top of the Player (e.g. a spinner), you can set [`showPosterWhenBuffering`](/docs/player/player#showposterwhenbuffering) to `true` and pass a [`renderPoster()`](/docs/player/player#renderposter) prop:

```

Rendering a custom spinner on top of the Player
tsx

import type {RenderPoster} from '@remotion/player';
import {Player} from '@remotion/player';

const MyApp: React.FC = () => {
  const renderPoster: RenderPoster = useCallback(({isBuffering}) => {
    if (isBuffering) {
      return (
        <AbsoluteFill style={{justifyContent: 'center', alignItems: 'center'}}>
          <Spinner />
        </AbsoluteFill>
      );
    }

    return null;
  }, []);

  return (
    <Player
      fps={30}
      component={Component}
      durationInFrames={100}
      compositionWidth={1080}
      compositionHeight={1080}
      renderPoster={renderPoster}
      showPosterWhenBuffering
    />
  );
};
```

```

Rendering a custom spinner on top of the Player
tsx

import type {RenderPoster} from '@remotion/player';
import {Player} from '@remotion/player';

const MyApp: React.FC = () => {
  const renderPoster: RenderPoster = useCallback(({isBuffering}) => {
    if (isBuffering) {
      return (
        <AbsoluteFill style={{justifyContent: 'center', alignItems: 'center'}}>
          <Spinner />
        </AbsoluteFill>
      );
    }

    return null;
  }, []);

  return (
    <Player
      fps={30}
      component={Component}
      durationInFrames={100}
      compositionWidth={1080}
      compositionHeight={1080}
      renderPoster={renderPoster}
      showPosterWhenBuffering
    />
  );
};
```

## Upcoming changes in Remotion 5.0 [​](\#upcoming-changes-in-remotion-50 "Direct link to Upcoming changes in Remotion 5.0")

In Remotion 4.0, media tags such as `<Audio>`, `<OffthreadVideo>` tags will need to opt-in to use the buffer state.

In Remotion 5.0, it is planned that `<Audio>`, `<Video>` and `<OffthreadVideo>` will automatically use the buffer state, but they can opt out of it.

## See also [​](\#see-also "Direct link to See also")

- [`useBufferState()`](/docs/use-buffer-state)
- [`<Player>`](/docs/player/player)
- [Avoiding flickering in `<Player>`](/docs/troubleshooting/player-flicker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/buffer-state.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Avoiding flickers](/docs/troubleshooting/player-flicker) [Next\
\
Preloading](/docs/player/preloading)

- [TL;DR](#tldr)
- [Mechanism](#mechanism)
  - [Activating the buffer state](#activating-the-buffer-state)
  - [Possible states](#possible-states)
  - [Listening to buffer events](#listening-to-buffer-events)
- [Components with built-in buffering](#components-with-built-in-buffering)
- [Indicating buffering in the UI](#indicating-buffering-in-the-ui)
- [Upcoming changes in Remotion 5.0](#upcoming-changes-in-remotion-50)
- [See also](#see-also)
