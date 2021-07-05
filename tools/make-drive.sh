source tools/log.sh; step "IMAGE"; set -e
mkbootimg support/part.json build/messageq.img
qemu-img convert -O qcow2 build/messageq.img build/messageq.qcow2