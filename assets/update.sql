
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
//

CREATE TABLE IF NOT EXISTS `vehicle_auths`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  `vehicle_id` varchar(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_vehicle_auths_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8;



