#1.执行命令拉取并加载httpserver镜像
docker run -it  -p 8080:80 569646632/task_3

#2.获取httpserver的PID
  #获取container id
  docker ps -f label=multi.label1="httpserver-task3"
  #根据container id 获取pid
  docker inspect container_id | grep -i pid

#3.根据pid进入容器执行命令
   nsenter -t pid -n ip a
