/*
Navicat MySQL Data Transfer

Source Server         : mysql57
Source Server Version : 50721
Source Host           : localhost:3307
Source Database       : swoft2

Target Server Type    : MYSQL
Target Server Version : 50721
File Encoding         : 65001

Date: 2019-12-02 17:25:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for book_metas
-- ----------------------------
DROP TABLE IF EXISTS `book_metas`;
CREATE TABLE `book_metas` (
  `meta_id` int(11) NOT NULL AUTO_INCREMENT,
  `meta_key` varchar(200) DEFAULT NULL,
  `meta_value` text,
  `item_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`meta_id`),
  UNIQUE KEY `meta_key` (`meta_key`,`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;
