package request

import "io"

type Whoami struct{}

func (w *Whoami) Path() (string, error) {
	return "/actions/whoami", nil
}

func (r *Whoami) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *Whoami) ContentType() string {
	return ""
}
