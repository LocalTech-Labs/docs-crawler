---
title: Third party integrations
source: Remotion Documentation
last_updated: 2024-12-22
---

# Third party integrations

- [Home page](/)
- Tooling
- Third party integrations

On this page

# Third party integrations

All animations in Remotion must be driven by the value returned by the [`useCurrentFrame()`](/docs/use-current-frame) hook. If you would like to use another way of animations in Remotion, you need an integration that supports synchronizing the timing with Remotion.

On this page, we maintain a list of integrations for popular ways of animating on the web, and provide status for popular requests.

## After Effects [​](\#after-effects "Direct link to After Effects")

See: [Lottie - Import from After Effects](/docs/after-effects)

## Anime.JS [​](\#animejs "Direct link to Anime.JS")

See [this repository](https://github.com/remotion-dev/anime-example) for an example.

Read [this article](https://revidcraft.com/posts/remotion-part-02) how one developer integrated Anime.js with Remotion and boosted their efficiency.

## CSS animations [​](\#css-animations "Direct link to CSS animations")

You can synchronize animations with Remotions time using the [`animation-play-state`](https://developer.mozilla.org/en-US/docs/Web/CSS/animation-play-state) and [`animation-delay`](https://developer.mozilla.org/en-US/docs/Web/CSS/animation-delay) CSS properties.

See an [example here](https://stackblitz.com/~/github.com/remotion-dev/css-animation-play-state).

## Framer Motion [​](\#framer-motion "Direct link to Framer Motion")

At the moment, we don't have a Framer Motion integration, but are discussing the matter on [GitHub Issues](https://github.com/remotion-dev/remotion/issues/399).

## GIFs [​](\#gifs "Direct link to GIFs")

Use the [`@remotion/gif`](/docs/gif) package.

## GreenSock [​](\#greensock "Direct link to GreenSock")

See: [How to integrate GreenSock with Remotion](https://archive.ph/dGQ19)

## Lottie [​](\#lottie "Direct link to Lottie")

Use the [`@remotion/lottie`](/docs/lottie) package.

## Matter.js [​](\#matterjs "Direct link to Matter.js")

There is no integration available, but there are two examples of how you can "bake" the Matter.js into a timeline in order to synchronize it later with Remotion:

- [Baking GitHub discussion](https://github.com/orgs/remotion-dev/discussions/4373)
- [Pool game animation](https://github.com/hylarucoder/remotion-physics)

## React Native Skia [​](\#react-native-skia "Direct link to React Native Skia")

Use the [`@remotion/skia`](/docs/skia) package.

## react-spring [​](\#react-spring "Direct link to react-spring")

There is no direct compatibility but Remotion provides it's own [`spring()`](/docs/spring) instead.

## Reanimated [​](\#reanimated "Direct link to Reanimated")

There is no integration available but Remotion shares some code with Reanimated, in particular [`interpolate()`](/docs/interpolate), [`spring()`](/docs/spring) and [`Easing`](/docs/easing). This makes it easier to refactor already existing animation from Reanimated.

## Rive [​](\#rive "Direct link to Rive")

Use the [`@remotion/rive`](/docs/rive) package.

## TailwindCSS [​](\#tailwindcss "Direct link to TailwindCSS")

See: [TailwindCSS](/docs/tailwind)

## Three.JS [​](\#threejs "Direct link to Three.JS")

Use the [`@remotion/three`](/docs/three) package.

## Vidstack [​](\#vidstack "Direct link to Vidstack")

Use the [Remotion Provider for Vidstack](https://www.vidstack.io/docs/player/api/providers/remotion).

## Other libraries [​](\#other-libraries "Direct link to Other libraries")

Are you interested in using other libraries with Remotion? You can [file a GitHub issue](https://github.com/remotion-dev/remotion/issues/new) to inquire it. While we cannot guarantee to help, you can register interest and kick off the discussion.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/third-party.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Plain JavaScript](/docs/javascript) [Next\
\
TypeScript aliases](/docs/typescript-aliases)

- [After Effects](#after-effects)
- [Anime.JS](#animejs)
- [CSS animations](#css-animations)
- [Framer Motion](#framer-motion)
- [GIFs](#gifs)
- [GreenSock](#greensock)
- [Lottie](#lottie)
- [Matter.js](#matterjs)
- [React Native Skia](#react-native-skia)
- [react-spring](#react-spring)
- [Reanimated](#reanimated)
- [Rive](#rive)
- [TailwindCSS](#tailwindcss)
- [Three.JS](#threejs)
- [Vidstack](#vidstack)
- [Other libraries](#other-libraries)
