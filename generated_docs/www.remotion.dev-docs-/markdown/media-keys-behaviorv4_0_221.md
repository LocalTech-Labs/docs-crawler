---
title: Media Keys Behaviorv4.0.221
source: Remotion Documentation
last_updated: 2024-12-22
---

# Media Keys Behaviorv4.0.221

- [Home page](/)
- [Player](/docs/player/)
- Media Keys

On this page

# Media Keys Behavior [v4.0.221](https://github.com/remotion-dev/remotion/releases/v4.0.221)

This document is about the behavior when a user:

- Presses the Play/Pause (⏯️) or Previous track (⏪) button on their keyboard
- Uses other controls such as Chromes built-in controls next to the user avatar to control the playback

These behaviors are controlled by the [`browserMediaControlsBehavior`](/docs/player/player#browsermediacontrolsbehavior) prop.

The underlying Web API used is the [Media Session API](https://developer.mozilla.org/en-US/docs/Web/API/Media_Session_API).

## Modes [​](\#modes "Direct link to Modes")

### `prevent-media-session` (default) [​](\#prevent-media-session-default "Direct link to prevent-media-session-default")

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'prevent-media-session',
      }}
      {...otherProps}
    />
  );
};
```

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'prevent-media-session',
      }}
      {...otherProps}
    />
  );
};
```

In this mode, Remotion will not act on the user's media keys, but map any keybord action to a no-op.

This prevents that the Player is paused but the audio tags get resumed by pressing the Play/Pause button on the keyboard, which was a problem prior to Remotion v4.0.221.

### `register-media-session` [​](\#register-media-session "Direct link to register-media-session")

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'register-media-session',
      }}
      {...otherProps}
    />
  );
};
```

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'register-media-session',
      }}
      {...otherProps}
    />
  );
};
```

In this mode, Remotion will use the Media Session API to register handlers:

- When the user presses the Play/Pause button on the keyboard
  - Toggle the Remotion Player's state
- When the user presses the Previous track button on the keyboard
  - Seek to the beginning of the video
- When the user presses the Fast Forward button on the keyboard
  - Seek 10 seconds forward
- When the user presses the Rewind button on the keyboard
  - Seek 10 seconds backward

Also, Remotion will react to seeking events and inform the device about the current playback position and duration.

### `do-nothing` [​](\#do-nothing "Direct link to do-nothing")

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'do-nothing',
      }}
      {...otherProps}
    />
  );
};
```

```

tsx

import {Player} from '@remotion/player';

export const MyComp: React.FC = () => {
  return (
    <Player
      browserMediaControlsBehavior={{
        mode: 'do-nothing',
      }}
      {...otherProps}
    />
  );
};
```

Reverts to the behavior prior to Remotion v4.0.221.

Remotion will not react to any media keys, leaving the browser to handle the media keys.

This leads to the problem that the user can resume any media tag by pressing the Play/Pause button on the keyboard, without the Remotion Player also resuming.

## When using multiple `<Player>`'s [​](\#when-using-multiple-players "Direct link to when-using-multiple-players")

Remotion's `register-media-session` handler is supposed to only work with 1 Player mounted.

It is not defined which Player reacts to the media keys.

When working with multiple Players, set one to `do-nothing` mode and the other to `register-media-session` mode to explicitly set the Media Keys for only 1 Player.

If you want all Players to react to the media keys, you need to use `do-nothing` mode and implement this behavior yourself with the Player API.

## In the Remotion Studio [​](\#in-the-remotion-studio "Direct link to In the Remotion Studio")

The behavior is set to `register-media-session` as of v4.0.221 and previously it behaved like `do-nothing`.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/media-keys.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Custom controls](/docs/player/custom-controls) [Next\
\
Overview](/docs/lambda)

- [Modes](#modes)
  - [`prevent-media-session` (default)](#prevent-media-session-default)
  - [`register-media-session`](#register-media-session)
  - [`do-nothing`](#do-nothing)
- [When using multiple `<Player>`'s](#when-using-multiple-players)
- [In the Remotion Studio](#in-the-remotion-studio)
