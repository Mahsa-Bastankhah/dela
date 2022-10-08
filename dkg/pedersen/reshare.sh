
#!/bin/sh
for n in 80
do
    for nc in $n  $(($n-1))  $(($n*3/4))
        do
            nc=${nc%.*}
            #echo $nc
            LLVL=warn go test -run TestResharingRecords -timeout 0 -args -nOld=$n -nCommon=$nc
            sleep 2m
        done
done

#echo "nOld,nCommon,nNew,setupTime,resharingTime" >> resharing_records.csv; \