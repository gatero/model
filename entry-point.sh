#!/bin/bash

if [ ! -f Gopkg.yml ] && [ ! -f Gopkg.lock ]; then
  dep init
fi

if [ ! -d vendor ]; then
  dep ensure -v 
fi

if [ -f refresh.yml ]; then
  refresh run
fi
