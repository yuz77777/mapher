package main

import (
	"math/rand"
	"strings"
	"time"
)

func Encryption(s string) string {
	sl := strings.Split(s, "")
	len := len(sl)
	var cipher string
	var idNumber int
	var keyStr string

	stoi := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"1": 27,
		"2": 28,
		"3": 29,
		"4": 30,
		"5": 31,
		"6": 32,
		"7": 33,
		"8": 34,
		"9": 35,
		"0": 36,
		"@": 37,
		"K": 38,
		"#": 39,
		"%": 40,
		"&": 41,
		"[": 42,
		"]": 43,
		"(": 44,
		")": 45,
		"?": 46,
		"!": 47,
		"-": 48,
		"+": 49,
		"=": 50,
		"~": 51,
		",": 52,
		".": 53,
		" ": 54,
		"A": 55,
		"B": 56,
		"C": 57,
		"D": 58,
		"E": 59,
		"F": 60,
		"G": 61,
		"H": 62,
		"I": 63,
		"J": 64,
		"L": 65,
		"M": 66,
		"N": 67,
		"O": 68,
		"P": 69,
		"Q": 70,
		"R": 71,
		"S": 72,
		"T": 73,
		"U": 74,
		"V": 75,
		"W": 76,
		"X": 77,
		"Y": 78,
		"Z": 79,
	}
	itos := map[int]string{
		0:  "å",
		1:  "∫",
		2:  "ç",
		3:  "∂",
		4:  "´",
		5:  "ƒ",
		6:  "©",
		7:  "˙",
		8:  "ˆ",
		9:  "∆",
		10: "˚",
		11: "¬",
		12: "µ",
		13: "Ø",
		14: "ø",
		15: "π",
		16: "œ",
		17: "®",
		18: "ß",
		19: "†",
		20: "¨",
		21: "√",
		22: "∑",
		23: "≈",
		24: `\`,
		25: "Ω",
		26: "¡",
		27: "™",
		28: "£",
		29: "¢",
		30: "∞",
		31: "§",
		32: "¶",
		33: "•",
		34: "ª",
		35: "º",
		36: "€",
		37: "",
		38: "‹",
		39: "ﬁ",
		40: "‡",
		41: "”",
		42: "’",
		43: "·",
		44: "‚",
		45: "¿",
		46: "⁄",
		47: "—",
		48: "±",
		49: "≠",
		50: "`",
		51: "≤",
		52: "≥",
		53: "◊",
		54: "♠",
		55: "♣",
		56: "♥",
		57: "♦",
		58: "∝",
		59: "ℜ",
		60: "ν",
		61: "ξ",
		62: "δ",
		63: "ψ",
		64: "ℵ",
		65: "ℑ",
		66: "Λ",
		67: "Ψ",
		68: "κ",
		69: "℘",
		70: "≅",
		71: "∈",
		72: "Ε",
		73: "÷",
		74: "η",
		75: "⇐",
		76: "↔",
		77: "λ",
		78: "↵",
	}

	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(78)

	for i := 0; i < len; i++ {
		idNumber = stoi[sl[i]]
		idNumber += key
		idNumber %= 79
		sl[i] = itos[idNumber]
		cipher += sl[i]
	}

	keyStr = itos[key]

	cipher = keyStr + cipher

	return cipher
}