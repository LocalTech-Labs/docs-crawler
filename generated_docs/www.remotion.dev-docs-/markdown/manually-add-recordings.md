---
title: Manually add recordings
source: Remotion Documentation
last_updated: 2024-12-22
---

# Manually add recordings

- [Home page](/)
- Recorder
- Record
- Manually add recordings

# Manually add recordings

You can also add recordings from other sources, like your phone.

Simply name the file `webcam` followed by any number, like `10`, followed by one of the supported file extensions: `.mp4`, `.webm`, `.mkv` or `.mov`.

The order your recordings appear in your video is based on the number suffix.

Here is a way you could structure your recordings:

```

txt

public
└── my-video
    ├── webcam10.mp4
    ├── display10.mp4
    ├── webcam20.mp4
    └── webcam30.mp4
```

```

txt

public
└── my-video
    ├── webcam10.mp4
    ├── display10.mp4
    ├── webcam20.mp4
    └── webcam30.mp4
```

If there is only a `webcam` recording with said number, the webcam will be fullscreen.

A `display` recording will only get picked up if there is also a `webcam` recording.

The webcam will be miniaturized in this case.

If you only have recorded a display, give it a `webcam` prefix nonetheless for it to get picked up.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/record/manually.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Using the interface](/docs/recorder/record/) [Next\
\
Delete recordings](/docs/recorder/record/delete)
