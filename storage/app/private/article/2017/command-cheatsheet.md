#### Linux command

**Add user exist to group www-data**

```
cut -d: -f1 /etc/group
groupadd
groupdel
sudo usermod -a -G www-data user_name
id vgite
groups vgite
```

#### Install LAMP

```
//For the case 
sudo apt-get install software-properties-common
sudo add-apt-repository ppa:ondrej/php5-5.6
sudo apt-get update
sudo apt-get upgrade 
```

**Install apache**
```
sudo apt-get install apache2
```

**Check connection**
```
ifconfig eth0 | grep inet | awk '{ print $2 }'
```

**Install mysql + php extensions**
```
sudo apt-get install mysql-server-5.6 libapache2-mod-auth-mysql php5-mysql
sudo apt-get install php-fpm php-cli php-mysql php-mcrypt php-memcached php-redis php-curl php-gd php-bz2 php-mbstring php-xml php-zip php-soap

sudo mysql_install_db
```

**Enable mod rewrite**
```
sudo a2enmod rewrite
```

**Web permission**
```
find ./ -type d -exec sudo chmod 2775 {} \;
find ./ -type f -exec sudo chmod 0664 {} \;
```

#### Alias for Apache

```
Alias /news /var/www/html/onelotto-blog
<Directory /var/www/html/onelotto/public>
        Options Indexes FollowSymLinks Includes ExecCGI
        AllowOverride All
        Require all granted
</Directory>

<Directory "/var/www/html/onelotto-blog">
        Options +Indexes
        AllowOverride All
 </Directory>
```

#### Mysql

**Create new database**

```
CREATE DATABASE database_name DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
```

**Privileges**


```
GRANT SELECT, INSERT, UPDATE, DELETE ON  database_name.* TO 'username'@'localhost' IDENTIFIED BY 'raw_password';

GRANT ALL ON mydb.* TO 'myuser'@'%';

REVOKE ALL PRIVILEGES ON database_name.* FROM 'username'@'localhost';

ALTER USER 'admin'@'%' IDENTIFIED BY 'raw_password';


```

**Mysql dump**

```
mysqldump -u username -p -v olddatabase > olddbdump.sql
```

**Clone database**


```
mysqladmin -u username -p create newdatabase
mysqldump -u username -v olddatabase -p | mysql -u username -p -D newdatabase
```

#### .htaccess

**Hotlink Protection**
```
RewriteEngine On
RewriteCond %{HTTP_REFERER} !^http://(.+\.)?domain\.com/ [NC]
RewriteCond %{HTTP_REFERER} !^$
RewriteRule .*\.(jpe?g|gif|bmp|png)$ http://i.imgur.com/donotsteal.gif [L] 
```

#### Memcached

**Install**

```
sudo apt-get install memcached
sudo apt-get install php5-memcached
//or: sudo yum install php[php-verson]-pecl-memcache
//ex: sudo yum install php55-pecl-memcache

sudo service apache2 restart
```

#### Composer

**Install**
```
curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer
```

#### Xdebug + [XHprof](https://docs.moodle.org/dev/Setting_up_xhprof_on_Moodle) + [Xhgui](https://github.com/phacility/xhprof)

**Xdebug**

```
sudo apt-get install php5-xdebug
```

__php.ini__

```
[xdebug]
zend_extension=xdebug.so
xdebug.remote_enable=1
xdebug.remote_handler=dbgp
xdebug.remote_mode=req
xdebug.remote_host=127.0.0.1
xdebug.remote_port=9000

xdebug.profiler_enable=0
xdebug.profiler_output_name = xdebug.out.%t
xdebug.profiler_output_dir=/tmp/profiler
xdebug.profiler_enable_trigger = 0

```

**Xhprof**

```
sudo apt-get install php5-xhprof
sudo pecl install -f xhprof
sudo apt-get install php5-common graphviz
```


####[Swap](https://www.digitalocean.com/community/tutorials/how-to-add-swap-on-ubuntu-14-04)

One of the easiest way of increasing the responsiveness of your server and guarding against out of memory errors in your applications is to add some swap space. Swap is an area on a hard drive that has been designated as a place where the operating system can temporarily store data that it can no longer hold in RAM

**Check the System for Swap Information**
```
sudo swapon -s
```

**Check free**
```
free -m
```

**Check Available Space on the Hard Drive Partition**
```
df -h --total
```

**Create a Swap File using fallocate programe**
```
sudo fallocate -l 4G /swapfile

//Double check
ls -lh /swapfile
```

**Enabling the Swap File**
```
//Permission
sudo chmod 600 /swapfile

//Verify
ls -lh /swapfile

//Turn on
sudo swapon /swapfile

//Verify
sudo swapon -s
```


#### Random password

```php
php -r "echo substr(str_shuffle('qwertyuiopasdfghjkl1234567890zxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM'), 0, 32);"
```

