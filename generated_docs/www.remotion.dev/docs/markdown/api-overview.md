---
title: API overview
source: Remotion Documentation
last_updated: 2024-12-21
---

# API overview

- [Home page](/)
- API overview

# API overview

[**Command line** \
\
Reference for the `npx remotion` commands](/docs/cli) [**Configuration file** \
\
Reference for the `remotion.config.ts` file](/docs/config)

## remotion

Core APIs: `useCurrentFrame()`, `interpolate()`, etc.

[**<Composition>** \
\
Define a video](/docs/composition) [**<Still>** \
\
Define a still](/docs/still) [**<Folder>** \
\
Organize compositions in the Studio sidebar](/docs/folder) [**registerRoot()** \
\
Initialize a Remotion project](/docs/register-root) [**useCurrentFrame()** \
\
Obtain the current time](/docs/use-current-frame) [**useVideoConfig()** \
\
Get the duration, dimensions and FPS of a composition](/docs/use-video-config) [**interpolate()** \
\
Map a range of values to another](/docs/interpolate) [**spring()** \
\
Physics-based animation primitive](/docs/spring) [**interpolateColors()** \
\
Map a range of values to colors](/docs/interpolate-colors) [**measureSpring()** \
\
Determine the duration of a spring](/docs/measure-spring) [**Easing** \
\
Customize animation curve of `interpolate()`](/docs/easing) [**<Img>** \
\
Render an `<img>` tag and wait for it to load](/docs/img) [**<Video>** \
\
Synchronize a `<video>` with Remotion's time](/docs/video) [**<Audio>** \
\
Synchronize `<audio>` with Remotion's time](/docs/audio) [**<OffthreadVideo>** \
\
Alternative to `<Video>`](/docs/offthreadvideo) [**<IFrame>** \
\
Render an `<iframe>` tag and wait for it to load](/docs/iframe) [**<Sequence>** \
\
Time-shifts it's children](/docs/sequence) [**<Series>** \
\
Display contents after another](/docs/series) [**<Freeze>** \
\
Freeze some content in time](/docs/freeze) [**<Loop>** \
\
Play some content repeatedly](/docs/loop) [**delayRender()** \
\
Block a render from continuing](/docs/delay-render) [**continueRender()** \
\
Unblock a render](/docs/continue-render) [**cancelRender()** \
\
Abort an error](/docs/cancel-render) [**getInputProps()** \
\
Receive the user-defined input data](/docs/get-input-props) [**getRemotionEnvironment()** \
\
Determine if you are currently previewing or rendering](/docs/get-remotion-environment) [**staticFile()** \
\
Access file from `public/` folder](/docs/staticfile) [**<AbsoluteFill>** \
\
Position content absolutely and in full size](/docs/absolute-fill) [**VERSION** \
\
Get the current version of Remotion](/docs/version)

## @remotion/bundler

Create a Webpack bundle from Node.JS

[**bundle()** \
\
Create a Webpack bundle](/docs/bundle)

## @remotion/player

Play a Remotion video in the browser.

[**<Player>** \
\
Embed a Remotion composition in a web app](/docs/player/player) [**<Thumbnail>** \
\
Embed a still in a web app](/docs/player/thumbnail)

## @remotion/lambda

Render videos and stills on AWS Lambda

[**estimatePrice()** \
\
Estimate the price of a render](/docs/lambda/estimateprice) [**deployFunction()** \
\
Create a new function in AWS Lambda](/docs/lambda/deployfunction) [**deleteFunction()** \
\
Delete a function in AWS Lambda](/docs/lambda/deletefunction) [**getFunctionInfo()** \
\
Gets information about a function](/docs/lambda/getfunctioninfo) [**getFunctions()** \
\
Lists available Remotion Lambda functions](/docs/lambda/getfunctions) [**getCompositionsOnLambda()** \
\
Gets list of compositions inside a Lambda function](/docs/lambda/getcompositionsonlambda) [**deleteSite()** \
\
Delete a bundle from S3](/docs/lambda/deletesite) [**deploySite()** \
\
Bundle and upload a site to S3](/docs/lambda/deploysite) [**getAwsClient()** \
\
Access the AWS SDK directly](/docs/lambda/getawsclient) [**getRegions()** \
\
Get all available regions](/docs/lambda/getregions) [**getSites()** \
\
Get all available sites](/docs/lambda/getsites) [**downloadMedia()** \
\
Download a render artifact from S3](/docs/lambda/downloadmedia) [**getUserPolicy()** \
\
Get the policy JSON for your AWS user](/docs/lambda/getuserpolicy) [**getRolePolicy()** \
\
Get the policy JSON for your AWS role](/docs/lambda/getrolepolicy) [**getOrCreateBucket()** \
\
Ensure a Remotion S3 bucket exists](/docs/lambda/getorcreatebucket) [**getRenderProgress()** \
\
Query the progress of a render](/docs/lambda/getrenderprogress) [**presignUrl()** \
\
Make a private file public to those with the link](/docs/lambda/presignurl) [**renderMediaOnLambda()** \
\
Trigger a video or audio render](/docs/lambda/rendermediaonlambda) [**renderStillOnLambda()** \
\
Trigger a still render](/docs/lambda/renderstillonlambda) [**simulatePermissions()** \
\
Ensure permissions are correctly set up](/docs/lambda/simulatepermissions) [**speculateFunctionName()** \
\
Get the lambda function name based on its configuration](/docs/lambda/speculatefunctionname) [**validateWebhookSignature()** \
\
Validate an incoming webhook request is authentic](/docs/lambda/validatewebhooksignature)

## @remotion/cloudrun

Render videos and stills on GCP Cloud Run

[**getServiceInfo()** \
\
Gets information about a service](/docs/cloudrun/getserviceinfo) [**deployService()** \
\
Create a new service in GCP Cloud Run](/docs/cloudrun/deployservice) [**deleteService()** \
\
Delete a service in GCP Cloud Run](/docs/cloudrun/deleteservice) [**getServices()** \
\
Lists available Remotion Cloud Run services](/docs/cloudrun/getservices) [**speculateServiceName()** \
\
Speculate a service name based on its configuration](/docs/cloudrun/speculateservicename) [**getRegions()** \
\
Get all available regions](/docs/cloudrun/getregions) [**deploySite()** \
\
Bundle and upload a site to Cloud Storage](/docs/cloudrun/deploysite) [**deleteSite()** \
\
Delete a bundle from Cloud Storage](/docs/cloudrun/deletesite) [**getSites()** \
\
Get all available sites from Cloud Storage](/docs/cloudrun/getsites) [**getOrCreateBucket()** \
\
Ensure a Remotion Cloud Storage bucket exists](/docs/cloudrun/getorcreatebucket) [**renderMediaOnCloudrun()** \
\
Trigger a video or audio render](/docs/cloudrun/rendermediaoncloudrun) [**renderStillOnCloudrun()** \
\
Trigger a still render](/docs/cloudrun/renderstilloncloudrun) [**testPermissions()** \
\
Ensure permissions are correctly set up in GCP](/docs/cloudrun/testpermissions)

## @remotion/captions

Common operations for subtitles.

[**Caption** \
\
An object shape for captions](/docs/captions/caption) [**parseSrt()** \
\
Parse a .srt file into a `Caption` array](/docs/captions/parse-srt) [**serializeSrt()** \
\
Serialize a .srt file into a `Caption` array](/docs/captions/serialize-srt) [**createTikTokStyleCaptions()** \
\
Structure the captions for TikTok-style display](/docs/captions/create-tiktok-style-captions)

## @remotion/gif

Include a GIF in your video.

[**<Gif>** \
\
Render a GIF](/docs/gif/gif) [**getGifDurationInSeconds()** \
\
Get the runtime of a GIF](/docs/gif/get-gif-duration-in-seconds) [**preloadGif()** \
\
Prepare a GIF for displaying in the Player](/docs/gif/preload-gif)

## @remotion/media-utils

Obtain info about video and audio.

[**audioBufferToDataUrl()** \
\
Serialize an audio buffer](/docs/audio-buffer-to-data-url) [**getAudioData()** \
\
Get metadata of an audio source](/docs/get-audio-data) [**getAudioDurationInSeconds()** \
\
Get the duration of an audio source](/docs/get-audio-duration-in-seconds) [**getVideoMetadata()** \
\
Get metadata of a video source](/docs/get-video-metadata) [**getWaveformPortion()** \
\
Trims audio data into a waveform](/docs/get-waveform-portion) [**useAudioData()** \
\
`getAudioData()` as a hook](/docs/use-audio-data) [**useWindowedAudioData()** \
\
Optimized for fetching only current data, works only with `.wav`](/docs/use-windowed-audio-data) [**visualizeAudio()** \
\
Processes a waveform for visualization](/docs/visualize-audio)

## @remotion/animation-utils

Obtain info about video and audio.

[**makeTransform()** \
\
Create a value for the CSS `transform` property](/docs/animation-utils/make-transform) [**interpolateStyles()** \
\
Map a range of values to CSS `style` values](/docs/animation-utils/interpolate-styles)

## @remotion/tailwind

Webpack override for using TailwindCSS

[**enableTailwind()** \
\
Override the Webpack config to enable TailwindCSS](/docs/tailwind/enable-tailwind)

## @remotion/enable-scss

Webpack override for enabling SASS/SCSS

[**enableScss()** \
\
Override the Webpack config to enable SCSS](/docs/enable-scss/enable-scss)

## @remotion/three

Create 3D videos using React Three Fiber

[**<ThreeCanvas>** \
\
A wrapper for React Three Fiber' Canvas](/docs/three-canvas) [**useVideoTexture(** \
\
Use a video in React Three Fiber](/docs/use-video-texture) [**useOffthreadVideoTexture()** \
\
Use an <OffthreadVideo> in React Three Fiber](/docs/use-offthread-video-texture)

## @remotion/skia

Low-level graphics using React Native Skia

[**enableSkia()** \
\
Webpack override for enabling Skia](/docs/skia/enable-skia) [**<SkiaCanvas>** \
\
React Native Skia <Canvas> wrapper](/docs/skia/skia-canvas)

## @remotion/lottie

Include a Lottie animation in your video

[**<Lottie>** \
\
Embed a Lottie animation in Remotion](/docs/lottie/lottie) [**getLottieMetadata()** \
\
Get metadata of a Lottie animation](/docs/lottie/getlottiemetadata) [**staticFile()** \
\
Load Lottie animations from a static file](/docs/lottie/staticfile)

## @remotion/preload

Preload media for the Player

[**preloadVideo()** \
\
Preload a video source](/docs/preload/preload-video) [**preloadAudio()** \
\
Preload an audio source](/docs/preload/preload-audio) [**preloadFont()** \
\
Preload a font](/docs/preload/preload-font) [**preloadImage()** \
\
Preload an image](/docs/preload/preload-image) [**resolveRedirect()** \
\
Get the definitive URL after all redirects](/docs/preload/preload-audio)

## @remotion/renderer

Render video, audio and stills from Node.JS or Bun

[**getCompositions()** \
\
List available compositions](/docs/renderer/get-compositions) [**selectComposition()** \
\
Get a composition](/docs/renderer/select-composition) [**renderMedia()** \
\
Render a video or audio](/docs/renderer/render-media) [**renderFrames()** \
\
Render a series of images](/docs/renderer/render-frames) [**renderStill()** \
\
Render a single image](/docs/renderer/render-still) [**stitchFramesToVideo()** \
\
Turn images into a video](/docs/renderer/stitch-frames-to-video) [**openBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/open-browser) [**ensureBrowser()** \
\
Open a Chrome browser to reuse across renders](/docs/renderer/ensure-browser) [**makeCancelSignal()** \
\
Create token to later cancel a render](/docs/renderer/make-cancel-signal) [**getVideoMetadata()** \
\
Get metadata from a video file in Node.js](/docs/renderer/get-video-metadata) [**getSilentParts()** \
\
Obtain silent portions of a video or audio](/docs/renderer/get-silent-parts) [**ensureFfmpeg()** \
\
Check for ffmpeg binary and install if not existing](/docs/renderer/ensure-ffmpeg) [**ensureFfprobe()** \
\
Check for ffprobe binary and install if not existing](/docs/renderer/ensure-ffprobe) [**getCanExtractFramesFast()** \
\
Probes for fast extraction for <OffthreadVideo>](/docs/renderer/get-can-extract-frames-fast)

## @remotion/paths

Manipulate and obtain info about SVG paths

[**getLength()** \
\
Obtain length of an SVG path](/docs/paths/get-length) [**getPointAtLength()** \
\
Get coordinates at a certain point of an SVG path](/docs/paths/get-point-at-length) [**getTangentAtLength()** \
\
Gets tangents `x` and `y` of a point which is on an SVG path](/docs/paths/get-tangent-at-length) [**reversePath()** \
\
Switch direction of an SVG path](/docs/paths/reverse-path) [**normalizePath()** \
\
Replace relative with absolute coordinates](/docs/paths/normalize-path) [**interpolatePath()** \
\
Interpolates between two SVG paths](/docs/paths/interpolate-path) [**evolvePath()** \
\
Animate an SVG path](/docs/paths/evolve-path) [**translatePath()** \
\
Translates the position of an path against X/Y coordinates](/docs/paths/translate-path) [**warpPath()** \
\
Remap the coordinates of a path](/docs/paths/warp-path) [**scalePath()** \
\
Grow or shrink the size of the path](/docs/paths/scale-path) [**getBoundingBox()** \
\
Get the bounding box of a SVG path](/docs/paths/get-bounding-box) [**resetPath()** \
\
Translates an SVG path to `(0, 0)`](/docs/paths/reset-path) [**extendViewBox()** \
\
Widen an SVG viewBox in all directions](/docs/paths/extend-viewbox) [**getSubpaths()** \
\
Split SVG path into its parts](/docs/paths/get-subpaths) [**parsePath()** \
\
Parse a string into an array of instructions](/docs/paths/parse-path) [**serializeInstructions()** \
\
Turn an array of instructions into a SVG path](/docs/paths/serialize-instructions) [**reduceInstructions()** \
\
Reduce the amount of instruction types](/docs/paths/reduce-instructions)

## @remotion/noise

Generate noise effects

[**noise2D()** \
\
Create 2D noise](/docs/noise/noise-2d) [**noise3D()** \
\
Create 3D noise](/docs/noise/noise-3d) [**noise4D()** \
\
Create 4D noise](/docs/noise/noise-4d)

## @remotion/shapes

Generate SVG shapes

[**makeRect()** \
\
Generate SVG Path for a rect](/docs/shapes/make-rect) [**<Rect/>** \
\
Render a rect](/docs/shapes/rect) [**makeCircle()** \
\
Generate SVG Path for a circle](/docs/shapes/make-circle) [**<Circle/>** \
\
Render a circle](/docs/shapes/circle) [**makePie()** \
\
Generate SVG Path for a pie](/docs/shapes/make-pie) [**<Pie/>** \
\
Render a pie](/docs/shapes/pie) [**makeEllipse()** \
\
Generate SVG Path for a ellipse](/docs/shapes/make-ellipse) [**<Ellipse/>** \
\
Render a ellipse](/docs/shapes/ellipse) [**makeTriangle()** \
\
Generate SVG Path for a triangle](/docs/shapes/make-triangle) [**<Triangle/>** \
\
Render a triangle](/docs/shapes/triangle) [**makeStar()** \
\
Generate SVG Path for a star](/docs/shapes/make-star) [**<Star/>** \
\
Render a star](/docs/shapes/star) [**makePolygon()** \
\
Generate SVG Path for a polygon](/docs/shapes/make-polygon) [**<Polygon/>** \
\
Render a polygon](/docs/shapes/polygon)

## @remotion/studio

APIs for controlling theRemotion Studio

[**getStaticFiles()** \
\
Get a list of files in the `public` folder](/docs/studio/get-static-files) [**watchPublicFolder()** \
\
Listen to changes in the public folder](/docs/studio/watch-public-folder) [**watchStaticFile()** \
\
Listen to changes of a static file](/docs/studio/watch-static-file) [**writeStaticFile()** \
\
Save content to a file in the public directory](/docs/studio/write-static-file) [**saveDefaultProps()** \
\
Save default props to the root file](/docs/studio/save-default-props) [**updateDefaultProps()** \
\
Update default props in the Props editor](/docs/studio/update-default-props) [**deleteStaticFile()** \
\
Delete a file from the public directory](/docs/studio/delete-static-file) [**restartStudio()** \
\
Restart the Studio Server.](/docs/studio/restart-studio) [**focusDefaultPropsPath()** \
\
Scrolls to a specific field in the default props editor](/docs/studio/focus-default-props-path) [**reevaluateComposition()** \
\
Re-runs calculateMetadata() on the current composition](/docs/studio/reevaluate-composition)

## @remotion/transitions

Transition between scenes

### Components

[**`<TransitionSeries>`** \
\
A `<Series>` with transitions inbetween](/docs/transitions/transitionseries)

### Timings

[**`springTiming()`** \
\
Transition with a `spring()`](/docs/transitions/timings/springtiming) [**`linearTiming()`** \
\
Transition linearly with optional Easing](/docs/transitions/timings/lineartiming)

### Presentations

Hover over an effect to see the preview.

[A\
\
B\
\
**`fade()`** \
\
Animate the opacity of the scenes](/docs/transitions/presentations/fade) [A\
\
B\
\
**`slide()`** \
\
Slide in and push out the previous scene](/docs/transitions/presentations/slide) [A\
\
B\
\
**`wipe()`** \
\
Slide over the previous scene](/docs/transitions/presentations/wipe) [A\
\
B\
\
**`flip()`** \
\
Rotate the previous scene](/docs/transitions/presentations/flip) [A\
\
B\
\
**`clockWipe()`** \
\
Reveal the new scene in a circular movement](/docs/transitions/presentations/clock-wipe) [A\
\
B\
\
**`cube()`** \
\
Paid\
\
Rotate both scenes with 3D perspective](/docs/transitions/presentations/cube) [A\
\
B\
\
**`none()`** \
\
Have no visual effect.](/docs/transitions/presentations/none)

## @remotion/layout-utils

Layout helpers

[**measureText()** \
\
Get dimensions of text](/docs/layout-utils/measure-text) [**fillTextBox()** \
\
Find line breaks and overflows in a text box](/docs/layout-utils/fill-text-box) [**fitText()** \
\
Get font size to fit text in a box](/docs/layout-utils/fit-text)

## @remotion/install-whisper-cpp

Whisper.cpp installation and transcription

[**installWhisperCpp()** \
\
Install the whisper.cpp software](/docs/install-whisper-cpp/install-whisper-cpp) [**downloadWhisperModel()** \
\
Download a Whisper model](/docs/install-whisper-cpp/download-whisper-model) [**transcribe()** \
\
Transcribe an audio file](/docs/install-whisper-cpp/transcribe) [**convertToCaptions()** \
\
Postprocessing for TikTok-style captions](/docs/install-whisper-cpp/convert-to-captions)

## @remotion/openai-whisper

Work with transcriptions from OpenAI Whisper

[**openAiWhisperApiToCaptions()** \
\
Turn OpenAI Whisper API transcriptions into an array of `Caption`](/docs/openai-whisper/openai-whisper-api-to-captions)

## @remotion/animated-emoji

Google Fonts Animated Emojis as Remotion Components

[**<AnimatedEmoji>** \
\
Component for rendering an animated emoji.](/docs/animated-emoji/animated-emoji) [**getAvailableEmoji()** \
\
Get a list of available emoji.](/docs/animated-emoji/get-available-emoji)

## @remotion/google-fonts

Load Google Fonts onto a page.

[**loadFont()** \
\
Load a Google Font](/docs/google-fonts/load-font) [**getAvailableFonts()** \
\
Static list of available fonts](/docs/google-fonts/get-available-fonts) [**getInfo()** \
\
Metadata about a specific font](/docs/google-fonts/get-info)

## @remotion/rive

Embed Rive animations in Remotion

[**<RemotionRiveCanvas>** \
\
Render a Rive animation](/docs/rive/remotionrivecanvas)

## @remotion/zod-types

Zod types enabling Remotion Studio UI

[**zColor()** \
\
A Zod Type for colors](/docs/zod-types/z-color) [**zTextarea()** \
\
A Zod Type for multiple-line text in a textarea](/docs/zod-types/z-textarea)

## @remotion/motion-blur

Apply motion blur effects to components

[**<Trail>** \
\
Add a trail effect to children](/docs/motion-blur/trail) [**<CameraMotionBlur>** \
\
Add a natural camera motion blur effect to children](/docs/motion-blur/camera-motion-blur)

## @remotion/fonts

Load font files onto a page.

[**loadFont()** \
\
Load a font from a URL or a local file](/docs/fonts-api/load-font)

## @remotion/media-parser

A pure JavaScript library for parsing video files

[**parseMedia()** \
\
Parse a media file.](/docs/media-parser/parse-media) [**nodeReader** \
\
Read a file from the local file system.](/docs/media-parser/node-reader) [**fetchReader** \
\
Read a file using `fetch()`.](/docs/media-parser/fetch-reader) [**webFileReader** \
\
Read a file from `<input type="file">`.](/docs/media-parser/web-file-reader)

## @remotion/webcodecs

Converting media using WebCodecs

[**convertMedia()** \
\
Converts a video using WebCodecs and Media Parser](/docs/webcodecs/convert-media) [**getAvailableContainers()** \
\
Get a list of containers `@remotion/webcodecs` supports.](/docs/webcodecs/get-available-containers) [**canReencodeVideoTrack()** \
\
Determine if a video track can be re-encoded](/docs/webcodecs/can-reencode-video-track) [**canReencodeAudioTrack()** \
\
Determine if a audio track can be re-encoded](/docs/webcodecs/can-reencode-audio-track) [**canCopyVideoTrack()** \
\
Determine if a video track can be copied without re-encoding](/docs/webcodecs/can-copy-video-track) [**canCopyAudioTrack()** \
\
Determine if a audio track can be copied without re-encoding](/docs/webcodecs/can-copy-audio-track) [**getDefaultAudioCodec()** \
\
Gets the default audio codec for a container if no other audio codec is specified.](/docs/webcodecs/get-default-audio-codec) [**getDefaultVideoCodec()** \
\
Gets the default video codec for a container if no other audio codec is specified.](/docs/webcodecs/get-default-video-codec) [**defaultOnAudioTrackHandler()** \
\
The default track transformation function for audio tracks.](/docs/webcodecs/default-on-audio-track-handler) [**defaultOnVideoTrackHandler()** \
\
The default track transformation function for video tracks.](/docs/webcodecs/default-on-video-track-handler) [**getAvailableAudioCodecs()** \
\
Get the audio codecs that can fit in a container.](/docs/webcodecs/get-available-audio-codecs) [**getAvailableVideoCodecs()** \
\
Get the video codecs that can fit in a container.](/docs/webcodecs/get-available-video-codecs)

## @remotion/licensing

Report and query company license usage

[**registerUsagePoint()** \
\
Register a cloud render or WebCodecs conversion](/docs/licensing/register-usage-point) [**getUsage()** \
\
Query usage of company license](/docs/licensing/get-usage)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/api.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 21, 2024**

[Next\
\
CLI reference](/docs/cli/)
