/*
 Navicat Premium Data Transfer

 Source Server         : myGo
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : mygo

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 03/01/2023 17:15:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `menuId` int(4) NOT NULL,
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `viewUrl` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `createTime` datetime NOT NULL,
  `updateTime` datetime NOT NULL,
  `level` int(2) NOT NULL,
  `pId` int(4) NOT NULL,
  `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`menuId`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '', NULL, '2022-11-15 10:33:11', '2022-11-15 10:33:14', 0, 0, '首页');
INSERT INTO `menu` VALUES (2, NULL, NULL, '2022-11-15 10:34:57', '2022-11-15 10:35:00', 0, 0, '系统管理');
INSERT INTO `menu` VALUES (3, NULL, NULL, '2022-11-15 10:35:36', '2022-11-15 10:35:38', 1, 2, '角色管理');
INSERT INTO `menu` VALUES (4, NULL, NULL, '2022-11-15 10:36:33', '2022-11-15 10:36:35', 1, 2, '菜单管理');
INSERT INTO `menu` VALUES (5, NULL, NULL, '2022-11-15 10:37:05', '2022-11-15 10:37:07', 1, 2, '用户管理');

SET FOREIGN_KEY_CHECKS = 1;
