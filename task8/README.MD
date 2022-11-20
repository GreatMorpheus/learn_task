#1 加载configmap
 kubectl apply -f httpserver_configmap.yaml

#2 拉取镜像,创建Pod
 kubectl apply -f httpserver.yaml

#3 创建service
 kubectl apply -f httpserver_service.yaml
