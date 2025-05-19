etz api --exec=scripts/executions.yaml -w 5 -r 5 -d 3s
sleep 5
etz api --exec=scripts/locations.yaml -w 5 -r 5 -d 3s
sleep 5
etz api --exec=scripts/calc.yaml -w=5 -r=5 -d=3s
sleep 5
etz file --exec=scripts/fileup.yaml -w 5 -r 5 -d 3s
sleep 5
rm -f uploads/*