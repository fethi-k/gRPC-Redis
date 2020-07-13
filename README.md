# gRPC-Redis
gRPC istemci ve gRPC sunucu kullanarak Redis veritabanına kayıt işlemi.

# Gerekli Kurulumlar
> [Docker for Windows](https://docs.docker.com/docker-for-windows/install/) indir ve kur.

# Çalıştırmak İçin
Proje klasöründe /src git ve
```
docker-compose build
docker-compose up
```
# Son Adım
Server kurulumu client kurulumundan uzun sürdüğü için server container'ının da ayağa kalkması uzun sürüyor. Bu nedenle client ilk başta bağlantı hatası veriyor. Fakat server kurulumu bittiğinde ekrana yazılan
```
gRPC server 4040 portunu dinliyor...
```
mesajından sonra client container'ı tekrar başlatıp sonuçları görebilirsiniz.

# --Düzeltme--
Yukarıdaki hatayı düzeltmek amacıyla docker-compose içerisindeki client altına 
```
restart: on-failure
```
satırını ekledim. Böylece server tarafı hazır olana dek client yeniden başlatılacak ve her şey düzgün çalıştığında sonuçlar gözükecek.
