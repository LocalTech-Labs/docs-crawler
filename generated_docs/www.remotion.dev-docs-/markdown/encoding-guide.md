---
title: Encoding Guide
source: Remotion Documentation
last_updated: 2024-12-22
---

# Encoding Guide

- [Home page](/)
- [Rendering](/docs/render)
- Encoding Guide

On this page

# Encoding Guide

Backed by [FFmpeg](https://ffmpeg.org/), Remotion allows you to configure a variety of encoding settings. The goal of this page is to help you navigate through the settings and to help you choose the right one.

## Choosing a codec [​](\#choosing-a-codec "Direct link to Choosing a codec")

Remotion supports 5 video codecs: `h264` ( _default_), `h265`, `vp8`, `vp9` and `prores`. While H264 will work well in most cases, sometimes it's worth going for a different codec. Refer to the table below to see the advantages and drawbacks of each codec.

CodecFile extensionFile sizeEncoding timeBrowser compatibilityHardware acceleration possible

H.264 also known as MPEG-4

.mp4, .mov or .mkvMediumVery fast[Very good](https://caniuse.com/mpeg4)No

H.265 also known as HEVC

.mp4 or .hevcMediumFast[Very poor](https://caniuse.com/hevc)NoVP8.webmSmallSlow[Okay](https://caniuse.com/webm)NoVP9.webmVery smallVery slow[Okay](https://caniuse.com/webm)NoProRes.movLargeFastNoneOn macOS

info

Click on a browser compatibility link to see exactly which browsers are supported on caniuse.com.

You can set a config using [`Config.setCodec()` in the config file](/docs/config#setcodec) or the [`--codec`](/docs/cli) CLI flag.

[Hardware acceleration](/docs/hardware-acceleration) is available from Remotion 4.0.228.

## Controlling quality using the CRF setting [​](\#controlling-quality-using-the-crf-setting "Direct link to Controlling quality using the CRF setting")

_Applies only to `h264`, `h265`, `vp8` and `vp9`._

No matter which codec you end up using, there's always a tradeoff between file size and video quality. You can control it by setting the so called CRF (Constant Rate Factor). The **lower the number, the better the quality**, the higher the number, the smaller the file is – of course at the cost of quality.

Be cautious: Every codec has it's own range of acceptable values and a different default. So while `23` will look very good on a H264 video, it will look terrible on a WebM video. Use this chart to determine which CRF value to use:

CodecMinimum - Best qualityMaximum - Best compressionDefaultH26415118H26505123VP84639VP906328

You can [set a CRF in the config file using the `Config.setCrf()`](/docs/config#setcrf) function or use the [`--crf`](/docs/cli/render#--jpeg-quality) command line flag.

note

If you enable [hardware acceleration](/docs/hardware-acceleration), you cannot set a `crf`.

## Controlling quality using `--video-bitrate` and `--audio-bitrate` [​](\#controlling-quality-using---video-bitrate-and---audio-bitrate "Direct link to controlling-quality-using---video-bitrate-and---audio-bitrate")

Use the following options to set the video and audio bitrate:

- In the Studio: Set video and audio bitrate in the Render Dialog
- In the CLI: Use the [`--video-bitrate`](/docs/cli/render#--video-bitrate) and [`--audio-bitrate`](/docs/cli/render#--audio-bitrate) flags
- In SSR, Lambda and Cloud Run APIs: Use the [`videoBitrate`](/docs/renderer/render-media#videobitrate) and [`audioBitrate`](http://localhost:3000/docs/renderer/render-media#audiobitrate) options.

This option is incompatible with other quality options.

## Controlling quality using ProRes profile [​](\#controlling-quality-using-prores-profile "Direct link to Controlling quality using ProRes profile")

_Applies only to `prores` codec_.

For ProRes, there is no CRF option, but there are profiles which you can set using the [`--prores-profile` flag](/docs/cli/render#--prores-profile) or the [`setProResProfile`](/docs/config#setproresprofile) config file option.

ValueFFmpeg settingBitrate[Supports alpha channel](/docs/transparent-videos)`"proxy"`0~45MbpsNo`"light"`1~102MbpsNo

`"standard"` (default)

2~147MbpsNo`"hq"`3~220MbpsNo`"4444"`4~330MbpsYes`"4444-xq"`4~500MbpsYes

Higher bitrate means higher quality and higher file size.

## Audio-only export [​](\#audio-only-export "Direct link to Audio-only export")

You can pass `mp3`, `wav` or `aac` as a codec. If you do it, an audio file will be output in the corresponding format. Quality settings will be ignored.

## GIFs [​](\#gifs "Direct link to GIFs")

You can also [render your video as a GIF](/docs/render-as-gif).

## Audio codec [​](\#audio-codec "Direct link to Audio codec")

_available from v3.3.42_

Using the [`--audio-codec`](/docs/config#setaudiocodec) flag, you can set the format of the audio that is embedded in the video. Not all codec and audio codec combinations are supported and certain combinations require a certain file extension and container format.

The container format will be automatically derived based on the file extension.

Video codecDefaultSupported audio codecsPossible file extensions`h264`✅`aac``.mp4` _(default)_, `.mkv`, `.mov``pcm-16``.mkv` _(default)_, `.mov``mp3``.mp4` _(default)_, `.mkv`, `.mov``h264-ts``pcm-16``.ts` _(default)_✅`aac``.ts` _(default)_`aac`✅`aac``.aac` _(default)_, `.3gp`, `.m4a`, `.m4b`, `.mpg`, `.mpeg``pcm-16``.wav` _(default)_`h265`✅`aac``.mp4` _(default)_, `.mkv`, `.hevc``pcm-16``.mkv` _(default)_`mp3`✅`mp3``.mp3` _(default)_`pcm-16``.wav` _(default)_`prores` \*`aac``.mov` _(default)_, `.mkv`, `.mxf`✅`pcm-16``.mov` _(default)_, `.mkv`, `.mxf``vp8`✅`opus``.webm` _(default)_`pcm-16``.mkv` _(default)_`vp9`✅`opus``.webm` _(default)_`pcm-16``.mkv` _(default)_`wav`✅`pcm-16``.wav` _(default)_

GIFs don't support audio.

\\* Note: In versions before `v4.0.0` the default audio codec for `ProRes` was `aac`. Now it's `pcm-16`.

## File extensions [​](\#file-extensions "Direct link to File extensions")

Specifying a file extension when rendering media will determine the default codec. You may override the codec using `--codec` as long as the combination is supported in the table above.

File extensionDefault codec.3gp`aac`.aac`aac`.gif`gif`.hevc`h265`.m4a`aac`.m4b`aac`.mkv`h264-mkv`.mov`prores`.mp3`mp3`.mp4`h264`.mpeg`aac`.mpg`aac`.mxf`prores`.wav`wav`.webm`vp8`.ts`h264-ts`

## What other settings do you need? [​](\#what-other-settings-do-you-need "Direct link to What other settings do you need?")

Which of the dozens of options that FFmpeg supports would you like to see exposed in Remotion? Let us know by opening an [issue on our issue tracker!](https://github.com/remotion-dev/remotion/issues)

## See also [​](\#see-also "Direct link to See also")

- [CLI Options](/docs/cli)
- [Configuration file](/docs/config)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/encoding.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Render your video](/docs/render) [Next\
\
Still images](/docs/stills)

- [Choosing a codec](#choosing-a-codec)
- [Controlling quality using the CRF setting](#controlling-quality-using-the-crf-setting)
- [Controlling quality using `--video-bitrate` and `--audio-bitrate`](#controlling-quality-using---video-bitrate-and---audio-bitrate)
- [Controlling quality using ProRes profile](#controlling-quality-using-prores-profile)
- [Audio-only export](#audio-only-export)
- [GIFs](#gifs)
- [Audio codec](#audio-codec)
- [File extensions](#file-extensions)
- [What other settings do you need?](#what-other-settings-do-you-need)
- [See also](#see-also)
