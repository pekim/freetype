package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
)

type Error struct {
	ftError C.FT_Error
	message string
}

func newError(err C.FT_Error, message string) error {
	if err == C.FT_Err_Ok {
		return nil
	}
	return Error{
		ftError: err,
		message: message,
	}
}

func (err Error) Error() string {
	return fmt.Sprintf("%s : %d", err.message, err.ftError)
}
