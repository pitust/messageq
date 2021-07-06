source tools/log.sh; step "BUILD KERNEL"
sh tools/build-supc.sh
sh tools/build-go.sh
sh tools/link.sh