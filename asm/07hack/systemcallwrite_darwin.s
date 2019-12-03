
// func SyscallWrite_Darwin(fd int, msg string) int
TEXT ·SyscallWrite_Darwin(SB), NOSPLIT, $0
	MOVQ $(0x2000000+4), AX // #define SYS_write 4
	MOVQ fd+0(FP),       DI
	MOVQ msg_data+8(FP), SI
	MOVQ msg_len+16(FP), DX
	SYSCALL
	MOVQ AX, ret+0(FP)
	RET
