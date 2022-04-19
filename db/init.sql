CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(11) DEFAULT NULL,
  `password` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `nameindex` (`name`(10))
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
