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

do_call_onstack:
    mov rax, rsp
    mov rsp, rdi
    push rax
    mov rdi, rdx
    call rsi
    pop rsp
    ret

global read_cr3
read_cr3:
    mov rax, cr3
    ret

global read_cr2
read_cr2:
    mov rax, cr3
    ret

global tlb_flush
tlb_flush:
    mov rax, cr3
    mov cr3, rax
    ret

global _start
extern kinit
_start:
    mov rsp, _stack_top
    call kinit
    .do_cli_hlt:
        cli
        hlt
        jmp .do_cli_hlt

global do_rdmsr
do_rdmsr:
    mov rcx, rdi
    rdmsr
    shl rdx, 32
    or rax, rdx
    ret
    

global do_wrmsr
do_wrmsr:
    mov rax, rsi
    mov rdx, rax
    shr rdx, 32
    mov rcx, rdi
    wrmsr
    ret

global get_int_stack
get_int_stack:
    mov rax, _int_stack_top
    ret

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
_int_stack_top:
    resb 32 * 1024 ; 32K
_stack_top: