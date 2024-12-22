---
title: Emojis
source: Remotion Documentation
last_updated: 2024-12-22
---

# Emojis

- [Home page](/)
- Miscellaneous
- Emojis

On this page

# Emojis

Emojis are always rendered as-is by the operating system.

By nature, they are different across different OS:

- [Windows Emoji Set](https://emojipedia.org/microsoft)
- [macOS and iOS Emoji Set](https://emojipedia.org/apple)
- Different Linux distributions have different sets of emojis, or none at all.

## Emojis on Linux and in Docker [​](\#emojis-on-linux-and-in-docker "Direct link to Emojis on Linux and in Docker")

When rendering in the cloud or in Linux in general, you will not get the same emojis as on your Desktop.

You can install open source emoji fonts for your operating system:

### Debian/Ubuntu: [​](\#debianubuntu "Direct link to Debian/Ubuntu:")

```

docker

apt-get install fonts-noto-color-emoji
```

```

docker

apt-get install fonts-noto-color-emoji
```

### Amazon Linux 2023: [​](\#amazon-linux-2023 "Direct link to Amazon Linux 2023:")

```

docker

yum install google-noto-emoji-color-fonts
```

```

docker

yum install google-noto-emoji-color-fonts
```

## Emojis on Remotion Lambda [​](\#emojis-on-remotion-lambda "Direct link to Emojis on Remotion Lambda")

[Noto Color Emoji are installed on Remotion Lambda](/docs/lambda/runtime#fonts) by default.

You can [enable Apple Emojis](/docs/lambda/deployfunction#runtimepreference).

Replacing Emojis is possible by [customizing Remotion Lambda](/docs/lambda/custom-layers).

note

Apple Emojis are intellectual property of Apple Inc.

You are responsible for the use of Apple Emojis in your project.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/emojis.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Standalone functions](/docs/standalone) [Next\
\
#t= Media Fragments](/docs/media-fragments)

- [Emojis on Linux and in Docker](#emojis-on-linux-and-in-docker)
  - [Debian/Ubuntu:](#debianubuntu)
  - [Amazon Linux 2023:](#amazon-linux-2023)
- [Emojis on Remotion Lambda](#emojis-on-remotion-lambda)
