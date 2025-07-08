#!/bin/bash

RUN_NAME="main"

OSLinux="linux"
OSMac="darwin"
OSWindows="windows"

ArchAMD64="amd64"
Arch386="386"

os=$1
BUILD_VERSION="1.0"

#mkdir -p output/bin output/conf
#cp script/bootstrap.sh output 2>/dev/null
#chmod +x output/bootstrap.sh
#cp script/bootstrap.sh output/bootstrap_staging.sh
#chmod +x output/bootstrap_staging.sh
#find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/
#
#go build -o output/bin/${RUN_NAME}


mkdir -p bin/${OSLinux}
rm -rf bin/

if [[ "${os}" == "${OSLinux}" ]]; then

    mkdir -p bin/${OSLinux}
    GOOS=${OSLinux} GOARCH=${ArchAMD64} go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -o ${RUN_NAME}

elif [[ "${os}" == "${OSMac}" ]]; then

    mkdir -p bin/${OSMac}
    GOOS=${OSMac} GOARCH=${ArchAMD64} go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -o bin/${OSMac}/${RUN_NAME}

else

    mkdir -p bin/${OSLinux}
    GOOS=${OSLinux} GOARCH=${ArchAMD64} go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -o bin/${OSLinux}/${RUN_NAME}
fi


rm -rf ./output/
rm -rf ./${RUN_NAME}
cp -rp ./bin/linux/ output/
cp -rp templates output/
cp -rp script/ output/




RUN_NAME="stack.examples.hertz"

mkdir -p output/bin output/conf
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
cp script/bootstrap.sh output/bootstrap_staging.sh
chmod +x output/bootstrap_staging.sh
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

go build -o output/bin/${RUN_NAME}