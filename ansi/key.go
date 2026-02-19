package ansi

// Input ANSI key
type Key = Input

const (
	KeyCtrlTilde       Key = 0x00
	KeyCtrl2           Key = 0x00
	KeyCtrlSpace       Key = 0x00
	KeyCtrlA           Key = 0x01
	KeyCtrlB           Key = 0x02
	KeyCtrlC           Key = 0x03
	KeyCtrlD           Key = 0x04
	KeyCtrlE           Key = 0x05
	KeyCtrlF           Key = 0x06
	KeyCtrlG           Key = 0x07
	KeyCtrlH           Key = 0x08
	KeyBackspace       Key = 0x08
	KeyCtrlI           Key = 0x09
	KeyTab             Key = 0x09
	KeyCtrlJ           Key = 0x0A
	KeyCtrlK           Key = 0x0B
	KeyCtrlL           Key = 0x0C
	KeyCtrlM           Key = 0x0D
	KeyEnter           Key = 0x0D
	KeyCtrlN           Key = 0x0E
	KeyCtrlO           Key = 0x0F
	KeyCtrlP           Key = 0x10
	KeyCtrlQ           Key = 0x11
	KeyCtrlR           Key = 0x12
	KeyCtrlS           Key = 0x13
	KeyCtrlT           Key = 0x14
	KeyCtrlU           Key = 0x15
	KeyCtrlV           Key = 0x16
	KeyCtrlW           Key = 0x17
	KeyCtrlBackspace   Key = 0x17
	KeyCtrlX           Key = 0x18
	KeyCtrlY           Key = 0x19
	KeyCtrlZ           Key = 0x20
	KeyEsc             Key = 0x1B
	KeyCtrlLsqBracket  Key = 0x1B
	KeyCtrl3           Key = 0x1B
	KeyCtrl4           Key = 0x1C
	KeyCtrlBackslash   Key = 0x1C
	KeyCtrl5           Key = 0x1D
	KeyCtrlRsqBracket  Key = 0x1D
	KeyCtrl6           Key = 0x1E
	KeyCtrl7           Key = 0x1F
	KeyCtrlSlash       Key = 0x1F
	KeyCtrlUnderscore  Key = 0x1F
	KeySpace           Key = 0x20
	KeyBackspace2      Key = 0x7F
	KeyCtrl8           Key = 0x7F
	KeyAltA            Key = 0x1B61
	KeyAltB            Key = 0x1B62
	KeyAltC            Key = 0x1B63
	KeyAltD            Key = 0x1B64
	KeyAltE            Key = 0x1B65
	KeyAltF            Key = 0x1B66
	KeyAltG            Key = 0x1B67
	KeyAltH            Key = 0x1B68
	KeyAltI            Key = 0x1B69
	KeyAltJ            Key = 0x1B6A
	KeyAltK            Key = 0x1B6B
	KeyAltL            Key = 0x1B6C
	KeyAltM            Key = 0x1B6D
	KeyAltN            Key = 0x1B6E
	KeyAltO            Key = 0x1B6F
	KeyAltP            Key = 0x1B70
	KeyAltQ            Key = 0x1B71
	KeyAltR            Key = 0x1B72
	KeyAltS            Key = 0x1B73
	KeyAltT            Key = 0x1B74
	KeyAltU            Key = 0x1B75
	KeyAltV            Key = 0x1B76
	KeyAltW            Key = 0x1B77
	KeyAltX            Key = 0x1B78
	KeyAltY            Key = 0x1B79
	KeyAltZ            Key = 0x1B7a
	KeyScrollUp        Key = 0x1B4F41
	KeyScrollDown      Key = 0x1B4F42
	KeyF1              Key = 0x1B4F50
	KeyF2              Key = 0x1B4F51
	KeyF3              Key = 0x1B4F52
	KeyF4              Key = 0x1B4F53
	KeyF5              Key = 0x1B5B31357E
	KeyCtrlF5          Key = 0x1B5B31353B357E
	KeyF6              Key = 0x1B5B31377E
	KeyCtrlF6          Key = 0x1B5B31373B357E
	KeyF7              Key = 0x1B5B31387E
	KeyF8              Key = 0x1B5B31397E
	KeyCtrlF8          Key = 0x1B5B31393B357E
	KeyF9              Key = 0x1B5B32307E
	KeyF10             Key = 0x1B5B32317E
	KeyCtrlF11         Key = 0x1B5B32333B357E
	KeyF12             Key = 0x1B5B32347E
	KeyInsert          Key = 0x1B5B327E
	KeyCtrlInsert      Key = 0x1B5B323B357E
	KeyDelete          Key = 0x1B5B337E
	KeyCtrlDelete      Key = 0x1B5B333B357E
	KeyHome            Key = 0x1B5B48
	KeyCtrlHome        Key = 0x1B5B313B3548
	KeyEnd             Key = 0x1B5B46
	KeyCtrlEnd         Key = 0x1B5B313B3546
	KeyPgup            Key = 0x1B5B357E
	KeyCtrlPgup        Key = 0x1B5B353B357E
	KeyPgdn            Key = 0x1B5B367E
	KeyCtrlPgdn        Key = 0x1B5B363B357E
	KeyArrowUp         Key = 0x1B5B41
	KeyShiftArrowUp    Key = 0x1B5B313B3241
	KeyAltArrowUp      Key = 0x1B5B313B3341
	KeyCtrlArrowUp     Key = 0x1B5B313B3541
	KeyArrowDown       Key = 0x1B5B42
	KeyShiftArrowDown  Key = 0x1B5B313B3242
	KeyAltArrowDown    Key = 0x1B5B313B3342
	KeyCtrlArrowDown   Key = 0x1B5B313B3542
	KeyArrowLeft       Key = 0x1B5B44
	KeyShiftArrowLeft  Key = 0x1B5B313B3244
	KeyAltArrowLeft    Key = 0x1B5B313B3344
	KeyCtrlArrowLeft   Key = 0x1B5B313B3544
	KeyArrowRight      Key = 0x1B5B43
	KeyShiftArrowRight Key = 0x1B5B313B3243
	KeyAltArrowRight   Key = 0x1B5B313B3343
	KeyCtrlArrowRight  Key = 0x1B5B313B3543
)

// Creates a key from the rune.
func KeyFromRune(r rune) Key {
	return Key(r)
}

func isKnownKeyEscape(k Input) bool {
	switch k {
	case
		KeyAltA,
		KeyAltB,
		KeyAltC,
		KeyAltD,
		KeyAltE,
		KeyAltF,
		KeyAltG,
		KeyAltH,
		KeyAltI,
		KeyAltJ,
		KeyAltK,
		KeyAltL,
		KeyAltM,
		KeyAltN,
		KeyAltO,
		KeyAltP,
		KeyAltQ,
		KeyAltR,
		KeyAltS,
		KeyAltT,
		KeyAltU,
		KeyAltV,
		KeyAltW,
		KeyAltX,
		KeyAltY,
		KeyAltZ,
		KeyScrollUp,
		KeyScrollDown,
		KeyF2,
		KeyF3,
		KeyF4,
		KeyF5,
		KeyCtrlF5,
		KeyF6,
		KeyCtrlF6,
		KeyF7,
		KeyF8,
		KeyCtrlF8,
		KeyF9,
		KeyF10,
		KeyCtrlF11,
		KeyF12,
		KeyInsert,
		KeyCtrlInsert,
		KeyDelete,
		KeyCtrlDelete,
		KeyHome,
		KeyCtrlHome,
		KeyEnd,
		KeyCtrlEnd,
		KeyPgup,
		KeyCtrlPgup,
		KeyPgdn,
		KeyCtrlPgdn,
		KeyArrowUp,
		KeyShiftArrowUp,
		KeyAltArrowUp,
		KeyCtrlArrowUp,
		KeyArrowDown,
		KeyShiftArrowDown,
		KeyAltArrowDown,
		KeyCtrlArrowDown,
		KeyArrowLeft,
		KeyShiftArrowLeft,
		KeyAltArrowLeft,
		KeyCtrlArrowLeft,
		KeyArrowRight,
		KeyShiftArrowRight,
		KeyAltArrowRight,
		KeyCtrlArrowRight:
		return true
	default:
		return false
	}
}
