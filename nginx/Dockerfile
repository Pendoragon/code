FROM ubuntu

RUN apt-get update && apt-get install -y --no-install-recommends \
        wget \
        geoip-bin \
        geoip-database \
        build-essential \
        libgd2-dev \
        libreadline-dev \
        libncurses5-dev \
        libpcre3-dev \
        libssl-dev \
        libxslt1-dev \
        perl \
        git \
        ca-certificates

# RUN cd /usr/share/GeoIP && \
#     wget -c http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz && \
#     gunzip GeoLiteCity.dat.gz && \
#     cp /etc/GeoIP.conf.default /etc/GeoIP.conf &&\
#     rm GeoLiteCity.dat.gz

RUN mkdir /root/.ssh && \
    wget http://nginx.org/download/nginx-1.9.14.tar.gz && \
    tar xvf nginx-1.9.14.tar.gz && \
    git clone https://bitbucket.org/nginx-goodies/nginx-sticky-module-ng.git && \
    cd nginx-1.9.14 && \
    ./configure \
    --prefix=/etc/nginx \
    --conf-path=/etc/nginx/nginx.conf \
    --sbin-path=/usr/sbin/nginx \
    --pid-path=/var/log/nginx/nginx.pid \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --with-http_geoip_module \
    --with-file-aio \
    --with-http_ssl_module \
    --with-http_spdy_module \
    --with-http_geoip_module \
    --with-http_gunzip_module \
    --with-http_gzip_static_module \
    --with-http_image_filter_module \
    --with-http_realip_module \
    --with-http_sub_module \
    --with-pcre-jit \
    --with-http_stub_status_module \
    --add-module=/nginx-sticky-module-ng && \
    make && \
    make install

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
	&& ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
