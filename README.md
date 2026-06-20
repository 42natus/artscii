# artscii

A Go implementation of an ASCII art renderer. This command-line tool transforms standard strings into stylized ASCII banners based on pre-configured font templates. It features control over layout positioning in the terminal (alignment, centering, and word-level justification), substring coloring (with the use of ANSI escape codes for color), and output redirection to text files.

---

## Features

- **Font Banner Rendering:** Supports dynamic loading of banner layout files containing the ASCII art templates (i.e. ASCII art "fonts"). The ones included here are: `standard`, `shadow`, `thinkertoy`.

- **Terminal Alignment Engine:** Fully responsive string positioning matching your current terminal context width via system IO calls:
    * `left`: Standard layout baseline.
    * `center`: Perfectly balances whitespace margins.
    * `right`: Right-aligns content to the right terminal edge.
    * `justify`: A tokenization-based layout engine that filters out spacing elements in the generated ASCII art in order to distribute the terminal's free space as clean gaps between the actual word components.


- **Targeted Substring Colorization:** Embeds 24-bit ANSI color codes in the generated ASCII art render, around specified substring matches or, if no substring is provided, the entire output. Colour set includes: `white`, `silver`, `gray`, `black`, `red`, `maroon`, `orange`, `yellow`, `olive`, `lime`, `green`, `cyan`, `aqua`, `teal`, `blue`, `navy`, `magenta`, `fuchsia`, and `purple`.

- **File Export Pipeline:** Diverts the ASCII art render straight to external `.txt` files, rather than standard output. <!-- without terminal formatting bleed.?? -->

---

## Project Architecture

The workspace is organized into a main runner, a flag verification checker, the font templates, and an isolated layout library package (`art`):

```text
├── main.go               # Command-line interface parser and flag interceptor
├── valid.go              # Verification guardrails checking formatting flags
├── banners/              # Font template directory
│   ├── standard.txt      # Reference template map files
    ├── shadow.txt
    └── thinkertoy.txt
└── art/                  # Core ASCII Art Engineering Suite
    ├── align.go          # Layout mechanics (Tokenization & padding distribution)
    ├── color.go          # ANSI color embedding filters
    ├── display.go        # Horizontal component stitching engine
    ├── draw.go           # Text parser & font-map lookup logic
    ├── output.go         # Pipeline for outputting to external file
    ├── template.go       # Font template generation operations
    └── word.go           # Type and method definitions for Word matrix

```

> **Architectural Note:** This layout engine models letters as distinct, multi-line structural elements (`type Word [][]string`). Characters are then stitched together row-by-row on a virtual canvas before formatting or outputting.

---

## Getting Started

### Prerequisites

* Go (version 1.22 or higher recommended)

### Setup

1. Clone the repository to your local runtime space:
```bash
git clone https://github.com/42natus/artscii
cd artscii

```

2. Verify that the font files are correctly situated in the layout file path:
```bash
ls banners/
# Should display standard.txt, shadow.txt, and thinkertoy.txt

```



---

## Usage Guide

The program handles formatting through three (optional) command flags: `--color`, `--output`, and `--align`.

```bash
go run . [OPTION] [STRING] [BANNER]

```

### Command Flags Reference

| Command Flag | Argument Pattern | Operation |
| --- | --- | --- |
| `--color` | `--color=<colorName> <substring>` | Colors a specific substring or the entire text block if no token match (`<substring>`) is provided. |
| `--output` | `--output=<filename>.txt` | Diverts the rendered ASCII art output away from standard output (the terminal) directly to disk (in a `.txt` file). |
| `--align` | `--align=<left\|center\|right\|justify>` | Formats text positioning using active terminal width detection. |

---

## Usage Examples

### 1. Simple Rendering

Render standard fonts via natural terminal output:

```bash
go run . "Hello World" standard

```

### 2. Layout Positioning Alignment

Place layout content directly against the right-hand margin of your terminal screen:

```bash
go run . --align=right "Right Aligned" shadow

```

### 3. Word-Level Text Justification

Distribute text block margins cleanly across both terminal edges using custom tokenization system:

```bash
go run . --align=justify "This is justified text" thinkertoy

```

### 4. Target Specific Term Colors

Apply 24-bit color to a distinct segment within the string input:

```bash
go run . --color=cyan "cyan" "painting a specific cyan word"

```

### 5. Multi-Flag Pipelines

Combine formatting, layout alignment, and disk export targets into a single command pipeline:

```bash
go run . --color=yellow --align=center "Centered Output"

```

```bash
go run . --color=lime --align=justify --output=english.txt "Centred vs. Centered"

```

---

## Deep Dive: The Justification Engine

The highlight of this implementation is the justification pipeline found in `art/align.go`.

When working with terminal layout spaces, the raw ASCII art rendering caused messy positioning in the terminal layouts, largely because:

- ANSI escape codes (for colored output) interfere with how the terminal counts character width.

- It’s necessary to differentiate between whitespace that’s part of the ASCII art (literal space characters) and whitespace that’s truly “empty” in the layout, so that the remaining unoccupied width is distributed correctly.

To solve this, the engine runs a custom extraction pass before performing alignment calculations. Take an input, "Hello World", for example:

```text
[ Raw Character Stream ]   ──>  [ Extract & Normalize ]  ──>  [ Clean Words ]
  H-e-l-l-o-[ ]-W-o-r-l-d        (Strips structural spaces     [Hello]  [World]
                                            & escape codes)            │
                                                                       ▼
[ Terminal Output Canvas ] <──    [ Inject Padding ]     <──  [ Calculate Gaps ]
  Hello             World        (Distributes spacing)          Gaps Needed: 1

```

1. **Token Isolation:** It loops through incoming elements and discards empty space tokens, isolating pure content words into a normalized structure (`cleanWords`).
2. **True Width Analysis:** It strips out underlying ANSI escape sequences to compute the literal visual layout width of the characters.
3. **Distribution Math:** It calculates the remaining terminal width and uses a remainder-distribution algorithm to evenly space out text gaps without breaking your words apart.

---

## Extensibility & Customization

* **Adding Fonts:** Feel free to drop any standard 8-line break character template file directly into the `banners/` directory. Ensure new font templates match standard ASCII alignment offsets starting from space (`' '`).
* **Custom Colors:** Extend the color mapping configurations inside `art/color.go` by appending desired ANSI tags to the internal `COLORS` map object.
