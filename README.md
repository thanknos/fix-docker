# fix-docker

#### 介绍
golang命令行工具cobra实现docker与docker-compose的一键部署安装，可在主流Linux操作系统上实现docker与docker-compose的命令行部署安装

#### 软件架构

golang命令行工具cobra实现

Debian及其衍生Linux发行版使用apt实现init初始化

基于Red Hat的Linux发行版使用rpm实现init初始化


#### 安装教程

1. 下载tar.gz安装包

   ```shell
   wget -c --referer=https://pan.baidu.com -O fix-docker.tar.gz "https://d.pcs.baidu.com/file/185e114efi4d99818399ccdad1dddb30?fid=1099665794657-250528-321969089793068&rt=pr&sign=FDtAERK-DCb740ccc5511e5e8fedcff06b081203-l%2FTN8uU1w5GoeSKJS10gPgkJ8OQ%3D&expires=8h&chkbd=0&chkv=0&dp-logid=2543806217839369865&dp-callid=0&dstime=1735278118&r=227283135&vuk=1099665794657&origin_appid=15195230&file_type=0&access_token=123.0efe10dc1300bf7a3cc51c8187f0fbe9.YlrgGltl7U03cf0FmUwOYVy8XQd_tVIxOCYQ8WL.CxYDtQ"
   ```

2. tar -zxvf解压

   ```shell
   tar -zxvf fix-docker.tar.gz
   ```

#### 使用说明

1.  **./fix-docker查看命令行**
2.  **./fix-docker init rpm/apt**
3.  **./fix-docker install docker/docker-compose**
4.  **./fix-docker uninstall docker/docker-compose**

#### 参与贡献

**ThankNos**

#### 联系方式

```
微信: 18779255721
QQ: 2164591400
```

