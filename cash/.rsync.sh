#!/bin/bash

# only execute the script when github token exists.
[ -z "$RSYNC_PASS" ] && echo "missing ssh key" && exit 2
[ -z "$DEST_PATH" ] && echo "missing deploy path" && exit 3
echo "Destination path: "$DEST_PATH

# write the rsync password.
mkdir -p /root/.ssh
echo -n "$RSYNC_PASS" > /root/.ssh/rsync.pass
chmod 600 /root/.ssh/rsync.pass



function rsync_file() {
    local src=$1
    local dest=$2
    local retry=3
    if [ -e $1 ];then
        echo "传输源文件： $src"
    else
        echo "传输源文件： $src 不存在"
        return 1
    fi
    if [ x_"${dest}" = x_ ];then
        echo "传输目的路径为空"
        return 1
    fi
    while [ $retry -gt 0 ] ;do
        rsync -avz --delete --timeout=60 --password-file=/root/.ssh/rsync.pass $src $dest
        if [ $? -eq 0 ] ;then
            echo "rsync file: ", $src , " to ", $dest, " success"
            return 0
        fi
        ((retry -= 1))
        sleep 10
    done
    echo "rsync file: ", $src , " to ", $dest, " success"
    return 1
}

# 传输编译文件
#rsync -avz --delete --timeout=60 --password-file=/root/.ssh/rsync.pass /go/src/management/release/server_linux* ${DEST_PATH}/server/
#rsync -avz --delete --timeout=60 --password-file=/root/.ssh/rsync.pass /go/src/management/release/wap_linux* ${DEST_PATH}/wap/
#rsync -avz --delete --timeout=60 --password-file=/root/.ssh/rsync.pass /go/src/management/release/front_linux* ${DEST_PATH}/front/
#rsync -avz --delete --timeout=60 --password-file=/root/.ssh/rsync.pass /go/src/management/release/admin_linux* ${DEST_PATH}/admin/

rsync_file "/go/src/management/release/server_linux" "${DEST_PATH}/server/"
rsync_file "/go/src/management/release/server_linux" "${DEST_PATH}/server/server_linux.${DRONE_BUILD_NUMBER}"
rsync_file "/go/src/management/release/server_linux.md5" "${DEST_PATH}/server/"
rsync_file "/go/src/management/release/server_linux.md5" "${DEST_PATH}/server/server_linux.${DRONE_BUILD_NUMBER}.md5"

rsync_file "/go/src/management/release/wap_linux" "${DEST_PATH}/wap/"
rsync_file "/go/src/management/release/wap_linux" "${DEST_PATH}/wap/wap_linux.${DRONE_BUILD_NUMBER}"
rsync_file "/go/src/management/release/wap_linux.md5" "${DEST_PATH}/wap/"
rsync_file "/go/src/management/release/wap_linux.md5" "${DEST_PATH}/wap/wap_linux.${DRONE_BUILD_NUMBER}.md5"

rsync_file "/go/src/management/release/admin_linux" "${DEST_PATH}/admin/"
rsync_file "/go/src/management/release/admin_linux" "${DEST_PATH}/admin/admin_linux.${DRONE_BUILD_NUMBER}"
rsync_file "/go/src/management/release/admin_linux.md5" "${DEST_PATH}/admin/"
rsync_file "/go/src/management/release/admin_linux.md5" "${DEST_PATH}/admin/admin_linux.${DRONE_BUILD_NUMBER}.md5"

rsync_file "/go/src/management/release/front_linux" "${DEST_PATH}/front/"
rsync_file "/go/src/management/release/front_linux" "${DEST_PATH}/front/front_linux.${DRONE_BUILD_NUMBER}"
rsync_file "/go/src/management/release/front_linux.md5" "${DEST_PATH}/front/"
rsync_file "/go/src/management/release/front_linux.md5" "${DEST_PATH}/front/front_linux.${DRONE_BUILD_NUMBER}.md5"
