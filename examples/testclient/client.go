package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/CyCoreSystems/audiosocket"
	"github.com/pkg/errors"
	"github.com/gofrs/uuid"
)

const serverAddr = ":9092"
const slinChunkSize = 320 // 8000Hz * 20ms * 2 bytes
const fileName = "test.slin"

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Generate a new UUID for the call ID
	callID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate call ID: %v", err)
	}

	// Send the call ID
	if _, err := conn.Write(audiosocket.IDMessage(callID)); err != nil {
		log.Fatalf("failed to send call ID: %v", err)
	}

	// Load the audio file data
	audioData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read audio file: %v", err)
	}

	// Send audio data
	if err = sendAudio(conn, audioData); err != nil {
		log.Fatalf("failed to send audio: %v", err)
	}

	// Send hangup message
	if _, err := conn.Write(audiosocket.HangupMessage()); err != nil {
		log.Fatalf("failed to send hangup message: %v", err)
	}

	log.Println("Audio sent, hangup message sent, closing connection.")
}

func sendAudio(w io.Writer, data []byte) error {
	var i int

	t := time.NewTicker(20 * time.Millisecond)
	defer t.Stop()

	for range t.C {
		if i >= len(data) {
			return nil
		}

		var chunkLen = slinChunkSize
		if i+slinChunkSize > len(data) {
			chunkLen = len(data) - i
		}
		if _, err := w.Write(audiosocket.SlinMessage(data[i : i+chunkLen])); err != nil {
			return errors.Wrap(err, "failed to write chunk to audiosocket")
		}
		i += chunkLen
	}
	return errors.New("ticker unexpectedly stopped")
}