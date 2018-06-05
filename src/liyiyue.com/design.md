# 设计思路

## reddit做法
- 1000 * 1000板子
  - redis - string-bit
    - 每个颜色4位（16色），顺序排列
- 记录
  - Cassandra
    - (x,y):{'timestamp':epochms,'author':user_name,'color':color}

## 我的做法
- 数据库：
  - redis
- 每条记录：
  - sorted-set，实时塞
  - 结构：时间戳，x-y，色
- 当前图像：不存，启动时加载，停止时