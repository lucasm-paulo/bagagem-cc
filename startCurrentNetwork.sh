#!/usr/bin/env bash

sleep 30
sudo docker start $(docker ps -a -q)
