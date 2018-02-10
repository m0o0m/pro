#!/bin/sh
set -e

MajorVerson=`date +%y%m%d`
# DRONE_BUILD_NUMBER=1
VERSION=${MajorVerson}.${DRONE_BUILD_NUMBER}
export VERSION
echo "Build Version: "${VERSION}

go get github.com/labstack/gommon/log 
# cd /go/src/management/src

# Check environment variables
[ -z ${BUILD_ENV} ] && echo "missing build environment variable" && exit 2
echo "Build environment: "${BUILD_ENV}

#if [ x_${BUILD_ENV} = x_"production" ];then
#    [ -z ${PRODUCTION_REDIS_HOSTS} ] && echo "missing CI environment redis ip list" && exit 3
#    [ -z ${PRODUCTION_MYSQL_HOSTS} ] && echo "missing CI environment mysql ip list" && exit 4
#    echo "CI Production environment redis ip: "${PRODUCTION_REDIS_HOSTS}
#    echo "CI Production environment mysql ip: "${PRODUCTION_REDIS_HOSTS}
#    export PRODUCTION_REDIS_HOSTS PRODUCTION_MYSQL_HOSTS
#else
#    [ -z ${REDIS_HOSTS} ] && echo "missing CI environment redis ip list" && exit 3
#    [ -z ${MYSQL_HOSTS} ] && echo "missing CI environment mysql ip list" && exit 4
#    echo "CI Testing environment redis ip: "${REDIS_HOSTS}
#    echo "CI Testing environment mysql ip: "${MYSQL_HOSTS}
#    export REDIS_HOSTS MYSQL_HOSTS
#fi

#cd /go/src/management/src
## build command tools
#go run ./build.go build tools
#
#[ -z ${MONGODB_MASTER_URL} ] && echo "WARN: missing MONGODB_MASTER_URL environment variable"
#export MONGODB_MASTER_URL
#
#cd /go/src/management/bin/
#echo "### Configure server MongoDB ##########################################"
#./tools -type conf -config "../bin/etc/conf.yaml"
#echo "### Configure admin MongoDB ##########################################"
#./tools -type admin -config "../bin/etc/admin.yaml"
#echo "### Configure front MongoDB ##########################################"
#./tools -type front -config "../bin/etc/front.yaml"
#echo "### Configure wap MongoDB ##########################################"
#./tools -type wap -config "../bin/etc/wap.yaml"

# build commond server
mkdir /go/src/management/release/
cd /go/src/management/src/

if [ x_${BUILD_ENV} = x_"production" ];then
    echo "Build production server"
    go run build.go --environment=production --goos=linux build server
elif [ x_${BUILD_ENV} = x_"development" ];then
    echo "Build development server"
    go run build.go --environment=development --goos=linux build server
else
    echo "Build testing server"
    go run build.go --environment=testing --goos=linux build server
fi

cp /go/src/management/bin/server_linux* /go/src/management/release/

# build commond admin
# cd /go/src/management/src/
if [ x_${BUILD_ENV} = x_"production" ];then
    echo "Build production admin"
    go run build.go --environment=production --goos=linux build admin
elif [ x_${BUILD_ENV} = x_"development" ];then
    echo "Build development admin"
    go run build.go --environment=development --goos=linux build admin
else
    echo "Build testing admin"
    go run build.go --environment=testing --goos=linux build admin
fi

cp /go/src/management/bin/admin_linux* /go/src/management/release/

# build commond front
# cd /go/src/management/src/
if [ x_${BUILD_ENV} = x_"production" ];then
    echo "Build production front"
    go run build.go --environment=production --goos=linux build front
elif [ x_${BUILD_ENV} = x_"development" ];then
    echo "Build development front"
    go run build.go --environment=development --goos=linux build front
else
    echo "Build testing front"
    go run build.go --environment=testing --goos=linux build front
fi

cp /go/src/management/bin/front_linux* /go/src/management/release/

# build commond front
# cd /go/src/management/src/
if [ x_${BUILD_ENV} = x_"production" ];then
    echo "Build production wap"
    go run build.go --environment=production --goos=linux build wap
elif [ x_${BUILD_ENV} = x_"development" ];then
    echo "Build development wap"
    go run build.go --environment=development --goos=linux build wap
else
    echo "Build testing front"
    go run build.go --environment=testing --goos=linux build wap
fi

cp /go/src/management/bin/wap_linux* /go/src/management/release/


# list release
ls -l /go/src/management/release/
