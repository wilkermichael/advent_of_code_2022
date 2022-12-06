package processor

import "os"

type ElfPacket struct {
	packet               []byte
	startPacketNumUnique int
	messageNumUnique     int
}

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

func (ep ElfPacket) GetStartSequenceIndex() int {
	return findIndexUniqueCharacterSet(string(ep.packet), ep.startPacketNumUnique)
}

func (ep ElfPacket) GetStartMessageIndex() int {
	return findIndexUniqueCharacterSet(string(ep.packet), ep.messageNumUnique)
}

func findIndexUniqueCharacterSet(s string, numUnique int) int {
	n := 0
	for n < len(s) {
		buf := make(map[string]bool)
		for i := n; i < (n + numUnique); i++ {
			buf[string(s[i])] = true
		}

		if len(buf) == numUnique {
			return n + numUnique
		}

		n++
	}

	return -1
}
