
create Database IF NOT EXISTS `vehicle`;
USE vehicle;

-- ----------------------------
-- Table structure for area_groups
-- ----------------------------
CREATE TABLE IF NOT EXISTS `area_groups`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `area_code` varchar(255) NULL DEFAULT NULL,
  `area_name` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_area_groups_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;
