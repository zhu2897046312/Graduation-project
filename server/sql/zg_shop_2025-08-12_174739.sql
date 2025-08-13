-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: 172.25.13.23    Database: zg_shop
-- ------------------------------------------------------
-- Server version	9.0.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cms_associated_place`
--

DROP TABLE IF EXISTS `cms_associated_place`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_associated_place` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '地点名称',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '地点代码',
  `full_pinyin` varchar(255) DEFAULT NULL COMMENT '地点名称全拼音',
  `initial_pinyin` varchar(255) DEFAULT NULL COMMENT '地点名称首字母拼音',
  `thumb_img` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `thumb_video` varchar(200) NOT NULL DEFAULT '' COMMENT '封面视频',
  `description` text COMMENT '地点简介',
  `content` text COMMENT '地点内容',
  `score` int NOT NULL DEFAULT '0' COMMENT '评分',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='地点信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_category`
--

DROP TABLE IF EXISTS `cms_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '栏目名称',
  `code` varchar(200) NOT NULL DEFAULT '' COMMENT '编码',
  `pid` bigint NOT NULL DEFAULT '0' COMMENT '上级栏目',
  `category_tpl` varchar(200) NOT NULL DEFAULT '' COMMENT '栏目模板',
  `cont_tpl` varchar(200) NOT NULL DEFAULT '' COMMENT '文档模板',
  `keyword` varchar(200) NOT NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `category_type` tinyint NOT NULL DEFAULT '0' COMMENT '类型:1=列表;2=单页面;3=外部链接',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_comment`
--

DROP TABLE IF EXISTS `cms_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `document_id` bigint NOT NULL DEFAULT '0' COMMENT '文档ID',
  `reply_id` bigint NOT NULL DEFAULT '0' COMMENT '回复顶级评论ID',
  `reply_to_reply_id` bigint NOT NULL DEFAULT '0' COMMENT '回复上级评论ID',
  `reply_to_username` varchar(200) NOT NULL DEFAULT '' COMMENT '回复上级评论人名称',
  `reply_to_avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '回复上级评论人头像',
  `reply_to_userid` bigint NOT NULL DEFAULT '0' COMMENT '回复上级评论人ID',
  `text` text NOT NULL COMMENT '评论内容',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已审核;2=待审核;3=审核不通过',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `ip` varchar(200) NOT NULL DEFAULT '' COMMENT 'ip',
  `comment_reply_total` int NOT NULL DEFAULT '0' COMMENT '评论回复数量',
  `like_num` int NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `bad_num` int NOT NULL DEFAULT '0' COMMENT '不喜欢数量',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_document`
--

DROP TABLE IF EXISTS `cms_document`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_document` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '文档名称',
  `category_id` bigint NOT NULL DEFAULT '0' COMMENT '所属栏目',
  `associated_place_id` bigint NOT NULL DEFAULT '0' COMMENT '关联地点ID',
  `cont_tpl` varchar(200) NOT NULL DEFAULT '' COMMENT '文档模板',
  `keyword` varchar(200) NOT NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `link_type` tinyint NOT NULL DEFAULT '0' COMMENT '链接类型:1=内部文档;2=外部链接',
  `document_type` tinyint NOT NULL DEFAULT '0' COMMENT '文档类型:1=图文;2=视频',
  `video_duration` int NOT NULL DEFAULT '0' COMMENT '视频时长(秒)',
  `send_time` datetime DEFAULT NULL COMMENT '发布时间',
  `author` varchar(200) NOT NULL DEFAULT '' COMMENT '作者',
  `source` varchar(200) NOT NULL DEFAULT '' COMMENT '来源',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '发布用户',
  `admin_id` bigint NOT NULL DEFAULT '0' COMMENT '发布管理员',
  `read_num` int NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `like_num` int NOT NULL DEFAULT '0' COMMENT '喜欢数量',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `code` varchar(512) DEFAULT NULL COMMENT '文档编码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文档主表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_document_archive`
--

DROP TABLE IF EXISTS `cms_document_archive`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_document_archive` (
  `document_id` bigint NOT NULL,
  `cont` mediumtext COMMENT '正文内容',
  `download_files` json DEFAULT NULL COMMENT '附件列表',
  PRIMARY KEY (`document_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图文内容表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_document_tag`
--

DROP TABLE IF EXISTS `cms_document_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_document_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `document_id` bigint NOT NULL DEFAULT '0' COMMENT '文档ID',
  `tag_id` bigint NOT NULL DEFAULT '0' COMMENT '标签ID',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签文档索引表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_document_video`
--

DROP TABLE IF EXISTS `cms_document_video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_document_video` (
  `document_id` bigint NOT NULL,
  `video_path` varchar(200) NOT NULL DEFAULT '' COMMENT '视频链接',
  PRIMARY KEY (`document_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图文内容表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_file`
--

DROP TABLE IF EXISTS `cms_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_file` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `file_name` varchar(200) NOT NULL DEFAULT '' COMMENT '文件名称',
  `file_path` varchar(200) NOT NULL DEFAULT '' COMMENT '文件路径',
  `file_size` int NOT NULL DEFAULT '0' COMMENT '文件大小',
  `file_type` varchar(50) NOT NULL DEFAULT '' COMMENT '文件类型',
  `file_md5` varchar(200) NOT NULL DEFAULT '' COMMENT '文件md5',
  `file_ext` varchar(50) NOT NULL DEFAULT '' COMMENT '文件后缀',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_recommend`
--

DROP TABLE IF EXISTS `cms_recommend`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_recommend` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '推荐位名称',
  `code` varchar(200) NOT NULL DEFAULT '' COMMENT '编码',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `more_link` varchar(200) NOT NULL DEFAULT '' COMMENT '更多链接',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='推荐位表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_recommend_index`
--

DROP TABLE IF EXISTS `cms_recommend_index`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_recommend_index` (
  `id` int NOT NULL AUTO_INCREMENT,
  `recommend_id` int NOT NULL DEFAULT '0' COMMENT '所属推荐位',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '推荐位名称',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `link` varchar(200) NOT NULL DEFAULT '' COMMENT '链接',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `product_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
  `document_id` int NOT NULL DEFAULT '0' COMMENT '文档id',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_rid` (`recommend_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='推荐位索引表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_scenic_spot`
--

DROP TABLE IF EXISTS `cms_scenic_spot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_scenic_spot` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `associated_place_id` bigint NOT NULL DEFAULT '0' COMMENT '关联地点ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '景区名称',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '景区代码',
  `full_pinyin` varchar(255) DEFAULT NULL COMMENT '景区名称全拼音',
  `initial_pinyin` varchar(255) DEFAULT NULL COMMENT '景区名称首字母拼音',
  `thumb_img` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `thumb_video` varchar(200) NOT NULL DEFAULT '' COMMENT '封面视频',
  `document_total` int NOT NULL DEFAULT '0' COMMENT '文档总数',
  `read_num` int NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `description` text COMMENT '景区简介',
  `content` text COMMENT '景区内容',
  `score` int NOT NULL DEFAULT '0' COMMENT '评分',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='景区表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_tag`
--

DROP TABLE IF EXISTS `cms_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '标签名称',
  `code` varchar(200) NOT NULL DEFAULT '' COMMENT '标签编码',
  `keyword` varchar(200) NOT NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `read_num` int NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cms_user_like_history`
--

DROP TABLE IF EXISTS `cms_user_like_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cms_user_like_history` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `document_id` bigint NOT NULL DEFAULT '0' COMMENT '文档ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `ip` varchar(200) NOT NULL DEFAULT '' COMMENT 'ip',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=点赞;2=取消点赞;',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户点赞记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_admin`
--

DROP TABLE IF EXISTS `core_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_admin` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nickname` varchar(128) NOT NULL DEFAULT '' COMMENT '昵称',
  `account` varchar(128) NOT NULL DEFAULT '' COMMENT '账户',
  `pwd` varchar(130) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
  `dept_id` bigint NOT NULL DEFAULT '0' COMMENT '所属部门',
  `admin_status` tinyint NOT NULL DEFAULT '0' COMMENT '管理员状态:1=启用;2=停用',
  `permission` json DEFAULT NULL COMMENT '权限',
  `last_pwd` datetime DEFAULT NULL COMMENT '最后次设置密码时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_admin_role_index`
--

DROP TABLE IF EXISTS `core_admin_role_index`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_admin_role_index` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` bigint NOT NULL DEFAULT '0' COMMENT '角色id',
  `admin_id` bigint NOT NULL DEFAULT '0' COMMENT '管理员id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台管理员分配角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_config`
--

DROP TABLE IF EXISTS `core_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_config` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `config_key` varchar(255) NOT NULL DEFAULT '' COMMENT '配置键',
  `config_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '配置值',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='系统配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_dept`
--

DROP TABLE IF EXISTS `core_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `pid` bigint NOT NULL DEFAULT '0' COMMENT '上级部门',
  `dept_name` varchar(128) NOT NULL DEFAULT '' COMMENT '部门名称',
  `connect_name` varchar(128) NOT NULL DEFAULT '' COMMENT '负责人姓名',
  `connect_mobile` varchar(32) NOT NULL DEFAULT '' COMMENT '负责人电话',
  `connect_address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',
  `leader_name` varchar(128) NOT NULL DEFAULT '' COMMENT '领导',
  `thumb` varchar(250) NOT NULL DEFAULT '' COMMENT '缩略图',
  `content` text COMMENT '描述',
  `organize` json DEFAULT NULL COMMENT '组织人员',
  `level` smallint NOT NULL DEFAULT '0' COMMENT '级别',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_permission`
--

DROP TABLE IF EXISTS `core_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_permission` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '权限名称',
  `code` varchar(128) NOT NULL DEFAULT '' COMMENT '权限唯一标识',
  `pid` bigint NOT NULL DEFAULT '0' COMMENT '上级权限单',
  `urls` json DEFAULT NULL COMMENT '关联url',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台权限表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_request_log`
--

DROP TABLE IF EXISTS `core_request_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_request_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `source` tinyint NOT NULL DEFAULT '0' COMMENT '来源（ 1=后台、2=用户端）',
  `tag` varchar(128) NOT NULL DEFAULT '' COMMENT '类别',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '操作说明',
  `request_url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求路径',
  `request_method` varchar(10) NOT NULL DEFAULT '' COMMENT '请求方式',
  `request_params` longtext COMMENT '请求参数（JSON格式）',
  `ip` varchar(200) NOT NULL DEFAULT '' COMMENT 'IP地址（支持IPv6）',
  `use_time` int NOT NULL DEFAULT '0' COMMENT '处理时间（毫秒）',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态：1=成功；2=失败',
  `response_content` longtext COMMENT '响应数据',
  `token` varchar(200) DEFAULT '' COMMENT '请求凭证',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_source` (`source`),
  KEY `idx_tag` (`tag`)
) ENGINE=InnoDB AUTO_INCREMENT=167 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='API请求日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `core_role`
--

DROP TABLE IF EXISTS `core_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `core_role` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) NOT NULL DEFAULT '' COMMENT '角色名称',
  `permission` json DEFAULT NULL COMMENT '权限',
  `role_status` tinyint NOT NULL DEFAULT '0' COMMENT '角色状态:1=启用;2=停用',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_order`
--

DROP TABLE IF EXISTS `mp_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_order` (
  `id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '支付订单id',
  `product_id` bigint NOT NULL DEFAULT '0' COMMENT '支付的产品id',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '支付的用户id',
  `pay_price` int NOT NULL DEFAULT '0' COMMENT '支付金额',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态1=待支付,2=支付中,3=支付成功4=支付失败,5=支付关闭',
  `pay_config_id` bigint NOT NULL DEFAULT '0' COMMENT '支付类型id',
  `third_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '第三方平台的订单号',
  `err_msg` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '失败原因',
  `remark` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
  `fail_time` datetime DEFAULT NULL COMMENT '订单失败时间',
  `close_time` datetime DEFAULT NULL COMMENT '订单关闭时间',
  `complete_time` datetime DEFAULT NULL COMMENT '订单完成时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单更新时间',
  PRIMARY KEY (`id`),
  KEY `aiv_order_state_IDX` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='产品订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_pay_config`
--

DROP TABLE IF EXISTS `mp_pay_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_pay_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '支付名称',
  `photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '图标',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '执行编码',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '开放状态:1=开放;2=不开放',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序字段',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='支付配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_product`
--

DROP TABLE IF EXISTS `mp_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_product` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '产品名称',
  `price` int NOT NULL DEFAULT '0' COMMENT '产品价格',
  `product_type` tinyint NOT NULL DEFAULT '0' COMMENT '产品类型：1=会员;2=积分',
  `terminal_type` tinyint NOT NULL DEFAULT '1' COMMENT '终端类型：0=未知;1=android;2=ios',
  `product_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '产品代码（主要用于ios支付）',
  `product_config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '产品配置',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '执行编码',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `state` int NOT NULL DEFAULT '0' COMMENT '是否开放:1=开放,2=不开放',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `show_type` tinyint NOT NULL DEFAULT '0' COMMENT '是否显示: 1=显示 ,2=不显示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='产品配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_reset_pwd_tokens`
--

DROP TABLE IF EXISTS `mp_reset_pwd_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_reset_pwd_tokens` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `expiration_time` datetime NOT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `count` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_user`
--

DROP TABLE IF EXISTS `mp_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'email',
  `email_verify` int NOT NULL DEFAULT '0' COMMENT '邮箱是否验证',
  `password` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码',
  `user_type` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型:1=普通用户;2=会员',
  `user_status` tinyint NOT NULL DEFAULT '0' COMMENT '用户状态:1=正常;2=停用',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录凭证',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mp_user_open_id_IDX` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='网站用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mp_user_token`
--

DROP TABLE IF EXISTS `mp_user_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mp_user_token` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `ip` varchar(200) NOT NULL DEFAULT '' COMMENT 'ip',
  `token` varchar(200) NOT NULL DEFAULT '' COMMENT 'token',
  `user_agent` varchar(250) NOT NULL DEFAULT '' COMMENT 'user agent',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`user_id`),
  KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户登录凭证表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `paypal_order_logs`
--

DROP TABLE IF EXISTS `paypal_order_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paypal_order_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `local_order_id` varchar(64) NOT NULL COMMENT '本地系统订单ID',
  `paypal_order_id` varchar(64) NOT NULL COMMENT 'PayPal订单ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_paypal_order` (`paypal_order_id`),
  KEY `idx_local_order` (`local_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='PayPal订单日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `paypal_webhook_logs`
--

DROP TABLE IF EXISTS `paypal_webhook_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paypal_webhook_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `local_order_id` varchar(64) DEFAULT NULL COMMENT '本地系统订单ID',
  `paypal_order_id` varchar(64) DEFAULT NULL COMMENT 'PayPal订单ID',
  `event_id` varchar(64) NOT NULL COMMENT 'PayPal事件ID',
  `event_type` varchar(128) NOT NULL COMMENT '事件类型',
  `event_body` json NOT NULL COMMENT '完整事件体(JSON格式)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `process_result` varchar(255) DEFAULT NULL COMMENT '处理结果或错误信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_event_id` (`event_id`),
  KEY `idx_paypal_order` (`paypal_order_id`),
  KEY `idx_local_order` (`local_order_id`),
  KEY `idx_event_type` (`event_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='PayPal Webhook日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `shop_tag`
--

DROP TABLE IF EXISTS `shop_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shop_tag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '标签名称',
  `code` varchar(200) NOT NULL DEFAULT '' COMMENT '标签编码',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态:1=已发布;2=未发布',
  `read_num` int NOT NULL DEFAULT '0' COMMENT '阅读数量',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `match_word` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品标签表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `shop_tag_index`
--

DROP TABLE IF EXISTS `shop_tag_index`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shop_tag_index` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
  `tag_id` int NOT NULL DEFAULT '0' COMMENT '标签id',
  `sort_num` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tid` (`tag_id`),
  KEY `idx_pid` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品标签索引表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `shop_tag_mate`
--

DROP TABLE IF EXISTS `shop_tag_mate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shop_tag_mate` (
  `id` int NOT NULL,
  `content` text NOT NULL COMMENT '内容',
  `seo_title` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo标题',
  `seo_keyword` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo关键词',
  `seo_description` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo描述',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品标签MATE表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_category`
--

DROP TABLE IF EXISTS `sp_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '父类目id',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '名称',
  `code` varchar(200) NOT NULL DEFAULT '' COMMENT '标识',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态:1=正常;2=停用',
  `icon` varchar(200) NOT NULL DEFAULT '' COMMENT '类目图标',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '类目图片',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `sort_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `seo_title` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo标题',
  `seo_keyword` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo关键词',
  `seo_description` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo描述',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='类目表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_order`
--

DROP TABLE IF EXISTS `sp_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(128) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `nickname` varchar(100) NOT NULL COMMENT '昵称',
  `email` varchar(100) NOT NULL COMMENT '邮箱',
  `total_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单总金额',
  `pay_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '实际支付总金额',
  `freight` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '运费',
  `pay_type` smallint unsigned NOT NULL DEFAULT '0' COMMENT '支付方式:1=货到付款',
  `source_type` smallint NOT NULL DEFAULT '0' COMMENT '订单来源：1=PC订单;2=移动端订单',
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '订单状态:1=待付款;2=待发货;3=已发货;4=已完成;5=已关闭;6=无效订单',
  `dispute_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '争议状态:0=无争议;1=买家发起争议;2=卖家处理中;3=平台仲裁中;4=已解决',
  `payment_time` datetime DEFAULT NULL COMMENT '支付时间',
  `delivery_time` datetime DEFAULT NULL COMMENT '发货时间',
  `receive_time` datetime DEFAULT NULL COMMENT '确认收货时间',
  `delivery_company` varchar(128) NOT NULL DEFAULT '' COMMENT '物流公司(配送方式)',
  `delivery_sn` varchar(200) NOT NULL DEFAULT '' COMMENT '物流单号',
  `remark` varchar(230) NOT NULL DEFAULT '' COMMENT '订单备注',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_order_item`
--

DROP TABLE IF EXISTS `sp_order_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_order_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '商品标题',
  `sku_title` varchar(200) NOT NULL DEFAULT '' COMMENT '商品SKU内容',
  `sku_code` varchar(200) NOT NULL DEFAULT '' COMMENT '商品SKU编码',
  `thumb` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片',
  `order_id` int unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
  `product_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `sku_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品SKUid',
  `total_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '总金额',
  `pay_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '实际支付金额',
  `quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '购买数量',
  `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单价',
  `cost_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '成本单价',
  `original_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '原价单价',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oid` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单节点表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_order_operate_history`
--

DROP TABLE IF EXISTS `sp_order_operate_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_order_operate_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` int unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
  `operate_user` varchar(200) NOT NULL DEFAULT '' COMMENT '操作人(用户;系统;管理员)',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '操作备注',
  `ip` varchar(200) NOT NULL DEFAULT '' COMMENT '操作IP',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oid` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单操作历史记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_order_receive_address`
--

DROP TABLE IF EXISTS `sp_order_receive_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_order_receive_address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` int unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
  `first_name` varchar(64) NOT NULL DEFAULT '' COMMENT '收发货人姓',
  `last_name` varchar(64) NOT NULL DEFAULT '' COMMENT '收发货人名',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '收货人邮箱',
  `phone` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人电话',
  `province` varchar(64) NOT NULL DEFAULT '' COMMENT '省/直辖市',
  `city` varchar(64) NOT NULL DEFAULT '' COMMENT '市',
  `region` varchar(64) NOT NULL DEFAULT '' COMMENT '区',
  `detail_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `postal_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_oid` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单收货地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_order_refund`
--

DROP TABLE IF EXISTS `sp_order_refund`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_order_refund` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` int unsigned NOT NULL DEFAULT '0' COMMENT '关联订单id',
  `refund_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '退款金额',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '退款状态:1=申请中;2=处理中;3=已退款;4=已拒绝;5=已取消',
  `reason` varchar(500) DEFAULT NULL COMMENT '退款原因',
  `refund_time` datetime DEFAULT NULL COMMENT '退款完成时间',
  `images` json DEFAULT NULL COMMENT '退款凭证图片集',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `refund_no` varchar(64) DEFAULT NULL COMMENT '退款单号',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='退款记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_prod_attributes`
--

DROP TABLE IF EXISTS `sp_prod_attributes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_prod_attributes` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(200) NOT NULL COMMENT '属性名称',
  `sort_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_prod_attributes_value`
--

DROP TABLE IF EXISTS `sp_prod_attributes_value`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_prod_attributes_value` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `prod_attributes_id` int unsigned NOT NULL COMMENT '商品属性id',
  `title` varchar(200) NOT NULL COMMENT '值名称',
  `sort_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_pai` (`prod_attributes_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性值表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_product`
--

DROP TABLE IF EXISTS `sp_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_product` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `category_id` int unsigned NOT NULL COMMENT '类目id',
  `title` varchar(200) NOT NULL COMMENT '商品名称',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态:1=上线;2=下线;3=待审核;4=审核不通过',
  `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '当前价格',
  `original_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '原价',
  `cost_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '成本价',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '库存',
  `open_sku` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否启用sku:1=是;2=否',
  `picture` varchar(200) NOT NULL COMMENT '商品主图',
  `picture_gallery` json DEFAULT NULL COMMENT '商品图片集',
  `description` varchar(200) NOT NULL COMMENT '描述',
  `sold_num` smallint unsigned DEFAULT NULL COMMENT '销量',
  `version` int unsigned NOT NULL DEFAULT '0' COMMENT '版本号',
  `sort_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `hot` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否热门:1=是;2=否',
  `putaway_time` datetime DEFAULT NULL COMMENT '上架时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `detail_url` varchar(512) DEFAULT NULL COMMENT '商品详情页URL',
  `price_locked` tinyint(1) DEFAULT '0' COMMENT '价格是否锁定',
  PRIMARY KEY (`id`),
  KEY `idx_cid` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_product_content`
--

DROP TABLE IF EXISTS `sp_product_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_product_content` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int unsigned NOT NULL COMMENT '商品id',
  `content` text NOT NULL COMMENT '内容',
  `seo_title` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo标题',
  `seo_keyword` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo关键词',
  `seo_description` varchar(200) NOT NULL DEFAULT '' COMMENT 'seo描述',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_proid` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品内容表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_product_property`
--

DROP TABLE IF EXISTS `sp_product_property`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_product_property` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int unsigned NOT NULL COMMENT '商品id',
  `title` varchar(200) NOT NULL COMMENT '属性名称',
  `value` varchar(200) NOT NULL COMMENT '属性值',
  `sort_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_proid` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品属性表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_sku`
--

DROP TABLE IF EXISTS `sp_sku`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_sku` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int unsigned NOT NULL COMMENT '商品id',
  `sku_code` varchar(200) NOT NULL COMMENT '商品编码(商品属性ID;商品属性ID)',
  `title` varchar(200) NOT NULL COMMENT '商品SKU名称',
  `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '当前价格',
  `original_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '原价',
  `cost_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '成本价',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '库存',
  `default_show` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否默认显示',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态:1=启用;2=停用',
  `version` int unsigned NOT NULL DEFAULT '0' COMMENT '版本号',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_proid` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品SKU表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_sku_index`
--

DROP TABLE IF EXISTS `sp_sku_index`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_sku_index` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` int unsigned NOT NULL COMMENT '商品id',
  `sku_id` int unsigned NOT NULL COMMENT '商品SKU id',
  `prod_attributes_id` int unsigned NOT NULL COMMENT '商品属性id',
  `prod_attributes_value_id` int unsigned NOT NULL COMMENT '商品属性值id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ski` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品SKU索引表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_user_address`
--

DROP TABLE IF EXISTS `sp_user_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_user_address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `title` varchar(200) NOT NULL COMMENT '地址名称',
  `default_status` tinyint(1) DEFAULT NULL COMMENT '是否默认收货地址：1=是;2=否',
  `first_name` varchar(64) NOT NULL DEFAULT '' COMMENT '收发货人姓',
  `last_name` varchar(64) NOT NULL DEFAULT '' COMMENT '收发货人名',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '收货人邮箱',
  `phone` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人电话',
  `province` varchar(64) NOT NULL DEFAULT '' COMMENT '省/直辖市',
  `city` varchar(64) NOT NULL DEFAULT '' COMMENT '市',
  `region` varchar(64) NOT NULL DEFAULT '' COMMENT '区',
  `detail_address` varchar(200) NOT NULL DEFAULT '' COMMENT '详细地址',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '国家',
  `postal_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_uid` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户收货地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sp_user_cart`
--

DROP TABLE IF EXISTS `sp_user_cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sp_user_cart` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `fingerprint` varchar(255) NOT NULL DEFAULT '' COMMENT '浏览器指纹数据',
  `title` varchar(200) NOT NULL COMMENT '商品名称',
  `sku_title` varchar(200) NOT NULL DEFAULT '' COMMENT '商品SKU内容',
  `sku_code` varchar(200) NOT NULL DEFAULT '' COMMENT '商品SKU编码',
  `thumb` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片',
  `product_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `sku_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商品SKUid',
  `total_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '总金额',
  `pay_amount` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '实际支付金额',
  `quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '购买数量',
  `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '单价',
  `original_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '原价单价',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户购物车';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'zg_shop'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-12 17:47:46
