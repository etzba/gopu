etz api --exec=scripts/distance.yaml -w=5 -r=5 -d=3s
sleep 2
etz api --exec=scripts/calc.yaml -w=5 -r=5 -d=3s
sleep 2
etz file --exec=scripts/fileup.yaml -w 5 -r 5 -d 3s
sleep 2
rm -f uploads/*