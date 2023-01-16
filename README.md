# slice

Slice is a command-line tool that allows you to extract specific parts of input using pipes or any other way. It works by providing it with standard input (STDIN). The tool can be used to filter and output specific lines of input.

## Installation

```
go install github.com/nahuelrm/slice@latest
```

## Usage

```
<command> | slice `<options>`
```

## Options

The `:` operator can be used to specify a range of lines to output. Multiple filter options can be used separated by a comma, for example: `1,3:-7,-2`.

Examples of possible options:

| Option | Description |
| :--- | :--- |
| `5` | This option will print the 5th line.
| `7:` | This option will print all the content from the 7th line to the end.
| `:7` | This option will print all the content from initial line to the 7th line.
| `3:7` | This option will print all the content from the 3rd line to the 7th line.
| `4:-1` | This option will print all the content from the 4th line to the penultimate line.
| `4,8,-1` | This option combines 3 different options and will print the 4th, 8th and the penultimate line.
