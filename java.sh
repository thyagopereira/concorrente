cd $1
count=0

javac WordCount.java
for d in *;
do
	result[$count]=$(java WordCount ./$d &);
	((count+=1));
done

sum=0
for i in ${resultado[@]}
do
	((sum+=$i))
done

echo $sum
