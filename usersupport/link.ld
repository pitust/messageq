SECTIONS
{
    . = 0x2000000000;
    _globals_start = .;
    .text : {
        KEEP(*(.text.boot)) *(.text .text.*)   /* code */
        *(.rodata .rodata.*)                   /* data */
        *(.data .data.*)
    }
    .bss (NOLOAD) : {                          /* bss */
        . = ALIGN(16);
        *(.bss .bss.*)
        *(COMMON)
    }
    .data : {                          /* data */
        . = ALIGN(16);
        *(.data .data.*)
        *(COMMON)
    }
    _globals_end = .;

    /DISCARD/ : { *(.eh_frame) *(.comment) }
}
