package letterboxed

import (
	"fmt"
	//"path/filepath"
    "log"
    "os"
    "encoding/json"
	"github.com/nyt-solvers/puzzle"
)

type LetterBoxedPuzzle struct {
    // todo: Why is this not unmarshalling the array of sides?
	Sides  []Side
	Words  []string
	Solved bool
    Title string
}

func (p *LetterBoxedPuzzle) IsSolved() (bool, error) {
	if p.Solved {
		return true, nil
	}
	return p.VerifySolution()
}

// todo: Make this dude return a bool error pair.
func (p LetterBoxedPuzzle) validWord(word string) (bool, error) {
	// letter boxed puzzle requires words of length >= 3
	if len(word) < 3 {
		return false, puzzle.PuzzleError{fmt.Sprintf("Word %s < 3 characters", word)}
	}
	charsSeen := 0
	for i, char := range word {
		for _, side := range p.Sides {
			if !side.Contains(char) {
				continue
			}
			charsSeen += 1
			// next character is on the same side which is a violation
			next := i + 1
			if len(word) > next && side.Contains(rune(word[next])) {
				return false, puzzle.PuzzleError{fmt.Sprintf("Word %s contains sequence %U %U which are on the same side.", word, char, next)}
			}
			break
		}
	}
	return len(word) == charsSeen, puzzle.PuzzleError{fmt.Sprintf("Word %s contains at least one character not in any of the sides.", word)}
}

func (p *LetterBoxedPuzzle) AddWord(word string) {
	p.Words = append(p.Words, word)
}

func (p *LetterBoxedPuzzle) Solve() {
	// todo
}

func (p *LetterBoxedPuzzle) View() string {
	return fmt.Sprintf("@[%p] %s", p, p.String())
}

func (p *LetterBoxedPuzzle) VerifySolution() (bool, error) {
	for _, side := range p.Sides {
		if !side.Solved() {
			return false, nil
		}
	}
	for i, word := range p.Words {
		valid, err := p.validWord(word)
		if !valid {
			return false, err // todo: Errors here
		}
		if i < len(p.Words)-1 {
			if word[len(word)-1] != p.Words[i+1][0] {
				err := puzzle.PuzzleError{fmt.Sprintf("Word %s ends with a different character than %s begins with", word, p.Words[i+1])}
				return false, err
			}
		}
	}
	p.Solved = true
	return true, nil
}

func (p *LetterBoxedPuzzle) DisplaySolution() string {
	var solution string
	if solved, _ := p.IsSolved(); solved {
		solution = "Solution: ("
	} else {
		solution = "Partial Solution: ("
	}
	for i, word := range p.Words {
		if i < len(p.Words)-1 {
			solution += fmt.Sprintf("'%s', ", word)
		} else {
			solution += fmt.Sprintf("'%s'", word)
		}
	}
	return solution + ")"
}

func (p *LetterBoxedPuzzle) String() string {
	var repr string = "Letterboxed Puzzle:\n\t- Sides:\n"
	for _, side := range p.Sides {
		repr += fmt.Sprintf("\t\t+ %v\n", side)
	}
	repr += "\n\t- Solution Words: "
	for _, word := range p.Words {
		repr += fmt.Sprintf("%s ", word)
	}
	repr += fmt.Sprintf("\n\t- IsSolved: %v", p.IsSolved)
	return repr
}

func ReadLetterboxedPuzzle(filepath string) *LetterBoxedPuzzle {
    content, err := os.ReadFile(filepath)
    if nil != err {
        log.Fatal(err) 
    }
    var puzzle LetterBoxedPuzzle
    json.Unmarshal(content, &puzzle)
    return &puzzle
}

func LetterboxedTest() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~ Letterboxed Module ~~~~~~~~~~~~~~~~~~~~")
	// todo: Fin Side implementation
	var sideA = Side{'a': 1, 'b': 1, 'c': 1, 'd': 1}
	var sideB = Side{'e': 1, 'f': 1, 'g': 1, 'h': 1}
	var sideC = Side{'i': 1, 'j': 1, 'k': 1, 'l': 1}
	var sideD = Side{'m': 1, 'n': 1, 'o': 1, 'p': 1}

	var puzz puzzle.Puzzle = &LetterBoxedPuzzle{
		[]Side{sideA, sideB, sideC, sideD},
		make([]string, 0),
		false,
        "test",
	}
	var solution []string = []string{"afkd", "dminjp", "pebgclho"}
	for _, w := range solution {
		puzz.AddWord(w)
	}
	fmt.Printf("%v\n", puzz.View())
	solved, err := puzz.IsSolved()
	fmt.Printf("%t, %v\n", solved, err)
	fmt.Println(puzz.DisplaySolution())
    var puzzleTwo LetterBoxedPuzzle = *ReadLetterboxedPuzzle("resource/sample_problems/3-21-letterboxed.json")
    fmt.Printf("PuzzleTwo\n\n%v\n", puzzleTwo)
    content, err := os.ReadFile("resource/sample_problems/test-side.json")
    if nil != err {
        log.Fatal("Oh fuggg, ReadFile err=%v", err)
    }
    var s Side
    err = s.UnmarshallJSON(content)
    if nil != err {
        log.Fatal("Oh fuggg, err=%v", err)
    }
    fmt.Printf("s=%v\n", s)
}
