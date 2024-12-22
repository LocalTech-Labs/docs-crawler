---
title: Setup EC2 for Docker with GPU
source: Remotion Documentation
last_updated: 2024-12-22
---

# Setup EC2 for Docker with GPU

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- GPU in the cloud (Docker)

On this page

# Setup EC2 for Docker with GPU

Follow these steps closely to render videos on EC2 in a Docker container.

These steps are opinionated, but specify a reference that works.

A word of warning: Deviating from the instructions, like:

- choosing a different AMI
- choosing a different Docker base
- choosing something else than EC2
- choosing a different host machine

may lead to the GPU not working. In this case, it is hard to debug.

We recommend to first follow these instructions and make changes once you have a working setup.

## Setup EC2 for Docker with GPU [​](\#setup-ec2-for-docker-with-gpu "Direct link to Setup EC2 for Docker with GPU")

[1](#1)

Follow the instructions for [GPUs on EC2](/docs/miscellaneous/cloud-gpu)
. You can skip installing Chrome, Node.js and cloning the repo to render a video.

[2](#2)

Install NVIDIA Container toolkit:

```

Add keyring
bash

curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
sudo apt-get update
```

```

Add keyring
bash

curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
sudo apt-get update
```

```

Install toolkit
bash

sudo apt-get install -y nvidia-container-toolkit
```

```

Install toolkit
bash

sudo apt-get install -y nvidia-container-toolkit
```

[3](#3)

Install Docker:

```

Add Docker's official GPG key
bash

sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

```

Add Docker's official GPG key
bash

sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

```

Add keyring
bash

echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

```

Add keyring
bash

echo "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

```

Install Docker
bash

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

```

Install Docker
bash

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

[4](#4)

Configure Docker to use the NVIDIA runtime

```

Configure the NVIDIA container runtime
bash

sudo nvidia-ctk runtime configure --runtime=docker
sudo systemctl restart docker
```

```

Configure the NVIDIA container runtime
bash

sudo nvidia-ctk runtime configure --runtime=docker
sudo systemctl restart docker
```

[5](#5)

Create two files, `Dockerfile` and `entrypoint.sh`
. You can for example create them using the `nano ./file-to-create` command.
Use `Ctrl`  `X` to save and quit.

```

Dockerfile
bash

FROM node:20-bookworm
RUN apt-get update
RUN apt-get install -y curl gnupg git
RUN rm -rf /var/lib/apt/lists/*

# Clone the repo
RUN git clone https://github.com/remotion-dev/gpu-scene.git
WORKDIR /gpu-scene
RUN npm install

# Copy the entrypoint script into the image
COPY entrypoint.sh .

CMD ["sh", "./entrypoint.sh"]
```

```

Dockerfile
bash

FROM node:20-bookworm
RUN apt-get update
RUN apt-get install -y curl gnupg git
RUN rm -rf /var/lib/apt/lists/*

# Clone the repo
RUN git clone https://github.com/remotion-dev/gpu-scene.git
WORKDIR /gpu-scene
RUN npm install

# Copy the entrypoint script into the image
COPY entrypoint.sh .

CMD ["sh", "./entrypoint.sh"]
```

```

entrypoint.sh
bash

#!/bin/bash

npx remotion render --gl=angle-egl Scene out/video.mp4
```

```

entrypoint.sh
bash

#!/bin/bash

npx remotion render --gl=angle-egl Scene out/video.mp4
```

[6](#6)

Build the container and run a sample render:

```

bash

sudo docker build . -t remotion-docker-gpu
sudo docker run --gpus all --runtime=nvidia -e "NVIDIA_DRIVER_CAPABILITIES=all" remotion-docker-gpu
```

```

bash

sudo docker build . -t remotion-docker-gpu
sudo docker run --gpus all --runtime=nvidia -e "NVIDIA_DRIVER_CAPABILITIES=all" remotion-docker-gpu
```

## Debugging [​](\#debugging "Direct link to Debugging")

Use the [`npx remotion gpu`](/docs/cli/gpu) command to get the output of `chrome://gpu` to verify that the GPU is working.

## See also [​](\#see-also "Direct link to See also")

- [GPUs on EC2](/docs/miscellaneous/cloud-gpu)
- [Using the GPU](/docs/gpu)

CONTRIBUTORS

[![UmungoBungo](https://github.com/UmungoBungo.png)\
\
**UmungoBungo** \
\
Workflow author](https://github.com/UmungoBungo) [![kaf-lamed-beyt](https://github.com/kaf-lamed-beyt.png)\
\
**kaf-lamed-beyt** \
\
Writeup](https://github.com/kaf-lamed-beyt)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/cloud-gpu-docker.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
GPU in the cloud (bare)](/docs/miscellaneous/cloud-gpu) [Next\
\
Comparison of SSR options](/docs/compare-ssr)

- [Setup EC2 for Docker with GPU](#setup-ec2-for-docker-with-gpu)
- [Debugging](#debugging)
- [See also](#see-also)
