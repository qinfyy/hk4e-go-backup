-- 基础信息
local base_info = {
	group_id = 133313100
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 100001, gadget_id = 70330450, pos = { x = -515.291, y = 50.581, z = 5727.504 }, rot = { x = 348.778, y = 83.169, z = 18.326 }, level = 32, area_id = 32 },
	{ config_id = 100002, gadget_id = 70500000, pos = { x = -515.291, y = 50.581, z = 5727.504 }, rot = { x = 348.778, y = 83.169, z = 18.326 }, level = 32, point_type = 3012, owner = 100001, area_id = 32 },
	{ config_id = 100003, gadget_id = 70330450, pos = { x = -507.480, y = 51.557, z = 5723.600 }, rot = { x = 7.302, y = 233.419, z = 357.754 }, level = 32, area_id = 32 },
	{ config_id = 100004, gadget_id = 70500000, pos = { x = -507.480, y = 51.557, z = 5723.600 }, rot = { x = 7.311, y = 233.419, z = 357.750 }, level = 32, point_type = 3012, owner = 100003, area_id = 32 },
	{ config_id = 100005, gadget_id = 70330450, pos = { x = -505.118, y = 51.675, z = 5735.426 }, rot = { x = 9.242, y = 331.240, z = 12.873 }, level = 32, area_id = 32 },
	{ config_id = 100006, gadget_id = 70500000, pos = { x = -505.118, y = 51.675, z = 5735.426 }, rot = { x = 9.242, y = 331.240, z = 12.873 }, level = 32, point_type = 3012, owner = 100005, area_id = 32 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { 100001, 100002, 100003, 100004, 100005, 100006 },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================