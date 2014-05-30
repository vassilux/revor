#!/bin/bash
#
# 
# Description : Prepare deploy revor package. 
# Author : vassilux
# Last modified : 2014-05-30 14:53:54  
#

set -e
DEPLOY_APP="revor"

revel package revor

rm -rf "$DEPLOY_DIR"

echo "Deploy finished."
