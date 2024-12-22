---
title: Wrongly mounted <Composition>
source: Remotion Documentation
last_updated: 2024-12-22
---

# Wrongly mounted <Composition>

- [Home page](/)
- Troubleshooting
- Wrongly mounted <Composition>

On this page

# Wrongly mounted <Composition>

## Problem [​](\#problem "Direct link to Problem")

You are facing the following error:

```

<Composition> mounted inside another composition.
```

```

<Composition> mounted inside another composition.
```

or, inside a player:

```

'<Composition> was mounted inside the `component` that was passed to the <Player>.'
```

```

'<Composition> was mounted inside the `component` that was passed to the <Player>.'
```

## Solution [​](\#solution "Direct link to Solution")

### In the Remotion Studio [​](\#in-the-remotion-studio "Direct link to In the Remotion Studio")

The cause for the error is that a [`<Composition>`](/docs/composition) was nested inside the `component` that was passed to another `<Composition>`.

```

❌
tsx

const MyComp: React.FC = () => {
  return (
    <Composition
      id="another-comp"
      width={1080}
      height={1080}
      durationInFrames={30}
      fps={30}
      component={AnotherComp}
    />
  );
};

const Index: React.FC = () => {
  return (
    <Composition
      id="my-comp"
      width={1080}
      height={1080}
      durationInFrames={30}
      fps={30}
      component={MyComp}
    />
  );
};
```

```

❌
tsx

const MyComp: React.FC = () => {
  return (
    <Composition
      id="another-comp"
      width={1080}
      height={1080}
      durationInFrames={30}
      fps={30}
      component={AnotherComp}
    />
  );
};

const Index: React.FC = () => {
  return (
    <Composition
      id="my-comp"
      width={1080}
      height={1080}
      durationInFrames={30}
      fps={30}
      component={MyComp}
    />
  );
};
```

note

There is no reason to nest compositions in Remotion.

- Do you want to include a component in another? Mount it directly instead by writing: `<AnotherComp />`
- Do you want to limit the duration or time-shift another component? Look into [`<Sequence>`](/docs/sequence).

To fix it, you must remove the nested compositions.

### In the Remotion Player [​](\#in-the-remotion-player "Direct link to In the Remotion Player")

The cause for the error is that in the component you passed in to the `component` prop of Remotion Player, you are returning a `<Composition>`.

```

❌
tsx

const MyComp: React.FC = () => {
  return (
    <Composition
      durationInFrames={300}
      fps={30}
      width={1080}
      height={1080}
      id="another-component"
      component={AnotherComp}
    />
  );
};

const Index: React.FC = () => {
  return (
    <Player
      durationInFrames={300}
      fps={30}
      compositionWidth={1080}
      compositionHeight={1080}
      component={MyComp}
    />
  );
};
```

```

❌
tsx

const MyComp: React.FC = () => {
  return (
    <Composition
      durationInFrames={300}
      fps={30}
      width={1080}
      height={1080}
      id="another-component"
      component={AnotherComp}
    />
  );
};

const Index: React.FC = () => {
  return (
    <Player
      durationInFrames={300}
      fps={30}
      compositionWidth={1080}
      compositionHeight={1080}
      component={MyComp}
    />
  );
};
```

note

There is no use for compositions in `<Player>`. Only use them during rendering and when using the Remotion Studio.

To fix it, pass the component directly to the player's [`component`](/docs/player/player#component-or-lazycomponent) prop and provide the [`durationInFrames`](/docs/player/player#durationinframes), [`fps`](/docs/player/player#fps), [`compositionHeight`](/docs/player/player#compositionheight) and [`compositionWidth`](/docs/player/player#compositionwidth) props to the player.

```

✅
tsx

const Index: React.FC = () => {
  return (
    <Player
      durationInFrames={300}
      fps={30}
      compositionWidth={1080}
      compositionHeight={1080}
      component={AnotherComp}
    />
  );
};
```

```

✅
tsx

const Index: React.FC = () => {
  return (
    <Player
      durationInFrames={300}
      fps={30}
      compositionWidth={1080}
      compositionHeight={1080}
      component={AnotherComp}
    />
  );
};
```

## See also [​](\#see-also "Direct link to See also")

- [`<Composition>`](/docs/composition)
- [`<Player>`](/docs/player)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/wrong-composition-mount.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Slow method to extract frame](/docs/slow-method-to-extract-frame) [Next\
\
staticFile() relative paths](/docs/staticfile-relative-paths)

- [Problem](#problem)
- [Solution](#solution)
  - [In the Remotion Studio](#in-the-remotion-studio)
  - [In the Remotion Player](#in-the-remotion-player)
- [See also](#see-also)
