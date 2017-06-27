-- phpMyAdmin SQL Dump
-- version phpStudy 2014
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2017 �?06 �?27 �?16:44
-- 服务器版本: 5.5.40
-- PHP 版本: 5.6.1

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `gotest`
--

-- --------------------------------------------------------

--
-- 表的结构 `articles`
--

CREATE TABLE IF NOT EXISTS `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cateid` int(11) DEFAULT '0' COMMENT '分类ID',
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `pic` varchar(255) DEFAULT NULL,
  `content` text,
  `addtime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=18 ;

--
-- 转存表中的数据 `articles`
--

INSERT INTO `articles` (`id`, `cateid`, `title`, `description`, `pic`, `content`, `addtime`) VALUES
(1, 4, '测试一下标题', '测试一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(2, 4, '测试一下标题', '测试一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(3, 5, '测试xxx一下标题', '测试xxx一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(4, 5, '测试xxx一下标题', '测试xxx一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(5, 5, '测试xxx一下标题', '测试xxx一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(6, 5, '测试xxx一下标题', '测试xxx一下描述', 'http://s2.juancdn.com/bao/160921/d/9/57e26ab2151ad118338b45a9_800x800.jpg', '测试一下内容', 1497704998),
(7, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(8, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(9, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(10, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(11, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(12, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(13, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(14, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(15, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(16, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998),
(17, 6, '测试yyy一下标题', '测试yyyy一下描述', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498366274764&di=ec6b5df0275bbd4364da7c50887657ad&imgtype=0&src=http%3A%2F%2Fpic.baike.soso.com%2Fp%2F20140110%2F20140110164023-78701636.jpg', '测试一下内容', 1497704998);

-- --------------------------------------------------------

--
-- 表的结构 `cate`
--

CREATE TABLE IF NOT EXISTS `cate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT '0',
  `name` varchar(255) DEFAULT NULL,
  `addtime` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='分类表' AUTO_INCREMENT=7 ;

--
-- 转存表中的数据 `cate`
--

INSERT INTO `cate` (`id`, `parent_id`, `name`, `addtime`) VALUES
(4, 0, '家居', 1497704998),
(5, 0, '电器', 1497705069),
(6, 0, '数码', 1497705107);

-- --------------------------------------------------------

--
-- 表的结构 `photos`
--

CREATE TABLE IF NOT EXISTS `photos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cateid` int(11) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `addtime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=23 ;

--
-- 转存表中的数据 `photos`
--

INSERT INTO `photos` (`id`, `cateid`, `image`, `title`, `addtime`) VALUES
(1, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607813&di=5f71d008119bc094817e116eb5a1c31b&imgtype=0&src=http%3A%2F%2Fwww.taopic.com%2Fuploads%2Fallimg%2F121219%2F267857-12121921013456.jpg', '好车图', 1497704998),
(2, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607813&di=29be14ecec834cdf0de64fa93976c7b1&imgtype=0&src=http%3A%2F%2Fxcmn.jsyks.com%2Fphoto%2Fimg.xgo-img.com.cn%2Fpics%2F1505%2F1504435.jpg', '好车图', 1497704998),
(3, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607813&di=b9f8888c9492ffe1c38be299b9d99945&imgtype=0&src=http%3A%2F%2Fpic33.nipic.com%2F20131011%2F11350592_170605502000_2.jpg', '好车图', 1497704998),
(4, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607813&di=b9f8888c9492ffe1c38be299b9d99945&imgtype=0&src=http%3A%2F%2Fpic33.nipic.com%2F20131011%2F11350592_170605502000_2.jpg', '好车图', 1497704998),
(5, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607813&di=b9f8888c9492ffe1c38be299b9d99945&imgtype=0&src=http%3A%2F%2Fpic33.nipic.com%2F20131011%2F11350592_170605502000_2.jpg', '好车图', 1497704998),
(6, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(7, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(8, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(9, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(10, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(11, 4, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607812&di=22ebc2bd8964ec4a51a0ff90ee02ed26&imgtype=0&src=http%3A%2F%2Fdl.bizhi.sogou.com%2Fimages%2F2012%2F02%2F06%2F74408.jpg', '好车图', 1497704998),
(12, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=fdd06037210a473ddc4e76d2ce1c7bd4&imgtype=0&src=http%3A%2F%2Ft1.niutuku.com%2F960%2F01%2F01-171601.jpg', '好车图', 1497704998),
(13, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=fdd06037210a473ddc4e76d2ce1c7bd4&imgtype=0&src=http%3A%2F%2Ft1.niutuku.com%2F960%2F01%2F01-171601.jpg', '好车图', 1497704998),
(14, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=fdd06037210a473ddc4e76d2ce1c7bd4&imgtype=0&src=http%3A%2F%2Ft1.niutuku.com%2F960%2F01%2F01-171601.jpg', '好车图', 1497704998),
(15, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=fdd06037210a473ddc4e76d2ce1c7bd4&imgtype=0&src=http%3A%2F%2Ft1.niutuku.com%2F960%2F01%2F01-171601.jpg', '好车图', 1497704998),
(16, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(17, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(18, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(19, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(20, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(21, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498377607811&di=d0ba7780b48369418bebf96141eea834&imgtype=0&src=http%3A%2F%2Fimg0.pcauto.com.cn%2Fpcauto%2F1405%2F30%2F4493541_100003604293146_thumb.gif', '好车图', 1497704998),
(22, 5, 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498396650834&di=406026c7c3b889b4ecf8db819cd90c88&imgtype=0&src=http%3A%2F%2Fds.com.cn%2Ftemplates%2Fprotostar%2Fimg%2Fcar%2Fds5ls%2Fblock2.jpg', '好车图', 1497704998);

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(200) DEFAULT NULL COMMENT '用户名',
  `password` varchar(200) DEFAULT NULL COMMENT '密码',
  `email` varchar(200) DEFAULT NULL COMMENT '邮件',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `addtime` int(11) DEFAULT '0' COMMENT '添加时间',
  `logintime` int(11) DEFAULT '0' COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表' AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `username`, `password`, `email`, `avatar`, `addtime`, `logintime`) VALUES
(1, 'test', '81dc9bdb52d04dc20036dbd8313ed055', 'test@admin.com', 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1498469950566&di=0e2bfc3c0ce2cbbae06c188da9606709&imgtype=jpg&src=http%3A%2F%2Fimg3.imgtn.bdimg.com%2Fit%2Fu%3D3797592229%2C3840448992%26fm%3D214%26gp%3D0.jpg', 1498472522, 1498476142);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
