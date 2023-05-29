#!/bin/bash
echo "启动容器orderer-----"
docker-compose -f docker-orderer.yaml up -d
echo "启动容器process的peer0节点和cli节点-----"
docker-compose -f docker-peer0-process.yaml up -d
echo "启动容器process的peer1节点-----"
docker-compose -f docker-peer1-process.yaml up -d
echo "启动容器company的peer0节点和cli1节点-----"
docker-compose -f docker-peer0-company.yaml up -d
echo "启动容器company的peer1节点-----"
docker-compose -f docker-peer1-company.yaml up -d
echo "启动容器hospital的peer0节点和cli2节点-----"
docker-compose -f docker-peer0-hospital.yaml up -d
echo "启动容hospital的peer1节点-----"
docker-compose -f docker-peer1-hospital.yaml up -d
