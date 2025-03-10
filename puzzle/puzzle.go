package puzzle

type Puzzle interface {
    CheckSolved() bool
    AddWord(string) 
    Solve()
    View() string
    VerifySolution() bool
    DisplaySolution() string
}

