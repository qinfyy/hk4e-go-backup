-- 基础信息
local base_info = {
	group_id = 133304021
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 21001, monster_id = 26090101, pos = { x = -1164.014, y = 166.932, z = 2489.993 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 30, drop_tag = "蕈兽", pose_id = 101, area_id = 21 },
	{ config_id = 21002, monster_id = 26090101, pos = { x = -1159.766, y = 166.939, z = 2483.343 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 30, drop_tag = "蕈兽", pose_id = 101, area_id = 21 }
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
		monsters = { 21001, 21002 },
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