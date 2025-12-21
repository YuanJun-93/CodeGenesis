CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userAccount` varchar(256) NOT NULL COMMENT 'User Account',
  `userPassword` varchar(512) NOT NULL COMMENT 'Password',
  `userName` varchar(256) DEFAULT NULL COMMENT 'User Name',
  `userAvatar` varchar(1024) DEFAULT NULL COMMENT 'User Avatar',
  `userProfile` varchar(512) DEFAULT NULL COMMENT 'User Profile',
  `userRole` varchar(256) NOT NULL DEFAULT 'user' COMMENT 'User Role: user/admin',
  `editTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Edit Time',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update Time',
  `isDelete` tinyint NOT NULL DEFAULT '0' COMMENT 'Is Deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_userAccount` (`userAccount`),
  KEY `idx_userName` (`userName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='User';