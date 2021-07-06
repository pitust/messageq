source tools/log.sh; step "BUILD USERLAND GO"; set -e
tinygo build -target x86_64-unknown-linux "-tags=baremetal do_unix_rt interrupt.none amd64 tinygo gc.conservative scheduler.coroutines" -o build/user.ll -size full -scheduler coroutines -gc conservative -opt 0 userapp/app.go
clang-12 -target x86_64-elf build/user.ll -o build/user.o -c -Wno-override-module -mcmodel=large -O0 -ggdb