---
title: Cannot read properties of undefined (reading 'decode')
source: Remotion Documentation
last_updated: 2024-12-22
---

# Cannot read properties of undefined (reading 'decode')

- [Home page](/)
- Recorder
- Troubleshooting
- Cannot read properties of undefined (reading 'decode')

# Cannot read properties of undefined (reading 'decode')

If you get the following error while starting the Studio:

```

Cannot read properties of undefined (reading 'decode')
  at new URLStateMachine
  at module.exports.basic
```

```

Cannot read properties of undefined (reading 'decode')
  at new URLStateMachine
  at module.exports.basic
```

You are experiencing an issue with Bun after upgrading Remotion.

You can resolve the issue by running

```

rm -rf node_modules
bun i
```

```

rm -rf node_modules
bun i
```

We are hoping to fix this issue soon.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/troubleshooting/cannot-read-properties-of-undefined.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Failed to execute getVideoMetadata()](/docs/recorder/troubleshooting/failed-to-execute-get-video-metadata) [Next\
\
Roadmap](/docs/recorder/roadmap)
