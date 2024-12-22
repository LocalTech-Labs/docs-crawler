---
title: Preloading assets
source: Remotion Documentation
last_updated: 2024-12-22
---

# Preloading assets

- [Home page](/)
- [Player](/docs/player/)
- Preloading

On this page

# Preloading assets

By default, assets such as videos, audio, or images will only be loaded as they enter the video. When using the [Remotion Player](/docs/terminology/player), you may want to preload those assets beforehand to make them play immediately once they enter the video.

Two ways of preloading are supported:

- Signaling to the browser that assets should be loaded using the [`@remotion/preload`](/docs/preload) APIs
- Fetching the assets beforehand and then playing them locally using [`prefetch()`](/docs/prefetch)

## Preloading videos using `@remotion/preload` [​](\#preloading-videos-using-remotionpreload "Direct link to preloading-videos-using-remotionpreload")

By preloading, a `<link type='preload'>` tag is placed on the page, signaling to the browser that it may start loading the media.

For videos, use [`preloadVideo()`](/docs/preload/preload-video) API, for audio use [`preloadAudio()`](/docs/preload/preload-audio), for images use [`preloadImage()`](/docs/preload/preload-image). Perform the preload outside a component or inside an [`useEffect()`](https://react.dev/reference/react/useEffect).

```

tsx

import {preloadAudio, preloadVideo} from '@remotion/preload';

const unpreloadVideo = preloadVideo('https://example.com/video.mp4');
const unpreloadAudio = preloadAudio(
  'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3',
);

// Later, you can optionally clean up the preload
unpreloadVideo();
unpreloadAudio();
```

```

tsx

import {preloadAudio, preloadVideo} from '@remotion/preload';

const unpreloadVideo = preloadVideo('https://example.com/video.mp4');
const unpreloadAudio = preloadAudio(
  'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3',
);

// Later, you can optionally clean up the preload
unpreloadVideo();
unpreloadAudio();
```

## Prefetching using `prefetch()` [​](\#prefetching-using-prefetch "Direct link to prefetching-using-prefetch")

_Available in v3.2.23_

By prefetching, the full media is downloaded and turned into a Blob URL using [`URL.createObjectURL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/createObjectURL).

If you pass the original URL into either an [`<Audio>`](/docs/audio), [`<Video>`](/docs/video), [`<OffthreadVideo>`](/docs/offthreadvideo) or [`<Img>`](/docs/img) tag and the asset is prefetched, those components will use Blob URL instead.

```

tsx

import {prefetch} from 'remotion';

const {free, waitUntilDone} = prefetch('https://example.com/video.mp4');

waitUntilDone().then(() => {
  console.log('Video has finished loading');
});

// Call free() if you want to un-prefetch and free up the memory:
free();
```

```

tsx

import {prefetch} from 'remotion';

const {free, waitUntilDone} = prefetch('https://example.com/video.mp4');

waitUntilDone().then(() => {
  console.log('Video has finished loading');
});

// Call free() if you want to un-prefetch and free up the memory:
free();
```

## `@remotion/preload` vs. `prefetch()` [​](\#remotionpreload-vs-prefetch "Direct link to remotionpreload-vs-prefetch")

[`prefetch()`](/docs/prefetch) is a more reliable way of ensuring the media is ready when it needs to displayed, but the asset needs to be downloaded in full for it.

[`@remotion/preload`](/docs/preload) is preferrable if the asset is large since you don't have to wait for it finish downloading,

`preloadAudio()`

`preloadVideo()`

`prefetch()`

Works withAll audio and video APIs, images and fonts[`<Audio/>`](/docs/audio)

,

[`<Video/>`](/docs/video)

,

[`<Img/>`](/docs/img)

,

[`<OffthreadVideo/>`](/docs/offthreadvideo)Mechanism

Inserts a `<link type='preload'>` tag

Fetches the asset and turns it into a URL blob or Base 64 URLReadiness

Ready to play when asset is partially loaded

Asset must be fully fetchedReliabilityNo guarantee, just a signal to the browser

Guaranteed to be ready, posssible to track progress of loading

Callback

No way to determine if ready

Ready when promise resolves

Package@remotion/preloadremotionHandles HTTP redirects

Must use

[`resolveRedirect()`](/docs/preload/resolve-redirect)Handled automatically[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

Resource must support CORS if

[`resolveRedirect()`](/docs/preload/resolve-redirect)

is used

Resource must support CORSAvailable fromv3.0.14v3.2.23

## Preloading GIFs [​](\#preloading-gifs "Direct link to Preloading GIFs")

You can preload and preparse GIFs using [`preloadGif()`](/docs/gif/preload-gif)

## See also [​](\#see-also "Direct link to See also")

- [`@remotion/preload`](/docs/preload)
- [`prefetch()`](/docs/prefetch)
- [Avoiding flickering in `<Player>`](/docs/troubleshooting/player-flicker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/player/preloading.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Buffer state](/docs/player/buffer-state) [Next\
\
Premounting](/docs/player/premounting)

- [Preloading videos using `@remotion/preload`](#preloading-videos-using-remotionpreload)
- [Prefetching using `prefetch()`](#prefetching-using-prefetch)
- [`@remotion/preload` vs. `prefetch()`](#remotionpreload-vs-prefetch)
- [Preloading GIFs](#preloading-gifs)
- [See also](#see-also)
