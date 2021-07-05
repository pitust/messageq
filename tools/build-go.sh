source tools/log.sh; step "BUILD GO"; set -e
tinygo build -target x86_64-unknown-linux "-tags=baremetal do_unix_rt interrupt.none amd64 tinygo gc.conservative scheduler.coroutines" -o build/main.ll -size full -scheduler coroutines -gc conservative -opt 0 main.go
clang-12 -target x86_64-elf build/main.ll -o build/main.o -c -Wno-override-module -mcmodel=kernel -O0 -ggdb