package blockchain

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)


type ChainCodeSpec struct {
	client      *channel.Client
	chaincodeId string
}

func Initialize(channelId, chainCodeId,orgName, userId, conf string) (*ChainCodeSpec, error) {
	//依据配置生成SDK操作失败-er
	config := beego.AppConfig.String(conf)
	sdk, err := getSDK(config)
	if err != nil {
		return nil, err
	}
	//通过SDK创建client
	clientChannelContext := sdk.ChannelContext(channelId, fabsdk.WithUser(userId), fabsdk.WithOrg(orgName))
	client, err := channel.New(clientChannelContext)

	if err != nil {
		beego.Error("Failed to create new client:", err)
		return nil, err
	}

	return &ChainCodeSpec{client, chainCodeId}, nil
}
//创建SDK
func getSDK(conf string) (*fabsdk.FabricSDK, error) {
	sdk, err := fabsdk.New(config.FromFile(conf))
	if err != nil {
		beego.Error("Failed to create new SDK:", err)
		return nil, err
	}
	return sdk, nil
}

func (c *ChainCodeSpec) ChainCodeUpdate(function string, args [][]byte,orgPeer string)  ([]byte, error) {
	beego.Info("orgPeer",orgPeer)
	response, err := c.client.Execute(channel.Request{ChaincodeID: c.chaincodeId, Fcn: function, Args: args},channel.WithTargetEndpoints(orgPeer))
	if err != nil {
		beego.Error("Failed to move funds:", err)
		return nil,err
	}
	return response.Payload,nil
}


func (c *ChainCodeSpec) ChainCodeQuery(function string, args [][]byte,orgPeer string) ([]byte, error) {
	response, err := c.client.Query(channel.Request{ChaincodeID: c.chaincodeId, Fcn: function, Args: args},channel.WithTargetEndpoints(orgPeer))
	if err != nil {
		beego.Error("Failed to query funds:", err)
		return nil, err
	}
	fmt.Println(response)

	return response.Payload, nil
}
