#include "textflag.h"

// func CopySlice_AVX2(dst, src []byte, len int)
TEXT ·CopySlice_AVX2(SB), NOSPLIT, $0
	MOVQ dst_data+0(FP),  DI
	MOVQ src_data+24(FP), SI
	MOVQ len+32(FP),      BX
	MOVQ $0,              AX

LOOP:
	VMOVDQU 0(SI)(AX*1), Y0
	VMOVDQU Y0, 0(DI)(AX*1)
	ADDQ $32, AX
	CMPQ AX, BX
	JL   LOOP
	RET
