source tools/log.sh; step "LINK"; set -e
ld.lld build/inthandler.o build/main.o build/sup.o build/support.o -o boot/kernel.elf -T support/link.ld