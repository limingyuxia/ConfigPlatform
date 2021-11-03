DROP DATABASE IF EXISTS `config_platform`;
CREATE DATABASE `config_platform`;
USE `config_platform`;

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `name` varchar(128) NOT NULL COMMENT '项目名称',
  `description` varchar(128) COMMENT '项目描述',
  `department` varchar(128) COMMENT '项目所属的部门',
  `admin` varchar(128) COMMENT '管理员',
  `project_user` varchar(50) NOT NULL COMMENT '项目创建者',
  `develop_user` varchar(128) COMMENT '项目的开发人员',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`name`)
) ENGINE = InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET = utf8mb4 COMMENT = '项目信息';

insert into project(`name`, `description`, `department`, `admin`, `project_user`, `develop_user`) values("配置管理平台", 
"可视化的编辑各种配置, 使开发更方便", "北方民族大学/电气信息工程学院/电子信息工程", "liming, pengshengzi", "superuser", "superuser");

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `nickname` varchar(50) COMMENT '昵称',
  `gender` int(4) COMMENT '性别',
  `region` varchar(128) COMMENT '地址',
  `phone` varchar(32) COMMENT '电话',
  `email` varchar(32) COMMENT '邮箱',
  `photo` varchar(128) COMMENT '头像',
  `introduction` varchar(256) COMMENT '简介',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`username`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8;

insert into user(`username`, `password`, `nickname`, `gender`, `region`, `phone`, `email`, `photo`)
  values("superuser", "supertoken", "super", 1, "中国", "12306", "super@icloud.com", "https://baidu.com", "");

-- ----------------------------
-- Table structure for seaweedfs
-- ----------------------------
DROP TABLE IF EXISTS `seaweedfs`;
CREATE TABLE `seaweedfs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) NOT NULL COMMENT '文件所属的用户',
  `seaweedfs_url` varchar(128) NOT NULL COMMENT 'seaweedfs 文件路径',
  `server_url` varchar(128) NOT NULL COMMENT '文件对外访问的路径',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for auth2
-- ----------------------------
DROP TABLE IF EXISTS `auth2`;
CREATE TABLE `auth2` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` int(11) NOT NULL COMMENT '用户的auth2 app',
  `qq_openid` varchar(64) COMMENT 'QQ用户的openid',
  `qq_username` varchar(64) COMMENT 'QQ用户的昵称',
  `qq_avatar` varchar(128) COMMENT 'QQ用户的头像',
  `wechat_openid` varchar(64) COMMENT '微信用户的openid',
  `wechat_username` varchar(64) COMMENT '微信用户的昵称',
  `wechat_avatar` varchar(128) COMMENT '微信用户的头像',
  `uid` varchar(64) COMMENT '微博用户的uid',
  `weibo_username` varchar(64) COMMENT '微博用户的昵称',
  `weibo_avatar` varchar(128) COMMENT '微博用户的头像',
  `github_id` varchar(64) COMMENT 'github的用户id',
  `github_username` varchar(64) COMMENT 'github用户的昵称',
  `github_avatar` varchar(128) COMMENT 'github用户的头像',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8;