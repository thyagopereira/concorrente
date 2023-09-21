
cd $1
count=0
for d in *;
do
	resultado[$count]=$(python3 ../python/word_count.py ./$d &);
        ((count+=1));
done

sum=0
for i in ${resultado[@]}
do
	((sum+=$i))
done
	
echo $sum

