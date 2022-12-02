package rps

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Strategy struct {
	code     map[string]string
	scoring  map[string]int
	solution []byte
}

type ScoreConfig struct {
	Rock     int
	Paper    int
	Scissors int
	Loss     int
	Draw     int
	Win      int
}

const (
	rockType     = "rock"
	paperType    = "paper"
	scissorsType = "Scissors"

	lossOutcome = "loss"
	drawOutcome = "draw"
	winOutcome  = "win"
)

var DefaultScoreConfig = ScoreConfig{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
	Loss:     0,
	Draw:     3,
	Win:      6,
}

func NewStrategy(fileName string, sc ScoreConfig) (*Strategy, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return &Strategy{}, err
	}

	strategy := &Strategy{
		code:     make(map[string]string),
		solution: b,
		scoring:  make(map[string]int),
	}

	strategy.SetScoring(sc)

	return strategy, nil
}

func (s *Strategy) CalculateTotalScore() int {
	r := bytes.NewReader(s.solution)
	scanner := bufio.NewScanner(r)
	score := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		// Get an array of codes eg. [a,x]
		arr := strings.Split(l, " ")
		opponent := s.code[arr[0]]
		player := s.code[arr[1]]

		outcomeScore := s.scoring[checkForOutcome(opponent, player)]
		choiceScore := s.scoring[player]

		score = score + outcomeScore + choiceScore
	}

	return score
}

func (s *Strategy) CalculateTotalScoreWithSuggestedMove() int {
	r := bytes.NewReader(s.solution)
	scanner := bufio.NewScanner(r)
	score := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		// Get an array of codes eg. [a,x]
		arr := strings.Split(l, " ")
		opponent := s.code[arr[0]]
		outcome := s.code[arr[1]]

		outcomeScore := s.scoring[outcome]
		choiceScore := s.scoring[determineMoveForDesiredOutcome(opponent, outcome)]

		score = score + outcomeScore + choiceScore
	}

	return score
}

func (s *Strategy) SetScoring(sc ScoreConfig) {
	s.scoring = map[string]int{
		rockType:     sc.Rock,
		paperType:    sc.Paper,
		scissorsType: sc.Scissors,
		winOutcome:   sc.Win,
		lossOutcome:  sc.Loss,
		drawOutcome:  sc.Draw,
	}
}

func (s *Strategy) EncodeRock(code string) {
	s.code[code] = rockType
}

func (s *Strategy) EncodePaper(code string) {
	s.code[code] = paperType
}

func (s *Strategy) EncodeScissors(code string) {
	s.code[code] = scissorsType
}

func (s *Strategy) EncodeWin(code string) {
	s.code[code] = winOutcome
}

func (s *Strategy) EncodeLoss(code string) {
	s.code[code] = lossOutcome
}

func (s *Strategy) EncodeDraw(code string) {
	s.code[code] = drawOutcome
}

func (s *Strategy) SetDefaultEncoding() {
	s.EncodeRock("A")
	s.EncodePaper("B")
	s.EncodeScissors("C")
	s.EncodeRock("X")
	s.EncodePaper("Y")
	s.EncodeScissors("Z")
}

func (s *Strategy) SetDefaultEncodingWithExpectedResults() {
	s.EncodeRock("A")
	s.EncodePaper("B")
	s.EncodeScissors("C")
	s.EncodeLoss("X")
	s.EncodeDraw("Y")
	s.EncodeWin("Z")
}

func checkForOutcome(opponent, player string) string {
	winMap, _ := GetWinLossMaps()

	if opponent == player {
		return drawOutcome
	} else if player == winMap[opponent] {
		return winOutcome
	} else {
		return lossOutcome
	}
}

func determineMoveForDesiredOutcome(opponent, desiredOutcome string) string {

	winMap, lossMap := GetWinLossMaps()

	if desiredOutcome == drawOutcome {
		return opponent
	} else if desiredOutcome == winOutcome {
		return winMap[opponent]
	} else {
		return lossMap[opponent]
	}
}

func GetWinLossMaps() (winMap, lossMap map[string]string) {
	winMap = map[string]string{
		scissorsType: rockType,
		paperType:    scissorsType,
		rockType:     paperType,
	}

	lossMap = map[string]string{
		rockType:     scissorsType,
		scissorsType: paperType,
		paperType:    rockType,
	}

	return winMap, lossMap
}
