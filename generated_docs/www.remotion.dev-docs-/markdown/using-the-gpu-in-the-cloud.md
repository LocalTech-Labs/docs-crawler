---
title: Using the GPU in the cloud
source: Remotion Documentation
last_updated: 2024-12-22
---

# Using the GPU in the cloud

- [Home page](/)
- [Server-side rendering](/docs/ssr)
- GPU in the cloud (bare)

On this page

# Using the GPU in the cloud

Here is an example workflow for how to use the GPU in the cloud to render your videos using EC2 instances.

## Launch an EC2 instance [​](\#launch-an-ec2-instance "Direct link to Launch an EC2 instance")

[1](#1) **Ensure you are allowed to use a GPU**

You might need to ask AWS for a limit increase for the number of GPUs you can use.

You can do this in the AWS console.

Go to "Service Quotas" -> "AWS Services" -> "Amazon Elastic Compute Cloud (Amazon EC2)" -> "Running On-Demand G and VT instances" -> "Request increase at account-level".

You may also click [here](https://us-east-1.console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-DB2E81BA) to go to the page directly for the `us-east-1` region.

[2](#2) **Launch an EC2 instance**

[Click here launch an EC2 instance on `us-east-1`](https://us-east-1.console.aws.amazon.com/ec2/home?region=us-east-1#LaunchInstances:).

Select "Browse more AMIs", search for `ami-053b0d53c279acc90`, select the "Community AMIs" tab and select the image with the right AMI ("ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20230516").

We recommend the `g4dn.xlarge` size - note that this instance costs $375 per month with the default configuration.

If you get a message "Subscribing to AMI is taking longer than expected", it is normal. You may need to wait a few minutes.

[3](#3) **Configure instance:**

Once you connected to the instance, run the following commands:

```

Upgrade the Linux Kernel to v6
bash

sudo bash -c "apt update && export DEBIAN_FRONTEND=noninteractive && export NEEDRESTART_MODE=a && apt upgrade -y && reboot"
```

```

Upgrade the Linux Kernel to v6
bash

sudo bash -c "apt update && export DEBIAN_FRONTEND=noninteractive && export NEEDRESTART_MODE=a && apt upgrade -y && reboot"
```

The instance will restart and therefore disconnect. Wait a few moments and then reconnect.

```

Install libvulkan
bash

sudo apt install -y build-essential libvulkan1
```

```

Install libvulkan
bash

sudo apt install -y build-essential libvulkan1
```

```

Install GPU drivers
bash

DRIVER_URL="https://us.download.nvidia.com/tesla/535.104.12/NVIDIA-Linux-x86_64-535.104.12.run"
DRIVER_NAME="NVIDIA-Linux-driver.run"
wget -O "$DRIVER_NAME" "$DRIVER_URL"
sudo sh "$DRIVER_NAME" --disable-nouveau --silent
rm "$DRIVER_NAME"
```

```

Install GPU drivers
bash

DRIVER_URL="https://us.download.nvidia.com/tesla/535.104.12/NVIDIA-Linux-x86_64-535.104.12.run"
DRIVER_NAME="NVIDIA-Linux-driver.run"
wget -O "$DRIVER_NAME" "$DRIVER_URL"
sudo sh "$DRIVER_NAME" --disable-nouveau --silent
rm "$DRIVER_NAME"
```

```

Configure startup service
bash

echo '[Unit]
Description=Run nvidia-smi at system startup

[Service]
ExecStart=/usr/bin/nvidia-smi
Type=oneshot
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target' | sudo tee /etc/systemd/system/nvidia-smi.service
sudo systemctl enable nvidia-smi.service
sudo systemctl start nvidia-smi.service
```

```

Configure startup service
bash

echo '[Unit]
Description=Run nvidia-smi at system startup

[Service]
ExecStart=/usr/bin/nvidia-smi
Type=oneshot
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target' | sudo tee /etc/systemd/system/nvidia-smi.service
sudo systemctl enable nvidia-smi.service
sudo systemctl start nvidia-smi.service
```

```

Install Chrome
bash

curl -fsSL https://dl.google.com/linux/linux_signing_key.pub | sudo gpg --dearmor -o /usr/share/keyrings/googlechrom-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/googlechrom-keyring.gpg] http://dl.google.com/linux/chrome/deb/ stable main" | sudo tee /etc/apt/sources.list.d/google-chrome.list
sudo apt update
sudo apt install -y google-chrome-stable curl gnupg ca-certificates
```

```

Install Chrome
bash

curl -fsSL https://dl.google.com/linux/linux_signing_key.pub | sudo gpg --dearmor -o /usr/share/keyrings/googlechrom-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/googlechrom-keyring.gpg] http://dl.google.com/linux/chrome/deb/ stable main" | sudo tee /etc/apt/sources.list.d/google-chrome.list
sudo apt update
sudo apt install -y google-chrome-stable curl gnupg ca-certificates
```

```

Install Node.js
bash

sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg

NODE_MAJOR=20
echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list

sudo apt-get update
sudo apt-get install nodejs -y
```

```

Install Node.js
bash

sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg

NODE_MAJOR=20
echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list

sudo apt-get update
sudo apt-get install nodejs -y
```

[4](#4) **Render a video with the GPU**

```

Clone a Remotion GPU demo
bash

git clone https://github.com/remotion-dev/gpu-scene
cd gpu-scene
npm i
npx remotion render --browser-executable=/usr/bin/google-chrome-stable --gl=angle-egl --enable-multiprocess-on-linux
```

```

Clone a Remotion GPU demo
bash

git clone https://github.com/remotion-dev/gpu-scene
cd gpu-scene
npm i
npx remotion render --browser-executable=/usr/bin/google-chrome-stable --gl=angle-egl --enable-multiprocess-on-linux
```

## See also [​](\#see-also "Direct link to See also")

- [Run this in a Docker container](/docs/miscellaneous/cloud-gpu-docker)
- [How To Enable Hardware Acceleration on Chrome, Chromium & Puppeteer on AWS in Headless mode](https://mirzabilal.com/how-to-enable-hardware-acceleration-on-chrome-chromium-puppeteer-on-aws-in-headless-mode), on which this document is based on.

CONTRIBUTORS

[![bilalmughal](https://github.com/bilalmughal.png)\
\
**bilalmughal** \
\
Research and writing](https://github.com/bilalmughal) [![UmungoBungo](https://github.com/UmungoBungo.png)\
\
**UmungoBungo** \
\
Research](https://github.com/UmungoBungo)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/miscellaneous/cloud-gpu.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Using the GPU](/docs/gpu) [Next\
\
GPU in the cloud (Docker)](/docs/miscellaneous/cloud-gpu-docker)

- [Launch an EC2 instance](#launch-an-ec2-instance)
- [See also](#see-also)
