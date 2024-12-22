---
title: Combatting autoplay issues
source: Remotion Documentation
last_updated: 2024-12-22
---

# Combatting autoplay issues

- [Home page](/)
- [Player](/docs/player/)
- Autoplay

On this page

# Combatting autoplay issues

Browsers place restrictions on websites that play audio without user interaction.

Read on how to properly use the Remotion Player so you don't run into a browser policy.

The most strict browser is **Mobile Safari**.

If your composition plays well there, you should have no problems elsewhere.

## Trigger the play from a user event [​](\#trigger-the-play-from-a-user-event "Direct link to Trigger the play from a user event")

If you autoplay audio (through an `<Audio>`, `<OffthreadVideo>` or `<Video>` tag) on site load, browsers will block the audio.

The [`autoPlay`](/docs/player/player#autoplay) prop of the [`<Player>`](/docs/player/player) is therefore discouraged if your video contains any audio.

Display a [play button](/docs/player/custom-controls) instead or use the [`controls`](/docs/player/player#controls) prop.

## Pass the event to the `play()` or `toggle()` method [​](\#pass-the-event-to-the-play-or-toggle-method "Direct link to pass-the-event-to-the-play-or-toggle-method")

If you don't use the default controls, make sure the call to [`.play()`](/docs/player/player#play) is initialized from a user gesture, for example a mouse click.

You should get a Javascript event like a `MouseEvent` from it.

Pass this event to the [`play()`](/docs/player/player#play) or [`toggle()`](/docs/player/player#toggle) function so audio may play automatically.

```

tsx

import {PlayerRef} from '@remotion/player';
import {useRef} from 'react';

export const App: React.FC = () => {
  const ref = useRef<PlayerRef>(null);

  return (
    <button
      onClickCapture={(e) => {
        const {current} = ref;
        // Pass the event to .play() or .toggle()
        current?.play(e);
      }}
    >
      Play
    </button>
  );
};
```

```

tsx

import {PlayerRef} from '@remotion/player';
import {useRef} from 'react';

export const App: React.FC = () => {
  const ref = useRef<PlayerRef>(null);

  return (
    <button
      onClickCapture={(e) => {
        const {current} = ref;
        // Pass the event to .play() or .toggle()
        current?.play(e);
      }}
    >
      Play
    </button>
  );
};
```

note

Prefer the `onClickCapture` event instead of the `onClick` event to make the event propagation work properly in Safari.

## Media that enters the video after the start [​](\#media-that-enters-the-video-after-the-start "Direct link to Media that enters the video after the start")

Media that contains audio and that plays without user interaction requires special attention.

Consider the following scenario:

```

tsx

export const MyComp = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile('audio.mp3')} />
      <Sequence from={120}>
        <Audio src={staticFile('audio2.mp3')} />
      </Sequence>
      <Sequence from={120}>
        <OffthreadVideo src={staticFile('video.mp4')} />
      </Sequence>
    </AbsoluteFill>
  );
};
```

```

tsx

export const MyComp = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile('audio.mp3')} />
      <Sequence from={120}>
        <Audio src={staticFile('audio2.mp3')} />
      </Sequence>
      <Sequence from={120}>
        <OffthreadVideo src={staticFile('video.mp4')} />
      </Sequence>
    </AbsoluteFill>
  );
};
```

There are three different cases here:

[1](#1)

An `<Audio>` tag that plays immediately. ✅ This is fine because we have user interaction.

[2](#2)

An `<Audio>` tag that plays after 120 frames. ⚠️ May be subject to browser autoplay policies, but there is a workaround
→ [Using the `numberOfSharedAudioTags` prop](/docs/player/autoplay#using-the-numberofsharedaudiotags-prop)

[3](#3)

A `<OffthreadVideo>` tag that plays after 120 frames. ⛔ May be subject to browser autoplay policies, which might restrict the video from playing with audio → [Dealing with an autoplay failure](/docs/player/autoplay#dealing-with-an-autoplay-failure)

## Using the `numberOfSharedAudioTags` prop [​](\#using-the-numberofsharedaudiotags-prop "Direct link to using-the-numberofsharedaudiotags-prop")

By default, Remotion will mount 5 auxiliary `<audio>` tags that will "warm" up with the first user interaction.

If an `<Audio>` tag is played after the first user interaction, it will be mounted into one of the auxiliary tags and will be allowed to play.

If you have more than 5 `<Audio>` tags that play at the same time, you can increase the number of auxiliary tags by passing the [`numberOfSharedAudioTags`](/docs/player/player#numberofsharedaudiotags) prop to the `<Player>`, because otherwise an error will be thrown.

## Dealing with an autoplay failure [v4.0.187](https://github.com/remotion-dev/remotion/releases/v4.0.187) [​](\#dealing-with-an-autoplay-failure "Direct link to dealing-with-an-autoplay-failure")

If the browser blocks a video from being autoplayed due to it having sound and not being started from a user interaction:

**By default**: The video will be muted, and played without audio. A console message will be printed.

**Custom behavior**: You may handle an autoplay failure using the [`onAutoPlayError`](/docs/offthreadvideo#onautoplayerror) prop of `<OffthreadVideo>` and `<Video>`:

```

tsx

import {OffthreadVideo, staticFile} from 'remotion';
import type {PlayerRef} from '@remotion/player';

export const MyComp: React.FC<{
  playerRef: React.RefObject<PlayerRef>;
}> = ({playerRef}) => {
  return (
    <OffthreadVideo
      src={staticFile('video.mp4')}
      onAutoPlayError={() => {
        playerRef.current?.pause();
      }}
    />
  );
};
```

```

tsx

import {OffthreadVideo, staticFile} from 'remotion';
import type {PlayerRef} from '@remotion/player';

export const MyComp: React.FC<{
  playerRef: React.RefObject<PlayerRef>;
}> = ({playerRef}) => {
  return (
    <OffthreadVideo
      src={staticFile('video.mp4')}
      onAutoPlayError={() => {
        playerRef.current?.pause();
      }}
    />
  );
};
```

In this example, the Player is simply paused.

If the user presses Play again, the video may then start playing with audio again due to it being started with user interaction.

You may show a [poster](/docs/player/player/#renderposter) with a hint - for example "Tap to resume".

## When will a browser restrict autoplay? [​](\#when-will-a-browser-restrict-autoplay "Direct link to When will a browser restrict autoplay?")

It is not entirely deterministic.

After playing a video with user gesture, Safari will allow more media to be played for some amount of time.

However, the WebKit code shows that it is also dependent on battery saver mode or thermal pressure on the device.

Therefore, it is better to attempt to play the video with audio and handle the failure gracefully.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/autoplay.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Code sharing](/docs/player/integration) [Next\
\
Getting the current time](/docs/player/current-time)

- [Trigger the play from a user event](#trigger-the-play-from-a-user-event)
- [Pass the event to the `play()` or `toggle()` method](#pass-the-event-to-the-play-or-toggle-method)
- [Media that enters the video after the start](#media-that-enters-the-video-after-the-start)
- [Using the `numberOfSharedAudioTags` prop](#using-the-numberofsharedaudiotags-prop)
- [Dealing with an autoplay failure](#dealing-with-an-autoplay-failure)
- [When will a browser restrict autoplay?](#when-will-a-browser-restrict-autoplay)