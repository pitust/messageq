source tools/log.sh; step "START QEMU"; set -e
qemu-system-x86_64 -hda build/messageq.qcow2 -debugcon stdio -s