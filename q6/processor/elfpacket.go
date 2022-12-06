package processor

import "os"

// ElfPacket describes the features of a packet sent by the
// Elf equipment
type ElfPacket struct {
	packet               []byte
	startPacketNumUnique int
	messageNumUnique     int
}

// NewElfPacket creates a new elf packet
func NewElfPacket(fileName string, startPacketNumUnique, messagePacketNumUnique int) (ElfPacket, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return ElfPacket{}, err
	}

	return ElfPacket{
		packet:               b,
		startPacketNumUnique: startPacketNumUnique,
		messageNumUnique:     messagePacketNumUnique,
	}, nil
}

// GetStartSequenceIndex returns the start sequence index whose uniqueness factor
// is defined by startPacketNumUnique
func (ep ElfPacket) GetStartSequenceIndex() int {
	return findIndexUniqueCharacterSet(string(ep.packet), ep.startPacketNumUnique)
}

// GetStartMessageIndex returns the message sequence whose uniqueness factor
// is defined by messageNumUnique
func (ep ElfPacket) GetStartMessageIndex() int {
	return findIndexUniqueCharacterSet(string(ep.packet), ep.messageNumUnique)
}

// findIndexUniqueCharacterSet uses a sliding window approach to find a unique
// set of characters
func findIndexUniqueCharacterSet(s string, numUnique int) int {
	n := 0
	for n < len(s) {
		// Use a window of size n : n + numUnique
		// Map entries need to be unique, so if we get numUnique map entries
		// for a particular window, then we know that everything in that window
		// is unique
		buf := make(map[string]bool)
		for i := n; i < (n + numUnique); i++ {
			buf[string(s[i])] = true
		}

		if len(buf) == numUnique {
			// Every value in the window created a new map entry, so return now!
			return n + numUnique
		}

		n++
	}

	return -1
}
