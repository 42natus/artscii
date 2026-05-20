# To-do for ascii-art-justify

1. make align work on color
2. figure out `align.Justify()`
3. order options around and observe output

## post-Coffee and Claude

1. ~~fix `'\n'` issue~~
2. cross-check position of checker for only `"\n"`s in input. (in `main.go`, lines 97 - 102)
3. display proper usage message for each case of invalid option usage
4. only output if filename follows <filename>.txt

## to implement

1. For each `[]Word` line, separate content words from space words — content at even indices, spaces at odd indices
2. Calculate `totalContentWidth` — sum of `word.Width()` for even-indexed words
3. Calculate `totalSpace = terminalWidth - totalContentWidth`
4. Calculate `numGaps = len(line) / 2` — number of space slots
5. Distribute space using `baseGap` and `remainder`
6. Build each output row by interleaving content lines with gap strings
