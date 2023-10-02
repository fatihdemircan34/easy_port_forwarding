Go İle DMZ'den İç Ağa Port Yönlendirici

Bu proje, DMZ (Demilitarized Zone) içerisinde belirli bir portta gelen bağlantıları, iç ağdaki belirli bir IP adresi ve porta yönlendirmek için tasarlanmıştır. Bu, özellikle DMZ'de bulunan bir sunucunun, iç ağdaki başka bir sunucuya trafik yönlendirmesi gerektiğinde kullanışlıdır.

Özellikler:

Komut satırından hedef IP, hedef port ve dinlenen port bilgilerini alabilme.
TCP bağlantılarını dinleyip, belirtilen hedef IP ve porta yönlendirme.
Hatalarla karşılaşıldığında loglama yapabilme.
Kullanım:
Proje, komut satırından parametre alarak çalıştırılır. Örneğin:

go
Copy code
go run dosya_adi.go -targetIP=192.168.1.10 -targetPort=8080 -listenPort=:9090
Bu komut, DMZ'de 9090 portunu dinleyip, gelen bağlantıları 192.168.1.10 IP adresindeki 8080 portuna yönlendirir.

Not:
Bu kod, temel bir port yönlendirme işlevi sağlar. Gerçek bir üretim ortamında kullanmadan önce güvenlik ve performans için ek optimizasyonlar ve kontrollerin yapılması önerilir.

Bu açıklama, projenin temel işlevselliğini ve nasıl kullanılacağını özetler. İhtiyaçlarınıza göre bu açıklamayı daha da genişletebilir veya özelleştirebilirsiniz.
