import socket
import time
import threading

# Configuration
HOST = 'localhost'
PORT = 6379
TOTAL_REQUESTS = 10000
NUM_THREADS = 10

# Function to send a RESP command
def send_command(sock, command):
    sock.sendall(command.encode())
    response = sock.recv(1024)
    return response

# Function to measure latency
def measure_latency():
    latencies = []
    for _ in range(TOTAL_REQUESTS):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((HOST, PORT))
        start_time = time.time()
        send_command(sock, "*1\r\n$4\r\nPING\r\n")
        end_time = time.time()
        latency = (end_time - start_time) * 1000  # Convert to milliseconds
        latencies.append(latency)
        sock.close()
    average_latency = sum(latencies) / len(latencies)
    print(f"Average latency: {average_latency:.2f} ms")

# Function to measure throughput
def measure_throughput():
    def worker():
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((HOST, PORT))
        for _ in range(TOTAL_REQUESTS // NUM_THREADS):
            send_command(sock, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")
        sock.close()

    threads = []
    start_time = time.time()
    for _ in range(NUM_THREADS):
        thread = threading.Thread(target=worker)
        thread.start()
        threads.append(thread)
    for thread in threads:
        thread.join()
    end_time = time.time()
    total_time = end_time - start_time
    throughput = TOTAL_REQUESTS / total_time
    print(f"Throughput: {throughput:.2f} operations/second")

# Run the measurements
measure_latency()
measure_throughput()
