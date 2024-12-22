Upgrading RemotionOn this pageUpgrading RemotionThe easiest way to do this is to run the following command in the root of your project:
npxpnpmyarnbunxbashnpx remotion upgrade
bashnpx remotion upgradebashpnpm exec remotion upgrade
bashpnpm exec remotion upgradebashyarn remotion upgrade
bashyarn remotion upgradebashbunx remotion upgrade
bashbunx remotion upgrade
noteYou need the @remotion/cli package installed for this.
Manually upgrading Remotion​
To upgrade Remotion to the latest version, all @remotion/* packages and remotion must be updated to the same version.
Edit the version number in your package.json for all packages.
Delete the ^ in front of the version number in your package.json in order to force the exact version you specified.
Breaking changes​
Remotion follows semantic versioning.
This means if the first number of the version is the same, you can upgrade and your code is backwards-compatible.
Example: If you are on 4.0.0, you can upgrade to 4.1.100 without changing your code.
However, to upgrade to 5.0, you will need to follow the migration guide.
Exceptions to the breaking change rule are APIs that are marked as experimental.
Changelog​
Visit remotion.dev/changelog to see a list of all changes.
Stable versions​
We maintain a repo with the latest stable version of Remotion for customers who need a higher level of stability.
Customers may get access to the repo on the remotion.pro portal.Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousCombining compositionsNextTerminologyManually upgrading RemotionBreaking changesChangelogStable versions