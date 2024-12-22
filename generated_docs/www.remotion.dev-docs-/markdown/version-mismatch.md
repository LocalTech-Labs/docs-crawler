---
title: Version mismatch
source: Remotion Documentation
last_updated: 2024-12-22
---

# Version mismatch

- [Home page](/)
- Troubleshooting
- Version mismatch

On this page

# Version mismatch

It is important that all Remotion packages you install ( `remotion`, `@remotion/player`, `@remotion/cli`, etc.) have the same version installed. If the versions differ, the packages might not be able to correctly communicate with each other which can introduce subtle bugs or even complete breakage.

## `^` character can be confusing [​](\#-character-can-be-confusing "Direct link to -character-can-be-confusing")

The following `package.json` does **not** ensure that all Remotion versions are the same:

```

package.json
json

{
  "dependencies": {
    "remotion": "^2.6.12",
    "@remotion/player": "^2.6.12",
    "@remotion/gif": "^2.6.12"
  }
}
```

```

package.json
json

{
  "dependencies": {
    "remotion": "^2.6.12",
    "@remotion/player": "^2.6.12",
    "@remotion/gif": "^2.6.12"
  }
}
```

however the following manifest does:

```

package.json
json

{
  "dependencies": {
    "remotion": "2.6.12",
    "@remotion/player": "2.6.12",
    "@remotion/gif": "2.6.12"
  }
}
```

```

package.json
json

{
  "dependencies": {
    "remotion": "2.6.12",
    "@remotion/player": "2.6.12",
    "@remotion/gif": "2.6.12"
  }
}
```

This is especially true if the packages are installed over time and not all at once. Removing the `^` character is a good way to ensure all Remotion packages have the same version.

## Packages that depend on Remotion [​](\#packages-that-depend-on-remotion "Direct link to Packages that depend on Remotion")

If you make a Remotion library and put it on npm, make Remotion a `peerDependency` and `devDependency` instead of a dependency:

```

package.json
json

{
  "name": "my-remotion-library",
  "peerDependencies": {
    "remotion": "*"
  },
  "devDependencies": {
    "remotion": "^2.6.11"
  }
}
```

```

package.json
json

{
  "name": "my-remotion-library",
  "peerDependencies": {
    "remotion": "*"
  },
  "devDependencies": {
    "remotion": "^2.6.11"
  }
}
```

This will make sure that when you then install `my-remotion-library` into your Remotion project, that no duplicates of Remotion will be introduced:

```

json

{
  "dependencies": {
    // No version mismatch will be introduced because `remotion`
    // is not a direct dependency of `my-remotion-library`
    "remotion": "2.7.0",
    "my-remotion-library": "1.0.0"
  }
}
```

```

json

{
  "dependencies": {
    // No version mismatch will be introduced because `remotion`
    // is not a direct dependency of `my-remotion-library`
    "remotion": "2.7.0",
    "my-remotion-library": "1.0.0"
  }
}
```

## Checking which versions have been installed [​](\#checking-which-versions-have-been-installed "Direct link to Checking which versions have been installed")

Just because the version in the package.json is `^2.6.11` does not mean that the version `2.6.11` got installed. Use the command:

```

bash

npx remotion versions
```

```

bash

npx remotion versions
```

to see which packages and versions your project contains.

## Testing prereleases [​](\#testing-prereleases "Direct link to Testing prereleases")

If we have given you a prerelease version of Remotion to test, the same principle applies, where you should not just upgrade the affected package, but all packages including `remotion` and packages starting with `@remotion`.

**See also:** [Testing prereleases](/docs/prereleases)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/version-mismatch.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Flickering](/docs/flickering) [Next\
\
ENAMETOOLONG](/docs/enametoolong)

- [`^` character can be confusing](#-character-can-be-confusing)
- [Packages that depend on Remotion](#packages-that-depend-on-remotion)
- [Checking which versions have been installed](#checking-which-versions-have-been-installed)
- [Testing prereleases](#testing-prereleases)
