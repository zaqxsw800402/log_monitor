# log monitor
改自網路上的課程 <br>
用來練習goroutine、channel <br>

# 功能
用來查看api每次執行的時間


# 架構圖

![image](https://github.com/zaqxsw800402/log_monitor/blob/master/picture/GinLog.png?raw=true)

## Read
讀取log檔新增的資料

## Process
萃取出需要的資料

## Write
寫進InfluxDB

## Grafana
將InfluxDB的資料圖示畫

## 示意圖
![image](https://github.com/zaqxsw800402/log_monitor/blob/master/picture/ginlog2.png?raw=true)

可以明顯看到redis緩存過期後，第一次的時間會明顯增加