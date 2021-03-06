# @author   : Yuana, Andhika
# @since    : January, 05 2019
# @github   : andhikayuana

declare -a files=(
    "hello" 
    "return-multiple-val"
    "type-switch"
    "error-handling"
    "go-interfaces"
    "type-casting"
    "fibonanci"
    "factorial"
    "delete-functions"
)

function run() {
    for i in "${files[@]}"
    do
        path="./$i/$i.go"
        echo "run : $path..."
        go run $path
    done
}


# start run here
run