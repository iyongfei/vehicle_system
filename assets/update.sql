
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


delimiter //
drop procedure if exists UpdateColumnUnlessExists;
create procedure UpdateColumnUnlessExists(
    IN dbName tinytext,
    IN tableName tinytext,
    IN fieldName tinytext,
    IN fieldDef text)
begin
    IF EXISTS (
            SELECT * FROM information_schema.COLUMNS
            WHERE column_name=fieldName
              and table_name=tableName
              and table_schema=dbName
        )
    THEN
        set @ddl=CONCAT('ALTER TABLE ',dbName,'.',tableName,
                        ' MODIFY ',fieldName,' ',fieldDef);
        prepare stmt from @ddl;
        execute stmt;
    END IF;
end;

call UpdateColumnUnlessExists('vehicle', 'gw_infos', 'gw_id', 'varchar(255) NULL DEFAULT NULL unique');
call UpdateColumnUnlessExists('vehicle', 'gw_infos', 'name', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'gw_infos', 'module', 'varchar(255) NULL DEFAULT NULL');

call UpdateColumnUnlessExists('vehicle', 'gw_info_uncerteds', 'gw_id', 'varchar(255) NULL DEFAULT NULL unique');
call UpdateColumnUnlessExists('vehicle', 'gw_info_uncerteds', 'name', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'gw_info_uncerteds', 'module', 'varchar(255) NULL DEFAULT NULL');

call UpdateColumnUnlessExists('vehicle', 'gw_assets', 'gw_id', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'gw_assets', 'asset_id', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'gw_assets', 'name', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'gw_assets', 'mac', 'varchar(255) NULL DEFAULT NULL');

call UpdateColumnUnlessExists('vehicle', 'threats', 'gw_id', 'varchar(255) NULL DEFAULT NULL');
call UpdateColumnUnlessExists('vehicle', 'threats', 'gw_asset_id', 'varchar(255) NULL DEFAULT NULL');

call UpdateColumnUnlessExists('vehicle', 'notifies', 'asset_name', 'varchar(255) NULL DEFAULT NULL');
//




