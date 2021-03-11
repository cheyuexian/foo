* go get是不会从本地仓库获取版本信息的，查看go get在module模式下工具链实现代码也可得出这个结论,只会拉取最新的代码
* go list -u -m all 更新查看
* go get -u 会更新mod文件版本