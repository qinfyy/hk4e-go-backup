-- 基础信息
local base_info = {
	group_id = 199001030
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 30001, monster_id = 21010301, pos = { x = 317.401, y = 128.231, z = 157.309 }, rot = { x = 0.000, y = 44.017, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9002, area_id = 402 },
	{ config_id = 30002, monster_id = 21010301, pos = { x = 325.603, y = 132.062, z = 150.285 }, rot = { x = 0.000, y = 135.461, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9013, area_id = 402 },
	{ config_id = 30003, monster_id = 21010301, pos = { x = 347.651, y = 139.006, z = 157.814 }, rot = { x = 0.000, y = 152.500, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9012, area_id = 402 },
	{ config_id = 30004, monster_id = 21010301, pos = { x = 350.124, y = 138.996, z = 155.385 }, rot = { x = 0.000, y = 325.123, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9012, area_id = 402 },
	{ config_id = 30005, monster_id = 21010301, pos = { x = 357.337, y = 146.662, z = 182.161 }, rot = { x = 0.000, y = 14.140, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9003, area_id = 402 },
	{ config_id = 30006, monster_id = 21010301, pos = { x = 356.284, y = 146.662, z = 183.755 }, rot = { x = 0.000, y = 14.140, z = 0.000 }, level = 20, drop_tag = "丘丘人", pose_id = 9003, area_id = 402 }
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
		monsters = { 30001, 30002, 30003, 30004, 30005, 30006 },
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