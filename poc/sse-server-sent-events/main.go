package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type StreamHandler struct {
	redisClient *redis.Client
}

func main() {
	rdb := createRedisClient(6379)
	stream := &StreamHandler{
		redisClient: rdb,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /stream/{symbol}", stream.handle)
	mux.Handle("GET /", http.FileServer(http.Dir("./static")))
	log.Printf("Server running on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", mux))
	defer rdb.Close()
}

func (s *StreamHandler) handle(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "sse not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	stockSymb := r.PathValue("symbol")
	sub := s.redisClient.Subscribe(r.Context(), stockSymb)
	defer sub.Close()

	go updateStockPrice(stockSymb, r.Context(), s.redisClient)

	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-sub.Channel():
			fmt.Fprintf(w, "data: %v\n\n", msg.Payload)
			f.Flush()
		}
	}
}

func createRedisClient(port int) *redis.Client {
	connStr := fmt.Sprintf("localhost:%v", port)
	rdb := redis.NewClient(&redis.Options{
		Addr: connStr,
	})

	status := rdb.Ping(context.Background()).Err()
	if status != nil {
		log.Fatalf("Error connecting to Redis on port: %v. Error message: %v", port, status.Error())
	}

	log.Printf("Redis connected on port %v\n", port)

	return rdb
}

func updateStockPrice(stockSymbol string, ctx context.Context, redisClient *redis.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := redisClient.Publish(ctx, stockSymbol, fmt.Sprintf("%.2f", getStockPrice(stockSymbol))).Err()
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	}
}

func getStockPrice(symbol string) float32 {
	symbolToDefaultPrice := map[string]float32{
		"petr4": 42,
		"vale3": 60,
	}

	defaultPrice := symbolToDefaultPrice[symbol]
	newPrice := (1 + float32(rand.Intn(30))/100) * defaultPrice

	return newPrice
}
