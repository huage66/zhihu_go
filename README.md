ql中

技术难点,当大批量数据进来时, 如何保证数据写入持久化磁盘中.

发送的微博是否有丢失的风险,如何把控

支持删除功能,那么删除是立即删除,还是延迟删除.

## 表设计

文章表

```sql
CREATE TABLE IF NOT EXISTS `zhihu_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `content` text NOT NULL COMMENT '文章内容',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zhihu_artical_user_id` (`user_id`,`id`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='文章表';
```

话题表

```sql
CREATE TABLE IF NOT EXISTS `zhihu_topic` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '话题标题',
  `description` text DEFAULT '' COMMENT '话题的描述信息',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zhihu_artical_user_id` (`user_id`,`id`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题表';
```

话题回答表

```sql
CREATE TABLE IF NOT EXISTS `zhihu_topic_reply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '话题标题',
  `content` text DEFAULT '' COMMENT '回答的内容',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `topic_id` bigint(20) NOT NULL COMMENT '话题id',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zhihu_artical_user_id` (`user_id`,`id`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题回答表';
```

文章评论表

```sql
CREATE TABLE IF NOT EXISTS `zhihu_article_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '话题标题',
  `content` text DEFAULT '' COMMENT '回答的内容',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `artical_id` bigint(20) NOT NULL COMMENT '话题id',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zhihu_artical_user_id` (`user_id`,`id`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='文章评论表';
```

话题评论表

```sql
CREATE TABLE IF NOT EXISTS `zhihu_topic_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL COMMENT '话题标题',
  `content` text DEFAULT '' COMMENT '回答的内容',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `topic_id` bigint(20) NOT NULL COMMENT '话题id',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `zhihu_artical_user_id` (`user_id`,`id`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题评论表';
```

用户表, 这种主表应该只需要用户的基本信息, 不经常改动的信息

```sql
CREATE TABLE IF NOT EXISTS `zhihu_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` varchar(64) NOT NULL COMMENT '用户名称',
  `avatar` varchar(256) NOT NULL COMMENT '头像',
  `phone` varchar(64) NOT NULL COMMENT '电话号码',
  `password` varchar(255) NOT NULL COMMENT '密码,加密形式',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题评论表';
```

用户详细信息表, 这种表应该是用户所关联或者活动的数据表, 该表的数据经常变动

```sql
CREATE TABLE IF NOT EXISTS `zhihu_user_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `scores` int(11) NOT NULL '积分, 积分代表用户等级',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题评论表';
```

头衔表, 头衔应当有多种类型, 每种类型它的达成方式不太一样

```sql
CREATE TABLE IF NOT EXISTS `zhihu_rank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `rank_name` int(11) NOT NULL '头衔名称',
  `rank_avatar` varchar(255) NOT NULL `头衔图片, 比较炫酷`,
  `rank_type` int(8) NOT NULL DEFAULT '0' COMMENT '0-默认升级头衔, 其他头衔根据业务需求来,待定',
  `delete` tinyint DEFAULT 0 COMMENT '是否删除, 默认0-未删除 1-删除',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '记录创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=1 COMMENT='话题评论表';
```

头衔达成规则表, 可以设置规则来让用户获取头衔, 而且有些头衔可能还有限时, 这种应当是业务表,设置不同规则门槛来达成不同的头衔.

头衔关联表, 一个用户可以拥有多个头衔, 但是注意, 有些头衔是有等级的, 那么到达下一等级后, 上一等级的头衔应当是被覆盖

管理员表, 一般用于运营人员审核内容或者内部员工查看数据

角色表

角色权限表

临时审核数据表, 审核s数据是否写表, 用户修改信息是否需要流程记录表, 记录用户的流转记录已经内部审核记录, 这样才能方便查看

## 项目梳理

基础表设计

根据这些数据可能衍生出话题圈(聚合话题的超集)

排行榜单(实时热点排行,根据热门话题的排行)

邀请回答(根据算法去邀请别人作答)

热门回答

热门收藏(根据收藏来变更)

热门专栏(专门的号来写特定的文章, 比如汽车栏目等等)

(话题,回答,文章等)有评论 赞同 分享 收藏 喜欢等功能, 还可以对话题进行举报等, 话题和文章有申请授权转载 不感兴趣 (设置屏蔽关键字等)

评论有点赞 回复 踩 举报等功能 (排序问题)

所有的踩, 点赞, 收藏等功能是否需要立即落盘, 还是先在缓存中存储, 然后定时同步数据, 缓存会导致数据丢失, 是否影响.

评论数据如何存储, 评论完成是否立马落盘, 如何回显评论, 评论是否应当先存入缓存,然后批量落盘,可能存在评论丢失问题, 评论还有删除功能,是否需要一直删除,应当做成微信那种俩分钟内可以删除评论.因为俩分钟后该评论会被写入磁盘.
