---
title: React 19
source: Remotion Documentation
last_updated: 2024-12-22
---

# React 19

- [Home page](/)
- Migration guides
- React 19

On this page

# React 19

## Updating packages [​](\#updating-packages "Direct link to Updating packages")

To use React 19's newest features, you need at least version `4.0.0` of Remotion.

You need to upgrade both `react` and `react-dom`:

```

diff

- "react": "18.3.1"
- "react-dom": "18.3.1"
+ "react": "19.0.0"
+ "react-dom": "19.0.0"
```

```

diff

- "react": "18.3.1"
- "react-dom": "18.3.1"
+ "react": "19.0.0"
+ "react-dom": "19.0.0"
```

If you use TypeScript, update to the newest types as well:

```

diff

- "@types/react": "18.3.1"
- "@types/react-dom": "18.3.1"
+ "@types/react": "19.0.0"
+ "@types/react-dom": "19.0.0"
```

```

diff

- "@types/react": "18.3.1"
- "@types/react-dom": "18.3.1"
+ "@types/react": "19.0.0"
+ "@types/react-dom": "19.0.0"
```

Run `npm i`, `bun i`, `yarn` or `pnpm i` afterwards, matching your package manager.

## Updated templates [​](\#updated-templates "Direct link to Updated templates")

We have updated all [templates](https://remotion.dev/templates) to use React 19 (exception: React Native Skia).

See the source code of the templates for examples on how to upgrade it to React 19.

## `HTMLRefElement` Type Change [​](\#htmlrefelement-type-change "Direct link to htmlrefelement-type-change")

If you have type errors related to React Refs, upgrade to v4.0.236 of Remotion, where we aligned the types with React 19.

## Ecosystem support [​](\#ecosystem-support "Direct link to Ecosystem support")

Some libraries that are used with Remotion need upgrading.

### React Three Fiber [​](\#react-three-fiber "Direct link to React Three Fiber")

Update to `9.0.0-rc.1` and `three` to `0.171.0`.

### `styled-components` [​](\#styled-components "Direct link to styled-components")

Update to v6.

### Next.js [​](\#nextjs "Direct link to Next.js")

Update to Next.js 15.

### React Native Skia [​](\#react-native-skia "Direct link to React Native Skia")

No React 19 support yet. Stay on React 18.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/react-19.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
List of resources](/docs/resources) [Next\
\
v5.0 Migration](/docs/5-0-migration)

- [Updating packages](#updating-packages)
- [Updated templates](#updated-templates)
- [`HTMLRefElement` Type Change](#htmlrefelement-type-change)
- [Ecosystem support](#ecosystem-support)
  - [React Three Fiber](#react-three-fiber)
  - [`styled-components`](#styled-components)
  - [Next.js](#nextjs)
  - [React Native Skia](#react-native-skia)
