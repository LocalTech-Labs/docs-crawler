---
title: February 2022 Outage
source: Remotion Documentation
last_updated: 2024-12-22
---

# February 2022 Outage

On this page

# February 2022 Outage

tip

**Update 2022/02/06:**

The problem is now solved. Your Lambda functions should work as normal. We recommend everyone to go back to use ARM64 Lambdas. We will consult with AWS support on how to prevent issues like this in the future.

On February 3rd 2022, AWS made a change to their Lambda micro-VMs that breaks Remotion Lambda. This document contains information for affected users.

## Hotfix solution [​](\#hotfix-solution "Direct link to Hotfix solution")

[Upgrade](/docs/lambda/upgrading) your Lambda functions to the latest version and deploy them with a `x86_64` architecture.

Via CLI:

```

npx remotion lambda functions deploy --architecture=x86_64
```

```

npx remotion lambda functions deploy --architecture=x86_64
```

Via Node.JS:

## Example [​](\#example "Direct link to Example")

```

ts

// @module: esnext
// @target: es2017

import { deployFunction } from "@remotion/lambda";

const { functionName } = await deployFunction({
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 1024,
  createCloudWatchLogGroup: true,
  architecture: "x86_64",
});
console.log(functionName);
```

```

ts

// @module: esnext
// @target: es2017

import { deployFunction } from "@remotion/lambda";

const { functionName } = await deployFunction({
  region: "us-east-1",
  timeoutInSeconds: 120,
  memorySizeInMb: 1024,
  createCloudWatchLogGroup: true,
  architecture: "x86_64",
});
console.log(functionName);
```

## Caveat [​](\#caveat "Direct link to Caveat")

The x86\_64 version **does not contain fonts for Japanese/Chinese/Korean**! Since the binaries are bigger on the **x86\_64** version, we exceeded the file limit and had to delete the biggest font in order to stay within the AWS bounds.

x86\_64 is also slower and has about a 30% higher cost/performance ratio.

## AWS mitigation [​](\#aws-mitigation "Direct link to AWS mitigation")

We contacted AWS and hope to receive an answer soon.

## Long-term solution [​](\#long-term-solution "Direct link to Long-term solution")

ARM64 is the preferred and default solution. Once the problem is solved, we recommend everyone to switch back.

## Contact [​](\#contact "Direct link to Contact")

Join the discussion in our Discord channel if you have questions.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/feb-2022-outage.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

- [Hotfix solution](#hotfix-solution)
- [Example](#example)
- [Caveat](#caveat)
- [AWS mitigation](#aws-mitigation)
- [Long-term solution](#long-term-solution)
- [Contact](#contact)
