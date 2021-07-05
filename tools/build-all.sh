source tools/log.sh; step "BUILD ALL"
sh tools/build-supc.sh
sh tools/build-go.sh
sh tools/link.sh
sh tools/make-drive.sh