#!/bin/sh

# only execute the script when github token exists.
[ -z "$SSH_KEY" ] && echo "missing ssh key" && exit 2
# [ -z "$DEPLOY_PATH" ] && echo "missing deploy path" && exit 3


if [ -n "$DEPLOY_SANDBOX_PATH" ];then
	DEST_PATH=$DEPLOY_SANDBOX_PATH
elif [ -n "$DEPLOY_PRODUCTION_PATH" ];then
	DEST_PATH=$DEPLOY_PRODUCTION_PATH
else
    echo "missing deploy path"
    exit 3
fi

[ -z "$DEST_PATH" ] && echo "missing deploy path" && exit 4
echo "Destination path: "$DEST_PATH

# write the ssh key.
mkdir -p /root/.ssh
echo -n "$SSH_KEY" > /root/.ssh/id_rsa
chmod 600 /root/.ssh/id_rsa

REMOTE_HOST_KEY="202.66.30.26 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJ29fK3KruzZSq3/MOhA9ZYeeRVMLOaj+rx4tolZAA5fA9NYdKqLf7JGgza0OONPvXghBayyCA0Q+j8KukTG8Rs="
echo "$REMOTE_HOST_KEY" >> /root/.ssh/known_hosts


rsync -u -az -r  -e "ssh -i /root/.ssh/id_rsa" ./* ${DEST_PATH}