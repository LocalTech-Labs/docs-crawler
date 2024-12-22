Building appsInstallation in existing projectOn this pageInstalling Remotion in an existing projectRemotion can be installed into existing projects, such as Next.JS, Remix, Vite or Create React App, as well as server-only projects that run on Node.JS. Get started by adding the following packages:
npmyarnpnpmbunnpm i --save-exact remotion@4.0.241 @remotion/cli@4.0.241Copynpm i --save-exact remotion@4.0.241 @remotion/cli@4.0.241Copypnpm i remotion@4.0.241 @remotion/cli@4.0.241Copypnpm i remotion@4.0.241 @remotion/cli@4.0.241Copybun i remotion@4.0.241 @remotion/cli@4.0.241Copybun i remotion@4.0.241 @remotion/cli@4.0.241Copyyarn --exact add remotion@4.0.241 @remotion/cli@4.0.241Copyyarn --exact add remotion@4.0.241 @remotion/cli@4.0.241CopyThis assumes you are currently using v4.0.241 of Remotion.Also update remotion and all `@remotion/*` packages to the same version.Remove all ^ character in front of the version numbers of it as it can lead to a version conflict.

If you'd like to embed a Remotion video in your existing React app, install @remotion/player as well.
If you'd like to render a video using the Node.JS APIs, install @remotion/renderer as well.
If you'd like to trigger a render on Remotion Lambda, install @remotion/lambda as well.

Setting up the folder structure​
Create a new folder for the Remotion files. It can be anywhere and assume any name, in this example we name it remotion and put it in the root of our project. Inside the folder you created, create 3 files:
remotion/Composition.tsxtsxexport const MyComposition = () => {  return null;};
remotion/Composition.tsxtsxexport const MyComposition = () => {  return null;};
remotion/Root.tsxtsximport React from 'react';import {Composition} from 'remotion';import {MyComposition} from './Composition'; export const RemotionRoot: React.FC = () => {  return (    <>      <Composition        id="Empty"        component={MyComposition}        durationInFrames={60}        fps={30}        width={1280}        height={720}      />    </>  );};
remotion/Root.tsxtsximport React from 'react';import {Composition} from 'remotion';import {MyComposition} from './Composition'; export const RemotionRoot: React.FC = () => {  return (    <>      <Composition        id="Empty"        component={MyComposition}        durationInFrames={60}        fps={30}        width={1280}        height={720}      />    </>  );};
remotion/index.tstsimport { registerRoot } from "remotion";import { RemotionRoot } from "./Root"; registerRoot(RemotionRoot);
remotion/index.tstsimport { registerRoot } from "remotion";import { RemotionRoot } from "./Root"; registerRoot(RemotionRoot);
The file that calls registerRoot() is now your Remotion entry point.
noteWatch out for import aliases in your tsconfig.json that will resolve import {...} from "remotion" to the remotion folder. We recommend to not use paths without a prefix.
Starting the Studio​
Start the Remotion Studio using the following command:
npx remotion studio remotion/index.ts
npx remotion studio remotion/index.ts
Replace remotion/index.ts with your entrypoint if necessary.
Render a video​
Render our a video using
npx remotion render remotion/index.ts MyComp out.mp4
npx remotion render remotion/index.ts MyComp out.mp4
Replace the entrypoint, composition name and output filename with the values that correspond to your usecase.
Install the ESLint Plugin​
Remotion has an ESLint plugin that warns about improper usage of Remotion APIs. To add it to your existing project, install it:
npmyarnpnpmbashnpm i @remotion/eslint-plugin
bashnpm i @remotion/eslint-pluginbashyarn add @remotion/eslint-plugin
bashyarn add @remotion/eslint-pluginbashpnpm i @remotion/eslint-plugin
bashpnpm i @remotion/eslint-plugin
This snippet will enable the recommended rules only for the Remotion files:
.eslintrcjson{  "plugins": ["@remotion"],  "overrides": [    {      "files": ["remotion/*.{ts,tsx}"],      "extends": ["plugin:@remotion/recommended"]    }  ]}
.eslintrcjson{  "plugins": ["@remotion"],  "overrides": [    {      "files": ["remotion/*.{ts,tsx}"],      "extends": ["plugin:@remotion/recommended"]    }  ]}
Embed a Remotion video into your React app​
You can use the <Player> component to display a Remotion video in your React project. Read the separate page about it for instructions.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousTelemetry in @remotion/webcodecsNextConvert Studio to an appSetting up the folder structureStarting the StudioRender a videoInstall the ESLint PluginEmbed a Remotion video into your React app