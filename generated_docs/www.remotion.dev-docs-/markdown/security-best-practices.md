---
title: Security Best Practices
source: Remotion Documentation
last_updated: 2024-12-22
---

# Security Best Practices

- [Home page](/)
- Miscellaneous
- Security

On this page

# Security Best Practices

Remotion is a software that can generate videos programmatically. It does so by running a headless browser, capturing images, and encoding them into a video. This in itself does not pose a security risk. However, Remotion interacts with the Web, can be used in the cloud and uses the NPM ecosystem.

## Environment variables [​](\#environment-variables "Direct link to Environment variables")

When starting the Remotion Studio or running a render via the command line, environment variables that are prefixed with `REMOTION_` and `.env` are passed to the headless browser. You are responsible to ensure those environment variables do not exposed to other services as you are interfacing with the Web.

## Credentials on the web [​](\#credentials-on-the-web "Direct link to Credentials on the web")

You should not try to call `renderMediaOnLambda()` or other APIs from `@remotion/lambda` from your app frontend. Those APIs can only be called with AWS credentials, meaning that you would have to expose your AWS credentials to the web. This is a security risk and should be avoided.

## `disableWebSecurity` flag [​](\#disablewebsecurity-flag "Direct link to disablewebsecurity-flag")

We provide as a workaround for CORS issues the `disableWebSecurity` flag. This flag disables the same-origin policy in the headless browser. If you are using it, you need to be aware of the implications of it and assess whether this can pose a security risk to you.

## DDoS attacks [​](\#ddos-attacks "Direct link to DDoS attacks")

APIs such as `renderMedia()` and `renderMediaOnLambda()` may be expensive (computationally or financially) to run. You should ensure that you are not exposing those APIs to the web without authentication or rate limiting. Otherwise, you may be vulnerable to DDoS attacks.

## General Node.JS security best practices [​](\#general-nodejs-security-best-practices "Direct link to General Node.JS security best practices")

Often you will use Remotion alongside other Node.js dependencies.

You should only install dependencies you trust, since they can run arbitrary code in their `postinstall` scripts.
Dependencies, including Remotion, may have known vulnerabilities which you can list with the `npm audit` command.

## Reporting of security vulnerabilities [​](\#reporting-of-security-vulnerabilities "Direct link to Reporting of security vulnerabilities")

If you believe to have found a security vulnerability, you can report it to [hi@remotion.dev](mailto:hi@remotion.dev)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/security.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Support Policy](/docs/support) [Next\
\
Chromium flags](/docs/chromium-flags)

- [Environment variables](#environment-variables)
- [Credentials on the web](#credentials-on-the-web)
- [`disableWebSecurity` flag](#disablewebsecurity-flag)
- [DDoS attacks](#ddos-attacks)
- [General Node.JS security best practices](#general-nodejs-security-best-practices)
- [Reporting of security vulnerabilities](#reporting-of-security-vulnerabilities)
