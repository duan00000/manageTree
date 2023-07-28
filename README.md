# 基于fabric的橡胶林管理平台
###  内容说明
本平台采用vue+beego+fabric。
        rubber文件为后端代码文件，trace为区块链代码与智能合约文件，vue-rub为前端代码文件。
        前端采用vue框架，后端采用beego框架，fabric是1.4版本，采用fabric-sdk-go与后端连接，具有添加和编辑链上橡胶林数据，判定林区否可割胶，判定割胶方式，割胶路径规划等功能。

### fabric端部署过程

#### 1.环境准备

若要搭建本平台具有的链必须准备以下环境：

①配置go环境，go的版本最好大于1.16。

②配置docker环境和docker-compose。

#### 2.底链部署

①将项目文件manageTree中的trace文件拷贝到服务器上的/data路径下。

②