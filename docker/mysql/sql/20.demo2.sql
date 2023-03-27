CREATE DATABASE IF NOT EXISTS demo2 DEFAULT CHARSET utf8mb4 DEFAULT COLLATE 'utf8mb4_bin';
grant select, insert, update, delete on demo2.* to writer@localhost;
grant select, insert, update, delete on demo2.* to writer@'%qmgo%';
grant select on demo2.* to demo@localhost;
grant select on demo2.* to demo@'%qmgo%';
