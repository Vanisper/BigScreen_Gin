-- 创建数据库
CREATE DATABASE IF NOT EXISTS `xxb_bigscreen` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建数据库的用户
-- 用于创建一个名为xxb_bigscreen 的用户，并设置密码为 xxb_bigscreen，可以在任何 IP 地址上登录。
-- 具体来说，这个 SQL 语句包含了以下内容：
    -- CREATE USER：创建一个新用户。
    -- 'xxb_bigscreen'@'%'：指定用户的名称和登录 IP 地址。'xxb_bigscreen' 是用户名，'%' 表示可以从任何 IP 地址登录。
    -- IDENTIFIED WITH mysql_native_password BY 'xxb_bigscreen'：指定用户的密码。mysql_native_password 是密码验证插件，'xxb_bigscreen' 是密码。
    -- PASSWORD EXPIRE NEVER：设置密码永不过期。
CREATE USER 'xxb_bigscreen'@'%' IDENTIFIED WITH mysql_native_password BY 'xxb_bigscreen' PASSWORD EXPIRE NEVER;
-- 授权数据库的用户
GRANT ALL PRIVILEGES ON xxb_bigscreen.* TO 'xxb_bigscreen'@'%';
-- 刷新权限
FLUSH PRIVILEGES;
