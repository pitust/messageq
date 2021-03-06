extern do_isr_handle

%macro idtend 1
    dq isr%1
%endmacro

%macro isrgen 1

isr%1:
%if (%1 >= 0x8 && %1 <= 0xE) || %1 == 0x11 || %1 == 0x1E
%else
    push 0
%endif
    push rbp
    push rax
    push rbx
    push rcx
    push rdx
    push rsi
    push rdi
    push r8
    push r9
    push r10
    push r11
    push r12
    push r13
    push r14
    push r15
    mov ax, 0
    push rax

    xor rbp, rbp
    xor rax, rax
    xor rbx, rbx
    xor rcx, rcx
    xor rdx, rdx
    xor rsi, rsi
    xor rdi, rdi
    xor r8, r8
    xor r9, r9
    xor r10, r10
    xor r11, r11
    xor r12, r12
    xor r13, r13
    xor r14, r14
    xor r15, r15

    mov rdi, %1
    lea rsi, [rsp + 8]
    call do_isr_handle
    pop rax
    ; mov ds, ax
    pop r15
    pop r14
    pop r13
    pop r12
    pop r11
    pop r10
    pop r9
    pop r8
    pop rdi
    pop rsi
    pop rdx
    pop rcx
    pop rbx
    pop rax
    pop rbp
    add rsp, 8 ; error code
    iretq

%endmacro

global get_isr
get_isr:
    lea rax, [rel idt_targets]
    mov rax, [rax + rdi * 8]
    ret

global do_int3
do_int3:
    int3
    ret

global get_magic_rip
get_magic_rip:
    mov rax, do_int3
    ret

global do_lidt
do_lidt:
    lidt [rdi]
    ret

global do_lgdt
do_lgdt:
    lgdt [rdi]
    mov ax, 0x28
    ltr ax
    ret

idt_targets:
%assign i 0
%rep 256
    idtend i
%assign i i+1
%endrep

%assign i 0
%rep 256
isrgen i
%assign i i+1
%endrep
