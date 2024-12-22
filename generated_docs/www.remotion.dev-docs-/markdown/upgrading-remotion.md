---
title: Upgrading Remotion
source: Remotion Documentation
last_updated: 2024-12-22
---

# Upgrading Remotion

- [Home page](/)
- Upgrading Remotion

On this page

# Upgrading Remotion

The easiest way to do this is to run the following command in the root of your project:

- npx
- pnpm
- yarn
- bunx

```

bash

npx remotion upgrade
```

```

bash

npx remotion upgrade
```

```

bash

pnpm exec remotion upgrade
```

```

bash

pnpm exec remotion upgrade
```

```

bash

yarn remotion upgrade
```

```

bash

yarn remotion upgrade
```

```

bash

bunx remotion upgrade
```

```

bash

bunx remotion upgrade
```

note

You need the `@remotion/cli` package installed for this.

## Manually upgrading Remotion [​](\#manually-upgrading-remotion "Direct link to Manually upgrading Remotion")

To upgrade Remotion to the latest version, all `@remotion/*` packages and `remotion` must be updated to the same version.

Edit the version number in your `package.json` for all packages.
Delete the `^` in front of the version number in your `package.json` in order to force the exact version you specified.

## Breaking changes [​](\#breaking-changes "Direct link to Breaking changes")

Remotion follows semantic versioning.

This means if the first number of the version is the same, you can upgrade and your code is backwards-compatible.

Example: If you are on 4.0.0, you can upgrade to 4.1.100 without changing your code.

However, to upgrade to 5.0, you will need to follow the [migration guide](/docs/5-0-migration).

Exceptions to the breaking change rule are APIs that are marked as experimental.

## Changelog [​](\#changelog "Direct link to Changelog")

Visit [remotion.dev/changelog](https://remotion.dev/changelog) to see a list of all changes.

## Stable versions [​](\#stable-versions "Direct link to Stable versions")

We maintain a repo with the latest stable version of Remotion for customers who need a higher level of stability.

Customers may get access to the repo on the [remotion.pro](https://remotion.pro) portal.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/upgrading.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Combining compositions](/docs/miscellaneous/snippets/combine-compositions) [Next\
\
Terminology](/docs/terminology)

- [Manually upgrading Remotion](#manually-upgrading-remotion)
- [Breaking changes](#breaking-changes)
- [Changelog](#changelog)
- [Stable versions](#stable-versions)
