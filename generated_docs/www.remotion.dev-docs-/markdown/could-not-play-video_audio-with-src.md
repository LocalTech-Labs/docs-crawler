---
title: Could not play video/audio with src
source: Remotion Documentation
last_updated: 2024-12-22
---

# Could not play video/audio with src

- [Home page](/)
- Troubleshooting
- Media playback error

On this page

# Could not play video/audio with src

The following error may occur in your Remotion project during development or while rendering:

```

diff

- Error: Could not play video with src [source] [object MediaError]
```

```

diff

- Error: Could not play video with src [source] [object MediaError]
```

or

```

diff

- Error: Could not play audio with src [source] [object MediaError]
```

```

diff

- Error: Could not play audio with src [source] [object MediaError]
```

This error happens when you are trying to embed a `<Video/>` or `<Audio/>` tag in your Remotion project but the browser is unable to load and play the media. Although the browser does not report the exact error, there are two common reasons for this error.

## Codec not supported by Chromium [​](\#codec-not-supported-by-chromium "Direct link to Codec not supported by Chromium")

Unlike Google Chrome, the Chromium Browser does not include proprietary codecs. This means you cannot play MP4/H.264 videos and some audio codecs (more codecs may not be supported).

**Workaround 1**: Use Chrome, not Chromium.

**Workaround 2**: Convert videos to WebM, [use Chrome instead of Chromium](/docs/config#setbrowserexecutable), or use [`<OffthreadVideo>`](/docs/offthreadvideo).

note

Since `v4.0.18`, [downloads a browser](/docs/miscellaneous/chrome-headless-shell) that includes codecs automatically if you don't have a local browser.

## Invalid source [​](\#invalid-source "Direct link to Invalid source")

Make sure you are trying to load a video that is either available locally or publicly on the internet. Open the DevTools, go to the Network tab, track the asset being loaded, and open it in a new tab. Does it actually start playing?

## Wrong headers or status code [​](\#wrong-headers-or-status-code "Direct link to Wrong headers or status code")

In order for the browser being able to play the media, it should be loaded with:

- a 200 status code
- a suitable `Content-Type` header
- A `Content-Range` header in order to support seeking.

Open the DevTools and go to the network tab to validate that this is the case.

## Internet Download Manager [​](\#internet-download-manager "Direct link to Internet Download Manager")

[Internet Download Manager](https://www.internetdownloadmanager.com/) is known to manipulate the network traffic causing problems for the browser to load media. Disable it if you have it.

## Other unsupported codecs [​](\#other-unsupported-codecs "Direct link to Other unsupported codecs")

You might be importing a video that is not compatible with Chrome at all, e.g. FLV.

## Too many video tags [​](\#too-many-video-tags "Direct link to Too many video tags")

The error message might contain `error creating media player`, appearing if too many video tags are created.

First check that you are not accidentially creating an infite render loop. For example, changing the `key` frequently will re-create the video tag on every frame:

```

tsx

import { Video } from "remotion";

export default function SBSVideo() {
  return (
    <>
      {[
        "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
        "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
      ].map((video, i) => {
        return <Video key={uuidv4()} src={video} />;
      })}
    </>
  );
}
```

```

tsx

import { Video } from "remotion";

export default function SBSVideo() {
  return (
    <>
      {[
        "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
        "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
      ].map((video, i) => {
        return <Video key={uuidv4()} src={video} />;
      })}
    </>
  );
}
```

If you ruled out this possibility, use [`<OffthreadVideo>`](/docs/offthreadvideo) instead as it does not rely on a `<video>` tag.

## Recover from this error [v3.3.89](https://github.com/remotion-dev/remotion/releases/v3.3.89) [​](\#recover-from-this-error "Direct link to recover-from-this-error")

You can handle this error and replace it with a different video by passing the `onError()` prop to the `<Video>` or `<OffthreadVideo>` component.

## See also [​](\#see-also "Direct link to See also")

- [Codec support](/docs/config#setbrowserexecutable)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-playback-error.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Target closed](/docs/target-closed) [Next\
\
Performance](/docs/performance)

- [Codec not supported by Chromium](#codec-not-supported-by-chromium)
- [Invalid source](#invalid-source)
- [Wrong headers or status code](#wrong-headers-or-status-code)
- [Internet Download Manager](#internet-download-manager)
- [Other unsupported codecs](#other-unsupported-codecs)
- [Too many video tags](#too-many-video-tags)
- [Recover from this error](#recover-from-this-error)
- [See also](#see-also)
