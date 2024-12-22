Cloud RunOn this pageCloud Run AlphaHelp us shape Remotion Cloud Run!
Cloud Run is an alternative option to Remotion Lambda. Where Lambda offers a cloud-based rendering solution on AWS (Amazon Web Services), Cloud Run makes use of GCP (Google Cloud Platform).
What to test​
We are looking for feedback on the experience of setting up a GCP Project for Remotion Cloud Run, as well as the required components for rendering on the cloud:

Deploy a rendering service (in Lambda, a service is known as a function).
Deploy a Remotion project to GCP Cloud Storage (in Lambda, the storage solution is S3).
Render a composition stored in Cloud Storage on a Cloud Run service.

We are welcoming any bug reports.
1. Install @remotion/cloudrun​
npmyarnpnpmbunnpm i --save-exact @remotion/cloudrun@4.0.241Copynpm i --save-exact @remotion/cloudrun@4.0.241Copypnpm i @remotion/cloudrun@4.0.241Copypnpm i @remotion/cloudrun@4.0.241Copybun i @remotion/cloudrun@4.0.241Copybun i @remotion/cloudrun@4.0.241Copyyarn --exact add @remotion/cloudrun@4.0.241Copyyarn --exact add @remotion/cloudrun@4.0.241CopyThis assumes you are currently using v4.0.241 of Remotion.Also update remotion and all `@remotion/*` packages to the same version.Remove all ^ character in front of the version numbers of it as it can lead to a version conflict.
From v4.0.18, Cloud run is distributed together with the main release of Remotion. Before that, you had to install the alpha release (see below).
Changelog (moved)​
From 4.0.18 on, see changes here.
4.0.18​
Remotion Cloud Run is now distributed together with the main release of Remotion. You no longer need to switch to the alpha release, although Remotion Cloud Run is still alpha software. The changelog is now part of the main changelog.
4.1.0-alpha12​
Includes features and bugfixes from v4.0.17.
Includes a fix for streaming progress sometimes throwing an exception.
4.1.0-alpha11​
Includes bugfixes from v4.0.12.
4.1.0-alpha9​
Known issues​

any internal errors created by Remotion from within the service are not currently sent back in the error response to the renderMediaOnCloudrun and renderStillOnCloudrun APIs (these APIs are also used within the CLI). For these errors, users will need to check the logs for now.

Improvements​

Artifact registry, used to store versioned images for deploying services, now has two folders - production and development.
Provide helpful response when Cloud Run crashes during render.

CLI alerts user there was a crash, fetches logs, determines if cause was likely memory or timeout issue.
API can receive a success or crash response
New response documented


Default concurrency for rendering media is now 100%. This will set the concurrency equal to the number of cores the deployed service has.

4.1.0-alpha5​

Fix input props not working for dynamic metadata
Apply changes from 4.0.0-alpha20.

4.1.0-alpha4​
Fixed schema error when invoking a render.
Bug fixes leading to public testing.
IssueResolutionRendering a still via CLI with defaults results in error - You can only pass the quality option if imageFormat is 'jpeg'.Migrated to V4 method, using internalRenderStill() instead of renderStill(). Noticed missing options, added them in and documented.
4.1.0-alpha3​
Bug fixes leading to public testing.
IssueResolutionWhen deploying a service, the image didn't exist in Google Artifact Registry.Added publish script that runs submit.mjs, automatically deploying the image, tagged with the version number.Functions folder wasn't included in dist folder, so no CLI commandswould work.Removed this from .npmignore, so that it is included.When using the CLI to request a render without passing a composition name, it fails to list out compositions to choose fromIssue raised, present in V4 for Lambda also.Service name structuring clips off alpha version denominator. During alpha, this will make it impossible to deploy multiple services spanning alpha versions.Create new name formatting that meets requirements. Added tests for this.CLI commands for rendering not aligned with Remotion Lambda.npx remotion cloudrun render media is now npx remotion cloudrun render.npx remotion cloudrun render still is now npx cloudrun remotion still.Documentation also updated.
4.1.0-alpha2​
Initial cloud run alpha release 🎉.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousFunction naming conventionNextOverviewWhat to test1. Install @remotion/cloudrunChangelog (moved)4.0.184.1.0-alpha124.1.0-alpha114.1.0-alpha94.1.0-alpha54.1.0-alpha44.1.0-alpha34.1.0-alpha2