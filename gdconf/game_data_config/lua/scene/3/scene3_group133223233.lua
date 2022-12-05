-- 基础信息
local base_info = {
	group_id = 133223233
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

-- 废弃数据
garbages = {
	gadgets = {
		{ config_id = 233001, gadget_id = 70220062, pos = { x = -6212.529, y = 199.229, z = -2811.825 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, area_id = 18 },
		{ config_id = 233002, gadget_id = 70220064, pos = { x = -6208.818, y = 200.082, z = -2812.702 }, rot = { x = 340.613, y = 0.000, z = 0.000 }, level = 33, area_id = 18 },
		{ config_id = 233003, gadget_id = 70220063, pos = { x = -6210.831, y = 199.845, z = -2813.513 }, rot = { x = 310.695, y = 330.121, z = 2.549 }, level = 33, area_id = 18 }
	}
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
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = ,
		monsters = { },
		gadgets = { },
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