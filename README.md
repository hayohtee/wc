# wc
wc is a command-line application that reads a value from the standard input (STDIN) and display the word count

## Features
- Count by words
- Count by lines
- Count by bytes

## Build from source
1. Clone the repository
  ```bash
  git clone git@github.com:hayohtee/wc.git
  ```
2. Change into the project directory
  ```bash
  cd wc
  ```
3. Compile
```bash
go build ./...
```
## Usage
Simply use pipe operator to sends the output of other process to wc\
Here is an example of using echo command.
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
