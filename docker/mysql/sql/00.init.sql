CREATE USER writer@localhost IDENTIFIED BY 'writer';
CREATE USER writer@'%qmgo%' IDENTIFIED BY 'writer';
CREATE USER demo@localhost IDENTIFIED BY 'demo';
CREATE USER demo@'%qmgo%' IDENTIFIED BY 'demo';
CREATE DATABASE IF NOT EXISTS demo DEFAULT CHARSET utf8mb4 DEFAULT COLLATE 'utf8mb4_bin';
grant select, insert, update, delete on demo.* to writer@localhost;
grant select, insert, update, delete on demo.* to writer@'%qmgo%';
grant select on demo.* to demo@localhost;
grant select on demo.* to demo@'%qmgo%';
