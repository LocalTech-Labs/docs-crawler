---
title: Debugging failed Lambda renders
source: Remotion Documentation
last_updated: 2024-12-22
---

# Debugging failed Lambda renders

- [Home page](/)
- [Lambda](/docs/lambda)
- Troubleshooting
- Debugging failures

On this page

# Debugging failed Lambda renders

[![](https://i.ytimg.com/vi/pwVEzTQ6VYE/hqdefault.jpg?sqp=-oaymwEbCKgBEF5IVfKriqkDDggBFQAAiEIYAXABwAEG&rs=AOn4CLDGudSHhI9u7kkShN_awBk3hbQ9mA)\
\
Also available as a 11min video\
\
**How to troubleshoot and debug for Remotion Lambda**](https://youtu.be/pwVEzTQ6VYE)

## General debugging strategy [​](\#general-debugging-strategy "Direct link to General debugging strategy")

[1](#1)

Reproduce a render via the [`npx remotion lambda render`](/docs/lambda/cli/render) command.

- Use `--props` to pass in a JSON file with the same input props as the render
you want to debug. - Use `--log=verbose` to get detailed debugging information.

[2](#2)

Wait for the render to fully time out or throw an error.

[3](#3) Read the error message carefully and use the log links that were
printed.

## Reasons a render may fail [​](\#reasons-a-render-may-fail "Direct link to Reasons a render may fail")

There are four reasons renders may fail:

- [1](#1) An error occurs in your React code.
- [2](#2) A timeout occurs after calling
[`delayRender()`](/docs/delay-render).
- [3](#3) The Lambda function rendering one of the chunks times out.
- [4](#4) The main Lambda function times out.

## Error in React code [​](\#error-in-react-code "Direct link to Error in React code")

If your code throws an error, you will see the error in the logs. If you are rendering via the CLI, the error will be symbolicated which can help you find the error in your code.
Try to fix the error, redeploy the site and re-render.

## Timeouts [​](\#timeouts "Direct link to Timeouts")

### Understanding which timeout occurred [​](\#understanding-which-timeout-occurred "Direct link to Understanding which timeout occurred")

If your error message reads that a [`delayRender()`](/docs/delay-render) call timed out, a function has called [`delayRender()`](/docs/delay-render) but did not call [`continueRender()`](/docs/delay-render) within the timeout period. This is usually caused by a bug in your code. See [Debugging timeout](/docs/timeout) for more tips.

note

You can increase the timeout by passing `timeoutInMilliseconds` to [`renderMediaOnLamba()`](/docs/lambda/rendermediaonlambda) or passing `--timeout` to the CLI when rendering a video. Don't confuse it with the `--timeout` flag for the CLI when deploying a function (see below).

If your error message reads that the main Lambda function has timed out, it
means that the render was still ongoing, but the maximum execution duration of
the Lambda function has been hit. This is caused by either:

- A render that takes too long to complete within the specified duration. In that case, you should increase the function timeout by passing `--timeout` to the CLI [when deploying a function](/docs/lambda/cli/functions).
- A render that has a bottleneck and completes slowly on Lambda. In that case, you should identify the bottleneck by measuring and optimizing the render. See [Optimizing for speed](/docs/lambda/optimizing-speed) for more tips.
- A chunk is getting stuck, making it impossible for the main function to complete the task of concatenating the chunks. You should identify the chunk that is getting stuck by looking at the logs. See below for how to do this.

### Inspecting the logs [​](\#inspecting-the-logs "Direct link to Inspecting the logs")

[1](#1)

Get the Log URL:

- If you use the CLI, add `--log=verbose` to the command. This will print a CloudWatch link to the console.
- When using [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda) add `logLevel: "verbose"` as an option. You will get a `cloudWatchLogs` field in the return value.

Open this link and log into AWS if needed. A log stream will open.

![](/img/cloudwatch.png)

By default, the filter is set to:

```

"method=renderer,renderId=[render-id]"
```

```

"method=renderer,renderId=[render-id]"
```

### Find the logs of a specific chunk [​](\#find-the-logs-of-a-specific-chunk "Direct link to Find the logs of a specific chunk")

Tweak the query to find the logs of a specific chunk:

```

method=renderer,renderId=[render-id],chunk=12
```

```

method=renderer,renderId=[render-id],chunk=12
```

will for example find the log stream of chunk 12. If your viewport is big enough, you will also see the chunk numbers in the summary view straight away.

Click the blue link in the `Log stream name` column to open the log stream. If you don't see any blue links, click "Display", then select "View in columns with details".

In the logstream, you will see debug logging from Remotion as well as any `console.log` statements from your React code.

### Find the logs of the main function [​](\#find-the-logs-of-the-main-function "Direct link to Find the logs of the main function")

Tweakt the query to find the logs of the main function:

```

method=launch,renderId=[render-id]
```

```

method=launch,renderId=[render-id]
```

You should get one result.
Click the blue link in the `Log stream name` column to open the log stream. If you don't see any blue links, click "Display", then select "View in columns with details".

### Finding the chunk that failed [​](\#finding-the-chunk-that-failed "Direct link to Finding the chunk that failed")

To find which chunks failed to render, add `--log=verbose` to the Lambda render while rendering via the CLI. Look for a log statement `Render folder: ` to find the link to the render folder in S3.

When using [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda), you will get a `folderInS3Console` field in the return value.

After opening the link, open the `chunks` subfolder. Go through each existing chunk and find the ones that are missing. The missing chunks are the ones that failed to render.

![](/img/chunks.png)

After you've identified which chunks are missing, [inspect the logs](#inspecting-the-logs) for that chunk to find the cause of the error.

### Increasing the timeout [​](\#increasing-the-timeout "Direct link to Increasing the timeout")

Note that two types of timeouts come into play in Remotion:

- The `delayRender()` timeout. This is the timeout that is defined using `--timeout` when calling [`npx remotion lambda render`](/docs/lambda/cli/render) or using `timeoutInMilliseconds` when calling [`renderMediaOnLambda()`](/docs/lambda/rendermediaonlambda).
- The Lambda function timeout. This is the timeout that is defined using `--timeout` when calling [`npx remotion lambda functions deploy`](/docs/lambda/cli/functions) or using `timeoutInSeconds` when calling [`deployFunction()`](/docs/lambda/deployfunction).

### Reporting an issue [​](\#reporting-an-issue "Direct link to Reporting an issue")

You may report an issue on [GitHub](https://remotion.dev/help) and [Discord](https://remotion.dev/discord).

Include the `progress.json` that lies in the S3 folder of the render, because it includes all informations to reproduce a render.

If you include CloudWatch logs, post the full ones (click "Load more messages" until there are no more to load.)

## See also [​](\#see-also "Direct link to See also")

- [Debugging timeouts](/docs/timeout)
- [Optimizing for speed](/docs/lambda/optimizing-cost)
- [Rate limits](/docs/lambda/troubleshooting/rate-limit)
- [Lambda Limits](/docs/lambda/limits)
- [Lambda insights](/docs/lambda/insights)

[Improve this page](https://github.com/remotion-dev/remotion/edit/main/packages/docs/docs/lambda/troubleshooting/debug.mdx)

[Ask on Discord](https://remotion.dev/discord)

[Get help](/docs/get-help)

Last updated on **Dec 20, 2024**

[Previous\
\
Auto-delete renders](/docs/lambda/autodelete) [Next\
\
Permissions](/docs/lambda/troubleshooting/permissions)

- [General debugging strategy](#general-debugging-strategy)
- [Reasons a render may fail](#reasons-a-render-may-fail)
- [Error in React code](#error-in-react-code)
- [Timeouts](#timeouts)
  - [Understanding which timeout occurred](#understanding-which-timeout-occurred)
  - [Inspecting the logs](#inspecting-the-logs)
  - [Find the logs of a specific chunk](#find-the-logs-of-a-specific-chunk)
  - [Find the logs of the main function](#find-the-logs-of-the-main-function)
  - [Finding the chunk that failed](#finding-the-chunk-that-failed)
  - [Increasing the timeout](#increasing-the-timeout)
  - [Reporting an issue](#reporting-an-issue)
- [See also](#see-also)
