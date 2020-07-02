
USE vehicle;

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

call AddColumnUnlessExists('vehicle', 'flows', 'src_dst_packets', 'BIGINT(20) UNSIGNED NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'dst_src_packets', 'BIGINT(20) UNSIGNED NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'host_name', 'varchar(255) NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'category', 'int(11) NULL DEFAULT NULL,');
call AddColumnUnlessExists('vehicle', 'flows', 'has_passive', 'tinyint(1) NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'iat_flow_avg', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'iat_flow_stddev', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'data_ratio', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'str_data_ratio', 'tinyint(3) UNSIGNED NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'pktlen_c_to_s_avg', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'pktlen_c_to_s_stddev', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'pktlen_s_to_c_avg', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'pktlen_s_to_c_stddev', 'double  NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'tls_client_info', 'varchar(255) NULL DEFAULT NULL');
call AddColumnUnlessExists('vehicle', 'flows', 'ja3c', 'varchar(255) NULL DEFAULT NULL');
//



CREATE TABLE IF NOT EXISTS `fprint_flows`  (
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

   `src_dst_packets`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
   `dst_src_packets`  BIGINT(20) UNSIGNED NULL DEFAULT NULL,
   `host_name` varchar(255) NULL DEFAULT NULL,
   `category` int(11) NULL DEFAULT NULL,
   `has_passive` tinyint(1) NULL DEFAULT NULL,
   `iat_flow_avg` double  NULL DEFAULT NULL,
   `iat_flow_stddev` double  NULL DEFAULT NULL,
   `data_ratio` double  NULL DEFAULT NULL,
   `str_data_ratio` tinyint(3) UNSIGNED NULL DEFAULT NULL,
   `pktlen_c_to_s_avg` double  NULL DEFAULT NULL,
   `pktlen_c_to_s_stddev` double  NULL DEFAULT NULL,
   `pktlen_s_to_c_avg` double  NULL DEFAULT NULL,
   `pktlen_s_to_c_stddev` double  NULL DEFAULT NULL,
   `tls_client_info` varchar(255) NULL DEFAULT NULL,
   `ja3c` varchar(255) NULL DEFAULT NULL,

   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE KEY `flow_id` (`flow_id`),

  INDEX `idx_fprint_flows_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `fprints`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `fprint_id` varchar(255)  NULL DEFAULT NULL,
  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  `asset_id` varchar(255) NULL DEFAULT NULL,

  `collect_time` int(11) NULL DEFAULT NULL,
  `collect_time_rate` double NULL DEFAULT NULL,

  `collect_proto_flows`  varchar(1000) NULL DEFAULT NULL,
  `collect_proto_rate` double NULL DEFAULT NULL,

  `categorys` varchar(255) NULL DEFAULT NULL,
--   `categorys_rate` double NULL DEFAULT NULL,

  `collect_host` varchar(255) NULL DEFAULT NULL,
  `collect_host_rate` double NULL DEFAULT NULL,

  `collect_tls` varchar(255) NULL DEFAULT NULL,
  `collect_tls_rate` double NULL DEFAULT NULL,

  `collect_bytes` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `collect_bytes_rate` double NULL DEFAULT NULL,

  `collect_total_rate` double NULL DEFAULT NULL,



  `collect_start` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
  `collect_finish` tinyint(1) NULL DEFAULT NULL,
  `auto_cate_id` varchar(255)  NULL DEFAULT NULL,
  `auto_cate_rate` double NULL DEFAULT NULL,

   PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fprints_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `asset_fprints`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `asset_fprint_id` varchar(255)  NULL DEFAULT NULL,
  `asset_id` varchar(255) NULL DEFAULT NULL,
  `cate_id` varchar(255)  NULL DEFAULT NULL,

   PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_asset_fprints_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;


