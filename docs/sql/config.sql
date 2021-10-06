DROP DATABASE IF EXISTS `config_plaform`;
CREATE DATABASE `config_plaform`;
USE `config_plaform`;

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
  PRIMARY KEY (`id`, `project_user`),
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
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`username`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8;

insert into user(`username`, `password`, `nickname`, `gender`, `region`, `phone`, `email`, `photo`)
  values("superuser", "supertoken", "super", 1, "中国", "12306", "super@icloud.com", "https://baidu.com");