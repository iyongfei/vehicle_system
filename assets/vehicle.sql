
create Database IF NOT EXISTS `vehicle`;
USE vehicle;



CREATE TABLE IF NOT EXISTS `area_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `area_code` varchar(255) NULL DEFAULT NULL,
  `area_name` varchar(255) NULL DEFAULT NULL,
  `parent_area_code` varchar(255) NULL DEFAULT NULL,
   `tree_area_code` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_area_groups_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for flows
-- ----------------------------
CREATE TABLE IF NOT EXISTS `flows`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `flow_id` int(11) UNSIGNED NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `hash` int(11) UNSIGNED NULL DEFAULT NULL,
  `src_ip` int(11) UNSIGNED NULL DEFAULT NULL,
  `src_port` int(11) NULL DEFAULT NULL,
  `dst_ip` int(11) UNSIGNED NULL DEFAULT NULL,
  `dst_port` int(11) NULL DEFAULT NULL,

  `protocol` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `flow_info` varchar(255) NULL DEFAULT NULL,
  `safe_type` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `safe_info` varchar(255) NULL DEFAULT NULL,
  `start_time`  int(11) UNSIGNED NULL DEFAULT NULL,
  `last_seen_time` int(11)  UNSIGNED NULL DEFAULT NULL,
  `src_dst_bytes`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `dst_src_bytes`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `stat`  tinyint(3) UNSIGNED NULL DEFAULT NULL,
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `flow_id` (`flow_id`),

  INDEX `idx_flows_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for firmware_updates
-- ----------------------------
CREATE TABLE IF NOT EXISTS `firmware_updates`  (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,

    `deploy_id`      varchar(255) NULL DEFAULT NULL,
    `vehicle_id`          varchar(255) NULL DEFAULT NULL,

    `update_version` varchar(255) NULL DEFAULT NULL,
    `upgrade_timestamp`  int(11) NULL DEFAULT NULL,
    `upgrade_status`   tinyint(3) NULL DEFAULT NULL,
    `timeout`  int(11) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_firmware_updates_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for firmware_infos
-- ----------------------------

CREATE TABLE IF NOT EXISTS `firmware_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `system` varchar(255) NULL DEFAULT NULL,
  `version` varchar(255) NULL DEFAULT NULL,
  `soctype` varchar(255) NULL DEFAULT NULL,
  `hardware_model` varchar(255) NULL DEFAULT NULL,
  `md5` varchar(255) NULL DEFAULT NULL,
  `firmware_type` varchar(255) NULL DEFAULT NULL,
  `firmware_name` varchar(255) NULL DEFAULT NULL,
  `size` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_info` varchar(255) NULL DEFAULT NULL,
  `upload_time` timestamp NULL DEFAULT NULL,
  `bin_path` varchar(255) NULL DEFAULT NULL,


  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_firmware_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for vehicle_infos
-- ----------------------------
CREATE TABLE IF NOT EXISTS `vehicle_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,

  `version` varchar(255) NULL DEFAULT NULL,
  `start_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `firmware_version` varchar(255) NULL DEFAULT NULL,
  `hardware_model` varchar(255) NULL DEFAULT NULL,
  `module` varchar(255) NULL DEFAULT NULL,
  `supply_id` varchar(255) NULL DEFAULT NULL,
  `up_router_ip` varchar(255) NULL DEFAULT NULL,
  `ip` varchar(255) NULL DEFAULT NULL,
  `type` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `mac` varchar(255) NULL DEFAULT NULL,
  `time_stamp` int(11) NULL DEFAULT NULL,
  `hb_timeout` int(11) NULL DEFAULT NULL,
  `deploy_mode` tinyint(3) NULL DEFAULT NULL,
  flow_idle_time_slot int(11) NULL DEFAULT NULL,

  `online_status` tinyint(1) NULL DEFAULT NULL,
  `protect_status` tinyint(3) NULL DEFAULT NULL,
  `leader_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `vehicle_id` (`vehicle_id`),
  INDEX `idx_vehicle_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `vehicle_leaders`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `leader_id` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `phone` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gw_leaders_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


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


CREATE TABLE IF NOT EXISTS `assets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `vehicle_id` varchar(255) DEFAULT NULL,
  `asset_id` varchar(255) DEFAULT NULL,

  `ip` varchar(255) DEFAULT NULL,
  `mac` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `trade_mark` varchar(255) DEFAULT NULL,

  `online_status` tinyint(1) DEFAULT NULL,
  `last_online` int(11) unsigned DEFAULT NULL,

  `internet_switch` tinyint(1) DEFAULT NULL,
  `protect_status` tinyint(1) DEFAULT NULL,
  `lan_visit_switch` tinyint(1) DEFAULT NULL,

  `asset_group` varchar(255) DEFAULT NULL,
  `asset_leader` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `asset_id` (`asset_id`),
  KEY `idx_gw_assets_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for threats
-- ----------------------------
CREATE TABLE IF NOT EXISTS `threats`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `threat_id` varchar(255) NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `asset_id` varchar(255) NULL DEFAULT NULL,

  `type` tinyint(3) NULL DEFAULT NULL,
  `content` varchar(255) NULL DEFAULT NULL,
  `status` tinyint(3) NULL DEFAULT NULL,
  `attact_time` int(11) NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `threat_id` (`threat_id`),
  INDEX `idx_threats_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;




-- ----------------------------
-- Table structure for white_lists
-- ----------------------------
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



CREATE TABLE IF NOT EXISTS `port_maps`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `port_map_id` varchar(255) NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `src_port` varchar(255) NULL DEFAULT NULL,

  `dst_port` varchar(255) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,

  `switch` tinyint(1) NULL DEFAULT NULL,
  `protocol_type` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_port_maps_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for strategies
-- ----------------------------

CREATE TABLE IF NOT EXISTS `strategies`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `strategy_id` varchar(255)  NULL DEFAULT NULL,
  `type` tinyint(3) NULL DEFAULT NULL,
  `handle_mode` tinyint(3) NULL DEFAULT NULL,
  `enable` tinyint(1) NULL DEFAULT NULL,

  `name` varchar(255) NULL DEFAULT NULL,
  `introduce` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategies_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `strategy_vehicles`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `strategy_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_vehicles_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `strategy_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `strategy_id` varchar(255)  NULL DEFAULT NULL,
  `group_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_groups_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `strategy_groups_learning_results`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `group_id` varchar(255) NULL DEFAULT NULL,
  `learning_result_id` varchar(255)  NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_groups_learning_results_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for FlowStrategy
-- ----------------------------

CREATE TABLE IF NOT EXISTS `flow_strategies`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `flow_strategy_id` varchar(255)  NULL DEFAULT NULL,
  `type` tinyint(3) NULL DEFAULT NULL,
  `handle_mode` tinyint(3) NULL DEFAULT NULL,
  `enable` tinyint(1) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `introduce` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_strategies_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `flow_strategy_vehicles`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `flow_strategy_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_strategy_vehicles_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `flow_strategy_items`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `flow_strategy_item_id` varchar(255)  NULL DEFAULT NULL,
  `dst_ip` int(11) UNSIGNED NULL DEFAULT NULL,
  `dst_port` int(11) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_strategy_items_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `flow_strategy_relate_items`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `flow_strategy_id` varchar(255)  NULL DEFAULT NULL,
  `flow_strategy_item_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_strategy_relate_items_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `flow_strategy_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `flow_strategy_id` varchar(255)  NULL DEFAULT NULL,
  `group_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_strategy_groups_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;






CREATE TABLE IF NOT EXISTS `samples`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `sample_id` varchar(255) NULL DEFAULT NULL,
  `start_time` timestamp NULL DEFAULT NULL,
  `remain_time` int(11) NULL DEFAULT NULL,
  `total_time` int(11) NULL DEFAULT NULL,

  `status` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `timeout` int(11) UNSIGNED NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `introduce` varchar(255) NULL DEFAULT NULL,

  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `study_origin_id` varchar(255) NULL DEFAULT NULL,

  `check` tinyint(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_collect_samples_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `sample_items`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `sample_item_id` varchar(255) NULL DEFAULT NULL,
  `sample_id` varchar(255) NULL DEFAULT NULL,
  `src_mac` varchar(255) NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `src_port` int(11) UNSIGNED NULL DEFAULT NULL,

  `dst_ip` varchar(255) NULL DEFAULT NULL,
  `dst_port` int(11) UNSIGNED NULL DEFAULT NULL,
  `url` varchar(255) NULL DEFAULT NULL,
  `fetch_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sample_items_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `study_origins`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `study_origin_id` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_study_origins_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;





CREATE TABLE IF NOT EXISTS `automated_learnings`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `learning_id` varchar(255) NULL DEFAULT NULL,
  `sample_id` varchar(255) NULL DEFAULT NULL,
  `file_name` varchar(255) NULL DEFAULT NULL,
  `automated_learning_id` varchar(255) NULL DEFAULT NULL,
  `description` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_automated_learnings_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;






CREATE TABLE IF NOT EXISTS `automated_learning_results`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `learning_result_id` varchar(255) NULL DEFAULT NULL,
  `sample_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_automated_learning_results_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


