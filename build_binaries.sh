#!/bin/sh

perl -pi -w -e 's/package romandateandtime/package main/g' *go

oss=(linux windows darwin freebsd)
archs=(amd64 arm64)

P="romandateandtime"
SUFF=""
OUT="./bin"

if [ ! -d "${OUT}" ]; then
  mkdir ${OUT}
fi

rm ${OUT}/*.zip

for arch in ${archs[@]}
do
  for os in ${oss[@]}
  do
    echo "building ${os}-${arch}"
    if [ ${os} == "windows" ]; then
      SUFF=".exe"
    else
      SUFF=""
    fi
    EXE=${P}-${os}-${arch}${SUFF}
	  env GOOS=${os} GOARCH=${arch} go build -o ${P}${SUFF}
	  zip -q ${EXE}.zip ${P}${SUFF}
	  mv ${EXE}.zip ${OUT}/
	  rm ${P}${SUFF}
	done
done

go build -o ${P}${SUFF}

perl -pi -w -e 's/package main/package romandateandtime/g' *go
