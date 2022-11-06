/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50734 (5.7.34)
 Source Host           : localhost:3306
 Source Schema         : cart

 Target Server Type    : MySQL
 Target Server Version : 50734 (5.7.34)
 File Encoding         : 65001

 Date: 06/11/2022 20:46:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '购物车id',
  `userId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `productId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `buyCount` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `checked` int(11) NOT NULL DEFAULT '0' COMMENT '是否选中,1=选中,0=未选',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `ix_user_id` (`userId`),
  KEY `ix_product_id` (`productId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';

SET FOREIGN_KEY_CHECKS = 1;
