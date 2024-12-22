---
title: Failed to execute getVideoMetadata()
source: Remotion Documentation
last_updated: 2024-12-22
---

# Failed to execute getVideoMetadata()

- [Home page](/)
- Recorder
- Troubleshooting
- Failed to execute getVideoMetadata()

# Failed to execute getVideoMetadata()

If you get the following error while loading a composition:

```

txt

Failed to execute getVideoMetadata(), Received a MediaError loading "/static-41edfae54cb6/welcome/webcam122.mp4". HTTP Status code of the file: 200. Check the path of the file and if it is a valid video.
```

```

txt

Failed to execute getVideoMetadata(), Received a MediaError loading "/static-41edfae54cb6/welcome/webcam122.mp4". HTTP Status code of the file: 200. Check the path of the file and if it is a valid video.
```

It could be that the video file is still being re-encoded and FFmpeg is in the process of writing it.

Wait until the process is finished and hard refresh the page to clear the cache (Cmd+Shift+R on Mac and Ctrl+Shift+R on Windows).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/troubleshooting/failed-to-execute-get-video-metadata.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Upgrading](/docs/recorder/upgrading) [Next\
\
Cannot read properties of undefined (reading 'decode')](/docs/recorder/troubleshooting/cannot-read-properties-of-undefined)
