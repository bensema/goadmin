/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : goadmin

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 24/03/2021 21:22:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `admin_id` varchar(50) NOT NULL COMMENT '管理员编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '管理员',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1:正常;2:禁用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `admin_id` (`admin_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='管理员';

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, '1', 'root', '$2a$10$muCPVqgBylixJjYfdhJorOfauVad9ywpFU.zdy1.XaMIoIZyVTECG', 1, '2021-02-18 15:24:46', '2021-02-18 15:24:46');
INSERT INTO `admin` VALUES (9, '9', 'demo996', '$2a$10$fMmFoKXbVpWxtSIy7fiTL.Ay4.QlTdIIuRdkjApPM6qsXBwCInL0q', 1, '2021-02-26 13:26:33', '2021-02-26 13:26:33');
COMMIT;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `admin_id` varchar(50) NOT NULL COMMENT '账户编号',
  `role_id` int(11) NOT NULL COMMENT '角色编号',
  PRIMARY KEY (`id`),
  KEY `admin_id` (`admin_id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COMMENT='管理员-角色';

-- ----------------------------
-- Records of admin_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_role` VALUES (61, '1', 1);
INSERT INTO `admin_role` VALUES (62, '1', 5);
INSERT INTO `admin_role` VALUES (63, '1', 8);
INSERT INTO `admin_role` VALUES (64, '9', 1);
INSERT INTO `admin_role` VALUES (65, '9', 8);
INSERT INTO `admin_role` VALUES (66, '', 2);
COMMIT;

-- ----------------------------
-- Table structure for log_admin_login
-- ----------------------------
DROP TABLE IF EXISTS `log_admin_login`;
CREATE TABLE `log_admin_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `admin_id` varchar(50) NOT NULL COMMENT '管理员编号',
  `name` varchar(36) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '管理员名',
  `location` varchar(100) NOT NULL DEFAULT '' COMMENT '位置',
  `os` text NOT NULL COMMENT '操作系统',
  `browser` text NOT NULL COMMENT '浏览器',
  `user_agent` text CHARACTER SET utf8 NOT NULL COMMENT '浏览器详情',
  `url` varchar(64) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT 'url',
  `result` tinyint(1) NOT NULL DEFAULT '1' COMMENT '2:失败;1:成功',
  `ip` varchar(16) CHARACTER SET utf8 NOT NULL DEFAULT '0.0.0.0' COMMENT 'IP',
  `record_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '记录时间',
  `remark` varchar(128) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_name` (`record_at`),
  KEY `ip` (`ip`,`record_at`),
  KEY `url` (`url`,`record_at`),
  KEY `record_time` (`record_at`),
  KEY `user_id` (`admin_id`,`record_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员登录日志';

-- ----------------------------
-- Records of log_admin_login
-- ----------------------------
BEGIN;
INSERT INTO `log_admin_login` VALUES (42, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.182', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-26 13:59:50', '');
INSERT INTO `log_admin_login` VALUES (43, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.182', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-26 14:03:35', '');
INSERT INTO `log_admin_login` VALUES (44, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.182', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-26 23:20:08', '');
INSERT INTO `log_admin_login` VALUES (45, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-27 15:43:31', '');
INSERT INTO `log_admin_login` VALUES (46, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-27 15:44:24', '');
INSERT INTO `log_admin_login` VALUES (47, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-27 15:45:39', '');
INSERT INTO `log_admin_login` VALUES (48, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-02-27 15:48:00', '');
INSERT INTO `log_admin_login` VALUES (49, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 12:19:02', '');
INSERT INTO `log_admin_login` VALUES (50, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 12:39:03', '');
INSERT INTO `log_admin_login` VALUES (51, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 13:29:47', '');
INSERT INTO `log_admin_login` VALUES (52, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 14:16:37', '');
INSERT INTO `log_admin_login` VALUES (53, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 14:21:11', '');
INSERT INTO `log_admin_login` VALUES (54, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 17:28:05', '');
INSERT INTO `log_admin_login` VALUES (55, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 19:30:23', '');
INSERT INTO `log_admin_login` VALUES (56, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 19:33:39', '');
INSERT INTO `log_admin_login` VALUES (57, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-02 19:37:12', '');
INSERT INTO `log_admin_login` VALUES (58, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-03 07:13:08', '');
INSERT INTO `log_admin_login` VALUES (59, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome88.0.4324.192', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-03 11:08:47', '');
INSERT INTO `log_admin_login` VALUES (60, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.82', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-24 12:07:55', '');
INSERT INTO `log_admin_login` VALUES (61, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.82', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-24 20:51:34', '');
COMMIT;

-- ----------------------------
-- Table structure for log_admin_operation
-- ----------------------------
DROP TABLE IF EXISTS `log_admin_operation`;
CREATE TABLE `log_admin_operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '操作编号',
  `admin_id` varchar(50) NOT NULL COMMENT '管理员编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '账户',
  `operation_code` varchar(50) NOT NULL DEFAULT '' COMMENT '行为编号',
  `operation_name` varchar(50) NOT NULL DEFAULT '' COMMENT '行为',
  `content` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '操作内容',
  `result` tinyint(1) NOT NULL DEFAULT '1' COMMENT '操作结果1:成功；2:失败',
  `message` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作消息',
  `ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作IP',
  `record_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`admin_id`,`operation_code`,`record_at`),
  KEY `user_id_2` (`admin_id`,`record_at`),
  KEY `action_module` (`operation_code`)
) ENGINE=InnoDB AUTO_INCREMENT=253 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员操作日志';

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `icon` varchar(50) NOT NULL DEFAULT '' COMMENT 'icon',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'url',
  `index_sort` int(11) NOT NULL DEFAULT '1' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COMMENT='菜单';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, '系统管理', 0, 'layui-icon-component', '', 1);
INSERT INTO `menu` VALUES (2, '管理员管理', 1, 'layui-icon-user', '/admin', 1);
INSERT INTO `menu` VALUES (3, '角色管理', 1, 'layui-icon-edit', '/role', 2);
INSERT INTO `menu` VALUES (7, '权限管理', 1, 'layui-icon-edit', '/permission', 3);
INSERT INTO `menu` VALUES (8, '资源管理', 1, 'layui-icon-edit', '/resources', 4);
INSERT INTO `menu` VALUES (16, '登录日志', 1, 'layui-icon-list', '/log/login', 5);
INSERT INTO `menu` VALUES (23, '操作日志', 1, 'layui-icon-list', '/log/operation', 6);
INSERT INTO `menu` VALUES (24, '内容管理', 0, 'layui-icon-component', '', 2);
INSERT INTO `menu` VALUES (25, '广告管理', 24, '', '', 0);
INSERT INTO `menu` VALUES (26, '公告管理', 24, '', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for operation
-- ----------------------------
DROP TABLE IF EXISTS `operation`;
CREATE TABLE `operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '编码',
  `method` varchar(50) NOT NULL DEFAULT '' COMMENT '方法',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'url',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COMMENT='操作';

-- ----------------------------
-- Records of operation
-- ----------------------------
BEGIN;
INSERT INTO `operation` VALUES (21, '查询管理员', 'search', 'get', '/api/v1/admin/pages', 2);
INSERT INTO `operation` VALUES (22, '管理员信息', 'admin_info', 'get', '/api/v1/admin/info', 2);
INSERT INTO `operation` VALUES (23, '更新管理员', 'update_admin', 'post', '/api/v1/admin/update', 2);
INSERT INTO `operation` VALUES (24, '删除管理员', 'delete_admin', 'post', '/api/v1/admin/delete', 2);
INSERT INTO `operation` VALUES (25, '添加管理员', 'add_admin', 'post', '/api/v1/admin/add', 2);
INSERT INTO `operation` VALUES (26, '查询角色', 'query_role', 'get', '/api/v1/role/pages', 3);
INSERT INTO `operation` VALUES (27, '角色信息', 'role_info', 'get', '/api/v1/role/info', 3);
INSERT INTO `operation` VALUES (28, '更新角色', 'update_role', 'post', '/api/v1/role/update', 3);
INSERT INTO `operation` VALUES (29, '添加角色', 'add_role', 'post', '/api/v1/role/add', 3);
INSERT INTO `operation` VALUES (30, '删除角色', 'delete_role', 'post', '/api/v1/role/delete', 3);
INSERT INTO `operation` VALUES (31, '查询权限', 'query_permission', 'get', '/api/v1/permission/pages', 7);
INSERT INTO `operation` VALUES (32, '添加权限', 'add_permission', 'post', '/api/v1/permission/add', 7);
INSERT INTO `operation` VALUES (33, '更新权限', 'update_permission', 'post', '/api/v1/permission/update', 7);
INSERT INTO `operation` VALUES (34, '删除权限', 'delete_permission', 'post', '/api/v1/permission/delete', 7);
INSERT INTO `operation` VALUES (35, '权限-菜单', 'query_permission_menu', 'get', '/api/v1/permission_menu/find', 7);
INSERT INTO `operation` VALUES (36, '权限-API操作', 'query_permission_operation', 'get', '/api/v1/permission_operation/find', 7);
INSERT INTO `operation` VALUES (37, '添加菜单', 'add_menu', 'post', '/api/v1/menu/add', 8);
INSERT INTO `operation` VALUES (38, '删除菜单', 'delete_menu', 'post', '/api/v1/menu/delete', 8);
INSERT INTO `operation` VALUES (39, '更新菜单', 'update_menu', 'post', '/api/v1/menu/update', 8);
INSERT INTO `operation` VALUES (40, '添加API操作', 'add_operation', 'post', '/api/v1/operation/add', 8);
INSERT INTO `operation` VALUES (41, '删除API操作', 'delete_operation', 'post', '/api/v1/operation/delete', 8);
INSERT INTO `operation` VALUES (42, '更新API操作', 'update_operation', 'post', '/api/v1/operation/update', 8);
INSERT INTO `operation` VALUES (43, '查询登录日志', 'query_log_login', 'get', '/api/v1/log_login/pages', 16);
INSERT INTO `operation` VALUES (44, '查询操作日志', 'query_log_operation', 'get', '/api/v1/log_operation/pages', 23);
INSERT INTO `operation` VALUES (45, '菜单-ALL', 'all_menu', 'get', '/api/v1/menu/all', 8);
INSERT INTO `operation` VALUES (46, 'API操作-ALL', 'operation_all', 'get', '/api/v1/operation/all', 8);
INSERT INTO `operation` VALUES (47, '角色-ALL', 'role_all', 'get', '/api/v1/role/all', 2);
INSERT INTO `operation` VALUES (48, '权限-ALL', 'permission_all', 'get', '/api/v1/permission/all', 3);
COMMIT;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='权限';

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` VALUES (1, '系统权限');
INSERT INTO `permission` VALUES (2, '操作员');
INSERT INTO `permission` VALUES (10, '查询');
COMMIT;

-- ----------------------------
-- Table structure for permission_menu
-- ----------------------------
DROP TABLE IF EXISTS `permission_menu`;
CREATE TABLE `permission_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `permission_id` int(11) NOT NULL COMMENT '权限编号',
  `menu_id` int(11) NOT NULL COMMENT '菜单编号',
  PRIMARY KEY (`id`),
  KEY `permission_id` (`permission_id`),
  KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=172 DEFAULT CHARSET=utf8mb4 COMMENT='权限-菜单';

-- ----------------------------
-- Records of permission_menu
-- ----------------------------
BEGIN;
INSERT INTO `permission_menu` VALUES (88, 2, 1);
INSERT INTO `permission_menu` VALUES (89, 2, 3);
INSERT INTO `permission_menu` VALUES (90, 2, 7);
INSERT INTO `permission_menu` VALUES (91, 2, 8);
INSERT INTO `permission_menu` VALUES (92, 2, 16);
INSERT INTO `permission_menu` VALUES (93, 2, 23);
INSERT INTO `permission_menu` VALUES (155, 1, 1);
INSERT INTO `permission_menu` VALUES (156, 1, 2);
INSERT INTO `permission_menu` VALUES (157, 1, 3);
INSERT INTO `permission_menu` VALUES (158, 1, 7);
INSERT INTO `permission_menu` VALUES (159, 1, 8);
INSERT INTO `permission_menu` VALUES (160, 1, 16);
INSERT INTO `permission_menu` VALUES (161, 1, 23);
INSERT INTO `permission_menu` VALUES (162, 1, 24);
INSERT INTO `permission_menu` VALUES (163, 1, 25);
INSERT INTO `permission_menu` VALUES (164, 1, 26);
INSERT INTO `permission_menu` VALUES (165, 10, 1);
INSERT INTO `permission_menu` VALUES (166, 10, 2);
INSERT INTO `permission_menu` VALUES (167, 10, 3);
INSERT INTO `permission_menu` VALUES (168, 10, 7);
INSERT INTO `permission_menu` VALUES (169, 10, 8);
INSERT INTO `permission_menu` VALUES (170, 10, 16);
INSERT INTO `permission_menu` VALUES (171, 10, 23);
COMMIT;

-- ----------------------------
-- Table structure for permission_operation
-- ----------------------------
DROP TABLE IF EXISTS `permission_operation`;
CREATE TABLE `permission_operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `permission_id` int(11) NOT NULL COMMENT '权限编号',
  `operation_id` int(11) NOT NULL COMMENT '操作编号',
  PRIMARY KEY (`id`),
  KEY `permission_id` (`permission_id`),
  KEY `operation_id` (`operation_id`)
) ENGINE=InnoDB AUTO_INCREMENT=336 DEFAULT CHARSET=utf8mb4 COMMENT='权限-操作';

-- ----------------------------
-- Records of permission_operation
-- ----------------------------
BEGIN;
INSERT INTO `permission_operation` VALUES (56, 2, 26);
INSERT INTO `permission_operation` VALUES (57, 2, 27);
INSERT INTO `permission_operation` VALUES (58, 2, 28);
INSERT INTO `permission_operation` VALUES (59, 2, 29);
INSERT INTO `permission_operation` VALUES (60, 2, 30);
INSERT INTO `permission_operation` VALUES (61, 2, 48);
INSERT INTO `permission_operation` VALUES (62, 2, 31);
INSERT INTO `permission_operation` VALUES (63, 2, 32);
INSERT INTO `permission_operation` VALUES (64, 2, 33);
INSERT INTO `permission_operation` VALUES (65, 2, 34);
INSERT INTO `permission_operation` VALUES (66, 2, 35);
INSERT INTO `permission_operation` VALUES (67, 2, 36);
INSERT INTO `permission_operation` VALUES (68, 2, 37);
INSERT INTO `permission_operation` VALUES (69, 2, 38);
INSERT INTO `permission_operation` VALUES (70, 2, 39);
INSERT INTO `permission_operation` VALUES (71, 2, 40);
INSERT INTO `permission_operation` VALUES (72, 2, 41);
INSERT INTO `permission_operation` VALUES (73, 2, 42);
INSERT INTO `permission_operation` VALUES (74, 2, 45);
INSERT INTO `permission_operation` VALUES (75, 2, 46);
INSERT INTO `permission_operation` VALUES (76, 2, 43);
INSERT INTO `permission_operation` VALUES (77, 2, 44);
INSERT INTO `permission_operation` VALUES (294, 1, 21);
INSERT INTO `permission_operation` VALUES (295, 1, 22);
INSERT INTO `permission_operation` VALUES (296, 1, 23);
INSERT INTO `permission_operation` VALUES (297, 1, 24);
INSERT INTO `permission_operation` VALUES (298, 1, 25);
INSERT INTO `permission_operation` VALUES (299, 1, 47);
INSERT INTO `permission_operation` VALUES (300, 1, 26);
INSERT INTO `permission_operation` VALUES (301, 1, 27);
INSERT INTO `permission_operation` VALUES (302, 1, 28);
INSERT INTO `permission_operation` VALUES (303, 1, 29);
INSERT INTO `permission_operation` VALUES (304, 1, 30);
INSERT INTO `permission_operation` VALUES (305, 1, 48);
INSERT INTO `permission_operation` VALUES (306, 1, 31);
INSERT INTO `permission_operation` VALUES (307, 1, 32);
INSERT INTO `permission_operation` VALUES (308, 1, 33);
INSERT INTO `permission_operation` VALUES (309, 1, 34);
INSERT INTO `permission_operation` VALUES (310, 1, 35);
INSERT INTO `permission_operation` VALUES (311, 1, 36);
INSERT INTO `permission_operation` VALUES (312, 1, 37);
INSERT INTO `permission_operation` VALUES (313, 1, 38);
INSERT INTO `permission_operation` VALUES (314, 1, 39);
INSERT INTO `permission_operation` VALUES (315, 1, 40);
INSERT INTO `permission_operation` VALUES (316, 1, 41);
INSERT INTO `permission_operation` VALUES (317, 1, 42);
INSERT INTO `permission_operation` VALUES (318, 1, 45);
INSERT INTO `permission_operation` VALUES (319, 1, 46);
INSERT INTO `permission_operation` VALUES (320, 1, 43);
INSERT INTO `permission_operation` VALUES (321, 1, 44);
INSERT INTO `permission_operation` VALUES (322, 10, 21);
INSERT INTO `permission_operation` VALUES (323, 10, 22);
INSERT INTO `permission_operation` VALUES (324, 10, 23);
INSERT INTO `permission_operation` VALUES (325, 10, 47);
INSERT INTO `permission_operation` VALUES (326, 10, 26);
INSERT INTO `permission_operation` VALUES (327, 10, 27);
INSERT INTO `permission_operation` VALUES (328, 10, 48);
INSERT INTO `permission_operation` VALUES (329, 10, 31);
INSERT INTO `permission_operation` VALUES (330, 10, 35);
INSERT INTO `permission_operation` VALUES (331, 10, 36);
INSERT INTO `permission_operation` VALUES (332, 10, 45);
INSERT INTO `permission_operation` VALUES (333, 10, 46);
INSERT INTO `permission_operation` VALUES (334, 10, 43);
INSERT INTO `permission_operation` VALUES (335, 10, 44);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='角色';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 'root');
INSERT INTO `role` VALUES (2, 'guest');
INSERT INTO `role` VALUES (5, '财务');
INSERT INTO `role` VALUES (6, '客服');
INSERT INTO `role` VALUES (7, '主管');
INSERT INTO `role` VALUES (8, '游客');
COMMIT;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `role_id` int(11) NOT NULL COMMENT '角色编号',
  `permission_id` int(11) NOT NULL COMMENT '权限编号',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `permission_id` (`permission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COMMENT='角色-权限';

-- ----------------------------
-- Records of role_permission
-- ----------------------------
BEGIN;
INSERT INTO `role_permission` VALUES (50, 5, 10);
INSERT INTO `role_permission` VALUES (51, 6, 10);
INSERT INTO `role_permission` VALUES (52, 7, 10);
INSERT INTO `role_permission` VALUES (55, 8, 10);
INSERT INTO `role_permission` VALUES (57, 1, 1);
INSERT INTO `role_permission` VALUES (58, 1, 10);
INSERT INTO `role_permission` VALUES (59, 2, 10);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
