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
dd if=/dev/zero of=/dev/${disk_name} bs=80M count=1



# 1 
echo "1. Приступаем к разметке диска ${disk_name}"

parted -a optimal -s /dev/${disk_name} \
        mklabel msdos \
        -- mkpart primary ext4 4096s -1 \
        set 1 boot on

echo "`fdisk -l /dev/${disk_name}`"

# 2
# Далеее накатываем файловую систему на размеченный диск.
echo "2. Создаем файловую систему ${disk_name}" 
mkfs.ext4 /dev/${disk_name}1 << EOF
y
EOF

