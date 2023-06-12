#!/usr/bin/env bash

# CAUTION: this script will replace every occurrence of the word
# `goledgerdev/cc-tools-demo` in the project folder with whatever argument
# you pass. Be very careful.

if [ $# -lt 1 ] ; then
  printf "Usage:\n$ ./renameProject.sh <profile-name/my-project-name>\n"
  exit
fi

grep -rl goledgerdev/cc-tools-demo . --exclude-dir={.git,node_modules} | xargs sed -i "s#goledgerdev/cc-tools-demo#$1#g"

