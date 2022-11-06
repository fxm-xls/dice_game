骰子游戏

1.后端接口文档
    两个接口：验证骰子点数、投掷骰子
2.数据存放数据库
    数据库设计：
        用户信息表：id、user_id、name、pwd、token
        用户游戏记录表：id、game_id、user_id、dice_num、create_time 
        游戏表：id、game_id、status、blue_user_id、blue_dice_num、red_user_id、red_dice_num、win_user_id、creat_time
3.不允许前端作弊骰子点数
    对前端传回来的骰子点数进行校验：
        1.时效性校验
        2.点数校验
        3.加入验证字段
    