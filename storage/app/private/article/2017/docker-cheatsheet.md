#### Docker cheatsheet

** Basic example **

```yml
# docker-compose.yml
version: '2'

services:
  nginx:
		image: nginx
	  links:
	    - fpm:php-fpm
	  ports:
	    - 80:80
	  volumes_from:
	    - fpm
	fpm:
	  build: .
	  volumes:
	    - .:/var/www/html
	  links:
	    - db:db
	    - redis:redis
	db:
	  image: mysql:5.6
	  volumes:
	    - /var/lib/mysql
	  environment:
	    MYSQL_ROOT_PASSWORD: password
	    MYSQL_DATABASE: database
	    MYSQL_USER: user
	    MYSQL_PASSWORD: password
	redis:
	  image: redis
```

** Docker compose command **

* [docker-compose ps](https://docs.docker.com/compose/reference/ps/)
Lists containers.

* [docker-compose up](https://docs.docker.com/compose/reference/up/)
Builds, (re)creates, starts, and attaches to containers for a service.

* [docker-compose down](https://docs.docker.com/compose/reference/down/)
Stops containers and removes containers, networks, volumes, and images created by up.

* [docker-compose stop](https://docs.docker.com/compose/reference/stop/)
Stops running containers without removing them. They can be started again with docker-compose start.

* [docker-compose start](https://docs.docker.com/compose/reference/start/)
Starts existing containers for a service.

#### Docker cheatsheet