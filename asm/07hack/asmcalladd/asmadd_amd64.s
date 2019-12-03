#include "textflag.h"

TEXT ·AsmCallCAdd(SB), NOSPLIT, $0
	MOVQ cfun+0(FP), AX // cfun
	MOVQ a+8(FP),    DI // a
	MOVQ b+16(FP),   SI // b
	CALL AX
	MOVQ AX, ret+24(FP)
	RET
