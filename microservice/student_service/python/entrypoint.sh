#!/bin/sh

# ตรวจสอบสิทธิ์ของโฟลเดอร์และไฟล์
ls -la /mariadb/data/
ls -la /mariadb/initdb/

# เปลี่ยนเจ้าของโฟลเดอร์และไฟล์ให้เป็น user ที่เหมาะสม
chown -R 999:999 /mariadb/data/
chown -R 999:999 /mariadb/initdb/

exec "$@"
