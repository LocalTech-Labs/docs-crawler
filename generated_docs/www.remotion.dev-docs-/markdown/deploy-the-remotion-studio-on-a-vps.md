---
title: Deploy the Remotion Studio on a VPS
source: Remotion Documentation
last_updated: 2024-12-22
---

# Deploy the Remotion Studio on a VPS

- [Home page](/)
- Studio
- Deploy to a VPS

On this page

# Deploy the Remotion Studio on a VPS

[![](https://i.ytimg.com/vi/430B9xSs06U/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCFJ4FoC-8enbWW4aU8uo9NLhEl6w)\
\
Also available as a 3min video\
\
**Deploy the Remotion Studio**](https://www.youtube.com/watch?v=430B9xSs06U)

_available from v4.0.46_

You can deploy the Remotion Studio to a long-running server in the cloud and make the render UI accessible to your team.

To do so, you need to:

- Install Node.js and Chrome
- Run `npx remotion studio` on the server
- Ensure port 3000 is available to the internet

The following examples have been tested with the [Hello World](/templates/hello-world) template initialized using `npx create-video@latest`.

## Dockerizing the Remotion Studio [​](\#dockerizing-the-remotion-studio "Direct link to Dockerizing the Remotion Studio")

```

Dockerfile
docker

FROM node:22-bookworm-slim

# Install Chrome dependencies
RUN apt-get update
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
  libatk-bridge2.0-0 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2

# Copy everything from your project to the Docker image. Adjust if needed.
COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src

# If you have a public folder:
COPY public ./public

# Install the right package manager and dependencies - see below for Yarn/PNPM
RUN npm i

# Install Chrome
RUN npx remotion browser ensure

CMD ["npx", "remotion studio"]
```

```

Dockerfile
docker

FROM node:22-bookworm-slim

# Install Chrome dependencies
RUN apt-get update
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
  libatk-bridge2.0-0 \
  libpango-1.0-0 \
  libcairo2 \
  libcups2

# Copy everything from your project to the Docker image. Adjust if needed.
COPY package.json package*.json yarn.lock* pnpm-lock.yaml* bun.lockb* tsconfig.json* remotion.config.* ./
COPY src ./src

# If you have a public folder:
COPY public ./public

# Install the right package manager and dependencies - see below for Yarn/PNPM
RUN npm i

# Install Chrome
RUN npx remotion browser ensure

CMD ["npx", "remotion studio"]
```

## Fly.io [​](\#flyio "Direct link to Fly.io")

To deploy the Remotion Studio to [Fly.io](https://fly.io):

- First add the above `Dockerfile` to the repo.
- Make the following change to the `Dockerfile` (works from v4.0.125):

```

diff

- CMD ["npx", "remotion studio"]
+ CMD ["npx", "remotion studio", "--ipv4"]
```

```

diff

- CMD ["npx", "remotion studio"]
+ CMD ["npx", "remotion studio", "--ipv4"]
```

- Ensure you are on a paid plan (free plan has too little memory)
- Use the following command:

```

bash

fly launch \
  # Get 2 CPU cores and 4GB of memory
  --vm-size=performance-2x \
  # Disable 2x replication
  --ha=false \
  # Use Remotion's port
  --internal-port=3000 \
  # Use Docker, not Node
  --dockerfile Dockerfile
```

```

bash

fly launch \
  # Get 2 CPU cores and 4GB of memory
  --vm-size=performance-2x \
  # Disable 2x replication
  --ha=false \
  # Use Remotion's port
  --internal-port=3000 \
  # Use Docker, not Node
  --dockerfile Dockerfile
```

For the following questions, answer no:

```

txt

? Would you like to set up a Postgresql database now? No
? Would you like to set up an Upstash Redis database now? No
? Create .dockerignore from 1 .gitignore files? No
```

```

txt

? Would you like to set up a Postgresql database now? No
? Would you like to set up an Upstash Redis database now? No
? Create .dockerignore from 1 .gitignore files? No
```

Answer Yes when asking if you want to deploy:

```

? Do you want to deploy now? Yes
```

```

? Do you want to deploy now? Yes
```

You should get a URL where the Studio was deployed!

## Render.com [​](\#rendercom "Direct link to Render.com")

To deploy the Remotion Studio to [Render.com](https://render.com):

- First add the above `Dockerfile` to the repo.
- Create a new "Web Service" and link your repository.
- Choose at least the "Standard" plan (2GB Memory).
- Deploy!

## DigitalOcean App Platform [​](\#digitalocean-app-platform "Direct link to DigitalOcean App Platform")

Is not working at the moment. The Render Button is disabled, because the DigitalOcean HTTP Proxy [does not support server-sent events](https://www.digitalocean.com/community/questions/does-app-platform-support-sse-server-sent-events-application).

A normal DigitalOcean droplet does work, though.

## Scaleway Serverless Container [​](\#scaleway-serverless-container "Direct link to Scaleway Serverless Container")

To deploy the [Scaleway Serverless Container](https://www.scaleway.com/en/serverless-containers/):

- First add the above `Dockerfile` to the repository
- Build your docker image and publish it to a Docker registry (for example a [Scaleway Docker Registry](https://www.scaleway.com/en/docs/containers/container-registry/quickstart/))
- Then [create your Serverless container namespace](https://www.scaleway.com/en/docs/serverless/containers/how-to/create-a-containers-namespace/) and use [your published docker image](https://www.scaleway.com/en/docs/serverless/containers/how-to/deploy-a-container-from-scaleway-container-registry/).

You can find [an example of a github action workflow that deploy Remotion Studio to Scaleway Serverless Container](https://github.com/lyonjs/shortvid.io/blob/main/.github/workflows/deploy-editor.yml).

## Example for deployment [​](\#example-for-deployment "Direct link to Example for deployment")

A test project to deploy the Remotion Studio is available [here](https://github.com/remotion-dev/shorts-customizer).

## See also [​](\#see-also "Direct link to See also")

- [Dockerizing an Remotion app](/docs/docker)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/studio/deploy-server.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Deploy as static site](/docs/studio/deploy-static) [Next\
\
Overview](/docs/ssr)

- [Dockerizing the Remotion Studio](#dockerizing-the-remotion-studio)
- [Fly.io](#flyio)
- [Render.com](#rendercom)
- [DigitalOcean App Platform](#digitalocean-app-platform)
- [Scaleway Serverless Container](#scaleway-serverless-container)
- [Example for deployment](#example-for-deployment)
- [See also](#see-also)
