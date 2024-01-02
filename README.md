# digester-bot
`digester-bot` is a telegram bot designed to help you study a particular subject. You can send a file or link to the bot and then ask him to do different things to help you learn about a subject.
## Features / modes
### Summary
`digester-bot` can make summaries to help you understand things. In order to do that, you just need to tell them which language, max amount of words and name some topic (like a chapter, title, etc) so that the bot can focus on that when he does the summary.
### Main concepts
Similar to summary, the bot can make a list of most important concepts. This can be a good alternative if you want a less verbose output. The parameters are pretty much the same as the ones for the summary, but instead of the max amount of words you provide the max amount of items in the list.
### Question
Of course you can ask a question about the content that you provided. No other parameters other than the question itself are needed.
### Test
This option can be very useful if you want to review the topic and test your knowledge about it. Basically the bot will ask you a question about a specific topic (this is the only parameter that you need to provide), then you answer that question and finally the bot will tell you if you're right or wrong, and provide an explanation.

# Try it out!
You can try `digester-bot` by using [this link](https://web.telegram.org/k/#@DigestAIBot).

# How does it work?
Internally, it is integrated with [GPT Assistants](https://platform.openai.com/docs/assistants/overview]). One assistant is created every time a user sends a content (either a link or a file) and then all the prompts will be made against that assistant. 

## How to run it
In order to run it you just need to set the env variables `CHAT_GPT_KEY` with your OpenAI key and `TBOT_TOKEN` with your telegram token. `digester-bot` does not persist anything, so all the data is stored in memory, nothing in files/db. This makes it not very reliable but it's okay for now, I just wanted to focus on the features first and improve the technical aspects later.

