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
	last_logon_ip NVARCHAR(15) not NULL DEFAULT '',
	error_logon_time INT NOT NULL DEFAULT 0,	-- 错误开始时间
	error_logon_count INT NOT NULL DEFAULT 0,	-- 错误次数
	lock_time INT NOT NULL DEFAULT 0 			-- 锁定结束时间
)
;

-- alter table admins add primary key pk_admins (id)  						-- 主键			
ALTER TABLE admins ADD UNIQUE INDEX ix_admins_user_name(user_name)  		-- UNIQUE INDEX 唯一索引
;

-- password:snai2024，保存时加盐md5(盐+password)
INSERT into admins(user_name,`password`,role_id,state,create_time)
VALUES('snai','86A6553DC16E23A7BB34B4306415C206',1,1,1550937600)
;

CREATE TABLE modules(
	id INT AUTO_INCREMENT PRIMARY KEY,
	parent_id int not NULL DEFAULT 0,
	title NVARCHAR(32) not NULL,
	controller NVARCHAR(32) not NULL DEFAULT '',
	action NVARCHAR(32) not NULL DEFAULT '',
	sort int not NULL DEFAULT 1,     			-- 小排在前
	state TINYINT NOT NULL DEFAULT 1			-- 1 启用，2 禁用
)
;

ALTER TABLE modules ADD UNIQUE INDEX ix_modules_parent_id_title(parent_id,title)  
;		
ALTER TABLE modules ADD INDEX ix_modules_controller_action(controller,action)  			
;

INSERT into modules(id,parent_id,title,controller,action,sort,state)
select 1,0,'首页','Home','Index',1,1
UNION ALL select 2,1,'首页','','',10,1
UNION ALL select 3,2,'登录信息','Home','LoginInfo',11,1
UNION ALL select 4,2,'修改密码','Home','UpdatePassword',12,1
UNION ALL select 5,0,'后台设置','BackManage','Index',2,1
UNION ALL select 6,5,'管理员管理','','',20,1
UNION ALL select 7,6,'账号管理','BackManage','AdminList',21,1
UNION ALL select 8,7,'添加修改账号','BackManage','ModifyAdmin',21,1
UNION ALL select 9,7,'禁启用账号','BackManage','UpdateAdminState',21,1
UNION ALL select 10,7,'解锁账号','BackManage','UnlockAdmin',21,1
UNION ALL select 11,7,'删除账号','BackManage','DeleteAdmin',21,1
UNION ALL select 12,6,'菜单管理','BackManage','ModuleList',22,1
UNION ALL select 13,12,'添加修改菜单','BackManage','ModifyModule',22,1
UNION ALL select 14,12,'禁启用菜单','BackManage','UpdateModuleState',22,1
UNION ALL select 15,12,'删除菜单','BackManage','DeleteModule',22,1
UNION ALL select 16,6,'角色管理','BackManage','RoleList',23,1
UNION ALL select 17,16,'添加修改角色','BackManage','ModifyRole',23,1
UNION ALL select 18,16,'禁启用角色','BackManage','UpdateRoleState',23,1
UNION ALL select 19,16,'删除角色','BackManage','DeleteRole',23,1
UNION ALL select 20,16,'分配权限','BackManage','ModifyRoleRight',23,1
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
select 1,1
UNION ALL select 1,2
UNION ALL select 1,3
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