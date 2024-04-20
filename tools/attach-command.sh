#!/bin/bash

cp /.gitconfig /home/user/.gitconfig
cp -r /.ssh/ /home/user/.ssh/
git config --global --add safe.directory /src
