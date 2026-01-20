CREATE TABLE `app_users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',

  `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间(软删)',

  -- 唯一身份识别码
  `casdoor_id` CHAR(36) DEFAULT NULL COMMENT 'Casdoor用户ID',
  `casdoor_owner` VARCHAR(100) NOT NULL COMMENT 'Casdoor组织名',
  `casdoor_name` VARCHAR(100) NOT NULL COMMENT 'Casdoor用户名/唯一标识',

  -- 基本信息
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT '昵称',
  `avatar` VARCHAR(500) DEFAULT NULL COMMENT '头像',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',

  -- 业务逻辑
  `source` VARCHAR(50) DEFAULT NULL COMMENT '来源: "google", "wechat", "password"',

  -- 登录统计
  `last_login_at` DATETIME(3) DEFAULT NULL COMMENT '最近登录时间',
  `login_count` INT UNSIGNED DEFAULT 0 COMMENT '登录次数',

  -- 唯一索引
  UNIQUE KEY `uk_casdoor_user` (`casdoor_owner`, `casdoor_name`),
  UNIQUE KEY `uk_casdoor_id` (`casdoor_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='C端用户表(对接Casdoor)';