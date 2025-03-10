package letterboxed

import (
    "fmt"
    "github.com/nyt-solvers/puzzle"
)


type struct Side map[rune]int

/* todo: Move this interface to a more appropriate module
type Puzzle interface {
    CheckSolved() bool
    AddWord(string) 
    Solve()
    View() string
    VerifySolution() bool
    DisplaySolution() string
}
*/

type struct LetterBoxedPuzzle {
    sides []Side
    words []string
    solved []bool
}

func (p *Puzzle) CheckSolved() bool { 
    return false // todo 
}

func (p *Puzzle) AddWord(word string) {
    // todo
}

func (p *Puzzle) Solve() {
    // todo
}

func (p *Puzzle) View() string {
    // todo
    return ""
}

func (p *Puzzle) VerifySolution() {
    // todo
}

func (p *Puzzle) DisplaySolution() string {
    // todo
    return ""
}

func (p *puzzle) String() string {
    var repr string = "Letterboxed Puzzle:\n\tSides:"
    for i := 0; i < len(p.sides); i++ {
        repr += fmt.Sprintf("\t\t%v\n", p.sides[i])
    }
    repr += "\n\tSolution Words:"
    for j := 0; j < len(p.sides); j++ {
        repr += fmt.Sprintf("%s ", p.words[j])
    }
    repr += fmt.Sprintf("\n\tSolved: %v", p.solved)
}

func LetterboxedTest() {
    fmt.Println("Main for letterboxed module")
}

