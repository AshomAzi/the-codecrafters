// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Squad Name]
// ───────────────────────────────────────────
// Input line types:
//   [list what your squad agreed on]
//
// Transformation rules (in order):
//   1. [Rule 1]
//   2. [Rule 2]
//   3. [Rule 3]
//   4. [Rule 4]
//   5. [Rule 5]
//
// Output format:
//   [Header: yes/no — exact text if yes]
//   [Line numbering format]
//   [Summary block: yes/no — fields if yes]
//
// Terminal summary fields:
//   [List what your squad agreed on]
// ═══════════════════════════════════════════

// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Azi Bulua]
// Squad:  [Pointers]

package main

import (
	"bufio"
	transformation "file-pipeline/Transformation"

	// "file-pipeline/Transformation"

	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	f1, err := os.Open("input.txt")
	check(err)
	defer f1.Close()

	f2, err := os.Create("output.txt")
	check(err)
	defer f2.Close()

	text := bufio.NewScanner(f1)
	writer := bufio.NewWriter(f2)
	for text.Scan() {
		line := text.Text()

		new := transformation.Title(line)

		new = transformation.Upper(new)


		_, err := writer.WriteString(new + "\n")
		check(err)
	}
	writer.Flush()
}
