# openai

This is a simple cli that uses the [OpenAI Chat API](https://platform.openai.com/docs/api-reference/introduction) to generate content.

## How to use it?

Build the command:

```sh
bin/build
```

And run it:

```sh
Usage: openai --content <code|quiz>
  -content string
        Define the kind of content to process. (default "code")
```

The cli works with two kind of contents:

1. [Code](https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences)
2. [Quiz](https://www.geeksforgeeks.org/quizzes/machine-learning-quiz-questions-and-answers/)

**Note**: only tested on Linux.
