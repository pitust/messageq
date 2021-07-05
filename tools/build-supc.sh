source tools/log.sh; step "BUILD SUPC"; set -e
clang-12 -target x86_64-elf support/support.c -o build/support.o -c $CFLAGS -mcmodel=kernel -ggdb
nasm -felf64 support/sup.s -o build/sup.o
nasm -felf64 support/inthandler.s -o build/inthandler.o