# Annotations_from_wxread

利用Go语言实现的微信读书笔记导出小工具，支持自定义导出模板

写这个小工具一是为了锻炼自己还在学习的Go语言能力，二才是为了把微信读书的一些笔记迁移到我的Obsidian中

支持自定义模板，暂不支持自定义路径

在 `config/config.json`
中可以配置源文件和导出目录，可以修改为绝对路径。
同时可以通过设置 `annotation_title_configurable`
，自定义是否需要设置笔记标题。

## Example

### 输入文件

```text

《海子的诗》
海子
2个笔记


◆  我，以及其他的证人

2023/2/26 发表想法
为自己的日子，在自己的脸上留下伤口
>> 为自己的日子
在自己的脸上留下伤口
因为没有别的一切为我们作证

>> 我和过去
隔着黑色的土地
我和未来
隔着无声的空气


```

### 模板文件 ([template/[annotation.source]/[annotation.name].md](https://github.com/Qnurye/Annotations_from_wxread/blob/main/template/%5Bannotation.source%5D/%5Bannotation.name%5D.md))

```markdown
---
type: annotation
tags: [annotations]
source: <% annotation.source %>
created_at: <% annotation.time %>
chapter: <% annotation.chapter %>
author: <% annotation.author %>
---

<% annotation.comment %>

> [!Quote]
> 
> <% annotation.quotation %>

```

### 执行程序

```shell
go run main.go
```

### 输出文件举例 (output/海子的诗/为自己的日子，在自己的脸上留下伤口.md)

```text
├── output
│ ├── 海子的诗
│ | ├── 为自己的日子，在自己的脸上留下伤口.md
│ │ └── 我，以及其他的证人-54.md
│ ├── 黑客与画家（10万册纪念版）
│ │ ├── 第二部分 黑客如何工作及影响世界-4.md
│ │ └── 第二部分 黑客如何工作及影响世界-5.md
│ └── 霍乱时期的爱情
│     ├── 1-15.md
│     ├── 1-16.md
│     ├── 1-17.md
│     ├── 1-18.md
│     ├── 1-19.md
│     ├── 1-20.md
│     ├── 1-21.md
│     ├── 1-22.md
│     ├── 1-23.md
│     ├── 1-24.md
```

```markdown
---
type: annotation
tags: [annotations]
source: 海子的诗
created_at: 2023-02-26T16:11
chapter: 我，以及其他的证人
author: 海子
---

为自己的日子，在自己的脸上留下伤口

> [!Quote]
> 
> 为自己的日子
在自己的脸上留下伤口
因为没有别的一切为我们作证

```

# TODO

本程序仅为本人go语言练手小项目，有很多地方尚未完善

- [ ] 根据模板文件位置，自定义导出笔记结构
- [ ] 构建跨平台程序
