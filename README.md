# goetry

_poetry generation in golang_

The goal of this project is to generate standard poetic forms (like sonnets, haikus, limericks, etc.) based on a corpus.

It will use a pronunciation dictionary (http://www.speech.cs.cmu.edu/cgi-bin/cmudict) achieve appropriate rhyming and metrical feet.

## Status

No poetry - not yet!

There are basic commands for generating text based on an input corpus. Hopefully we'll get to poetry someday. Just not today.

## How to run

### Get a Corpus Ready

You'll need a basic corpus file in the following json format:

```
[
  {"Person": "Al", "Line": "Why am I soft in the middle?"},
  {"Person": "Al", "Line": "I can call you \"Betty;\" and, Betty, when you call me you can call me \"Al.\""}
]
```

### Run a command

#### List People
If your corpus has multiple "persons" in it, this command will list you all the persons in the corpus.

#### Generate Sentences
You can run the `generate-sentences` command to generate _n_ number of sentences.

Required: The corpus file
Optional: Specific person (if unspecified, uses all the text in the corpus)
Optional: Number of words (default 10)

_Warning! If your corpus is small, or the specific person has very little training data, this command could result in an infinite loop._

#### Generate Words
You can run the `generate-words` command to generate _n_ number of words.

Required: The corpus file
Optional: Specific person (if unspecified, uses all the text in the corpus)
Optional: Number of words (default 10)