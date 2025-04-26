package letterboxed

import (
    "fmt"
    "encoding/json"
)



type Side map[rune]int

func (s Side) Solved() bool {
    for _, count := range s {
        if count < 1 {
            return false
        }
    }
    return true
}

func (s Side) Contains(c rune) bool {
    _, ok := s[c]
    return ok
}

func (s Side) Missing() []rune {
    // todo
    return nil
}

func (s *Side) Visit(char rune) {
    // todo
}

func (s Side) GetRunes() map[rune]int {
    // todo
    return nil
}

func (s *Side) UnmarshallJSON(data []byte) error {
    // TODO: Why isn't this unmarshaller working? It makes empty map every time
    if len(data) < 2 {
        return fmt.Errorf("Data bytes too few for valid json")
    }
    var sideMap map[string]int
    *s = make(Side)
    side := *s
    err := json.Unmarshal(data, &sideMap)
    if nil != err {
       return err 
    }
    for k, v := range sideMap { 
        kRune := []rune(k)[0]
        side[kRune] = v
    }
    return nil
}


