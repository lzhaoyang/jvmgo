# get current dir name
current_dir=${PWD##*/}

# change dir to parent dir
cd ..
#  go build -o bin/ch01.exe jvmgo/ch01

echo "build exe for " "$current_dir"

# because of unix like platform don't judge binary file by extension name
# so you can compile program to .exe on unix platform, and run as normal
go build -o bin/"$current_dir".exe jvmgo/"$current_dir"

echo "build success"

echo "change to bin dir for testing"

cd bin && ls -al