CREATE DATABASE IF NOT EXISTS demo1 DEFAULT CHARSET utf8mb4 DEFAULT COLLATE 'utf8mb4_bin';
grant select, insert, update, delete on demo1.* to writer@localhost;
grant select, insert, update, delete on demo1.* to writer@'%qmgo%';
grant select on demo1.* to demo@localhost;
grant select on demo1.* to demo@'%qmgo%';
