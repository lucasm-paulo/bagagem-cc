#!/usr/bin/env bash

# CAUTION: this script will replace every occurrence of the word
# `cc-tools-demo` in the project folder with the arguments you pass. Be very careful.

if [ $# -lt 2 ]; then
  printf "Usage:\n$ ./renameProject.sh <my-profile-name> <my-project-name>\n"
  exit
fi

old_project_name="cc-tools-demo"
new_project_name=$2
new_profile_name=$1

grep -rl "$old_project_name" . --exclude-dir={.git,node_modules} | xargs sed -i "s/$old_project_name/$new_project_name/g"
grep -rl "goledgerdev/$old_project_name" . --exclude-dir={.git,node_modules} | xargs sed -i "s/goledgerdev\/$old_project_name/$new_profile_name\/$new_project_name/g"
