FROM prom/alertmanager:v0.20.0

COPY . /etc/alertmanager/

CMD [ "--config.file=/etc/alertmanager/alertmanager.yml", \
      "--web.route-prefix=/alertmanager" ]
