 `  go build frwd ` 


Go İle DMZ'den İç Ağa Port Yönlendirici

Bu proje, DMZ (Demilitarized Zone) içerisinde belirli bir portta gelen bağlantıları, iç ağdaki belirli bir IP adresi ve porta yönlendirmek için tasarlanmıştır. Bu, özellikle DMZ'de bulunan bir sunucunun, iç ağdaki başka bir sunucuya trafik yönlendirmesi gerektiğinde kullanışlıdır.

Özellikler:

Komut satırından hedef IP, hedef port ve dinlenen port bilgilerini alabilme.
TCP bağlantılarını dinleyip, belirtilen hedef IP ve porta yönlendirme.
Hatalarla karşılaşıldığında loglama yapabilme.
Kullanım:
Proje, komut satırından parametre alarak çalıştırılır. Örneğin:

 

` go run frwd.go -targetIP=192.168.1.10 -targetPort=8080 -listenPort=:9090` 

Bu komut, DMZ'de 9090 portunu dinleyip, gelen bağlantıları 192.168.1.10 IP adresindeki 8080 portuna yönlendirir.

Not:
Bu kod, temel bir port yönlendirme işlevi sağlar. Gerçek bir üretim ortamında kullanmadan önce güvenlik ve performans için ek optimizasyonlar ve kontrollerin yapılması önerilir.

Bu açıklama, projenin temel işlevselliğini ve nasıl kullanılacağını özetler. İhtiyaçlarınıza göre bu açıklamayı daha da genişletebilir veya özelleştirebilirsiniz.

------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Go-based DMZ to Internal Network Port Forwarder

This project is designed to forward connections coming on a specific port within the DMZ (Demilitarized Zone) to a specified IP address and port in the internal network. This is particularly useful when a server within the DMZ needs to forward traffic to another server inside the internal network.

Features:

Ability to take target IP, target port, and listening port information from the command line.
Listen to TCP connections and forward them to the specified target IP and port.
Logging capabilities when encountering errors.
Usage:
The project is run by taking parameters from the command line. For example:

 
 `   go run frwd.go -targetIP=192.168.1.10 -targetPort=8080 -listenPort=:9090` 
This command listens on port 9090 in the DMZ and forwards incoming connections to port 8080 on the IP address 192.168.1.10.

Note:
This code provides basic port forwarding functionality. It is recommended to perform additional optimizations and checks for security and performance before using it in a real production environment.

This description summarizes the basic functionality of the project and how to use it. You can further expand or customize this description based on your needs.
