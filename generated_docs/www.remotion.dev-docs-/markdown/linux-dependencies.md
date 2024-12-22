---
title: Linux Dependencies
source: Remotion Documentation
last_updated: 2024-12-22
---

# Linux Dependencies

- [Home page](/)
- Miscellaneous
- Linux Dependencies

On this page

# Linux Dependencies

If you are on Linux, Chrome Headless Shell requires some shared libraries to be installed.

## Ubuntu 24.04 [​](\#ubuntu-2404 "Direct link to Ubuntu 24.04")

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libasound2t64 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libgbm-dev \
  libatk-bridge2.0-0
```

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libasound2t64 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libgbm-dev \
  libatk-bridge2.0-0
```

## Older versions of Ubuntu [​](\#older-versions-of-ubuntu "Direct link to Older versions of Ubuntu")

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libgbm-dev \
  libatk-bridge2.0-0
```

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libgbm-dev \
  libatk-bridge2.0-0
```

## Debian [​](\#debian "Direct link to Debian")

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2 \
  libatk-bridge2.0-0
```

```

bash

RUN apt install -y \
  libnss3 \
  libdbus-1-3 \
  libatk1.0-0 \
  libgbm-dev \
  libasound2 \
  libxrandr2 \
  libxkbcommon-dev \
  libxfixes3 \
  libxcomposite1 \
  libxdamage1 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2 \
  libatk-bridge2.0-0
```

## Amazon Linux 2023 [​](\#amazon-linux-2023 "Direct link to Amazon Linux 2023")

```

bash

RUN yum install -y \
  mesa-libgbm	\
  libX11 \
  libXrandr\
  libdrm \
  libXdamage \
  libXfixes \
  dbus -libs\
  libXcomposite \
  alsa -lib \
  nss \
  dbus \
  at-spi2-core \
  atk \
  at-spi2-atk
```

```

bash

RUN yum install -y \
  mesa-libgbm	\
  libX11 \
  libXrandr\
  libdrm \
  libXdamage \
  libXfixes \
  dbus -libs\
  libXcomposite \
  alsa -lib \
  nss \
  dbus \
  at-spi2-core \
  atk \
  at-spi2-atk
```

## Alpine Linux [​](\#alpine-linux "Direct link to Alpine Linux")

Not supported due to to unsupported Libc symbols.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/linux-dependencies.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell) [Next\
\
--gl flag](/docs/gl-options)

- [Ubuntu 24.04](#ubuntu-2404)
- [Older versions of Ubuntu](#older-versions-of-ubuntu)
- [Debian](#debian)
- [Amazon Linux 2023](#amazon-linux-2023)
- [Alpine Linux](#alpine-linux)
