---
title: Chromium flags
source: Remotion Documentation
last_updated: 2024-12-22
---

# Chromium flags

- [Home page](/)
- Miscellaneous
- Chromium flags

On this page

# Chromium flags

We allow you to set the following flags in Chromium and Google Chrome since Remotion 2.6.5:

## `--disable-web-security` [​](\#--disable-web-security "Direct link to --disable-web-security")

This will most notably disable CORS among other security features.

note

Remotion will automatically append the `--user-data-dir` flag.

### Via Node.JS APIs [​](\#via-nodejs-apis "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions), [`renderStill()`](/docs/renderer/render-still#disablewebsecurity), [`renderMedia()`](/docs/renderer/render-media#disablewebsecurity), [`renderFrames()`](/docs/renderer/render-frames#disablewebsecurity), [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda#disablewebsecurity), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#disablewebsecurity) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#disablewebsecurity), you can pass [`chromiumOptions.disableWebSecurity`](/docs/renderer/render-still#disablewebsecurity).

### Via CLI flag [​](\#via-cli-flag "Direct link to Via CLI flag")

Pass [`--disable-web-security`](/docs/cli/render#--disable-web-security) in one of the following commands: [`remotion render`](/docs/cli/render), [`remotion compositions`](/docs/cli/compositions), [`remotion still`](/docs/cli/still), [`remotion lambda render`](/docs/lambda/cli/render), [`remotion lambda still`](/docs/lambda/cli/still), [`remotion lambda compositions`](/docs/lambda/cli/compositions).

### Via config file [​](\#via-config-file "Direct link to Via config file")

Use [setChromiumDisableWebSecurity()](/docs/config#setchromiumdisablewebsecurity).

```

tsx

Config.setChromiumDisableWebSecurity(true);
```

```

tsx

Config.setChromiumDisableWebSecurity(true);
```

note

Prior to `v3.3.39`, the option was called `Config.Puppeteer.setChromiumDisableWebSecurity()`.

## `--ignore-certificate-errors` [​](\#--ignore-certificate-errors "Direct link to --ignore-certificate-errors")

Results in invalid SSL certificates, such as self-signed ones, being ignored.

### Via Node.JS APIs [​](\#via-nodejs-apis-1 "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions), [`renderStill()`](/docs/renderer/render-still#ignorecertificateerrors), [`renderMedia()`](/docs/renderer/render-media#ignorecertificateerrors), [`renderFrames()`](/docs/renderer/render-frames#ignorecertificateerrors), [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda#disablewebsecurity), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#ignorecertificateerrors) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#ignorecertificateerrors), you can pass [`chromiumOptions.ignoreCertificateErrors`](/docs/renderer/render-still#ignorecertificateerrors).

### Via CLI flag [​](\#via-cli-flag-1 "Direct link to Via CLI flag")

Pass [`--ignore-certificate-errors`](/docs/cli/render#--ignore-certificate-errors) in one of the following commands: [`remotion render`](/docs/cli/render), [`remotion compositions`](/docs/cli/compositions), [`remotion still`](/docs/cli/still), [`remotion lambda render`](/docs/lambda/cli/render), [`remotion lambda still`](/docs/lambda/cli/still), [`remotion lambda compositions`](/docs/lambda/cli/compositions).

### Via config file [​](\#via-config-file-1 "Direct link to Via config file")

Use [setChromiumIgnoreCertificateErrors()](/docs/config#setchromiumignorecertificateerrors).

```

tsx

Config.setChromiumIgnoreCertificateErrors(true);
```

```

tsx

Config.setChromiumIgnoreCertificateErrors(true);
```

note

Prior to `v3.3.39`, the option was called `Config.Puppeteer.setChromiumIgnoreCertificateErrors()`.

## `--disable-headless` [​](\#--disable-headless "Direct link to --disable-headless")

Deprecated - will be removed in 5.0.0. With the migration to [Chrome Headless Shell](/docs/miscellaneous/chrome-headless-shell), this option is not functional anymore.

If disabled, the render will open an actual Chrome window where you can see the render happen. The default is headless mode.

### Via Node.JS APIs [​](\#via-nodejs-apis-2 "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions), [`renderStill()`](/docs/renderer/render-still#headless), [`renderMedia()`](/docs/renderer/render-media#headless) and [`renderFrames()`](/docs/renderer/render-frames#headless), you can pass [`chromiumOptions.headless`](/docs/renderer/render-still#headless). You cannot set this option in Lambda.

### Via CLI flag [​](\#via-cli-flag-2 "Direct link to Via CLI flag")

Pass [`--disable-headless`](/docs/cli/render#--disable-headless) in one of the following commands: [`remotion compositions`](/docs/cli/compositions), [`remotion render`](/docs/cli/render), [`remotion still`](/docs/cli/still).

### Via config file [​](\#via-config-file-2 "Direct link to Via config file")

Use [setChromiumHeadlessMode()](/docs/config#setchromiumheadlessmode).

```

tsx

Config.setChromiumHeadlessMode(false);
```

```

tsx

Config.setChromiumHeadlessMode(false);
```

note

Prior to `v3.3.39`, the option was called `Config.Puppeteer.setChromiumHeadlessMode()`.

## `--gl` [​](\#--gl "Direct link to --gl")

Changelog

- From Remotion v2.6.7 until v3.0.7, the default for Remotion Lambda was `swiftshader`, but from v3.0.8 the default is `swangle` (Swiftshader on Angle) since Chrome 101 added support for it.
- From Remotion v2.4.3 until v2.6.6, the default was `angle`, however it turns out to have a small memory leak that could crash long Remotion renders.

Select the OpenGL renderer backend for Chromium.

Accepted values:

- `"angle"`
- `"egl"`
- `"swiftshader"`
- `"swangle"`
- `"vulkan"` ( _from Remotion v4.0.41_)
- `"angle-egl"` ( _from Remotion v4.0.51_)

The default is `null`, letting Chrome decide, except on Lambda where the default is `"swangle"`

### Via Node.JS APIs [​](\#via-nodejs-apis-3 "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions#chromiumoptions), [`renderStill()`](/docs/renderer/render-still#gl), [`renderMedia()`](/docs/renderer/render-media#gl), [`renderFrames()`](/docs/renderer/render-frames#gl), [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda#gl), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#gl) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#gl), you can pass [`chromiumOptions.gl`](/docs/renderer/render-still#gl).

### Via CLI flag [​](\#via-cli-flag-3 "Direct link to Via CLI flag")

Pass [`--gl=swiftshader`](/docs/cli) in one of the following commands: [`remotion render`](/docs/cli/render), [`remotion compositions`](/docs/cli/compositions), [`remotion still`](/docs/cli/still), [`remotion lambda render`](/docs/lambda/cli/render), [`remotion lambda still`](/docs/lambda/cli/still), [`remotion lambda compositions`](/docs/lambda/cli/compositions).

### Via config file [​](\#via-config-file-3 "Direct link to Via config file")

```

tsx

Config.setChromiumOpenGlRenderer('swiftshader');
```

```

tsx

Config.setChromiumOpenGlRenderer('swiftshader');
```

note

Prior to `v3.3.39`, the option was called `Config.Puppeteer.setChromiumOpenGlRenderer()`.

## `--user-agent` [v3.3.83](https://github.com/remotion-dev/remotion/releases/v3.3.83) [​](\#--user-agent "Direct link to --user-agent")

### Via Node.JS APIs [​](\#via-nodejs-apis-4 "Direct link to Via Node.JS APIs")

In [`getCompositions()`](/docs/renderer/get-compositions#chromiumoptions), [`renderStill()`](/docs/renderer/render-still#useragent), [`renderMedia()`](/docs/renderer/render-media#useragent), [`renderFrames()`](/docs/renderer/render-frames#useragent), [`getCompositionsOnLambda()`](/docs/lambda/getcompositionsonlambda#useragent), [`renderStillOnLambda()`](/docs/lambda/renderstillonlambda#useragent) and [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda#chromiumoptions), you can pass [`chromiumOptions.userAgent`](/docs/renderer/render-still#useragent).

### Via CLI flag [​](\#via-cli-flag-4 "Direct link to Via CLI flag")

Pass [`--user-agent`](/docs/cli) in one of the following commands: [`remotion render`](/docs/cli/render), [`remotion compositions`](/docs/cli/compositions), [`remotion still`](/docs/cli/still), [`remotion lambda render`](/docs/lambda/cli/render), [`remotion lambda still`](/docs/lambda/cli/still), [`remotion lambda compositions`](/docs/lambda/cli/compositions).

## Need more flags? [​](\#need-more-flags "Direct link to Need more flags?")

Open a [GitHub issue](https://github.com/remotion-dev/remotion/issues/new?assignees=&labels=&template=feature_request.md&title=) to request it.

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/chromium-flags.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Security](/docs/security) [Next\
\
Temporary directory](/docs/miscellaneous/changing-temp-dir)

- [`--disable-web-security`](#--disable-web-security)
  - [Via Node.JS APIs](#via-nodejs-apis)
  - [Via CLI flag](#via-cli-flag)
  - [Via config file](#via-config-file)
- [`--ignore-certificate-errors`](#--ignore-certificate-errors)
  - [Via Node.JS APIs](#via-nodejs-apis-1)
  - [Via CLI flag](#via-cli-flag-1)
  - [Via config file](#via-config-file-1)
- [`--disable-headless`](#--disable-headless)
  - [Via Node.JS APIs](#via-nodejs-apis-2)
  - [Via CLI flag](#via-cli-flag-2)
  - [Via config file](#via-config-file-2)
- [`--gl`](#--gl)
  - [Via Node.JS APIs](#via-nodejs-apis-3)
  - [Via CLI flag](#via-cli-flag-3)
  - [Via config file](#via-config-file-3)
- [`--user-agent`](#--user-agent)
  - [Via Node.JS APIs](#via-nodejs-apis-4)
  - [Via CLI flag](#via-cli-flag-4)
- [Need more flags?](#need-more-flags)
