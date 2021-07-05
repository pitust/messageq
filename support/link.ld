mmio        = 0xfffffffff8000000;              /* these are configurable for level 2 loaders */
fb          = 0xfffffffffc000000;
bootboot    = 0xffffffffffe00000;
environment = 0xffffffffffe01000;

/* TODO: fuck the init stack (it's too small on L1 loaders anyway, so we switch really quickly in support/support.c) */
/* initstack = 1024; */
PHDRS
{
  boot PT_LOAD;                                /* one single loadable segment */
}
SECTIONS
{
    . = 0xffffffffffe02000;
    _globals_start = .;
    .text : {
        KEEP(*(.text.boot)) *(.text .text.*)   /* code */
        *(.rodata .rodata.*)                   /* data */
        *(.data .data.*)
    } :boot
    .bss (NOLOAD) : {                          /* bss */
        . = ALIGN(16);
        *(.bss .bss.*)
        *(COMMON)
    } :boot
    .data : {                          /* data */
        . = ALIGN(16);
        *(.data .data.*)
        *(COMMON)
    } :boot
    _globals_end = .;

    /DISCARD/ : { *(.eh_frame) *(.comment) }
}
