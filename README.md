# Annotations_from_wxread

利用Go语言实现的微信读书笔记导出小工具，支持自定义导出模板

写这个小工具一是为了锻炼自己还在学习的Go语言能力，二才是为了把微信读书的一些笔记迁移到我的Obsidian中

支持自定义模板，暂不支持自定义路径

在 `config/config.json`
中可以配置源文件和导出目录，可以修改为绝对路径。
同时可以通过设置 `annotation_title_configurable`
，自定义是否需要设置笔记标题。
