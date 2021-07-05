source tools/log.sh; step "LINK"; set -e
ld.lld build/*.o -o boot/kernel.elf -T support/link.ld