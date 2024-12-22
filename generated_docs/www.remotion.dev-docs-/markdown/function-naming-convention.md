---
title: Function naming convention
source: Remotion Documentation
last_updated: 2024-12-22
---

# Function naming convention

- [Home page](/)
- [Lambda](/docs/lambda)
- Function naming convention

On this page

# Function naming convention

A Remotion Lambda function has the following naming convention:

```

remotion-render-4-0-140-mem2048mb-disk2048mb-240sec
                ^^^^^^^    ^^^^       ^^^^   ^^^
                  |         |          |      |-- Timeout in seconds
                  |         |          |--------- Disk size in MB
                  |         |------------------- Memory size in MB
                  |----------------------------- Remotion version with dots replaced by dashes
```

```

remotion-render-4-0-140-mem2048mb-disk2048mb-240sec
                ^^^^^^^    ^^^^       ^^^^   ^^^
                  |         |          |      |-- Timeout in seconds
                  |         |          |--------- Disk size in MB
                  |         |------------------- Memory size in MB
                  |----------------------------- Remotion version with dots replaced by dashes
```

When you deploy a function, the name is hardcoded. This has two functions:

- [1](#1) Avoid having multiple functions with the same configuration unnecessarily.
- [2](#2) Being able to guess the function name using [`speculateFunctionName()`](/docs/lambda/speculatefunctionname) so you don't have to fetch a list of functions in your AWS account.

## Why can't I rename the function? [​](\#why-cant-i-rename-the-function "Direct link to Why can't I rename the function?")

- The [`npx remotion lambda render`](/docs/lambda/cli/render) command looks for functions that match this convention.
- With the default user policy, Remotion Lambda restricts itself from accessing functions that don't match this convention.
- You can use the [`speculateFunctionName()`](/docs/lambda/speculatefunctionname) function save one API call.
- The function is more likely to be warm from a previous invocation if you only have 1 function with the same configuration.
- There is no benefit to renaming the function.

## What if I want to have two functions for two different projects? [​](\#what-if-i-want-to-have-two-functions-for-two-different-projects "Direct link to What if I want to have two functions for two different projects?")

A function is not tied to a project.

Each function is a binary that contains the same code.

Every Remotion Lambda user runs the exact same code in their function.

The React code you write is not contained in the function, it is hosted on the Serve URL.

Each function invocation is isolated and they cannot conflict each other. There is a [concurrency limit](/docs/lambda/troubleshooting/rate-limit#default-concurrency-limits), but it is per region, not per function.

## I need to separate production, staging and development [​](\#i-need-to-separate-production-staging-and-development "Direct link to I need to separate production, staging and development")

Function invocations don't conflict each other.

Functions also don't contain any code that you write, they are binary and every Remotion Lambda user runs the exact same code.

It is impossible for a bad staging deployment to affect the production function.

Therefore, we recommend to use the same function for all environments.

## I want to have different configurations for different functions [​](\#i-want-to-have-different-configurations-for-different-functions "Direct link to I want to have different configurations for different functions")

This is supported! You can have multiple functions with different configurations.

It is only not possible to have multiple functions that share the exact same configuration:

- Timeout
- RAM
- Region
- Disk size
- Remotion version

To distinguish which function should be used, pass the function name explicitly to [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda).

You can pass [`--function-name`](/docs/lambda/cli/render#--function-name) to [`npx remotion lambda render`](/docs/lambda/cli/render)

## I want to deploy multiple projects [​](\#i-want-to-deploy-multiple-projects "Direct link to I want to deploy multiple projects")

It is possible deploy multiple sites under different [Serve URLs](/docs/terminology/serve-url).

This convention only applies to functions, which do not contain any code that you write.

## I want to deploy multiple functions to load-balance between them [​](\#i-want-to-deploy-multiple-functions-to-load-balance-between-them "Direct link to I want to deploy multiple functions to load-balance between them")

You do not need to do this, because you can invoke a function multiple times concurrently.

There is no concurrency limit per function, but a concurrency limit per region and account.

Therefore there is no benefit in having multiple identical functions in the same region and account for load-balancing.

## See also [​](\#see-also "Direct link to See also")

- [`speculateFunctionName()`](/docs/lambda/speculatefunctionname)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/naming-convention.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Cannot create public bucket](/docs/lambda/s3-public-access) [Next\
\
Cloud Run Alpha](/docs/cloudrun-alpha)

- [Why can't I rename the function?](#why-cant-i-rename-the-function)
- [What if I want to have two functions for two different projects?](#what-if-i-want-to-have-two-functions-for-two-different-projects)
- [I need to separate production, staging and development](#i-need-to-separate-production-staging-and-development)
- [I want to have different configurations for different functions](#i-want-to-have-different-configurations-for-different-functions)
- [I want to deploy multiple projects](#i-want-to-deploy-multiple-projects)
- [I want to deploy multiple functions to load-balance between them](#i-want-to-deploy-multiple-functions-to-load-balance-between-them)
- [See also](#see-also)