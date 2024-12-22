---
title: Upgrade to React 18
source: Remotion Documentation
last_updated: 2024-12-22
---

# Upgrade to React 18

On this page

# Upgrade to React 18

## Updating packages [​](\#updating-packages "Direct link to Updating packages")

To use React 18's newest features, you need at least version `3.0.0` of Remotion.

```

diff

- "remotion": "2.6.15"
- "@remotion/bundler": "2.6.15"
- "@remotion/cli": "2.6.15"
- "@remotion/renderer": "2.6.15"
+ "remotion": "3.0.0"
+ "@remotion/bundler": "3.0.0"
+ "@remotion/cli": "3.0.0"
+ "@remotion/renderer": "3.0.0"
```

```

diff

- "remotion": "2.6.15"
- "@remotion/bundler": "2.6.15"
- "@remotion/cli": "2.6.15"
- "@remotion/renderer": "2.6.15"
+ "remotion": "3.0.0"
+ "@remotion/bundler": "3.0.0"
+ "@remotion/cli": "3.0.0"
+ "@remotion/renderer": "3.0.0"
```

You need to upgrade both `react` and `react-dom`:

```

diff

- "react": "17.0.1"
- "react-dom": "17.0.1"
+ "react": "18.2.0"
+ "react-dom": "18.2.0"
```

```

diff

- "react": "17.0.1"
- "react-dom": "17.0.1"
+ "react": "18.2.0"
+ "react-dom": "18.2.0"
```

If you use TypeScript, update to the newest types as well:

```

diff

- "@types/react": "17.0.3"
+ "@types/react": "18.0.0"
```

```

diff

- "@types/react": "17.0.3"
+ "@types/react": "18.0.0"
```

Run `npm i`, or `yarn`, or `pnpm i` afterwards, matching your package manager.

## Fixing `React.FC` types [​](\#fixing-reactfc-types "Direct link to fixing-reactfc-types")

The types for `React.FC` have changed to no longer include `children`. If you get a type error, change

```

tsx

const MyComp: React.FC = ({ children }) => <div>{children}</div>;
```

```

tsx

const MyComp: React.FC = ({ children }) => <div>{children}</div>;
```

to:

```

tsx

const MyComp: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => <div>{children}</div>;
```

```

tsx

const MyComp: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => <div>{children}</div>;
```

## No action required for `createRoot()` [​](\#no-action-required-for-createroot "Direct link to no-action-required-for-createroot")

If React 18 is installed, Remotion will automatically use `createRoot()` from `react-dom/client` instead of `render` from `react-dom`. You don't need to do anything.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/react-18.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

- [Updating packages](#updating-packages)
- [Fixing `React.FC` types](#fixing-reactfc-types)
- [No action required for `createRoot()`](#no-action-required-for-createroot)
