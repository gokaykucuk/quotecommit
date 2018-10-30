# quotecommit

[![Maintainability](https://api.codeclimate.com/v1/badges/5387272a1509c322c196/maintainability)](https://codeclimate.com/github/gokaykucuk/quotecommit/maintainability)
[![Build Status](https://travis-ci.org/gokaykucuk/quotecommit.svg?branch=master)](https://travis-ci.org/gokaykucuk/quotecommit)
![license](https://img.shields.io/github/license/mashape/apistatus.svg)

quotecommit is a tool that i use from time to time while i shouldn't use it at all. I mean seriously, I am not even sure why i created this abomination. It simply breaks a fundemental system in Git just so I can be lazy for .4 seconds. 

Anyways, It just puts a random quote about computers into commit message and commits the current git repository.

# Building

On this project, I've experimented using go runtime to do some precompilation. Because of this process, you might need
 to either put `PROJECT_URI` (default: `github.com/gokaykucuk/quotecommit`) to your path, or directly go and modify the code in place in `main.go / projectFolder()` function. Actually check the func

# Installation

Just run `sudo make install` after running `make build`, don't trust the binary in repo. It's better if you build it.

# Usage

Just navigate to directory and use rhe command `gitqq`. If you also want to push what you've commited please add p as first argument.

Ex: `gitqq p` 

# Interesting Note

My ide, ( Jetbrains ), completely stopped responding to me, so with the commit this note included, i am actually going to use quotecommit to commit. You will see quote_placeholder during this process as commit message until it starts pulling the names.
 