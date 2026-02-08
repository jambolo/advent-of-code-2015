# advent-of-code-2015

[![Build Status](https://github.com/jambolo/advent-of-code-2015/actions/workflows/validation-go.yml/badge.svg)](https://github.com/jambolo/advent-of-code-2015/actions/workflows/validation-go.yml)

My solutions to Advent Of Code 2015 implemented in Go. The development environment is VS Code and WSL Ubuntu.

## Day 1

Trivial.

| Part | Answer |
|------|--------|
|    1 |    280 |
|    2 |   1797 |

I am using Claude to help me write code and learn Go, but I have to be careful because it knows about Advent Of Code 2015 and it will write a complete solution if I let it.

## Day 2

Trivial.

| Part |  Answer |
|------|---------|
|    1 | 1588178 |
|    2 | 3783758 |

## Day 3

Trivial.

| Part | Answer |
|------|--------|
|    1 |   2565 |
|    2 |   2639 |

## Day 4

Trivial. Go has a simple MD5 library.

| Part |  Answer |
|------|---------|
|    1 |  346386 |
|    2 | 9958218 |

## Day 5

Trivial. I wanted to use regex, but it turned out to be a lot easier to do the checking manually. I'm letting Claude to most of the work here, because the solutions are trivial, so I'm being lazy.

| Part | Answer |
|------|--------|
|    1 |    238 |
|    2 |     69 |

## Day 6

Trivial. Go is starting to look to me like just a better version of C. I haven't seen any really interesting syntax or capabilities. In fact, it feels anemic compared to other modern languages.

| Part |  Answer  |
|------|----------|
|    1 |   543903 |
|    2 | 14687245 |

## Day 7

Fun -- some recursion, some regexes, some caching. Not too hard to figure out.

| Part | Answer |
|------|--------|
|    1 |  46065 |
|    2 |  14134 |

## Day 8

Trivial.

| Part | Answer |
|------|--------|
|    1 |   1371 |
|    2 |   2117 |

## Day 9

So disappointing... I was expecting some kind of combinatorial explosion with this traveling salesman problem, so I spent time preparing for the inevitable optimizations that would have to be implemented. Part 2 let me down. None of that was necessary, so I removed it.

| Part | Answer |
|------|--------|
|    1 |    141 |
|    2 |    736 |

## Day 10

Well, part 1 is pretty easy. I was lazy and just used the same code for part 2 and waited a long time for the result. I thought about caching substrings, but that seemed complicated and would probably take longer to implement than just waiting for the brute force result. After refactoring, the results were much faster. It turns out that the slowness was caused by reallocating the entire string after each character was appended.

| Part |  Answer |
|------|---------|
|    1 |  329356 |
|    2 | 4666278 |

## Day 11

Trivial. Not much opportunity to be clever.

| Part |  Answer  |
|------|----------|
|    1 | cqjxxyzz |
|    2 | cqkaabcc |

## Day 12

Go's JSON library made it simple.

| Part | Answer |
|------|--------|
|    1 | 191164 |
|    2 |  87842 |

## Day 13

Trivial. The permutation generator I made for a previous day came in very handy. These puzzles have been very simple so far. AI (Github Copilot, specifically) has been very handy doing 80% of the typing for me. I learned about the range keyword with loops. That's nice.

| Part | Answer |
|------|--------|
|    1 |    709 |
|    2 |    668 |

## Day 14

Trivial, except that I misread the description and overlooked the input data.

| Part | Answer |
|------|--------|
|    1 |   2660 |
|    2 |   1256 |

## Day 15

Trivial. Claude and Copilot have been very helpful pointing out ways to make my code more idiomatic.

| Part |  Answer  |
|------|----------|
|    1 | 13882464 |
|    2 | 11171160 |

## Day 16

Trivial.

| Part | Answer |
|------|--------|
|    1 |     40 |
|    2 |    241 |
