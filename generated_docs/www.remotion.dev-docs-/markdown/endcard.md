---
title: Endcard
source: Remotion Documentation
last_updated: 2024-12-22
---

# Endcard

- [Home page](/)
- Recorder
- Editing
- Endcard

# Endcard

In the endcard you can add a call to follow, as well as links to your social media profiles or other websites.
This can be setup in the [`config/endcard.ts`](https://github.com/remotion-dev/recorder/tree/main/config/endcard.ts) file by adding your information as `Channel` object.

```

tsx

export const channels: {
  [key in Channel]: ChannelConfig & {
    isLinkedInBusinessPage: boolean;
  };
} = {
  jonny: {
    instagram: null,
    linkedin: "Jonny Burger",
    x: "@JNYBGR",
    youtube: "/JonnyBurger",
    discord: null,
    isLinkedInBusinessPage: false,
  },
  remotion: {
    instagram: "@remotion",
    linkedin: "Remotion",
    x: "@remotion",
    youtube: "@remotion_dev",
    discord: null,
    isLinkedInBusinessPage: true,
  },
};
```

```

tsx

export const channels: {
  [key in Channel]: ChannelConfig & {
    isLinkedInBusinessPage: boolean;
  };
} = {
  jonny: {
    instagram: null,
    linkedin: "Jonny Burger",
    x: "@JNYBGR",
    youtube: "/JonnyBurger",
    discord: null,
    isLinkedInBusinessPage: false,
  },
  remotion: {
    instagram: "@remotion",
    linkedin: "Remotion",
    x: "@remotion",
    youtube: "@remotion_dev",
    discord: null,
    isLinkedInBusinessPage: true,
  },
};
```

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/recorder/editing/endcard.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Scenes](/docs/recorder/editing/scenes) [Next\
\
Transitions](/docs/recorder/editing/transitions)
