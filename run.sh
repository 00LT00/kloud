go build -o kloud
lsof -i:1121 | awk 'NR>1 {print $2}'|xargs kill -2
nohup ./kloud > run.log 2>&1 &
echo "start kloud"