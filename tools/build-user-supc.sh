source tools/log.sh; step "BUILD USERLAND SUPC"; set -e
clang-12 -target x86_64-elf usersupport/support.c -o build/usersupport.o -c $CFLAGS -mcmodel=large -ggdb
nasm -felf64 usersupport/sup.s -o build/usersup.o