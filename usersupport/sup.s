extern tinygo_scanstack
global tinygo_scanCurrentStack
global do_call_onstack

tinygo_scanCurrentStack:
    ; Save callee-saved registers.
    push rbx
    push rbp
    push r12
    push r13
    push r14
    push r15

    ; Scan the stack.
    sub rsp, 8 ; adjust the stack before the call to maintain 16-byte alignment
    mov rdi, rsp
    call tinygo_scanstack

    ; Restore the stack pointer. Registers do not need to be restored as they
    ; were only pushed to be discoverable by the GC.
    add rsp, 56
    ret


%macro intfunc 1
global int%1
int%1:
    int %1
    ret
%endmacro

%assign i 0
%rep 256
intfunc i
%assign i i+1
%endrep

section .bss
global _heap_start
global _heap_end
global _stack
global _stack_top

_heap_start:
    resb 1024 * 1024 ; 1MB
_heap_end:

_stack:
    resb 32 * 1024 ; 32K
_stack_top: