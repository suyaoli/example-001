/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50643
 Source Host           : localhost
 Source Database       : user_db

 Target Server Type    : MySQL
 Target Server Version : 50643
 File Encoding         : utf-8

 Date: 11/12/2019 12:07:39 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `tb_orders`
-- ----------------------------
DROP TABLE IF EXISTS `tb_orders`;
CREATE TABLE `tb_orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_sn` varchar(16) NOT NULL COMMENT '订单编号',
  `order_status` tinyint(4) NOT NULL COMMENT '订单状态(1-待付款，2-已取消，3-已付款，4-已关闭)',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Records of `tb_orders`
-- ----------------------------
BEGIN;
INSERT INTO `tb_orders` VALUES ('1', '201911120936', '1', '2019-11-11 21:36:30'), ('2', '201911100824', '2', '2019-11-11 21:36:53'), ('3', '201911090726', '1', '2019-11-11 21:41:06'), ('4', '201911100753', '3', '2019-11-11 21:41:55'), ('5', '201911230587', '4', '2019-11-11 21:42:10');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
