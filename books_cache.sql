CREATE TABLE `books_cache` (
  `item_id` int(11) NOT NULL AUTO_INCREMENT,
  `cache_content` varchar(255) CHARACTER SET latin1 NOT NULL,
  `cache_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1代表是排行榜 其他待扩展',
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`item_id`) USING BTREE,
  UNIQUE KEY `cache_type` (`cache_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
