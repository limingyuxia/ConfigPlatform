-- ----------------------------
-- Table structure for drug_source
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `name` varchar(32) NOT NULL COMMENT '项目名称',
  `description` varchar(32) COMMENT '项目描述',
  `department` varchar(32) COMMENT '项目所属的部门',
  `admin` varchar(32) COMMENT '管理员',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY(`name`)
) ENGINE = InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET = utf8mb4 COMMENT = '项目信息';
