package kafkaWriter

import (
	"fmt"

	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
	"github.com/segmentio/kafka-go"
)

func init() {
	model.HookMap["kafkaGoWriter"] = new(KafkaWriter)
}

type KafkaWriter struct {
}

func (h *KafkaWriter) Hook() {
	w := &kafka.Writer{}
	err := gohook.HookMethod(w, "WriteMessages", WriteMessages, WriteMessagesT)
	if err != nil {
		fmt.Println(err, "kafkaGoWriter")
	} else {
		fmt.Println("kafkaGoWriter")
	}
}

func (h *KafkaWriter) UnHook() {
	cl := &kafka.Writer{}
	gohook.UnHookMethod(cl, "WriteMessages")
}
