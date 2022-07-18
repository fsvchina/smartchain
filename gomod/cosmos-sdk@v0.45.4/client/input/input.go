package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bgentry/speakeasy"
	isatty "github.com/mattn/go-isatty"
)


const MinPassLength = 8



func GetPassword(prompt string, buf *bufio.Reader) (pass string, err error) {
	if inputIsTty() {
		pass, err = speakeasy.FAsk(os.Stderr, prompt)
	} else {
		pass, err = readLineFromBuf(buf)
	}

	if err != nil {
		return "", err
	}

	if len(pass) < MinPassLength {


		return pass, fmt.Errorf("password must be at least %d characters", MinPassLength)
	}

	return pass, nil
}




func GetConfirmation(prompt string, r *bufio.Reader, w io.Writer) (bool, error) {
	if inputIsTty() {
		fmt.Fprintf(w, "%s [y/N]: ", prompt)
	}

	response, err := readLineFromBuf(r)
	if err != nil {
		return false, err
	}

	response = strings.TrimSpace(response)
	if len(response) == 0 {
		return false, nil
	}

	response = strings.ToLower(response)
	if response[0] == 'y' {
		return true, nil
	}

	return false, nil
}


func GetString(prompt string, buf *bufio.Reader) (string, error) {
	if inputIsTty() && prompt != "" {
		fmt.Fprintf(os.Stderr, "> %s\n", prompt)
	}

	out, err := readLineFromBuf(buf)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out), nil
}




func inputIsTty() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}




func readLineFromBuf(buf *bufio.Reader) (string, error) {
	pass, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(pass), nil
}
