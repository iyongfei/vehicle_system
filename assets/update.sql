
USE vehicle;


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


delimiter //
drop procedure if exists AddColumnUnlessExists;
create procedure AddColumnUnlessExists(
    IN dbName tinytext,
    IN tableName tinytext,
    IN fieldName tinytext,
    IN fieldDef text)
begin
    IF NOT EXISTS (
            SELECT * FROM information_schema.COLUMNS
            WHERE column_name=fieldName
              and table_name=tableName
              and table_schema=dbName
        )
    THEN
        set @ddl=CONCAT('ALTER TABLE ',dbName,'.',tableName,
                        ' ADD COLUMN ',fieldName,' ',fieldDef);
        prepare stmt from @ddl;
        execute stmt;
    END IF;
end;
call AddColumnUnlessExists('vehicle', 'assets', 'access_net', 'tinyint(1) UNSIGNED NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'fstrategies', 'name', 'varchar(255)  NULL DEFAULT NULL');
//

