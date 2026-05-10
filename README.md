# ASCII-Art CLI & Reverse Tool

A comprehensive **Go** command-line utility for generating stylized ASCII art and reversing ASCII graphics back into plain text.

## Features

*   **ASCII Generation**: Convert any string into high-quality ASCII art.
*   **Reverse Mode**: Use the `--reverse=<fileName>` flag to decode an ASCII art file back into a normal string.
*   **Color Support**: Apply ANSI colors to specific substrings or the entire output using `--color=<colorName>`.
*   **Alignment**: Align your art within the terminal using `--align=<type>` (left, center, right, justify).
*   **File Export**: Save your generated ASCII art directly to a `.txt` file with the `--output=<fileName>` flag.
*   **Banner Support**: Supports multiple styles: `standard`, `shadow`, and `thinkertoy`.

## Installation

Ensure you have **Go** installed (version 1.16 or higher). Clone the repository and ensure the banner files (`.txt`) are in the root directory.

## Usage

### 1. Basic Generation
**bash**

```
go run . "Hello World" standard
```

### 2. Output to File

**bash**

```
go run . --output=result.txt "ASCII Art" shadow
```

### 3. Reverse Mode (Decoding)

- To convert an existing ASCII art file back to text:

**bash**

```
go run . --reverse=result.txt
```

### 4. Colors and Alignment

**bash**

```
go run . --color=red --align=center "Warning" thinkertoy
```

### Options & Flags
Flag	Description	Example

```
--reverse	Decodes an ASCII file to text	--reverse=file.txt
--output	Saves the output to a file	--output=art.txt
--color	Sets the text color	--color=blue
--align	Sets text alignment	--align=center
```

### How to run the tests

- Open your terminal in the project folder and type:

**bash**

```
go test -v
```

### Technical Implementation
- Pattern Matching: The Reverse algorithm scans input files column-by-column to match 8-line blocks against banner templates.
- Dynamic Width: Alignment is calculated using tput cols to detect the current terminal size.
- Color Mapping: Uses a custom mapping logic to identify and color specific substrings within the ASCII output.
### Requirements
**Go 1.16+**
### Banner files: 
- standard.txt, shadow.txt, thinkertoy.txt