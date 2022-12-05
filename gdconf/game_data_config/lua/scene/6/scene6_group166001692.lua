-- 基础信息
local base_info = {
	group_id = 166001692
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
		{ config_id = 692001, gadget_id = 70900382, pos = { x = 731.935, y = 981.345, z = 1071.149 }, rot = { x = 358.593, y = 18.115, z = 149.140 }, level = 36, area_id = 300 },
		{ config_id = 692002, gadget_id = 70900382, pos = { x = 726.815, y = 982.575, z = 1071.032 }, rot = { x = 2.108, y = 198.466, z = 180.434 }, level = 36, area_id = 300 },
		{ config_id = 692003, gadget_id = 70900382, pos = { x = 722.303, y = 981.567, z = 1067.463 }, rot = { x = 1.950, y = 295.438, z = 210.868 }, level = 36, area_id = 300 },
		{ config_id = 692004, gadget_id = 70900382, pos = { x = 728.621, y = 980.987, z = 1065.338 }, rot = { x = 337.016, y = 280.940, z = 207.646 }, level = 36, area_id = 300 },
		{ config_id = 692005, gadget_id = 70900382, pos = { x = 725.698, y = 979.076, z = 1063.094 }, rot = { x = 342.130, y = 287.907, z = 238.818 }, level = 36, area_id = 300 },
		{ config_id = 692006, gadget_id = 70900382, pos = { x = 729.498, y = 976.718, z = 1061.192 }, rot = { x = 17.044, y = 82.303, z = 110.169 }, level = 36, area_id = 300 }
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
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================