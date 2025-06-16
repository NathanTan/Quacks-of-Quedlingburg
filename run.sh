
echo -e "Building client server"
cd ./client && go build &
pwd
pid=$! # Get the process ID of the background task

# Show loading dots while the process is running
while kill -0 $pid 2>/dev/null; do
  echo -n "." # Print a dot without a newline
  sleep 0.5   # Wait for half a second
done

echo -e "Complete!\\n"

echo -e "Building game server"
cd ./game && go build &
pid=$! # Get the process ID of the background task

# Show loading dots while the process is running
while kill -0 $pid 2>/dev/null; do
  echo -n "." # Print a dot without a newline
  sleep 0.5   # Wait for half a second
done

echo -e "Complete!\\n"

sleep 1

echo -e "Starting game server"

touch "./logs/gameServerLogs.txt"
./game/main.exe > ./logs/gameServerLogs.txt &

sleep 5

echo -e "Starting client server"

touch "./logs/clientServerLogs.txt"
./client/main.exe > ./logs/clientServerLogs.txt &
