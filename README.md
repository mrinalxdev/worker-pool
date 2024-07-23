# Distributed Task Processing System

This project implements a distributed task processing system using Golang's concurrency features. It demonstrates a worker pool pattern where multiple process jobs concurrently, showcasing Golang's powerful and channels.

## Project Structure

'''bash
worker-pool/
├── main.go
├── worker/
│ └── worker.go
├── job/
│ └── job.go
├── dispatcher/
│ └── dispatcher.go
├── result/
│ └── result.go
└── config/
└── config.go
'''

## Key Features

- Concurrent job processing
- Scalable worker pool
- Configurable number of workers and jobs
- Proper synchronization to prevent race conditions
