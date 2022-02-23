/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : goadmin

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 23/02/2022 14:17:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '管理员',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `status` enum('Enable','Disable') NOT NULL DEFAULT 'Enable' COMMENT 'Enable/Disable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `remark` text NOT NULL COMMENT 'remark',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COMMENT='管理员';

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` (`id`, `name`, `password`, `status`, `created_at`, `updated_at`, `remark`) VALUES (1, 'root', '$2a$10$muCPVqgBylixJjYfdhJorOfauVad9ywpFU.zdy1.XaMIoIZyVTECG', 'Enable', '2021-02-18 15:24:46', '2021-02-18 15:24:46', '');
COMMIT;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `admin_id` int(11) NOT NULL COMMENT '账户编号',
  `role_id` int(11) NOT NULL COMMENT '角色编号',
  PRIMARY KEY (`id`),
  KEY `admin_id` (`admin_id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4 COMMENT='管理员-角色';

-- ----------------------------
-- Records of admin_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_role` (`id`, `admin_id`, `role_id`) VALUES (186, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `api_group` varchar(255) NOT NULL DEFAULT '' COMMENT '分组',
  `method` varchar(50) NOT NULL DEFAULT '' COMMENT '方法',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COMMENT='操作';

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (21, '查询管理员', 'admin', 'get', '/api/v1/admin/pages');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (22, '管理员信息', 'admin', 'get', '/api/v1/admin/info');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (23, '更新管理员', 'admin', 'post', '/api/v1/admin/update');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (24, '删除管理员', 'admin', 'post', '/api/v1/admin/delete');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (25, '添加管理员', 'admin', 'post', '/api/v1/admin/add');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (26, '查询角色', 'role', 'get', '/api/v1/role/pages');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (27, '角色信息', 'role', 'get', '/api/v1/role/info');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (28, '更新角色', 'role', 'post', '/api/v1/role/update');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (29, '添加角色', 'role', 'post', '/api/v1/role/add');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (30, '删除角色', 'role', 'post', '/api/v1/role/delete');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (31, '查询权限', 'permission', 'get', '/api/v1/permission/pages');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (32, '添加权限', 'permission', 'post', '/api/v1/permission/add');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (33, '更新权限', 'permission', 'post', '/api/v1/permission/update');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (34, '删除权限', 'permission', 'post', '/api/v1/permission/delete');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (35, '权限-菜单', 'permission', 'get', '/api/v1/permission_menu/find');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (36, '权限-API操作', 'permission', 'get', '/api/v1/permission_operation/find');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (37, '添加菜单', 'menu', 'post', '/api/v1/menu/add');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (38, '删除菜单', 'menu', 'post', '/api/v1/menu/delete');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (39, '更新菜单', 'menu', 'post', '/api/v1/menu/update');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (40, '添加API操作', 'api', 'post', '/api/v1/operation/add');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (41, '删除API操作', 'api', 'post', '/api/v1/operation/delete');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (42, '更新API操作', 'api', 'post', '/api/v1/operation/update');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (43, '查询登录日志', '', 'get', '/api/v1/log_login/pages');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (45, '菜单-ALL', '', 'get', '/api/v1/menu/all');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (46, 'API操作-ALL', '', 'get', '/api/v1/operation/all');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (47, '角色-ALL', '', 'get', '/api/v1/role/all');
INSERT INTO `api` (`id`, `name`, `api_group`, `method`, `url`) VALUES (48, '权限-ALL', '', 'get', '/api/v1/permission/all');
COMMIT;

-- ----------------------------
-- Table structure for log_admin_login
-- ----------------------------
DROP TABLE IF EXISTS `log_admin_login`;
CREATE TABLE `log_admin_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增编号',
  `admin_id` int(11) NOT NULL COMMENT '管理员编号',
  `name` varchar(36) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '管理员名',
  `location` varchar(100) NOT NULL DEFAULT '' COMMENT '位置',
  `os` text NOT NULL COMMENT '操作系统',
  `browser` text NOT NULL COMMENT '浏览器',
  `user_agent` text CHARACTER SET utf8 NOT NULL COMMENT '浏览器详情',
  `url` varchar(64) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT 'url',
  `ip` varchar(16) CHARACTER SET utf8 NOT NULL DEFAULT '0.0.0.0' COMMENT 'IP',
  `result` enum('SUCCESS','FAIL') NOT NULL DEFAULT 'SUCCESS',
  `record_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '记录时间',
  `remark` varchar(128) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_name` (`record_at`),
  KEY `ip` (`ip`,`record_at`),
  KEY `url` (`url`,`record_at`),
  KEY `record_time` (`record_at`),
  KEY `user_id` (`admin_id`,`record_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员登录日志';

-- ----------------------------
-- Records of log_admin_login
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for log_admin_operation
-- ----------------------------
DROP TABLE IF EXISTS `log_admin_operation`;
CREATE TABLE `log_admin_operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '操作编号',
  `admin_id` int(11) NOT NULL COMMENT '管理员编号',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '账户',
  `module` varchar(64) NOT NULL DEFAULT '' COMMENT '模块',
  `action` varchar(64) NOT NULL DEFAULT '' COMMENT '行为',
  `content` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '操作内容',
  `result` enum('SUCCESS','FAIL') NOT NULL DEFAULT 'SUCCESS',
  `message` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作消息',
  `ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作IP',
  `record_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`admin_id`,`action`,`record_at`),
  KEY `user_id_2` (`admin_id`,`record_at`),
  KEY `action_module` (`action`)
) ENGINE=InnoDB AUTO_INCREMENT=465 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员操作日志';

-- ----------------------------
-- Records of log_admin_operation
-- ----------------------------
BEGIN;
COMMIT;

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
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (1, 'Dashboard', 0, 'Dashboard', '/dashboard', 1);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (2, 'Management', 0, 'ManageAccountsIcon', '', 1);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (3, 'Admins', 2, '', '/admins', 1);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (4, 'Roles', 2, '', '/roles', 2);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (7, 'Permissions', 2, '', '/permissions', 3);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (8, 'Menus', 2, '', '/menus', 4);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (9, 'APIs', 2, '', '/myapis', 5);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (16, 'History', 27, 'FindInPageIcon', '/log/login', 88);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (23, 'Logs', 27, 'FindInPageIcon', '/log/operation', 99);
INSERT INTO `menu` (`id`, `name`, `pid`, `icon`, `url`, `index_sort`) VALUES (27, 'Logs', 0, '', '', 8);
COMMIT;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `permission_group` varchar(255) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='权限';

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` (`id`, `name`, `permission_group`, `remark`) VALUES (11, 'admin view', 'admin', 'demo');
INSERT INTO `permission` (`id`, `name`, `permission_group`, `remark`) VALUES (12, 'admin create', 'admin', 'nothing');
COMMIT;

-- ----------------------------
-- Table structure for permission_api
-- ----------------------------
DROP TABLE IF EXISTS `permission_api`;
CREATE TABLE `permission_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `api_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4 COMMENT='权限-API';

-- ----------------------------
-- Records of permission_api
-- ----------------------------
BEGIN;
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (43, 12, 21);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (44, 12, 22);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (45, 12, 24);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (46, 12, 23);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (47, 12, 25);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (66, 11, 21);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (67, 11, 22);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (68, 11, 26);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (69, 11, 27);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (70, 11, 31);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (71, 11, 43);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (72, 11, 48);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (73, 11, 47);
INSERT INTO `permission_api` (`id`, `permission_id`, `api_id`) VALUES (74, 11, 46);
COMMIT;

-- ----------------------------
-- Table structure for permission_menu
-- ----------------------------
DROP TABLE IF EXISTS `permission_menu`;
CREATE TABLE `permission_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8mb4 COMMENT='权限-menu';

-- ----------------------------
-- Records of permission_menu
-- ----------------------------
BEGIN;
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (103, 12, 2);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (104, 12, 1);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (105, 12, 4);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (124, 11, 23);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (125, 11, 16);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (126, 11, 27);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (127, 11, 2);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (128, 11, 3);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (129, 11, 4);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (130, 11, 7);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (131, 11, 8);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (132, 11, 9);
INSERT INTO `permission_menu` (`id`, `permission_id`, `menu_id`) VALUES (133, 11, 1);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COMMENT='角色';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (1, 'root', '');
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (2, 'guest', '');
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (3, '财务', '');
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (4, '客服', '');
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (15, 'ddd+3', 'jjjjj');
INSERT INTO `role` (`id`, `name`, `remark`) VALUES (16, 'ookk', 'dddddd');
COMMIT;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COMMENT='角色-权限';

-- ----------------------------
-- Records of role_permission
-- ----------------------------
BEGIN;
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (35, 15, 6);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (36, 15, 7);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (37, 15, 5);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (38, 15, 4);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (39, 15, 8);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (40, 15, 1);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (41, 15, 3);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (42, 15, 2);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (43, 16, 6);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (45, 18, 1);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (48, 2, 11);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (51, 0, 0);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (52, 1, 11);
INSERT INTO `role_permission` (`id`, `role_id`, `permission_id`) VALUES (53, 1, 12);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
