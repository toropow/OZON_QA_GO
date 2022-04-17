## Задание
- Нужно прислать результат запуска `make tools-version`
- Команда выполнилась без ошибок и содержит нужные приложения:
    - curl
    - golangci-lint
    - libprotoc
    - Docker
    - docker-compose

## Результат
```
anton@anton-Virtual-Machine:~/GolandProjects/device-api$ make tools-version
curl 7.58.0 (x86_64-pc-linux-gnu) libcurl/7.58.0 OpenSSL/1.1.1 zlib/1.2.11 libidn2/2.0.4 libpsl/0.19.1 (+libidn2/2.0.4) nghttp2/1.30.0 librtmp/2.3
Release-Date: 2018-01-24
Protocols: dict file ftp ftps gopher http https imap imaps ldap ldaps pop3 pop3s rtmp rtsp smb smbs smtp smtps telnet tftp
Features: AsynchDNS IDN IPv6 Largefile GSS-API Kerberos SPNEGO NTLM NTLM_WB SSL libz TLS-SRP HTTP2 UnixSockets HTTPS-proxy PSL
golangci-lint has version 1.45.0 built from 1f4c1ed9 on 2022-03-18T15:08:39Z
libprotoc 3.0.0
Docker version 20.10.14, build a224086
docker-compose version 1.29.2, build 5becea4c
```