package kafkaWriter

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/model/request"
	"github.com/HXSecurity/DongTai-agent-go/utils"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/metadata"
)

const (
	TraceId = iota
	AgentId
	RoutineId
	NextKey
	OnlyKey
)

func WriteMessages(w *kafka.Writer, ctx context.Context, msgs ...kafka.Message) error {
	for idx, msg := range msgs {
		traceId := getTraceId(ctx)
		msgs[idx].Headers = append(msg.Headers, kafka.Header{
			Key:   "dt-traceid",
			Value: []byte(traceId),
		})

		fmt.Println("traceId:", traceId, "msg:", msg.Value)
		v := reflect.ValueOf(msg.Value)
		request.FmtHookPool(request.PoolReq{
			Args:            request.Collect(msg.Value),
			NeedHook:        request.Collect(v.Pointer()),
			Source:          false,
			OriginClassName: "kafka.(*Writer)",
			MethodName:      "WriteMessages",
			ClassName:       "kafka.(*Writer)",
		})
	}

	err := WriteMessagesT(w, ctx, msgs...)
	return err
}

func getTraceId(ctx context.Context) string {
	outmd, _ := metadata.FromIncomingContext(ctx)
	worker, _ := utils.NewWorker(global.AgentId)
	var tranceid string
	if len(outmd.Get("dt-traceid")) > 0 {
		tranceid = outmd.Get("dt-traceid")[0]
	}
	if tranceid == "" {
		tranceid = global.TargetTraceId + "." + strconv.Itoa(global.AgentId) + ".0.1." + strconv.Itoa(int(worker.GetId()))
	} else {
		four := strconv.Itoa(int(worker.GetId()))
		tranceids := strings.Split(tranceid, ".")
		tranceids[AgentId] = strconv.Itoa(global.AgentId)
		num, _ := strconv.Atoi(tranceids[NextKey])
		tranceids[NextKey] = strconv.Itoa(num + 1)
		tranceids[OnlyKey] = four
		newId := ""
		for i := 0; i < len(tranceids); i++ {
			if i == OnlyKey {
				newId += tranceids[i]
			} else {
				newId += tranceids[i] + "."
			}
		}
		tranceid = newId
	}

	return tranceid
}

func WriteMessagesT(w *kafka.Writer, ctx context.Context, msgs ...kafka.Message) error {
	return nil
}
