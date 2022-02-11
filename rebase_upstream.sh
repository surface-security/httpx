#!/bin/sh

git checkout master
git pull
git pull https://github.com/projectdiscovery/httpx.git master:temporary
git checkout -b upgrade_branch
# use merge (instead of rebase) for clear commit in history
git merge temporary
