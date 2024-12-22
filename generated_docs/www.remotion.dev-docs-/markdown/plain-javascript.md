---
title: Plain JavaScript
source: Remotion Documentation
last_updated: 2024-12-22
---

# Plain JavaScript

- [Home page](/)
- Tooling
- Plain JavaScript

On this page

# Plain JavaScript

Since Remotion 1.3, you can opt out of Typescript and it's type checking advantages in Remotion. Continue at your own risk.

## Opting out of Typescript [​](\#opting-out-of-typescript "Direct link to Opting out of Typescript")

You may import `.js` and `.jsx` files as normal. If you would like to completely move to JS, rename `index.ts` and `Root.tsx` so they have a `.jsx` file extension. Remove types such as `React.FC` and `SpringConfig`.

## Upgrading [​](\#upgrading "Direct link to Upgrading")

If you upgrade from Remotion 1.2 or older, consider changing the `npm test` command to also include JavaScript files, and to remove the `tsc` command:

```

diff

-  "test": "eslint src --ext ts,tsx && tsc"
+  "test": "eslint src --ext ts,tsx,js,jsx"
```

```

diff

-  "test": "eslint src --ext ts,tsx && tsc"
+  "test": "eslint src --ext ts,tsx,js,jsx"
```

## See also [​](\#see-also "Direct link to See also")

- [Custom Webpack config](/docs/webpack) for more advanced tweaking

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/jsx-support.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Using legacy Babel transpilation](/docs/legacy-babel) [Next\
\
Third party integrations](/docs/third-party)

- [Opting out of Typescript](#opting-out-of-typescript)
- [Upgrading](#upgrading)
- [See also](#see-also)
