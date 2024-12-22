MiscellaneousLinux DependenciesOn this pageLinux DependenciesIf you are on Linux, Chrome Headless Shell requires some shared libraries to be installed.
Ubuntu 24.04​
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libasound2t64 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libgbm-dev \  libatk-bridge2.0-0
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libasound2t64 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libgbm-dev \  libatk-bridge2.0-0
Older versions of Ubuntu​
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libasound2 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libgbm-dev \  libatk-bridge2.0-0
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libasound2 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libgbm-dev \  libatk-bridge2.0-0
Debian​
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libgbm-dev \  libasound2 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libpango-1.0-0 \  libcairo2 \  libcups2 \  libatk-bridge2.0-0
bashRUN apt install -y \  libnss3 \  libdbus-1-3 \  libatk1.0-0 \  libgbm-dev \  libasound2 \  libxrandr2 \  libxkbcommon-dev \  libxfixes3 \  libxcomposite1 \  libxdamage1 \  libpango-1.0-0 \  libcairo2 \  libcups2 \  libatk-bridge2.0-0
Amazon Linux 2023​
bashRUN yum install -y \  mesa-libgbm	\  libX11 \  libXrandr\  libdrm \  libXdamage \  libXfixes \  dbus -libs\  libXcomposite \  alsa -lib \  nss \  dbus \  at-spi2-core \  atk \  at-spi2-atk
bashRUN yum install -y \  mesa-libgbm	\  libX11 \  libXrandr\  libdrm \  libXdamage \  libXfixes \  dbus -libs\  libXcomposite \  alsa -lib \  nss \  dbus \  at-spi2-core \  atk \  at-spi2-atk
Alpine Linux​
Not supported due to to unsupported Libc symbols.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousChrome Headless ShellNext--gl flagUbuntu 24.04Older versions of UbuntuDebianAmazon Linux 2023Alpine Linux