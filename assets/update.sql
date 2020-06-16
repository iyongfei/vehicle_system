
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


