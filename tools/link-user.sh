source tools/log.sh; step "LINK USER"; set -e
ld.lld build/user.o build/usersup.o build/usersupport.o -o boot/userland.elf -T usersupport/link.ld