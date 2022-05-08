# Лабораторная 1 по предмету Hyperledger Fabric
## Запуск сети, установление чейнкода
Для запуска сети и установки чейнкода нужно запустить скрипт [run.sh](test-network/run.sh) из папки test-network. Он запустит сеть с 1 организацией с 2 пирами и с одним ордерером. Так же эта команда создат канал population в котором установит чейнкод population. 
```
./run.sh
```
## Расположение чейнкода 
Чейнкод располагается [тут](population/chaincode-go)
## Расположение приложения
Приложение располагается [тут](population/application-go)
## Запуск приложения
Для запуска приложения нужно запустить команду
```
go run population.go
```
После чего будет запущено консольное приложение. Для помощи введите команду help. Ввод параметров онеобходимо вводить 1 на 1 строку.

### Пример работы 
```
============ application-golang starts ============
Initialization
============ Populating wallet ============
 [fabsdk/core] 2022/05/08 07:59:09 UTC - cryptosuite.GetDefault -> INFO No default cryptosuite found, using default SW implementation
Initialization finished
Write commands separately
Print help for help inforamtion
insert
Address:
Moscow,1
City:
Moscow 
Id:
1
Name:
Ivan
Status:
Unknown 
Surname:
Ivanov
TelephoneNumber:
88005353535
Enter command
read 
Id:
1
{"Address":"Moscow,1","City":"Moscow","Id":"1","Name":"Ivan","Status":"Unknown","Surname":"Ivanov","TelephoneNumber":"88005353535"}
Enter command
exit
```
