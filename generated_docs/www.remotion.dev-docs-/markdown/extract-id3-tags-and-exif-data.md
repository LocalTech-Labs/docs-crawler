---
title: Extract ID3 tags and EXIF data
source: Remotion Documentation
last_updated: 2024-12-22
---

# Extract ID3 tags and EXIF data

- [Home page](/)
- [Media Parser](/docs/media-parser/)
- Extract ID3 tags

On this page

# Extract ID3 tags and EXIF data

[`@remotion/media-parser`](/docs/media-parser) allows you to extract metadata that is embedded in video files, such as ID3 tags or EXIF data.

```

Extracting metadata
tsx

import {parseMedia} from '@remotion/media-parser';

const {metadata} = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    metadata: true,
  },
});

console.log(metadata);

/*
[
  {
    key: 'com.apple.quicktime.model',
    trackId: null,
    value: 'iPhone 15 Pro',
  },
  {
    key: 'encoder',
    trackId: null,
    value: 'Lavf57.19.100',
  }
]
*/
```

```

Extracting metadata
tsx

import {parseMedia} from '@remotion/media-parser';

const {metadata} = await parseMedia({
  src: 'https://example.com/my-video.mp4',
  fields: {
    metadata: true,
  },
});

console.log(metadata);

/*
[
  {
    key: 'com.apple.quicktime.model',
    trackId: null,
    value: 'iPhone 15 Pro',
  },
  {
    key: 'encoder',
    trackId: null,
    value: 'Lavf57.19.100',
  }
]
*/
```

Some metadata refers to a certain track, such as the 'encoder' field.

In this case, the `trackId` field will contain a non-null value. You can get a list of tracks by setting the [`tracks`](/docs/media-parser/parse-media#tracks) field to `true`.

## Common metadata fields [​](\#common-metadata-fields "Direct link to Common metadata fields")

Here is a non-exhaustive list of commonly seen metadata that you may want to parse and display.

Use the [`container`](/docs/media-parser/parse-media#container) field to get the container format of the video.

### ISO Base Media (.mp4, .mov) [​](\#iso-base-media-mp4-mov "Direct link to ISO Base Media (.mp4, .mov)")

FieldDescriptionExample value`artist`ID3 Artist Tag`"Blender Foundation 2008, Janus Bager Kristensen 2013"``album`ID3 Album Tag`"The Resistance"``composer`ID3 Composer Tag`"Sacha Goedegebure"``comment`ID3 Comment Tag`"Made with Remotion v4.0.234"``releaseDate`ID3 Release Date Tag`genre`ID3 Genre Tag`"Animation"``title`ID3 Title Tag`"Big Buck Bunny, Sunflower version"``writer`ID3 Writer Tag`"Sacha Goedegebure"``director`ID3 Director Tag`producer`ID3 Producer Tag`description`ID3 Description Tag`duration`ID3 Duration Tag`"00:00:10.000000000"``encoder`Software used to encode the video/stream`"Lavf60.16.100"``copyright`Copyright information`"Blender Foundation"``major_brand`The major brand of the file`"isom"``minor_version`The minor version of the file`"512"``compatible_brands`Container compatibility`"isomav01iso2mp41"``handler_name`Software used to mux the video`"GPAC ISO Video Handler"``com.apple.quicktime.camera.focal_length.35mm_equivalent`[35mm-equivalent focal length](https://en.wikipedia.org/wiki/35_mm_equivalent_focal_length)`23``com.apple.quicktime.camera.lens_model`Lens of Apple Device`"iPhone 15 Pro back camera 6.765mm f/1.78"``com.apple.quicktime.creationdate`Creation date of video`"2024-10-01T12:46:18+0200"``com.apple.quicktime.software`Operating System Version (e.g. iOS) of Device`"18.0"``com.apple.quicktime.model`Device used to capture video`"iPhone 15 Pro"``com.apple.quicktime.make`Maker of the device used to capture the video`"Apple"``com.apple.quicktime.live-photo.vitality-score`[An identifier that represents the vitality score of a Live Photo movie](https://developer.apple.com/documentation/avfoundation/avmetadataidentifier/quicktimemetadatalivephotovitalityscore)`1``com.apple.quicktime.live-photo.vitality-scoring-version`[An identifier that represents the version of the algorithm responsible for scoring the Live Photo movie for vitality](https://developer.apple.com/documentation/avfoundation/avmetadataidentifier/quicktimemetadatalivephotovitalityscoringversion)`4``com.apple.quicktime.content.identifier`[An identifier that represents the content identifier in QuickTime](https://developer.apple.com/documentation/avfoundation/avmetadataidentifier/quicktimemetadatacontentidentifier)`"2C1C7C94-E977-45D0-9B3F-9A9CA8EFB47D"``com.apple.quicktime.full-frame-rate-playback-intent`[An key that represents whether this movie should play at full frame rate](https://developer.apple.com/documentation/avfoundation/avmetadatakey/quicktimemetadatakeyfullframerateplaybackintent)`1``com.apple.quicktime.live-photo.auto`[An identifier that represents whether the live photo movie used auto mode](https://developer.apple.com/documentation/avfoundation/avmetadataidentifier/quicktimemetadataautolivephoto?language=objc)`1``com.apple.quicktime.location.accuracy.horizontal`[An identifier that represents the horizontal accuracy of the location data](https://developer.apple.com/documentation/avfoundation/avmetadataidentifier/quicktimemetadatalocationhorizontalaccuracyinmeters)`1``com.apple.quicktime.information`General information`"Captured with VisionCamera by mrousavy"`

### Matroska (.mkv, .webm) [​](\#matroska-mkv-webm "Direct link to Matroska (.mkv, .webm)")

FieldDescriptionExample value`duration`Duration of a video / stream`"00:00:00.333000000"``encoder`Software used to encode stream/Video`"Lavc60.31.102 libx264"``comment`Software used to encode stream/Video`"Made with Remotion 4.0.192"`

### RIFF (.avi) [​](\#riff-avi "Direct link to RIFF (.avi)")

FieldDescriptionExample value`encoder`Software used to encode stream/Video`"Lavf57.19.100"`

## Fields are not unique [​](\#fields-are-not-unique "Direct link to Fields are not unique")

Some fields may appear multiple times per video, such as `encoder`, `comment` or `duration`.

Oftentimes they are scoped to a certain track, which is why the `trackId` field is present in the metadata object.

## See also [​](\#see-also "Direct link to See also")

- [`parseMedia()`](/docs/media-parser/parse-media)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/media-parser/tags.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Fast and slow operations](/docs/media-parser/fast-and-slow) [Next\
\
Runtime support](/docs/media-parser/support)

- [Common metadata fields](#common-metadata-fields)
  - [ISO Base Media (.mp4, .mov)](#iso-base-media-mp4-mov)
  - [Matroska (.mkv, .webm)](#matroska-mkv-webm)
  - [RIFF (.avi)](#riff-avi)
- [Fields are not unique](#fields-are-not-unique)
- [See also](#see-also)