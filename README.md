# wc
wc is a command-line application that reads a value from the standard input (STDIN) and display the word count

## Features
- Count by words
- Count by lines
- Count by bytes

## Usage
1. Compile the application
```bash

  git git@github.com:hayohtee/wc.git
  cd wc
  go build

```
2. Run wc command
```bash
  echo "This is my first command-line app" | ./wc
```
## Options
wc contains flags for performing different counts. By defaults, it count by words.
```bash
  
  -l // For performing line count
  -b // For performing bytes count

```
**NOTE:** If both -l and -b flags are specified, -b will take precedence.
