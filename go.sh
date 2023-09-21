cd $1

count=0
for d in *;
do
        resultado[$count]=$(go run  ../go/word_count_changed.go ./$d &);
        ((count+=1));
done

sum=0
for i in ${resultado[@]}
do
        ((sum+=$i))
done
echo $sum

