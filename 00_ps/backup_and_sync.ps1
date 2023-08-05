# Нахренах?
# Задумано так подключившись по рдп к пк, делается архив текущей папки 
# После производится rsync проброшенной папки Documents средствами программы FreeFileSync
# Переменные
$work_dir="C:\Users\user\что_бекапим_папка\"
$backup_dir="D:\куда_букапим_папка\"
$remote_dir="\\tsclient\Documents\"
$data_time=$(Get-Date -Format yyyyMMdd_Hm_s)
$name_zip_file="name_file_arhive"


# Архивируем каталог из $work_dir в $backup_dir
Compress-Archive -Path "${work_dir}" -DestinationPath "${backup_dir}${name_zip_file}.${data_time}.zip"


# Если каталог есть производим синхронизацию
# Использую стороннию программу FreeFileSync для синхронизации
# Проверяем наличие подключенной удаленой папки
if(test-path ${remote_dir})
{
 echo "Отлично каталог ${remote_dir} есть"
 echo "Производим синхронизацию"
#BatchRun.ffs.batch - файл создан в FreeFileSync, он и выполняет rsync
 C:\folder\BatchRun.ffs_batch
}


