
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



CREATE TABLE IF NOT EXISTS `asset_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `area_code` varchar(255) NULL DEFAULT NULL,
  `area_name` varchar(255) NULL DEFAULT NULL,
  `parent_area_code` varchar(255) NULL DEFAULT NULL,
   `tree_area_code` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_asset_groups_deleted_at`(`deleted_at`) USING BTREE
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
  `asset_id` varchar(255) NULL DEFAULT NULL,
  `hash` int(11) UNSIGNED NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `src_port` int(11) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
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



CREATE TABLE IF NOT EXISTS `temp_flows`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `flow_id` int(11) UNSIGNED NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
   `asset_id` varchar(255) NULL DEFAULT NULL,
  `hash` int(11) UNSIGNED NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `src_port` int(11) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
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

  INDEX `idx_temp_flows_deleted_at`(`deleted_at`) USING BTREE
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
  `group_id` varchar(255) NULL DEFAULT NULL,
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



CREATE TABLE IF NOT EXISTS `asset_leaders`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `leader_id` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `phone` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_asset_leaders_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `assets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `vehicle_id` varchar(255) DEFAULT NULL,
  `asset_id` varchar(255) DEFAULT NULL,

  `ip` varchar(255) NULL DEFAULT NULL,
  `mac` varchar(255) NULL DEFAULT NULL,
  `name` varchar(255) NULL DEFAULT NULL,
  `trade_mark` varchar(255) NULL DEFAULT NULL,

  `online_status` tinyint(1) NULL DEFAULT NULL,
  `last_online` int(11) unsigned NULL DEFAULT NULL,

  `internet_switch` tinyint(1) NULL DEFAULT NULL,
  `protect_status` tinyint(1) NULL DEFAULT NULL,
  `lan_visit_switch` tinyint(1) NULL DEFAULT NULL,
  `access_net` tinyint(1) UNSIGNED NULL DEFAULT NULL,

  `asset_group` varchar(255) NULL DEFAULT NULL,
  `asset_leader` varchar(255) NULL DEFAULT NULL,
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

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategies_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `strategy_vehicles`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `strategy_vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `strategy_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_vehicles_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `strategy_vehicle_learning_results`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `strategy_vehicle_id` varchar(255) NULL DEFAULT NULL,
  `learning_result_id` varchar(255)  NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_vehicle_learning_results_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `strategy_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `strategy_group_id` varchar(255)  NULL DEFAULT NULL,
  `strategy_id` varchar(255)  NULL DEFAULT NULL,
  `group_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_strategy_groups_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;





CREATE TABLE IF NOT EXISTS `fstrategy_items`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fstrategy_item_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
  `dst_port` int(11) UNSIGNED NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fstrategy_items_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `fstrategies`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `fstrategy_id` varchar(255)  NULL DEFAULT NULL,
  `type` tinyint(3) NULL DEFAULT NULL,
  `handle_mode` tinyint(3) NULL DEFAULT NULL,
  `enable` tinyint(1) NULL DEFAULT NULL,
  `csv_path` varchar(255)  NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fstrategies_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `fstrategy_vehicles`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fstrategy_vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `fstrategy_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fstrategy_vehicles_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;




CREATE TABLE IF NOT EXISTS `fstrategy_vehicle_items`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fstrategy_vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `fstrategy_item_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fstrategy_vehicle_items_deleted_at`(`deleted_at`) USING BTREE
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
  INDEX `idx_samples_deleted_at`(`deleted_at`) USING BTREE
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
  `origin_id` varchar(255) NULL DEFAULT NULL,
  `origin_type` tinyint(3) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_automated_learning_results_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `flow_statistics`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `interface_name` varchar(255)  NULL DEFAULT NULL,

  `receivex`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `uploadx`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `flow_count`  int(11) UNSIGNED NULL DEFAULT NULL,
  `pub_flow`  int(11) UNSIGNED NULL DEFAULT NULL,
  `notlocal_flow`  int(11) UNSIGNED NULL DEFAULT NULL,
  `white_count`  int(11) UNSIGNED NULL DEFAULT NULL,


  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_flow_statistics_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `disks`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `monitor_id` varchar(255)  NULL DEFAULT NULL,
  `path` varchar(255)  NULL DEFAULT NULL,
  `disk_rate` double  NULL DEFAULT NULL,
   `gather_time` BIGINT(20) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_disks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `redis_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `monitor_id` varchar(255)  NULL DEFAULT NULL,
  `active` tinyint(1)  NULL DEFAULT NULL,
  `cpu_rate` double  NULL DEFAULT NULL,
  `mem_rate` double  NULL DEFAULT NULL,
  `mem` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
   `gather_time` BIGINT(20) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_redis_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;




CREATE TABLE IF NOT EXISTS `vhalo_nets`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

 `monitor_id` varchar(255)  NULL DEFAULT NULL,
  `active` tinyint(1)  NULL DEFAULT NULL,
  `cpu_rate` double  NULL DEFAULT NULL,
  `mem_rate` double  NULL DEFAULT NULL,
  `mem` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
   `gather_time` BIGINT(20) NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_vhalo_nets_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `alayer_protos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `proto_id` varchar(255)  NULL DEFAULT NULL,
  `protocol` varchar(255)  NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_alayer_protos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `categories`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `cate_id` varchar(255)  NULL DEFAULT NULL,
  `name` varchar(255)  NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_categories_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `finger_prints`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_id` varchar(255)  NULL DEFAULT NULL,
  `cate_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `device_mac` varchar(255)  NULL DEFAULT NULL,
  `flow_ids` varchar(255)  NULL DEFAULT NULL,
  `proto_rate`   varchar(500)  NULL DEFAULT NULL,
  `collect_type`   tinyint(3) UNSIGNED NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_finger_prints_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `fprint_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_info_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `device_mac` varchar(255)  NULL DEFAULT NULL,
   `trade_mark` varchar(255)  NULL DEFAULT NULL,
  `os`   varchar(255) NULL DEFAULT NULL,
  `dst_port`   int(11) NULL DEFAULT NULL,
  `examine_net` varchar(255)  NULL DEFAULT NULL,
  `access_net` tinyint(1) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fprint_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `fprint_info_actives`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_info_id` varchar(255)  NULL DEFAULT NULL,
  `os`   varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fprint_info_actives_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `fprint_info_passives`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_info_id` varchar(255)  NULL DEFAULT NULL,
  `dst_port`   int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fprint_info_passives_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `fprint_passive_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_info_id` varchar(255)  NULL DEFAULT NULL,
  `hash` int(11) UNSIGNED NULL DEFAULT NULL,
  `src_ip` varchar(255) NULL DEFAULT NULL,
  `src_port` int(11) NULL DEFAULT NULL,
  `dst_ip` varchar(255) NULL DEFAULT NULL,
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

  INDEX `idx_fprint_passive_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `white_assets`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `white_asset_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255)  NULL DEFAULT NULL,
  `device_mac` varchar(255)  NULL DEFAULT NULL,
   `trade_mark` varchar(255)  NULL DEFAULT NULL,
  `examine_net` varchar(255)  NULL DEFAULT NULL,
  `access_net` tinyint(1) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_white_assets_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

