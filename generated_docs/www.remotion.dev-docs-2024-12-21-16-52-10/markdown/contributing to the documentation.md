ContributingWriting docsOn this pageContributing to the documentationGood documentation makes the difference between frustration and joy! We happily accept contributions to the Remotion documentation.
Editing an existing page​
At the bottom of each page, the Improve this page button is the easiest way to edit a page. You don't even need to setup the repository locally.
Submitting a new page​
Basic setup (quicker)Full docsThis method will avoid having to compile all docs. You cannot validate if links to other docs are valid.1  Set up the Remotion repository according to the instructions here. 
2  Run node new-doc.mjs and follow the instructions.1  Set up the Remotion repository according the instructions here. 
2  Create a new .md document in the packages/docs/docs folder. 
3  Add the document to packages/docs/sidebars.js.
4  Write what you have to say!
5  Run node render-cards.mjs in packages/docs to generate preview cards that will show up if the documentation page is shared on social media.
Language guidelines​

Keep it brief: Developers don't like to read, so adding too much words will lead to information being lost.
Link to terminology: When using a Remotion-specific term, link to the terminology page that explains it.
Avoid emotions and filler language ("Great! Let's move on to the next step"). Usually it can be removed without losing any information.
Separate into multiple paragraphs if a section is becoming too long.
Address the reader as "you" and not "we".
Don't blame the user: Instead of "You have provided the wrong input", say "The input is invalid" or "the wrong input was provided".
Don't assume it is easy: Avoid using the words "simply" and "just" as beginners might not find it as simple as you do.

Adding code snippets​
You can add codesnippets by beginning a paragraph with three backticks: ```. The code will be highlighted according to the language you specify after the backticks.
```ts will highlight the code as TypeScript.
Type safe snippets​
```ts twoslash will make a snippet type-safe - it will be checked against the TypeScript compiler. This is the preferred way of writing docs, but if it is too hard, you don't have to do it.
When writing typesafe snippets, sometimes it does not make sense to list all import statements at the top.You can add a line stating // ---cut--- and only the content below will be displayed.
Titles​
Add a title to a code snippet by adding a line with ```ts twoslash title="file.ts":
file.tstsconsole.log("Hello World");
file.tstsconsole.log("Hello World");
Special formatting​
Steps​
Use <Step> to create lists:
md- <Step>1</Step> Step 1- <Step>2</Step> Step 2
md- <Step>1</Step> Step 1- <Step>2</Step> Step 2

1  Step 1
2  Step 2

Mark as experimental​
Use <ExperimentalBadge /> to mark something as experimental:
md<ExperimentalBadge><p>This feature is still experimental.</p></ExperimentalBadge>
md<ExperimentalBadge><p>This feature is still experimental.</p></ExperimentalBadge>
EXPERIMENTALThis feature is still experimental.
Demos​
Using <Demo type="[demo-name]" /> you can render a Remotion Player and specify props that can be updated.
md<Demo type="rect"/>
md<Demo type="rect"/>
width200height200cornerRadius0edgeRoundnessdebug
The demo needs to be implemented in packages/docs/components/demos/index.tsx.
See also​

General information
Implementing a new feature
Implementing a new option
How to take a bounty issue
Improve this pageAsk on DiscordGet helpLast updated on Dec 21, 2024PreviousImplementing a new optionNextFormattingEditing an existing pageSubmitting a new pageLanguage guidelinesAdding code snippetsType safe snippetsTitlesSpecial formattingStepsMark as experimentalDemosSee also