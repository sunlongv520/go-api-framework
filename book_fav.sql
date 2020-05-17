DROP TABLE IF EXISTS `book_fav`;
CREATE TABLE `book_fav`  (
  `item_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `book_id` int(11) NULL DEFAULT NULL COMMENT '图书ID',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户ID',
  `add_time` datetime NOT NULL COMMENT '收藏时间',
  PRIMARY KEY (`item_id`) USING BTREE,
  UNIQUE INDEX `book_id`(`book_id`, `user_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;