#!/bin/bash


_deploy() {
    docker-compose rm
    docker-compose kill

    docker rm $(docker ps -al -q) -f
    
    # sudo rm -R ./volumes/mysql
    # sudo rm -R ./volumes/wordpress
    # sudo rm -R ./volumes/dump
    
    mkdir -p ./volumes/mysql
    mkdir -p ./volumes/wordpress
    mkdir -p ./volumes/dump
    
    # cp /Users/altus/Projects/bls/azapi/databases/zanew/zanew.sql ./volumes/dump
    # rsync -rv --exclude=.htaccess /Users/altus/Projects/bls/azapi/quarantine/za/ ./volumes/wordpress

    docker-compose up --build --force-recreate 
}

if [[ $1 =~ ^(deploy|start|startDatabase|reateTables|test|killDatabase|initiateDatabase|killAPI|createDomain|deleteDomain)$ ]]; then
  COMMAND=_$1
  shift
  $COMMAND "$@"
else
  exit 1
fi