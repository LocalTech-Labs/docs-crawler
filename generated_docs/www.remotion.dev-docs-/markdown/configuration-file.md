---
title: Configuration file
source: Remotion Documentation
last_updated: 2024-12-22
---

# Configuration file

- [Home page](/)
- Configuration file

On this page

# Configuration file

To configure Remotion, create a `remotion.config.ts` file in the root of your Remotion project.

These options will apply to CLI commands such as `npx remotion studio` and `npx remotion render`.

warning

The configuration file has no effect when using [SSR](/docs/renderer) APIs.

You can control several behaviors of Remotion here.

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.setConcurrency(8);
Config.setPixelFormat('yuv444p');
Config.setCodec('h265');
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.setConcurrency(8);
Config.setPixelFormat('yuv444p');
Config.setCodec('h265');
```

## `overrideWebpackConfig()` [v1.1.0](https://github.com/remotion-dev/remotion/releases/v1.1.0) [​](\#overridewebpackconfig "Direct link to overridewebpackconfig")

Allows you to insert your custom Webpack config. [See the page about custom Webpack configs](/docs/webpack) for more information.

```

remotion.config.ts
ts

Config.overrideWebpackConfig((currentConfiguration) => {
  // Return a new Webpack configuration
  return {
    ...currentConfiguration,
    // new configuration
  };
});
```

```

remotion.config.ts
ts

Config.overrideWebpackConfig((currentConfiguration) => {
  // Return a new Webpack configuration
  return {
    ...currentConfiguration,
    // new configuration
  };
});
```

## `setCachingEnabled()` [v2.0.0](https://github.com/remotion-dev/remotion/releases/v2.0.0) [​](\#setcachingenabled "Direct link to setcachingenabled")

Enable or disable Webpack caching. Default is `true` which will make the Webpack step in the first run a bit slower but will massively speed up subsequent runs.

```

remotion.config.ts
ts

Config.setCachingEnabled(false);
```

```

remotion.config.ts
ts

Config.setCachingEnabled(false);
```

The [command line flag](/docs/cli/render#--bundle-cache) `--bundle-cache` will take precedence over this option.

## `setStudioPort()` [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61) [​](\#setstudioport "Direct link to setstudioport")

Set the HTTP port for the Studio.

```

remotion.config.ts
ts

Config.setStudioPort(3003);
```

```

remotion.config.ts
ts

Config.setStudioPort(3003);
```

The [command line flag](/docs/cli/studio#--port) `--port` will take precedence over this option.

## `setRendererPort()` [v4.0.61](https://github.com/remotion-dev/remotion/releases/v4.0.61) [​](\#setrendererport "Direct link to setrendererport")

Set the port to be used to host the Webpack bundle.

```

remotion.config.ts
ts

Config.setStudioPort(3004);
```

```

remotion.config.ts
ts

Config.setStudioPort(3004);
```

The [command line flag](/docs/cli/render#--port) `--port` will take precedence over this option.

## `setPublicDir()` [v3.2.13](https://github.com/remotion-dev/remotion/releases/v3.2.13) [​](\#setpublicdir "Direct link to setpublicdir")

Define the location of the [`public/ directory`](/docs/terminology/public-dir). If not defined, Remotion will assume the location is the \`public\` folder in your Remotion root.

```

remotion.config.ts
ts

Config.setPublicDir('./custom-public-dir');
```

```

remotion.config.ts
ts

Config.setPublicDir('./custom-public-dir');
```

The [command line flag](/docs/cli/render#--public-dir) `--public-dir` will take precedence over this option.

## `setEntryPoint()` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#setentrypoint "Direct link to setentrypoint")

Sets the Remotion [entry point](/docs/terminology/entry-point), you don't have to specify it for CLI commands.

```

remotion.config.ts
ts

Config.setEntryPoint('./src/index.ts');
```

```

remotion.config.ts
ts

Config.setEntryPoint('./src/index.ts');
```

If you pass an entry point as a CLI argument, it will take precedence.

## `setLevel()` [v2.0.1](https://github.com/remotion-dev/remotion/releases/v2.0.1) [​](\#setlevel "Direct link to setlevel")

Increase or decrease the amount of log messages in the CLI.
Acceptable values:

- `error`: Silent except error messages.
- `warn`: Only showing errors and warnings.
- `info` ( _default_): Default output - besides errors and warnings, prints progress and output location.
- `verbose`: All of the above, plus browser logs and other debug info.

```

remotion.config.ts
ts

Config.setLevel('verbose');
```

```

remotion.config.ts
ts

Config.setLevel('verbose');
```

The [command line flag](/docs/cli/render#--log) `--log` will take precedence over this option.

## `setMaxTimelineTracks()` [v2.1.10](https://github.com/remotion-dev/remotion/releases/v2.1.10) [​](\#setmaxtimelinetracks "Direct link to setmaxtimelinetracks")

Set how many tracks are being displayed in the timeline in the Studio at most. This does not affect your video, just the amount of tracks shown when previewing. Default `15`.

```

remotion.config.ts
ts

Config.setMaxTimelineTracks(20);
```

```

remotion.config.ts
ts

Config.setMaxTimelineTracks(20);
```

## `setKeyboardShortcutsEnabled()` [v3.2.11](https://github.com/remotion-dev/remotion/releases/v3.2.11) [​](\#setkeyboardshortcutsenabled "Direct link to setkeyboardshortcutsenabled")

Whether the Studio should react to keyboard shortcuts. Default `true`.

```

remotion.config.ts
ts

Config.setKeyboardShortcutsEnabled(false);
```

```

remotion.config.ts
ts

Config.setKeyboardShortcutsEnabled(false);
```

The [command line flag](/docs/cli/studio#--disable-keyboard-shortcuts) `--disable-keyboard-shortcuts` will take precedence over this option.

## `setWebpackPollingInMilliseconds()` [v3.3.11](https://github.com/remotion-dev/remotion/releases/v3.3.11) [​](\#setwebpackpollinginmilliseconds "Direct link to setwebpackpollinginmilliseconds")

Enables Webpack polling instead of the file system event listeners for hot reloading.
This is useful if you are inside a virtual machine or have a remote file system.

```

remotion.config.ts
ts

Config.setWebpackPollingInMilliseconds(1000);
```

```

remotion.config.ts
ts

Config.setWebpackPollingInMilliseconds(1000);
```

The [command line flag](/docs/cli/studio#--webpack-poll) `--webpack-poll` will take precedence over this option.

## `setNumberOfSharedAudioTags()` [v3.3.2](https://github.com/remotion-dev/remotion/releases/v3.3.2) [​](\#setnumberofsharedaudiotags "Direct link to setnumberofsharedaudiotags")

How many shared audio tags should be mounted in the Studio. Shared audio tags can help prevent playback issues due to audio autoplay policies of the browser. See [this article](/docs/player/autoplay#use-the-numberofsharedaudiotags-property) which covers the same option but for the Player. Default `0`, meaning no autoplay policies are circumvented.

```

remotion.config.ts
ts

Config.setNumberOfSharedAudioTags(5);
```

```

remotion.config.ts
ts

Config.setNumberOfSharedAudioTags(5);
```

## `setShouldOpenBrowser()` [v3.3.19](https://github.com/remotion-dev/remotion/releases/v3.3.19) [​](\#setshouldopenbrowser "Direct link to setshouldopenbrowser")

Whether Remotion should open a browser when starting the Studio. Default `true`.

```

remotion.config.ts
ts

Config.setShouldOpenBrowser(false);
```

```

remotion.config.ts
ts

Config.setShouldOpenBrowser(false);
```

## `setBrowserExecutable()` [v1.5.0](https://github.com/remotion-dev/remotion/releases/v1.5.0) [​](\#setbrowserexecutable "Direct link to setbrowserexecutable")

Set a custom Chrome or Chromium executable path. By default Remotion will try to find an existing version of Chrome on your system and if not found, it will download one. This flag is useful if you don't have Chrome installed in a standard location and you want to prevent downloading an additional browser or need [support for the H264 codec](/docs/video#codec-support).

```

remotion.config.ts
ts

Config.setBrowserExecutable('/usr/bin/google-chrome-stable');
```

```

remotion.config.ts
ts

Config.setBrowserExecutable('/usr/bin/google-chrome-stable');
```

The [command line flag](/docs/cli/render#--browser-executable) `--browser-executable` will take precedence over this option.

## `setDelayRenderTimeoutInMilliseconds()` [v2.6.3](https://github.com/remotion-dev/remotion/releases/v2.6.3) [​](\#setdelayrendertimeoutinmilliseconds "Direct link to setdelayrendertimeoutinmilliseconds")

_previously named "setTimeoutInMilliseconds"_

Define how long a single frame may take to resolve all [`delayRender()`](/docs/delay-render) calls [before it times out](/docs/timeout). Default: `30000`

```

remotion.config.ts
ts

Config.setDelayRenderTimeoutInMilliseconds(60000);
```

```

remotion.config.ts
ts

Config.setDelayRenderTimeoutInMilliseconds(60000);
```

The [command line flag](/docs/cli/render#--timeout) `--timeout` will take precedence over this option.

## `setChromiumDisableWebSecurity()` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#setchromiumdisablewebsecurity "Direct link to setchromiumdisablewebsecurity")

This will most notably disable CORS among other security features during rendering.

```

remotion.config.ts
tsx

Config.setChromiumDisableWebSecurity(true);
```

```

remotion.config.ts
tsx

Config.setChromiumDisableWebSecurity(true);
```

The [command line flag](/docs/cli/render#--disable-web-security) `--disable-web-security` will take precedence over this option.

## `setChromiumIgnoreCertificateErrors()` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#setchromiumignorecertificateerrors "Direct link to setchromiumignorecertificateerrors")

Results in invalid SSL certificates, such as self-signed ones, being ignored during rendering.

```

remotion.config.ts
tsx

Config.setChromiumIgnoreCertificateErrors(true);
```

```

remotion.config.ts
tsx

Config.setChromiumIgnoreCertificateErrors(true);
```

The [command line flag](/docs/cli/render#--ignore-certificate-errors) `--ignore-certificate-errors` will take precedence over this option.

## `setChromiumHeadlessMode()` [v2.6.5](https://github.com/remotion-dev/remotion/releases/v2.6.5) [​](\#setchromiumheadlessmode "Direct link to setchromiumheadlessmode")

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

```

remotion.config.ts
tsx

Config.setChromiumHeadlessMode(false);
```

```

remotion.config.ts
tsx

Config.setChromiumHeadlessMode(false);
```

## `setChromiumMultiProcessOnLinux()` [v4.0.42](https://github.com/remotion-dev/remotion/releases/v4.0.42) [​](\#setchromiummultiprocessonlinux "Direct link to setchromiummultiprocessonlinux")

Removes the `--single-process` flag that gets passed to Chromium on Linux by default. This will make the render faster because multiple processes can be used, but may cause issues with some Linux distributions or if window server libraries are missing.

Default: `false` until v4.0.136, then `true` from v4.0.137 on because newer Chrome versions don't allow rendering with the `--single-process` flag.

This flag will be removed in Remotion v5.0.

```

remotion.config.ts
tsx

Config.setChromiumMultiProcessOnLinux(true);
```

```

remotion.config.ts
tsx

Config.setChromiumMultiProcessOnLinux(true);
```

The [command line flag](/docs/cli/render#--enable-multiprocess-on-linux) `--enable-multiprocess-on-linux` will take precedence over this option.

## `setChromiumOpenGlRenderer()` [​](\#setchromiumopenglrenderer "Direct link to setchromiumopenglrenderer")

Changelog

- From Remotion v2.6.7 until v3.0.7, the default for Remotion Lambda was `swiftshader`, but from v3.0.8 the default is `swangle` (Swiftshader on Angle) since Chrome 101 added support for it.
- From Remotion v2.4.3 until v2.6.6, the default was `angle`, however it turns out to have a small memory leak that could crash long Remotion renders.

Select the OpenGL renderer backend for Chromium.

Accepted values:

- `"angle"`
- `"egl"`
- `"swiftshader"`
- `"swangle"`
- `"vulkan"` ( _from Remotion v4.0.41_)
- `"angle-egl"` ( _from Remotion v4.0.51_)

The default is `null`, letting Chrome decide, except on Lambda where the default is `"swangle"`

```

remotion.config.ts
tsx

Config.setChromiumOpenGlRenderer('angle');
```

```

remotion.config.ts
tsx

Config.setChromiumOpenGlRenderer('angle');
```

The [command line flag](/docs/cli/render#--gl) `--gl` will take precedence over this option.

## `setConcurrency()` [​](\#setconcurrency "Direct link to setconcurrency")

Sets how many Puppeteer instances will work on rendering your video in parallel.
Default: `null`, meaning **half of the threads** available on your CPU.

```

remotion.config.ts
ts

Config.setConcurrency(8);
```

```

remotion.config.ts
ts

Config.setConcurrency(8);
```

The [command line flag](/docs/cli/render#--concurrency) `--concurrency` will take precedence over this option.

tip

Try to set your concurrency to `os.cpus().length` to all the threads available on your CPU for faster rendering. The drawback is that other parts of your system might slow down.

## `setVideoImageFormat()` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#setvideoimageformat "Direct link to setvideoimageformat")

Determines which in which image format to render the frames. Either:

- `jpeg` \- the fastest option (default)
- `png` \- slower, but supports transparency
- `none` \- don't render images, just calculate audio information

```

remotion.config.ts
ts

Config.setVideoImageFormat('png');
```

```

remotion.config.ts
ts

Config.setVideoImageFormat('png');
```

## `setStillImageFormat()` [v4.0.0](https://github.com/remotion-dev/remotion/releases/v4.0.0) [​](\#setstillimageformat "Direct link to setstillimageformat")

Determines which in which image format to render the frames. Either:

- `png` (default)
- `jpeg`
- `pdf`
- `webp`

```

remotion.config.ts
ts

Config.setStillImageFormat('pdf');
```

```

remotion.config.ts
ts

Config.setStillImageFormat('pdf');
```

## `setScale()` [v2.6.7](https://github.com/remotion-dev/remotion/releases/v2.6.7) [​](\#setscale "Direct link to setscale")

[Scales the output frames by the factor you pass in.](/docs/scaling) For example, a 1280x720px frame will become a 1920x1080px frame with a scale factor of `1.5`. Vector elements like fonts and HTML markups will be rendered with extra details. Default: `1`.

```

remotion.config.ts
ts

Config.setScale(2);
```

```

remotion.config.ts
ts

Config.setScale(2);
```

The [command line flag](/docs/cli/render#--scale) `--scale` will take precedence over this option.

## `setMuted()` [v3.2.1](https://github.com/remotion-dev/remotion/releases/v3.2.1) [​](\#setmuted "Direct link to setmuted")

Disables audio output. Default `false`.

```

remotion.config.ts
ts

Config.setMuted(true);
```

```

remotion.config.ts
ts

Config.setMuted(true);
```

The [command line flag](/docs/cli/render#--muted) `--muted` will take precedence over this option.

## `setEnforceAudioTrack()` [v3.2.1](https://github.com/remotion-dev/remotion/releases/v3.2.1) [​](\#setenforceaudiotrack "Direct link to setenforceaudiotrack")

Render a silent audio track if there would be none otherwise. Default `false`.

```

remotion.config.ts
ts

Config.setEnforceAudioTrack(true);
```

```

remotion.config.ts
ts

Config.setEnforceAudioTrack(true);
```

The [command line flag](/docs/cli/render#--enforce-audio-track) `--enforce-audio-track` will take precedence over this option.

## `setForSeamlessAacConcatenation()` [v4.0.123](https://github.com/remotion-dev/remotion/releases/v4.0.123) [​](\#setforseamlessaacconcatenation "Direct link to setforseamlessaacconcatenation")

If enabled, the audio is trimmed to the nearest AAC frame, which is required for seamless concatenation of AAC files. This is a requirement if you later want to combine multiple video snippets seamlessly.

This option is used internally. There is currently no documentation yet for to concatenate the audio chunks.

```

remotion.config.ts
ts

Config.setForSeamlessAacConcatenation(true);
```

```

remotion.config.ts
ts

Config.setForSeamlessAacConcatenation(true);
```

The [command line flag](/docs/cli/render#--for-seamless-aac-concatenation) `--for-seamless-aac-concatenation` will take precedence over this option.

## `setFrameRange()` [v2.0.0](https://github.com/remotion-dev/remotion/releases/v2.0.0) [​](\#setframerange "Direct link to setframerange")

Pass a number to render a still frame or a tuple to render a subset of a video. The frame sequence is zero-indexed.

```

remotion.config.ts
ts

Config.setFrameRange(90); // To render only the 91st frame
```

```

remotion.config.ts
ts

Config.setFrameRange(90); // To render only the 91st frame
```

or

```

remotion.config.ts
ts

Config.setFrameRange([0, 20]); // Render a video only containing the first 21 frames
```

```

remotion.config.ts
ts

Config.setFrameRange([0, 20]); // Render a video only containing the first 21 frames
```

The [command line flag](/docs/cli/render#--frames) `--frames` will take precedence over this option.

## `setJpegQuality()` [​](\#setjpegquality "Direct link to setjpegquality")

The JPEG quality of each frame. Must be a number between 0 and 100. Will not work if you render PNG frames. [Default: 80](https://github.com/chromium/chromium/blob/99314be8152e688bafbbf9a615536bdbb289ea87/headless/lib/browser/protocol/headless_handler.cc#L32).

```

remotion.config.ts
ts

Config.setJpegQuality(90);
```

```

remotion.config.ts
ts

Config.setJpegQuality(90);
```

The [command line flag](/docs/cli/render#--jpeg-quality) `--jpeg-quality` will take precedence over this option.

## `setDotEnvLocation()` [​](\#setdotenvlocation "Direct link to setdotenvlocation")

Set a custom location for a [`.env`](https://www.npmjs.com/package/dotenv) file. You can specify an absolute path or a relative path in which case it gets resolved based on the current working directory.

```

remotion.config.ts
ts

Config.setDotEnvLocation('.my-env');
```

```

remotion.config.ts
ts

Config.setDotEnvLocation('.my-env');
```

The [command line flag](/docs/cli/render#--env-file) `--env-file` will take precedence over this option.

## `setEveryNthFrame()` [​](\#seteverynthframe "Direct link to seteverynthframe")

This option may only be set when rendering GIFs. [It determines how many frames are rendered, while the other ones gets skipped in order to lower the FPS of the GIF.](/docs/render-as-gif)

For example, if the `fps` is 30, and `everyNthFrame` is 2, the FPS of the GIF is `15`.

```

remotion.config.ts
ts

Config.setEveryNthFrame(2);
```

```

remotion.config.ts
ts

Config.setEveryNthFrame(2);
```

The [command line flag](/docs/cli/render#--every-nth-frame) `--every-nth-frame` will take precedence over this option.

## `setNumberOfGifLoops()` [​](\#setnumberofgifloops "Direct link to setnumberofgifloops")

Allows you to set the number of loops as follows:

- `null` (or omitting in the CLI) plays the GIF indefinitely.
- `0` disables looping
- `1` loops the GIF once (plays twice in total)
- `2` loops the GIF twice (plays three times in total) and so on.

```

remotion.config.ts
ts

Config.setNumberOfGifLoops(2);
```

```

remotion.config.ts
ts

Config.setNumberOfGifLoops(2);
```

The [command line flag](/docs/cli/render#--number-of-gif-loops) `--number-of-gif-loops` will take precedence over this option.

## `setOutputLocation()` [v3.1.6](https://github.com/remotion-dev/remotion/releases/v3.1.6) [​](\#setoutputlocation "Direct link to setoutputlocation")

Set the output location of the video or still, relative to the current working directory. The default is `out/{composition}.{container}`. For example, `out/HelloWorld.mp4`.

```

remotion.config.ts
ts

Config.setOutputLocation('out/video.mp4');
```

```

remotion.config.ts
ts

Config.setOutputLocation('out/video.mp4');
```

If you pass another argument to the render command, it will take precedence: `npx remotion render src/index.ts HelloWorld out/video.mp4`.

## `setOverwriteOutput()` [​](\#setoverwriteoutput "Direct link to setoverwriteoutput")

Set this to `false` to prevent overwriting Remotion outputs when they already exists.

```

remotion.config.ts
ts

Config.setOverwriteOutput(false);
```

```

remotion.config.ts
ts

Config.setOverwriteOutput(false);
```

info

In version 1.x, the default behavior was inverse - Remotion would not override by default.

## `setPixelFormat()` [​](\#setpixelformat "Direct link to setpixelformat")

Controls the pixel format in FFmpeg. [Read more about it here.](https://trac.ffmpeg.org/wiki/Chroma%20Subsampling) Acceptable values: `yuv420p`, `yuv422p`, `yuv444p`, `yuv420p10le`, `yuv422p10le`, `yuv444p10le`. Since v1.4, `yuva420p` is also supported for transparent WebM videos. Since v2.1.7, `yuva444p10le` is also supported for transparent ProRes videos
Default value: `yuv420p`

```

remotion.config.ts
ts

Config.setPixelFormat('yuv420p');
```

```

remotion.config.ts
ts

Config.setPixelFormat('yuv420p');
```

The [command line flag](/docs/cli/render#--pixel-format) `--pixel-format` will take precedence over this option.

## `setCodec()` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#setcodec "Direct link to setcodec")

Choose one of the supported codecs: `h264` _(default)_, `h265`, `vp8`, `vp9`.

- `h264` is the classic MP4 file as you know it.
- `h265` is the successor of H264, with smaller file sizes. Also known as HEVC. Poor browser compatibility.
- `vp8` is the codec for WebM.
- `vp9` is the next-generation codec for WebM. Lower file size, longer compression time.
- `prores` is a common codec if you want to import the output into another video editing program _(available from v2.1.6)_
- `mp3` will export audio only as an MP3 file _(available from v2.0)_
- `wav` will export audio only as an WAV file _(available from v2.0)_
- `aac` will export audio only as an AAC file _(available from v2.0)_

```

remotion.config.ts
ts

Config.setCodec('h265');
```

```

remotion.config.ts
ts

Config.setCodec('h265');
```

The [command line flag](/docs/cli/render#--codec) `--codec` will take precedence over this option.

**See also**: [Encoding guide](/docs/encoding)

## `setAudioCodec()` [​](\#setaudiocodec "Direct link to setaudiocodec")

```

remotion.config.ts
ts

Config.setAudioCodec('pcm-16');
```

```

remotion.config.ts
ts

Config.setAudioCodec('pcm-16');
```

Choose the encoding of your audio.

- The default is dependent on the chosen `codec`.
- Choose `pcm-16` if you need uncompressed audio.
- Not all video containers support all audio codecs.
- This option takes precedence if the `codec` option also specifies an audio codec.

The [command line flag](/docs/cli/render#--audio-codec) `--audio-codec` will take precedence over this option.

Refer to the [Encoding guide](/docs/encoding) to see defaults and supported combinations.

## `setProResProfile()` [v2.1.6](https://github.com/remotion-dev/remotion/releases/v2.1.6) [​](\#setproresprofile "Direct link to setproresprofile")

Set the ProRes profile. This option is only valid if the codec has been set to `prores`.
Possible values: `4444-xq`, `4444`, `hq`, `standard`, `light`, `proxy`.
See [here](https://video.stackexchange.com/a/14715) for explanation of possible values.
Default: `hq`

```

remotion.config.ts
ts

Config.setProResProfile('4444');
```

```

remotion.config.ts
ts

Config.setProResProfile('4444');
```

The [command line flag](/docs/cli/render#--prores-profile) `--prores-profile` will take precedence over this option.

## `setX264Preset()` [v4.2.2](https://github.com/remotion-dev/remotion/releases/v4.2.2) [​](\#setx264preset "Direct link to setx264preset")

Set the Preset profile. This option is only valid if the codec has been set to `h264`.
Possible values: `superfast`, `veryfast`, `faster`, `fast`, `medium`, `slow`, `slower`, `veryslow`, `placebo`,
Default: `medium`

```

remotion.config.ts
ts

Config.setX264Preset('fast');
```

```

remotion.config.ts
ts

Config.setX264Preset('fast');
```

The [command line flag](/docs/cli/render#--x264-preset) `--x264-preset` will take precedence over this option.

**See also**: [Encoding guide](/docs/encoding), [Transparent videos](/docs/transparent-videos)

## `setImageSequence()` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#setimagesequence "Direct link to setimagesequence")

Set to true if you want to output an image sequence instead of a video.

```

remotion.config.ts
ts

Config.setImageSequence(true);
```

```

remotion.config.ts
ts

Config.setImageSequence(true);
```

The [command line flag](/docs/cli/render#--sequence) `--sequence` will take precedence over this option.

## `overrideHeight()` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#overrideheight "Direct link to overrideheight")

Overrides the height of the rendered video.

```

remotion.config.ts
ts

Config.overrideHeight(600);
```

```

remotion.config.ts
ts

Config.overrideHeight(600);
```

The [command line flag](/docs/cli/render#--height) `--height` will take precedence over this option.
(see h264 validation?)

## `overrideWidth()` [v3.2.40](https://github.com/remotion-dev/remotion/releases/v3.2.40) [​](\#overridewidth "Direct link to overridewidth")

Overrides the width of the rendered video.

```

remotion.config.ts
ts

Config.overrideWidth(900);
```

```

remotion.config.ts
ts

Config.overrideWidth(900);
```

The [command line flag](/docs/cli/render#--width) `--width` will take precedence over this option

## `setCrf()` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#setcrf "Direct link to setcrf")

The "Constant Rate Factor" (CRF) of the output. [Use this setting to tell FFmpeg how to trade off size and quality.](/docs/encoding#controlling-quality-using-the-crf-setting)

Ranges for CRF scale, by codec:

- `h264` crf range is 1-51 where crf 18 is _default_.
- `h265` crf range is 0-51 where crf 23 is _default_.
- `vp8` crf range is 4-63 where crf 9 is _default_.
- `vp9` crf range is 0-63 where crf 28 is _default_.

The lowest value is lossless, and the highest value is the worst quality possible. Higher values decrease the filesize at the cost of quality.

The range is exponential, so increasing the CRF value +6 results in roughly half the bitrate / file size, while -6 leads to roughly twice the bitrate.

Choose the highest CRF value that still provides an acceptable quality. If the output looks good, then try a higher value. If it looks bad, choose a lower value.

```

remotion.config.ts
ts

Config.setCrf(16);
```

```

remotion.config.ts
ts

Config.setCrf(16);
```

The [command line flag](/docs/cli/render#--crf) `--crf` will take precedence over this option.

## `setVideoBitrate()` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#setvideobitrate "Direct link to setvideobitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:v` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact video bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps.

```

remotion.config.ts
ts

Config.setVideoBitrate('1M');
```

```

remotion.config.ts
ts

Config.setVideoBitrate('1M');
```

The [command line flag](/docs/cli/render#--video-bitrate) `--video-bitrate` will take precedence over this option.

## `setEncodingBufferSize()` [v4.0.78](https://github.com/remotion-dev/remotion/releases/v4.0.78) [​](\#setencodingbuffersize "Direct link to setencodingbuffersize")

The value for the `-bufsize` flag of FFmpeg. Should be used in conjunction with the encoding max rate flag.

```

remotion.config.ts
ts

Config.setEncodingBufferSize('10000k');
```

```

remotion.config.ts
ts

Config.setEncodingBufferSize('10000k');
```

The [command line flag](/docs/cli/render#--buffer-size) `--buffer-size` will take precedence over this option.

## `setEncodingMaxRate()` [v4.0.78](https://github.com/remotion-dev/remotion/releases/v4.0.78) [​](\#setencodingmaxrate "Direct link to setencodingmaxrate")

The value for the `-maxrate` flag of FFmpeg. Should be used in conjunction with the encoding buffer size flag.

```

remotion.config.ts
ts

Config.setEncodingMaxRate('5000k');
```

```

remotion.config.ts
ts

Config.setEncodingMaxRate('5000k');
```

The [command line flag](/docs/cli/render#--max-rate) `--max-rate` will take precedence over this option.

## `setAudioBitrate()` [v3.2.32](https://github.com/remotion-dev/remotion/releases/v3.2.32) [​](\#setaudiobitrate "Direct link to setaudiobitrate")

Specify the target bitrate for the generated video. The syntax for FFmpeg's `-b:a` parameter should be used. FFmpeg may encode the video in a way that will not result in the exact audio bitrate specified. Example values: `512K` for 512 kbps, `1M` for 1 Mbps. Default: `320k`

```

remotion.config.ts
ts

Config.setAudioBitrate('128K');
```

```

remotion.config.ts
ts

Config.setAudioBitrate('128K');
```

The [command line flag](/docs/cli/render#--audio-bitrate) `--audio-bitrate` will take precedence over this option.

## `setEnableFolderExpiry()` [v4.0.32](https://github.com/remotion-dev/remotion/releases/v4.0.32) [​](\#setenablefolderexpiry "Direct link to setenablefolderexpiry")

When deploying sites, enable or disable S3 Lifecycle policies which allow for renders to auto-delete after a certain time. Default is `null`, which does not change any lifecycle policies of the S3 bucket. See: [Lambda autodelete](/docs/lambda/autodelete).

```

remotion.config.ts
ts

Config.setEnableFolderExpiry(true);
```

```

remotion.config.ts
ts

Config.setEnableFolderExpiry(true);
```

## `setLambdaInsights()` [v4.0.115](https://github.com/remotion-dev/remotion/releases/v4.0.115) [​](\#setlambdainsights "Direct link to setlambdainsights")

Enable [Lambda Insights in AWS CloudWatch](https://remotion.dev/docs/lambda/insights). For this to work, you may have to update your role permission.

```

remotion.config.ts
ts

Config.setLambdaInsights(true);
```

```

remotion.config.ts
ts

Config.setLambdaInsights(true);
```

## `setDeleteAfter()` [v4.0.32](https://github.com/remotion-dev/remotion/releases/v4.0.32) [​](\#setdeleteafter "Direct link to setdeleteafter")

Automatically delete the render after a certain period. Accepted values are `1-day`, `3-days`, `7-days` and `30-days`.

For this to work, your bucket needs to have [lifecycles enabled](/docs/lambda/autodelete).

```

remotion.config.ts
ts

Config.setDeleteAfter('3-days');
```

```

remotion.config.ts
ts

Config.setDeleteAfter('3-days');
```

## `setBeepOnFinish()` [v4.0.84](https://github.com/remotion-dev/remotion/releases/v4.0.84) [​](\#setbeeponfinish "Direct link to setbeeponfinish")

Whether the Remotion Studio tab should beep when the render is finished.

```

remotion.config.ts
ts

Config.setBeepOnFinish(true);
```

```

remotion.config.ts
ts

Config.setBeepOnFinish(true);
```

The [command line flag](/docs/cli/studio#--beep-on-finish) `--beep-on-finish` will take precedence over this option.

## `setBufferStateDelayInMilliseconds()` [v4.0.111](https://github.com/remotion-dev/remotion/releases/v4.0.111) [​](\#setbufferstatedelayinmilliseconds "Direct link to setbufferstatedelayinmilliseconds")

Set the amount of milliseconds after which the Player in the Studio will display a buffering UI after the Player has entered a buffer state. Default `300`.

```

remotion.config.ts
ts

Config.setBufferStateDelayInMilliseconds(0);
```

```

remotion.config.ts
ts

Config.setBufferStateDelayInMilliseconds(0);
```

## `setBinariesDirectory()` [v4.0.120](https://github.com/remotion-dev/remotion/releases/v4.0.120) [​](\#setbinariesdirectory "Direct link to setbinariesdirectory")

The directory where the platform-specific binaries and libraries that Remotion needs are located. Those include an `ffmpeg` and `ffprobe` binary, a Rust binary for various tasks, and various shared libraries. If the value is set to `null`, which is the default, then the path of a platform-specific package located at `node_modules/@remotion/compositor-*` is selected.

This option is useful in environments where Remotion is not officially supported to run like bundled serverless functions or Electron.

```

remotion.config.ts
ts

Config.setBinariesDirectory('/path/to/custom/directory');
```

```

remotion.config.ts
ts

Config.setBinariesDirectory('/path/to/custom/directory');
```

## `setPreferLosslessAudio()` [v4.0.123](https://github.com/remotion-dev/remotion/releases/v4.0.123) [​](\#setpreferlosslessaudio "Direct link to setpreferlosslessaudio")

Uses a lossless audio codec, if one is available for the codec. If you set `audioCodec`, it takes priority over `preferLossless`.

```

remotion.config.ts
ts

Config.setPreferLosslessAudio(true);
```

```

remotion.config.ts
ts

Config.setPreferLosslessAudio(true);
```

## `setHardwareAcceleration()` [v4.0.228](https://github.com/remotion-dev/remotion/releases/v4.0.228) [​](\#sethardwareacceleration "Direct link to sethardwareacceleration")

One of
"disable", "if-possible", or "required"
. Default "disable". Encode using a hardware-accelerated encoder if
available. If set to "required" and no hardware-accelerated encoder is
available, then the render will fail.

```

remotion.config.ts
ts

Config.setHardwareAcceleration('if-possible');
```

```

remotion.config.ts
ts

Config.setHardwareAcceleration('if-possible');
```

## `overrideFfmpegCommand()` [v3.2.22](https://github.com/remotion-dev/remotion/releases/v3.2.22) [​](\#overrideffmpegcommand "Direct link to overrideffmpegcommand")

Modifies the FFmpeg command that Remotion uses under the hood. It works reducer-style, meaning that you pass a function that takes a command as an argument and returns a new command.

```

remotion.config.ts
tsx

Config.overrideFfmpegCommand(({args}) => {
  // Define the custom FFmpeg options as an array of strings
  const customFfmpegOptions = [
    '-profile:v',
    'main',
    '-video_track_timescale',
    '90000',
    '-color_primaries',
    'bt709',
    '-color_trc',
    'bt709',
    '-strict',
    'experimental',
  ];
  // The customFfmpegOptions are inserted before the last element to ensure
  // they appear before the ffmpeg's output path
  args.splice(args.length - 1, 0, ...customFfmpegOptions);
  return args;
});
```

```

remotion.config.ts
tsx

Config.overrideFfmpegCommand(({args}) => {
  // Define the custom FFmpeg options as an array of strings
  const customFfmpegOptions = [
    '-profile:v',
    'main',
    '-video_track_timescale',
    '90000',
    '-color_primaries',
    'bt709',
    '-color_trc',
    'bt709',
    '-strict',
    'experimental',
  ];
  // The customFfmpegOptions are inserted before the last element to ensure
  // they appear before the ffmpeg's output path
  args.splice(args.length - 1, 0, ...customFfmpegOptions);
  return args;
});
```

The function you pass must accept an object as it's only parameter which contains the following properties:

- `type`: Either `"stitcher"` or `"pre-stitcher"`. If enough memory and CPU is available, Remotion may use parallel rendering and encoding, which means that a pre-stitcher process gets spawned before all frames are rendered. You can tell whether parallel encoding is enabled by adding `--log=verbose` to your render command.
- `args`: An array of strings that is passed as arguments to the FFmpeg command.

Your function must return a modified array of strings.

warning

Using this feature is discouraged. Before using it, we want to make you aware of some caveats:

- The render command can change with any new Remotion version, even when it is a patch upgrade. This might break your usage of this feature.
- Depending on the selected codec, available CPU and RAM, Remotion may or may not use "parallel encoding" which will result in multiple FFmpeg commands being executed. Your function must be able to handle being called multiple times.
- The FFmpeg binary provided by Remotion supports only a small subset of FFmpeg commands, therefore not every passed option will be applied.
- This feature is not available when using Remotion Lambda.

Before you use this hack, reach out to the Remotion team on [Discord](https://remotion.dev/discord) and ask us if we are open to implement the feature you need in a clean way - we often do implement new features quickly based on users feedback.

## `setQuality()` [​](\#setquality "Direct link to setquality")

Renamed to `setJpegQuality` in `v4.0.0`.

## `setFfmpegExecutable()` [​](\#setffmpegexecutable "Direct link to setffmpegexecutable")

_removed in v4.0_

Allows you to use a custom FFmpeg binary. Must be an absolute path. By default, this is null and the FFmpeg in `PATH` will be used.

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setFfmpegExecutable('/path/to/custom/ffmpeg');
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setFfmpegExecutable('/path/to/custom/ffmpeg');
```

The [command line flag](/docs/cli/render#--ffmpeg-executable) `--ffmpeg-executable` will take precedence over this option.

## `setFfprobeExecutable()` [​](\#setffprobeexecutable "Direct link to setffprobeexecutable")

_removed in v4.0_

Allows you to use a custom `ffprobe` binary. Must be an absolute path. By default, this is null and the `ffprobe` in `PATH` will be used.

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setFfprobeExecutable('/path/to/custom/ffprobe');
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setFfprobeExecutable('/path/to/custom/ffprobe');
```

The [command line flag](/docs/cli/render#--ffprobe-executable-) `--ffprobe-executable` will take precedence over this option.

## `setPort()` [​](\#setport "Direct link to setport")

_deprecated in v4.0.61 - use [`setStudioPort()`](/docs/config#setstudioport)_
and [`setRendererPort()`](/docs/config#setrendererport) instead.

Define on which port Remotion should start it's HTTP servers.

By default, Remotion will try to find a free port.

If you specify a port, but it's not available, Remotion will throw an error.

warning

Setting this option will break rendering in the Remotion Studio, because this option controls two settings at the same time:

- When starting the [Remotion Studio](/docs/terminology/studio), a server will be started to host it ( [`setStudioPort()`](/docs/config#setstudioport)).
- During rendering, a HTTP server is also started in the background to serve the Webpack [bundle](/docs/terminology/bundle) ( [`setRendererPort()`](/docs/config#setrendererport)).

Use the options individually.

```

remotion.config.ts
ts

Config.setPort(3003);
```

```

remotion.config.ts
ts

Config.setPort(3003);
```

The [command line flag](/docs/cli/render#--port) `--port` will take precedence over this option. If set on `npx remotion studio`, it will set the Studio port, otherwise the renderer port.

## `setOutputFormat()` [​](\#setoutputformat "Direct link to setoutputformat")

_Removed in v4.0.0_ _Deprecated_. Use `setCodec()` and `setImageSequence()` instead.

Either `'mp4'` or `'png-sequence'`.

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setOutputFormat('mp4');
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setOutputFormat('mp4');
```

The [command line flags](/docs/cli) `--sequence` and `--codec` will take precedence over this option.

The [command line flag](/docs/cli) `--quality` will take precedence over this option.

## `setImageFormat()` [v1.4.0](https://github.com/remotion-dev/remotion/releases/v1.4.0) [​](\#setimageformat "Direct link to setimageformat")

_Removed in v4.0_

Replaced in v4.0 with `setVideoImageFormat()` and `setStillImageFormat()`

Determines which in which image format to render the frames. Either:

- `jpeg` \- the fastest option (default from v1.1)
- `png` \- slower, but supports transparency
- `none` \- don't render images, just calculate audio information (available from v2.0)

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setImageFormat('png');
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';
// ---cut---
Config.setImageFormat('png');
```

The [command line flag](/docs/cli/render#--image-format) `--image-format` will take precedence over this option.

## Importing ES Modules [​](\#importing-es-modules "Direct link to Importing ES Modules")

The [config file](/docs/config) gets executed in a CommonJS environment. If you want to import ES modules to override the Webpack config, you can pass an async function to `Config.overrideWebpackConfig()`:

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig(async (currentConfiguration) => {
  const {enableSass} = await import('./src/enable-sass');
  return enableSass(currentConfiguration);
});
```

```

remotion.config.ts
ts

import {Config} from '@remotion/cli/config';

Config.overrideWebpackConfig(async (currentConfiguration) => {
  const {enableSass} = await import('./src/enable-sass');
  return enableSass(currentConfiguration);
});
```

## Old config file format [​](\#old-config-file-format "Direct link to Old config file format")

In v3.3.39, a new config file format was introduced which flattens the options so they can more easily be discovered using TypeScript autocompletion.

Previously, each config option was two levels deep:

```

remotion.config.ts
ts

Config.Bundling.setCachingEnabled(false);
```

```

remotion.config.ts
ts

Config.Bundling.setCachingEnabled(false);
```

From v3.3.39 on, all options can be accessed directly from the `Config` object.

```

remotion.config.ts
ts

Config.setCachingEnabled(false);
```

```

remotion.config.ts
ts

Config.setCachingEnabled(false);
```

The old way is deprecated, but will work for the foreseeable future.

## See also [​](\#see-also "Direct link to See also")

- [Encoding guide](/docs/encoding)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/config.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
help](/docs/cli/help) [Next\
\
remotion](/docs/remotion)

- [`overrideWebpackConfig()`](#overridewebpackconfig)
- [`setCachingEnabled()`](#setcachingenabled)
- [`setStudioPort()`](#setstudioport)
- [`setRendererPort()`](#setrendererport)
- [`setPublicDir()`](#setpublicdir)
- [`setEntryPoint()`](#setentrypoint)
- [`setLevel()`](#setlevel)
- [`setMaxTimelineTracks()`](#setmaxtimelinetracks)
- [`setKeyboardShortcutsEnabled()`](#setkeyboardshortcutsenabled)
- [`setWebpackPollingInMilliseconds()`](#setwebpackpollinginmilliseconds)
- [`setNumberOfSharedAudioTags()`](#setnumberofsharedaudiotags)
- [`setShouldOpenBrowser()`](#setshouldopenbrowser)
- [`setBrowserExecutable()`](#setbrowserexecutable)
- [`setDelayRenderTimeoutInMilliseconds()`](#setdelayrendertimeoutinmilliseconds)
- [`setChromiumDisableWebSecurity()`](#setchromiumdisablewebsecurity)
- [`setChromiumIgnoreCertificateErrors()`](#setchromiumignorecertificateerrors)
- [`setChromiumHeadlessMode()`](#setchromiumheadlessmode)
- [`setChromiumMultiProcessOnLinux()`](#setchromiummultiprocessonlinux)
- [`setChromiumOpenGlRenderer()`](#setchromiumopenglrenderer)
- [`setConcurrency()`](#setconcurrency)
- [`setVideoImageFormat()`](#setvideoimageformat)
- [`setStillImageFormat()`](#setstillimageformat)
- [`setScale()`](#setscale)
- [`setMuted()`](#setmuted)
- [`setEnforceAudioTrack()`](#setenforceaudiotrack)
- [`setForSeamlessAacConcatenation()`](#setforseamlessaacconcatenation)
- [`setFrameRange()`](#setframerange)
- [`setJpegQuality()`](#setjpegquality)
- [`setDotEnvLocation()`](#setdotenvlocation)
- [`setEveryNthFrame()`](#seteverynthframe)
- [`setNumberOfGifLoops()`](#setnumberofgifloops)
- [`setOutputLocation()`](#setoutputlocation)
- [`setOverwriteOutput()`](#setoverwriteoutput)
- [`setPixelFormat()`](#setpixelformat)
- [`setCodec()`](#setcodec)
- [`setAudioCodec()`](#setaudiocodec)
- [`setProResProfile()`](#setproresprofile)
- [`setX264Preset()`](#setx264preset)
- [`setImageSequence()`](#setimagesequence)
- [`overrideHeight()`](#overrideheight)
- [`overrideWidth()`](#overridewidth)
- [`setCrf()`](#setcrf)
- [`setVideoBitrate()`](#setvideobitrate)
- [`setEncodingBufferSize()`](#setencodingbuffersize)
- [`setEncodingMaxRate()`](#setencodingmaxrate)
- [`setAudioBitrate()`](#setaudiobitrate)
- [`setEnableFolderExpiry()`](#setenablefolderexpiry)
- [`setLambdaInsights()`](#setlambdainsights)
- [`setDeleteAfter()`](#setdeleteafter)
- [`setBeepOnFinish()`](#setbeeponfinish)
- [`setBufferStateDelayInMilliseconds()`](#setbufferstatedelayinmilliseconds)
- [`setBinariesDirectory()`](#setbinariesdirectory)
- [`setPreferLosslessAudio()`](#setpreferlosslessaudio)
- [`setHardwareAcceleration()`](#sethardwareacceleration)
- [`overrideFfmpegCommand()`](#overrideffmpegcommand)
- [`setQuality()`](#setquality)
- [`setFfmpegExecutable()`](#setffmpegexecutable)
- [`setFfprobeExecutable()`](#setffprobeexecutable)
- [`setPort()`](#setport)
- [`setOutputFormat()`](#setoutputformat)
- [`setImageFormat()`](#setimageformat)
- [Importing ES Modules](#importing-es-modules)
- [Old config file format](#old-config-file-format)
- [See also](#see-also)
