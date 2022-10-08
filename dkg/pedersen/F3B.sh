
#!/bin/sh
for n in 64
do
    LLVL=error go test -run TestF3Brecords -timeout 0 -args -n=$n
    sleep 2m
done




#echo "nOld,nCommon,nNew,setupTime,resharingTime" >> resharing_records.csv; \