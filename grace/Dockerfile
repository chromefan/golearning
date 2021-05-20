#源镜像
FROM centos:7
#设置工作目录
RUN mkdir /app
#将服务器的go工程代码加入到docker容器中
ADD . /app
#go构建可执行文件
WORKDIR /app

RUN sh control.sh start
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./grace"]