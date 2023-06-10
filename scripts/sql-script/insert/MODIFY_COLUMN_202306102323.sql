ALTER TABLE `mall_go`.`ums_admin_login_log`
MODIFY COLUMN `user_agent` varchar (255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '浏览器登录类型' AFTER `address`;