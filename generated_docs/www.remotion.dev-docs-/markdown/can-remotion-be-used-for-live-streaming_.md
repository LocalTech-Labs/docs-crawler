---
title: Can Remotion be used for live streaming?
source: Remotion Documentation
last_updated: 2024-12-22
---

# Can Remotion be used for live streaming?

- [Home page](/)
- FAQ
- Live streaming

# Can Remotion be used for live streaming?

Remotion is not designed for broadcasting. It's use case is to render fixed-length videos which content is deterministic from start to end.

Live streaming requires a different approach. Content is not known in advance, being able to react to events is often necessary, the video is not fixed-length and broadcasting is extremely performance sensitive, with all frames needing to be rendered in real-time.

Remotion declarative approach based on browser APIs is not suitable for this use case.

For streamers wanting to leverage web technologies in their livestream, we recommend that they build a web app with transparent background and add it as a browser source to [OBS](https://obsproject.com/).

You may create stinger transitions for OBS with Remotion, see an [example](https://github.com/JonnyBurger/ledevevent-stinger) here.

If you want to create live streams programmatically with React, we recommend [Live Compositor](https://compositor.live/) from our friends at [Software Mansion](https://swmansion.com/).

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/live-streaming.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Absolute paths](/docs/miscellaneous/absolute-paths) [Next\
\
parseMedia() vs. getVideoMetadata()](/docs/miscellaneous/parse-media-vs-get-video-metadata)
