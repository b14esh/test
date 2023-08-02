#!/bin/sh
disk_name=$1
if [ "$1" == "" ]
then
         echo -e "ВНИМАНИЕ !!!"
         echo -e "\t Убедись что это не твойcbcntvysq диск!!!"
         echo -e "Передайте имя диска аргументом при запусе. \nПример:\n\t $0 sdb"
         echo -e "Доступные диски для записи: \n\t\t`ls /dev/sd*`"       
         exit
fi

# 0
# Если аргумент был передан, то пошли дальше  
# Для начала очистим диск
echo "0. Стираем начало диска ${disk_name}"
dd if=/dev/zero of=/dev/${disk_name} bs=8M count=1

# Теперь создадим разметку
# n                     # создать новый раздел
# p                     # основной раздел
# номер                 # номер раздела
# 2048                  # создать раздел с отступом 2MiB
# +64M          # размер раздела, тут будет лежать kernel файл с переменными и dts файл
# +2G                   # этот раздел для ubuntu 
# +2G                   # этот в дальнейшем станет для каскадной файловой системы
#                       # а в этом разделе будет весь остальной размер
# |-2048-64M-|-2G-|-2G-|-all-| # желаемая разметка

# 1
echo "1. Приступаем к разметке диска ${disk_name}"
## Пустые строки ниже сделаны не просто так!!!
## -w, --wipe <mode> wipe signatures (auto, always or never) 
##  -W, --wipe-partitions <mode>  wipe signatures from new partitions (auto, always or never) 
## fdisk /dev/sda <<EOF n p 1 2048 +64M....EOF - not work more


parted -a optimal -s /dev/${disk_name} \
        mklabel gpt \
        mkpart primary ext2 10M 74M \
        mkpart primary ext4 74M 2G \
        -- mkpart primary ext4 2G -1

echo "`fdisk -l /dev/${disk_name}`"

# 2
# Далеее накатываем файловую систему на размеченный диск.
echo "2. Создаем файловую систему ${disk_name}" 
# раздел с kernel должен быть либо fat16 либо ext2
mkfs.ext2 /dev/${disk_name}1 << EOF
y
EOF
mkfs.ext4 /dev/${disk_name}2 << EOF
y
EOF
mkfs.ext4 /dev/${disk_name}3 << EOF
y
EOF
#mkfs.ext4 /dev/${disk_name}4  << EOF
#y
#EOF
