package game

import (
	pb "google.golang.org/protobuf/proto"
	"hk4e/gs/constant"
	"hk4e/gs/model"
	"hk4e/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"
	"time"
)

func (g *GameManager) GetShopmallDataReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user get shop mall, uid: %v", player.PlayerID)

	// PacketGetShopmallDataRsp
	getShopmallDataRsp := new(proto.GetShopmallDataRsp)
	getShopmallDataRsp.ShopTypeList = []uint32{900, 1052, 902, 1001, 903}
	g.SendMsg(cmd.GetShopmallDataRsp, player.PlayerID, player.ClientSeq, getShopmallDataRsp)
}

func (g *GameManager) GetShopReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user get shop, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.GetShopReq)
	shopType := req.ShopType

	if shopType != 1001 {
		return
	}

	nextRefreshTime := uint32(time.Now().Add(time.Hour * 24 * 30).Unix())

	// PacketGetShopRsp
	getShopRsp := new(proto.GetShopRsp)
	getShopRsp.Shop = &proto.Shop{
		GoodsList: []*proto.ShopGoods{
			{
				MinLevel:        1,
				EndTime:         2051193600,
				Hcoin:           160,
				GoodsId:         102001,
				NextRefreshTime: nextRefreshTime,
				MaxLevel:        99,
				BeginTime:       1575129600,
				GoodsItem: &proto.ItemParam{
					ItemId: 223,
					Count:  1,
				},
			},
			{
				MinLevel:        1,
				EndTime:         2051193600,
				Hcoin:           160,
				GoodsId:         102002,
				NextRefreshTime: nextRefreshTime,
				MaxLevel:        99,
				BeginTime:       1575129600,
				GoodsItem: &proto.ItemParam{
					ItemId: 224,
					Count:  1,
				},
			},
		},
		NextRefreshTime: nextRefreshTime,
		ShopType:        1001,
	}
	g.SendMsg(cmd.GetShopRsp, player.PlayerID, player.ClientSeq, getShopRsp)
}

func (g *GameManager) BuyGoodsReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user buy goods, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.BuyGoodsReq)
	buyItemId := req.Goods.GoodsItem.ItemId
	buyItemCount := req.BuyCount
	costHcoinCount := req.Goods.Hcoin * buyItemCount

	if buyItemId != 223 && buyItemId != 224 {
		return
	}

	if player.GetItemCount(201) < costHcoinCount {
		return
	}
	g.CostUserItem(player.PlayerID, []*UserItem{{
		ItemId:      201,
		ChangeCount: costHcoinCount,
	}})

	g.AddUserItem(player.PlayerID, []*UserItem{{
		ItemId:      buyItemId,
		ChangeCount: buyItemCount,
	}}, true, constant.ActionReasonConst.Shop)
	req.Goods.BoughtNum = player.GetItemCount(buyItemId)

	// PacketBuyGoodsRsp
	buyGoodsRsp := new(proto.BuyGoodsRsp)
	buyGoodsRsp.ShopType = req.ShopType
	buyGoodsRsp.BuyCount = req.BuyCount
	buyGoodsRsp.GoodsList = []*proto.ShopGoods{req.Goods}
	g.SendMsg(cmd.BuyGoodsRsp, player.PlayerID, player.ClientSeq, buyGoodsRsp)
}

func (g *GameManager) McoinExchangeHcoinReq(player *model.Player, payloadMsg pb.Message) {
	logger.LOG.Debug("user mcoin exchange hcoin, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.McoinExchangeHcoinReq)
	if req.Hcoin != req.McoinCost {
		return
	}
	count := req.Hcoin

	if player.GetItemCount(203) < count {
		return
	}
	g.CostUserItem(player.PlayerID, []*UserItem{{
		ItemId:      203,
		ChangeCount: count,
	}})

	g.AddUserItem(player.PlayerID, []*UserItem{{
		ItemId:      201,
		ChangeCount: count,
	}}, false, 0)

	// PacketMcoinExchangeHcoinRsp
	mcoinExchangeHcoinRsp := new(proto.McoinExchangeHcoinRsp)
	mcoinExchangeHcoinRsp.Hcoin = req.Hcoin
	mcoinExchangeHcoinRsp.McoinCost = req.McoinCost
	g.SendMsg(cmd.McoinExchangeHcoinRsp, player.PlayerID, player.ClientSeq, mcoinExchangeHcoinRsp)
}