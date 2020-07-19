# get current dir name
current_dir=${PWD##*/}

# change dir to parent dir
cd ..
#  go build -o bin/ch01.exe jvmgo/ch01

echo "build exe for " $current_dir

go build -o bin/$current_dir.exe jvmgo/$current_dir

echo "build success"

echo "change to bin dir for testing"

cd bin && ls -al