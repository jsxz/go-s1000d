CREATE TABLE `element` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `version` int(11) DEFAULT NULL COMMENT '4.0 1,\r\n4.1 2,\r\n4.2 4\r\n4.0,4.2=1+4=5',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin