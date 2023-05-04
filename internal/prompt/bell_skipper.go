package prompt

import "os"

const bellCharacter = 7

type BellSkipperStdout struct{}

func NewBellSkipperStdout() *BellSkipperStdout {
	return &BellSkipperStdout{}
}

func (bs *BellSkipperStdout) Write(b []byte) (int, error) {
	if len(b) == 1 && b[0] == bellCharacter {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (bs *BellSkipperStdout) Close() error {
	return os.Stderr.Close()
}
