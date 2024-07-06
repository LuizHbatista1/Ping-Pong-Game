# Ping-Pong Game with Concurrency in Go

This project is a simple Ping-Pong game implemented in Go to help understand the concept of concurrency. It uses goroutines and channels to simulate a game of ping-pong between two players.

## Overview

In this Ping-Pong game, two players (goroutines) hit a ball (message) back and forth using a channel. The game continues for a specified number of hits, and then it stops.

## Prerequisites

Go (version 1.16+ recommended)

A basic understanding of Go programming

## Key Concepts
Goroutines: Lightweight threads managed by the Go runtime. The player function runs as a goroutine.

Channels: Used to communicate between goroutines. The table channel simulates the ping-pong table.

Concurrency: Allows multiple goroutines to run simultaneously, making it possible to simulate real-time interaction between players.

## Customization

You can customize the game by:

Changing the duration of the game by modifying time.Sleep(10 * time.Second) in the main function.

Adjusting the delay between hits by changing time.Sleep(1 * time.Second) in the player function.

Adding more players by starting additional goroutines.

## Conclusion

This Ping-Pong game is a simple yet effective way to understand concurrency in Go. By using goroutines and channels, you can see how concurrent programming can be used to simulate real-world interactions.

Feel free to explore and modify the code to deepen your understanding of Go's concurrency model.
