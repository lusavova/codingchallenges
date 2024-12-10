# Coding Challenges

This repository contains solutions to various challenges from [Coding Challenges](https://codingchallenges.fyi/).

## [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc)
A command-line tool thhat count lines, words, bytes and chas in a file or standard input.

### Build and Run The Project
**Prerequisites:**
- Make sure Go is installed.

1) Navigate to directory:
```
cd wc
```

2) Build the Executable:

```
go build -o ccwc
```
3) Run the Executable:

```
./ccwc -l -w test.txt
```
OR
```
  cat test.txt | ./ccwc
```
