package dik // This package contains DIK(DirectInputKey) scan-code constants taken from dinput.h

//#include "dinputkeys.h"
import "C"

// These are dikey values. Mapped to the scan codes.
const (
	KeyReleased uint8 = 0x0
	KeyPressed  uint8 = 0x80
)

// Listed are keyboard scan code constants.
const (
	Escape uint8 = iota + 1
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	Key0
	Minus
	Equals
	Back
	Tab
	Q
	W
	E
	R
	T
	Y
	U
	I
	O
	P
	LeftBracket
	RightBracket
	Return
	LeftControl
	A
	S
	D
	F
	G
	H
	J
	K
	L
	SemiColon
	Apostrophe
	Grave
	LeftShift
	BackSlash
	Z
	X
	C
	V
	B
	N
	M
	Comma
	Period
	Slash
	RightShift
	Multiply
	LeftMenu
	Space
	Capital
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	NumLock
	Scroll
	Numpad7
	Numpad8
	Numpad9
	NumPadSubtract
	Numpad4
	Numpad5
	Numpad6
	NumpadAdd
	Numpad1
	Numpad2
	Numpad3
	Numpad0
	NumpadDecimal
)

// Arrow keys
const (
	Up    uint8 = 0xC8
	Down  uint8 = 0xD0
	Left  uint8 = 0xCB
	Right uint8 = 0xCD
)

// #define DIK_F11 0x57
// #define DIK_F12 0x58

// #define DIK_F13 0x64 /* (NEC PC98) */
// #define DIK_F14 0x65 /* (NEC PC98) */
// #define DIK_F15 0x66 /* (NEC PC98) */

// #define DIK_KANA 0x70 /* (Japanese keyboard) */
// #define DIK_CONVERT 0x79 /* (Japanese keyboard) */
// #define DIK_NOCONVERT 0x7B /* (Japanese keyboard) */
// #define DIK_YEN 0x7D /* (Japanese keyboard) */
// #define DIK_NUMPADEQUALS 0x8D /* = on numeric keypad (NEC PC98) */
// #define DIK_CIRCUMFLEX 0x90 /* (Japanese keyboard) */
// #define DIK_AT 0x91 /* (NEC PC98) */
// #define DIK_COLON 0x92 /* (NEC PC98) */
// #define DIK_UNDERLINE 0x93 /* (NEC PC98) */
// #define DIK_KANJI 0x94 /* (Japanese keyboard) */
// #define DIK_STOP 0x95 /* (NEC PC98) */
// #define DIK_AX 0x96 /* (Japan AX) */
// #define DIK_UNLABELED 0x97 /* (J3100) */
// #define DIK_NUMPADENTER 0x9C /* Enter on numeric keypad */
// #define DIK_RCONTROL 0x9D
// #define DIK_NUMPADCOMMA 0xB3 /* , on numeric keypad (NEC PC98) */
// #define DIK_DIVIDE 0xB5 /* / on numeric keypad */
// #define DIK_SYSRQ 0xB7
// #define DIK_RMENU 0xB8 /* right Alt */
// #define DIK_HOME 0xC7 /* Home on arrow keypad */
// #define DIK_UP 0xC8 /* UpArrow on arrow keypad */
// #define DIK_PRIOR 0xC9 /* PgUp on arrow keypad */
// #define DIK_LEFT 0xCB /* LeftArrow on arrow keypad */
// #define DIK_RIGHT 0xCD /* RightArrow on arrow keypad */
// #define DIK_END 0xCF /* End on arrow keypad */
// #define DIK_DOWN 0xD0 /* DownArrow on arrow keypad */
// #define DIK_NEXT 0xD1 /* PgDn on arrow keypad */
// #define DIK_INSERT 0xD2 /* Insert on arrow keypad */
// #define DIK_DELETE 0xD3 /* Delete on arrow keypad */
// #define DIK_LWIN 0xDB /* Left Windows key */
// #define DIK_RWIN 0xDC /* Right Windows key */
// #define DIK_APPS 0xDD /* AppMenu key */
