/**
  @Author ZXQ-Administrator
  @Date 2021-09-08
  @node:
**/
package models

type Goods struct {
	Id            int     // 主键 自增
	Title         string  //商品标题
	SubTitle      string  //商品二级标题
	GoodsSn       string  //商品sn号
	CateId        int     //商品所属分类与所属分类进行关联
	ClickCount    int     //点击量
	GoodsNumber   int     //库存
	Price         float64 //价格
	MarketPrice   float64 //市场价格
	RelationGoods string  //关联商品
	GoodsAttr     string  //商品属性
	GoodsColor    string  //商品颜色
	GoodsVersion  string  //商品版本
	GoodsImg      string  //商品图片
	GoodsGift     string  //商品赠品
	GoodsFitting  string  //商品配件
	GoodsKeywords string  //商品关键词
	GoodsDesc     string  //商品描述
	GoodsContent  string //商品详情
	IsDelete      int
	IsHot         int
	IsBest        int
	IsNew         int
	GoodsTypeId   int //与商品类型关联 = goods_type中的id
	Sort          int
	Status        int
	AddTime       int
}

func (Goods) TableName() string {
	return "goods"
}

/*
   根据商品分类获取推荐商品
   @param {Number} cateId - 分类id
   @param {String} goodsType -  hot  best  new
   @param {Number} limitNum -  数量

顶级分类
1
	21
	22

二级分类
22

*/

func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []Goods {
	goods := []Goods{}
	//判断CateId是不是顶级分类
	goodsCate := []GoodsCate{}
	DB.Where("pid=?", cateId).Find(&goodsCate)
	var tempSlice []int
	if len(goodsCate) > 0 { //顶级分类
		for i := 0; i < len(goodsCate); i++ {
			tempSlice = append(tempSlice, goodsCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
	case "best":
		where += " AND is_best=1"
	case "new":
		where += " AND is_new=1"
	default:
		break
	}
	DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort desc").Find(&goods)
	return goods
}
