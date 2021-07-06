__env_script() { # (name)
    NAME=$1
    PWD=`pwd`
    echo "alias $NAME='sh $PWD/tools/$NAME.sh';"
}
eval `__env_script build-all`
eval `__env_script build-go`
eval `__env_script build-kernel`
eval `__env_script build-supc`
eval `__env_script build-user-go`
eval `__env_script build-user-supc`
eval `__env_script build-user`
eval `__env_script env`
eval `__env_script link-user`
eval `__env_script link`
eval `__env_script log`
eval `__env_script make-drive`
eval `__env_script run-quick`
eval `__env_script start-qemu`