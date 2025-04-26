package puzzle

type Puzzle interface {
    IsSolved() (bool, error)
    AddWord(string) 
    Solve()
    View() string
    VerifySolution() (bool, error)
    DisplaySolution() string
}

type PuzzleError struct {
    Cause string
}


func (err PuzzleError) Error() string {
    return err.Cause
}

type PuzzleReader interface {
    Read(string) Puzzle
}
