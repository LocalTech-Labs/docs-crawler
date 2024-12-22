---
title: Using audio
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using audio

- [Home page](/)
- Designing visuals
- Audio

On this page

# Using audio

## Import audio [​](\#import-audio "Direct link to Import audio")

[Put an audio file into the `public/` folder](/docs/assets) and use [`staticFile()`](/docs/staticfile) to reference it.

Add an [`<Audio/>`](/docs/audio) tag to your component to add sound to it.

```

tsx

import { AbsoluteFill, Audio, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} />
    </AbsoluteFill>
  );
};
```

You can also add remote audio by passing a URL:

```

tsx

import { AbsoluteFill, Audio } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src="https://example.com/audio.mp3" />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src="https://example.com/audio.mp3" />
    </AbsoluteFill>
  );
};
```

By default, the audio will play from the start, at full volume and full length.
You can mix multiple tracks together by adding more audio tags.

## Cutting or trimming the audio [​](\#cutting-or-trimming-the-audio "Direct link to Cutting or trimming the audio")

The `<Audio />` tag supports the `startFrom` and `endAt` props. In the following example, we assume that the [`fps`](/docs/composition#fps) of the composition is `30`.

```

tsx

import { AbsoluteFill, Audio, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} startFrom={60} endAt={120} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} startFrom={60} endAt={120} />
    </AbsoluteFill>
  );
};
```

By passing `startFrom={60}`, the playback starts immediately, but with the first 2 seconds of the audio trimmed away.

By passing `endAt={120}`, any audio after the 4 second mark in the file will be trimmed away.

The audio will play the range from `00:02:00` to `00:04:00`, meaning the audio will play for 2 seconds.

## Delaying audio [​](\#delaying-audio "Direct link to Delaying audio")

Use a `<Sequence>` with a positive `from` value to delay the audio from playing.
In the following example, the audio will start playing (from the beginning) after 100 frames.

```

tsx

import { AbsoluteFill, Audio, Sequence, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Sequence from={100}>
        <Audio src={staticFile("audio.mp3")} />
      </Sequence>
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio, Sequence, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Sequence from={100}>
        <Audio src={staticFile("audio.mp3")} />
      </Sequence>
    </AbsoluteFill>
  );
};
```

## Controlling volume [​](\#controlling-volume "Direct link to Controlling volume")

You can use the [`volume`](/docs/audio#volume) prop to control the volume.

The simplest way is to pass a number between 0 and 1. `1` is the maximum volume and `0` mutes the audio.

```

tsx

import { Audio } from "remotion";
import audio from "./audio.mp3";

export const MyComposition = () => {
  return (
    <div>
      <div>Hello World!</div>
      <Audio src={audio} volume={0.5} />
    </div>
  );
};
```

```

tsx

import { Audio } from "remotion";
import audio from "./audio.mp3";

export const MyComposition = () => {
  return (
    <div>
      <div>Hello World!</div>
      <Audio src={audio} volume={0.5} />
    </div>
  );
};
```

You can also change volume over time by passing in a function that takes a frame number and returns the volume.

```

tsx

import { AbsoluteFill, Audio, interpolate, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio
        src={staticFile("audio.mp3")}
        volume={(f) =>
          interpolate(f, [0, 30], [0, 1], { extrapolateLeft: "clamp" })
        }
      />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio, interpolate, staticFile } from "remotion";

export const MyComposition = () => {
  return (
    <AbsoluteFill>
      <Audio
        src={staticFile("audio.mp3")}
        volume={(f) =>
          interpolate(f, [0, 30], [0, 1], { extrapolateLeft: "clamp" })
        }
      />
    </AbsoluteFill>
  );
};
```

In this example we are using the [interpolate()](/docs/interpolate) function to fade the audio in over 30 frames. Note that because values below 0 are not allowed, we need to set the `extrapolateLeft: 'clamp'` option to ensure no negative values.

Inside the callback function, the first frame at which audio is being played is numbered `0`, regardless of the value of `useCurrentFrame()`.

Prefer using a callback function if the volume is changing. This will enable Remotion to draw a volume curve in the timeline and is more performant.

note

When using the [`<Player>`](/docs/player), note that Mobile Safari [does not support the `volume` property](https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/Using_HTML5_Audio_Video/Device-SpecificConsiderations/Device-SpecificConsiderations.html#//apple_ref/doc/uid/TP40009523-CH5-SW11). The audio mix may not play as intended.

## `muted` property [​](\#muted-property "Direct link to muted-property")

You may pass in the `muted` and it may change over time. When `muted` is true, audio will be omitted at that time. In the following example, we are muting the track between frame 40 and 60.

```

tsx

import { AbsoluteFill, Audio, staticFile, useCurrentFrame } from "remotion";

export const MyComposition = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} muted={frame >= 40 && frame <= 60} />
    </AbsoluteFill>
  );
};
```

```

tsx

import { AbsoluteFill, Audio, staticFile, useCurrentFrame } from "remotion";

export const MyComposition = () => {
  const frame = useCurrentFrame();

  return (
    <AbsoluteFill>
      <Audio src={staticFile("audio.mp3")} muted={frame >= 40 && frame <= 60} />
    </AbsoluteFill>
  );
};
```

## Use audio from `<Video />` tags [​](\#use-audio-from-video--tags "Direct link to use-audio-from-video--tags")

Audio from [`<Video />`](/docs/video) and [`<OffthreadVideo />`](/docs/offthreadvideo) tags are also included in the output. You may use the [`startFrom`, `endAt`](/docs/video/#startfrom), [`volume`](/docs/video/#volume) and [`muted`](/docs/video/#muted) props in the same way.

## Controlling playback speed [​](\#controlling-playback-speed "Direct link to Controlling playback speed")

[v2.2](https://github.com/remotion-dev/remotion/releases/v2.2)

You can use the `playbackRate` prop to control the speed of the audio. `1` is the default and means regular speed, `0.5` slows down the audio so it's twice as long and `2` speeds up the audio so it's twice as fast.

While Remotion doesn't limit the range of possible playback speeds, in development mode the [`HTMLMediaElement.playbackRate`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/playbackRate) API is used which throws errors on extreme values. At the time of writing, Google Chrome throws an exception if the playback rate is below `0.0625` or above `16`.

## Audio visualization [​](\#audio-visualization "Direct link to Audio visualization")

You can obtain audio data and create visualizations based on it. See the page [Audio visualization](/docs/audio-visualization) for more info.

## Rendering audio only [​](\#rendering-audio-only "Direct link to Rendering audio only")

Exporting as `mp3`, `aac` and `wav` is supported:

- CLI
- renderMedia()
- Lambda

To render only the audio via CLI, specify an extension when exporting via CLI:

```

npx remotion render src/index.ts my-comp out/audio.mp3
```

```

npx remotion render src/index.ts my-comp out/audio.mp3
```

or use the `--codec` flag to automatically choose a good output file name:

```

npx remotion render src/index.ts my-comp --codec=mp3
```

```

npx remotion render src/index.ts my-comp --codec=mp3
```

To render only the audio via Node.JS, use [`renderMedia()`](/docs/renderer/render-media) and set the [`codec`](/docs/renderer/render-media#codec) to an audio codec.

```

tsx

await renderMedia({
  composition,
  serveUrl: bundleLocation,
  codec: "mp3",
  outputLocation,
  inputProps,
});
```

```

tsx

await renderMedia({
  composition,
  serveUrl: bundleLocation,
  codec: "mp3",
  outputLocation,
  inputProps,
});
```

To render only the audio via Lambda, use [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) and set the [`codec`](/docs/lambda/rendermediaonlambda#codec) to an audio codec and [`imageFormat`](/docs/lambda/rendermediaonlambda#imageformat) to `none`.

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  framesPerLambda: 20,
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "mp3",
  imageFormat: "none",
  maxRetries: 1,
  privacy: "public",
});
```

```

tsx

const { bucketName, renderId } = await renderMediaOnLambda({
  region: "us-east-1",
  functionName: "remotion-render-bds9aab",
  composition: "MyVideo",
  framesPerLambda: 20,
  serveUrl:
    "https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw",
  inputProps: {},
  codec: "mp3",
  imageFormat: "none",
  maxRetries: 1,
  privacy: "public",
});
```

To render via the Lambda CLI, use the [`npx remotion lambda render`](/docs/lambda/cli/render) command and pass the [`--codec`](/docs/lambda/cli/render#--codec) flag:

```

npx remotion lambda render --codec=mp3 https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw my-comp
```

```

npx remotion lambda render --codec=mp3 https://remotionlambda-qg35eyp1s1.s3.eu-central-1.amazonaws.com/sites/bf2jrbfkw my-comp
```

## Order of operations [v4.0.141](https://github.com/remotion-dev/remotion/releases/v4.0.141) [​](\#order-of-operations "Direct link to order-of-operations")

Before Remotion v4.0.141, it was not defined in which order audio operations would be applied and behavior in preview and render could deviate.

Since Remotion v4.0.141, the order of operations is guaranteed to be the following:

1. Trim audio (using [`startFrom`](/docs/audio#startfrom--endat)).
2. Offset audio (by putting it in a [sequence](/docs/terminology/sequence)).
3. Stretch audio (by adding a [`playbackRate`](/docs/audio#playbackrate)).

Example for a 30 FPS composition which is 60 frames long:

1. An [`<Audio>`](/docs/audio) tag has a [`startFrom`](/docs/audio#startfrom--endat) value of 45. The first 1.5 seconds of the audio get trimmed off.
2. The [`<Audio>`](/docs/audio) tag is in a [`<Sequence>`](/docs/sequence) which starts at `30`. The audio only begins playing at the 1.0 second timeline mark at the 1.5 second audio position.
3. The [`<Audio>`](/docs/audio) has a [`playbackRate`](/docs/audio#playbackrate) of `2`. The audio gets sped up by 2x, but the starting position and start offset is not affected.
4. The composition is 60 frames long, so the audio must stop at the 3.5 second mark:

> (comp\_duration - offset) \* playback\_rate + start\_from
>
> (60 - 30) \* 2 + 45 => frame 105 or the 3.5 second mark

5. Result: The section of 1.5sec - 3.5sec gets cut out of the audio and is played in the Remotion timeline between frames 30 and 59 at 2x speed.

## See also [​](\#see-also "Direct link to See also")

- [Importing assets](/docs/assets)
- [Audio visualization](/docs/audio-visualization)
- [`<Audio />`](/docs/audio) tag

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/using-audio.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Transitions](/docs/transitioning) [Next\
\
Fonts](/docs/fonts)

- [Import audio](#import-audio)
- [Cutting or trimming the audio](#cutting-or-trimming-the-audio)
- [Delaying audio](#delaying-audio)
- [Controlling volume](#controlling-volume)
- [`muted` property](#muted-property)
- [Use audio from `<Video />` tags](#use-audio-from-video--tags)
- [Controlling playback speed](#controlling-playback-speed)
- [Audio visualization](#audio-visualization)
- [Rendering audio only](#rendering-audio-only)
- [Order of operations](#order-of-operations)
- [See also](#see-also)
