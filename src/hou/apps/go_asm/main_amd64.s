#include "textflag.h"
TEXT ·neg(SB), NOSPLIT, $0
	MOVQ 	x+0(FP), AX
	NEGQ 	AX
	MOVQ 	AX, y+8(FP)
	RET
