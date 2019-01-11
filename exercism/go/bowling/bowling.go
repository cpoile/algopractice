package bowling

import (
	"errors"
)

type Game struct {
	frameNum int
	frame    []Frame
}

type Frame struct {
	rolls []int
	next  *Frame
}

func NewGame() *Game {
	frames := make([]Frame, 10)
	for i := 0; i < 9; i++ {
		frames[i].next = &frames[i+1]
	}

	return &Game{frame: frames}
}

func (g *Game) Roll(pins int) error {
	if g.frameNum == 10 {
		return errors.New("game is over")
	}

	curFrame := &g.frame[g.frameNum]
	curFrame.rolls = append(curFrame.rolls, pins)

	// check inputs
	if pins < 0 || pins > 10 {
		return errors.New("invalid role")
	}
	if g.frameNum < 9 && curFrame.sum() > 10 {
		return errors.New("invalid role")
	}
	if g.frameNum == 9 && len(curFrame.rolls) == 3 {
		if curFrame.rolls[0] == 10 && curFrame.rolls[1] != 10 {
			if curFrame.rolls[1]+curFrame.rolls[2] > 10 {
				return errors.New("invalid role")
			}
		}
	}

	if g.frameNum == 9 { // tenth frame
		if len(curFrame.rolls) == 3 {
			g.frameNum++
		}
		if len(curFrame.rolls) == 2 {
			if curFrame.rolls[0] == 10 || curFrame.rolls[0]+curFrame.rolls[1] == 10 {
				// extra frame
			} else {
				g.frameNum++
			}
		}
	} else {
		if curFrame.sum() == 10 || len(curFrame.rolls) == 2 {
			g.frameNum++
		}
	}

	return nil
}

func (g *Game) Score() (int, error) {
	if g.frameNum < 10 {
		return 0, errors.New("can't score, not finished")
	}
	var sum int
	for i := 0; i < 10; i++ {
		curFrame := g.frame[i]
		sum += curFrame.sum()
		if curFrame.sum() == 10 && len(curFrame.rolls) == 2 {
			sum += curFrame.next.nextXRolls(1)
		} else if curFrame.sum() == 10 && len(curFrame.rolls) == 1 {
			sum += curFrame.next.nextXRolls(2)
		}
	}
	return sum, nil
}

// get the next 1 or 2 rolls after the current. Handle 10th frame.
func (f *Frame) nextXRolls(x int) int {
	if f == nil || x == 0 {
		return 0
	}
	if len(f.rolls) == 1 {
		return f.rolls[0] + f.next.nextXRolls(x-1)
	} else {
		if x == 1 {
			return f.rolls[0]
		}
		return f.rolls[0] + f.rolls[1]
	}
}

func (f *Frame) sum() (res int) {
	for _, r := range f.rolls {
		res += r
	}
	return
}
