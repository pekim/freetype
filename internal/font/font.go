package font

import _ "embed"

//go:embed DejaVuSans/DejaVuSans.ttf
var DejaVuSans []byte

//go:embed DejaVuSans/DejaVuSansMono.ttf
var DejaVuSansMono []byte

//go:embed Noto_Color_Emoji/NotoColorEmoji-Regular.ttf
var NotoColorEmoji []byte
