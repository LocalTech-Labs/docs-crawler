---
title: Upgrading
source: Remotion Documentation
last_updated: 2024-12-22
---

# Upgrading

- [Home page](/)
- Recorder
- Upgrading

On this page

# Upgrading

Since the Recorder ships as source code and allows you to hack around with it, your version will diverge from our template, at least once we release an update.

## Pull updates [​](\#pull-updates "Direct link to Pull updates")

You may run

```

bun update.ts
```

```

bun update.ts
```

to copy all files from our template into your version of the Recorder, except the `public`, `config` and `remotion/Root.tsx` folders.
This upgrades your local Remotion Recorder to the latest version.

If necessary, compare the changes that we have made the Recorder by visiting the following URL:

## Comparing changes [​](\#comparing-changes "Direct link to Comparing changes")

```

https://github.com/remotion-dev/recorder/compare/COMMIT_ID...main
```

```

https://github.com/remotion-dev/recorder/compare/COMMIT_ID...main
```

Where `COMMIT_ID` is the commit at which you forked the Recorder.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/upgrading.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
External recordings](/docs/recorder/external-recordings) [Next\
\
Failed to execute getVideoMetadata()](/docs/recorder/troubleshooting/failed-to-execute-get-video-metadata)

- [Pull updates](#pull-updates)
- [Comparing changes](#comparing-changes)
