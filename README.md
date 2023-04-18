# Doubletapp Prometheus+Grafana

### Запуск

1. Создать файл .env
    ```
    $ make env
    ```

    `TELEGRAM_BOT_API_KEY` - API токен Телеграма
    `PORT` - порт Nginx, проксирующий запросы к Prometheus и Alertmanager
    `GF_PORT` - порт Grafana
    
2. Запустить приложение
    ```
    $ make
    ```

3. Импортировать дашборды в Grafana

    Дашборды находятся в ```grafana/dashboards```
    В веб-интерфейсе Grafana: ```Create -> Import -> Upload JSON file ```

### Конфигурация

- prometheus/prometheus.yml - Настройка сбора метрик, опций отправки алертов Alertmanger
- prometheus/alerts.yml - Алерты
- alertmanager/alertmanager.yml - Обработка алертов, переотправка вебхукам
