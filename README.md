# 导出csdn博客markdown源码
之前在csdn写了一些博客，想迁移到github pages上，因此写一个导出博客源码的工具

博客的元信息基本都有导出，可以很方便地通过二次开发直接生成hexo、vuepress静态页面
## 功能
- 导出个人csdn账号下，每篇博客的标题、内容、创建时间、分类、标签等信息
- 所有博客都输出到一个json文件中
- 无markdown源码的博客，content字段为html代码

## 编译
见build.sh

## 使用

- 编译
  - 直接下载对应的release也行
- 填写conf.yml配置文件
- 运行

*conf.yml必须与可执行文件在同一目录下*