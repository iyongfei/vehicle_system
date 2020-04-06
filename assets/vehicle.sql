
create Database IF NOT EXISTS `vehicle`;
USE vehicle;

-- ----------------------------
-- Table structure for users
-- ----------------------------
CREATE TABLE IF NOT EXISTS `users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `user_id` varchar(255) NULL DEFAULT NULL,
  `user_name` varchar(255) NULL DEFAULT NULL,
  `password` varchar(255) NULL DEFAULT NULL,
  `type` int(11) NULL DEFAULT NULL,
  `email` varchar(255) NULL DEFAULT NULL,
  `phone` varchar(255) NULL DEFAULT NULL,
  `marks` varchar(255) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `white_lists`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `white_list_id` varchar(255) NULL DEFAULT NULL,
  `dest_ip` varchar(255) NULL DEFAULT NULL,
  `url` varchar(255) NULL DEFAULT NULL,
  `source_mac` varchar(255) NULL DEFAULT NULL,
  `source_ip` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `white_list_id` (`white_list_id`),
  INDEX `idx_white_lists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `vehicle_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `bind_ip` varchar(255) NULL DEFAULT NULL,
  `mac` varchar(255) NULL DEFAULT NULL,
  `version` varchar(255) NULL DEFAULT NULL,
  `start_time` timestamp NULL DEFAULT NULL,
  `firmware_version` varchar(255) NULL DEFAULT NULL,
  `hardware_model` varchar(255) NULL DEFAULT NULL,
  `supply_id` varchar(255) NULL DEFAULT NULL,
  `up_router_ip` varchar(255) NULL DEFAULT NULL,
  `module` varchar(256) NULL DEFAULT NULL,
  `type` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `gw_timeout` int(11) NULL DEFAULT NULL,
  `online_status` tinyint(1) NULL DEFAULT NULL,
  `protect_status` tinyint(3) NULL DEFAULT NULL,
  `recent_active_time` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `deploy_mode` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `vehicle_id` (`vehicle_id`),
  INDEX `idx_vehicle_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `threats`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `threat_id` varchar(255) NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `type` int(11) NULL DEFAULT NULL,
  `content` varchar(255) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `attact_time` timestamp NULL DEFAULT NULL,
  `is_readed` tinyint(1) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
  `dest_type` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `threat_id` (`threat_id`),
  INDEX `idx_threats_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;