#!/bin/bash

# 定义变量
DB_NAME="aigc-video"
DB_USER="root"
DB_PASS="123456"
DB_HOST="host.docker.internal"
DB_PORT="3306"

INIT_DB_TABLE="casbin_rule sys_apis sys_authorities sys_authority_menus sys_base_menus sys_data_authority_id sys_dictionaries sys_dictionary_details sys_export_templates sys_user_authority sys_users"
BASE_DB_TABLE="sys_apis sys_base_menus sys_base_menu_parameters sys_base_menu_btns sys_dictionaries sys_dictionary_details"
DEV_DB_TABLE="sys_auto_code_packages sys_auto_code_histories"

rm -rf sql/*.sql
mkdir -p sql
touch sql/base.config.sql
touch sql/dev.config.sql

# 初始化脚本，默认关掉（需要时取消注释）
# docker run --rm \
#    --platform linux/amd64 \
#    -e MYSQL_PWD="$DB_PASS" \
#    -v "$(pwd)/sql:/data" \
#    hub.roky.work/library/mysql:5.7.39 \
#    sh -c "mysqldump -h $DB_HOST -P $DB_PORT -u $DB_USER -t $DB_NAME $INIT_DB_TABLE --replace --skip-triggers" >> $(pwd)/sql/init.sql

# 使用docker run命令导出数据
docker run --rm \
    --platform linux/amd64 \
    -e MYSQL_PWD="$DB_PASS" \
    -v "$(pwd)/sql:/data" \
    hub.roky.work/library/mysql:5.7.39 \
    sh -c "mysqldump -h $DB_HOST -P $DB_PORT -u $DB_USER -t $DB_NAME $BASE_DB_TABLE --replace --skip-triggers" >> $(pwd)/sql/base.config.sql

docker run --rm \
    --platform linux/amd64 \
    -e MYSQL_PWD="$DB_PASS" \
    -v "$(pwd)/sql:/data" \
    hub.roky.work/library/mysql:5.7.39 \
    sh -c "mysqldump -h $DB_HOST -P $DB_PORT -u $DB_USER -t $DB_NAME $DEV_DB_TABLE --replace --skip-triggers" >> $(pwd)/sql/dev.config.sql