#!/bin/bash
docker run -p 1234:1234 -v ${PWD}:/home/user/app ubuntu:16.04 /home/user/app/ElementaryFileServer

