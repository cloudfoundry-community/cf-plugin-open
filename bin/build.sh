#!/bin/bash

if [[ "$(which gox)X" == "X" ]]; then
  echo "Please install gox. https://github.com/mitchellh/gox#readme"
  exit 1
fi


rm -f cf-plugin-open*

gox -os linux -os windows -arch 386
gox -os darwin -os linux -os windows -arch amd64

rm -rf out
mkdir -p out
mv cf-plugin-open* out/
