
CREATE TABLE `user_info`
(
    `id`            int          NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name`          varchar(10)  NOT NULL COMMENT '用户名',
    `password`      varchar(100) NOT NULL COMMENT '密码',
    `ware_house_id` varchar(100) DEFAULT NULL COMMENT '仓库ID',
    `create_time`   timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `ware_house`
(
    `id`            int NOT NULL AUTO_INCREMENT COMMENT '仓库ID',
    `storage_space` int   DEFAULT NULL COMMENT '存储空间',
    `used_space`    float DEFAULT NULL COMMENT '使用空间',
    `user_id`       int NOT NULL COMMENT '所属用户ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `folder`
(
    `id`            int NOT NULL AUTO_INCREMENT COMMENT '文件夹ID',
    `name`          varchar(100) DEFAULT NULL COMMENT '名称',
    `ware_house_id` int          DEFAULT NULL COMMENT '仓库ID',
    `folder_id`     int          DEFAULT NULL COMMENT '父级文件夹ID',
    `included`      int          DEFAULT NULL COMMENT '包含的文件数',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `file`
(
    `id`            int          NOT NULL AUTO_INCREMENT COMMENT '文件ID',
    `name`          varchar(100) NOT NULL COMMENT '文件名',
    `md5`           varchar(100) NOT NULL COMMENT '文件md5',
    `path`          varchar(100) NOT NULL COMMENT '保存路径',
    `ware_house_id` int DEFAULT NULL COMMENT '仓库ID',
    `folder_id`     int DEFAULT NULL COMMENT '文件夹ID',
    `status`        tinyint(1) DEFAULT NULL COMMENT '文件上传状态',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
