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

 Date: 08/05/2021 21:13:25
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
INSERT INTO `admin` VALUES (11, '6578644', 'admin996', '$2a$10$U/stEM.FBeMCh3sdj1vOyu31IDebfEluJiabA4kS64dtP2RQlzixi', 1, '2021-05-08 21:13:07', '2021-05-08 21:13:07');
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
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COMMENT='管理员-角色';

-- ----------------------------
-- Records of admin_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_role` VALUES (61, '1', 1);
INSERT INTO `admin_role` VALUES (62, '1', 5);
INSERT INTO `admin_role` VALUES (63, '1', 8);
INSERT INTO `admin_role` VALUES (66, '', 2);
INSERT INTO `admin_role` VALUES (74, '6578644', 10);
INSERT INTO `admin_role` VALUES (75, '6578644', 5);
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
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='管理员登录日志';

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
INSERT INTO `log_admin_login` VALUES (62, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-25 19:44:06', '');
INSERT INTO `log_admin_login` VALUES (63, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-26 10:52:35', '');
INSERT INTO `log_admin_login` VALUES (64, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-26 18:53:03', '');
INSERT INTO `log_admin_login` VALUES (65, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-27 09:15:41', '');
INSERT INTO `log_admin_login` VALUES (66, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-27 19:15:26', '');
INSERT INTO `log_admin_login` VALUES (67, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-28 09:05:34', '');
INSERT INTO `log_admin_login` VALUES (68, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-29 12:27:45', '');
INSERT INTO `log_admin_login` VALUES (69, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-29 21:48:55', '');
INSERT INTO `log_admin_login` VALUES (70, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-30 14:31:30', '');
INSERT INTO `log_admin_login` VALUES (71, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-03-31 10:30:29', '');
INSERT INTO `log_admin_login` VALUES (72, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-01 10:37:10', '');
INSERT INTO `log_admin_login` VALUES (73, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-01 20:03:30', '');
INSERT INTO `log_admin_login` VALUES (74, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-02 11:40:04', '');
INSERT INTO `log_admin_login` VALUES (75, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-02 19:41:33', '');
INSERT INTO `log_admin_login` VALUES (76, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-02 19:41:33', '');
INSERT INTO `log_admin_login` VALUES (77, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome89.0.4389.90', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-04-04 08:50:49', '');
INSERT INTO `log_admin_login` VALUES (78, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome90.0.4430.93', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-05-06 13:33:13', '');
INSERT INTO `log_admin_login` VALUES (79, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome90.0.4430.93', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-05-07 07:56:10', '');
INSERT INTO `log_admin_login` VALUES (80, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome90.0.4430.93', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-05-08 10:49:22', '');
INSERT INTO `log_admin_login` VALUES (81, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome90.0.4430.93', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-05-08 15:07:52', '');
INSERT INTO `log_admin_login` VALUES (82, '1', 'root', '0 内网IP', 'Intel Mac OS X 10_15_7', 'Chrome90.0.4430.93', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36', '/api/v1/login', 1, '127.0.0.1', '2021-05-08 21:08:28', '');
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
) ENGINE=InnoDB AUTO_INCREMENT=302 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='管理员操作日志';

-- ----------------------------
-- Records of log_admin_operation
-- ----------------------------
BEGIN;
INSERT INTO `log_admin_operation` VALUES (253, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:广告管理;菜单编号:25;名称:广告管理;pid:24;icon:;url:advertise;index_sort:0;', 1, '', '127.0.0.1', '2021-03-26 11:14:11');
INSERT INTO `log_admin_operation` VALUES (254, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:广告管理;菜单编号:25;名称:广告管理;pid:24;icon:;url:/advertise;index_sort:1;', 1, '', '127.0.0.1', '2021-03-26 11:14:34');
INSERT INTO `log_admin_operation` VALUES (255, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:查询;权限编号:10;', 1, '', '127.0.0.1', '2021-03-26 14:49:27');
INSERT INTO `log_admin_operation` VALUES (256, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:查询;权限编号:10;', 1, '', '127.0.0.1', '2021-03-26 14:49:47');
INSERT INTO `log_admin_operation` VALUES (257, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:查询;权限编号:10;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:25;操作:添加管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-03-26 14:49:54');
INSERT INTO `log_admin_operation` VALUES (258, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:查询;权限编号:10;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-03-26 14:50:01');
INSERT INTO `log_admin_operation` VALUES (259, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:公告管理;菜单编号:26;名称:公告管理;pid:24;icon:;url:/announcements;index_sort:1;', 1, '', '127.0.0.1', '2021-03-27 14:54:57');
INSERT INTO `log_admin_operation` VALUES (260, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:广告管理;菜单编号:25;名称:广告管理;pid:24;icon:;url:/advertise;index_sort:2;', 1, '', '127.0.0.1', '2021-03-27 14:55:03');
INSERT INTO `log_admin_operation` VALUES (261, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:游戏管理;菜单编号:28;', 1, '', '127.0.0.1', '2021-03-27 16:41:06');
INSERT INTO `log_admin_operation` VALUES (262, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:系统权限;权限编号:1;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;菜单编号:24;菜单:内容管理;菜单编号:25;菜单:广告管理;菜单编号:26;菜单:公告管理;菜单编号:28;菜单:游戏管理;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:24;操作:删除管理员;操作编号:25;操作:添加管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:28;操作:更新角色;操作编号:29;操作:添加角色;操作编号:30;操作:删除角色;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:32;操作:添加权限;操作编号:33;操作:更新权限;操作编号:34;操作:删除权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:37;操作:添加菜单;操作编号:38;操作:删除菜单;操作编号:39;操作:更新菜单;操作编号:40;操作:添加API操作;操作编号:41;操作:删除API操作;操作编号:42;操作:更新API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-03-27 16:50:45');
INSERT INTO `log_admin_operation` VALUES (263, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:游戏管理;菜单编号:28;名称:游戏管理;pid:0;icon:;url:;index_sort:3;', 1, '', '127.0.0.1', '2021-03-27 16:51:06');
INSERT INTO `log_admin_operation` VALUES (264, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:游戏管理;菜单编号:28;名称:游戏管理;pid:0;icon:layui-icon-component;url:;index_sort:3;', 1, '', '127.0.0.1', '2021-03-27 16:51:16');
INSERT INTO `log_admin_operation` VALUES (265, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:游戏;菜单编号:29;', 1, '', '127.0.0.1', '2021-03-27 16:51:58');
INSERT INTO `log_admin_operation` VALUES (266, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:开奖;菜单编号:30;', 1, '', '127.0.0.1', '2021-03-27 16:52:09');
INSERT INTO `log_admin_operation` VALUES (267, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:玩法;菜单编号:31;', 1, '', '127.0.0.1', '2021-03-27 16:52:30');
INSERT INTO `log_admin_operation` VALUES (268, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:订单;菜单编号:32;', 1, '', '127.0.0.1', '2021-03-27 16:52:37');
INSERT INTO `log_admin_operation` VALUES (269, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:系统权限;权限编号:1;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;菜单编号:24;菜单:内容管理;菜单编号:25;菜单:广告管理;菜单编号:26;菜单:公告管理;菜单编号:28;菜单:游戏管理;菜单编号:29;菜单:游戏;菜单编号:30;菜单:开奖;菜单编号:31;菜单:玩法;菜单编号:32;菜单:订单;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:24;操作:删除管理员;操作编号:25;操作:添加管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:28;操作:更新角色;操作编号:29;操作:添加角色;操作编号:30;操作:删除角色;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:32;操作:添加权限;操作编号:33;操作:更新权限;操作编号:34;操作:删除权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:37;操作:添加菜单;操作编号:38;操作:删除菜单;操作编号:39;操作:更新菜单;操作编号:40;操作:添加API操作;操作编号:41;操作:删除API操作;操作编号:42;操作:更新API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-03-27 16:52:45');
INSERT INTO `log_admin_operation` VALUES (270, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:游戏;菜单编号:29;名称:游戏设置;pid:28;icon:;url:game;index_sort:0;', 1, '', '127.0.0.1', '2021-03-27 16:54:39');
INSERT INTO `log_admin_operation` VALUES (271, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:开奖;菜单编号:30;名称:赔率设置;pid:28;icon:;url:;index_sort:2;', 1, '', '127.0.0.1', '2021-03-27 16:55:00');
INSERT INTO `log_admin_operation` VALUES (272, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:游戏设置;菜单编号:29;名称:游戏设置;pid:28;icon:;url:game;index_sort:1;', 1, '', '127.0.0.1', '2021-03-27 16:55:07');
INSERT INTO `log_admin_operation` VALUES (273, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:订单;菜单编号:32;名称:订单管理;pid:28;icon:;url:;index_sort:3;', 1, '', '127.0.0.1', '2021-03-27 16:55:47');
INSERT INTO `log_admin_operation` VALUES (274, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:玩法;菜单编号:31;名称:开奖;pid:28;icon:;url:;index_sort:4;', 1, '', '127.0.0.1', '2021-03-27 16:56:08');
INSERT INTO `log_admin_operation` VALUES (275, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:开奖;菜单编号:31;名称:开奖设置;pid:28;icon:;url:;index_sort:4;', 1, '', '127.0.0.1', '2021-03-27 16:56:23');
INSERT INTO `log_admin_operation` VALUES (276, '1', 'root', 'add_menu', '添加菜单', '添加菜单:权菜单限:游戏统计;菜单编号:33;', 1, '', '127.0.0.1', '2021-03-27 16:57:12');
INSERT INTO `log_admin_operation` VALUES (277, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:系统权限;权限编号:1;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;菜单编号:24;菜单:内容管理;菜单编号:25;菜单:广告管理;菜单编号:26;菜单:公告管理;菜单编号:28;菜单:游戏管理;菜单编号:29;菜单:游戏设置;菜单编号:30;菜单:赔率设置;菜单编号:31;菜单:开奖设置;菜单编号:32;菜单:订单管理;菜单编号:33;菜单:游戏统计;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:24;操作:删除管理员;操作编号:25;操作:添加管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:28;操作:更新角色;操作编号:29;操作:添加角色;操作编号:30;操作:删除角色;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:32;操作:添加权限;操作编号:33;操作:更新权限;操作编号:34;操作:删除权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:37;操作:添加菜单;操作编号:38;操作:删除菜单;操作编号:39;操作:更新菜单;操作编号:40;操作:添加API操作;操作编号:41;操作:删除API操作;操作编号:42;操作:更新API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-03-27 16:57:20');
INSERT INTO `log_admin_operation` VALUES (278, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:游戏统计;菜单编号:33;名称:游戏统计;pid:28;icon:;url:;index_sort:5;', 1, '', '127.0.0.1', '2021-03-27 16:58:41');
INSERT INTO `log_admin_operation` VALUES (279, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:开奖设置;菜单编号:31;名称:开奖设置;pid:28;icon:;url:/game_result;index_sort:4;', 1, '', '127.0.0.1', '2021-04-01 14:57:55');
INSERT INTO `log_admin_operation` VALUES (280, '1', 'root', 'add_role', '添加角色', '添加角色:角色:;角色编号:0;', 1, '', '127.0.0.1', '2021-05-08 15:01:48');
INSERT INTO `log_admin_operation` VALUES (281, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:游戏统计;菜单编号:33;', 1, '', '127.0.0.1', '2021-05-08 20:31:24');
INSERT INTO `log_admin_operation` VALUES (282, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:订单管理;菜单编号:32;', 1, '', '127.0.0.1', '2021-05-08 20:31:27');
INSERT INTO `log_admin_operation` VALUES (283, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:开奖设置;菜单编号:31;', 1, '', '127.0.0.1', '2021-05-08 20:31:30');
INSERT INTO `log_admin_operation` VALUES (284, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:赔率设置;菜单编号:30;', 1, '', '127.0.0.1', '2021-05-08 20:31:32');
INSERT INTO `log_admin_operation` VALUES (285, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:游戏设置;菜单编号:29;', 1, '', '127.0.0.1', '2021-05-08 20:31:35');
INSERT INTO `log_admin_operation` VALUES (286, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:游戏管理;菜单编号:28;', 1, '', '127.0.0.1', '2021-05-08 20:31:38');
INSERT INTO `log_admin_operation` VALUES (287, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:公告管理;菜单编号:26;', 1, '', '127.0.0.1', '2021-05-08 20:31:44');
INSERT INTO `log_admin_operation` VALUES (288, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:广告管理;菜单编号:25;', 1, '', '127.0.0.1', '2021-05-08 20:31:48');
INSERT INTO `log_admin_operation` VALUES (289, '1', 'root', 'delete_menu', '删除菜单', '删除菜单:菜单:内容管理;菜单编号:24;', 1, '', '127.0.0.1', '2021-05-08 20:31:53');
INSERT INTO `log_admin_operation` VALUES (290, '1', 'root', 'permission_update', '修改权限', '修改权限:权限:系统权限;权限编号:1;菜单编号:1;菜单:系统管理;菜单编号:2;菜单:管理员管理;菜单编号:3;菜单:角色管理;菜单编号:7;菜单:权限管理;菜单编号:8;菜单:资源管理;菜单编号:16;菜单:登录日志;菜单编号:23;菜单:操作日志;操作编号:21;操作:查询管理员;操作编号:22;操作:管理员信息;操作编号:23;操作:更新管理员;操作编号:24;操作:删除管理员;操作编号:25;操作:添加管理员;操作编号:47;操作:角色-ALL;操作编号:26;操作:查询角色;操作编号:27;操作:角色信息;操作编号:28;操作:更新角色;操作编号:29;操作:添加角色;操作编号:30;操作:删除角色;操作编号:48;操作:权限-ALL;操作编号:31;操作:查询权限;操作编号:32;操作:添加权限;操作编号:33;操作:更新权限;操作编号:34;操作:删除权限;操作编号:35;操作:权限-菜单;操作编号:36;操作:权限-API操作;操作编号:37;操作:添加菜单;操作编号:38;操作:删除菜单;操作编号:39;操作:更新菜单;操作编号:40;操作:添加API操作;操作编号:41;操作:删除API操作;操作编号:42;操作:更新API操作;操作编号:45;操作:菜单-ALL;操作编号:46;操作:API操作-ALL;操作编号:43;操作:查询登录日志;操作编号:44;操作:查询操作日志;', 1, '', '127.0.0.1', '2021-05-08 20:32:06');
INSERT INTO `log_admin_operation` VALUES (291, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:管理员管理;菜单编号:2;名称:管理员;pid:1;icon:layui-icon-user;url:/admin;index_sort:1;', 1, '', '127.0.0.1', '2021-05-08 20:55:36');
INSERT INTO `log_admin_operation` VALUES (292, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:管理员;菜单编号:2;名称:用户管理;pid:1;icon:layui-icon-user;url:/admin;index_sort:1;', 1, '', '127.0.0.1', '2021-05-08 20:56:19');
INSERT INTO `log_admin_operation` VALUES (293, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:权限管理;菜单编号:7;名称:权限管理;pid:1;icon:layui-icon-auz;url:/permission;index_sort:3;', 1, '', '127.0.0.1', '2021-05-08 21:02:45');
INSERT INTO `log_admin_operation` VALUES (294, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:角色管理;菜单编号:3;名称:角色管理;pid:1;icon:layui-icon-username;url:/role;index_sort:2;', 1, '', '127.0.0.1', '2021-05-08 21:03:12');
INSERT INTO `log_admin_operation` VALUES (295, '1', 'root', 'update_menu', '更新菜单', '修改菜单:菜单:资源管理;菜单编号:8;名称:资源管理;pid:1;icon:layui-icon-form;url:/resources;index_sort:4;', 1, '', '127.0.0.1', '2021-05-08 21:04:21');
INSERT INTO `log_admin_operation` VALUES (296, '1', 'root', 'delete_admin', '删除账户', '删除管理员:账户:demo996;账户编号:9;', 1, '', '127.0.0.1', '2021-05-08 21:10:09');
INSERT INTO `log_admin_operation` VALUES (297, '1', 'root', 'add_admin', '添加账户', '添加管理员:账户:;账户编号:0;', 1, '', '127.0.0.1', '2021-05-08 21:10:49');
INSERT INTO `log_admin_operation` VALUES (298, '1', 'root', 'delete_admin', '删除账户', '删除管理员:账户:test996;账户编号:10;', 1, '', '127.0.0.1', '2021-05-08 21:11:09');
INSERT INTO `log_admin_operation` VALUES (299, '1', 'root', 'delete_role', '删除角色', '删除角色:角色:aa;角色编号:9;', 1, '', '127.0.0.1', '2021-05-08 21:11:16');
INSERT INTO `log_admin_operation` VALUES (300, '1', 'root', 'add_role', '添加角色', '添加角色:角色:;角色编号:0;', 1, '', '127.0.0.1', '2021-05-08 21:12:24');
INSERT INTO `log_admin_operation` VALUES (301, '1', 'root', 'add_admin', '添加账户', '添加管理员:账户:;账户编号:0;', 1, '', '127.0.0.1', '2021-05-08 21:13:07');
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
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COMMENT='菜单';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, '系统管理', 0, 'layui-icon-component', '', 1);
INSERT INTO `menu` VALUES (2, '用户管理', 1, 'layui-icon-user', '/admin', 1);
INSERT INTO `menu` VALUES (3, '角色管理', 1, 'layui-icon-username', '/role', 2);
INSERT INTO `menu` VALUES (7, '权限管理', 1, 'layui-icon-auz', '/permission', 3);
INSERT INTO `menu` VALUES (8, '资源管理', 1, 'layui-icon-form', '/resources', 4);
INSERT INTO `menu` VALUES (16, '登录日志', 1, 'layui-icon-list', '/log/login', 5);
INSERT INTO `menu` VALUES (23, '操作日志', 1, 'layui-icon-list', '/log/operation', 6);
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
) ENGINE=InnoDB AUTO_INCREMENT=235 DEFAULT CHARSET=utf8mb4 COMMENT='权限-菜单';

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
INSERT INTO `permission_menu` VALUES (179, 10, 1);
INSERT INTO `permission_menu` VALUES (180, 10, 2);
INSERT INTO `permission_menu` VALUES (181, 10, 3);
INSERT INTO `permission_menu` VALUES (182, 10, 7);
INSERT INTO `permission_menu` VALUES (183, 10, 8);
INSERT INTO `permission_menu` VALUES (184, 10, 16);
INSERT INTO `permission_menu` VALUES (185, 10, 23);
INSERT INTO `permission_menu` VALUES (228, 1, 1);
INSERT INTO `permission_menu` VALUES (229, 1, 2);
INSERT INTO `permission_menu` VALUES (230, 1, 3);
INSERT INTO `permission_menu` VALUES (231, 1, 7);
INSERT INTO `permission_menu` VALUES (232, 1, 8);
INSERT INTO `permission_menu` VALUES (233, 1, 16);
INSERT INTO `permission_menu` VALUES (234, 1, 23);
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
) ENGINE=InnoDB AUTO_INCREMENT=477 DEFAULT CHARSET=utf8mb4 COMMENT='权限-操作';

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
INSERT INTO `permission_operation` VALUES (351, 10, 21);
INSERT INTO `permission_operation` VALUES (352, 10, 22);
INSERT INTO `permission_operation` VALUES (353, 10, 23);
INSERT INTO `permission_operation` VALUES (354, 10, 47);
INSERT INTO `permission_operation` VALUES (355, 10, 26);
INSERT INTO `permission_operation` VALUES (356, 10, 27);
INSERT INTO `permission_operation` VALUES (357, 10, 48);
INSERT INTO `permission_operation` VALUES (358, 10, 31);
INSERT INTO `permission_operation` VALUES (359, 10, 35);
INSERT INTO `permission_operation` VALUES (360, 10, 36);
INSERT INTO `permission_operation` VALUES (361, 10, 45);
INSERT INTO `permission_operation` VALUES (362, 10, 46);
INSERT INTO `permission_operation` VALUES (363, 10, 43);
INSERT INTO `permission_operation` VALUES (364, 10, 44);
INSERT INTO `permission_operation` VALUES (449, 1, 21);
INSERT INTO `permission_operation` VALUES (450, 1, 22);
INSERT INTO `permission_operation` VALUES (451, 1, 23);
INSERT INTO `permission_operation` VALUES (452, 1, 24);
INSERT INTO `permission_operation` VALUES (453, 1, 25);
INSERT INTO `permission_operation` VALUES (454, 1, 47);
INSERT INTO `permission_operation` VALUES (455, 1, 26);
INSERT INTO `permission_operation` VALUES (456, 1, 27);
INSERT INTO `permission_operation` VALUES (457, 1, 28);
INSERT INTO `permission_operation` VALUES (458, 1, 29);
INSERT INTO `permission_operation` VALUES (459, 1, 30);
INSERT INTO `permission_operation` VALUES (460, 1, 48);
INSERT INTO `permission_operation` VALUES (461, 1, 31);
INSERT INTO `permission_operation` VALUES (462, 1, 32);
INSERT INTO `permission_operation` VALUES (463, 1, 33);
INSERT INTO `permission_operation` VALUES (464, 1, 34);
INSERT INTO `permission_operation` VALUES (465, 1, 35);
INSERT INTO `permission_operation` VALUES (466, 1, 36);
INSERT INTO `permission_operation` VALUES (467, 1, 37);
INSERT INTO `permission_operation` VALUES (468, 1, 38);
INSERT INTO `permission_operation` VALUES (469, 1, 39);
INSERT INTO `permission_operation` VALUES (470, 1, 40);
INSERT INTO `permission_operation` VALUES (471, 1, 41);
INSERT INTO `permission_operation` VALUES (472, 1, 42);
INSERT INTO `permission_operation` VALUES (473, 1, 45);
INSERT INTO `permission_operation` VALUES (474, 1, 46);
INSERT INTO `permission_operation` VALUES (475, 1, 43);
INSERT INTO `permission_operation` VALUES (476, 1, 44);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='角色';

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
INSERT INTO `role` VALUES (10, 'admin');
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
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COMMENT='角色-权限';

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
INSERT INTO `role_permission` VALUES (61, 10, 2);
INSERT INTO `role_permission` VALUES (62, 10, 10);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
