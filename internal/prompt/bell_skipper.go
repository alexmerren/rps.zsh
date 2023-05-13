package prompt

import (
	"os"
)

const bellCharacter = 7

type BellSkipperStdout struct{}

func NewBellSkipperStdout() *BellSkipperStdout {
	return &BellSkipperStdout{}
}

func (bs *BellSkipperStdout) Write(bytes []byte) (int, error) {
	if len(bytes) == 1 && bytes[0] == bellCharacter {
		return 0, nil
	}

	//nolint:wrapcheck // This breaks the bellskipper if the error is wrapped.
	return os.Stderr.Write(bytes)
}

func (bs *BellSkipperStdout) Close() error {
	//nolint:wrapcheck // This breaks the bellskipper if the error is wrapped.
	return os.Stderr.Close()
}
