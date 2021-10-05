CREATE DATABASE ConfigPlaform;
USE ConfigPlaform;

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
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`name`)
) ENGINE = InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET = utf8mb4 COMMENT = '项目信息';

insert into project(`name`, `description`, `department`, `admin`) values("配置管理平台", "可视化的编辑各种配置, 使开发更方便", "electronic information engineering", "liming, pengshengzi");

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(50) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;