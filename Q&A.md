# Q&A
## yarn : 无法加载文件 C:\Users\10993\AppData\Roaming\npm\yarn.ps1，因为在此系统上禁止运行脚本。
- 1、搜索powershell，右键以管理员身份运行
- 2、若要在本地计算机上运行您编写的未签名脚本和来自其他用户的签名脚本，请使用以下命令将计算机上的 执行策略更改为 RemoteSigned
执行：set-ExecutionPolicy RemoteSigned
- 3、查看执行策略：get-ExecutionPolicy
- 4、 关闭命令窗口 即可