---
title: parseMedia() vs. getVideoMetadata()
source: Remotion Documentation
last_updated: 2024-12-22
---

# parseMedia() vs. getVideoMetadata()

- [Home page](/)
- FAQ
- parseMedia() vs. getVideoMetadata()

On this page

# parseMedia() vs. getVideoMetadata()

There are multiple ways to retrieve information about a media such as:

- Duration
- Width and height
- Codec, container, framerate

## Recommendation: `parseMedia()` [​](\#recommendation-parsemedia "Direct link to recommendation-parsemedia")

[`parseMedia()`](/docs/media-parser/parse-media) is a library developed by Remotion that parses video files using JavaScript.

### Advantages [​](\#advantages "Direct link to Advantages")

✅ It supports more containers than [`getVideoMetadata()`](/docs/get-video-metadata): MP4, MOV, WebM, AVI, MKV

✅ It supports getting get dimensions of videos shot on iPhones, a common pitfall when using [`getVideoMetadata()`](/docs/get-video-metadata) on Linux

✅ It also works on the server in Node.js and Bun.

✅ It only fetches the information you request, needing to read less bytes to get the information you need.

### Disadvantages [​](\#disadvantages "Direct link to Disadvantages")

❌ The API is still unstable, although the signature for simple use cases should not change.

❌ When using in the browser: Assets need to come from the same origin or be CORS-enabled.

## `getVideoMetadata()` from `@remotion/media-utils` [​](\#getvideometadata-from-remotionmedia-utils "Direct link to getvideometadata-from-remotionmedia-utils")

[`getVideoMetadata()`](/docs/get-video-metadata) will mount a `<video>` tag in the browser and returns the metadata after it is available.

### Advantages [​](\#advantages-1 "Direct link to Advantages")

✅ In the browser, it does not require the assets to be CORS-enabled or be on the same origin.

### Disadvantages [​](\#disadvantages-1 "Direct link to Disadvantages")

❌ It only works in the browser.

❌ It does not support some codecs, such as the ones used by iPhone videos.

## `getVideoMetadata()` from `@remotion/renderer` [​](\#getvideometadata-from-remotionrenderer "Direct link to getvideometadata-from-remotionrenderer")

[`getVideoMetadata()`](/docs/renderer/get-video-metadata) will read metadata using FFmpeg through a Rust interface.

### Advantages [​](\#advantages-2 "Direct link to Advantages")

✅ The widest format compatibility.

### Disadvantages [​](\#disadvantages-2 "Direct link to Disadvantages")

❌ It only works in Node.js and Bun, so you cannot use it in [`calculateMetadata()`](/docs/calculate-metadata).

❌ It invokes a Rust binary, so bundling it is not straightforward.

## See also [​](\#see-also "Direct link to See also")

- [`parseMedia()`](/docs/media-parser/parse-media)
- [`@remotion/media-utils`: `getVideoMetadata()`](/docs/get-video-metadata)
- [`@remotion/renderer`: `getVideoMetadata()`](/docs/renderer/get-video-metadata)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/parse-media-vs-get-video-metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Live streaming](/docs/miscellaneous/live-streaming) [Next\
\
Difference to Motion Canvas](/docs/compare/motion-canvas)

- [Recommendation: `parseMedia()`](#recommendation-parsemedia)
  - [Advantages](#advantages)
  - [Disadvantages](#disadvantages)
- [`getVideoMetadata()` from `@remotion/media-utils`](#getvideometadata-from-remotionmedia-utils)
  - [Advantages](#advantages-1)
  - [Disadvantages](#disadvantages-1)
- [`getVideoMetadata()` from `@remotion/renderer`](#getvideometadata-from-remotionrenderer)
  - [Advantages](#advantages-2)
  - [Disadvantages](#disadvantages-2)
- [See also](#see-also)
