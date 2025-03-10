package letterboxed

import (
    "fmt"
    "github.com/nyt-solvers/puzzle"
)


type Side map[rune]int

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

type LetterBoxedPuzzle struct {
    sides []Side
    words []string
    solved bool
}

func (p *LetterBoxedPuzzle) CheckSolved() bool { 
    return false // todo 
}

func (p *LetterBoxedPuzzle) AddWord(word string) {
    // todo
}

func (p *LetterBoxedPuzzle) Solve() {
    // todo
}

func (p *LetterBoxedPuzzle) View() string {
    // todo
    return ""
}

func (p *LetterBoxedPuzzle) VerifySolution() bool {
    // todo
    return false
}

func (p *LetterBoxedPuzzle) DisplaySolution() string {
    // todo
    return ""
}

func (p *LetterBoxedPuzzle) String() string {
    var repr string = "Letterboxed Puzzle:\n\t- Sides:\n"
    for _, side := range p.sides {
        repr += fmt.Sprintf("\t\t+ %v\n", side)
    }
    repr += "\n\t- Solution Words: "
    for _, word := range p.words {
        repr += fmt.Sprintf("%s ", word)
    }
    repr += fmt.Sprintf("\n\t- Solved: %v", p.solved)
    return repr
}

func LetterboxedTest() {
    fmt.Println("~~~~~~~~~~~~~~~~~~~~ Letterboxed Module ~~~~~~~~~~~~~~~~~~~~")
    var sideA = Side{'a':0, 'b':0, 'c':0, 'd':0}
    var sideB = Side{'e':0, 'f':0, 'g':0, 'h':0}
    var sideC = Side{'i':0, 'j':0, 'k':0, 'l':0}
    var sideD = Side{'m':0, 'n':0, 'o':0, 'p':0}

    var puzz puzzle.Puzzle = &LetterBoxedPuzzle{ 
        []Side{sideA, sideB, sideC, sideD},
        []string{"foo", "bar",},
        false,
    }  
    fmt.Printf("%v\n", puzz)
}

