#!/bin/bash
#
# 
# Description : Prepare deploy revor package. 
# Author : vassilux
# Last modified : 2014-05-30 14:53:54  
#

set -e

VERSION=$(cat VERSION)

DEPLOY_APP="revor"


DEPLOY_APP_VER="${DEPLOY_APP}_${VERSION}"
DEPLOY_DIR="${DEPLOY_APP_VER}"
#
if [ -d releases ]; then
        rm -rf  releases
fi
#
if [ -d "${DEPLOY_DIR}" ]; then
	rm -rf ${DEPLOY_DIR}
fi

if [ -f "${DEPLOY_APP}.tar.gz" ]; then
	rm -rf "${DEPLOY_APP}.tar.gz"
fi

mkdir "${DEPLOY_DIR}"
mkdir "${DEPLOY_DIR}/samples"
cp -R "./samples/revor.supervisor.conf" "${DEPLOY_DIR}/samples/revor.supervisor.conf"

revel package "${DEPLOY_APP}"

#
pandoc -o "$DEPLOY_DIR/INSTALL.html" ./docs/INSTALL.md
pandoc -o "$DEPLOY_DIR/ReleaseNotes.html" ./docs/ReleaseNotes.md
#
cp "$DEPLOY_DIR/INSTALL.html" .
cp "$DEPLOY_DIR/ReleaseNotes.html" .
#
mv "${DEPLOY_APP}.tar.gz" "$DEPLOY_DIR"
cd "${DEPLOY_DIR}"
tar xzf "${DEPLOY_APP}.tar.gz"
rm -rf "${DEPLOY_APP}.tar.gz"
cd ..
tar czf "${DEPLOY_APP_VER}.tar.gz" ${DEPLOY_DIR}

if [ ! -d releases ]; then
        mkdir releases
fi

mv "${DEPLOY_APP_VER}.tar.gz" ./releases
mv INSTALL.* ./releases
mv ReleaseNotes.* ./releases

rm -rf "$DEPLOY_DIR"


rm -rf "${DEPLOY_APP_VER}"

echo "Deploy finished."
