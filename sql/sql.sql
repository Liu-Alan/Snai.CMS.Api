/*
DROP DATABASE snai_cms
;
*/
CREATE DATABASE snai_cms CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
;

USE snai_cms
;

CREATE TABLE admins(
	id INT AUTO_INCREMENT PRIMARY KEY,			-- 自增列需为主键
	user_name NVARCHAR(32) NOT NULL,
	`password` VARCHAR(32) NOT NULL,
	role_id INT NOT NULL,
	state TINYINT NOT NULL DEFAULT 1,			-- 1 启用，2 禁用
	create_time INT NOT NULL DEFAULT 0,
	update_time INT NOT NULL DEFAULT 0,
	last_logon_time INT NOT NULL DEFAULT 0,
	last_logon_ip NVARCHAR(45) not NULL DEFAULT '',
	error_logon_time INT NOT NULL DEFAULT 0,	-- 错误开始时间
	error_logon_count INT NOT NULL DEFAULT 0,	-- 错误次数
	lock_time INT NOT NULL DEFAULT 0 			-- 锁定结束时间
)
;

-- alter table admins add primary key pk_admins (id)  					-- 主键			
ALTER TABLE admins ADD UNIQUE INDEX ix_admins_user_name(user_name)  	-- UNIQUE INDEX 唯一索引
;

-- password:snai2024，保存时加盐md5(盐+password)
INSERT into admins(user_name,`password`,role_id,state,create_time)
VALUES('snai','86a6553dc16e23a7bb34b4306415c206',1,1,1550937600)
;

CREATE TABLE tokens(
	id INT AUTO_INCREMENT PRIMARY KEY,			-- 自增列需为主键
	token NVARCHAR(512) NOT NULL,
	user_id INT NOT NULL,
	state TINYINT NOT NULL DEFAULT 1,			-- 1 登录，2 退出
	create_time INT NOT NULL DEFAULT 0
)
;
		
ALTER TABLE tokens ADD INDEX ix_tokens_user_id_state(user_id,state)
;

CREATE TABLE modules(
	id INT AUTO_INCREMENT PRIMARY KEY,
	parent_id int not NULL DEFAULT 0,
	title NVARCHAR(32) not NULL,
	name NVARCHAR(32) not NULL DEFAULT '',		-- 用于前端判断是否显示，如创建，删除等
	router NVARCHAR(64) not NULL DEFAULT '',	-- api路由
	ui_router NVARCHAR(64) not NULL DEFAULT '', -- 前端路由
	menu TINYINT NOT NULL DEFAULT 2,			-- 是否菜单：1 是，2 否
	sort int not NULL DEFAULT 1,     			-- 小排在前
	state TINYINT NOT NULL DEFAULT 1			-- 1 启用，2 禁用
	
)
;

ALTER TABLE modules ADD UNIQUE INDEX ix_modules_parent_id_title(parent_id,title)  
;		
ALTER TABLE modules ADD INDEX ix_modules_router(router)  			
;

INSERT into modules(id,parent_id,title,router,sort,state,menu,ui_router,name)
select 1,-1,'修改密码','/api/home/changepassword',1,1,2,'',''
UNION ALL select 2,-1,'退出','/api/home/logout',1,1,2,'',''
UNION ALL select 3,0,'系统管理','',100,1,1,'','manage'
UNION ALL select 4,3,'账号管理','/api/admin/list',110,1,1,'/admins','admins'
UNION ALL select 5,4,'添加账号','/api/admin/add',1,1,2,'','addadmin'
UNION ALL select 6,4,'修改账号','/api/admin/update',1,1,2,'','updateadmin'
UNION ALL select 7,4,'删除账号','/api/admin/delete',1,1,2,'','deleteadmin'
UNION ALL select 8,4,'禁启用账号','/api/admin/endisable',1,1,2,'','endisableadmin'
UNION ALL select 9,4,'解锁账号','/api/admin/unlock',1,1,2,'','unlockadmin'
UNION ALL select 10,3,'模块管理','/api/module/list',120,1,1,'/modules','modules'
UNION ALL select 11,10,'添加模块','/api/module/add',1,1,2,'','addmodule'
UNION ALL select 12,10,'修改模块','/api/module/update',1,1,2,'','updatemodule'
UNION ALL select 13,10,'删除模块','/api/module/delete',1,1,2,'','deletemodule'
UNION ALL select 14,10,'禁启用模块','/api/module/endisable',1,1,2,'','endisablemodule'
UNION ALL select 15,3,'角色管理','/api/role/list',130,1,1,'/roles','roles'
UNION ALL select 16,15,'添加角色','/api/role/add',1,1,2,'','addrole'
UNION ALL select 17,15,'修改角色','/api/role/update',1,1,2,'','updaterole'
UNION ALL select 18,15,'删除角色','/api/role/delete',1,1,2,'','deleterole'
UNION ALL select 19,15,'禁启用角色','/api/role/endisable',1,1,2,'','endisablerole'
UNION ALL select 20,15,'分配权限','/api/role/assignperm',1,1,2,'','assignperm'

;

CREATE TABLE roles(
	id int AUTO_INCREMENT PRIMARY KEY,
	title NVARCHAR(32) not NULL,
	state TINYINT NOT NULL DEFAULT 1			-- 1 启用，2 禁用
)
;

ALTER TABLE roles ADD UNIQUE INDEX ix_roles_title(title)  		
;

INSERT into roles(title,state)
select '超级管理员',1
;

CREATE TABLE role_module(
	id int AUTO_INCREMENT PRIMARY KEY,
	role_id int not NULL,
	module_id int not NULL
)
;

ALTER TABLE role_module ADD UNIQUE INDEX ix_role_module_role_id_module_id(role_id,module_id) 
;

INSERT into role_module(role_id,module_id)
select 1,3
UNION ALL select 1,4
UNION ALL select 1,5
UNION ALL select 1,6
UNION ALL select 1,7
UNION ALL select 1,8
UNION ALL select 1,9
UNION ALL select 1,10
UNION ALL select 1,11
UNION ALL select 1,12
UNION ALL select 1,13
UNION ALL select 1,14
UNION ALL select 1,15
UNION ALL select 1,16
UNION ALL select 1,17
UNION ALL select 1,18
UNION ALL select 1,19
UNION ALL select 1,20
;