### 部署过程记录：

#### Fabric配置和证书生成脚本

```shell
cd /data/trace


cryptogen generate --config=./crypto-config.yaml



configtxgen -profile OrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block


configtxgen \
-profile OrgsChannel \
-outputCreateChannelTx ./channel-artifacts/channel.tx \
-channelID mychannel	



```

```shell
./start.sh	
```

### Fabric操作

#### 进入cli

```shell
docker exec -it cli bash
```

#### 创建channel

```shell
peer channel create \
-o orderer.trace.com:7050 \
-c mychannel \
-f ./channel-artifacts/channel.tx \
--tls \
--cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem 
```

#### 加入通道

peer0-org1加入通道

```shell
peer channel join -b mychannel.block
```

#### 链码安装和初始化

```shell
peer chaincode install \
-n mycc1 -v 1.0 \
-p github.com/chaincode/chaincode_example02/go/


```

```shell
peer chaincode instantiate \
-o orderer.trace.com:7050 \
--tls \
--cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/trace.com/orderers/orderer.trace.com/msp/tlscacerts/tlsca.trace.com-cert.pem \
-C mychannel \
-n mycc1 \
-v 1.0 \
-c '{"Args": [""]}' \
-P "OR('OrgProcessMSP.member','OrgCompanyMSP.member','OrgHospitalMSP.member')"
```

#### 测试链码

```shell
peer chaincode query -C mychannel -n mycc1 -c '{"Args": ["query", "a"]}'
100  #结果输出100

```

### 关停服务

```shell
./stop.sh
```
