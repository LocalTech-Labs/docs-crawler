---
title: Variable duration and dimensions
source: Remotion Documentation
last_updated: 2024-12-22
---

# Variable duration and dimensions

- [Home page](/)
- [Parameterized videos](/docs/parameterized-rendering)
- Variable duration

On this page

# Variable duration and dimensions

You may change the duration of the video based on some asynchronously determined data.

The same goes for the width, height and framerate of the video.

## Using the `calculateMetadata()` function [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#using-the-calculatemetadata-function "Direct link to using-the-calculatemetadata-function")

Consider a scenario where a video is dynamically specified as a background and the duration of the composition should be aligned with the duration of the video.

Pass a [`calculateMetadata`](/docs/composition#calculatemetadata) callback function to a [`<Composition>`](/docs/composition). This function should take the [combined props](/docs/props-resolution) and calculate the metadata.

```

src/Root.tsx
tsx

import {parseMedia} from '@remotion/media-parser';
import {Composition, Video} from 'remotion';

type MyCompProps = {
  src: string;
};

const MyComp: React.FC<MyCompProps> = ({src}) => {
  return <Video src={src} />;
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      }}
      calculateMetadata={async ({props}) => {
        const {durationInSeconds} = await parseMedia({
          src: props.src,
          fields: {durationInSeconds: true},
        });
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        return {
          durationInFrames: Math.floor(durationInSeconds * 30),
        };
      }}
    />
  );
};
```

```

src/Root.tsx
tsx

import {parseMedia} from '@remotion/media-parser';
import {Composition, Video} from 'remotion';

type MyCompProps = {
  src: string;
};

const MyComp: React.FC<MyCompProps> = ({src}) => {
  return <Video src={src} />;
};

export const Root: React.FC = () => {
  return (
    <Composition
      id="MyComp"
      component={MyComp}
      durationInFrames={300}
      fps={30}
      width={1920}
      height={1080}
      defaultProps={{
        src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      }}
      calculateMetadata={async ({props}) => {
        const {durationInSeconds} = await parseMedia({
          src: props.src,
          fields: {durationInSeconds: true},
        });
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        return {
          durationInFrames: Math.floor(durationInSeconds * 30),
        };
      }}
    />
  );
};
```

The `props` defaults to the `defaultProps` specified, but may be overriden by [passing input props to the render](/docs/props-resolution).

Return an object with `durationInFrames` to change the duration of the video.
In addition or instead, you may also return `fps`, `width` and `height` to update the video's resolution and framerate.

It is also possible to transform the props [passed to the component by returning a `props` field](/docs/data-fetching) at the same time.

## With `useEffect()` and `getInputProps()` [​](\#with-useeffect-and-getinputprops "Direct link to with-useeffect-and-getinputprops")

In the following example, Remotion is instructed to wait for the [`parseMedia()`](/docs/media-parser/parse-media) promise to resolve before evaluating the composition.

By calling [`delayRender()`](/docs/delay-render), Remotion will be blocked from proceeding until [`continueRender()`](/docs/continue-render) is called.

```

src/Root.tsx
tsx

import {parseMedia} from '@remotion/media-parser';

export const Index: React.FC = () => {
  const [handle] = useState(() => delayRender());
  const [duration, setDuration] = useState(1);

  useEffect(() => {
    parseMedia({
      src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      fields: {
        durationInSeconds: true,
      },
    })
      .then(({durationInSeconds}) => {
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        setDuration(Math.round(durationInSeconds * 30));
        continueRender(handle);
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, [handle]);

  return (
    <Composition
      id="dynamic-duration"
      component={VideoTesting}
      width={1080}
      height={1080}
      fps={30}
      durationInFrames={duration}
    />
  );
};
```

```

src/Root.tsx
tsx

import {parseMedia} from '@remotion/media-parser';

export const Index: React.FC = () => {
  const [handle] = useState(() => delayRender());
  const [duration, setDuration] = useState(1);

  useEffect(() => {
    parseMedia({
      src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      fields: {
        durationInSeconds: true,
      },
    })
      .then(({durationInSeconds}) => {
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        setDuration(Math.round(durationInSeconds * 30));
        continueRender(handle);
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, [handle]);

  return (
    <Composition
      id="dynamic-duration"
      component={VideoTesting}
      width={1080}
      height={1080}
      fps={30}
      durationInFrames={duration}
    />
  );
};
```

To dynamically pass a video asset, you [may pass input props](/docs/props-resolution) when rendering and retrieve them inside your React code using [`getInputProps()`](/docs/get-input-props).

```

src/Root.tsx
tsx

import {getInputProps} from 'remotion';

const inputProps = getInputProps();
const src = inputProps.src;
```

```

src/Root.tsx
tsx

import {getInputProps} from 'remotion';

const inputProps = getInputProps();
const src = inputProps.src;
```

### Drawbacks [​](\#drawbacks "Direct link to Drawbacks")

This technique is not recommended anymore since v4.0, because the `useEffect()` does not only get executed when Remotion is initially calculating the metadata of the video, but also when spawning a render worker.

Since a render process might be highly concurrent, this might lead to unnecessary API calls and rate limitations.

## Using together with dimension overrides [​](\#using-together-with-dimension-overrides "Direct link to Using together with dimension overrides")

Override parameters such as [`--width`](/docs/cli/render#--width) will be given priority and override the variable dimensions you set using `calculateMetadata()`.

The [`--scale`](/docs/scaling) parameter has the highest priority and will be applied after override parameters and `calculateMetadata()`.

## Changing dimensions and FPS after the video was designed [​](\#changing-dimensions-and-fps-after-the-video-was-designed "Direct link to Changing dimensions and FPS after the video was designed")

If you designed your video with certain dimensions and then want to render a different resolution (e.g. 4K instead of Full HD), you can use [output scaling](/docs/scaling).

If you designed your video with certain FPS and then want to change the frame rate, you should refactor the composition to ensure [the timing stays the same with changing frame rates](/docs/multiple-fps).

## With the `<Player>` [​](\#with-the-player "Direct link to with-the-player")

The [`<Player>`](/docs/player) will react if the metadata being passed to it changes. There are two viable ways to do dynamically set the metadata of the Player:

[1](#1)

Perform data fetching using `useEffect()`:

```

MyApp.tsx
tsx

import {parseMedia} from '@remotion/media-parser';
import {useEffect, useState} from 'react';
import {Player} from '@remotion/player';

export const Index: React.FC = () => {
  const [duration, setDuration] = useState(1);

  useEffect(() => {
    parseMedia({
      src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      fields: {
        durationInSeconds: true,
      },
    })
      .then(({durationInSeconds}) => {
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        setDuration(Math.round(durationInSeconds * 30));
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, []);

  return (
    <Player
      component={VideoTesting}
      compositionWidth={1080}
      compositionHeight={1080}
      fps={30}
      durationInFrames={duration}
    />
  );
};
```

```

MyApp.tsx
tsx

import {parseMedia} from '@remotion/media-parser';
import {useEffect, useState} from 'react';
import {Player} from '@remotion/player';

export const Index: React.FC = () => {
  const [duration, setDuration] = useState(1);

  useEffect(() => {
    parseMedia({
      src: 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4',
      fields: {
        durationInSeconds: true,
      },
    })
      .then(({durationInSeconds}) => {
        if (durationInSeconds === null) {
          throw new Error('Could not read duration of video');
        }

        setDuration(Math.round(durationInSeconds * 30));
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, []);

  return (
    <Player
      component={VideoTesting}
      compositionWidth={1080}
      compositionHeight={1080}
      fps={30}
      durationInFrames={duration}
    />
  );
};
```

[2](#2)

Call your own `calculateMetadata()` function reused
from your Remotion project:

```

MyApp.tsx
tsx

import {Player} from '@remotion/player';

type Props = {};

const calculateMetadataFunction: CalculateMetadataFunction<Props> = () => {
  return {
    props: {},
    durationInFrames: 1,
    width: 100,
    height: 100,
    fps: 30,
  };
};

type Metadata = {
  durationInFrames: number;
  compositionWidth: number;
  compositionHeight: number;
  fps: number;
  props: Props;
};

export const Index: React.FC = () => {
  const [metadata, setMetadata] = useState<Metadata | null>(null);

  useEffect(() => {
    Promise.resolve(
      calculateMetadataFunction({
        defaultProps: {},
        props: {},
        abortSignal: new AbortController().signal,
        compositionId: 'MyComp',
      }),
    )
      .then(({durationInFrames, props, width, height, fps}) => {
        setMetadata({
          durationInFrames: durationInFrames as number,
          compositionWidth: width as number,
          compositionHeight: height as number,
          fps: fps as number,
          props: props as Props,
        });
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, []);

  if (!metadata) {
    return null;
  }

  return <Player component={VideoTesting} {...metadata} />;
};
```

```

MyApp.tsx
tsx

import {Player} from '@remotion/player';

type Props = {};

const calculateMetadataFunction: CalculateMetadataFunction<Props> = () => {
  return {
    props: {},
    durationInFrames: 1,
    width: 100,
    height: 100,
    fps: 30,
  };
};

type Metadata = {
  durationInFrames: number;
  compositionWidth: number;
  compositionHeight: number;
  fps: number;
  props: Props;
};

export const Index: React.FC = () => {
  const [metadata, setMetadata] = useState<Metadata | null>(null);

  useEffect(() => {
    Promise.resolve(
      calculateMetadataFunction({
        defaultProps: {},
        props: {},
        abortSignal: new AbortController().signal,
        compositionId: 'MyComp',
      }),
    )
      .then(({durationInFrames, props, width, height, fps}) => {
        setMetadata({
          durationInFrames: durationInFrames as number,
          compositionWidth: width as number,
          compositionHeight: height as number,
          fps: fps as number,
          props: props as Props,
        });
      })
      .catch((err) => {
        console.log(`Error fetching metadata: ${err}`);
      });
  }, []);

  if (!metadata) {
    return null;
  }

  return <Player component={VideoTesting} {...metadata} />;
};
```

## See also [​](\#see-also "Direct link to See also")

- [How props get resolved](/docs/props-resolution)
- [Data fetching](/docs/data-fetching)
- [How do I make the composition the same duration as my video?](/docs/miscellaneous/snippets/align-duration)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/dynamic-metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Data fetching](/docs/data-fetching) [Next\
\
How props get resolved](/docs/props-resolution)

- [Using the `calculateMetadata()` function](#using-the-calculatemetadata-function)
- [With `useEffect()` and `getInputProps()`](#with-useeffect-and-getinputprops)
  - [Drawbacks](#drawbacks)
- [Using together with dimension overrides](#using-together-with-dimension-overrides)
- [Changing dimensions and FPS after the video was designed](#changing-dimensions-and-fps-after-the-video-was-designed)
- [With the `<Player>`](#with-the-player)
- [See also](#see-also)
