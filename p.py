import requests
import time
import matplotlib.pyplot as plt

def test_api_speed(api_url, num_requests=10000):
    response_times = []
    for _ in range(num_requests):
        start_time = time.time()
        response = requests.get(api_url)
        end_time = time.time()
        response_time = end_time - start_time
        response_times.append(response_time)
        # You can print response.status_code or response.json() if you want to validate the response

    average_time = sum(response_times) / num_requests
    print(f"Average time taken for {num_requests} requests: {average_time:.4f} seconds")

    # Plotting
    plt.figure(figsize=(10, 5))
    plt.plot(range(1, num_requests + 1), response_times, marker='o', linestyle='-', color='b')
    plt.title('Response Times of API Requests')
    plt.xlabel('Request Number')
    plt.ylabel('Time (seconds)')
    plt.xticks(range(1, num_requests + 1))
    plt.grid(True)
    plt.tight_layout()
    plt.show()

# Example usage:
api_url = 'http://localhost:8080/movies'  # Replace with your API endpoint 
test_api_speed(api_url)
