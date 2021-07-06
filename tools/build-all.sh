source tools/log.sh; step "BUILD ALL"
sh tools/build-user.sh
sh tools/build-kernel.sh
sh tools/make-drive.sh