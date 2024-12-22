---
title: Setting video metadata
source: Remotion Documentation
last_updated: 2024-12-22
---

# Setting video metadata

- [Home page](/)
- [Rendering](/docs/render)
- Setting video metadata

On this page

# Setting video metadata

Video files can store metadata in them.

This metadata can be viewed using the `npx remotion ffprobe` command:

```

Getting metadata
txt

$ npx remotion ffprobe bigbuckbunny.mp4
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'bigbuckbunny.mp4':
  Metadata:
    major_brand     : mp42
    minor_version   : 0
    compatible_brands: isomavc1mp42
    creation_time   : 2010-01-10T08:29:06.000000Z
    comment         : This is a comment
    artist          : Remotion
  Duration: 00:09:56.47, start: 0.000000, bitrate: 2119 kb/s
  Stream #0:0[0x1](und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 125 kb/s (default)
      Metadata:
        creation_time   : 2010-01-10T08:29:06.000000Z
        handler_name    : (C) 2007 Google Inc. v08.13.2007.
        vendor_id       : [0][0][0][0]
  Stream #0:1[0x2](und): Video: h264 (High) (avc1 / 0x31637661), yuv420p(progressive), 1280x720 [SAR 1:1 DAR 16:9], 1991 kb/s, 24 fps, 24 tbr, 24k tbn (default)
      Metadata:
        creation_time   : 2010-01-10T08:29:06.000000Z
        handler_name    : (C) 2007 Google Inc. v08.13.2007.
        vendor_id       : [0][0][0][0]
```

```

Getting metadata
txt

$ npx remotion ffprobe bigbuckbunny.mp4
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'bigbuckbunny.mp4':
  Metadata:
    major_brand     : mp42
    minor_version   : 0
    compatible_brands: isomavc1mp42
    creation_time   : 2010-01-10T08:29:06.000000Z
    comment         : This is a comment
    artist          : Remotion
  Duration: 00:09:56.47, start: 0.000000, bitrate: 2119 kb/s
  Stream #0:0[0x1](und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 125 kb/s (default)
      Metadata:
        creation_time   : 2010-01-10T08:29:06.000000Z
        handler_name    : (C) 2007 Google Inc. v08.13.2007.
        vendor_id       : [0][0][0][0]
  Stream #0:1[0x2](und): Video: h264 (High) (avc1 / 0x31637661), yuv420p(progressive), 1280x720 [SAR 1:1 DAR 16:9], 1991 kb/s, 24 fps, 24 tbr, 24k tbn (default)
      Metadata:
        creation_time   : 2010-01-10T08:29:06.000000Z
        handler_name    : (C) 2007 Google Inc. v08.13.2007.
        vendor_id       : [0][0][0][0]
```

## Adding metadata [v4.0.216](https://github.com/remotion-dev/remotion/releases/v4.0.216) [​](\#adding-metadata "Direct link to adding-metadata")

You can set metadata when rendering videos with Remotion using

- [`--metadata`](/docs/cli/render#--metadata) in `npx remotion render`
- [`--metadata`](/docs/lambda/cli/render#--metadata) in `npx remotion lambda render`
- [`--metadata`](/docs/cloudrun/cli/render#--metadata) in `npx remotion cloudrun render`
- [`metadata`](/docs/renderer/render-media#metadata) in `renderMedia()`
- [`metadata`](/docs/lambda/rendermediaonlambda) in `renderMediaOnLambda()`
- [`metadata`](/docs/cloudrun/rendermediaoncloudrun) in `renderMediaOnCloudrun()`

## Accepted metadata [​](\#accepted-metadata "Direct link to Accepted metadata")

ISO Base Media Format videos ( `.mp4`, `.mov`) only accepts the following metadata:

- `title`
- `artist`
- `album_artist`
- `composer`
- `album`
- `date`
- `encoding_tool` (Already set by FFmpeg)
- `comment` (Already set by Remotion)
- `genre`
- `copyright`
- `grouping`
- `lyrics`
- `description`
- `synopsis`
- `show`
- `episode_id`
- `network`
- `keywords`
- `episode_sort` (int8)
- `season_number` (int8)
- `media_type` (int8)
- `hd_video` (int8)
- `gapless_playback` (int8)
- `compilation` (int8)

Fields designated with "int8" expect a numeric value between 0 and 255.

Matroska videos ( `.webm`, `.mkv`) accept arbitrary key-value fields.

The keys are always case-insensitive.

## Predefined metadata [​](\#predefined-metadata "Direct link to Predefined metadata")

Remotion already sets the `comment` field to `Made with Remotion v[VERSION]`.

If you set a custom comment, it will be merged with Remotion's comment.

FFmpeg also already sets the `encoder` field to `Lavc[VERSION]` (stands for "libavcodec", FFmpeg's video codec library).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Emitting Artifacts](/docs/artifacts) [Next\
\
Hardware acceleration](/docs/hardware-acceleration)

- [Adding metadata](#adding-metadata)
- [Accepted metadata](#accepted-metadata)
- [Predefined metadata](#predefined-metadata)