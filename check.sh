inputs=("supplied" "n-ties" "all-ties")
for input in "${inputs[@]}"
do
    echo 'program reports for '$input' input...\n'
    go run . './sample-inputs/'$input'-input.txt'
    echo '\nexpected output for '$input' \n'
    cat './sample-inputs/'$input'-output.txt'
    echo '\n'
done