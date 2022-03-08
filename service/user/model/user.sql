CREATE TABLE IF NOT EXISTS `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
    `email` varchar(350) NOT NULL DEFAULT '' COMMENT '邮箱',
    `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;