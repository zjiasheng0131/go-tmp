## Device Provision

**备注**：创建流程仅配置IP/MAC,用户名、密码，机器名（如果需要且可行）；其它配置、安装步骤在配置流程完成

| 设备类型                            | Cloud-init/WinRM | 创建流程                                                                                                                                                                                                                                                           | 配置流程                                               | 模版（仅调试用）          | 模版维护方     |维护模式| 进展   |备注
  |:--------------------------------|:-----------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------------------------------------|:------------------|:----------|:-----|:-----|:-----|
| FortiGate                       | Yes(类似)/No       | **按需创建模版**<br/>**创建设备**<br/> 1. Clone<br/>2. BasicConfig<br/>3. IP&Mac<br/>4. Boot<br/>5. Hostname config<br/>6. Verify<br/>7. Deliver<br/>8. Import licence                                                                                                   | Adapter通过管理网，SSH配置（优先）<br/>~~或者调用FortiGate API配置~~ | 9201->9205(136)   | Kimi组     |手动| 已验证  |
| FortiCSP-Site                   | No/No            | **按需创建模版**<br/>1. 下载镜像<br/> 2. 创建VM<br/> 3. 启动VM<br/> 4. 代码创建~~临时管理IP~~串口，连接VM<br/>5. 交互配置用户名/密码（串口）；配置VM为DHCP方式；<br/> 6. 关机<br/> 7. 设置为模版<br/> 8. 注册模版<br/> **创建设备**<br/>  1. Clone<br/>2. BasicConfig<br/>3. IP&Mac<br/>4. Boot<br/>5. Verify<br/>6. Deliver | Adapter通过管理网，SSH配置                                 | 9204              | Kimi组（昆仑） |自动| 待验证  |
| FortiCSP-Tester&Agent(XP)       | No/No            | **创建设备**<br/> 1. Clone<br/>2. BasicConfig<br/>3. IP&Mac（DHCP）<br/>4. Boot<br/>5. 通过CSP tester配置基础信息<br/>6. Verify<br/>7. Deliver                                                                                                                               | Adapter通过管理网，调用CSP tester配置(当前只有更改机器名API)          | 9008,9009(32-bit) | 测试人员维护    |手动| 待验证  |内置 Tester&Agent，通过其实现软件安装和配置
| FortiCSP-Tester&Agent(XP above) | No/Yes           | **创建设备**<br/> 1. Clone<br/>2. BasicConfig<br/>3. IP&Mac（DHCP）<br/>4. Boot<br/>5. 通过WinRM config<br/>6. Verify<br/>7. Deliver                                                                                                                                   | Adapter通过管理网，WinRM配置（优先）<br/>或者调用CSP tester配置      | 9101-9109         | 测试人员维护    |手动| 待验证  |
| Linux                           | Yes/No           | **创建VM**<br/> 1. Clone<br/>2. BasicConfig<br/>3. IP&Mac<br/>4. Boot<br/>5. Verify<br/>6. Deliver                                                                                                                                                               | Adapter通过管理网，SSH配置                                 | 9002              | Kimi组     |手动| 已验证  
| Windows                         | No/Yes           | **创建VM**<br/> 1. Clone<br/>2. BasicConfig<br/>3. IP&Mac（DHCP）<br/>4. Boot<br/>5. 通过WinRM config<br/>6. Verify<br/>7. Deliver                                                                                                                                   | Adapter通过管理网，WinRM配置                               | 9101-9109         | 测试人员维护    |手动| 已验证  
| ~~FortiCSP-Manager(Server)~~    | N/A              | N/A                                                                                                                                                                                                                                                            | N/A                                                | N/A               | N/A       |N/A| 暂不考虑 


### Fortigate
```
# 下载镜像
https://info.fortinet.com/builds/?project_id=410&show_interim=false
例如：
https://info.fortinet.com/files/FortiOS/v7.00/images/build2573/FGT_VM64_KVM-v7.4.3.F-build2573-FORTINET.out.kvm.zip

# 制作镜像
qm disk import 9205 /mnt/pve/cephfs-img/template/iso/FGT-v7.4.3.F-build2573.img VM-Storage-Pool
  
# show/config system global
      set hostname ‘gkl-fortigate-1’
      set timezone 'US/Pacific'
      set alias "lab180-fortigate-1"
      set gui-auto-upgrade-setup-warning disable
  end
  
# config system admin
    edit "admin"
        set password ENC SH2RQwpwjZFFiUwShFlBEmWp/Jqg1/lfzPv1azBVbwoeDbNmaFIlVZpBGuaTHI=
    next
end  

# show system dhcp server
# config system dhcp server
    edit 12
        set dns-service default
        set default-gateway 10.188.0.254
        set netmask 255.255.255.0
        set interface "PV_VM_VLAN_188"
        config ip-range
            edit 1
                set start-ip 10.188.0.2
                set end-ip 10.188.0.200
            next
        end
        config reserved-address
            edit 1
                set ip 10.188.0.158
                set mac bc:24:11:ad:71:b9
            next
            edit 2
                set ip 10.188.0.20
                set mac bc:24:11:e3:23:83
            next
        end
    next
end
            

-- 获得网卡IP
# get system interface physical port1
https://10.65.184.21 admin/admin
模版 9201
```

### FortiCSP-Site
```shell
  # 创建并外部配置VM
  1. wget ftp://test:test@10.65.10.243/images/fortiot/csp_site/v6.1.0-build7005/FortiCSP-KVM-v6.1.0-build7005.kvm.zip 并解压
  2. Create VM并配置特定VM参数
  3. 导入disk 
      qm disk import 9209 FCSP-Manager-v6.0.0-build0075.qcow2  VM-Storage-Pool  或者
      qm import 9208 FortiCSP-Manager-KVM-v6.1.0-build0172.out --storage VM-Storage-Pool
  4. Attach disk to VM
  5. Add serial0
  6. 增加辅助磁盘 32G 存日志（仅对于FortiGate）
  7. ~~ 做光驱，注册licence（仅对于FortiGate，已不需要）~~
  8. Change boot sequence
  9. Start the VM
  10. 创建websocket->串口->terminal通讯连接（需封装为类paramiko）
      - Websocket通讯，流管理，go routine实现异步处理（封装底层 WebSocket 的 send 和 recv 逻辑，提供易于使用的 API 层，类似于 Paramiko 的 SSHClient）
      - 会话管理（封装 WebSocket 连接，管理会话的建立、关闭，以及发送/接收数据）
      - 命令发送和解析（类似 Paramiko，提供接口 ExecuteCommand 来发送命令并获取执行结果。使用同步的方式处理命令的发送和响应获取（在底层通过 WebSocket 实现异步））
      - 自动登录
      - 错误处理与超时机制
            
  # 利用串口在VM内部配置 FortiCSP-Site（自动做template）
  1. 设置模版密码（admin->enter->enter） 
  2. 改DHCP
  3. 设置防火墙（https,ping,ssh）
  （改机器名放在创建VM的流程里）
  
  # 后续
  1. 关机
  2. 配置为模版
  3. 在模版管理模块注册并分发镜像至别的集群
  
  # 模版管理
  - 模版ID范围及生成管理
  - 模版生命周期管理（创建-删除）
  
  # 总流程调整
  - 模版创建作为一个嵌入流程，怎么融合进整个流程
  - 如果模版不存在，返回什么（2009）
  - 如果模版的image在FTP Server不存在，返回什么（2010）
  
  # FortiCSP/FortiCSPManager/FortiGate
  # 联合VM创建流程调试
  
  模版9202？？？
  # 显示网卡 
  show system interface
  # 获取IP
  get system interface
  # 终端密码第一次进去设置；Web默认账号密码是 "admin/password"
  # ip 默认为192.169.1.99/24， 需要更改为dhcp
  # template 9204, console 账号密码是 "admin/Forti1@#", dhcp 获得IP
  config system interface
  edit port1
  set mode dhcp
  end 
  
  
```

### FortiCSP-tester
#### APIs
* CSP Tester API: http://WINDOWS_VM_IP:8086/command
* Installation package: ftp://test:test@10.65.10.243/images/fortiot/csp_tester/windows/x86/
* Example
  - Post Data:
    - RUN CMD: { “command”: “ipconfig” }
    - RUN PowerShell: { “command”: “powershell -command \”get-service\””}
  - Configure examples:
    - Setup Hostname
      - CMD: WMIC ComputerSystem Where Name="%ComputerName%" Call Rename "new-hostname"
      - PowerShell: Rename-Computer -NewName test-123 -Force -Restart
    - Setup existing user’s password
      - CMD: net user ${USERNAME} ${PASSWORD}
      - PowerShell: Set-LocalUser -Name test -Password (ConvertTo-SecureString -AsPlainText 'Forti1@#' -Force)
    - Add new user
      - CMD: net user ${USERNAME} ${PASSWORD} /ADD
      - PowerShell: New-LocalUser -Name test -Password (ConvertTo-SecureString -AsPlainText 'Forti1@#' -Force)
    - Add Route rules
      - CMD:
        - route add 172.16.0.0 mask 255.255.255.0 10.65.184.254 METRIC 100
        - route print
        - route delete 172.16.0.0
      - PowerShell:
        - New-NetRoute -DestinationPrefix “172.16.0.0/24” -NextHop 10.65.184.254 –InterfaceAlias Ethernet -RouteMetric 100
        - Get-NetRoute
        - Remove-NetRoute –DestinationPrefix “172.16.0.0/24” –Confirm:$false
* API endpoint list      
  ```json lines
  {
      "Operation": {
          "POST REQUEST_A_FILE_OPERATION": "/file",
          "POST REQUEST_A_REGISTRY_OPERATION": "/reg",
          "GET  LIST_AVAILABLE_VOLUMES": "/volume",
          "POST REQUEST_A_VOLUME_OPERATION": "/volume",
          "POST REQUEST_A_APP_OPERATION": "/app"
      },
      "Command": {
          "POST EXECUTE_A_COMMAND": "/command"
      },
      "Agent": {
          "GET CHECK_AGENT_CONFIG_UPDATE": "/agent/config",
          "GET CHECK_AGENT_TIME": "/agent/time",
          "GET CSP_AGENT_ID": "/agent/id",
          "GET CSP_AGENT_PROXY": "/agent/proxy",
          "GET CSP_AGENT_VERSION": "/agent/version",
          "GET CSP_AGENT_STATUS": "/agent/status",
          "POST SYNC_AGENT_TIME": "/agent/time",
          "POST UPGRADE_CSP_AGENT": "/agent/upgrade",
          "POST INSTALL_CSP_AGENT": "/agent/install",
          "POST UNINSTALL_CSP_AGENT": "/agent/uninstall"
      },
      "Utility": {
          "GET CSP_TESTER_VERSION": "/version",
          "GET CSP_TESTER_LATEST_VERSION": "/version/latest",
          "GET HEALTH_CHECK": "/system/healthcheck",
          "GET SYSTEM_INFO": "/system/info",
          "GET WIN_DEFENDER_STATUS": "/system/windows/defender",
          "GET WIN_USER_SID": "/system/windows/user/sid",
          "GET WIN_PROXY": "/system/windows/proxy",
          "GET WIN_EVENTVIEWER_LOGS": "/system/windows/eventviewer",
          "POST UPGRADE_CSP_TESTER": "/upgrade",
          "POST REQUEST_A_TIME_SYNC": "/system/time",
          "POST DOWNLOAD_A_FILE": "/system/download",
          "POST REBOOT_SYSTEM": "/system/reboot",
          "POST SET_SYSTEM_NETWORK": "/system/network",
          "POST SET_SYSTEM_DNS": "/system/dns",
          "POST SET_WIN_PROXY": "/system/windows/proxy",
          "POST SET_WIN_UAC": "/system/windows/uac",
          "POST SET_WIN_HOSTNAME": "/system/windows/hostname",
          "POST SET_WIN_DESCRIPTION": "/system/windows/description",
          "POST SET_WIN_NETSHARE": "/system/windows/netshare",
          "POST SET_WIN_AUTOLOGON": "/system/windows/autologon",
          "POST SET_WIN_DEFENDER": "/system/windows/defender",
          "POST SET_WIN_UPDATE": "/system/windows/update",
          "POST SET_WIN_STARTUP": "/system/windows/startup",
          "POST SET_WIN_CSP_TESTER_STARTUP": "/system/windows/startup/csptester"
      }
  }
  ```
### WinRM
- Do not need to install WINRM packages separately except Windows XP and Vista
  - WINRM setup README, include how to setup WINRM server examples for winrs (windows client tool) and python script
    - https://dops-git.fortinet-us.com/monitor/infrastructure_as_code/-/tree/main/windows
    - WINRM: http://10.65.184.146:5985
    - 用户密码 test 123456


### Appendix
  ![compatibility.png](../images/compatibility.png)