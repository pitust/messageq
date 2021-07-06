source tools/log.sh; step "BUILD KERNEL"
sh tools/build-user-supc.sh
sh tools/build-user-go.sh
sh tools/link-user.sh