FROM eboraas/apache
MAINTAINER Ed Boraas <ed@boraas.ca>

RUN apt-get update && apt-get -y install php5 php5-mysql && apt-get clean && rm -rf 
/var/lib/apt/lists/*

EXPOSE 80
EXPOSE 443

CMD ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]
