---
title: @remotion/licensing
source: Remotion Documentation
last_updated: 2024-12-22
---

# @remotion/licensing

- [Home page](/)
- @remotion/licensing

On this page

# @remotion/licensing

_available from v4.0.237_

This package allows holders of the [company license](https://remotion.pro/license) to send events to Remotion to track the following usage of Remotion products:

- Cloud renders
- WebCodecs conversions

Also, it offers an API to programmatically check the usage in order to implement spend controls.

### How do I use this package? [​](\#how-do-i-use-this-package "Direct link to How do I use this package?")

On your company license dashboard on [remotion.pro](https://remotion.pro), you can find your public and secret API key.

You can use the public API key to register usage points and the secret API key to get the current usage.

### Do I need to use this package? [​](\#do-i-need-to-use-this-package "Direct link to Do I need to use this package?")

**For Remotion, the framework**: It is voluntary to use this package.

If you are not eligible for the free license, you need to get a company license and are required to keep your seat count up to date to cover all your renders.

This package makes it easier to count the number of Remotion renders without having to set up any infrastructure yourself.

**For [`@remotion/webcodecs`](/docs/webcodecs)**: This package is a dependency and conversions will trigger a call to [`registerUsagePoint()`](/docs/licensing/register-usage-point) automatically.

### Can I use this package to count renders if I am eligible for the free license? [​](\#can-i-use-this-package-to-count-renders-if-i-am-eligible-for-the-free-license "Direct link to Can I use this package to count renders if I am eligible for the free license?")

Yes, you may still create a project on [remotion.pro](https://remotion.pro) and use this package to count renders and WebCodecs conversions.

You do not have to pay anything for the renders. This package is only used to count the number of renders, not for billing.

### Will I get charged based on the usage? [​](\#will-i-get-charged-based-on-the-usage "Direct link to Will I get charged based on the usage?")

This package currently only counts renders, it does not bill you based on them.

If you are a company license holder, you need to manually adjust your seat count on [remotion.pro](https://remotion.pro) to cover all your renders.

However, the plan is to introduce a system that automatically bills you based on the usage, eliminating the need for you to manually adjust the seat count.

## Installation [​](\#installation "Direct link to Installation")

- npm
- yarn
- pnpm
- bun

```

npm i --save-exact @remotion/licensing@4.0.242Copy
```

```

npm i --save-exact @remotion/licensing@4.0.242Copy
```

```

pnpm i @remotion/licensing@4.0.242Copy
```

```

pnpm i @remotion/licensing@4.0.242Copy
```

```

bun i @remotion/licensing@4.0.242Copy
```

```

bun i @remotion/licensing@4.0.242Copy
```

```

yarn --exact add @remotion/licensing@4.0.242Copy
```

```

yarn --exact add @remotion/licensing@4.0.242Copy
```

This assumes you are currently using v4.0.242 of Remotion.

Also update `remotion` and all `` `@remotion/*` `` packages to the same version.

Remove all `^` character in front of the version numbers of it as it can lead to a version conflict.

## API [​](\#api "Direct link to API")

[**registerUsagePoint()** \
\
Register a cloud render or WebCodecs conversion](/docs/licensing/register-usage-point) [**getUsage()** \
\
Query usage of company license](/docs/licensing/get-usage)

## License [​](\#license "Direct link to License")

[Remotion License](https://remotion.dev/license)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/licensing/index.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
loadFont()](/docs/fonts-api/load-font) [Next\
\
registerUsagePoint()](/docs/licensing/register-usage-point)

- [How do I use this package?](#how-do-i-use-this-package)
- [Do I need to use this package?](#do-i-need-to-use-this-package)
- [Can I use this package to count renders if I am eligible for the free license?](#can-i-use-this-package-to-count-renders-if-i-am-eligible-for-the-free-license)
- [Will I get charged based on the usage?](#will-i-get-charged-based-on-the-usage)
- [Installation](#installation)
- [API](#api)
- [License](#license)
