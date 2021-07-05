source tools/log.sh; step "QUICK RUN"; set -e
sh tools/build-go.sh
sh tools/link.sh
sh tools/make-drive.sh
sh tools/start-qemu.sh
